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
	name := c.PostForm("name")
	amountStr := c.PostForm("amount")
	reservedStr := c.PostForm("reserved")
	mintQuantityStr := c.PostForm("mint_quantity")
	startTimeStr := c.PostForm("start_time")
	endTimeStr := c.PostForm("end_time")
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
	fairLaunchInfo, err = services.ProcessFairLaunchInfo(name, amount, reserved, mintQuantity, startTime, endTime)
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
	// TODO: get batch key
	//	 TODO: update batch state
	//  TODO: get batch txid

	//  TODO:  Use BatchTxidAnchorToAssetId to update

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
