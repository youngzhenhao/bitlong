package routers

import (
	"github.com/gin-gonic/gin"
	"trade/config"
	"trade/handlers"
)

func setupFileServerRouter(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		config.GetLoadConfig().BasicAuth[0].Username: config.GetLoadConfig().BasicAuth[0].Password,
	}))
	authorized.POST("/upload", handlers.FileUpload)
	authorized.GET("/download", handlers.FileDownload)
	return router
}
