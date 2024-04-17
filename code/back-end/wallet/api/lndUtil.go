package api

import (
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/protobuf/proto"
)

func GetRespJSON(resp proto.Message) string {
	jsonBytes, err := lnrpc.ProtoJSONMarshalOpts.Marshal(resp)
	if err != nil {
		fmt.Printf("%s unable to decode response: %v\n", GetTimeNow(), err)
		return ""
	}
	return string(jsonBytes)
}
