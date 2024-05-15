package main

import (
	"fmt"
	"trade/config"
	"trade/middleware"
	"trade/routers"
)

func main() {

	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	if true {
		if loadConfig.Routers.Login {
			// Initialize the database connection
			middleware.DbConnect()
			// If you need to migrate the database table structure
			// models.Migrate()
			// Initialize the Redis connection
			middleware.RedisConnect()
		}
		r := routers.SetupRouter()
		// Read the port number from the configuration file

		port := loadConfig.GinConfig.Port
		if port == "" {
			// Default port number
			port = "8080"
		}
		// Start the server
		err = r.Run(fmt.Sprintf(":%s", port))
		if err != nil {
			return
		}
	}
}
