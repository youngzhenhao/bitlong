package main

import (
	"flag"
	"github.com/wallet/api"
)

func main() {
	var lnu string
	var amount string
	flag.StringVar(&lnu, "lnu", "", "recommend uuid")
	flag.StringVar(&amount, "amount", "", "username")
	api.LnurlPayToLnu(lnu, amount)
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
}
