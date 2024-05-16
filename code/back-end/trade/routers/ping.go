package routers

import (
	"github.com/gin-gonic/gin"
	"trade/config"
	"trade/handlers"
)

func SetupPingRouter(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		config.GetLoadConfig().AdminUsers[0].Username: config.GetLoadConfig().AdminUsers[0].Password,
	}))
	authorized.GET("/ping", handlers.PingHandler)
	return router
}
