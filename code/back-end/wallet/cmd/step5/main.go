package main

import (
	"flag"
	"github.com/wallet/api"
)

func main() {
	var id string
	var name string
	var localPort string
	var remotePort string
	flag.StringVar(&id, "id", "", "recommend uuid")
	flag.StringVar(&name, "name", "", "username")
	flag.StringVar(&localPort, "localPort", "", "use 9090")
	flag.StringVar(&remotePort, "remotePort", "", "port remote client wanna to request")
	flag.Parse()
	api.LnurlUploadUserInfo(id, name, localPort, remotePort)
}
