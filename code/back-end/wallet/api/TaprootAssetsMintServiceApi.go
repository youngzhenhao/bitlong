package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/mintrpc"

	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"path/filepath"
)

func InitMint() bool {
	const (
		grpcHost = "127.0.0.1:10029"
	)
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "data/testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := ioutil.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)

	cert, err := ioutil.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := mintrpc.NewMintClient(conn)

	assetMeta := &taprpc.AssetMeta{
		Type: 1,
		Data: []byte("bullFire97014 is great"),
	}
	mintAsset := &mintrpc.MintAsset{
		AssetVersion:    1,
		AssetType:       0,
		Name:            "bullFire97014",
		AssetMeta:       assetMeta,
		Amount:          100,
		NewGroupedAsset: true,
	}
	request := &mintrpc.MintAssetRequest{Asset: mintAsset, ShortResponse: true}
	response, err := client.MintAsset(context.Background(), request)
	log.Print(response)
	return true
}

func FinalizeMint() bool {
	const (
		grpcHost = "127.0.0.1:10029"
	)
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "data/testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := ioutil.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)

	cert, err := ioutil.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := mintrpc.NewMintClient(conn)
	request := &mintrpc.FinalizeBatchRequest{}
	response, err := client.FinalizeBatch(context.Background(), request)
	log.Print(response)
	return true
}
func GetTapRootAddr() bool {
	const (
		grpcHost = "127.0.0.1:10029"
	)
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "data/testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := ioutil.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)

	cert, err := ioutil.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	btye, err := hex.DecodeString("58fac0a06471220d3b086d8d35189ad380586ab9260b500f3c060ef549c955e2")
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.NewAddrRequest{
		AssetId: btye,
		Amt:     10,
	}
	response, err := client.NewAddr(context.Background(), request)
	log.Print(response)
	return true
}

func SendAssets() bool {
	const (
		grpcHost = "127.0.0.1:10029"
	)
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "data/testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := ioutil.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)

	cert, err := ioutil.ReadFile(tlsCertPath)
	if err != nil {
		log.Fatalf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatalf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.SendAssetRequest{
		TapAddrs: []string{"taptb1qqqsqqspqqzzqsczw2wfw56q2d5stk87yzw9vpnzpchaqrqp3y0umaexckc75lhwq5ssxcfat33js52y49pgqzlald48wltsmkse4k5cadnxgug0d7zmecmzqcssy3jra2y3kwd93v399hmpwttke8v6sg09evsvmjdfqmvj06k7335mpqss8jrgusekgysyszexkzlg2lhmmractu9eu8e7qhwx99j7xeennmzhpgqsxrpkw4hxjan9wfek2unsvvaz7tm5v4ehgmn9wsh82mnfwejhyum99ekxjemgw3hxjmn89enxjmnpde3k2w33xqcrywgj07486"},
		FeeRate:  0,
	}
	response, err := client.SendAsset(context.Background(), request)
	log.Print(response)
	return true
}
