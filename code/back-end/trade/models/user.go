package models

import (
	"AssetsTrade/middleware"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func Migrate() {
	err := middleware.DB.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
