package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
)

func GetRespJSON(resp proto.Message) string {
	jsonBytes, err := lnrpc.ProtoJSONMarshalOpts.Marshal(resp)
	if err != nil {
		fmt.Printf("%s unable to decode response: %v\n", GetTimeNow(), err)
		return ""
	}
	return string(jsonBytes)
}

type JsonResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

func MakeJsonResult(success bool, error string, data any) string {
	jsr := JsonResult{
		Success: success,
		Error:   error,
		Data:    data,
	}
	jstr, err := json.Marshal(jsr)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return string(jstr)
}

func MultipleRespJsonString(resp proto.Message) string {
	jsonBytes, err := lnrpc.ProtoJSONMarshalOpts.Marshal(resp)
	if err != nil {
		fmt.Println("unable to decode response: ", err)
		return "false"
	}
	return string(jsonBytes)
}

type JsonResult_ONLY_FOR_TEST struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    string `json:"data,omitempty"`
}

func MakeJsonResult_ONLY_FOR_TEST(success bool, error string, data string) string {
	data = strings.Replace(data, "\n", "", -1)
	data = strings.Replace(data, "\t", "", -1)
	data = strings.Replace(data, " ", "", -1)
	jstr := "{\"Success\":\"" + strconv.FormatBool(success) + "\",\"Error\":\"" + error + "\",\"Data\":" + data + "}"
	var restr bytes.Buffer
	_ = json.Indent(&restr, []byte(jstr), "", "\t")
	return restr.String()
}
