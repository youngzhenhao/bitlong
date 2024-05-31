package routers

import (
	"github.com/gin-gonic/gin"
	"trade/handlers"
	"trade/middleware"
)

func SetupFeeRouter(router *gin.Engine) *gin.Engine {
	fee := router.Group("/fee")
	fee.Use(middleware.AuthMiddleware())
	{
		query := fee.Group("/query")
		{
			query.GET("/rate", handlers.QueryFeeRate)
		}
	}
	return router
}
