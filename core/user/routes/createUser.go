package userRoutes

import (
	"github.com/arnaugomez/buscopartida/core/common/view"
	userView "github.com/arnaugomez/buscopartida/core/user/view"
	"github.com/arnaugomez/buscopartida/ctx"
	"github.com/gin-gonic/gin"
	"net/http"
)


func createUser(c *gin.Context, ctx *ctx.Ctx) {
	// Parse body with Credentials
	body := &userView.Credentials{}
	if err := c.BindJSON(body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(view.ErrorMissingBody, err))
		return
	}
	credentials := userView.ToCredentialsDomain(body)

	// Check if user already exists
	if ctx.UserRepo.UserAlreadyExists(credentials) {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(userView.ErrorUserExists, nil))
		return
	}

	// Create the User
	user, err := ctx.UserRepo.CreateUser(credentials)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view.ErrorToJson(view.ServerError, err))
	}

	c.JSON(http.StatusOK, user)
}
