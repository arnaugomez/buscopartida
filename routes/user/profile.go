package userRoutes

import (
	"errors"
	"github.com/arnaugomez/buscopartida/ctx"
	view2 "github.com/arnaugomez/buscopartida/routes/common/view"
	userView "github.com/arnaugomez/buscopartida/routes/user/view"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func getMyProfile(c *gin.Context, ctx *ctx.Ctx, userId uint) {
	p, err := ctx.UserRepo.GetProfileByUserId(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusBadRequest, view2.ErrorToJson(userView.ErrorProfileDoesNotExist, err))
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view2.ErrorToJson(view2.ServerError, err))
		return
	}
	c.JSON(http.StatusOK, userView.ToProfileView(p))
}

