package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path"
	"trade/utils"
)

func FileUpload(c *gin.Context) {
	//fileName := c.Query("filename")
	saveName := uuid.New().String()
	saveDir := "files"
	file, header, err := c.Request.FormFile("file")
	result := true
	if err != nil {
		result = false
		fmt.Printf("%s %v\n", utils.GetTimeNow(), err)
		c.JSON(http.StatusOK, gin.H{
			"time":   utils.GetTimeNow(),
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
			fmt.Printf("%s mkdir %s error %v\n", utils.GetTimeNow(), saveDir, err)
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
		fmt.Printf("%s create file %s error %v\n", utils.GetTimeNow(), savePath, err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Printf("%s close file %s error %v\n", utils.GetTimeNow(), savePath, err)
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
	filePath := "files/" + c.Query("file")
	fileTmp, err := os.Open(filePath)
	defer func(fileTmp *os.File) {
		err := fileTmp.Close()
		if err != nil {
			fmt.Printf("%s close file error %v\n", utils.GetTimeNow(), err)
		}
	}(fileTmp)
	fileName := path.Base(filePath)
	if filePath == "" || fileName == "" || err != nil {
		fmt.Printf("%s file not found %v\n", utils.GetTimeNow(), err)
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
