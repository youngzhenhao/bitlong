package routers

import (
	"github.com/gin-gonic/gin"
	"trade/config"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	if !config.GetLoadConfig().RouterBlock.Login {
		SetupLoginRouter(r)
	}
	if !config.GetLoadConfig().RouterBlock.FairLaunch {
		setupFairLaunchRouter(r)
	}
	if !config.GetLoadConfig().RouterBlock.Fee {
		SetupFeeRouter(r)
	}
	if !config.GetLoadConfig().RouterBlock.CustodyAccount {
		SetupCustodyAccountRouter(r)
	}
	if !config.GetLoadConfig().RouterBlock.Ping {
		SetupPingRouter(r)
	}
	return r
}
