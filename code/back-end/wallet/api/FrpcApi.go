package api

import (
	"github.com/fatedier/frp/cmd/frpc/sub"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/wallet/base"
	"strconv"
)

func FrpcRun(id, remotePortStr string) {
	remotePort, _ := strconv.Atoi(remotePortStr)
	_ = WriteConfig(base.QueryConfigByKey("serverAddr"), 7000, id, "tcp", "127.0.0.1", 9090, remotePort)
	system.EnableCompatibilityMode()
	sub.Execute()
}
