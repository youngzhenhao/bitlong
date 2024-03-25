package api

import (
	"fmt"
	"github.com/wallet/base"
	"os"
)

func CreateFileWithContent(dir, filename, content string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Printf("MkdirAll error: %v", err)
			return false
		}
	}
	filePath := dir + "/" + filename + ".txt"
	err := os.WriteFile(filePath, []byte(content), 0777)
	if err != nil {
		fmt.Printf("WriteFile error: %v", err)
		return false
	}
	fmt.Printf("Successed!\n")
	return true
}
func TESTWriteFile() bool {
	folderPath := base.Configure("tapd")
	return CreateFileWithContent(folderPath, "TESTWriteFile", "TEST MESSAGE.")
}
