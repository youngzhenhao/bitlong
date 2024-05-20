package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trade/cmd-test/deprecated"
	"trade/models"
)

func GetFairLaunchInfo(c *gin.Context) {
	id := c.Param("id")
	fairLaunch := deprecated.GetFairLaunch(id)
	if fairLaunch == nil {
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "not found fair launch info or error occurs.",
			Data:    nil,
		})
	} else {
		c.JSON(http.StatusOK, models.JsonResult{
			Success: true,
			Error:   "",
			Data:    fairLaunch,
		})
	}
}

func GetMintedInfo(c *gin.Context) {
	id := c.Param("id")
	mintedInfo := deprecated.GetMinted(id)
	if mintedInfo == nil {
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "not found fair launch info or fair launch no minted info or error occurs.",
			Data:    nil,
		})
	} else {
		c.JSON(http.StatusOK, models.JsonResult{
			Success: true,
			Error:   "",
			Data:    mintedInfo,
		})
	}
}

func SetFairLaunchInfo(c *gin.Context) {
	c.PostForm("asset_id")
	// TODO: need to complete
	deprecated.SetFairLaunch()

}

func MintFairLaunch(c *gin.Context) {
	c.PostForm("asset_id")
	// TODO: need to complete
	deprecated.FairLaunchMint()

}

func GetAllFairLaunchInfo(c *gin.Context) {
	allFairLaunch := deprecated.GetAllFairLaunch()
	if allFairLaunch == nil {
		c.JSON(http.StatusOK, models.JsonResult{
			Success: false,
			Error:   "null fair launch info or error occurs.",
			Data:    nil,
		})
	} else {
		c.JSON(http.StatusOK, models.JsonResult{
			Success: true,
			Error:   "",
			Data:    allFairLaunch,
		})
	}
}
