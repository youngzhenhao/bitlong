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
func GetPayMintFeeState(invoice string) error {
	// TODO: need to complete
	utils.LogInfo("GetPayMintFeeState triggered. This function did nothing, need to complete.")
	return nil
}

func IsMintFeePaid(invoice string) bool {
	_ = GetPayMintFeeState(invoice)
	// TODO: need to complete
	utils.LogInfo("IsMintFeePaid triggered. This function did nothing, need to complete.")
	return true
}

// TODO: variables need to modify
func GetPayIssuanceFeeState(invoice string) error {
	// TODO: need to complete
	utils.LogInfo("GetPayIssuanceFeeState triggered. This function did nothing, need to complete.")
	return nil
}

func IsIssuanceFeePaid(invoice string) bool {
	_ = GetPayIssuanceFeeState(invoice)
	// TODO: need to complete
	utils.LogInfo("IsIssuanceFeePaid triggered. This function did nothing, need to complete.")
	return true
}
