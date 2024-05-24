package services

type CronService struct{}

func (sm *CronService) FairLaunchIssuance() {
	CUST.Info("start cron job: FairLaunchIssuance")
	FairLaunchIssuance()
}

func (sm *CronService) PollPaymentCron() {
	CUST.Info("start cron job: PollPayment")
	pollPayment()
}
func (sm *CronService) PollInvoiceCron() {
	CUST.Info("start cron job: PollInvoice")
	pollInvoice()
}
