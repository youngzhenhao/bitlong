package routers

import (
	"github.com/gin-gonic/gin"
	"trade/config"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	if config.GetConfig().Routers.Login {
		SetupLoginRouter(r)
	}
	if config.GetConfig().Routers.FileServer {
		setupFileServerRouter(r)
	}
	if config.GetConfig().Routers.FairLaunch {
		setupFairLaunchRouter(r)
	}
	return r
}
