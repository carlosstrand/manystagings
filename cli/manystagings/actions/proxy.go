package actions

import (
	"context"

	"github.com/carlosstrand/manystagings/core/orchestrator"
	"github.com/urfave/cli"
)

func (a *Actions) ProxyDeployment(c *cli.Context) error {
	appList, err := a.client.GetEnvironmentApplications(context.TODO(), a.config.EnvironmentID)
	if err != nil {
		return err
	}
	appName := c.Args().First()
	for _, app := range appList.Data {
		if app.Name == appName {
			a.orchestratorCLI.ProxyDeployment(context.TODO(), &orchestrator.Deployment{
				Name:      app.Name,
				Namespace: app.Environment.Namespace,
			})
		}
	}
	return ErrAppNotFound
}
