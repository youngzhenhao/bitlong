package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"

	"path/filepath"
)

func SendPaymentV2(invoice string) *lnrpc.Payment {

	const (
		grpcHost = "202.79.173.41:10009"
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

	client := routerrpc.NewRouterClient(conn)
	request := &routerrpc.SendPaymentRequest{
		PaymentRequest: invoice,
		TimeoutSeconds: 60,
	}
	response, err := client.SendPaymentV2(context.Background(), request)

	if err != nil {
		log.Fatalf("client.SendPaymentV2 :%v", err)
	}
	temp, _ := response.Recv()
	log.Print(response.Recv())
	return temp
}
func TrackPaymentV2(payhash []byte) bool {
	const (
		grpcHost = "202.79.173.41:10009"
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

	client := routerrpc.NewRouterClient(conn)
	request := &routerrpc.TrackPaymentRequest{
		PaymentHash: payhash,
	}
	response, err := client.TrackPaymentV2(context.Background(), request)

	if err != nil {
		log.Fatalf("client.SendPaymentV2 :%v", err)
	}
	temp, _ := response.Recv()
	log.Print(temp.Status)

	return true
}
func SendToRouteV2(payhash []byte, route *lnrpc.Route) {
	const (
		grpcHost = "202.79.173.41:10009"
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

	client := routerrpc.NewRouterClient(conn)
	request := &routerrpc.SendToRouteRequest{
		PaymentHash: cert,
		Route:       route,
	}
	response, err := client.SendToRouteV2(context.Background(), request)

	if err != nil {
		log.Fatalf("client.SendPaymentV2 :%v", err)
	}

	log.Print(response)
}
func EstimateRouteFee(dest []byte, amtsat int64) string {
	const (
		grpcHost = "202.79.173.41:10009"
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

	client := routerrpc.NewRouterClient(conn)
	request := &routerrpc.RouteFeeRequest{
		Dest:   dest,
		AmtSat: amtsat,
	}
	response, err := client.EstimateRouteFee(context.Background(), request)

	if err != nil {
		log.Fatalf("client.SendPaymentV2 :%v", err)
	}

	log.Print(response.RoutingFeeMsat)
	return response.String()
}
