package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"trade/middleware"
	"trade/models"
)

func Login(creds models.User) (string, error) {
	var user models.User
	result := middleware.DB.Where("user_name = ?", creds.Username).First(&user)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 如果有其他数据库错误，返回错误
			return "", result.Error
		} else {
			password, err := hashPassword(creds.Password)
			if err != nil {
				return "", err
			}
			creds.Password = password
			createUserErr := CreateUser(&creds)
			if createUserErr != nil {
				return "", createUserErr
			}
		}
	}
	if !CheckPassword(user.Password, creds.Password) {
		return "", errors.New("invalid credentials")
	}
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
	if user.Password != creds.Password {
		return "", errors.New("invalid credentials")
	}
	if !CheckPassword(user.Password, creds.Password) {
		return "", errors.New("invalid credentials")
	}
	token, err := middleware.GenerateToken(creds.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (sm *CronService) FiveSecondTask() {
	fmt.Println("5 secs runs")
	log.Println("5 secs runs")
}
