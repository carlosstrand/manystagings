package controllers

import (
	"github.com/carlosstrand/manystagings/core/service"
	"github.com/go-zepto/zepto/plugins/linker"
	"github.com/go-zepto/zepto/web"
)

type Controllers struct {
	svc    *service.Service
	linker *linker.Linker
}

type Options struct {
	Service *service.Service
	Linker  *linker.Linker
}

func NewControllers(opts Options) *Controllers {
	return &Controllers{
		svc:    opts.Service,
		linker: opts.Linker,
	}
}

func renderError(ctx web.Context, err error, statusCode int) error {
	ctx.SetStatus(statusCode)
	ctx.RenderJson(map[string]string{
		"error": err.Error(),
	})
	return nil
}

func renderAccepted(ctx web.Context) error {
	ctx.SetStatus(202)
	ctx.RenderJson(map[string]bool{
		"status": true,
	})
	return nil
}
