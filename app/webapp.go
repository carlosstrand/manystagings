package app

import "github.com/carlosstrand/manystagings/controllers"

func (a *App) setupWebapp() {
	a.Zepto.Get("/", controllers.WebappIndex)
}
