package services

type CronService struct{}

func (sm *CronService) FairLaunchIssuance() {
	CUST.Info("start cron job: FairLaunchIssuance")
	FairLaunchIssuance()
}
