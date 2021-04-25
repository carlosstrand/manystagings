package controllers

import "github.com/go-zepto/zepto/web"

func WebappIndex(ctx web.Context) error {
	return ctx.Render("webapp/index")
}
