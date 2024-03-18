package lnurl

import (
	"github.com/gin-gonic/gin"
	"github.com/wallet/api"
	"net/http"
	"strconv"
)

func LnurlRouterRun() {
	router := setupRouter()
	router.Run(":8888")
}

// setupRouter
//
//	@Description: 请求服务器的LND节点开发票
//	@return *gin.Engine
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/pay", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pay.html", gin.H{
			"title": "pay",
		})
	})

	router.POST("/addInvoice", func(c *gin.Context) {
		amountStr := c.PostForm("amount")
		amountInt, _ := strconv.Atoi(amountStr)
		invoiceStr := api.AddInvoice(int64(amountInt)).PaymentRequest
		c.JSON(http.StatusOK, gin.H{
			"invoice": invoiceStr,
		})
	})
	return router
}
