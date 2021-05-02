package kubernetes

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/carlosstrand/manystagings/core/orchestrator"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var ErrCouldNotLoadKubeconfig = errors.New("could not load kubeconfig")

type Kubernetes struct {
	clientset *kubernetes.Clientset
	logger    *logrus.Logger
	settings  map[string]interface{}
}

type Options struct {
	LogLevel logrus.Level
}

type NamespaceResourceMap struct {
	Pods     map[string]v1.Pod
	Services map[string]v1.Service
}

func loadConfigYaml() (string, error) {
	kubeconfigBase64 := os.Getenv("KUBERNETES_KUBECONFIG_BASE64")
	if kubeconfigBase64 != "" {
		decoded, err := base64.StdEncoding.DecodeString(kubeconfigBase64)
		if err != nil {
			return "", ErrCouldNotLoadKubeconfig
		}
		return string(decoded), nil
	}
	file, err := ioutil.ReadFile(path.Join(userHomeDir(), ".kube", "/config"))
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func encodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func NewKubernetesProvider(opts Options) *Kubernetes {
	configYaml, err := loadConfigYaml()
	if err != nil {
		panic(err)
	}
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(configYaml))
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(opts.LogLevel)
	return &Kubernetes{
		clientset: clientset,
		logger:    logger,
		settings: map[string]interface{}{
			"KUBECONFIG_BASE_64": encodeBase64(configYaml),
		},
	}
}

func (k *Kubernetes) Provider() string {
	return "kubernetes"
}

func (k *Kubernetes) Settings() map[string]interface{} {
	return k.settings
}

func (k *Kubernetes) CreateNamespace(ctx context.Context, namespace string) error {
	_, err := k.clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
	if k8serrors.IsNotFound(err) {
		k.logger.Debugf("namespace '%s' doesn't exist. Creating...", namespace)
		nsSpec := &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}
		_, err = k.clientset.CoreV1().Namespaces().Create(ctx, nsSpec, metav1.CreateOptions{})
		return err
	}
	k.logger.Debugf("namespace '%s' already exists, nothing to do.", namespace)
	return nil
}

func (k *Kubernetes) checkDeploymentExists(ctx context.Context, deployment *orchestrator.Deployment) (exists bool, err error) {
	deploymentsClient := k.clientset.AppsV1().Deployments(deployment.Namespace)
	_, err = deploymentsClient.Get(ctx, deployment.Name, metav1.GetOptions{})
	if err == nil {
		return true, nil
	} else if k8serrors.IsNotFound(err) {
		return false, nil
	}
	return false, err
}

func (k *Kubernetes) checkServiceExists(ctx context.Context, deployment *orchestrator.Deployment) (exists bool, err error) {
	serviceClient := k.clientset.CoreV1().Services(deployment.Namespace)
	_, err = serviceClient.Get(ctx, deployment.Name, metav1.GetOptions{})
	if err == nil {
		return true, nil
	} else if k8serrors.IsNotFound(err) {
		return false, nil
	}
	return false, err
}

func (k *Kubernetes) createService(ctx context.Context, deployment *orchestrator.Deployment) error {
	serviceClient := k.clientset.CoreV1().Services(deployment.Namespace)
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
			Labels: map[string]string{
				"run": deployment.Name,
			},
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"run": deployment.Name,
			},
			Ports: []apiv1.ServicePort{
				{
					Name:       "http",
					Port:       deployment.Port,
					TargetPort: intstr.FromInt(int(deployment.ContainerPort)),
				},
			},
		},
	}
	_, err := serviceClient.Create(ctx, service, metav1.CreateOptions{})
	return err
}

func (k *Kubernetes) deleteService(ctx context.Context, deployment *orchestrator.Deployment) error {
	serviceClient := k.clientset.CoreV1().Services(deployment.Namespace)
	return serviceClient.Delete(ctx, deployment.Name, metav1.DeleteOptions{})
}

func envMapToK8sEnvVar(envMap map[string]string) []apiv1.EnvVar {
	vars := make([]apiv1.EnvVar, 0)
	for key, value := range envMap {
		vars = append(vars, apiv1.EnvVar{
			Name:  key,
			Value: value,
		})
	}
	return vars
}

func (k *Kubernetes) CreateDeployment(ctx context.Context, deployment *orchestrator.Deployment, recreate bool) error {
	exists, err := k.checkDeploymentExists(ctx, deployment)
	if err != nil {
		return err
	}
	if exists && recreate {
		k.logger.Debugf("deployment '%s' already exists. Deleting current deployment...", deployment.Name)
		err := k.DeleteDeployment(ctx, deployment)
		if err != nil {
			return err
		}
	}
	replicas := int32(1)
	k.logger.Debugf("Creating '%s' deployment...", deployment.Name)
	k8sDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deployment.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"run": deployment.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"run": deployment.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  deployment.Name,
							Image: deployment.DockerImage.ToString(),
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
							Env: envMapToK8sEnvVar(deployment.Env),
						},
					},
				},
			},
		},
	}
	if !exists || recreate {
		deploymentsClient := k.clientset.AppsV1().Deployments(deployment.Namespace)
		_, err = deploymentsClient.Create(ctx, k8sDeployment, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	exists, err = k.checkServiceExists(ctx, deployment)
	if exists && recreate {
		k.logger.Debugf("Service '%s' already exists and recreate is enabled. Deleting current service...", deployment.Name)
		err := k.deleteService(ctx, deployment)
		if err != nil {
			return err
		}
	}
	if !exists || recreate {
		k.logger.Debugf("Creating service '%s'...", deployment.Name)
		return k.createService(ctx, deployment)
	}
	return nil
}

func (k *Kubernetes) DeleteDeployment(ctx context.Context, deployment *orchestrator.Deployment) error {
	deploymentsClient := k.clientset.AppsV1().Deployments(deployment.Namespace)
	return deploymentsClient.Delete(ctx, deployment.Name, metav1.DeleteOptions{})
}

// apiVersion: networking.k8s.io/v1beta1
// kind: Ingress
// metadata:
//   annotations:
//     kubernetes.io/ingress.class: "nginx"
//     cert-manager.io/cluster-issuer: "letsencrypt-prod"
//   name: manystagings
// spec:
//   rules:
//     - host: ms.carlosstrand.com
//       http:
//         paths:
//           - backend:
//               serviceName: manystagings
//               servicePort: 80
//             path: /
//   # This section is only required if TLS is to be enabled for the Ingress
//   tls:
//       - hosts:
//           - ms.carlosstrand.com
//         secretName: manystagings-tls-secret

func (k *Kubernetes) CreatePublicURL(ctx context.Context, deployment *orchestrator.Deployment, opts orchestrator.PublicURLOptions) (string, error) {
	host := fmt.Sprintf("%s.%s", opts.Subdomain, opts.Host)
	k.logger.Debugf("Creating public url '%s' (host=%s)...", deployment.Name, host)
	ingClient := k.clientset.ExtensionsV1beta1().Ingresses(deployment.Namespace)
	_, err := ingClient.Create(ctx, &v1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
			Annotations: map[string]string{
				"kubernetes.io/ingress.class":    "nginx",
				"cert-manager.io/cluster-issuer": "letsencrypt-prod",
			},
		},
		Spec: v1beta1.IngressSpec{
			Rules: []v1beta1.IngressRule{
				{
					Host: host,
					IngressRuleValue: v1beta1.IngressRuleValue{
						HTTP: &v1beta1.HTTPIngressRuleValue{
							Paths: []v1beta1.HTTPIngressPath{
								{
									Path: "/",
									Backend: v1beta1.IngressBackend{
										ServiceName: deployment.Name,
										ServicePort: intstr.IntOrString{IntVal: deployment.Port},
									},
								},
							},
						},
					},
				},
			},
			TLS: []v1beta1.IngressTLS{
				{
					Hosts: []string{
						host,
					},
					SecretName: fmt.Sprintf("%s-tls", deployment.Name),
				},
			},
		},
	}, metav1.CreateOptions{})
	url := fmt.Sprintf("https://%s", host)
	return url, err
}

func (k *Kubernetes) DeletePublicURL(ctx context.Context, deployment *orchestrator.Deployment) error {
	ingClient := k.clientset.ExtensionsV1beta1().Ingresses(deployment.Namespace)
	return ingClient.Delete(ctx, deployment.Name, metav1.DeleteOptions{})
}

func statusPhaseToOrchestratorStatus(phase string) string {
	p := strings.ToUpper(phase)
	switch p {
	case "PENDING", "RUNNING", "FAILED":
		return p
	}
	return "UNKNOWN"
}

func (k *Kubernetes) getNamespaceResourcesMap(namespace string) (*NamespaceResourceMap, error) {
	rm := NamespaceResourceMap{
		Pods:     make(map[string]v1.Pod),
		Services: make(map[string]v1.Service),
	}
	pods, err := k.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, pod := range pods.Items {
		rm.Pods[pod.ObjectMeta.Labels["run"]] = pod
	}
	services, err := k.clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, s := range services.Items {
		rm.Services[s.Name] = s
	}
	return &rm, nil
}

func envVarListTopMap(vars []v1.EnvVar) map[string]string {
	varMap := make(map[string]string)
	for _, v := range vars {
		varMap[v.Name] = v.Value
	}
	return varMap
}

func deploymentFromResourceMap(rm *NamespaceResourceMap, appName string) *orchestrator.Deployment {
	pod := rm.Pods[appName]
	svc := rm.Services[appName]
	if len(pod.Spec.Containers) == 0 || len(svc.Spec.Ports) == 0 {
		return nil
	}
	svcPort := svc.Spec.Ports[0]
	container := pod.Spec.Containers[0]
	dockerImage := orchestrator.NewDeploymentDockerImageFromString(container.Image)
	return &orchestrator.Deployment{
		Name:          appName,
		Namespace:     pod.Namespace,
		DockerImage:   *dockerImage,
		Port:          svcPort.Port,
		ContainerPort: svcPort.TargetPort.IntVal,
		Env:           envVarListTopMap(container.Env),
	}
}

func (k *Kubernetes) DeploymentStatuses(ctx context.Context, namespace string) ([]orchestrator.DeploymentStatus, error) {
	rm, err := k.getNamespaceResourcesMap(namespace)
	if err != nil {
		return nil, err
	}
	statuses := make([]orchestrator.DeploymentStatus, 0)
	for appName := range rm.Services {
		status := orchestrator.DeploymentStatus{
			Status:     statusPhaseToOrchestratorStatus(string(rm.Pods[appName].Status.Phase)),
			Deployment: deploymentFromResourceMap(rm, appName),
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}
