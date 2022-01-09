package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type userBody struct {
	User string `json:"user"`
	Name string `json:"name"`
}

func getJsonBody(c *gin.Context) {
	body := userBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, body)
}
