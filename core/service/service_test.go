package service_test

import (
	"context"
	"testing"

	"github.com/carlosstrand/manystagings/app"
	"github.com/carlosstrand/manystagings/consts"
	"github.com/carlosstrand/manystagings/core/orchestrator"
	"github.com/carlosstrand/manystagings/core/orchestrator/providers/orchestratormock"
	"github.com/carlosstrand/manystagings/core/service"
	"github.com/carlosstrand/manystagings/models"
	"github.com/carlosstrand/manystagings/seeds"
	"github.com/go-zepto/zepto/plugins/linker/filter"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewTestApp(orchestrator orchestrator.Orchestrator) *app.App {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	if err := seeds.DropAll(db); err != nil {
		panic(err)
	}
	if err := app.AutoMigrateDB(db); err != nil {
		panic(err)
	}
	if err := seeds.RunSeeds(db, seeds.DEFAULT); err != nil {
		panic(err)
	}
	a := app.NewApp(app.Options{
		DB:           db,
		Orchestrator: orchestrator,
	})
	a.Init()
	return a
}

func getEnvironmentByNamespace(t *testing.T, a *app.App, namespace string) *models.Environment {

	var envs []models.Environment
	a.Linker.RepositoryDecoder("Environment").Find(context.TODO(), nil, &envs)

	var env *models.Environment
	err := a.Linker.RepositoryDecoder("Environment").FindOne(context.TODO(), &filter.Filter{
		Where: &map[string]interface{}{
			"namespace": map[string]interface{}{
				"eq": namespace,
			},
		},
	}, &env)
	require.NoError(t, err)
	return env
}

func TestGetInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := orchestratormock.NewMockOrchestrator(ctrl)

	kubeSettings := map[string]interface{}{
		"KUECONFIG": "some-kube-config-file",
	}

	m.EXPECT().
		Provider().
		Return("kubernetes")

	m.EXPECT().
		Settings().
		Return(kubeSettings)

	app := NewTestApp(m)

	info := app.Service.GetInfo()

	assert.Equal(t, &service.Info{
		Version:              consts.VERSION,
		OrchestratorProvider: "kubernetes",
		OrchestratorSettings: kubeSettings,
	}, info)
}

func TestApplyDeployment_All(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := orchestratormock.NewMockOrchestrator(ctrl)

	m.EXPECT().
		CreateNamespace(gomock.Any(), "qa").
		Return(nil)

	m.EXPECT().
		CreateDeployment(gomock.Any(), &orchestrator.Deployment{
			Name:      "node-api",
			Namespace: "qa",
			DockerImage: orchestrator.DeploymentDockerImage{
				Name: "dockercloud/hello-world",
				Tag:  "latest",
			},
			Env: map[string]string{
				"NODE_ENV": "development",
				"SITE_URL": "http://staging.mysite.com",
			},
		}).
		Return(nil)

	m.EXPECT().
		CreateDeployment(gomock.Any(), &orchestrator.Deployment{
			Name:      "db",
			Namespace: "qa",
			DockerImage: orchestrator.DeploymentDockerImage{
				Name: "postgres",
				Tag:  "13.2-alpine",
			},
			Env: map[string]string{
				"POSTGRES_USER":     "staging",
				"POSTGRES_PASSWORD": "staging",
			},
		}).
		Return(nil)

	app := NewTestApp(m)

	environment := getEnvironmentByNamespace(t, app, "qa")
	err := app.Service.EnvironmentApplyDeployment(context.Background(), environment, []string{})
	assert.NoError(t, err)
}

func TestApplyDeployment_Single_App(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := orchestratormock.NewMockOrchestrator(ctrl)

	m.EXPECT().
		CreateNamespace(gomock.Any(), "qa").
		Return(nil)

	m.EXPECT().
		CreateDeployment(gomock.Any(), &orchestrator.Deployment{
			Name:      "node-api",
			Namespace: "qa",
			DockerImage: orchestrator.DeploymentDockerImage{
				Name: "dockercloud/hello-world",
				Tag:  "latest",
			},
			Env: map[string]string{
				"NODE_ENV": "development",
				"SITE_URL": "http://staging.mysite.com",
			},
		}).
		Return(nil)

	app := NewTestApp(m)

	environment := getEnvironmentByNamespace(t, app, "qa")
	err := app.Service.EnvironmentApplyDeployment(context.Background(), environment, []string{"node-api"})
	assert.NoError(t, err)
}

func TestApplyDeployment_Apps_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := orchestratormock.NewMockOrchestrator(ctrl)

	app := NewTestApp(m)

	environment := getEnvironmentByNamespace(t, app, "qa")
	err := app.Service.EnvironmentApplyDeployment(context.Background(), environment, []string{"node-api", "not-found-1", "not-found-2"})
	assert.Error(t, err, "apps not found: not-found-1, not-found-2")
}
