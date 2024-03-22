package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wallet/api"
)

func main() {

	//var password string
	//flag.StringVar(&password, "password", "", "password must have at least 8 characters")
	//flag.Parse()
	//
	//fmt.Printf("%s password: %v\n", api.GetTimeNow(), password)

	//api.StarLnd()
	//api.StopDaemon()
	//api.StartTapRoot()

}

func s2json(value any) string {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("%s %v", api.GetTimeNow(), err)
	}
	var str bytes.Buffer
	err = json.Indent(&str, jsonBytes, "", "\t")
	if err != nil {
		fmt.Printf("%s %v", api.GetTimeNow(), err)
	}
	result := str.String()
	return result
}
