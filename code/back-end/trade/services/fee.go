package services

import (
	"errors"
	"trade/api"
)

var (
	GasFeeRateOfNumber0  float64 = 1
	GasFeeRateOfNumber1          = 1.2
	GasFeeRateOfNumber2  float64 = 2
	GasFeeRateOfNumber3  float64 = 3
	GasFeeRateOfNumber4  float64 = 4
	GasFeeRateOfNumber5          = 4.5
	GasFeeRateOfNumber6  float64 = 5
	GasFeeRateOfNumber7          = 5.4
	GasFeeRateOfNumber8          = 5.7
	GasFeeRateOfNumber9  float64 = 6
	GasFeeRateOfNumber10         = 6.5
)

func NumberToGasFeeRate(number int) (gasFeeRate float64, err error) {
	if number < 0 || number > 10 {
		err = errors.New("number out of range")
		// Max rate
		return GasFeeRateOfNumber10, err
	} else if number == 1 {
		return GasFeeRateOfNumber1, nil
	} else if number == 2 {
		return GasFeeRateOfNumber2, nil
	} else if number == 3 {
		return GasFeeRateOfNumber3, nil
	} else if number == 4 {
		return GasFeeRateOfNumber4, nil
	} else if number == 5 {
		return GasFeeRateOfNumber5, nil
	} else if number == 6 {
		return GasFeeRateOfNumber6, nil
	} else if number == 7 {
		return GasFeeRateOfNumber7, nil
	} else if number == 8 {
		return GasFeeRateOfNumber8, nil
	} else if number == 9 {
		return GasFeeRateOfNumber9, nil
	} else if number == 10 {
		return GasFeeRateOfNumber10, nil
	} else {
		return GasFeeRateOfNumber0, nil
	}
}

// BTC/kB
func EstimateSmartFeeRate(blocks int) (gasFeeRate float64, err error) {
	feeResult, err := api.EstimateSmartFeeAndGetResult(blocks)
	if err != nil {
		FEE.Error("Estimate SmartFee And GetResult", err)
		return 0, err
	}
	if feeResult.Errors != nil || feeResult.Blocks != int64(blocks) || *feeResult.FeeRate == 0 {
		err = errors.New("fee result got error or blocks is not same or fee rate is zero")
		FEE.Error("Invalid fee rate result", err)
		return 0, err
	}
	return *feeResult.FeeRate, nil
}

func EstimateSmartFeeRateSatPerKw(blocks int) (estimatedFeeSatPerKw int, err error) {
	estimatedFee, err := EstimateSmartFeeRate(blocks)
	if err != nil {
		FEE.Error("Estimate Smart FeeRate", err)
		return 0, err
	}
	estimatedFeeSatPerKw = BtcPerKbToSatPerKw(estimatedFee)
	return estimatedFeeSatPerKw, nil
}

// BtcPerKbToSatPerKw
// @Description: 1 sat/vB = 0.25 sat/wu
// https://bitcoin.stackexchange.com/questions/106333/different-fee-rate-units-sat-vb-sat-perkw-sat-perkb
func BtcPerKbToSatPerKw(btcPerKb float64) (satPerKw int) {
	// @dev: 1 BTC/kB = 1e8 sat/kB 1e5 sat/B = 0.25e5 sat/w = 0.25e8 sat/kw
	return int(0.25e8 * btcPerKb)
}

func FeeRateSatPerKwToSatPerB(feeRateSatPerKw int) (feeRateSatPerB int) {
	return feeRateSatPerKw * 4 / 1000
}

// BTC/kB
func CalculateGasFeeRateBtcPerKb(number int, blocks int) (float64, error) {
	rate, err := NumberToGasFeeRate(number)
	if err != nil {
		FEE.Error("Number To Gas FeeRate", err)
		return 0, err
	}
	estimatedFee, err := EstimateSmartFeeRate(blocks)
	if err != nil {
		FEE.Error("Estimate Smart FeeRate", err)
		return 0, err
	}
	feeRateBtcPerKb := rate * estimatedFee
	return feeRateBtcPerKb, nil
}

// sat/kw
func CalculateGasFeeRateSatPerKw(number int, blocks int) (feeRateSatPerKw int, err error) {
	feeRateBtcPerKb, err := CalculateGasFeeRateBtcPerKb(number, blocks)
	if err != nil {
		FEE.Error("Calculate Gas FeeRate BtcPerKb", err)
		return 0, err
	}
	feeRateSatPerKw = BtcPerKbToSatPerKw(feeRateBtcPerKb)
	return feeRateSatPerKw, nil
}

// sat/B
func CalculateGasFeeRateSatPerB(number int, blocks int) (feeRateSatPerB int, err error) {
	feeRateBtcPerKb, err := CalculateGasFeeRateBtcPerKb(number, blocks)
	if err != nil {
		FEE.Error("Calculate Gas FeeRate BtcPerKb", err)
		return 0, err
	}
	feeRateSatPerB = int(feeRateBtcPerKb * 1e8 / 1e3)
	return feeRateSatPerB, nil
}

// @dev: not actual value
func GetTransactionByteSize() int {
	// TODO: need to complete
	return 250
}

func CalculateGasFee(number int, blocks int, byteSize int) (int, error) {
	calculatedGasFeeRateSatPerB, err := CalculateGasFeeRateSatPerB(number, blocks)
	if err != nil {
		FEE.Error("Calculate GasFeeRate SatPerB", err)
		return 0, err
	}
	gasFee := byteSize * calculatedGasFeeRateSatPerB
	return gasFee, nil
}

func GetPayMintFeeState(paidId int) (int, error) {
	balance, err := GetBalance(uint(paidId))
	if err != nil {
		FEE.Error("GetBalance", err)
		return 0, err
	}
	return int(balance.State), nil
}

func IsMintFeePaid(paidId int) bool {
	state, err := GetPayMintFeeState(paidId)
	if err != nil {
		FEE.Error("GetBalance", err)
		return false
	}
	if state == PAY_SUCCESS {
		return true
	}
	return false
}

func GetPayIssuanceFeeState(paidId int) (int, error) {
	balance, err := GetBalance(uint(paidId))
	if err != nil {
		FEE.Error("GetBalance", err)
		return 0, err
	}
	return int(balance.State), nil
}

func IsIssuanceFeePaid(paidId int) bool {
	state, err := GetPayIssuanceFeeState(paidId)
	if err != nil {
		FEE.Error("GetBalance", err)
		return false
	}
	if state == PAY_SUCCESS {
		return true
	}
	return false
}

func PayMintFee(userId int, feeRateSatPerKw int) (mintFeePaidId int, err error) {
	fee := FeeRateSatPerKwToSatPerB(feeRateSatPerKw) * GetTransactionByteSize()
	return PayGasFee(userId, fee)
}

func PayIssuanceFee(userId int, feeRateSatPerKw int) (IssuanceFeePaidId int, err error) {
	// TODO: User need to pay more fee
	fee := FeeRateSatPerKwToSatPerB(feeRateSatPerKw) * GetTransactionByteSize()
	return PayGasFee(userId, fee)
}

func PayGasFee(payUserId int, gasFee int) (int, error) {
	id, err := PayAmountToAdmin(uint(payUserId), uint64(gasFee), 0)
	return int(id), err
}
