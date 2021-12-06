package userRoutes

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	userRoutes := r.Group("/user")
	userRoutes.GET("/", getUser)
	userRoutes.GET("/id/:id", getUserById)
	userRoutes.GET("/json", getJsonBody)
}