package routes

import (
	"github.com/arnaugomez/buscopartida/ctx"
	"github.com/gin-gonic/gin"
)

type routeHandler func(r *gin.Context, ctx *ctx.Ctx)

func Register(g *gin.RouterGroup, path string, handler routeHandler, ctx *ctx.Ctx) {
	g.POST(path, func(context *gin.Context) {
		handler(context, ctx)
	})
}