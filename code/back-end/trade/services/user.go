package services

import (
	"errors"
	"gorm.io/gorm"
	"trade/dao"
	"trade/middleware"
	"trade/models"
)

//	func Migrate() {
//		err := middleware.DB.AutoMigrate(&models.User{})
//		if err != nil {
//			return
//		}
//	}
func ValidateUser(creds models.User) (string, error) {
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

// CreateUser creates a new user record
func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

// GetUser retrieves a user by ID
func GetUser(id uint) (*models.User, error) {
	var user models.User
	err := dao.DB.First(&user, id).Error
	return &user, err
}

// UpdateUser updates an existing user
func UpdateUser(user *models.User) error {
	return dao.DB.Save(user).Error
}

// DeleteUser soft deletes a user by ID
func DeleteUser(id uint) error {
	var user models.User
	return dao.DB.Delete(&user, id).Error
}
