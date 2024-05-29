package services

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"trade/api"
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

func GetFairLaunchMintedInfo(id int) (*models.FairLaunchMintedInfo, error) {
	f := FairLaunchStore{DB: middleware.DB}
	return f.ReadFairLaunchMintedInfo(uint(id))
}

func GetFairLaunchMintedInfosByFairLaunchId(fairLaunchId int) (*[]models.FairLaunchMintedInfo, error) {
	f := FairLaunchStore{DB: middleware.DB}
	var fairLaunchMintedInfos []models.FairLaunchMintedInfo
	//err := f.DB.Where("fair_launch_info_id = ?", int(uint(id))).Find(&fairLaunchMintedInfos).Error
	err := f.DB.Where(&models.FairLaunchMintedInfo{FairLaunchInfoID: int(uint(fairLaunchId))}).Find(&fairLaunchMintedInfos).Error
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
		AssetType:              taprpc.AssetType(assetType),
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
func ProcessFairLaunchMintedInfo(fairLaunchInfoID int, mintedNumber int, mintedFeeRateSatPerKw int, addr string, userId int) (*models.FairLaunchMintedInfo, error) {
	var fairLaunchMintedInfo models.FairLaunchMintedInfo
	isFairLaunchMintTimeRight, err := IsFairLaunchMintTimeRight(fairLaunchInfoID)
	if err != nil {
		return nil, err
	}
	if !isFairLaunchMintTimeRight {
		err = errors.New("not valid mint time")
		return nil, err
	}
	decodedAddrInfo, err := api.GetDecodedAddrInfo(addr)
	if err != nil {
		return nil, err
	}
	//calculatedGasFeeRateSatPerKw, _ := CalculateGasFeeRateSatPerKw(mintedNumber, 6)
	fairLaunchMintedInfo = models.FairLaunchMintedInfo{
		FairLaunchInfoID:      fairLaunchInfoID,
		MintedNumber:          mintedNumber,
		MintedFeeRateSatPerKw: mintedFeeRateSatPerKw,
		EncodedAddr:           addr,
		UserID:                userId,
		AssetID:               hex.EncodeToString(decodedAddrInfo.AssetId),
		AssetType:             decodedAddrInfo.AssetType.String(),
		AddrAmount:            int(decodedAddrInfo.Amount),
		ScriptKey:             hex.EncodeToString(decodedAddrInfo.ScriptKey),
		InternalKey:           hex.EncodeToString(decodedAddrInfo.InternalKey),
		TaprootOutputKey:      hex.EncodeToString(decodedAddrInfo.TaprootOutputKey),
		ProofCourierAddr:      decodedAddrInfo.ProofCourierAddr,
		MintTime:              utils.GetTimestamp(),
		State:                 models.FairLaunchMintedStateNoPay,
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
			State:            models.FairLaunchInventoryStateOpen,
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
		IsFairLaunch:   true,
		FairLaunchID:   int(fairLaunchInfo.ID),
		State:          models.AssetIssuanceStatePending,
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
	err := middleware.DB.Where("fair_launch_info_id = ? AND status = ?", fairLaunchInfoId, models.StatusNormal).Find(&fairLaunchInventoryInfos).Error
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
	err := middleware.DB.Where("fair_launch_info_id = ? AND status = ? AND is_minted = ? AND state = ?", fairLaunchInfoId, models.StatusNormal, false, models.FairLaunchInventoryStateOpen).Find(&fairLaunchInventoryInfos).Error
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

// IsMintAvailable
// @Description: Is Mint Available by fairLaunchInfoId and number
// @param id
// @param number
// @return bool
func IsMintAvailable(fairLaunchInfoId int, number int) bool {
	inventoryNumber, err := GetNumberOfInventoryCouldBeMinted(fairLaunchInfoId)
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
		return false
	}
	return inventoryNumber >= number
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

// LockInventoryByFairLaunchMintedIdAndMintNumber
// @Description: Calculate MintAmount By id and MintNumber, then Update State, this function will lock inventory
// @param fairLaunchInfoId
// @param number
// @return *[]models.FairLaunchInventoryInfo
// @return error
func LockInventoryByFairLaunchMintedIdAndMintNumber(fairLaunchMintedInfoId int, number int) (*[]models.FairLaunchInventoryInfo, error) {
	if number <= 0 {
		err := errors.New("mint number must be greater than zero")
		FairLaunchDebugLogger.Error("", err)
		return nil, err
	}
	fairLaunchMintedInfo, err := GetFairLaunchMintedInfo(fairLaunchMintedInfoId)
	if err != nil {
		FairLaunchDebugLogger.Error("Get FairLaunchMintedInfo", err)
		return nil, err
	}
	fairLaunchInventoryInfos, err := GetInventoryCouldBeMintedByFairLaunchInfoId(fairLaunchMintedInfo.FairLaunchInfoID)
	if err != nil {
		FairLaunchDebugLogger.Error("Get Inventory Could Be Minted By FairLaunchInfoId", err)
		return nil, err
	}
	allNum := len(*fairLaunchInventoryInfos)
	if allNum < number {
		err = errors.New("not enough mint amount")
		FairLaunchDebugLogger.Error("", err)
		return nil, err
	}
	mintInventoryInfos := (*fairLaunchInventoryInfos)[:number]
	//for _, inventory := range mintInventoryInfos {
	//	inventory.Status = models.StatusPending
	//}
	err = middleware.DB.Model(&mintInventoryInfos).Updates(map[string]any{"state": models.FairLaunchInventoryStateLocked, "minted_id": fairLaunchMintedInfoId}).Error
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
// maybe deprecated
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

func UpdateFairLaunchInfoPaidId(fairLaunchInfo *models.FairLaunchInfo, paidId int) (err error) {
	fairLaunchInfo.IssuanceFeePaidID = paidId
	f := FairLaunchStore{DB: middleware.DB}
	return f.UpdateFairLaunchInfo(fairLaunchInfo)
}

func ChangeFairLaunchInfoState(fairLaunchInfo *models.FairLaunchInfo, state models.FairLaunchState) (err error) {
	fairLaunchInfo.State = state
	f := FairLaunchStore{DB: middleware.DB}
	return f.UpdateFairLaunchInfo(fairLaunchInfo)
}

func UpdateFairLaunchInfoBatchKeyAndBatchState(fairLaunchInfo *models.FairLaunchInfo, batchKey string, batchState string) (err error) {
	fairLaunchInfo.BatchKey = batchKey
	fairLaunchInfo.BatchState = batchState
	f := FairLaunchStore{DB: middleware.DB}
	return f.UpdateFairLaunchInfo(fairLaunchInfo)
}

func UpdateFairLaunchInfoBatchTxidAndAssetId(fairLaunchInfo *models.FairLaunchInfo, batchTxidAnchor string, batchState string, assetId string) (err error) {
	fairLaunchInfo.BatchTxidAnchor = batchTxidAnchor
	fairLaunchInfo.BatchState = batchState
	fairLaunchInfo.AssetID = assetId
	f := FairLaunchStore{DB: middleware.DB}
	return f.UpdateFairLaunchInfo(fairLaunchInfo)
}

func FairLaunchTapdMint(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// @dev: 1.taprpc MintAsset
	var isCollectible bool
	if taprpc.AssetType(fairLaunchInfo.AssetType) == taprpc.AssetType_COLLECTIBLE {
		isCollectible = true
	}
	newMeta := api.NewMeta(fairLaunchInfo.Description, fairLaunchInfo.ImageData)
	mintResponse, err := api.MintAssetAndGetResponse(fairLaunchInfo.Name, isCollectible, newMeta, fairLaunchInfo.Amount, false)
	if err != nil {
		FairLaunchDebugLogger.Error("Tapd Mint Asset.", err)
		return err
	}
	// @dev: 2.update batchKey and batchState
	batchKey := hex.EncodeToString(mintResponse.GetPendingBatch().GetBatchKey())
	batchState := mintResponse.GetPendingBatch().GetState().String()
	err = UpdateFairLaunchInfoBatchKeyAndBatchState(fairLaunchInfo, batchKey, batchState)
	if err != nil {
		FairLaunchDebugLogger.Error("Update FairLaunchInfo BatchKey And BatchState", err)
		return err
	}
	return nil
}

func FairLaunchTapdMintFinalize(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	finalizeResponse, err := api.FinalizeBatchAndGetResponse(fairLaunchInfo.FeeRate)
	if err != nil {
		FairLaunchDebugLogger.Error("Tapd Mint finalize.", err)
		return err
	}
	if hex.EncodeToString(finalizeResponse.GetBatch().GetBatchKey()) != fairLaunchInfo.BatchKey {
		err = errors.New("finalize batch key is not equal mint batch key")
		FairLaunchDebugLogger.Error("Tapd Mint finalize.", err)
		return err
	}
	batchTxidAnchor := finalizeResponse.GetBatch().GetBatchTxid()
	batchState := finalizeResponse.GetBatch().GetState().String()
	assetId, err := api.BatchTxidAnchorToAssetId(batchTxidAnchor)
	if err != nil {
		FairLaunchDebugLogger.Error("Batch Anchor Txid To AssetId.", err)
		return err
	}
	err = UpdateFairLaunchInfoBatchTxidAndAssetId(fairLaunchInfo, batchTxidAnchor, batchState, assetId)
	if err != nil {
		FairLaunchDebugLogger.Error("Update FairLaunchInfo BatchTxid And AssetId.", err)
		return err
	}
	return nil
}

func GetTransactionConfirmedNumber(txid string) (mumConfirmations int, err error) {
	response, err := api.GetListChainTransactions()
	if err != nil {
		FairLaunchDebugLogger.Error("Get List ChainTransactions", err)
		return 0, err
	}
	for _, transaction := range *response {
		if txid == transaction.TxHash {
			return transaction.NumConfirmations, nil
		}
	}
	err = errors.New("did not match transaction hash")
	return 0, err
}

func IsTransactionConfirmed(txid string) bool {
	mumConfirmations, err := GetTransactionConfirmedNumber(txid)
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
		return false
	}
	return mumConfirmations > 0
}

func UpdateFairLaunchInfoReservedCouldMintAndState(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	fairLaunchInfo.ReservedCouldMint = true
	fairLaunchInfo.State = models.FairLaunchStateIssued
	f := FairLaunchStore{DB: middleware.DB}
	return f.UpdateFairLaunchInfo(fairLaunchInfo)
}

func GetFairLaunchInfoState(fairLaunchId int) (fairLaunchState models.FairLaunchState, err error) {
	var fairLaunchInfo *models.FairLaunchInfo
	fairLaunchInfo, err = GetFairLaunchInfo(fairLaunchId)
	if err != nil {
		FairLaunchDebugLogger.Error("Get FairLaunchInfo", err)
		return 0, err
	}
	return fairLaunchInfo.State, nil
}

func IsFairLaunchIssued(fairLaunchId int) bool {
	state, err := GetFairLaunchInfoState(fairLaunchId)
	if err != nil {
		FairLaunchDebugLogger.Error("Get FairLaunchInfo State", err)
		return false
	}
	return state == models.FairLaunchStateIssued
}

// FairLaunchInfos Procession

func ProcessFairLaunchStateNoPayInfoService(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// @dev: 1.pay fee
	paidId, err := PayIssuanceFee(fairLaunchInfo.UserID, fairLaunchInfo.FeeRate)
	if err != nil {
		FairLaunchDebugLogger.Error("Pay Mint Fee.", err)
		return nil
	}
	// @dev: 2.Store paidId
	err = UpdateFairLaunchInfoPaidId(fairLaunchInfo, paidId)
	if err != nil {
		FairLaunchDebugLogger.Error("Update FairLaunchInfo PaidId", err)
		return err
	}
	// @dev: 3.Change state
	err = ChangeFairLaunchInfoState(fairLaunchInfo, models.FairLaunchStatePaidPending)
	if err != nil {
		FairLaunchDebugLogger.Error("Change FairLaunchInfo State.", err)
		return err
	}
	return nil
}

func ProcessFairLaunchStatePaidPendingInfoService(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// @dev: 1.fee paid
	if IsIssuanceFeePaid(fairLaunchInfo.IssuanceFeePaidID) {
		// @dev: Change state
		err = ChangeFairLaunchInfoState(fairLaunchInfo, models.FairLaunchStatePaidNoIssue)
		if err != nil {
			FairLaunchDebugLogger.Error("Change FairLaunchInfo State.", err)
			return err
		}
		return nil
	}
	// @dev: fee has not been paid
	FairLaunchDebugLogger.Info("fairLaunchInfo:", fairLaunchInfo.ID, "is in Paid Pending State:", fairLaunchInfo.IssuanceFeePaidID)
	return nil
}

func ProcessFairLaunchStatePaidNoIssueInfoService(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// @dev: 1.tapd mint, add to batch, finalize
	err = FairLaunchTapdMint(fairLaunchInfo)
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
		return err
	}
	// @TODO: Consider whether to use scheduled task to finalize
	err = FairLaunchTapdMintFinalize(fairLaunchInfo)
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
		return err
	}
	// @dev: 2.Update asset issuance table
	err = CreateAssetIssuanceInfoByFairLaunchInfo(fairLaunchInfo)
	// @dev: 3.update inventory
	err = CreateInventoryInfoByFairLaunchInfo(fairLaunchInfo)
	// @dev: Change state
	err = ChangeFairLaunchInfoState(fairLaunchInfo, models.FairLaunchStateIssuedPending)
	if err != nil {
		FairLaunchDebugLogger.Error("Change FairLaunchInfo State.", err)
		return err
	}
	return nil
}

func ProcessFairLaunchStateIssuedPendingInfoService(fairLaunchInfo *models.FairLaunchInfo) (err error) {
	// @dev: 1.Is Transaction Confirmed
	if IsTransactionConfirmed(fairLaunchInfo.BatchTxidAnchor) {
		// @dev: Update FairLaunchInfo ReservedCouldMint And Change State
		err = UpdateFairLaunchInfoReservedCouldMintAndState(fairLaunchInfo)
		if err != nil {
			FairLaunchDebugLogger.Error("Update FairLaunchInfo ReservedCouldMint And Change State.", err)
			return err
		}
		// @dev: Update Asset Issuance
		var a = AssetIssuanceStore{DB: middleware.DB}
		var assetIssuance *models.AssetIssuance
		assetIssuance, err = a.ReadAssetIssuanceByFairLaunchId(fairLaunchInfo.ID)
		if err != nil {
			FairLaunchDebugLogger.Error("Read AssetIssuance By FairLaunchId.", err)
			return err
		}
		assetIssuance.State = models.AssetIssuanceStateIssued
		err = a.UpdateAssetIssuance(assetIssuance)
		if err != nil {
			FairLaunchDebugLogger.Error("Update AssetIssuance.", err)
			return err
		}
		return nil
	}
	// @dev: Transaction has not been Confirmed
	FairLaunchDebugLogger.Info("fairLaunchInfo:", fairLaunchInfo.ID, "is in Issued Pending State:", fairLaunchInfo.BatchTxidAnchor)
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

func UpdateFairLaunchMintedInfoPaidId(fairLaunchMintedInfo *models.FairLaunchMintedInfo, paidId int) (err error) {
	fairLaunchMintedInfo.MintFeePaidID = paidId
	f := FairLaunchStore{DB: middleware.DB}
	return f.UpdateFairLaunchMintedInfo(fairLaunchMintedInfo)
}

func ChangeFairLaunchMintedInfoState(fairLaunchMintedInfo *models.FairLaunchMintedInfo, state models.FairLaunchMintedState) (err error) {
	fairLaunchMintedInfo.State = state
	f := FairLaunchStore{DB: middleware.DB}
	return f.UpdateFairLaunchMintedInfo(fairLaunchMintedInfo)
}

func LockInventoryByFairLaunchMintedInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (lockedInventory *[]models.FairLaunchInventoryInfo, err error) {
	//fairLaunchId := fairLaunchMintedInfo.FairLaunchInfoID
	//mintNumber := fairLaunchMintedInfo.MintedNumber
	lockedInventory, err = LockInventoryByFairLaunchMintedIdAndMintNumber(int(fairLaunchMintedInfo.ID), fairLaunchMintedInfo.MintedNumber)
	if err != nil {
		FairLaunchDebugLogger.Error("Lock Inventory By FairLaunchId And MintNumber", err)
		return nil, err
	}
	return lockedInventory, nil
}

func GetAllUnsentFairLaunchMintedInfos() (fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, err error) {
	err = middleware.DB.Where("status = ? AND state = ? AND is_addr_sent = ?", models.StatusNormal, models.FairLaunchInventoryStateLocked, false).Find(fairLaunchMintedInfos).Error
	if err != nil {
		utils.LogError("Get all fairLaunch minted infos error. ", err)
		return nil, err
	}
	return fairLaunchMintedInfos, err
}

// @dev: dprecated
func UpdateFairLaunchMintedInfosIsAddrSent(fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, isAddrSent bool) (err error) {
	return middleware.DB.Model(&fairLaunchMintedInfos).Update("is_addr_sent", isAddrSent).Error
}

func SendAssetResponseScriptKeyAndInternalKeyToOutpoint(sendAssetResponse *taprpc.SendAssetResponse, scriptKey string, internalKey string) (outpoint string, err error) {
	for _, output := range sendAssetResponse.Transfer.Outputs {
		outputScriptKey := hex.EncodeToString(output.ScriptKey)
		outputAnchorInternalKey := hex.EncodeToString(output.Anchor.InternalKey)
		if outputScriptKey == scriptKey && outputAnchorInternalKey == internalKey {
			return output.Anchor.Outpoint, nil
		}
	}
	err = errors.New("can not find anchor outpoint value")
	return "", err
}

// GetTransactionAndIndexByOutpoint
// @dev: Split outpoint
func GetTransactionAndIndexByOutpoint(outpoint string) (transaction string, index string) {
	result := strings.Split(outpoint, ":")
	return result[0], result[1]
}

func GetListChainTransactionsOutpointAddress(outpoint string) (address string, err error) {
	response, err := api.GetListChainTransactions()
	if err != nil {
		FairLaunchDebugLogger.Error("Get List ChainTransactions", err)
		return "", err
	}
	tx, indexStr := GetTransactionAndIndexByOutpoint(outpoint)
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		FairLaunchDebugLogger.Error("strconv.Atoi(indexStr)", err)
		return "", err
	}
	for _, transaction := range *response {
		if transaction.TxHash == tx {
			return transaction.DestAddresses[index], nil
		}
	}
	err = errors.New("did not match transaction outpoint")
	return "", err
}

// @dev: Updated outpoint and is_addr_sent
func UpdateFairLaunchMintedInfosBySendAssetResponse(fairLaunchMintedInfos *[]models.FairLaunchMintedInfo, sendAssetResponse *taprpc.SendAssetResponse) (err error) {
	// deprecate anchor tx hash
	_ = hex.EncodeToString(sendAssetResponse.Transfer.AnchorTxHash)
	for _, fairLaunchMintedInfo := range *fairLaunchMintedInfos {
		scriptKey := fairLaunchMintedInfo.ScriptKey
		internalKey := fairLaunchMintedInfo.InternalKey
		var outpoint string
		outpoint, err = SendAssetResponseScriptKeyAndInternalKeyToOutpoint(sendAssetResponse, scriptKey, internalKey)
		fairLaunchMintedInfo.OutpointTxHash, _ = GetTransactionAndIndexByOutpoint(outpoint)
		if err != nil {
			FairLaunchDebugLogger.Error("Send Asset Response ScriptKey And InternalKey To Outpoint", err)
			return err
		}
		// @dev: Update outpoint and isAddrSent
		fairLaunchMintedInfo.Outpoint = outpoint
		fairLaunchMintedInfo.IsAddrSent = true
		var address string
		address, err = GetListChainTransactionsOutpointAddress(outpoint)
		if err != nil {
			FairLaunchDebugLogger.Error("Get List Chain Transactions Outpoint Address", err)
			return err
		}
		fairLaunchMintedInfo.Address = address
	}
	return middleware.DB.Save(fairLaunchMintedInfos).Error
}

// @dev: Trigger after ProcessFairLaunchMintedStatePaidNoSendInfo
func SendFairLaunchMintedAssetLocked() (err error) {
	// @dev: all unsent
	unsentFairLaunchMintedInfos, err := GetAllUnsentFairLaunchMintedInfos()
	if err != nil {
		FairLaunchDebugLogger.Error("Get All Unsent FairLaunchMintedInfos", err)
		return err
	}
	// @dev: addr Slice
	var addrSlice []string
	for _, fairLaunchMintedInfo := range *unsentFairLaunchMintedInfos {
		addrSlice = append(addrSlice, fairLaunchMintedInfo.EncodedAddr)
	}
	feeRateSatPerKw, err := EstimateSmartFeeRateSatPerKw()
	if err != nil {
		return err
	}
	// @dev: Send Asset
	response, err := api.SendAssetAddrSliceAndGetResponse(addrSlice, feeRateSatPerKw)
	if err != nil {
		FairLaunchDebugLogger.Error("Send Asset AddrSlice And Get Response", err)
		return err
	}
	// @dev: Update minted info
	err = UpdateFairLaunchMintedInfosBySendAssetResponse(unsentFairLaunchMintedInfos, response)
	if err != nil {
		FairLaunchDebugLogger.Error("Update By FairLaunchMintedInfos And SendAssetResponse", err)
		return err
	}
	return nil
}

func GetAllLockedInventoryByFairLaunchMintedInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (*[]models.FairLaunchInventoryInfo, error) {
	var fairLaunchInventoryInfos []models.FairLaunchInventoryInfo
	err := middleware.DB.Where("status = ? AND state = ? AND minted_id = ?", models.StatusNormal, models.FairLaunchInventoryStateLocked, fairLaunchMintedInfo.ID).Find(&fairLaunchInventoryInfos).Error
	if err != nil {
		FairLaunchDebugLogger.Error("DB Find by state AND minted_id", err)
		return nil, err
	}
	return &fairLaunchInventoryInfos, nil
}

func UpdateLockedInventoryByFairLaunchMintedInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	fairLaunchMintedInfos, err := GetAllLockedInventoryByFairLaunchMintedInfo(fairLaunchMintedInfo)
	if err != nil {
		FairLaunchDebugLogger.Error("Get All Locked Inventory By FairLaunchMintedInfo", err)
		return err
	}
	// @dev: Update
	err = middleware.DB.Model(&fairLaunchMintedInfos).Updates(map[string]any{"is_minted": true, "state": models.FairLaunchInventoryStateMinted}).Error
	if err != nil {
		FairLaunchDebugLogger.Error("DB Updates is_minted, state", err)
		return err
	}
	return nil
}

func UpdateMintedNumberAndIsMintAllOfFairLaunchInfoByFairLaunchMintedInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	fairLaunchInfoId := fairLaunchMintedInfo.FairLaunchInfoID
	fairLaunchInfo, err := GetFairLaunchInfo(fairLaunchInfoId)
	if err != nil {
		FairLaunchDebugLogger.Error("Get FairLaunchInfo", err)
		return err
	}
	var isMintAll bool
	if fairLaunchInfo.MintedNumber+fairLaunchMintedInfo.MintedNumber >= fairLaunchInfo.MintNumber {
		isMintAll = true
	}
	fairLaunchInfo.MintedNumber += fairLaunchMintedInfo.MintedNumber
	fairLaunchInfo.IsMintAll = isMintAll
	return middleware.DB.Save(fairLaunchInfo).Error
}

// FairLaunchMintedInfos Procession

func ProcessFairLaunchMintedStateNoPayInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	// @dev: 1.pay fee
	paidId, err := PayMintFee(fairLaunchMintedInfo.UserID, fairLaunchMintedInfo.MintedFeeRateSatPerKw)
	if err != nil {
		FairLaunchDebugLogger.Error("Pay Issuance Fee.", err)
		return nil
	}
	// @dev: 2.Store paidId
	err = UpdateFairLaunchMintedInfoPaidId(fairLaunchMintedInfo, paidId)
	if err != nil {
		FairLaunchDebugLogger.Error("Update FairLaunchMintedInfo PaidId", err)
		return err
	}
	// @dev: 3.Change state
	err = ChangeFairLaunchMintedInfoState(fairLaunchMintedInfo, models.FairLaunchMintedStatePaidPending)
	if err != nil {
		FairLaunchDebugLogger.Error("Change FairLaunchMintedInfo State.", err)
		return err
	}
	return nil
}

func ProcessFairLaunchMintedStatePaidPendingInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	// @dev: 1.fee paid
	if IsMintFeePaid(fairLaunchMintedInfo.MintFeePaidID) {
		// @dev: Change state
		err = ChangeFairLaunchMintedInfoState(fairLaunchMintedInfo, models.FairLaunchMintedStatePaidNoSend)
		if err != nil {
			FairLaunchDebugLogger.Error("Change FairLaunchMintedInfo State.", err)
			return err
		}
		return nil
	}
	// @dev: fee has not been paid
	FairLaunchDebugLogger.Info("fairLaunchMintedInfo:", fairLaunchMintedInfo.ID, "is in Paid Pending State:", fairLaunchMintedInfo.MintFeePaidID)
	return nil
}

func ProcessFairLaunchMintedStatePaidNoSendInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	// @dev: Locked Inventory
	lockedInventory, err := LockInventoryByFairLaunchMintedInfo(fairLaunchMintedInfo)
	if err != nil {
		FairLaunchDebugLogger.Error("Lock Inventory By FairLaunchMintedInfo", err)
		return err
	}
	// @dev: Calculate mint amount
	calculatedMintAmount := CalculateMintAmountByFairLaunchInventoryInfos(lockedInventory)
	if calculatedMintAmount != fairLaunchMintedInfo.AddrAmount {
		err = errors.New("calculated amount is not equal fairLaunchMintedInfo's addr amount")
		FairLaunchDebugLogger.Error("calculatedMintAmount != fairLaunchMintedInfo.AddrAmount", err)
		return err
	}
	// @dev: Change state
	err = ChangeFairLaunchMintedInfoState(fairLaunchMintedInfo, models.FairLaunchMintedStateSentPending)
	if err != nil {
		FairLaunchDebugLogger.Error("Change FairLaunchMintedInfo State.", err)
		return err
	}
	return nil
}

func ProcessFairLaunchMintedStateSentPendingInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) (err error) {
	// @dev: 1.Is Transaction Confirmed
	if IsTransactionConfirmed(fairLaunchMintedInfo.OutpointTxHash) {
		// @dev: Change state
		err = ChangeFairLaunchMintedInfoState(fairLaunchMintedInfo, models.FairLaunchMintedStateSent)
		if err != nil {
			FairLaunchDebugLogger.Error("Change FairLaunchMintedInfo State.", err)
			return err
		}
		// @dev: Update MintedNumber and IsMintAll
		err = UpdateMintedNumberAndIsMintAllOfFairLaunchInfoByFairLaunchMintedInfo(fairLaunchMintedInfo)
		if err != nil {
			FairLaunchDebugLogger.Error("Update MintedNumber And IsMintAll Of FairLaunchInfo By FairLaunchMintedInfo", err)
			return err
		}
		// Update Inventory
		err = UpdateLockedInventoryByFairLaunchMintedInfo(fairLaunchMintedInfo)
		if err != nil {
			FairLaunchDebugLogger.Error("Update Locked Inventory By FairLaunchMintedInfo", err)
			return err
		}
		// Update minted user
		f := FairLaunchStore{DB: middleware.DB}
		err = f.CreateFairLaunchMintedUserInfo(&models.FairLaunchMintedUserInfo{
			UserID:           fairLaunchMintedInfo.UserID,
			FairLaunchInfoID: fairLaunchMintedInfo.FairLaunchInfoID,
			MintedNumber:     fairLaunchMintedInfo.MintedNumber,
		})
		if err != nil {
			FairLaunchDebugLogger.Error("Create FairLaunch Minted UserInfo", err)
			return err
		}
		return nil
	}
	// @dev: Transaction has not been Confirmed
	FairLaunchDebugLogger.Info("fairLaunchMintedInfo:", fairLaunchMintedInfo.ID, "is in Sent Pending State:", fairLaunchMintedInfo.OutpointTxHash)
	return nil
}

// FairLaunchIssuance
// @Description: Scheduled Task
func FairLaunchIssuance() {
	processionResult, err := ProcessAllFairLaunchInfos()
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
		return
	}
	FairLaunchDebugLogger.Error(utils.ValueJsonString(processionResult))
}

// FairLaunchMint
// @Description: Scheduled Task
func FairLaunchMint() {
	processionResult, err := ProcessAllFairLaunchMintedInfos()
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
		return
	}
	FairLaunchDebugLogger.Error(utils.ValueJsonString(processionResult))
}

// SendFairLaunchAsset
// @Description: Scheduled Task
func SendFairLaunchAsset() {
	err := SendFairLaunchMintedAssetLocked()
	if err != nil {
		FairLaunchDebugLogger.Error("", err)
		return
	}
}

func SendFairLaunchReserved(fairLaunchInfo *models.FairLaunchInfo, addr string) (response *taprpc.SendAssetResponse, err error) {
	decodedAddrInfo, err := api.GetDecodedAddrInfo(addr)
	if err != nil {
		FairLaunchDebugLogger.Error("Get Decoded Addr Info", err)
		return nil, err
	}
	if int(decodedAddrInfo.Amount) != fairLaunchInfo.ReserveTotal {
		err = errors.New("wrong addr amount value")
		FairLaunchDebugLogger.Error("", err)
		return nil, err
	}
	// send
	addrSlice := []string{addr}
	feeRateSatPerKw, err := EstimateSmartFeeRateSatPerKw()
	response, err = api.SendAssetAddrSliceAndGetResponse(addrSlice, feeRateSatPerKw)
	if err != nil {
		FairLaunchDebugLogger.Error("Send Asset AddrSlice And Get Response", err)
		return nil, err
	}
	return response, nil
}
