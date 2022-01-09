package main

import (
	"fmt"
	ctxDomain "github.com/arnaugomez/buscopartida/ctx/domain"
)

func main() {
  ctx := ctxDomain.Setup()
  fmt.Println("Contexto")
  fmt.Println(ctx)

  // r := gin.Default()

  // userRoutes.Setup(r)


  // r.Run(":3000")
}
