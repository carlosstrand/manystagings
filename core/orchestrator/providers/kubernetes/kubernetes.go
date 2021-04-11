package kubernetes

import (
	"context"

	"github.com/carlosstrand/manystagings/core/orchestrator"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Kubernetes struct {
	clientset *kubernetes.Clientset
	logger    *logrus.Logger
}

type Options struct {
	LogLevel logrus.Level
}

func NewKubernetesProvider(opts Options) *Kubernetes {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/carlosstrand/.kube/config")
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
	}
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

func (k *Kubernetes) CreateDeployment(ctx context.Context, deployment *orchestrator.Deployment) error {
	exists, err := k.checkDeploymentExists(ctx, deployment)
	if err != nil {
		return err
	}
	if exists {
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
					"app": deployment.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deployment.Name,
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
						},
					},
				},
			},
		},
	}
	deploymentsClient := k.clientset.AppsV1().Deployments(deployment.Namespace)
	_, err = deploymentsClient.Create(ctx, k8sDeployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (k *Kubernetes) DeleteDeployment(ctx context.Context, deployment *orchestrator.Deployment) error {
	deploymentsClient := k.clientset.AppsV1().Deployments(deployment.Namespace)
	return deploymentsClient.Delete(ctx, deployment.Name, metav1.DeleteOptions{})
}
