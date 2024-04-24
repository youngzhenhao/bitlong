package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
)

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

func MakeJsonResult1(success bool, error string, data string) string {
	data = strings.Replace(data, "\n", "", -1)
	data = strings.Replace(data, "\t", "", -1)
	data = strings.Replace(data, " ", "", -1)
	jstr := "{\"Success\":\"" + strconv.FormatBool(success) + "\",\"Error\":\"" + error + "\",\"Data\":" + data + "}"
	var restr bytes.Buffer
	_ = json.Indent(&restr, []byte(jstr), "", "\t")
	return restr.String()
}
func Base64Decode(s string) string {
	byte1, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "DECODE_ERROR"
	}
	return string(byte1)
}
