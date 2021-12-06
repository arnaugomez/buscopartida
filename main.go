package main

import (
  "fmt"
  "github.com/arnaugomez/buscopartida/di"
  userRoutes "github.com/arnaugomez/buscopartida/routes/user"
  "github.com/gin-gonic/gin"
)

func main() {
  di := di.Setup()
  fmt.Println(di)

  r := gin.Default()

  userRoutes.Setup(r)


  r.Run(":3000")
}
