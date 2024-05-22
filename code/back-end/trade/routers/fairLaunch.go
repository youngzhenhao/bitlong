package routers

import (
	"github.com/gin-gonic/gin"
	"trade/handlers"
	"trade/middleware"
)

func setupFairLaunchRouter(router *gin.Engine) *gin.Engine {
	fairLaunch := router.Group("/fair_launch")
	fairLaunch.Use(middleware.AuthMiddleware())
	{
		fairLaunch.GET("/all", handlers.GetAllFairLaunchInfo)
		fairLaunch.GET("/info/:id", handlers.GetFairLaunchInfo)
		fairLaunch.GET("/minted/:id", handlers.GetMintedInfo)
		fairLaunch.POST("/set", handlers.SetFairLaunchInfo)
		// TODO: add query amount available
		fairLaunch.POST("/mint", handlers.MintFairLaunch)
	}
	return router
}
