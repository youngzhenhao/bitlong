package services

import (
	"trade/middleware"
	"trade/models"
)

func Migrate() {
	err := middleware.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
