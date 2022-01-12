package userRoutes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func getUserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"user": "Mantarelo",
		"id": id,
	})
}
