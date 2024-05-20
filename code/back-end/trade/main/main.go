package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"trade/config"
	"trade/dao"
	"trade/routers"
	"trade/services"
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
		dao.InitMysql()
		dao.RedisConnect()
		//err = dao.DB.AutoMigrate(&models.Account{})
		//err = dao.DB.AutoMigrate(&models.Balance{})
		//err = dao.DB.AutoMigrate(&models.BalanceExt{})
		//err = dao.DB.AutoMigrate(&models.Invoice{})
		//err = dao.DB.AutoMigrate(&models.User{})
		//if err != nil {
		//	utils.LogError("AutoMigrate error", err)
		//	return
		//}

		jobs, err := services.LoadJobs()
		if err != nil {
			log.Fatal(err)
		}

		c := cron.New(cron.WithSeconds())
		for _, job := range jobs {
			// Schedule each job using cron
			_, err := c.AddFunc(job.CronExpression, func() {
				services.ExecuteWithLock(job.Name)
			})
			if err != nil {
				log.Printf("Error scheduling job %s: %v\n", job.Name, err)
			}
		}

		c.Start()
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
