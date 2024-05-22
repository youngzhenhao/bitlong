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
		fairLaunch.POST("/mint", handlers.MintFairLaunch)

		query := fairLaunch.Group("/query")
		{
			query.POST("/mint", handlers.QueryMintIsAvailable)
		}

	}
	return router
}
