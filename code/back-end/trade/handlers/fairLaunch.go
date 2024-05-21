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
		c.JSON(http.StatusOK, utils.MakeJsonResult(false, "Can not get all fair launch infos. "+err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.MakeJsonResult(true, "", allFairLaunch))
}

func GetFairLaunchInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.LogError("id is not valid int", err)
		c.JSON(http.StatusOK, utils.MakeJsonResult(false, "id is not valid int. "+err.Error(), ""))
		return
	}
	fairLaunch, err := services.GetFairLaunch(id)
	if err != nil {
		utils.LogError("Get fair launch info", err)
		c.JSON(http.StatusOK, utils.MakeJsonResult(false, "Can not get fair launch info. "+err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.MakeJsonResult(true, "", fairLaunch))
}

func GetMintedInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.LogError("id is not valid int", err)
		c.JSON(http.StatusOK, utils.MakeJsonResult(false, "id is not valid int. "+err.Error(), ""))
		return
	}
	minted, err := services.GetMinted(id)
	if err != nil {
		utils.LogError("Get fair launch minted info", err)
		c.JSON(http.StatusOK, utils.MakeJsonResult(false, "Can not get fair launch minted info. "+err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.MakeJsonResult(true, "", minted))

}

func SetFairLaunchInfo(c *gin.Context) {
	var fairLaunchInfo models.FairLaunchInfo
	err := c.ShouldBind(&fairLaunchInfo)
	if err != nil {
		utils.LogError("Wrong json to bind.", err)
		c.JSON(http.StatusOK, utils.MakeJsonResult(false, "Wrong json to bind. "+err.Error(), ""))
		return
	}
	err = services.SetFairLaunch(&fairLaunchInfo)
	if err != nil {
		utils.LogError("Set fair launch error.", err)
		c.JSON(http.StatusOK, utils.MakeJsonResult(false, "Set fair launch error. "+err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.MakeJsonResult(true, "", nil))
}

func MintFairLaunch(c *gin.Context) {
	// TODO: receive mint info
	// TODO: Pay Fee
	// TODO: call services FairLaunchMint
	// TODO:

	// TODO: need to complete
	utils.LogInfo("MintFairLaunch triggered. This function did nothing, need to complete.")

}
