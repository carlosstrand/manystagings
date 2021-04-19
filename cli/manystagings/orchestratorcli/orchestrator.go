package orchestratorcli

import (
	"context"

	"github.com/carlosstrand/manystagings/core/orchestrator"
)

type OrchestratorCLI interface {
	ProxyDeployment(ctx context.Context, deployment *orchestrator.Deployment) error
	ExecDeployment(ctx context.Context, deployment *orchestrator.Deployment) error
	// ApplyEnvironmentDeployment(ctx context.Context, appNames []string) error
}
