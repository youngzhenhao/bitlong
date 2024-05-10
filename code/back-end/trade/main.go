package main

import (
	"AssetsTrade/config"
	"AssetsTrade/handlers"
	"AssetsTrade/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	// Initialize the database connection
	middleware.DbConnect()

	// If you need to migrate the database table structure
	// models.Migrate()

	// Initialize the Redis connection
	middleware.RedisConnect()

	r := gin.Default()

	// Login routing
	r.POST("/login", handlers.LoginHandler)

	// Refresh the route for the token
	r.POST("/refresh", handlers.RefreshTokenHandler)

	// A routing group that requires authentication
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/userinfo", handlers.UserInfoHandler)
	}

	// Read the port number from the configuration file
	port := loadConfig.Server.Port
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
