package routers

import (
	"github.com/gin-gonic/gin"
	"trade/handlers"
)

func setupFileServerRouter(router *gin.Engine) *gin.Engine {
	router.POST("/upload", handlers.FileUpload)
	router.GET("/download", handlers.FileDownload)
	return router
}
