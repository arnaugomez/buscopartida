package userRoutes

import (
	"errors"
	"github.com/arnaugomez/buscopartida/core/common/utils/hashing"
	"github.com/arnaugomez/buscopartida/core/common/view"
	"github.com/arnaugomez/buscopartida/core/user/view"
	"github.com/arnaugomez/buscopartida/ctx"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func login(c *gin.Context, ctx *ctx.Ctx) {
	body := &userView.Credentials{}
	if err := c.BindJSON(body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(view.ErrorMissingBody, err))
		return
	}
	credentials := userView.ToCredentialsDomain(body)

	usr, err := ctx.UserRepo.GetUserByEmail(credentials.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(userView.ErrorUserDoesNotExist, err))
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.ErrorToJson(view.ServerError, err))
		return
	}
	if err := hashing.CheckHash(credentials.Password, usr.PasswordHash); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(userView.WrongPassword, err))
		return
	}

	token, err := ctx.JwtRepo.CreateJWT(usr.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.ErrorToJson(view.ServerError, err))
		return
	}

	c.JSON(http.StatusOK, userView.ToUserWithJwtView(token, usr))
}

func jwtLogin(c *gin.Context, ctx *ctx.Ctx, userId uint) {

	usr, err := ctx.UserRepo.GetUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(userView.ErrorUserDoesNotExist, err))
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.ErrorToJson(view.ServerError, err))
		return
	}

	token, err := ctx.JwtRepo.CreateJWT(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.ErrorToJson(view.ServerError, err))
		return
	}

	c.JSON(http.StatusOK, userView.ToUserWithJwtView(token, usr))
}
