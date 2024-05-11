package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	SetupLoginRouter(r)
	setupFileServerRouter(r)
	setupFairLaunchRouter(r)
	return r
}
