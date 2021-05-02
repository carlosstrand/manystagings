package actions

import (
	"context"

	"github.com/carlosstrand/manystagings/core/orchestrator"
)

func (a *Actions) ProxyDeployment(appName string, port int32) error {
	appList, err := a.client.GetEnvironmentApplications(context.TODO(), a.config.EnvironmentID)
	if err != nil {
		return err
	}
	for _, app := range appList.Data {
		if app.Name == appName {
			proxyPort := port
			if port == -1 {
				proxyPort = app.Port
			}
			a.orchestratorCLI.ProxyDeployment(context.TODO(), &orchestrator.Deployment{
				Name:          app.Name,
				Namespace:     app.Environment.Namespace,
				Port:          proxyPort,
				ContainerPort: app.ContainerPort,
			})
		}
	}
	return ErrAppNotFound
}
