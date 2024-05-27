package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc/tapdevrpc"
	"github.com/wallet/api/connect"
)

// ImportProof
//
//	@Description:ImportProof attempts to import a proof file into the daemon.
//	If successful, a new asset will be inserted on disk, spendable using the specified target script key, and internal key.
//	@return bool
func ImportProof(proofFile, genesisPoint string) bool {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := tapdevrpc.NewTapDevClient(conn)
	_proofFileByteSlice, _ := hex.DecodeString(proofFile)
	request := &tapdevrpc.ImportProofRequest{
		ProofFile:    _proofFileByteSlice,
		GenesisPoint: genesisPoint,
	}
	response, err := client.ImportProof(context.Background(), request)
	if err != nil {
		fmt.Printf("%s tapdevrpc QueryRfqAcceptedQuotes Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}
