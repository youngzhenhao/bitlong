package routers

import (
	"github.com/gin-gonic/gin"
	"trade/handlers"
	"trade/middleware"
)

func SetupCustodyAccountRouter(router *gin.Engine) *gin.Engine {

	// A routing group that requires authentication
	custody := router.Group("/custodyAccount")

	custody.Use(middleware.AuthMiddleware())
	{
		custody.POST("/create", handlers.CreateCustodyAccount)
		Invoice := custody.Group("/invoice")
		{
			Invoice.POST("/apply", handlers.ApplyInvoiceCA)
			Invoice.POST("/pay", handlers.PayInvoice)
			Invoice.POST("/querybalance", handlers.QueryBalance)
		}
	}
	return router
}
