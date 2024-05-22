package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
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
	// TODO: 1.Query info
	fairLaunchInfo, err := GetFairLaunch(fairLaunchMintedInfo.FairLaunchInfoID)
	if err != nil {
		utils.LogError("Get fair launch by id of fairLaunchMintedInfo", err)
		return err
	}
	// TODO: 2.Calculate number of asset
	amt := calculateAmount(int(fairLaunchInfo.ID), fairLaunchMintedInfo.AddrAmount)
	if amt == 0 {
		err = errors.New("amount of asset to send is zero")
		utils.LogError("", err)
		return err
	}
	// TODO: 3.Request an addr
	addr := GetAddr(fairLaunchInfo.AssetID, amt)
	// TODO: 4.Pay asset to addr
	result, err := api.SendAssetBool(addr, 0)
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

	// TODO: 5.Write to database of asset release

	return nil
}

func GetAddr(id string, amt int) string {
	// TODO: need to complete
	// TODO: Request User to generate a addr of receiving asset
	// TODO: Response struct
	// TODO: Call PostPhoneToNewAddr
	utils.LogInfo("GetAddr triggered. This function did nothing, need to complete.")
	return ""
}

func PostPhoneToNewAddr(remotePort string, assetId string, amount int) *taprpc.Addr {
	frpsForwardSocket := fmt.Sprintf("%s:%s", config.GetLoadConfig().FrpsServer, remotePort)
	targetUrl := "http://" + frpsForwardSocket + "/newAddr"
	payload := url.Values{"asset_id": {assetId}, "amount": {strconv.Itoa(amount)}}
	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		utils.LogError("http.PostForm.", err)
		return nil
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var addrResponse struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Data    *taprpc.Addr
	}
	if err := json.Unmarshal(bodyBytes, &addrResponse); err != nil {
		utils.LogError("PPTNA json.Unmarshal.", err)
		return nil
	}
	return addrResponse.Data
}

func calculateAmount(id int, amount int) int {
	// TODO: need to complete
	// TODO: add number logic, or judge number when mint
	utils.LogInfo("calculateAmount triggered. This function did nothing, need to complete.")
	return 0
}

func ProcessFairLaunchMintedInfo(fairLaunchMintedInfo *models.FairLaunchMintedInfo) *models.FairLaunchMintedInfo {
	// TODO: need to complete
	utils.LogInfo("ProcessFairLaunchMintedInfo triggered. This function did nothing, need to complete.")
	return nil
}
