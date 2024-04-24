package api

import (
	"github.com/fatedier/frp/cmd/frpc/sub"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/wallet/base"
	"path/filepath"
	"strconv"
)

// FrpcConfig
// Need to test to find the "current path" of the Android
// 2024-04-03 15:22:47.852 [I] [sub/root.go:142] start frpc service for config file [./frpc.ini]
// @dev: modified sub/root.go
func FrpcConfig(id, remotePortStr string) {
	remotePort, _ := strconv.Atoi(remotePortStr)
	_ = WriteConfig(filepath.Join(base.QueryConfigByKey("dirpath"), "frpc.ini"), base.QueryConfigByKey("serverAddr"), 7000, id, "tcp", "127.0.0.1", 9090, remotePort)
}

func FrpcRun() {
	system.EnableCompatibilityMode()
	sub.Execute()
}
