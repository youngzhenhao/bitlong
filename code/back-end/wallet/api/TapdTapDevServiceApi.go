package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc/tapdevrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"path/filepath"
)

// ImportProof
//
//	@Description:ImportProof attempts to import a proof file into the daemon.
//	If successful, a new asset will be inserted on disk, spendable using the specified target script key, and internal key.
//	@return bool
func ImportProof(proofFile, genesisPoint string) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
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
