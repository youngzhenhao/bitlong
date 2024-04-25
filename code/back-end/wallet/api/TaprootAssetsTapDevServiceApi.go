package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc/tapdevrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
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
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

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
