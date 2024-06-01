package api

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/protobuf/proto"
	"log"
	"math"
	"os"
	"time"
)

type ErrCode int

const (
	OLD     ErrCode = -1
	SUCCESS ErrCode = 200
)

type JsonResult struct {
	Success bool    `json:"success"`
	Error   string  `json:"error"`
	Code    ErrCode `json:"code"`
	Data    any     `json:"data"`
}

// Deprecated: Use MakeJsonErrorResult instead
func MakeJsonResult(success bool, error string, data any) string {
	jsr := JsonResult{
		Success: success,
		Error:   error,
		Code:    -1,
		Data:    data,
	}
	jstr, err := json.Marshal(jsr)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return string(jstr)
}

func MakeJsonErrorResult(code ErrCode, error string, data any) string {
	jsr := JsonResult{
		Success: true,
		Error:   error,
		Code:    code,
		Data:    data,
	}
	jstr, err := json.Marshal(jsr)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return string(jstr)
}

func LnMarshalRespString(resp proto.Message) string {
	jsonBytes, err := lnrpc.ProtoJSONMarshalOpts.Marshal(resp)
	if err != nil {
		fmt.Printf("%s unable to decode response: %v\n", GetTimeNow(), err)
		return ""
	}
	return string(jsonBytes)
}

func TapMarshalRespString(resp proto.Message) string {
	jsonBytes, err := taprpc.ProtoJSONMarshalOpts.Marshal(resp)
	if err != nil {
		fmt.Printf("%s unable to decode response: %v\n", GetTimeNow(), err)
		return ""
	}
	return string(jsonBytes)
}

func B64DecodeToHex(s string) string {
	byte1, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "DECODE_ERROR"
	}
	return hex.EncodeToString(byte1)
}

func GetTimeNow() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func GetTimeSuffixString() string {
	return time.Now().Format("20060102150405")
}

func RoundToDecimalPlace(number float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(number*shift) / shift
}

func GetEnv(key string, filename ...string) string {
	err := godotenv.Load(filename...)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	value := os.Getenv(key)
	return value
}

func ToBTC(sat int) float64 {
	return float64(sat / 1e8)
}

func ToSat(btc float64) int {
	return int(btc * 1e8)
}

func LogInfo(info string) {
	fmt.Printf("%s %s\n", GetTimeNow(), info)
}

func LogInfos(infos ...string) {
	var info string
	for i, _info := range infos {
		if i != 0 {
			info += " "
		}
		info += _info
	}
	fmt.Printf("%s %s\n", GetTimeNow(), info)
}

func LogError(description string, err error) {
	fmt.Printf("%s %s :%v\n", GetTimeNow(), description, err)
}

// FeeRateBtcPerKbToSatPerKw
// @Description: BTC/Kb to sat/kw
// 1 sat/vB = 0.25 sat/wu
// https://bitcoin.stackexchange.com/questions/106333/different-fee-rate-units-sat-vb-sat-perkw-sat-perkb
func FeeRateBtcPerKbToSatPerKw(btcPerKb float64) (satPerKw int) {
	// @dev: 1 BTC/kB = 1e8 sat/kB 1e5 sat/B = 0.25e5 sat/w = 0.25e8 sat/kw
	return int(0.25e8 * btcPerKb)
}

// FeeRateBtcPerKbToSatPerB
// @Description: BTC/Kb to sat/b
// @param btcPerKb
// @return satPerB
func FeeRateBtcPerKbToSatPerB(btcPerKb float64) (satPerB int) {
	return int(1e5 * btcPerKb)
}

// FeeRateSatPerKwToBtcPerKb
// @Description: sat/kw to BTC/Kb
// @param feeRateSatPerKw
// @return feeRateBtcPerKb
func FeeRateSatPerKwToBtcPerKb(feeRateSatPerKw int) (feeRateBtcPerKb float64) {
	return RoundToDecimalPlace(float64(feeRateSatPerKw)/0.25e8, 8)
}

// FeeRateSatPerKwToSatPerB
// @Description: sat/kw to sat/b
// @param feeRateSatPerKw
// @return feeRateSatPerB
func FeeRateSatPerKwToSatPerB(feeRateSatPerKw int) (feeRateSatPerB int) {
	return feeRateSatPerKw * 4 / 1000
}

// FeeRateSatPerBToBtcPerKb
// @Description: sat/b to BTC/Kb
// @param feeRateSatPerB
// @return feeRateBtcPerKb
func FeeRateSatPerBToBtcPerKb(feeRateSatPerB int) (feeRateBtcPerKb float64) {
	return RoundToDecimalPlace(float64(feeRateSatPerB)/100000, 8)
}

// FeeRateSatPerBToSatPerKw
// @Description: sat/b to sat/kw
// @param feeRateSatPerB
// @return feeRateSatPerKw
func FeeRateSatPerBToSatPerKw(feeRateSatPerB int) (feeRateSatPerKw int) {
	return feeRateSatPerB * 1000 / 4
}

func ValueJsonString(value any) string {
	resultJSON, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		LogError("MarshalIndent error", err)
		return ""
	}
	return string(resultJSON)
}
