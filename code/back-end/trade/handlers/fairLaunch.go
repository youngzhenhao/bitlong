package handlers

import (
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"net/http"
	"strconv"
	"trade/api"
	"trade/middleware"
	"trade/models"
	"trade/services"
	"trade/utils"
)

func GetAllFairLaunchInfo(c *gin.Context) {
	allFairLaunch, err := services.GetAllFairLaunch()
	if err != nil {
		utils.LogError("Get all fair launch infos", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Can not get all fair launch infos. " + err.Error(),
			Data:    "",
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResult{
		Success: true,
		Error:   "",
		Data:    allFairLaunch,
	})
}

func GetFairLaunchInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.LogError("id is not valid int", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "id is not valid int. " + err.Error(),
			Data:    "",
		})
		return
	}
	fairLaunch, err := services.GetFairLaunch(id)
	if err != nil {
		utils.LogError("Get fair launch info", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Can not get fair launch info. " + err.Error(),
			Data:    "",
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResult{
		Success: true,
		Error:   "",
		Data:    fairLaunch,
	})
}

func GetMintedInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.LogError("id is not valid int", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "id is not valid int. " + err.Error(),
			Data:    "",
		})
		return
	}
	minted, err := services.GetMinted(id)
	if err != nil {
		utils.LogError("Get fair launch minted info", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Can not get fair launch minted info. " + err.Error(),
			Data:    "",
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResult{
		Success: true,
		Error:   "",
		Data:    minted,
	})
}

func SetFairLaunchInfo(c *gin.Context) {
	var fairLaunchInfo *models.FairLaunchInfo
	// TODO: Use SetFairLaunchInfoRequest c.ShouldBind
	imageData := c.PostForm("image_data")
	name := c.PostForm("name")
	assetTypeStr := c.PostForm("asset_type")
	amountStr := c.PostForm("amount")
	reservedStr := c.PostForm("reserved")
	mintQuantityStr := c.PostForm("mint_quantity")
	startTimeStr := c.PostForm("start_time")
	endTimeStr := c.PostForm("end_time")
	description := c.PostForm("description")
	feeRateStr := c.PostForm("fee_rate")
	username := c.MustGet("username").(string)
	releaseFeeInvoice := c.PostForm("release_fee_invoice")
	userId, err := services.NameToId(username)
	if err != nil {
		utils.LogError("Query user id by name.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Query user id by name." + err.Error(),
			Data:    "",
		})
		return
	}
	assetType, err := strconv.Atoi(assetTypeStr)
	if err != nil {
		utils.LogError("strconv string to int.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "strconv string to int." + err.Error(),
			Data:    "",
		})
		return
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		utils.LogError("strconv string to int.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "strconv string to int." + err.Error(),
			Data:    "",
		})
		return
	}
	reserved, err := strconv.Atoi(reservedStr)
	if err != nil {
		utils.LogError("strconv string to int.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "strconv string to int." + err.Error(),
			Data:    "",
		})
		return
	}
	mintQuantity, err := strconv.Atoi(mintQuantityStr)
	if err != nil {
		utils.LogError("strconv string to int.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "strconv string to int." + err.Error(),
			Data:    "",
		})
		return
	}
	startTime, err := strconv.Atoi(startTimeStr)
	if err != nil {
		utils.LogError("strconv string to int.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "strconv string to int." + err.Error(),
			Data:    "",
		})
		return
	}
	endTime, err := strconv.Atoi(endTimeStr)
	if err != nil {
		utils.LogError("strconv string to int.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "strconv string to int." + err.Error(),
			Data:    "",
		})
		return
	}
	feeRate, err := strconv.Atoi(feeRateStr)
	if err != nil {
		utils.LogError("strconv string to int.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "strconv string to int." + err.Error(),
			Data:    "",
		})
		return
	}
	// TODO: add judge logic

	// TODO: check release fee paid
	isReleaseFeePaid := services.IsReleaseFeePaid(releaseFeeInvoice)
	if !isReleaseFeePaid {
		err = errors.New("release fee did not been paid")
		utils.LogError("", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "" + err.Error(),
			Data:    "",
		})
		return
	}
	//assetTypeQuery, _ := api.QueryAssetType(assetType)
	var isCollectible bool
	if taprpc.AssetType(assetType) == taprpc.AssetType_COLLECTIBLE {
		isCollectible = true
	}
	newMeta := api.NewMeta(description, imageData)
	mintResponse, err := api.MintAssetAndGetResponse(name, isCollectible, newMeta, amount, false)
	if err != nil {
		utils.LogError("Mint asset.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Mint asset." + err.Error(),
			Data:    "",
		})
		return
	}
	batchKey := hex.EncodeToString(mintResponse.GetPendingBatch().GetBatchKey())
	batchState := mintResponse.GetPendingBatch().GetState().String()
	//utils.LogInfos("Batch state:", batchState)
	finalizeResponse, err := api.FinalizeBatchAndGetResponse(feeRate)
	if err != nil {
		utils.LogError("Finalize batch.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Finalize batch." + err.Error(),
			Data:    "",
		})
		return
	}
	if hex.EncodeToString(finalizeResponse.GetBatch().GetBatchKey()) != batchKey {
		err = errors.New("finalize batch key is not equal mint batch key")
		utils.LogError("", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   err.Error(),
			Data:    "",
		})
		return
	}
	batchTxidAnchor := finalizeResponse.GetBatch().GetBatchTxid()
	batchState = finalizeResponse.GetBatch().GetState().String()
	//utils.LogInfos("Batch state:", batchState)
	assetId, err := api.BatchTxidAnchorToAssetId(batchTxidAnchor)
	if err != nil {
		utils.LogError("Batch Anchor Txid To AssetId.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Batch Anchor Txid To AssetId." + err.Error(),
			Data:    "",
		})
		return
	}
	// @dev: Process struct
	fairLaunchInfo, err = services.ProcessFairLaunchInfo(imageData, name, assetType, amount, reserved, mintQuantity, startTime, endTime, description, batchKey, batchState, batchTxidAnchor, assetId, userId, releaseFeeInvoice)
	if err != nil {
		utils.LogError("Process fair launch info.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Process fair launch info." + err.Error(),
			Data:    "",
		})
		return
	}
	// @dev: Update dbs
	err = services.SetFairLaunch(fairLaunchInfo)
	if err != nil {
		utils.LogError("Set fair launch error.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Set fair launch error. " + err.Error(),
			Data:    "",
		})
		return
	}
	// @dev: update inventory
	err = services.CreateInventoryInfoByFairLaunchInfo(fairLaunchInfo)
	if err != nil {
		utils.LogError("Create Inventory Info By FairLaunchInfo.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Create Inventory Info By FairLaunchInfo. " + err.Error(),
			Data:    "",
		})
		return
	}
	// @dev: Update asset_releases
	err = services.CreateAssetReleaseInfoByFairLaunchInfo(fairLaunchInfo)
	if err != nil {
		utils.LogError("Create Asset Release Info By FairLaunchInfo.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Create Asset Release Info By FairLaunchInfo. " + err.Error(),
			Data:    "",
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResult{
		Success: true,
		Error:   "",
		Data:    fairLaunchInfo,
	})
	// TODO: update FairLaunchInfo later
	//		IsReservedSent
	//		MintedNumber
	//		Status
}

// @Previous: before this operation, user send request to check avaible amount and number of inventory and modify inverntory status
// TODO: User should finish paying mint fee before this operation

// TODO: use scheduled task
func MintFairLaunch(c *gin.Context) {
	var fairLaunchMintedInfo *models.FairLaunchMintedInfo
	// @dev: Get id, addr and pay-fee-invoice
	idStr := c.PostForm("id")
	mintFeeInvoice := c.PostForm("mint_fee_invoice")
	addr := c.PostForm("addr")
	fairLaunchInfoID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.LogError("id is not valid int", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "id is not valid int. " + err.Error(),
			Data:    "",
		})
		return
	}
	fairLaunchMintedInfo, _ = services.ProcessFairLaunchMintedInfo(fairLaunchInfoID, addr, mintFeeInvoice)
	// @dev: ShouldBind
	//if err != nil {
	//	utils.LogError("Wrong json to bind.", err)
	//	c.JSON(http.StatusOK, models.JsonResult{
	//		Success: false,
	//		Error:   "Wrong json to bind. " + err.Error(),
	//		Data:    "",
	//	})
	//	return
	//}

	// @previous: 1.Pay Fee

	//amount := fairLaunchMintedInfo.AddrAmount

	// @previous: CalculatedFee
	//_, err = services.CalculateFee(amount)
	//if err != nil {
	//	utils.LogError("Calculate fee error", err)
	//	c.JSON(http.StatusOK, models.JsonResult{
	//		Success: false,
	//		Error:   "Calculate fee error. " + err.Error(),
	//		Data:    "",
	//	})
	//	return
	//}

	username := c.MustGet("username").(string)
	// @dev: userId
	userId, err := services.NameToId(username)
	if err != nil {
		utils.LogError("Name to id error", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Name to id error. " + err.Error(),
			Data:    "",
		})
		return
	}
	// @dev: check whether mint fee is paid
	isMintFeePaid := services.IsMintFeePaid(mintFeeInvoice)
	if !isMintFeePaid {
		err = errors.New("mint fee did not been paid")
		utils.LogError("", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "" + err.Error(),
			Data:    "",
		})
		return
	}
	var fairLaunchMintedUserInfo models.FairLaunchMintedUserInfo
	err = middleware.DB.Where("user_id = ? AND fair_launch_info_id =? AND status = ?", userId, fairLaunchInfoID, models.StatusNormal).First(&fairLaunchMintedUserInfo).Error
	var noMintedUserInfo bool
	if err != nil {
		utils.LogError("select fair launch minted user info.", err)
		//c.JSON(http.StatusOK, models.JsonResult{
		//	Success: false,
		//	Error:   "select fair launch minted user info. " + err.Error(),
		//	Data:    "",
		//})
		noMintedUserInfo = true
	}
	fairLaunchInfo, err := services.GetFairLaunch(fairLaunchInfoID)
	if err != nil {
		utils.LogError("Get fair launch info.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Get fair launch info. " + err.Error(),
			Data:    "",
		})
	}
	mintNumber := services.AmountAndQuantityToNumber(fairLaunchMintedInfo.AddrAmount, fairLaunchInfo.MintQuantity)
	mintedNumber := fairLaunchMintedUserInfo.MintedNumber
	if mintNumber > models.MintMaxNumber-mintedNumber {
		err = errors.New("mint number out of max")
		utils.LogError("", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "" + err.Error(),
			Data:    "",
		})
	}
	// TODO: Fair launch mint service
	err = services.FairLaunchMint(fairLaunchMintedInfo)
	if err != nil {
		utils.LogError("Fair launch mint error.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Fair launch mint error. " + err.Error(),
			Data:    "",
		})
		return
	}
	// TODO: update db
	//		minted
	//		minted user
	if noMintedUserInfo {
		// TODO: Create
	} else {
		// TODO: update
		//err = middleware.DB.Where().Update()
	}

	c.JSON(http.StatusOK, models.JsonResult{
		Success: true,
		Error:   "",
		Data:    nil,
	})
	// TODO: update db status later

}

func QueryMintIsAvailable(c *gin.Context) {
	// TODO: Check if a specific amount of minting asset is valid
	//id := c.PostForm("id")
	//amount := c.PostForm("amount")

}

func MintFairLaunchReserved(c *gin.Context) {
	// TODO: need to complete

}

func QueryInventory(c *gin.Context) {
	// call GetNumberOfInventoryCouldBeMinted
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.LogError("strconv string to int.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "strconv string to int." + err.Error(),
			Data:    "",
		})
		return
	}
	inventory, err := services.GetInventoryCouldBeMintedByFairLaunchInfoId(id)
	if err != nil {
		utils.LogError("Get inventory could be minted by fair launch info id.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Get inventory could be minted by fair launch info id." + err.Error(),
			Data:    "",
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResult{
		Success: true,
		Error:   "",
		Data:    inventory,
	})
}
