package userRoutes

import (
	"github.com/arnaugomez/buscopartida/core/common/utils/routes"
	"github.com/arnaugomez/buscopartida/ctx"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, ctx *ctx.Ctx) {
	g := r.Group("/user")
	routes.Register(g, "/create", createUser, ctx )
}