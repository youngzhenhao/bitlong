package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	// @dev: Update db to State models.FairLaunchStateNoPay
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
}

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
