package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wallet/api"
	"io"
	"net/http"
	"os"
	"path"
)

func RunFileServerRouter() {
	router := setupFileServerRouter()
	err := router.Run("0.0.0.0:6080")
	if err != nil {
		return
	}
}

func setupFileServerRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/upload", UploadFile)
	router.GET("/download", FileDownload)
	return router
}

// UploadFile
// @Description: Upload a single file
// @param c
func UploadFile(c *gin.Context) {
	//fileName := c.Query("filename")
	saveName := uuid.New().String()
	saveDir := "files"
	file, header, err := c.Request.FormFile("file")
	result := true
	if err != nil {
		result = false
		fmt.Printf("%s %v\n", api.GetTimeNow(), err)
		c.JSON(http.StatusOK, gin.H{
			"time":   api.GetTimeNow(),
			"result": result,
			"info":   err,
		})
		return
	}
	uploadFileNameWithSuffix := path.Base(header.Filename)
	//uploadFileType := path.Ext(uploadFileNameWithSuffix)
	//saveName = fileName + uploadFileType
	saveName = uploadFileNameWithSuffix
	savePath := saveDir + "/" + saveName
	var localFileInfo os.FileInfo
	localFileInfo, err = os.Stat(saveDir)
	if err != nil || !localFileInfo.IsDir() {
		err := os.MkdirAll(saveDir, 0755)
		if err != nil {
			fmt.Printf("%s mkdir %s error %v\n", api.GetTimeNow(), saveDir, err)
			result = false
			c.JSON(http.StatusOK, gin.H{
				"time":   api.GetTimeNow(),
				"result": result,
				"info":   err,
			})
			return
		}
	}
	out, err := os.Create(savePath)
	if err != nil {
		fmt.Printf("%s create file %s error %v\n", api.GetTimeNow(), savePath, err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Printf("%s close file %s error %v\n", api.GetTimeNow(), savePath, err)
		}
	}(out)
	_, err = io.Copy(out, file)
	if err != nil {
		result = false
		c.JSON(http.StatusOK, gin.H{
			"time":   api.GetTimeNow(),
			"result": result,
			"info":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"time":   api.GetTimeNow(),
		"result": result,
		"info":   "uploaded successfully",
	})
	return
}

func FileDownload(c *gin.Context) {
	filePath := "files/" + c.Query("file")
	fileTmp, err := os.Open(filePath)
	defer func(fileTmp *os.File) {
		err := fileTmp.Close()
		if err != nil {
			fmt.Printf("%s close file error %v\n", api.GetTimeNow(), err)
		}
	}(fileTmp)
	fileName := path.Base(filePath)
	if filePath == "" || fileName == "" || err != nil {
		fmt.Printf("%s file not found %v\n", api.GetTimeNow(), err)
		c.Redirect(http.StatusFound, "/404")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(filePath)
	return
}
