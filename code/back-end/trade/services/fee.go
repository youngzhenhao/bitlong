package services

import (
	"errors"
	"trade/api"
	"trade/config"
	"trade/middleware"
	"trade/models"
)

type (
	FeeRateInfoName string
)

var (
	GasFeeRateOfNumber0    float64         = 1
	GasFeeRateOfNumber1                    = 1.2
	GasFeeRateOfNumber2    float64         = 2
	GasFeeRateOfNumber3    float64         = 3
	GasFeeRateOfNumber4    float64         = 4
	GasFeeRateOfNumber5                    = 4.5
	GasFeeRateOfNumber6    float64         = 5
	GasFeeRateOfNumber7                    = 5.4
	GasFeeRateOfNumber8                    = 5.7
	GasFeeRateOfNumber9    float64         = 6
	GasFeeRateOfNumber10                   = 6.5
	GasFeeRateNameBitcoind FeeRateInfoName = "bitcoind"
	GasFeeRateNameDefault                  = GasFeeRateNameBitcoind
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

func EstimateSmartFeeRateSatPerKw() (estimatedFeeSatPerKw int, err error) {
	err = CheckIfUpdateFeeRateInfo()
	if err != nil {
		return 0, err
	}
	estimatedFee, err := GetEstimateSmartFeeRate()
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
	err = CheckIfUpdateFeeRateInfo()
	if err != nil {
		return 0, err
	}
	estimatedFee, err := GetEstimateSmartFeeRate()
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

func GetPayMintFeeState(paidId int) (bool, error) {
	var balance models.Balance
	err := middleware.DB.Where("state = ?", PAY_SUCCESS).First(&balance, paidId).Error
	if err != nil {
		FEE.Error("GetBalance", err)
		return false, err
	}
	return true, nil
}

func IsMintFeePaid(paidId int) bool {
	state, err := GetPayMintFeeState(paidId)
	if err != nil {
		FEE.Error("GetBalance", err)
		return false
	}
	return state
}

func IsPayIssuanceFeeStatePaid(paidId int) (bool, error) {
	var balance models.Balance
	err := middleware.DB.Where("state = ?", PAY_SUCCESS).First(&balance, paidId).Error
	if err != nil {
		FEE.Error("GetBalance", err)
		return false, err
	}
	return true, nil
}

func IsIssuanceFeePaid(paidId int) bool {
	state, err := IsPayIssuanceFeeStatePaid(paidId)
	if err != nil {
		FEE.Error("GetBalance", err)
		return false
	}
	return state
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

func GetFeeRateInfoByName(name string) (feeRateInfo *models.FeeRateInfo, err error) {
	err = middleware.DB.Where("name = ?", name).First(&feeRateInfo).Error
	if err != nil {
		FEE.Error("Find FeeRateInfo", err)
		return nil, err
	}
	return feeRateInfo, nil
}

func GetFeeRateInfoEstimateSmartFeeRateByName(name string) (estimateSmartFeeRate float64, err error) {
	var feeRateInfo *models.FeeRateInfo
	feeRateInfo, err = GetFeeRateInfoByName(name)
	if err != nil {
		FEE.Error("Get FeeRateInfo By Name", err)
		return 0, err
	}
	return feeRateInfo.EstimateSmartFeeRate, nil
}

func UpdateFeeRateInfoByBitcoind() (err error) {
	var feeRateInfo *models.FeeRateInfo
	var f = FeeRateInfoStore{DB: middleware.DB}
	feeRateInfo, err = GetFeeRateInfoByName(string(GasFeeRateNameBitcoind))
	if err != nil {
		FEE.Error("Get FeeRateInfo By Bitcoind", err)
		//	Create FeeRateInfo
		feeRateInfo = &models.FeeRateInfo{
			Name: string(GasFeeRateNameBitcoind),
		}
		err = f.CreateFeeRateInfo(feeRateInfo)
		if err != nil {
			FEE.Error("Create FeeRate Info", err)
			return err
		}
	}
	feeRateInfo.EstimateSmartFeeRate, err = EstimateSmartFeeRate(config.GetLoadConfig().FairLaunchConfig.EstimateSmartFeeRateBlocks)
	if err != nil {
		FEE.Error("Estimate Smart FeeRate", err)
		return err
	}
	err = f.UpdateFeeRateInfo(feeRateInfo)
	if err != nil {
		FEE.Error("Update FeeRateInfo", err)
		return err
	}
	return nil
}

// @dev: 1.update fee rate or not
func CheckIfUpdateFeeRateInfo() (err error) {
	if config.GetLoadConfig().FairLaunchConfig.IsAutoUpdateFeeRate {
		err = UpdateFeeRateInfoByBitcoind()
		if err != nil {
			FEE.Error("Update FeeRateInfo By Bitcoind", err)
			return err
		}
	}
	return nil
}

// @dev: 2.get fee rate
func GetEstimateSmartFeeRate() (estimateSmartFeeRate float64, err error) {
	return GetFeeRateInfoEstimateSmartFeeRateByName(string(GasFeeRateNameDefault))
}
