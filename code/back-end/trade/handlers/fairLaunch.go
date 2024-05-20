package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetFairLaunchInfo(c *gin.Context) {
	// TODO: need to complete
}

func GetMintedInfo(c *gin.Context) {
	// TODO: need to complete

}

func SetFairLaunchInfo(c *gin.Context) {
	c.PostForm("asset_id")
	// TODO: need to complete

}

func MintFairLaunch(c *gin.Context) {
	c.PostForm("asset_id")
	// TODO: need to complete

}

func GetAllFairLaunchInfo(c *gin.Context) {
	// TODO: need to complete

}
