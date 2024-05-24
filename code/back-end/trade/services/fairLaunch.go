package services

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"gorm.io/gorm"
	"io"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"
	"trade/api"
	"trade/config"
	"trade/middleware"
	"trade/models"
	"trade/utils"
)

// TODO: Need to test
func GetAllFairLaunch() (*[]models.FairLaunchInfo, error) {
	f := FairLaunchStore{DB: middleware.DB}
	var fairLaunchInfos []models.FairLaunchInfo
	err := f.DB.Find(&fairLaunchInfos).Error
	return &fairLaunchInfos, err
}

// TODO: Need to test
func GetFairLaunch(id int) (*models.FairLaunchInfo, error) {
	f := FairLaunchStore{DB: middleware.DB}
	return f.ReadFairLaunchInfo(uint(id))
}

// TODO: Need to test
func GetMinted(id int) (*[]models.FairLaunchMintedInfo, error) {
	f := FairLaunchStore{DB: middleware.DB}
	var fairLaunchMintedInfos []models.FairLaunchMintedInfo
	//err := f.DB.Where("fair_launch_info_id = ?", int(uint(id))).Find(&fairLaunchMintedInfos).Error
	err := f.DB.Where(&models.FairLaunchMintedInfo{FairLaunchInfoID: int(uint(id))}).Find(&fairLaunchMintedInfos).Error
	return &fairLaunchMintedInfos, err
}

// TODO: Need to test
func SetFairLaunch(fairLaunchInfo *models.FairLaunchInfo) error {
	f := FairLaunchStore{DB: middleware.DB}
	return f.CreateFairLaunchInfo(fairLaunchInfo)
}

// TODO: Need to test
func FairLaunchMint(fairLaunchMintedInfo *models.FairLaunchMintedInfo) error {
	// @dev: 1.Query info
	// TODO: Use this fairLaunchInfo
	fairLaunchInfo, err := GetFairLaunch(fairLaunchMintedInfo.FairLaunchInfoID)
	if err != nil {
		utils.LogError("Get fair launch by id of fairLaunchMintedInfo", err)
		return err
	}
	if fairLaunchInfo.Status != 1 {
		utils.LogError("fair launch info status is 1.", err)
		return errors.New("fair launch status is not valid")
	}
	// TODO: check time now whether between start and end

	// TODO: check mint is valid

	//// @previous: 2.Calculate number of asset
	//amt := calculateAmount(int(fairLaunchInfo.ID), fairLaunchMintedInfo.AddrAmount)
	//if amt == 0 {
	//	err = errors.New("amount of asset to send is zero")
	//	utils.LogError("", err)
	//	return err
	//}
	//// @previous: 3.Request an addr
	//addr := GetAddr(fairLaunchInfo.AssetID, amt)

	// @dev: 4.Pay asset to addr
	// TODO: feeRate need to set
	result, err := api.SendAssetBool(fairLaunchMintedInfo.EncodedAddr, 0)
	if !result {
		utils.LogError("Send asset error", err)
		return err
	}
	// TODO: 5.Write to database
	f := FairLaunchStore{DB: middleware.DB}
	err = f.CreateFairLaunchMintedInfo(fairLaunchMintedInfo)
	if err != nil {
		utils.LogError("Create fair launch minted info error.", err)
		return err
	}
	// TODO: need to complete

	return nil
}

func GetAddr(id string, amt int) string {
	// TODO: need to complete
	// TODO: Request User to generate an addr of receiving asset
	// TODO: Call PostPhoneToNewAddr
	// TODO: User local_port and remote_port
	utils.LogInfo("GetAddr triggered. This function did nothing, need to complete.")
	return ""
}

func PostPhoneToNewAddr(remotePort string, assetId string, amount int) (*taprpc.Addr, error) {
	frpsForwardSocket := fmt.Sprintf("%s:%s", config.GetLoadConfig().FrpsServer, remotePort)
	targetUrl := "http://" + frpsForwardSocket + "/newAddr"
	payload := url.Values{"asset_id": {assetId}, "amount": {strconv.Itoa(amount)}}
	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		utils.LogError("http.PostForm.", err)
		return nil, err
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var addrResponse struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Data    *taprpc.Addr
	}
	if err := json.Unmarshal(bodyBytes, &addrResponse); err != nil {
		utils.LogError("PPTNA json.Unmarshal.", err)
		return nil, err
	}
	return addrResponse.Data, nil
}

func CalculateAmount(id int, amount int) int {
	// TODO: need to complete
	// TODO: add number logic, or judge number when mint
	// TODO: Verify amount is valid
	utils.LogInfo("calculateAmount triggered. This function did nothing, need to complete.")
	return 0
}

func ProcessFairLaunchMintedInfo(fairLaunchInfoID int, addr string, mintFeeInvoice string) (*models.FairLaunchMintedInfo, error) {
	var fairLaunchMintedInfo models.FairLaunchMintedInfo
	response, err := api.GetDecodedAddrInfo(addr)
	if err != nil {
		utils.LogError("Decoded Addr error.", err)
		return nil, err
	}
	fairLaunchMintedInfo = models.FairLaunchMintedInfo{
		Model:            gorm.Model{},
		FairLaunchInfoID: fairLaunchInfoID,
		EncodedAddr:      addr,
		MintFeeInvoice:   mintFeeInvoice,
		AssetID:          hex.EncodeToString(response.AssetId),
		AssetType:        response.AssetType.String(),
		AddrAmount:       int(response.Amount),
		ScriptKey:        hex.EncodeToString(response.ScriptKey),
		InternalKey:      hex.EncodeToString(response.InternalKey),
		TaprootOutputKey: hex.EncodeToString(response.TaprootOutputKey),
		ProofCourierAddr: response.ProofCourierAddr,
		MintTime:         utils.GetTimestamp(),
		// TODO: update
		//		Outpoint
		//		Address
	}
	return &fairLaunchMintedInfo, nil
}

func ProcessFairLaunchInfo(imageData string, name string, assetType int, amount int, reserved int, mintQuantity int, startTime int, endTime int, description string, batchKey string, batchState string, batchTxidAnchor string, assetId string, userId int, releaseFeeInvoice string) (*models.FairLaunchInfo, error) {
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
		ActualReserved:         calculateSeparateAmount.ActualReserved,
		ReserveTotal:           calculateSeparateAmount.ReserveTotal,
		MintNumber:             calculateSeparateAmount.MintNumber,
		IsFinalEnough:          calculateSeparateAmount.IsFinalEnough,
		FinalQuantity:          calculateSeparateAmount.FinalQuantity,
		MintTotal:              calculateSeparateAmount.MintTotal,
		ActualMintTotalPercent: calculateSeparateAmount.ActualMintTotalPercent,
		CalculationExpression:  calculateSeparateAmount.CalculationExpression,
		BatchKey:               batchKey,
		BatchState:             batchState,
		BatchTxidAnchor:        batchTxidAnchor,
		AssetID:                assetId,
		UserID:                 userId,
		ReleaseFeeInvoice:      releaseFeeInvoice,
	}
	return &fairLaunchInfo, nil
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

func CalculationExpressionBySeparateAmount(calculateSeparateAmount *CalculateSeparateAmount) (string, error) {
	calculated := calculateSeparateAmount.ReserveTotal + calculateSeparateAmount.MintQuantity*(calculateSeparateAmount.MintNumber-1) + calculateSeparateAmount.FinalQuantity
	if reflect.DeepEqual(calculated, calculateSeparateAmount.Amount) {
		return fmt.Sprintf("%d+%d*%d+%d=%d", calculateSeparateAmount.ReserveTotal, calculateSeparateAmount.MintQuantity, calculateSeparateAmount.MintNumber-1, calculateSeparateAmount.FinalQuantity, calculated), nil
	}
	return "", errors.New("calculated result is not equal amount")
}

// TODO: update IsMinted and MintedID after fairLaunchMint

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

func CreateAssetReleaseInfoByFairLaunchInfo(fairLaunchInfo *models.FairLaunchInfo) error {
	assetRelease := models.AssetRelease{
		AssetName:     fairLaunchInfo.Name,
		AssetId:       fairLaunchInfo.AssetID,
		AssetType:     fairLaunchInfo.AssetType,
		ReleaseUserId: fairLaunchInfo.UserID,
		ReleaseTime:   utils.GetTimestamp(),
	}
	a := AssetReleaseStore{DB: middleware.DB}
	return a.CreateAssetRelease(&assetRelease)
}

// TODO: Query all inventory by FairLaunchInfo id

func GetAllInventoryInfoByFairLaunchInfoId(fairLaunchInfoId int) (*[]models.FairLaunchInventoryInfo, error) {
	var fairLaunchInventoryInfos []models.FairLaunchInventoryInfo
	err := middleware.DB.Where("fair_launch_info_id = ? AND status = ?", fairLaunchInfoId, 1).Find(&fairLaunchInventoryInfos).Error
	if err != nil {
		utils.LogError("Get all inventory info by fair launch id. ", err)
		return nil, err
	}
	return &fairLaunchInventoryInfos, err
}

func GetInventoryCouldBeMintedByFairLaunchInfoId(fairLaunchInfoId int) (*[]models.FairLaunchInventoryInfo, error) {
	var fairLaunchInventoryInfos []models.FairLaunchInventoryInfo
	err := middleware.DB.Where("fair_launch_info_id = ? AND status = ? AND is_minted = ?", fairLaunchInfoId, models.StatusNormal, false).Find(&fairLaunchInventoryInfos).Error
	if err != nil {
		utils.LogError("Get all inventory info could be minted by fair launch id. ", err)
		return nil, err
	}
	return &fairLaunchInventoryInfos, err
}

func GetNumberOfInventoryCouldBeMinted(fairLaunchInfoId int) (int, error) {
	fairLaunchInventoryInfos, err := GetInventoryCouldBeMintedByFairLaunchInfoId(fairLaunchInfoId)
	if err != nil {
		utils.LogError("", err)
		return 0, err
	}
	return len(*fairLaunchInventoryInfos), err
}

// TODO: calculate mint number to mint amount

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

// @previous
// TODO: consider pay fee before change status to prevent user lock asset,or add unlock logic
func CalculateMintAmountByMintNumberAndUpdateState(fairLaunchInfoId int, number int) (*[]models.FairLaunchInventoryInfo, error) {
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
	var amount int
	mintInventoryInfos := (*fairLaunchInventoryInfos)[:number]
	for _, inventory := range mintInventoryInfos {
		inventory.Status = models.StatusPending
		amount += inventory.Quantity
	}
	err = middleware.DB.Model(&mintInventoryInfos).Update("status", models.StatusPending).Error
	// TODO: consider whether amount should be deprecated
	_ = amount
	return &mintInventoryInfos, err
}

// IsDuringMintTime
// @dev: timestamp now is between start and end
func IsDuringMintTime(start int, end int) bool {
	now := int(time.Now().Unix())
	return now >= start && now < end
}

func AmountAndQuantityToNumber(amount int, quantity int) int {
	return int(math.Ceil(float64(amount) / float64(quantity)))
}

// TODO: update fairLaunchInventoryInfos which could be mint, set status to 2, set is_minted and minted_id after mint

// TODO: check mint is valid, query inventory db

// TODO: update
//		IsReservedSent
//		MintedNumber
//		Status
