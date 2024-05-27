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
func GetPayMintFeeState(paidId int) error {
	// TODO: need to complete
	utils.LogInfo("GetPayMintFeeState triggered. This function did nothing, need to complete.")
	return nil
}

func IsMintFeePaid(paidId int) bool {
	_ = GetPayMintFeeState(paidId)
	// TODO: need to complete
	utils.LogInfo("IsMintFeePaid triggered. This function did nothing, need to complete.")
	return true
}

// TODO: variables need to modify
func GetPayIssuanceFeeState(paidId int) error {
	// TODO: need to complete
	utils.LogInfo("GetPayIssuanceFeeState triggered. This function did nothing, need to complete.")
	return nil
}

func IsIssuanceFeePaid(paidId int) bool {
	_ = GetPayIssuanceFeeState(paidId)
	// TODO: need to complete
	utils.LogInfo("IsIssuanceFeePaid triggered. This function did nothing, need to complete.")
	return true
}

func PayMintFee() (mintFeePaidId int, err error) {
	// TODO: need to complete
	utils.LogInfo("PayMintFee triggered. This function did nothing, need to complete.")
	return 0, nil
}

func PayIssuanceFee() (IssuanceFeePaidId int, err error) {
	// TODO: need to complete
	utils.LogInfo("PayIssuanceFee triggered. This function did nothing, need to complete.")
	return 0, nil
}

// Consider latest block all transactions, use fee median
