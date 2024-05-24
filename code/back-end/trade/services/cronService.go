package services

type CronService struct{}

func (sm *CronService) FairLaunchIssuance() {
	FairLaunchDebugLogger.Info("start cron job: FairLaunchIssuance")
	FairLaunchIssuance()
}
