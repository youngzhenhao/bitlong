package routers

import (
	"github.com/gin-gonic/gin"
	"trade/handlers"
)

func setupFairLaunchRouter(router *gin.Engine) *gin.Engine {
	fairLaunch := router.Group("/fair_launch")
	// TODO: Enable AuthMiddleware
	//fairLaunch.Use(middleware.AuthMiddleware())
	{
		// TODO: Add and modify routers
		fairLaunch.GET("/all", handlers.GetAllFairLaunchInfo)
		fairLaunch.GET("/info/:id", handlers.GetFairLaunchInfo)
		fairLaunch.GET("/minted/:id", handlers.GetMintedInfo)
		fairLaunch.POST("/set", handlers.SetFairLaunchInfo)

		fairLaunch.POST("/mint", handlers.SetFairLaunchMintedInfo)

		fairLaunch.POST("/mint_reserved", handlers.MintFairLaunchReserved)

		query := fairLaunch.Group("/query")
		{
			query.POST("/:id", handlers.QueryInventory)
			query.POST("/mint", handlers.QueryMintIsAvailable)
		}

	}
	return router
}
