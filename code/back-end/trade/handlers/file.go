package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path"
	"trade/utils"
)

func FileUpload(c *gin.Context) {
	saveName := uuid.New().String()
	saveDir := "files"
	file, header, err := c.Request.FormFile("file")
	result := true
	if err != nil {
		result = false
		utils.LogError("", err)
		c.JSON(http.StatusOK, gin.H{
			"time":   utils.GetTimeNow(),
			"result": result,
			"info":   err,
		})
		return
	}
	uploadFileNameWithSuffix := path.Base(header.Filename)
	saveName = uploadFileNameWithSuffix
	savePath := saveDir + "/" + saveName
	var localFileInfo os.FileInfo
	localFileInfo, err = os.Stat(saveDir)
	if err != nil || !localFileInfo.IsDir() {
		err := os.MkdirAll(saveDir, 0755)
		if err != nil {
			utils.LogInfos("mkdir", saveDir, "error.", err.Error())
			result = false
			c.JSON(http.StatusOK, gin.H{
				"time":   utils.GetTimeNow(),
				"result": result,
				"info":   err,
			})
			return
		}
	}
	out, err := os.Create(savePath)
	if err != nil {
		utils.LogInfos("create file", saveDir, "error.", err.Error())
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			utils.LogInfos("close file", saveDir, "error.", err.Error())
		}
	}(out)
	_, err = io.Copy(out, file)
	if err != nil {
		result = false
		c.JSON(http.StatusOK, gin.H{
			"time":   utils.GetTimeNow(),
			"result": result,
			"info":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"time":   utils.GetTimeNow(),
		"result": result,
		"info":   "uploaded successfully",
	})
	return
}

func FileDownload(c *gin.Context) {
	filePath := "/root/files/" + c.Query("file")
	fileTmp, err := os.Open(filePath)
	defer func(fileTmp *os.File) {
		err := fileTmp.Close()
		if err != nil {
			utils.LogError("close file error", err)
		}
	}(fileTmp)
	fileName := path.Base(filePath)
	if filePath == "" || fileName == "" || err != nil {
		utils.LogError("file not found", err)
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
