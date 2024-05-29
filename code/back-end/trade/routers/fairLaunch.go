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
		fairLaunch.POST("/set", handlers.SetFairLaunchInfo)
		fairLaunch.POST("/mint", handlers.SetFairLaunchMintedInfo)
		fairLaunch.POST("/mint_reserved/:id", handlers.MintFairLaunchReserved)
		query := fairLaunch.Group("/query")
		{
			query.GET("/all", handlers.GetAllFairLaunchInfo)
			query.GET("/info/:id", handlers.GetFairLaunchInfo)
			query.GET("/minted/:id", handlers.GetMintedInfo)
			query.GET("/inventory/:id", handlers.QueryInventory)
			query.GET("/mint", handlers.QueryMintIsAvailable)
		}
	}
	return router
}
