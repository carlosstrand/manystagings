package actions

import (
	"context"

	"github.com/carlosstrand/manystagings/core/orchestrator"
)

func (a *Actions) ProxyDeployment(appName string) error {
	appList, err := a.client.GetEnvironmentApplications(context.TODO(), a.config.EnvironmentID)
	if err != nil {
		return err
	}
	for _, app := range appList.Data {
		if app.Name == appName {
			a.orchestratorCLI.ProxyDeployment(context.TODO(), &orchestrator.Deployment{
				Name:          app.Name,
				Namespace:     app.Environment.Namespace,
				Port:          app.Port,
				ContainerPort: app.ContainerPort,
			})
		}
	}
	return ErrAppNotFound
}
