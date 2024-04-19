package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileServerRouter() {
	router := setupFileServerRouter()
	err := router.Run("0.0.0.0:10080")
	if err != nil {
		return
	}
}

// TODO: simple file upload and download service
func setupFileServerRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		result := true
		c.JSON(http.StatusOK, gin.H{
			"time":   GetTimeNow(),
			"ping":   "pong",
			"result": result,
		})
	})

	return router
}
