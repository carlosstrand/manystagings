package kubernetescli

import (
	"encoding/base64"
	"errors"

	"github.com/carlosstrand/manystagings/cli/ms/utils/msconfig"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var ErrKubeconfigNofConfigured = errors.New("kubeconfig not configured")

type KubernetesCLIProvider struct {
	clientset  *kubernetes.Clientset
	restconfig *rest.Config
	logger     *logrus.Logger
}

type Options struct {
	Config *msconfig.ManyStagingsConfig
}

func NewKubernetesCLIProvider(opts Options) *KubernetesCLIProvider {
	kubeconfigBase64, exists := opts.Config.OrchestratorSettings["KUBECONFIG_BASE_64"]
	if !exists {
		panic(ErrKubeconfigNofConfigured)
	}
	kb64 := kubeconfigBase64.(string)
	kyaml, err := base64.StdEncoding.DecodeString(kb64)
	if err != nil {
		panic(err)
	}
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kyaml))
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
	loglevel, err := logrus.ParseLevel(opts.Config.LogLevel)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(loglevel)
	return &KubernetesCLIProvider{
		clientset:  clientset,
		logger:     logger,
		restconfig: config,
	}
}
