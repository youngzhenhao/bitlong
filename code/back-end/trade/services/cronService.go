package services

import (
	"fmt"
	"trade/middleware"
	"trade/models"
)

type CronService struct{}

func (cs *CronService) FairLaunchIssuance() {
	FairLaunchDebugLogger.Info("start cron job: FairLaunchIssuance")
	FairLaunchIssuance()
}

func (cs *CronService) FairLaunchMint() {
	FairLaunchDebugLogger.Info("start cron job: FairLaunchMint")
	FairLaunchMint()
}

func (cs *CronService) SendFairLaunchAsset() {
	FairLaunchDebugLogger.Info("start cron job: FairLaunchMint")
	SendFairLaunchAsset()
}

func CreateScheduledTask(scheduledTask *models.ScheduledTask) (err error) {
	s := ScheduledTaskStore{DB: middleware.DB}
	return s.CreateScheduledTask(scheduledTask)
}

func CreateFairLaunchIssuance() (err error) {
	return CreateScheduledTask(&models.ScheduledTask{
		Name:           "FairLaunchIssuance",
		CronExpression: "*/30 * * * * *",
		FunctionName:   "FairLaunchIssuance",
		Package:        "services",
	})
}

func CreateFairLaunchMint() (err error) {
	return CreateScheduledTask(&models.ScheduledTask{
		Name:           "FairLaunchMint",
		CronExpression: "*/30 * * * * *",
		FunctionName:   "FairLaunchMint",
		Package:        "services",
	})
}

func CreateSendFairLaunchAsset() (err error) {
	return CreateScheduledTask(&models.ScheduledTask{
		Name:           "SendFairLaunchAsset",
		CronExpression: "*/30 * * * * *",
		FunctionName:   "SendFairLaunchAsset",
		Package:        "services",
	})
}

func CreateFairLaunchScheduledTasks() {
	err := CreateFairLaunchIssuance()
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
	}
	err = CreateFairLaunchMint()
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
	}
	err = CreateSendFairLaunchAsset()
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
	}
	fmt.Println("Create FairLaunch ScheduledTasks Procession finished!")
}

func (sm *CronService) PollPaymentCron() {
	CUST.Info("start cron job: PollPayment")
	pollPayment()
}
func (sm *CronService) PollInvoiceCron() {
	CUST.Info("start cron job: PollInvoice")
	pollInvoice()
}
