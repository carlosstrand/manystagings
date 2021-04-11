package controllers

import "github.com/go-zepto/zepto/web"

func (c *Controllers) Info(ctx web.Context) error {
	info := c.svc.GetInfo()
	return ctx.RenderJson(info)
}
