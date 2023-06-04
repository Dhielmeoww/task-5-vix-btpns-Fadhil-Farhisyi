package main

import (
	"task-5-vix-btpns-Fadhil-Farhisyi/initializers"
	"task-5-vix-btpns-Fadhil-Farhisyi/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
