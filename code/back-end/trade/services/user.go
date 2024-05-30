package services

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
func CreateUser(user *models.User) error {
	return middleware.DB.Create(user).Error
}

// ReadUser retrieves a user by ID
func ReadUser(id uint) (*models.User, error) {
	var user models.User
	err := middleware.DB.First(&user, id).Error
	return &user, err
}

// ReadUserByUsername retrieves a user by username
func ReadUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := middleware.DB.Where(&models.User{Username: username}).First(&user).Error
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
	user := models.User{}
	err := middleware.DB.Where("user_name = ?", name).First(&user).Error
	return int(user.ID), err
}

func hashPassword(password string) (string, error) {
	// Passwords are encrypted using the bcrypt algorithm
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(hashedPassword, password string) bool {
	// bcrypt.CompareHashAndPassword Compare the hashed password with the password entered by the user. If there is a match, nil is returned.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
