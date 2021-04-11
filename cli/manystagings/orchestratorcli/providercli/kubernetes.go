package kubernetes

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesCLIProvider struct {
	Name string
}

type Options struct {
	LogLevel logrus.Level
}

func NewKubernetesCLIProvider(opts Options) *KubernetesCLIProvider {
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
