package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"trade/middleware"
	"trade/models"
)

func ValidateUser(creds models.User) (string, error) {
	var user models.User
	result := middleware.DB.Where("username = ?", creds.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("invalid credentials")
	}
	if user.Password != creds.Password {
		return "", errors.New("invalid credentials")
	}
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

// ReadUser retrieves a user by ID
func ReadUser(id uint) (*models.User, error) {
	var user models.User
	err := middleware.DB.First(&user, id).Error
	return &user, err
}

// UpdateUser updates an existing user
func UpdateUser(user *models.User) error {
	return middleware.DB.Save(user).Error
}

// DeleteUser soft deletes a user by ID
func DeleteUser(id uint) error {
	var user models.User
	return middleware.DB.Delete(&user, id).Error
}

func (sm *CronService) SixSecondTask() {
	fmt.Println("6 secs runs")
	log.Println("6 secs runs")
}

func NameToId(name string) (int, error) {
	user := models.User{Username: name}
	err := middleware.DB.First(&user).Error
	return int(user.ID), err
}
