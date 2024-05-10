package services

import (
	"AssetsTrade/middleware"
	"AssetsTrade/models"
	"errors"
	"gorm.io/gorm"
)

func Login(creds models.User) (string, error) {
	var user models.User
	result := middleware.DB.Where("username = ?", creds.Username).First(&user)

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
	result := middleware.DB.Where("username = ?", creds.Username).First(&user)

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
