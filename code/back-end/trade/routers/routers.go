package routers

import (
	"github.com/gin-gonic/gin"
	"trade/config"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	SetupLoginRouter(r)
	if config.GetLoadConfig().Routers.FairLaunch {
		setupFairLaunchRouter(r)
	}
	if config.GetLoadConfig().Routers.Ping {
		SetupPingRouter(r)
	}
	SetupCustodyAccountRouter(r)
	return r
}
