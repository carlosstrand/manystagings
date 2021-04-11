package controllers

import (
	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/web"
)

func (c *Controllers) EnvironmentApplyDeployment(ctx web.Context) error {
	envID := ctx.Params()["id"]
	var env *models.Environment
	if err := c.linker.RepositoryDecoder("Environment").FindById(ctx, envID, &env); err != nil {
		return renderError(ctx, err, 400)
	}
	if err := c.svc.EnvironmentApplyDeployment(ctx, env); err != nil {
		return renderError(ctx, err, 400)
	}
	return renderAccepted(ctx)
}

func (c *Controllers) EnvironmentGetList(ctx web.Context) error {
	envList, err := c.svc.EnvironmentGetList(ctx)
	if err != nil {
		return renderError(ctx, err, 500)
	}
	return ctx.RenderJson(envList)
}

func (c *Controllers) EnvironmentGetById(ctx web.Context) error {
	envID := ctx.Params()["id"]
	env, err := c.svc.EnvironmentGetById(ctx, envID)
	if err != nil {
		return renderError(ctx, err, 500)
	}
	return ctx.RenderJson(env)
}
