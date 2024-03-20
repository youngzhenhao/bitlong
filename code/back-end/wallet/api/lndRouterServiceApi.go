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
	"io"
	"log"
	"os"

	"path/filepath"
)

func SendPaymentV2(invoice string, feelimit int64) string {
	const (
		grpcHost = "202.79.173.41:10009"
	)
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
		log.Printf("Failed to read cert file: %s", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Printf("Failed to append cert")
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close err: %v", err)
		}
	}(conn)
	client := routerrpc.NewRouterClient(conn)
	request := &routerrpc.SendPaymentRequest{
		PaymentRequest: invoice,
		FeeLimitSat:    feelimit,
		TimeoutSeconds: 60,
	}
	stream, err := client.SendPaymentV2(context.Background(), request)
	if err != nil {
		log.Printf("client.SendPaymentV2 :%v", err)
		return "false"
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// 流已经关闭，退出循环
				log.Printf("err == io.EOF, err: %v\n", err)
				return "false"
			}
			log.Printf("stream Recv err: %v\n", err)
			return "false"
		}
		log.Printf("%v\n", response)
		return response.PaymentHash
	}
}

func TrackPaymentV2(payhash string) string {
	const (
		grpcHost = "202.79.173.41:10009"
	)
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
		log.Printf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Printf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close err: %v", err)
		}
	}(conn)
	client := routerrpc.NewRouterClient(conn)
	_payhashByteSlice, _ := hex.DecodeString(payhash)
	request := &routerrpc.TrackPaymentRequest{
		PaymentHash: _payhashByteSlice,
	}
	stream, err := client.TrackPaymentV2(context.Background(), request)

	if err != nil {
		log.Printf("client.SendPaymentV2 :%v", err)
		return "false"
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// 流已经关闭，退出循环
				log.Printf("err == io.EOF, err: %v\n", err)
				return "false"
			}
			log.Printf("stream Recv err: %v\n", err)
			return "false"
		}
		log.Printf("%v\n", response)
		return response.Status.String()
	}
}

func SendToRouteV2(payhash []byte, route *lnrpc.Route) {
	const (
		grpcHost = "202.79.173.41:10009"
	)
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
		log.Printf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Printf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close err: %v", err)
		}
	}(conn)
	client := routerrpc.NewRouterClient(conn)
	request := &routerrpc.SendToRouteRequest{
		PaymentHash: cert,
		Route:       route,
	}
	response, err := client.SendToRouteV2(context.Background(), request)

	if err != nil {
		log.Printf("client.SendPaymentV2 :%v", err)
	}

	log.Print(response)
}

func EstimateRouteFee(dest string, amtsat int64) string {
	const (
		grpcHost = "202.79.173.41:10009"
	)
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
		log.Printf("Failed to read cert file: %s", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		log.Printf("Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close err: %v", err)
		}
	}(conn)
	client := routerrpc.NewRouterClient(conn)

	bDest, _ := hex.DecodeString(dest)
	request := &routerrpc.RouteFeeRequest{
		Dest:   bDest,
		AmtSat: amtsat,
	}
	response, err := client.EstimateRouteFee(context.Background(), request)
	if err != nil {
		log.Printf("client.Sen下dPaymentV2 :%v", err)
	}
	log.Print(response.RoutingFeeMsat)
	return response.String()
}
