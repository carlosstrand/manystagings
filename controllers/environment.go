package controllers

import (
	"encoding/json"

	"github.com/carlosstrand/manystagings/models"
	"github.com/go-zepto/zepto/web"
)

type EnvironmentApplyDeploymentRequest struct {
	Apps     []string `json:"apps"`
	Recreate bool     `json:"recreate"`
}

type EnvironmentDeleteDeploymentRequest struct {
	Apps []string `json:"apps"`
}

func (c *Controllers) EnvironmentApplyDeployment(ctx web.Context) error {
	envID := ctx.Params()["id"]
	var req EnvironmentApplyDeploymentRequest
	err := json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		renderError(ctx, err, 500)
	}
	var env *models.Environment
	if err := c.linker.RepositoryDecoder("Environment").FindById(ctx, envID, &env); err != nil {
		return renderError(ctx, err, 400)
	}
	if err := c.svc.EnvironmentApplyDeployment(ctx, env, req.Apps, req.Recreate); err != nil {
		return renderError(ctx, err, 400)
	}
	return renderAccepted(ctx)
}

func (c *Controllers) EnvironmentDeleteDeployment(ctx web.Context) error {
	envID := ctx.Params()["id"]
	var req EnvironmentDeleteDeploymentRequest
	err := json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		renderError(ctx, err, 500)
	}
	var env *models.Environment
	if err := c.linker.RepositoryDecoder("Environment").FindById(ctx, envID, &env); err != nil {
		return renderError(ctx, err, 400)
	}
	if err := c.svc.EnvironmentDeleteDeployment(ctx, env, req.Apps); err != nil {
		return renderError(ctx, err, 400)
	}
	return renderAccepted(ctx)
}

func (c *Controllers) EnvironmentAppStatuses(ctx web.Context) error {
	envID := ctx.Params()["id"]
	var env *models.Environment
	if err := c.linker.RepositoryDecoder("Environment").FindById(ctx, envID, &env); err != nil {
		return renderError(ctx, err, 400)
	}
	statuses, err := c.svc.EnvironmentAppStatuses(ctx, env)
	if err != nil {
		return renderError(ctx, err, 400)
	}
	return json.NewEncoder(ctx.Response()).Encode(statuses)
}
