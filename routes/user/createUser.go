package userRoutes

import (
	"fmt"
	"github.com/arnaugomez/buscopartida/core/user"
	"github.com/arnaugomez/buscopartida/ctx"
	view2 "github.com/arnaugomez/buscopartida/routes/common/view"
	userView2 "github.com/arnaugomez/buscopartida/routes/user/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createUser(c *gin.Context, ctx *ctx.Ctx) {
	// Parse body with Credentials
	body := &userView2.Credentials{}
	if err := c.BindJSON(body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view2.ErrorToJson(view2.ErrorMissingBody, err))
		return
	}
	credentials := userView2.ToCredentialsDomain(body)

	// Check if user already exists
	if ctx.UserRepo.UserAlreadyExists(credentials) {
		c.AbortWithStatusJSON(http.StatusBadRequest, view2.ErrorToJson(userView2.ErrorUserExists, nil))
		return
	}

	// Create the User
	usr, err := ctx.UserRepo.CreateUser(credentials)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view2.ErrorToJson(view2.ServerError, err))
	}

	// Create the profile
	profile, err := ctx.UserRepo.CreateProfile(&user.Profile{UserID: usr.ID, Slug: usr.Name})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, view2.ErrorToJson(view2.ServerError, err))
	}
	fmt.Println(profile)

	c.JSON(http.StatusOK, userView2.ToUserView(usr))
}
