package services

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"time"
	"trade/config"
	"trade/dao"
	"trade/task"
)

var (
	db *sql.DB
)

type Job struct {
	Name           string
	CronExpression string
	FunctionName   string
	PackagePath    string
}

type CronService struct{}

func init() {
	var err error
	// Initialize MySQL Database
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
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadJobs() ([]Job, error) {
	var jobs []Job
	rows, err := db.Query("SELECT name, cron_expression, function_name, package FROM scheduled_tasks")
	if err != nil {
		log.Fatal("Failed to load tasks:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.Name, &job.CronExpression, &job.FunctionName, &job.PackagePath); err != nil {
			log.Fatal("Failed to read task data:", err)
			return nil, err
		}
		fn := getFunction(job.PackagePath, job.FunctionName)
		if fn.IsValid() {
			taskFunc := func() {
				fn.Call(nil)
			}
			task.RegisterTask(job.Name, taskFunc)

			jobs = append(jobs, job)
		} else {
			log.Printf("Function %s not found in package %s", job.FunctionName, job.PackagePath)
		}
	}
	return jobs, nil
}

func ExecuteWithLock(taskName string) {
	lockKey := "lock:" + taskName
	flag := dao.AcquireLock(lockKey, 1, 10*time.Minute)
	if !flag {
		fmt.Println("Failed to acquire lock for", taskName)
		return
	}
	defer dao.ReleaseLock(lockKey)
	task.ExecuteTask(taskName)
}
func getFunction(pkgName, funcName string) reflect.Value {
	switch pkgName {
	case "services":
		// Assuming there is a struct that encapsulates the methods
		manager := CronService{} // You need to define this struct
		return reflect.ValueOf(&manager).MethodByName(funcName)
	default:
		return reflect.Value{}
	}
}
