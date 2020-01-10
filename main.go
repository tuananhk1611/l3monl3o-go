package main

import (
	controller "./controller"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	client := r.Group("/api")
	{
		client.GET("/user/:id", controller.Get)
		client.POST("/user/create", controller.Create)
		client.PATCH("/user/update/:id", controller.Update)
		client.DELETE("/user/delete/:id", controller.Delete)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":1611") // listen and serve at port http 1611
}
