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

		}
		middleware.DbConnect()
		middleware.RedisConnect()
		//users, err := services.GenericQuery[models.User](&models.User{}, services.QueryParams{"Username": "John", "Status": "active"})
		//condition := &models.User{Username: "John", Status: 1}

		// 调用GenericQuery来根据条件查询用户
		//users, err := services.GenericQueryByObject[models.User](condition)
		//if err != nil {
		//	fmt.Printf("Error while fetching users: %v\n", err)
		//	return
		//}
		//
		//fmt.Printf("Found %d users\n", len(users))
		//for _, user := range users {
		//	fmt.Println(user)
		//}
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
