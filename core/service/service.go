package service

import (
	"context"

	"github.com/carlosstrand/manystagings/consts"
	"github.com/carlosstrand/manystagings/core/orchestrator"
	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/plugins/linker"
	"github.com/go-zepto/zepto/plugins/linker/filter"
	"github.com/go-zepto/zepto/plugins/linker/filter/include"
	"gorm.io/gorm"
)

const MAX_ENVIRONMENTS = 20

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

type EnvironmentList struct {
	Data  []*models.Environment `json:"data"`
	Count int64                 `json:"count"`
}

type Info struct {
	Version      string `json:"version"`
	Orchestrator string `json:"orchestrator"`
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

func (s *Service) GetInfo() *Info {
	info := &Info{
		Version:      consts.VERSION,
		Orchestrator: s.orchestrator.Provider(),
	}
	return info
}

func (s *Service) EnvironmentApplyDeployment(ctx context.Context, environment *models.Environment) error {
	var apps []models.Application
	s.linker.RepositoryDecoder("Application").Find(ctx, &filter.Filter{
		Where: &map[string]interface{}{
			"environment_id": map[string]interface{}{
				"eq": environment.ID,
			},
		},
		Include: []include.Include{
			{
				Relation: "ApplicationEnvVars",
			},
		},
	}, &apps)

	err := s.orchestrator.CreateNamespace(ctx, environment.Namespace)
	if err != nil {
		return err
	}

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
