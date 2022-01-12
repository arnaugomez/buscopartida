package main

import (
	"fmt"
	userRoutes "github.com/arnaugomez/buscopartida/core/user/routes"
	ctxDomain "github.com/arnaugomez/buscopartida/ctx/domain"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := ctxDomain.CreateCtx()

	r := gin.Default()

	userRoutes.Register(r, ctx)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
