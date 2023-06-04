package main

import (
	"User_endpoint/controllers"
	"User_endpoint/initializers"
	"User_endpoint/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDataBase()
}

func main() {
	r := gin.Default()

	r.POST("/users/register", controllers.Signup)
	r.POST("/users/login", controllers.Login)
	r.GET("/users/login", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
