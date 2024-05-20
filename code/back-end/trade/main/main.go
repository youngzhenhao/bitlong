package main

import (
	"fmt"
	"trade/config"
	"trade/dao"
	"trade/models"
	"trade/routers"
	"trade/utils"
)

func main() {
	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	dao.InitMysql()
	dao.RedisConnect()
	err = dao.DB.AutoMigrate(&models.Account{})
	err = dao.DB.AutoMigrate(&models.Balance{})
	err = dao.DB.AutoMigrate(&models.BalanceExt{})
	err = dao.DB.AutoMigrate(&models.Invoice{})
	err = dao.DB.AutoMigrate(&models.User{})
	if err != nil {
		utils.LogError("AutoMigrate error", err)
		return
	}
	r := routers.SetupRouter()
	port := loadConfig.GinConfig.Port
	if port == "" {
		port = "8080"
	}
	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return
	}
}
