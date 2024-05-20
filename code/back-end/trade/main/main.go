package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"trade/config"
	"trade/dao"
	"trade/routers"
	"trade/task"
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

		jobs, err := task.LoadJobs()
		if err != nil {
			log.Fatal(err)
		}

		c := cron.New(cron.WithSeconds())
		for _, job := range jobs {
			// Schedule each job using cron
			_, err := c.AddFunc(job.CronExpression, func() {
				task.ExecuteWithLock(job.Name)
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
	r := routers.SetupRouter()
	bind := loadConfig.GinConfig.Bind
	port := loadConfig.GinConfig.Port
	if port == "" {
		port = "8080"
	}
	err = r.Run(fmt.Sprintf("%s:%s", bind, port))
}
