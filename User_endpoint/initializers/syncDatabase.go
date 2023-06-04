package initializers

import (
	"User_endpoint/models"
)

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
}
