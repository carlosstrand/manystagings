package actions

import (
	"context"

	"github.com/carlosstrand/manystagings/core/orchestrator"
)

func (a *Actions) ExecDeployment(appName string, command []string) error {
	appList, err := a.client.GetEnvironmentApplications(context.TODO(), a.config.EnvironmentID)
	if err != nil {
		return err
	}
	for _, app := range appList.Data {
		if app.Name == appName {
			return a.orchestratorCLI.ExecDeployment(context.TODO(), &orchestrator.Deployment{
				Name:      app.Name,
				Namespace: app.Environment.Namespace,
			}, command)
		}
	}
	return ErrAppNotFound
}
