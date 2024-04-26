package api

import (
	"fmt"
	"io"
	"os"
)

func CreateFile(dir, filename, content string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Printf("MkdirAll error: %v", err)
			return false
		}
	}
	filePath := dir + "/" + filename
	err := os.WriteFile(filePath, []byte(content), 0777)
	if err != nil {
		fmt.Printf("WriteFile error: %v", err)
		return false
	}
	fmt.Printf("Successed!\n")
	return true
}

// testCreateFile ONLY_FOR_TEST IN WINDOWS
func testCreateFile() bool {
	//folderPath := base.Configure("tapd")
	return CreateFile("C:\\mySpace\\bitlong\\code\\back-end\\wallet\\base", "TESTWriteFile.txt", "TEST MESSAGE.")
}

func ReadFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return ""
	}
	return string(content)
}

// testReadFile ONLY_FOR_TEST IN WINDOWS
func testReadFile() bool {
	content := ReadFile("C:\\mySpace\\bitlong\\code\\back-end\\wallet\\config.txt")
	if content == "" {
		fmt.Printf("content is EMPTY or read fail!\n")
		return false
	}
	fmt.Printf("%v\n", content)
	return true
}

func CopyFile(srcPath, destPath string) bool {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Printf("Error opening source file: %v", err)
		return false
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			fmt.Printf("Error srcFile.Close: %v", err)
		}
	}(srcFile)
	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Printf("Error creating destination file: %v", err)
		return false
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			fmt.Printf("Error destFile.Close: %v", err)
		}
	}(destFile)
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Printf("Error copying file: %v", err)
		return false
	}
	fmt.Printf("File copied from %s to %s", srcPath, destPath)
	return true
}

// testCopyFile ONLY_FOR_TEST IN WINDOWS
func testCopyFile() bool {
	return CopyFile("C:\\mySpace\\bitlong\\code\\back-end\\wallet\\config.txt", "C:\\mySpace\\bitlong\\code\\back-end\\wallet\\base\\config.txt")
}

func DeleteFile(filePath string) bool {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error deleting file: %v", err)
		return false
	}
	fmt.Printf("File deleted: %s", filePath)
	return true
}

// testDeleteFile ONLY_FOR_TEST IN WINDOWS
func testDeleteFile() bool {
	return DeleteFile("C:\\mySpace\\bitlong\\code\\back-end\\wallet\\base\\config.txt")
}
