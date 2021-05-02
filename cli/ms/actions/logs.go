package actions

import (
	"context"

	"github.com/carlosstrand/manystagings/cli/ms/orchestratorcli"
	"github.com/carlosstrand/manystagings/core/orchestrator"
)

func (a *Actions) LogsDeployment(appName string, opts orchestratorcli.LogsOptions) error {
	appList, err := a.client.GetEnvironmentApplications(context.TODO(), a.config.EnvironmentID)
	if err != nil {
		return err
	}
	for _, app := range appList.Data {
		if app.Name == appName {
			return a.orchestratorCLI.LogsDeployment(context.TODO(), &orchestrator.Deployment{
				Name:      app.Name,
				Namespace: app.Environment.Namespace,
			}, opts)
		}
	}
	return ErrAppNotFound
}
