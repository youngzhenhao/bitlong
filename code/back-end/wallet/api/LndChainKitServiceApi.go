package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightningnetwork/lnd/lnrpc/chainrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
)

func GetBlockWrap(blockHash string) string {
	response, err := GetBlock(blockHash)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	msgBlock := &wire.MsgBlock{}
	blockReader := bytes.NewReader(response.RawBlock)
	err = msgBlock.Deserialize(blockReader)
	return MakeJsonResult(true, "", msgBlock)
}

func GetBlock(blockHashStr string) (response *chainrpc.GetBlockResponse, err error) {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
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
		fmt.Printf(GetTimeNow() + "Failed to append cert")
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)

	blockHash, err := chainhash.NewHashFromStr(blockHashStr)
	client := chainrpc.NewChainKitClient(conn)
	request := &chainrpc.GetBlockRequest{
		BlockHash: blockHash.CloneBytes(),
	}
	response, err = client.GetBlock(context.Background(), request)
	return response, err
}
