package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trade/dao"
	"trade/task"
	"trade/utils"

	"github.com/robfig/cron/v3"
	"trade/config"
	"trade/middleware"
	"trade/routers"
)

func main() {
	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// Initialize the database connection
	if err := middleware.InitMysql(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	if err := middleware.RedisConnect(); err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}
	if !checkConfig() {
		return
	}
	if config.GetLoadConfig().IsAutoMigrate {
		err = dao.Migrate()
	}
	if err != nil {
		utils.LogError("AutoMigrate error", err)
		return
	}
	// Setup cron jobs
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
	defer c.Stop() // Ensure cron scheduler is stopped on shutdown
	// Setup HTTP server
	r := routers.SetupRouter()
	/**
	for _, routeInfo := range r.Routes() {
		fmt.Println(routeInfo.Method, routeInfo.Path)
	}
	*/
	bind := loadConfig.GinConfig.Bind
	port := loadConfig.GinConfig.Port
	if port == "" {
		port = "8080"
	}
	err = r.Run(fmt.Sprintf("%s:%s", bind, port))
	if err != nil {
		return
	}
	// Create a channel to listen to interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	// Start HTTP server in a goroutine
	go func() {
		if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	// Block until a signal is received
	sig := <-signalChan
	log.Printf("Received signal: %s", sig)
	// Gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Close Redis connection
	if err := middleware.Client.Close(); err != nil {
		log.Printf("Failed to close Redis connection: %v", err)
	} else {
		log.Println("Redis connection closed successfully.")
	}
	// Close database connection
	if db, err := middleware.DB.DB(); err == nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
	// Perform any other shutdown tasks here
	log.Println("Shutting down the server...")
	os.Exit(0)
}

// check config
func checkConfig() bool {
	cfg := config.GetConfig()
	if cfg.ApiConfig.CustodyAccount.MacaroonDir == "" {
		fmt.Println("Custody account MacaroonDir is not set")
		return false
	}
	return true
}
