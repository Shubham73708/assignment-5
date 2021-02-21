package main

import (
	"fmt"

	"fruit-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")

	router := SetupRouter()
	router.Run(":8081")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/fruit", controllers.Create)
		v1.GET("/fruit/:id", controllers.GetFruit)
		v1.GET("/fruits", controllers.GetAllFruit)
		v1.DELETE("/fruit/:id", controllers.DeleteFruit)
		v1.GET("/check", controllers.HealthCheck)
		v1.POST("/signup/", controllers.Signup)
		v1.POST("/signin", controllers.Signin)

	}
	return router
}
