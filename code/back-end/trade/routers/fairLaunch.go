package routers

import (
	"github.com/gin-gonic/gin"
	"trade/handlers"
)

//	 TODO:	Set fair launch
//			Query fair launch info
//			Do fair launch Mint
//			Query fair launch minted info
//			Consider lock time and relative lock time
func setupFairLaunchRouter(router *gin.Engine) *gin.Engine {
	fairLaunch := router.Group("/fair_launch")
	{
		fairLaunch.GET("/all", handlers.GetAllFairLaunchInfo)
		fairLaunch.GET("/info/:id", handlers.GetFairLaunchInfo)
		fairLaunch.GET("/minted/:id", handlers.GetMintedInfo)
		fairLaunch.POST("/set", handlers.SetFairLaunchInfo)
		fairLaunch.POST("/mint", handlers.MintFairLaunch)
	}
	return router
}
