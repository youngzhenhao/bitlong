package lnurl

import (
	"github.com/gin-gonic/gin"
	"github.com/wallet/api"
	"net/http"
	"strconv"
)

func LnurlRouterRun() {
	router := setupRouter()
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("lnurl/page/pay.html")

	router.GET("/pay", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pay.html", gin.H{
			"title": "pay",
		})
	})
	//  请求服务器的LND节点开发票
	router.POST("/addInvoice", func(c *gin.Context) {
		amountStr := c.PostForm("amount")
		amountInt, _ := strconv.Atoi(amountStr)
		invoiceStr := api.AddInvoice(int64(amountInt))
		c.JSON(http.StatusOK, gin.H{
			"invoice": invoiceStr,
		})
	})

	channelGroup := router.Group("/channel")
	{
		channelGroup.GET("/GetChanInfo", func(c *gin.Context) {
			chanIdStr, _ := strconv.Atoi(c.Query("chanId"))
			c.JSON(http.StatusOK, gin.H{
				"GetChanInfo": api.GetChanInfo(chanIdStr),
			})
		})
		channelGroup.GET("/ListChannels", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ListChannels": api.ListChannels(),
			})
		})
		channelGroup.GET("/ClosedChannels", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ClosedChannels": api.ClosedChannels(),
			})
		})
		channelGroup.GET("/PendingChannels", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"PendingChannels": api.PendingChannels(),
			})
		})
	}
	router.GET("/WatchtowerGetInfo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"GetInfo": api.WatchtowerGetInfo(),
		})
	})
	router.GET("/GetState", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"GetState": api.GetState(),
		})
	})

	walletKitGroup := router.Group("/walletKit")
	{
		walletKitGroup.GET("/ListAddress", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ListAddress": api.ListAddress(),
			})
		})

		walletKitGroup.GET("/ListAccounts", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ListAccounts": api.ListAccounts(),
			})
		})
		walletKitGroup.GET("/ListLeases", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ListLeases": api.ListLeases(),
			})
		})
		walletKitGroup.GET("/ListSweeps", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ListSweeps": api.ListSweeps(),
			})
		})
		walletKitGroup.GET("/ListUnspent", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ListUnspent": api.ListUnspent(),
			})
		})
		walletKitGroup.GET("/NextAddr", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"NextAddr": api.NextAddr(),
			})
		})

	}

	return router
}
