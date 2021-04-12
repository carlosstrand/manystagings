package orchestratorcli

import (
	"context"

	"github.com/carlosstrand/manystagings/core/orchestrator"
)

type OrchestratorCLI interface {
	ProxyDeployment(ctx context.Context, deployment *orchestrator.Deployment) error
}
