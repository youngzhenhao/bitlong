package main

import (
	"fmt"
	"trade/config"
	"trade/middleware"
	"trade/routers"
	"trade/utils"
)

func main() {
	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	if loadConfig.Routers.Login {
		middleware.DbConnect()
		// If you need to migrate the database table structure
		// models.Migrate()
		// Initialize the Redis connection
		middleware.RedisConnect()
	}
	r := routers.SetupRouter()
	bind := loadConfig.GinConfig.Bind
	port := loadConfig.GinConfig.Port
	addr := fmt.Sprintf("%s:%s", bind, port)
	if port == "" {
		port = "8080"
	}
	err = r.Run(addr)
	if err != nil {
		utils.LogError("", err)
		return
	}
}
