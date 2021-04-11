package app

import (
	"github.com/carlosstrand/manystagings/controllers"
)

func (a *App) setupControllers() {
	ctrls := controllers.NewControllers(controllers.Options{
		Service: a.Service,
		Linker:  a.Linker,
	})

	// Controller Routes
	a.apiRouter.Post("/environments/{environment_id}/apply-deployment", ctrls.EnvironmentApplyDeployment)
}
