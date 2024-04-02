package main

import (
	"flag"
	"github.com/wallet/api"
)

func main() {
	var id string
	var portRemote string
	flag.StringVar(&id, "id", "", "recommend uuid")
	flag.StringVar(&portRemote, "portRemote", "", "port remote client wanna to request")
	flag.Parse()

	api.LnurlFrpcRun(id, portRemote)
}
