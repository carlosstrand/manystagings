package app

import (
	"github.com/carlosstrand/manystagings/controllers"
)

func (a *App) setupControllers() {
	ctrls := controllers.NewControllers(controllers.Options{
		Service: a.Service,
		Linker:  a.Linker,
	})

	// Additional CLI routes
	a.apiRouter.Get("/info", ctrls.Info)
	a.apiRouter.Post("/environments/{id}/apply-deployment", ctrls.EnvironmentApplyDeployment)
	a.apiRouter.Post("/environments/{id}/delete-deployment", ctrls.EnvironmentDeleteDeployment)
	a.apiRouter.Get("/environments/{id}/status", ctrls.EnvironmentAppStatuses)
}
