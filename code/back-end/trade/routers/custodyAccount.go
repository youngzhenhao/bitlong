package routers

import (
	"github.com/gin-gonic/gin"
	"trade/handlers"
)

func SetupCustodyAccountRouter(router *gin.Engine) *gin.Engine {

	// A routing group that requires authentication
	custody := router.Group("/custodyAccount")

	custody.GET("/create", handlers.CreateCustodyAccount)

	Invoice := custody.Group("/invoice")
	Invoice.GET("/apply", handlers.ApplyInvoiceCA)
	Invoice.GET("/pay", handlers.PayInvoice)

	//custody.Use(middleware.AuthMiddleware())
	//{
	//}
	return router
}
