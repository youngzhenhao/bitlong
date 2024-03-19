package main

import (
	"bytes"
	"encoding/json"
	"github.com/wallet/lnurl"
	"log"
)

func main() {
	//api.StarLnd()
	//api.StartTapRoot()

	lnurl.LnurlRouterRun()
}

func s2json(value any) string {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		log.Printf("%v", err)
	}
	var str bytes.Buffer
	err = json.Indent(&str, jsonBytes, "", "\t")
	if err != nil {
		log.Printf("%v", err)
	}
	result := str.String()
	return result
}
