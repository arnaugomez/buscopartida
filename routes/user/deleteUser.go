package userRoutes

import (
	"github.com/arnaugomez/buscopartida/ctx"
	view2 "github.com/arnaugomez/buscopartida/routes/common/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func deleteUser(c *gin.Context, ctx *ctx.Ctx, userId uint) {
	err := ctx.UserRepo.DeleteUserById(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view2.ErrorToJson(view2.ServerError, err))
		return
	}
	c.Status(http.StatusOK)
}
