package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"trade/dao"
	"trade/middleware"
	"trade/models"
)

func ValidateUser(creds models.User) (string, error) {
	var user models.User
	result := dao.DB.Where("username = ?", creds.Username).First(&user)
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

type CronService struct{}

func (sm *CronService) SixSecondTask() {
	fmt.Println("6 secs runs")
	log.Println("6 secs runs")
}
