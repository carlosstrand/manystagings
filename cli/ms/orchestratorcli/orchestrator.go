package orchestratorcli

import (
	"context"
	"time"

	"github.com/carlosstrand/manystagings/core/orchestrator"
)

type LogsOptions struct {
	Follow       bool
	Timestamps   bool
	LimitBytes   int64
	Tail         int64
	SinceTime    string
	SinceSeconds time.Duration
}

type OrchestratorCLI interface {
	ProxyDeployment(ctx context.Context, deployment *orchestrator.Deployment) error
	ExecDeployment(ctx context.Context, deployment *orchestrator.Deployment, command []string) error
	LogsDeployment(ctx context.Context, deployment *orchestrator.Deployment, opts LogsOptions) error
}
