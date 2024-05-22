package services

import "trade/utils"

var (
	FeeLimit = 100000000
)

func CalculateFee(amount int) (int, error) {
	// TODO: need to complete
	utils.LogInfo("CalculateFee triggered. This function did nothing, need to complete.")
	return 0, nil
}

func PayFee(fee int, limit int, userId int) error {
	// TODO: need to complete
	utils.LogInfo("PayFee triggered. This function did nothing, need to complete.")
	return nil
}
