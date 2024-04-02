package api

import (
	"github.com/fatedier/frp/cmd/frpc/sub"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/wallet/base"
	"strconv"
)

// FrpcConfig
// TODO: Need to test to find the "current path" of the Android
func FrpcConfig(id, remotePortStr string) {
	remotePort, _ := strconv.Atoi(remotePortStr)
	_ = WriteConfig(".\\frpc.ini", base.QueryConfigByKey("serverAddr"), 7000, id, "tcp", "127.0.0.1", 9090, remotePort)
}

func FrpcRun() {
	system.EnableCompatibilityMode()
	sub.Execute()
}
