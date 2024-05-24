package services

type CronService struct{}

func (cs *CronService) FairLaunchIssuance() {
	FairLaunchDebugLogger.Info("start cron job: FairLaunchIssuance")
	FairLaunchIssuance()
}

func (cs *CronService) FairLaunchMint() {
	FairLaunchDebugLogger.Info("start cron job: FairLaunchMint")
	FairLaunchMint()
}

func (sm *CronService) PollPaymentCron() {
	CUST.Info("start cron job: PollPayment")
	pollPayment()
}
func (sm *CronService) PollInvoiceCron() {
	CUST.Info("start cron job: PollInvoice")
	pollInvoice()
}
