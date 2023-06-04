package main

import (
	"task-5-vix-btpns-Fadhil-Farhisyi/controllers"
	"task-5-vix-btpns-Fadhil-Farhisyi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionToDB()
}

func main() {
	r := gin.Default()
	r.POST("/photos", controllers.PostsCreate)
	r.GET("/photos", controllers.PostsIndex)
	r.GET("/photos/:id", controllers.PostsOne)
	r.PUT("/:photoid", controllers.PostsUpdate)
	r.DELETE("/:photoid", controllers.PostsDelete)
	r.Run()
}
