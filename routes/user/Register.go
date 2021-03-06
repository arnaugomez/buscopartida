package userRoutes

import (
	"github.com/arnaugomez/buscopartida/core/common/utils/routes"
	"github.com/arnaugomez/buscopartida/ctx"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, ctx *ctx.Ctx) {
	g := r.Group("/user")
	routes.Register(g, "/create", createUser, ctx)
	routes.Register(g, "/login", login, ctx)
	routes.Protected(g, "/login/jwt", jwtLogin, ctx)
	routes.Protected(g, "/delete", deleteUser, ctx)
	routes.Protected(g, "/password/update", updatePassword, ctx)
	routes.Protected(g, "/my-profile/get", getMyProfile, ctx)
}
