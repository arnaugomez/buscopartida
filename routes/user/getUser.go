package userRoutes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// getUser, createUser, updateUser, deleteUser

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "Mantarelo",
	})
}
