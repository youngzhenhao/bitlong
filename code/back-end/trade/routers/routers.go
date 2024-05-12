package routers

import (
	"AssetsTrade/config"
	"github.com/gin-gonic/gin"
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
