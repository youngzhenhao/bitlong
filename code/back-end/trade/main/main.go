package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"trade/config"
	"trade/dao"
	"trade/routers"
	"trade/task"
	"trade/utils"
)

func main() {
	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		utils.LogError("failed to load config: ", err)
	}
	if loadConfig.Routers.Login {
		dao.InitMysql()
		dao.RedisConnect()
	}
	err = dao.Migrate()
	jobs, err := task.LoadJobs()
	if err != nil {
		utils.LogError("", err)
	}
	c := cron.New(cron.WithSeconds())
	for _, job := range jobs {
		_, err := c.AddFunc(job.CronExpression, func() {
			task.ExecuteWithLock(job.Name)
		})
		if err != nil {
			utils.LogError("Error scheduling job "+job.Name, err)
		}
	}
	c.Start()
	if err != nil {
		utils.LogError("AutoMigrate error", err)
		return
	}
	r := routers.SetupRouter()
	bind := loadConfig.GinConfig.Bind
	port := loadConfig.GinConfig.Port
	if port == "" {
		port = "8080"
	}
	err = r.Run(fmt.Sprintf("%s:%s", bind, port))
	if err != nil {
		return
	}
}
