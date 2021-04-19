package service

import (
	"context"
	"fmt"
	"strings"

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
	Version              string                 `json:"version"`
	OrchestratorProvider string                 `json:"orchestrator_provider"`
	OrchestratorSettings map[string]interface{} `json:"orchestrator_settings"`
}

type AppStatus struct {
	Application *models.Application `json:"application"`
	Status      string              `json:"status"`
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
		Version:              consts.VERSION,
		OrchestratorProvider: s.orchestrator.Provider(),
		OrchestratorSettings: s.orchestrator.Settings(),
	}
	return info
}

func appsContainsName(apps []models.Application, name string) bool {
	for _, a := range apps {
		if a.Name == name {
			return true
		}
	}
	return false
}

func (s *Service) EnvironmentApplyDeployment(ctx context.Context, environment *models.Environment, appNames []string) error {
	var apps []models.Application
	where := map[string]interface{}{
		"environment_id": map[string]interface{}{
			"eq": environment.ID,
		},
	}
	if len(appNames) > 0 {
		where["name"] = map[string]interface{}{
			"in": appNames,
		}
	}
	s.linker.RepositoryDecoder("Application").Find(ctx, &filter.Filter{
		Where: &where,
		Include: []include.Include{
			{
				Relation: "ApplicationEnvVars",
			},
		},
	}, &apps)

	if len(appNames) > 0 && len(appNames) != len(apps) {
		notFoundApps := make([]string, 0)
		for _, appName := range appNames {
			if !appsContainsName(apps, appName) {
				notFoundApps = append(notFoundApps, appName)
			}
		}
		return fmt.Errorf("apps not found: %s", strings.Join(notFoundApps, ", "))
	}

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
			Port:          app.Port,
			ContainerPort: app.ContainerPort,
			Namespace:     environment.Namespace,
			Env:           appEnvVarsToMap(app.ApplicationEnvVars),
		})
	}

	return nil
}

func statusFromAppName(statuses []orchestrator.DeploymentStatus, appName string) string {
	for _, ds := range statuses {
		if ds.Deployment != nil && ds.Deployment.Name == appName {
			return ds.Status
		}
	}
	return "NOT RUNNING"
}

func (s *Service) EnvironmentAppStatuses(ctx context.Context, environment *models.Environment) ([]AppStatus, error) {
	var apps []*models.Application
	s.linker.RepositoryDecoder("Application").Find(ctx, &filter.Filter{
		Where: &map[string]interface{}{
			"environment_id": map[string]interface{}{
				"eq": environment.ID,
			},
		},
	}, &apps)
	statuses, err := s.orchestrator.DeploymentStatuses(ctx, environment.Namespace)
	if err != nil {
		return nil, err
	}
	appStatuses := make([]AppStatus, 0)
	for _, app := range apps {
		appStatuses = append(appStatuses, AppStatus{
			Application: app,
			Status:      statusFromAppName(statuses, app.Name),
		})
	}
	return appStatuses, nil
}
