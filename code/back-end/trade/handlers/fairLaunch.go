package handlers

import (
	"github.com/gin-gonic/gin"
	"trade/services"
)

func GetFairLaunchInfo(c *gin.Context) {
	id := c.Param("id")
	_ = id
	services.GetFairLaunch()
	// TODO: need to complete

}

func GetMintedInfo(c *gin.Context) {
	id := c.Param("id")
	_ = id
	services.GetMinted()
	// TODO: need to complete

}

func SetFairLaunchInfo(c *gin.Context) {
	c.PostForm("asset_id")
	// TODO: need to complete
	services.SetFairLaunch()

}

func MintFairLaunch(c *gin.Context) {
	c.PostForm("asset_id")
	// TODO: need to complete
	services.FairLaunchMint()

}
