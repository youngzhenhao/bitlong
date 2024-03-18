package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightningnetwork/lnd/lnrpc/walletrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ListAddress() bool {
	const (
		grpcHost = "127.0.0.1:10009"
	)
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
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
	client := walletrpc.NewWalletKitClient(conn)
	request := &walletrpc.ListAddressesRequest{}
	response, err := client.ListAddresses(context.Background(), request)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println(response)
	return true
}

type macaroonCredential struct {
	macaroon string
}

func newMacaroonCredential(macaroon string) *macaroonCredential {
	return &macaroonCredential{macaroon: macaroon}
}

func (c *macaroonCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"macaroon": c.macaroon}, nil
}

func (c *macaroonCredential) RequireTransportSecurity() bool {
	return true
}
