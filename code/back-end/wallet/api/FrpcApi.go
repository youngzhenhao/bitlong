package api

import (
	"github.com/fatedier/frp/cmd/frpc/sub"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/wallet/base"
	"strconv"
)

// FrpcConfig
// Need to test to find the "current path" of the Android
// TODO: modify source code to set configuration file path
// 2024-04-03 15:22:47.852 [I] [sub/root.go:142] start frpc service for config file [./frpc.ini]
// 2024-04-03 15:22:47.874 [I] [client/service.go:294] try to connect to server...
// 2024-04-03 15:22:48.157 [I] [client/service.go:286] [a967e563d3cd75bb] login to server success, get run id [a967e563d3cd75bb]
// 2024-04-03 15:22:48.158 [I] [proxy/proxy_manager.go:173] [a967e563d3cd75bb] proxy added: [14da130b-5757-45bb-bd72-80df85958f54]
// 2024-04-03 15:22:48.276 [I] [client/control.go:170] [a967e563d3cd75bb] [14da130b-5757-45bb-bd72-80df85958f54] start proxy success
func FrpcConfig(id, remotePortStr string) {
	remotePort, _ := strconv.Atoi(remotePortStr)
	_ = WriteConfig(".\\frpc.ini", base.QueryConfigByKey("serverAddr"), 7000, id, "tcp", "127.0.0.1", 9090, remotePort)
}

func FrpcRun() {
	system.EnableCompatibilityMode()
	sub.Execute()
}
