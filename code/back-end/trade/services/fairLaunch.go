package services

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
	"trade/middleware"
	"trade/models"
	"trade/utils"
)

func GetAllFairLaunchInfos() (*[]models.FairLaunchInfo, error) {
	f := FairLaunchStore{DB: middleware.DB}
	var fairLaunchInfos []models.FairLaunchInfo
	err := f.DB.Find(&fairLaunchInfos).Error
	return &fairLaunchInfos, err
}

func GetFairLaunchInfo(id int) (*models.FairLaunchInfo, error) {
	f := FairLaunchStore{DB: middleware.DB}
	return f.ReadFairLaunchInfo(uint(id))
}

func GetFairLaunchMintedInfo(id int) (*[]models.FairLaunchMintedInfo, error) {
	f := FairLaunchStore{DB: middleware.DB}
	var fairLaunchMintedInfos []models.FairLaunchMintedInfo
	//err := f.DB.Where("fair_launch_info_id = ?", int(uint(id))).Find(&fairLaunchMintedInfos).Error
	err := f.DB.Where(&models.FairLaunchMintedInfo{FairLaunchInfoID: int(uint(id))}).Find(&fairLaunchMintedInfos).Error
	return &fairLaunchMintedInfos, err
}

func SetFairLaunchInfo(fairLaunchInfo *models.FairLaunchInfo) error {
	f := FairLaunchStore{DB: middleware.DB}
	return f.CreateFairLaunchInfo(fairLaunchInfo)
}

func SetFairLaunchMintedInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) error {
	f := FairLaunchStore{DB: middleware.DB}
	return f.CreateFairLaunchMintedInfo(fairLaunchMintedInfo)
}

// ProcessFairLaunchInfo
// @Description: Process fairLaunchInfo
// @param imageData
// @param name
// @param assetType
// @param amount
// @param reserved
// @param mintQuantity
// @param startTime
// @param endTime
// @param description
// @param feeRate
// @param userId
// @return *models.FairLaunchInfo
// @return error
func ProcessFairLaunchInfo(imageData string, name string, assetType int, amount int, reserved int, mintQuantity int, startTime int, endTime int, description string, feeRate int, userId int) (*models.FairLaunchInfo, error) {
	calculateSeparateAmount, err := AmountReservedAndMintQuantityToReservedTotalAndMintTotal(amount, reserved, mintQuantity)
	if err != nil {
		utils.LogError("Calculate separate amount", err)
		return nil, err
	}
	var fairLaunchInfo models.FairLaunchInfo
	fairLaunchInfo = models.FairLaunchInfo{
		ImageData:              imageData,
		Name:                   name,
		AssetType:              assetType,
		Amount:                 amount,
		Reserved:               reserved,
		MintQuantity:           mintQuantity,
		StartTime:              startTime,
		EndTime:                endTime,
		Description:            description,
		FeeRate:                feeRate,
		ActualReserved:         calculateSeparateAmount.ActualReserved,
		ReserveTotal:           calculateSeparateAmount.ReserveTotal,
		MintNumber:             calculateSeparateAmount.MintNumber,
		IsFinalEnough:          calculateSeparateAmount.IsFinalEnough,
		FinalQuantity:          calculateSeparateAmount.FinalQuantity,
		MintTotal:              calculateSeparateAmount.MintTotal,
		ActualMintTotalPercent: calculateSeparateAmount.ActualMintTotalPercent,
		CalculationExpression:  calculateSeparateAmount.CalculationExpression,
		UserID:                 userId,
		State:                  models.FairLaunchStateNoPay,
	}
	return &fairLaunchInfo, nil
}

// ProcessFairLaunchMintedInfo
// @Description: Process fairLaunchMintedInfo
// @param fairLaunchInfoID
// @param mintedNumber
// @param userId
// @return *models.FairLaunchMintedInfo
// @return error
func ProcessFairLaunchMintedInfo(fairLaunchInfoID int, mintedNumber int, userId int) (*models.FairLaunchMintedInfo, error) {
	var fairLaunchMintedInfo models.FairLaunchMintedInfo
	isFairLaunchMintTimeRight, err := IsFairLaunchMintTimeRight(fairLaunchInfoID)
	if err != nil {
		return nil, err
	}
	if !isFairLaunchMintTimeRight {
		err = errors.New("not valid mint time")
		return nil, err
	}
	fairLaunchMintedInfo = models.FairLaunchMintedInfo{
		FairLaunchInfoID: fairLaunchInfoID,
		MintedNumber:     mintedNumber,
		UserID:           userId,
		MintTime:         utils.GetTimestamp(),
		State:            models.FairLaunchMintedStateNoPay,
	}
	return &fairLaunchMintedInfo, nil
}

type CalculateSeparateAmount struct {
	Amount                 int
	Reserved               int
	ActualReserved         float64
	ReserveTotal           int
	MintQuantity           int
	MintNumber             int
	IsFinalEnough          bool
	FinalQuantity          int
	MintTotal              int
	ActualMintTotalPercent float64
	CalculationExpression  string
}

// AmountReservedAndMintQuantityToReservedTotalAndMintTotal
// @Description: return Calculated result struct
// @param amount
// @param reserved
// @param mintQuantity
// @return *CalculateSeparateAmount
// @return error
func AmountReservedAndMintQuantityToReservedTotalAndMintTotal(amount int, reserved int, mintQuantity int) (*CalculateSeparateAmount, error) {
	if amount <= 0 || reserved <= 0 || mintQuantity <= 0 {
		return nil, errors.New("amount reserved and mint amount must be greater than zero")
	}
	if reserved > 99 {
		return nil, errors.New("reserved amount must be less equal than 99")
	}
	if amount <= mintQuantity {
		return nil, errors.New("amount must be greater than mint quantity")
	}
	reservedTotal := int(math.Ceil(float64(amount) * float64(reserved) / 100))
	mintTotal := amount - reservedTotal
	remainder := mintTotal % mintQuantity
	var finalQuantity int
	var isFinalEnough bool
	if remainder == 0 {
		isFinalEnough = true
		finalQuantity = mintQuantity
	} else {
		isFinalEnough = false
		finalQuantity = remainder
	}
	if mintTotal <= 0 || mintTotal < mintQuantity {
		return nil, errors.New("insufficient mint total amount")
	}
	reservedTotal = amount - mintTotal
	if reservedTotal <= 0 {
		return nil, errors.New("reserved amount is less equal than zero")
	}

	mintNumber := int(math.Ceil(float64(mintTotal) / float64(mintQuantity)))
	if mintNumber <= 0 {
		return nil, errors.New("mint number is less equal than zero")
	}
	actualReserved := float64(reservedTotal) * 100 / float64(amount)
	actualReserved = utils.RoundToDecimalPlace(actualReserved, 8)
	actualMintTotalPercent := 100 - actualReserved
	calculatedSeparateAmount := CalculateSeparateAmount{
		Amount:                 amount,
		Reserved:               reserved,
		ActualReserved:         actualReserved,
		ReserveTotal:           reservedTotal,
		MintQuantity:           mintQuantity,
		MintNumber:             mintNumber,
		IsFinalEnough:          isFinalEnough,
		FinalQuantity:          finalQuantity,
		MintTotal:              mintTotal,
		ActualMintTotalPercent: actualMintTotalPercent,
	}
	var err error
	calculatedSeparateAmount.CalculationExpression, err = CalculationExpressionBySeparateAmount(&calculatedSeparateAmount)
	if err != nil {
		utils.LogError("CalculationExpressionBySeparateAmount error.", err)
		return nil, err
	}
	return &calculatedSeparateAmount, nil
}

// CalculationExpressionBySeparateAmount
// @Description: Generate Calculation Expression By Separate Amount
// @param calculateSeparateAmount
// @return string
// @return error
func CalculationExpressionBySeparateAmount(calculateSeparateAmount *CalculateSeparateAmount) (string, error) {
	calculated := calculateSeparateAmount.ReserveTotal + calculateSeparateAmount.MintQuantity*(calculateSeparateAmount.MintNumber-1) + calculateSeparateAmount.FinalQuantity
	if reflect.DeepEqual(calculated, calculateSeparateAmount.Amount) {
		return fmt.Sprintf("%d+%d*%d+%d=%d", calculateSeparateAmount.ReserveTotal, calculateSeparateAmount.MintQuantity, calculateSeparateAmount.MintNumber-1, calculateSeparateAmount.FinalQuantity, calculated), nil
	}
	return "", errors.New("calculated result is not equal amount")
}

// CreateInventoryInfoByFairLaunchInfo
// @Description: Create Inventory Info By FairLaunchInfo
// @param fairLaunchInfo
// @return error
func CreateInventoryInfoByFairLaunchInfo(fairLaunchInfo *models.FairLaunchInfo) error {
	var FairLaunchInventoryInfos []models.FairLaunchInventoryInfo
	items := fairLaunchInfo.MintNumber - 1
	for ; items > 0; items -= 1 {
		FairLaunchInventoryInfos = append(FairLaunchInventoryInfos, models.FairLaunchInventoryInfo{
			FairLaunchInfoID: int(fairLaunchInfo.ID),
			Quantity:         fairLaunchInfo.MintQuantity,
		})
	}
	FairLaunchInventoryInfos = append(FairLaunchInventoryInfos, models.FairLaunchInventoryInfo{
		FairLaunchInfoID: int(fairLaunchInfo.ID),
		Quantity:         fairLaunchInfo.FinalQuantity,
	})
	f := FairLaunchStore{DB: middleware.DB}
	return f.CreateFairLaunchInventoryInfos(&FairLaunchInventoryInfos)
}

// CreateAssetIssuanceInfoByFairLaunchInfo
// @Description: Create Asset Issuance Info By FairLaunchInfo
// @param fairLaunchInfo
// @return error
func CreateAssetIssuanceInfoByFairLaunchInfo(fairLaunchInfo *models.FairLaunchInfo) error {
	assetIssuance := models.AssetIssuance{
		AssetName:      fairLaunchInfo.Name,
		AssetId:        fairLaunchInfo.AssetID,
		AssetType:      fairLaunchInfo.AssetType,
		IssuanceUserId: fairLaunchInfo.UserID,
		IssuanceTime:   utils.GetTimestamp(),
	}
	a := AssetIssuanceStore{DB: middleware.DB}
	return a.CreateAssetIssuance(&assetIssuance)
}

// GetAllInventoryInfoByFairLaunchInfoId
// @Description: Query all inventory by FairLaunchInfo id
// @param fairLaunchInfoId
// @return *[]models.FairLaunchInventoryInfo
// @return error
func GetAllInventoryInfoByFairLaunchInfoId(fairLaunchInfoId int) (*[]models.FairLaunchInventoryInfo, error) {
	var fairLaunchInventoryInfos []models.FairLaunchInventoryInfo
	err := middleware.DB.Where("fair_launch_info_id = ? AND status = ?", fairLaunchInfoId, 1).Find(&fairLaunchInventoryInfos).Error
	if err != nil {
		utils.LogError("Get all inventory info by fair launch id. ", err)
		return nil, err
	}
	return &fairLaunchInventoryInfos, err
}

// GetInventoryCouldBeMintedByFairLaunchInfoId
// @Description: Get all Inventory Could Be Minted By FairLaunchInfoId
// @param fairLaunchInfoId
// @return *[]models.FairLaunchInventoryInfo
// @return error
func GetInventoryCouldBeMintedByFairLaunchInfoId(fairLaunchInfoId int) (*[]models.FairLaunchInventoryInfo, error) {
	var fairLaunchInventoryInfos []models.FairLaunchInventoryInfo
	err := middleware.DB.Where("fair_launch_info_id = ? AND status = ? AND is_minted = ?", fairLaunchInfoId, models.StatusNormal, false).Find(&fairLaunchInventoryInfos).Error
	if err != nil {
		utils.LogError("Get all inventory info could be minted by fair launch id. ", err)
		return nil, err
	}
	return &fairLaunchInventoryInfos, err
}

// GetNumberOfInventoryCouldBeMinted
// @Description: call GetInventoryCouldBeMintedByFairLaunchInfoId
// @param fairLaunchInfoId
// @return int
// @return error
func GetNumberOfInventoryCouldBeMinted(fairLaunchInfoId int) (int, error) {
	fairLaunchInventoryInfos, err := GetInventoryCouldBeMintedByFairLaunchInfoId(fairLaunchInfoId)
	if err != nil {
		utils.LogError("", err)
		return 0, err
	}
	return len(*fairLaunchInventoryInfos), err
}

// GetMintAmountByFairLaunchMintNumber
// @Description: Get Mint Amount By FairLaunch id and MintNumber
// @param fairLaunchInfoId
// @param number
// @return amount
// @return err
func GetMintAmountByFairLaunchMintNumber(fairLaunchInfoId int, number int) (amount int, err error) {
	if number <= 0 {
		err = errors.New("mint number must be greater than zero")
		utils.LogError("", err)
		return 0, err
	}
	fairLaunchInventoryInfos, err := GetInventoryCouldBeMintedByFairLaunchInfoId(fairLaunchInfoId)
	if err != nil {
		utils.LogError("", err)
		return 0, err
	}
	allNum := len(*fairLaunchInventoryInfos)
	if allNum < number {
		err = errors.New("not enough mint amount")
		utils.LogError("", err)
		return 0, err
	}
	mintInventoryInfos := (*fairLaunchInventoryInfos)[:number]
	for _, inventory := range mintInventoryInfos {
		amount += inventory.Quantity
	}
	return amount, err
}

// LockInventoryByFairLaunchIdAndMintNumber
// @Description: Calculate MintAmount By id and MintNumber, then Update State, this function will lock inventory
// @param fairLaunchInfoId
// @param number
// @return *[]models.FairLaunchInventoryInfo
// @return error
func LockInventoryByFairLaunchIdAndMintNumber(fairLaunchInfoId int, number int) (*[]models.FairLaunchInventoryInfo, error) {
	if number <= 0 {
		err := errors.New("mint number must be greater than zero")
		utils.LogError("", err)
		return nil, err
	}
	fairLaunchInventoryInfos, err := GetInventoryCouldBeMintedByFairLaunchInfoId(fairLaunchInfoId)
	if err != nil {
		utils.LogError("", err)
		return nil, err
	}
	allNum := len(*fairLaunchInventoryInfos)
	if allNum < number {
		err = errors.New("not enough mint amount")
		utils.LogError("", err)
		return nil, err
	}
	mintInventoryInfos := (*fairLaunchInventoryInfos)[:number]
	for _, inventory := range mintInventoryInfos {
		inventory.Status = models.StatusPending
	}
	err = middleware.DB.Model(&mintInventoryInfos).Update("status", models.StatusPending).Error
	return &mintInventoryInfos, err
}

// CalculateMintAmountByFairLaunchInventoryInfos
// @Description: Calculate MintAmount By FairLaunchInventoryInfos
// @param fairLaunchInventoryInfos
// @return amount
func CalculateMintAmountByFairLaunchInventoryInfos(fairLaunchInventoryInfos *[]models.FairLaunchInventoryInfo) (amount int) {
	for _, inventory := range *fairLaunchInventoryInfos {
		amount += inventory.Quantity
	}
	return amount
}

// IsDuringMintTime
// @Description: timestamp now is between start and end
// @param start
// @param end
// @return bool
func IsDuringMintTime(start int, end int) bool {
	now := int(time.Now().Unix())
	return now >= start && now < end
}

// IsFairLaunchInfoMintTimeValid
// @Description: call IsDuringMintTime
// @param fairLaunchInfo
// @return bool
func IsFairLaunchInfoMintTimeValid(fairLaunchInfo *models.FairLaunchInfo) bool {
	return IsDuringMintTime(fairLaunchInfo.StartTime, fairLaunchInfo.EndTime)
}

// IsFairLaunchMintTimeRight
// @Description: call GetFairLaunchInfo and IsFairLaunchInfoMintTimeValid
// @param fairLaunchInfoId
// @return bool
// @return error
func IsFairLaunchMintTimeRight(fairLaunchInfoId int) (bool, error) {
	fairLaunchInfo, err := GetFairLaunchInfo(fairLaunchInfoId)
	if err != nil {
		utils.LogError("", err)
		return false, err
	}
	return IsFairLaunchInfoMintTimeValid(fairLaunchInfo), nil
}

// AmountAndQuantityToNumber
// @Description: calculate Number by Amount And Quantity
// @param amount
// @param quantity
// @return int
func AmountAndQuantityToNumber(amount int, quantity int) int {
	return int(math.Ceil(float64(amount) / float64(quantity)))
}

// CreateInventoryAndAssetIssuanceInfoByFairLaunchInfo
// @Description: Update inventory and asset issuance
// @param fairLaunchInfo
// @return err
func CreateInventoryAndAssetIssuanceInfoByFairLaunchInfo(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	err = CreateInventoryInfoByFairLaunchInfo(fairLaunchInfo)
	if err != nil {
		return err
	}
	err = CreateAssetIssuanceInfoByFairLaunchInfo(fairLaunchInfo)
	if err != nil {
		return err
	}
	return nil
}

// FairLaunchInfos

func GetAllFairLaunchInfoByState(state models.FairLaunchState) (fairLaunchInfos *[]models.FairLaunchInfo, err error) {
	err = middleware.DB.Where("status = ? AND state = ?", models.StatusNormal, state).Find(fairLaunchInfos).Error
	if err != nil {
		utils.LogError("Get all fairLaunch info by state. ", err)
		return nil, err
	}
	return fairLaunchInfos, err
}

func GetAllFairLaunchStateNoPayInfos() (fairLaunchInfos *[]models.FairLaunchInfo, err error) {
	return GetAllFairLaunchInfoByState(models.FairLaunchStateNoPay)
}

func GetAllFairLaunchStatePaidPendingInfos() (fairLaunchInfos *[]models.FairLaunchInfo, err error) {
	return GetAllFairLaunchInfoByState(models.FairLaunchStatePaidPending)
}

func GetAllFairLaunchStatePaidNoIssueInfos() (fairLaunchInfos *[]models.FairLaunchInfo, err error) {
	return GetAllFairLaunchInfoByState(models.FairLaunchStatePaidNoIssue)
}

func GetAllFairLaunchStateIssuedPendingInfos() (fairLaunchInfos *[]models.FairLaunchInfo, err error) {
	return GetAllFairLaunchInfoByState(models.FairLaunchStateIssuedPending)
}

func GetAllFairLaunchStateIssuedInfos() (fairLaunchInfos *[]models.FairLaunchInfo, err error) {
	return GetAllFairLaunchInfoByState(models.FairLaunchStateIssued)
}

func GetAllValidFairLaunchInfos() (fairLaunchInfos *[]models.FairLaunchInfo, err error) {
	err = middleware.DB.Where("status = ?", models.StatusNormal).Find(fairLaunchInfos).Error
	if err != nil {
		utils.LogError("Get all fairLaunch infos error. ", err)
		return nil, err
	}
	return fairLaunchInfos, err
}

type ProcessionResult struct {
	id int
	models.JsonResult
}

func ProcessAllFairLaunchInfos() (*[]ProcessionResult, error) {
	var processionResults []ProcessionResult
	allFairLaunchInfos, err := GetAllValidFairLaunchInfos()
	if err != nil {
		FairLaunchDebugLogger.Error("Get all fairLaunch infos error. ", err)
	}
	for _, fairLaunchInfo := range *allFairLaunchInfos {
		if fairLaunchInfo.State == models.FairLaunchStateNoPay {
			err = ProcessFairLaunchStateNoPayInfoService(&fairLaunchInfo)
			if err != nil {
				FairLaunchDebugLogger.Error("Process FairLaunch info Service error. ", err)
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchInfo.ID),
					JsonResult: models.JsonResult{
						Success: false,
						Error:   err.Error(),
						Data:    nil,
					},
				})
				continue
			} else {
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchInfo.ID),
					JsonResult: models.JsonResult{
						Success: true,
						Error:   "",
						Data:    nil,
					},
				})
			}
		} else if fairLaunchInfo.State == models.FairLaunchStatePaidPending {
			err = ProcessFairLaunchStatePaidPendingInfoService(&fairLaunchInfo)
			if err != nil {
				FairLaunchDebugLogger.Error("Process FairLaunch info Service error. ", err)
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchInfo.ID),
					JsonResult: models.JsonResult{
						Success: false,
						Error:   err.Error(),
						Data:    nil,
					},
				})
				continue
			} else {
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchInfo.ID),
					JsonResult: models.JsonResult{
						Success: true,
						Error:   "",
						Data:    nil,
					},
				})
			}
		} else if fairLaunchInfo.State == models.FairLaunchStatePaidNoIssue {
			err = ProcessFairLaunchStatePaidNoIssueInfoService(&fairLaunchInfo)
			if err != nil {
				FairLaunchDebugLogger.Error("Process FairLaunch info Service error. ", err)
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchInfo.ID),
					JsonResult: models.JsonResult{
						Success: false,
						Error:   err.Error(),
						Data:    nil,
					},
				})
				continue
			} else {
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchInfo.ID),
					JsonResult: models.JsonResult{
						Success: true,
						Error:   "",
						Data:    nil,
					},
				})
			}
		} else if fairLaunchInfo.State == models.FairLaunchStateIssuedPending {
			err = ProcessFairLaunchStateIssuedPendingInfoService(&fairLaunchInfo)
			if err != nil {
				FairLaunchDebugLogger.Error("Process FairLaunch info Service error. ", err)
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchInfo.ID),
					JsonResult: models.JsonResult{
						Success: false,
						Error:   err.Error(),
						Data:    nil,
					},
				})
				continue
			} else {
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchInfo.ID),
					JsonResult: models.JsonResult{
						Success: true,
						Error:   "",
						Data:    nil,
					},
				})
			}
		}
	}
	if processionResults == nil || len(processionResults) == 0 {
		err = errors.New("procession results error")
		return nil, err
	}
	return &processionResults, nil
}

func ProcessFairLaunchStateNoPayInfoService(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// TODO: need to complete
	// TODO: Call
	return nil
}

func ProcessFairLaunchStatePaidPendingInfoService(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// TODO: need to complete
	return nil
}

func ProcessFairLaunchStatePaidNoIssueInfoService(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// TODO: need to complete
	return nil
}

func ProcessFairLaunchStateIssuedPendingInfoService(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// TODO: need to complete
	return nil
}

// FairLaunchMintedInfos

func GetAllFairLaunchMintedInfoByState(state models.FairLaunchMintedState) (fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, err error) {
	err = middleware.DB.Where("status = ? AND state = ?", models.StatusNormal, state).Find(fairLaunchMintedInfos).Error
	if err != nil {
		utils.LogError("Get all fairLaunch minted info by state. ", err)
		return nil, err
	}
	return fairLaunchMintedInfos, err
}

func GetAllFairLaunchMintedStateNoPayInfo() (fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, err error) {
	return GetAllFairLaunchMintedInfoByState(models.FairLaunchMintedStateNoPay)
}

func GetAllFairLaunchMintedStatePaidPendingInfo() (fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, err error) {
	return GetAllFairLaunchMintedInfoByState(models.FairLaunchMintedStatePaidPending)
}

func GetAllFairLaunchMintedStatePaidNoSendInfo() (fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, err error) {
	return GetAllFairLaunchMintedInfoByState(models.FairLaunchMintedStatePaidNoSend)
}

func GetAllFairLaunchMintedStateSentPendingInfo() (fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, err error) {
	return GetAllFairLaunchMintedInfoByState(models.FairLaunchMintedStateSentPending)
}

func GetAllFairLaunchMintedStateSentInfo() (fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, err error) {
	return GetAllFairLaunchMintedInfoByState(models.FairLaunchMintedStateSent)
}

func GetAllValidFairLaunchMintedInfos() (fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, err error) {
	err = middleware.DB.Where("status = ?", models.StatusNormal).Find(fairLaunchMintedInfos).Error
	if err != nil {
		utils.LogError("Get all fairLaunch minted infos error. ", err)
		return nil, err
	}
	return fairLaunchMintedInfos, err
}

func ProcessAllFairLaunchMintedInfos() (*[]ProcessionResult, error) {
	var processionResults []ProcessionResult
	allFairLaunchMintedInfos, err := GetAllValidFairLaunchMintedInfos()
	if err != nil {
		FairLaunchDebugLogger.Error("Get all fairLaunch minted infos error. ", err)
	}
	for _, fairLaunchMintedInfo := range *allFairLaunchMintedInfos {
		if fairLaunchMintedInfo.State == models.FairLaunchMintedStateNoPay {
			err = ProcessFairLaunchMintedStateNoPayInfo(&fairLaunchMintedInfo)
			if err != nil {
				FairLaunchDebugLogger.Error("Process FairLaunch info Service error. ", err)
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchMintedInfo.ID),
					JsonResult: models.JsonResult{
						Success: false,
						Error:   err.Error(),
						Data:    nil,
					},
				})
				continue
			} else {
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchMintedInfo.ID),
					JsonResult: models.JsonResult{
						Success: true,
						Error:   "",
						Data:    nil,
					},
				})
			}
		} else if fairLaunchMintedInfo.State == models.FairLaunchMintedStatePaidPending {
			err = ProcessFairLaunchMintedStatePaidPendingInfo(&fairLaunchMintedInfo)
			if err != nil {
				FairLaunchDebugLogger.Error("Process FairLaunch info Service error. ", err)
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchMintedInfo.ID),
					JsonResult: models.JsonResult{
						Success: false,
						Error:   err.Error(),
						Data:    nil,
					},
				})
				continue
			} else {
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchMintedInfo.ID),
					JsonResult: models.JsonResult{
						Success: true,
						Error:   "",
						Data:    nil,
					},
				})
			}
		} else if fairLaunchMintedInfo.State == models.FairLaunchMintedStatePaidNoSend {
			err = ProcessFairLaunchMintedStatePaidNoSendInfo(&fairLaunchMintedInfo)
			if err != nil {
				FairLaunchDebugLogger.Error("Process FairLaunch info Service error. ", err)
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchMintedInfo.ID),
					JsonResult: models.JsonResult{
						Success: false,
						Error:   err.Error(),
						Data:    nil,
					},
				})
				continue
			} else {
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchMintedInfo.ID),
					JsonResult: models.JsonResult{
						Success: true,
						Error:   "",
						Data:    nil,
					},
				})
			}
		} else if fairLaunchMintedInfo.State == models.FairLaunchMintedStateSentPending {
			err = ProcessFairLaunchMintedStateSentPendingInfo(&fairLaunchMintedInfo)
			if err != nil {
				FairLaunchDebugLogger.Error("Process FairLaunch info Service error. ", err)
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchMintedInfo.ID),
					JsonResult: models.JsonResult{
						Success: false,
						Error:   err.Error(),
						Data:    nil,
					},
				})
				continue
			} else {
				processionResults = append(processionResults, ProcessionResult{
					id: int(fairLaunchMintedInfo.ID),
					JsonResult: models.JsonResult{
						Success: true,
						Error:   "",
						Data:    nil,
					},
				})
			}
		}
	}
	if processionResults == nil || len(processionResults) == 0 {
		err = errors.New("procession results error")
		return nil, err
	}
	return &processionResults, nil
}

func ProcessFairLaunchMintedStateNoPayInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	// TODO: need to complete
	return nil
}

func ProcessFairLaunchMintedStatePaidPendingInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	// TODO: need to complete
	return nil
}

func ProcessFairLaunchMintedStatePaidNoSendInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	// TODO: need to complete
	return nil
}

func ProcessFairLaunchMintedStateSentPendingInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	// TODO: need to complete
	return nil
}

// TODO: need to process all
func FairLaunchIssuance() {
	// TODO: need to complete
	FairLaunchDebugLogger.Info("FairLaunchIssuance triggered.")

}

// TODO: need to process all
func FairLaunchMint() {
	// TODO: need to complete
	FairLaunchDebugLogger.Info("FairLaunchMint triggered.")

}
