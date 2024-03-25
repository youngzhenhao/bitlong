package api

import (
	"fmt"
	"github.com/wallet/base"
	"os"
)

// CreateFileWithContent 在指定的目录中创建一个新的txt文件，并写入内容
func CreateFileWithContent(dir, filename, content string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return fmt.Errorf("MkdirAll error: %v", err)
		}
	}
	filePath := dir + "/" + filename + ".txt"
	err := os.WriteFile(filePath, []byte(content), 0777)
	if err != nil {
		return fmt.Errorf("WriteFile error: %v", err)
	}
	fmt.Printf("Successed!\n")
	return nil
}
func TESTWriteFile() bool {
	folderPath := base.Configure("tapd")
	err := CreateFileWithContent(folderPath, "TESTWriteFile", "TEST MESSAGE.")
	if err != nil {
		return false
	}
	return true
}
