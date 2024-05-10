package routers

import (
	"AssetsTrade/handlers"
	"github.com/gin-gonic/gin"
)

func setupFileServerRouter(router *gin.Engine) *gin.Engine {
	router.POST("/upload", handlers.FileUpload)
	router.GET("/download", handlers.FileDownload)
	return router
}
