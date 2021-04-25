package app

import (
	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/plugins/auth"
	"github.com/go-zepto/zepto/plugins/auth/datasources/gorm"
)

func (app *App) setupAuth() {
	app.Zepto.AddPlugin(auth.NewAuthTokenPlugin(auth.AuthTokenOptions{
		Datasource: gorm.NewGormAuthDatasoruce(gorm.GormAuthDatasourceOptions{
			UserModel: &models.User{},
			DB:        app.DB,
		}),
	}))

	// Block any API request from unauthorized user
	// app.apiRouter.Use(func(next web.RouteHandler) web.RouteHandler {
	// 	return func(ctx web.Context) error {
	// 		a := auth.InstanceFromCtx(ctx)
	// 		loggedUserPID := a.LoggedPIDFromCtx(ctx)
	// 		if loggedUserPID == nil {
	// 			ctx.SetStatus(401)
	// 			ctx.RenderJson(map[string]string{
	// 				"error": "unauthorized",
	// 			})
	// 		}
	// 		return next(ctx)
	// 	}
	// })
}
