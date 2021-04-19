package app

import (
	"github.com/carlosstrand/manystagings/core/service"
)

func (a *App) setupService() {
	a.Service = service.NewService(service.Options{
		DB:           a.DB,
		Linker:       a.Linker,
		Orchestrator: a.Orchestrator,
	})
}
