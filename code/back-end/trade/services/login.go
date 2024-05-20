package services

import (
	"errors"
	"gorm.io/gorm"
	"trade/dao"
	"trade/middleware"
	"trade/models"
)

func Login(creds models.User) (string, error) {
	var user models.User
	result := dao.DB.Where("username = ?", creds.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("invalid credentials")
	}
	// Verify your password
	if user.Password != creds.Password {
		return "", errors.New("invalid credentials")
	}
	// Generate an initial token
	token, err := middleware.GenerateToken(creds.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateUserAndGenerateToken(creds models.User) (string, error) {
	var user models.User
	result := dao.DB.Where("username = ?", creds.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("invalid credentials")
	}
	// Verify your password
	if user.Password != creds.Password {
		return "", errors.New("invalid credentials")
	}
	// Generate tokens for subsequent requests
	token, err := middleware.GenerateToken(creds.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
