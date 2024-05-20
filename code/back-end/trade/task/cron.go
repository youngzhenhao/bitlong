package task

import (
	"database/sql"
	"log"
	"reflect"
	"time"
	"trade/middleware"
	"trade/services"
)

var (
	db *sql.DB
)

type Job struct {
	Name           string
	CronExpression string
	FunctionName   string
	Package        string
}

func LoadJobs() ([]Job, error) {
	var jobs []Job
	// 使用GORM的方法进行查询
	err := middleware.DB.Table("scheduled_tasks").Select("name, cron_expression, function_name, package").Scan(&jobs).Error
	if err != nil {
		log.Fatal("Failed to load tasks:", err)
		return nil, err
	}
	for _, job := range jobs {
		fn := getFunction(job.Package, job.FunctionName)
		if fn.IsValid() {
			taskFunc := func() {
				fn.Call(nil)
			}
			RegisterTask(job.Name, taskFunc)

			jobs = append(jobs, job)
		} else {
			log.Printf("Function %s not found in package %s", job.FunctionName, job.Package)
		}
	}
	return jobs, nil
}

func ExecuteWithLock(taskName string) {
	lockKey := "lock:" + taskName
	expiration := 10 * time.Second
	// 尝试获取锁
	identifier, acquired := middleware.AcquireLock(lockKey, expiration)
	if !acquired {
		log.Println("Failed to acquire lock" + lockKey)
		return
	}
	defer middleware.ReleaseLock(lockKey, identifier) //
	ExecuteTask(taskName)
}
func getFunction(pkgName, funcName string) reflect.Value {
	switch pkgName {
	case "services":
		// Assuming there is a struct that encapsulates the methods
		manager := services.CronService{} // You need to define this struct
		return reflect.ValueOf(&manager).MethodByName(funcName)
	default:
		return reflect.Value{}
	}
}
