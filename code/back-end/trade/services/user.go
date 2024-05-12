package services

import (
	"AssetsTrade/middleware"
	"AssetsTrade/models"
)

func Migrate() {
	err := middleware.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
