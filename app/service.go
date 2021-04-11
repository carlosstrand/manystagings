package app

import (
	"github.com/carlosstrand/manystagings/core/orchestrator/providers/kubernetes"
	"github.com/carlosstrand/manystagings/core/service"
	"github.com/sirupsen/logrus"
)

func (a *App) setupService() {
	a.Service = service.NewService(service.Options{
		DB:     a.DB,
		Linker: a.Linker,
		Orchestrator: kubernetes.NewKubernetesProvider(kubernetes.Options{
			LogLevel: logrus.DebugLevel,
		}),
	})
}
