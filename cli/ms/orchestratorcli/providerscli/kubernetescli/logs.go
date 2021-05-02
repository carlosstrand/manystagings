package kubernetescli

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/carlosstrand/manystagings/cli/ms/orchestratorcli"
	"github.com/carlosstrand/manystagings/core/orchestrator"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ParseRFC3339 parses an RFC3339 date in either RFC3339Nano or RFC3339 format.
func ParseRFC3339(s string, nowFn func() metav1.Time) (metav1.Time, error) {
	if t, timeErr := time.Parse(time.RFC3339Nano, s); timeErr == nil {
		return metav1.Time{Time: t}, nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return metav1.Time{}, err
	}
	return metav1.Time{Time: t}, nil
}

func (k *KubernetesCLIProvider) LogsDeployment(ctx context.Context, deployment *orchestrator.Deployment, opts orchestratorcli.LogsOptions) error {
	podName, err := k.getPodByDeployment(ctx, deployment)
	if err != nil {
		return err
	}
	podLogOptions := corev1.PodLogOptions{
		Follow:     opts.Follow,
		Timestamps: opts.Timestamps,
	}

	if opts.Tail != 0 {
		podLogOptions.TailLines = &opts.Tail
	}

	if len(opts.SinceTime) > 0 {
		t, err := ParseRFC3339(opts.SinceTime, metav1.Now)
		if err != nil {
			return err
		}

		podLogOptions.SinceTime = &t
	}

	if opts.SinceSeconds != 0 {
		// round up to the nearest second
		sec := int64(opts.SinceSeconds.Round(time.Second).Seconds())
		podLogOptions.SinceSeconds = &sec
	}

	if opts.LimitBytes != 0 {
		podLogOptions.LimitBytes = &opts.LimitBytes
	}

	podLogRequest := k.clientset.CoreV1().
		Pods(deployment.Namespace).
		GetLogs(podName, &podLogOptions)
	stream, err := podLogRequest.Stream(context.TODO())
	if err != nil {
		return err
	}
	defer stream.Close()

	for {
		buf := make([]byte, 2000)
		numBytes, err := stream.Read(buf)
		if numBytes == 0 {
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		message := string(buf[:numBytes])
		fmt.Print(message)
		if !opts.Follow {
			return nil
		}
	}
	return nil
}
