package actions

import (
	"context"
)

func (a *Actions) Kill(appNames []string) error {
	return a.client.DeleteEnvironmentDeployment(context.TODO(), a.config.EnvironmentID, appNames)
}
