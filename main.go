package main

import (
	"github.com/gin-gonic/gin"
	controller "./controller"
)

func main() {
	r := gin.Default()
	r.GET("/user/:id", controller.Read)
	r.Run(":1611") // listen and serve at port http 1611
}
