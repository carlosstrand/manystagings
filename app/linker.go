package app

import (
	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/plugins/linker"
	gormds "github.com/go-zepto/zepto/plugins/linker/datasource/linkergorm"
)

func (a *App) setupLinker() {
	a.Linker = linker.NewLinker(a.apiRouter)

	a.Linker.AddResource(linker.Resource{
		Name:       "Environment",
		Datasource: gormds.NewGormDatasource(a.DB, &models.Environment{}),
	})

	a.Linker.AddResource(linker.Resource{
		Name:       "Application",
		Datasource: gormds.NewGormDatasource(a.DB, &models.Application{}),
	})

	a.Linker.AddResource(linker.Resource{
		Name:       "ApplicationEnvVar",
		Datasource: gormds.NewGormDatasource(a.DB, &models.ApplicationEnvVar{}),
	})

	a.Linker.AddResource(linker.Resource{
		Name:       "Config",
		Datasource: gormds.NewGormDatasource(a.DB, &models.Config{}),
	})

	a.Zepto.AddPlugin(linker.NewLinkerPlugin(linker.Options{
		Linker: a.Linker,
	}))
}
