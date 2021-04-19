package actions

import (
	"context"
)

func (a *Actions) Up(appNames []string) error {
	return a.client.ApplyEnvironmentDeployment(context.TODO(), a.config.EnvironmentID, appNames)
}
