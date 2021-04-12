package kubernetescli

import (
	"context"
	"errors"

	"github.com/carlosstrand/manystagings/core/orchestrator"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ErrCouldNotFindRunningPod = errors.New("could not find running pod")

func (k *KubernetesCLIProvider) getPodByDeployment(ctx context.Context, deployment *orchestrator.Deployment) (string, error) {
	podList, err := k.clientset.CoreV1().Pods(deployment.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	for _, pod := range podList.Items {
		validPod := pod.Status.Phase == v1.PodPending || pod.Status.Phase == v1.PodRunning
		if pod.Labels["run"] == deployment.Name && validPod {
			return pod.Name, nil
		}
	}
	return "", ErrCouldNotFindRunningPod
}
