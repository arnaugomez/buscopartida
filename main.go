package main

import (
	"fmt"
	ctxDomain "github.com/arnaugomez/buscopartida/ctx/domain"
)

func main() {
	ctx := ctxDomain.CreateCtx()
	fmt.Println("Contexto")
	fmt.Println(ctx)

	// r := gin.Default()

	// userRoutes.CreateCtx(r)

	// r.Run(":3000")
}
