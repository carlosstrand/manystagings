package kubernetescli

import (
	"context"
	"io"
	"os"

	"github.com/carlosstrand/manystagings/core/orchestrator"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func (k *KubernetesCLIProvider) ExecDeployment(ctx context.Context, deployment *orchestrator.Deployment, command []string) error {
	podName, err := k.getPodByDeployment(ctx, deployment)
	if err != nil {
		return ErrCouldNotFindRunningPod
	}
	return k.ExecCmd(ExecCmdOptions{
		Namespace: deployment.Namespace,
		PodName:   podName,
		Command:   command,
		Stdin:     os.Stdin,
		Stdout:    os.Stdout,
		Stderr:    os.Stderr,
	})
}

type ExecCmdOptions struct {
	Namespace string
	PodName   string
	Command   []string
	Stdin     io.Reader
	Stdout    io.Writer
	Stderr    io.Writer
}

func (k *KubernetesCLIProvider) ExecCmd(opts ExecCmdOptions) error {
	req := k.clientset.CoreV1().RESTClient().Post().Resource("pods").Name(opts.PodName).
		Namespace(opts.Namespace).SubResource("exec")
	option := &v1.PodExecOptions{
		Command: opts.Command,
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}
	if opts.Stdin == nil {
		option.Stdin = false
	}
	req.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(k.restconfig, "POST", req.URL())
	if err != nil {
		return err
	}
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  opts.Stdin,
		Stdout: opts.Stdout,
		Stderr: opts.Stderr,
	})
	if err != nil {
		return err
	}

	return nil
}
