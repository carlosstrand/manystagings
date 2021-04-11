package app

import (
	"github.com/carlosstrand/manystagings/controllers"
)

func (a *App) setupControllers() {
	ctrls := controllers.NewControllers(controllers.Options{
		Service: a.Service,
		Linker:  a.Linker,
	})

	// Controller CLI Router is
	cliRouter := a.Zepto.Router("/cli/api")
	cliRouter.Get("/environments", ctrls.EnvironmentGetList)
	cliRouter.Get("/environments/{id}", ctrls.EnvironmentGetById)
	cliRouter.Post("/environments/{id}/apply-deployment", ctrls.EnvironmentApplyDeployment)
}
