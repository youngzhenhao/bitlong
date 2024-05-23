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

// TODO: variables need to modify
func PayMintFee(fee int, limit int, userId int) error {
	// TODO: need to complete
	utils.LogInfo("PayMintFee triggered. This function did nothing, need to complete.")
	return nil
}

// TODO: variables need to modify
func PayReleaseFee() error {
	// TODO: need to complete
	utils.LogInfo("PayReleaseFee triggered. This function did nothing, need to complete.")
	return nil
}
