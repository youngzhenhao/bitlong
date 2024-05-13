package middleware

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trade/config"
	"trade/utils"
)

var (
	DB *gorm.DB
)

// TODO: Rename function
func DbConnect() {
	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		loadConfig.GormConfig.Mysql.Username,
		loadConfig.GormConfig.Mysql.Password,
		loadConfig.GormConfig.Mysql.Host,
		loadConfig.GormConfig.Mysql.Port,
		loadConfig.GormConfig.Mysql.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.LogError("failed to connect database", err)
		return
	}
	DB = db
}
