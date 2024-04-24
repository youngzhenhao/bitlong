package frpc

import (
	"fmt"
	_ "github.com/fatedier/frp/assets/frpc"
	"github.com/google/uuid"
	"github.com/wallet/api/lnurl"
	"github.com/wallet/base"
	"os"
	"strconv"
)

func WriteConfigFrpcRunTest() {
	id := uuid.New().String()
	//	@dev: Get available port twice to compare to prevent being taken
	port := strconv.Itoa(RequestServerGetPortAvailable(base.QueryConfigByKey("LnurlServerHost")))
	FrpcConfig(id, port)
	FrpcRun()
}

func InitPhoneDBTest() {
	err := lnurl.InitPhoneDB()
	if err != nil {
		fmt.Println("init phone db error,", err)
	}
}

func WriteConfig(filePath string, serverAddr string, serverPort int, proxyName string, proxyType string, localIP string, localPort int, remotePort int) bool {
	content := fmt.Sprintf("serverAddr = \"%s\"\nserverPort = %d\n\n[[proxies]]\nname = \"%s\"\ntype = \"%s\"\nlocalIP = \"%s\"\nlocalPort = %d\nremotePort = %d",
		serverAddr, serverPort, proxyName, proxyType, localIP, localPort, remotePort)
	contentByte := []byte(content)
	err := os.WriteFile(filePath, contentByte, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
