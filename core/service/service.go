package service

import (
	"context"

	"github.com/carlosstrand/manystagings/core/orchestrator"
	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/plugins/linker"
	"github.com/go-zepto/zepto/plugins/linker/filter"
	"github.com/go-zepto/zepto/plugins/linker/filter/include"
	"gorm.io/gorm"
)

type Options struct {
	DB           *gorm.DB
	Linker       *linker.Linker
	Orchestrator orchestrator.Orchestrator
}

type Service struct {
	db           *gorm.DB
	linker       *linker.Linker
	orchestrator orchestrator.Orchestrator
}

func NewService(opts Options) *Service {
	return &Service{
		db:           opts.DB,
		linker:       opts.Linker,
		orchestrator: opts.Orchestrator,
	}
}

func appEnvVarsToMap(appEnvVars []models.ApplicationEnvVar) map[string]string {
	env := make(map[string]string)
	for _, v := range appEnvVars {
		env[v.Key] = v.Value
	}
	return env
}

func (s *Service) ApplyEnvironmentDeployment(ctx context.Context, environment *models.Environment) error {
	var apps []models.Application
	s.linker.RepositoryDecoder("Application").Find(ctx, &filter.Filter{
		Where: &map[string]interface{}{
			"environment_id": map[string]interface{}{
				"eq": environment.ID,
			},
		},
		Include: []include.Include{
			{
				Relation: "ApplicationEnvVar",
			},
		},
	}, &apps)

	for _, app := range apps {
		s.orchestrator.CreateDeployment(ctx, &orchestrator.Deployment{
			Name: app.Name,
			DockerImage: orchestrator.DeploymentDockerImage{
				Name: app.DockerImageName,
				Tag:  app.DockerImageTag,
			},
			Namespace: environment.Namespace,
			Env:       appEnvVarsToMap(app.ApplicationEnvVars),
		})
	}

	return nil
}
