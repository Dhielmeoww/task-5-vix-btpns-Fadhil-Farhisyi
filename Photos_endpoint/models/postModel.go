package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string
	Caption  string
	PhotoUrl string
	USC      int32
}
