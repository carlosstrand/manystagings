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
