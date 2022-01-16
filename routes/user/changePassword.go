package userRoutes

import (
	"errors"
	"github.com/arnaugomez/buscopartida/core/common/utils/hashing"
	"github.com/arnaugomez/buscopartida/ctx"
	"github.com/arnaugomez/buscopartida/routes/common/view"
	"github.com/arnaugomez/buscopartida/routes/user/view"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func updatePassword(c *gin.Context, ctx *ctx.Ctx, userId uint) {
	body := &userView.ChangePasswordPayload{}
	if err := c.BindJSON(body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(view.ErrorMissingBody, err))
		return
	}


	usr, err := ctx.UserRepo.GetUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(userView.ErrorUserDoesNotExist, err))
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.ErrorToJson(view.ServerError, err))
		return
	}

	if err := hashing.CheckHash(body.PreviousPassword, usr.PasswordHash); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(userView.WrongPassword, err))
		return
	}


	if err := ctx.UserRepo.UpdatePassword(userId, body.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.ErrorToJson(view.ServerError, err))
		return
	}

	c.Status(http.StatusOK)
}
