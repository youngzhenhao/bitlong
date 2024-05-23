package handlers

import (
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"net/http"
	"strconv"
	"trade/api"
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
	utils.LogInfos("Batch state:", batchState)
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
	utils.LogInfos("Batch state:", batchState)
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
	fairLaunchInfo, err = services.ProcessFairLaunchInfo(imageData, name, assetType, amount, reserved, mintQuantity, startTime, endTime, description, batchKey, batchState, batchTxidAnchor, assetId, userId)
	if err != nil {
		utils.LogError("Process fair launch info.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Process fair launch info." + err.Error(),
			Data:    "",
		})
		return
	}
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
	c.JSON(http.StatusOK, models.JsonResult{
		Success: true,
		Error:   "",
		Data:    nil,
	})
	// TODO: update FairLaunchInfo later
}

func MintFairLaunch(c *gin.Context) {
	var fairLaunchMintedInfo *models.FairLaunchMintedInfo
	// @dev: Get id and addr
	idStr := c.PostForm("id")
	addr := c.PostForm("addr")
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
	fairLaunchMintedInfo, _ = services.ProcessFairLaunchMintedInfo(id, addr)
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
	// TODO: 1.Pay Fee
	amount := fairLaunchMintedInfo.AddrAmount
	fee, err := services.CalculateFee(amount)
	if err != nil {
		utils.LogError("Calculate fee error", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Calculate fee error. " + err.Error(),
			Data:    "",
		})
		return
	}
	limit := services.FeeLimit
	username := c.MustGet("username").(string)
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
	err = services.PayFee(fee, limit, userId)
	if err != nil {
		utils.LogError("Pay fee error", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Pay fee error. " + err.Error(),
			Data:    "",
		})
		return
	}

	// TODO: 2.Fair launch mint service
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
	c.JSON(http.StatusOK, models.JsonResult{
		Success: true,
		Error:   "",
		Data:    nil,
	})
}

func QueryMintIsAvailable(c *gin.Context) {
	// TODO: Check if a specific amount of minting asset is valid

}
