package main

import (
	"github.com/wallet/api"
)

const PATH = "D:\\share\\bitlong\\code\\back-end\\wallet\\config"
const PATH2 = "/home/en/test"

func main() {
	//api.StartLitd()
	println(api.GetInfoOfTap())
	println(api.GetAssetInfo("f0397b0183d21978ffd5c2c00c5d4d417758c360d096ffa2c3a9fe065df306a1"))
}

func init() {
	api.SetPath(PATH)
}
