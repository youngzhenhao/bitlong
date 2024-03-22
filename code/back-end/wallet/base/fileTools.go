package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ReadConfigFile(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("f.Close err: %v\n", err)
		}
	}(f)
	if err != nil {
		fmt.Printf("open file err: %v\n", err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("read file err: %v\n", err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config

}

func Configure(appName string) string {
	fileConfig := ReadConfigFile("C:\\mySpace\\bitlong\\code\\back-end\\wallet\\config.txt")
	//fileConfig := ReadConfigFile("/data/user/0/io.bitlong/files/NewFolderBit/config.txt")
	//fileConfig := ReadConfigFile("D:\\share\\bitlong\\code\\back-end\\wallet\\config.txt")
	//fileConfig := ReadConfigFile("/home/en/test/config.txt")
	dirPath := fileConfig["dirpath"]
	folderPath := filepath.Join(dirPath, "."+appName)
	return folderPath
}

func QueryConfigByKey(key string) (value string) {
	fileConfig := ReadConfigFile("C:\\mySpace\\bitlong\\code\\back-end\\wallet\\config.txt")
	//fileConfig := ReadConfigFile("D:\\share\\bitlong\\code\\back-end\\wallet\\config.txt")
	//fileConfig := ReadConfigFile("/home/en/test/config.txt")
	value = fileConfig[key]
	return
}
