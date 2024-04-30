package api

import (
	"fmt"
	"github.com/wallet/base"
)

func SetPath(path string) {
	err := base.SetFilePath(path)
	if err != nil {
		fmt.Println("set file path error. ", err)
		return
	}
}

func GetPath() string {
	return base.GetFilePath()
}

func FileTestConfig() bool {
	return base.FileConfig(GetPath())
}

func ReadConfigFile() {
	base.ReadConfig(GetPath())
}

func ReadConfigFile1() {
	base.ReadConfig1(GetPath())
}

func ReadConfigFile2() {
	base.ReadConfig2(GetPath())
}

func CreateDir() {
	base.CreateDir(GetPath())
}

func CreateDir2() {
	base.CreateDir2(GetPath())
}

func Visit() {
	base.VisitAll()
}
