package routes

import (
	"errors"
	jwtView "github.com/arnaugomez/buscopartida/core/jwt/view"
	"github.com/arnaugomez/buscopartida/ctx"
	"github.com/arnaugomez/buscopartida/routes/common/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type routeHandler func(r *gin.Context, ctx *ctx.Ctx)

func Register(g *gin.RouterGroup, path string, handler routeHandler, ctx *ctx.Ctx) {
	g.POST(path, func(context *gin.Context) {
		handler(context, ctx)
	})
}

type protectedRouteHandler func(r *gin.Context, ctx *ctx.Ctx, userId uint)

func Protected(g *gin.RouterGroup, path string, handler protectedRouteHandler, ctx *ctx.Ctx) {
	g.POST(path, func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		jwtToken := strings.Replace(authHeader, "Bearer ", "", 1)
		if len(jwtToken) == 0 {
			context.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(jwtView.ErrorNoJwt, errors.New("no jwt provided")))
			return
		}

		userId, err := ctx.JwtRepo.GetUserIdByJWT(jwtToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(jwtView.ErrorInvalidJwt, err))
			return
		}

		handler(context, ctx, userId)
	})
}
