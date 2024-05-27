package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trade/models"
	"trade/services"
	"trade/utils"
)

func GetAllFairLaunchInfo(c *gin.Context) {
	allFairLaunch, err := services.GetAllFairLaunchInfos()
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
	fairLaunch, err := services.GetFairLaunchInfo(id)
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
	minted, err := services.GetFairLaunchMintedInfosByFairLaunchId(id)
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
	// TODO: Use MustGet. alice ONLY FOR TEST
	//username := c.MustGet("username").(string)
	username := "alice"
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
	// @dev: Use SetFairLaunchInfoRequest c.ShouldBind
	var setFairLaunchInfoRequest models.SetFairLaunchInfoRequest
	err = c.ShouldBindJSON(&setFairLaunchInfoRequest)
	if err != nil {
		utils.LogError("Should Bind JSON setFairLaunchInfoRequest.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Should Bind JSON setFairLaunchInfoRequest. " + err.Error(),
			Data:    "",
		})
		return
	}
	imageData := setFairLaunchInfoRequest.ImageData
	name := setFairLaunchInfoRequest.Name
	assetType := setFairLaunchInfoRequest.AssetType
	amount := setFairLaunchInfoRequest.Amount
	reserved := setFairLaunchInfoRequest.Reserved
	mintQuantity := setFairLaunchInfoRequest.MintQuantity
	startTime := setFairLaunchInfoRequest.StartTime
	endTime := setFairLaunchInfoRequest.EndTime
	description := setFairLaunchInfoRequest.Description
	feeRate := setFairLaunchInfoRequest.FeeRate
	// @dev: Process struct, update later
	// @notice: State is 0 now
	fairLaunchInfo, err = services.ProcessFairLaunchInfo(imageData, name, assetType, amount, reserved, mintQuantity, startTime, endTime, description, feeRate, userId)
	if err != nil {
		utils.LogError("Process fair launch info.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Process fair launch info." + err.Error(),
			Data:    "",
		})
		return
	}
	// @dev: Update db, State models.FairLaunchStateNoPay
	err = services.SetFairLaunchInfo(fairLaunchInfo)
	if err != nil {
		utils.LogError("Set fair launch info.", err)
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
}

func SetFairLaunchMintedInfo(c *gin.Context) {
	var fairLaunchMintedInfo *models.FairLaunchMintedInfo
	var mintFairLaunchRequest models.MintFairLaunchRequest
	// @notice: only receive id and number
	err := c.ShouldBindJSON(&mintFairLaunchRequest)
	if err != nil {
		utils.LogError("Should Bind JSON mintFairLaunchRequest.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Should Bind JSON mintFairLaunchRequest. " + err.Error(),
			Data:    "",
		})
		return
	}
	// @dev: Ensure time is valid
	isTimeRight, err := services.IsFairLaunchMintTimeRight(mintFairLaunchRequest.FairLaunchInfoID)
	if err != nil {
		utils.LogError("Is FairLaunch Mint Time Right.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Is FairLaunch Mint Time Right. " + err.Error(),
			Data:    "",
		})
		return
	}
	if !isTimeRight {
		utils.LogError("It is not Right FairLaunch Mint Time now.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "It is not Right FairLaunch Mint Time now. " + err.Error(),
			Data:    "",
		})
		return
	}
	// TODO: Use MustGet. bob ONLY FOR TEST
	username := "bob"
	//username := c.MustGet("username").(string)
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
	fairLaunchInfoID := mintFairLaunchRequest.FairLaunchInfoID
	mintedNumber := mintFairLaunchRequest.MintedNumber
	addr := mintFairLaunchRequest.EncodedAddr
	fairLaunchMintedInfo, err = services.ProcessFairLaunchMintedInfo(fairLaunchInfoID, mintedNumber, addr, userId)
	if err != nil {
		utils.LogError("Process FairLaunchMintedInfo.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Process FairLaunchMintedInfo " + err.Error(),
			Data:    "",
		})
		return
	}
	// @dev: Update db, State models.FairLaunchMintedStateNoPay
	err = services.SetFairLaunchMintedInfo(fairLaunchMintedInfo)
	if err != nil {
		utils.LogError("Set fair launch minted info.", err)
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "Set fair launch minted info. " + err.Error(),
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

// TODO: add more query

func QueryMintIsAvailable(c *gin.Context) {
	// TODO: Check if a specific id and amount of minting asset is valid

}

func MintFairLaunchReserved(c *gin.Context) {
	// TODO: need to complete

}
