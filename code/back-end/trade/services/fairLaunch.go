package services

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

	//
	//// TODO: 2.Calculate number of asset
	//amt := calculateAmount(int(fairLaunchInfo.ID), fairLaunchMintedInfo.AddrAmount)
	//if amt == 0 {
	//	err = errors.New("amount of asset to send is zero")
	//	utils.LogError("", err)
	//	return err
	//}
	//// TODO: 3.Request an addr
	//addr := GetAddr(fairLaunchInfo.AssetID, amt)
	//

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
	// TODO: 6.Write to database of asset release

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

func ProcessFairLaunchMintedInfo(id int, addr string) (*models.FairLaunchMintedInfo, error) {
	var fairLaunchMintedInfo models.FairLaunchMintedInfo
	response, err := api.GetDecodedAddrInfo(addr)
	if err != nil {
		utils.LogError("Decoded Addr error.", err)
		return nil, err
	}
	fairLaunchMintedInfo = models.FairLaunchMintedInfo{
		Model:            gorm.Model{},
		FairLaunchInfoID: id,
		EncodedAddr:      addr,
		AssetID:          hex.EncodeToString(response.AssetId),
		AssetType:        response.AssetType.String(),
		AddrAmount:       int(response.Amount),
		ScriptKey:        hex.EncodeToString(response.ScriptKey),
		InternalKey:      hex.EncodeToString(response.InternalKey),
		TaprootOutputKey: hex.EncodeToString(response.TaprootOutputKey),
		ProofCourierAddr: response.ProofCourierAddr,
		MintTime:         utils.GetTimestamp(),
		//deferred update
		Outpoint: "",
		Address:  "",
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
	MintTotal              int
	ActualMintTotalPercent float64
}

// TODO: need to test
// TODO: need to modify logic
// TODO: / and %
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
	mintPercent := 100 - reserved
	mintTotal := amount * mintPercent / 100
	for mintTotal%mintQuantity != 0 {
		mintTotal--
	}
	if mintTotal <= 0 || mintTotal < mintQuantity {
		return nil, errors.New("insufficient mint total amount")
	}
	ReservedTotal := amount - mintTotal
	if ReservedTotal <= 0 {
		return nil, errors.New("reserved amount is less equal than zero")
	}
	// mintTotal = mintQuantity * mintNumber
	mintNumber := mintTotal / mintQuantity
	if mintNumber <= 0 {
		return nil, errors.New("mint number is less equal than zero")
	}
	actualReserved := float64(ReservedTotal) * 100 / float64(amount)
	actualReserved = utils.RoundToDecimalPlace(actualReserved, 2)
	calculatedSeparateAmount := CalculateSeparateAmount{
		Amount:                 amount,
		Reserved:               reserved,
		ActualReserved:         actualReserved,
		ReserveTotal:           ReservedTotal,
		MintQuantity:           mintQuantity,
		MintNumber:             mintNumber,
		MintTotal:              mintTotal,
		ActualMintTotalPercent: 100 - actualReserved,
	}
	return &calculatedSeparateAmount, nil
}
