package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ErrorToJson(errorType string, err error) gin.H  {
	return gin.H{
		"error": errorType,
		"trace": fmt.Sprint(err),
	}
}