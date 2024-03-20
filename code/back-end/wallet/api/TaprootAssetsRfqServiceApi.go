package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc/rfqrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"os"
	"path/filepath"
)

func AddAssetBuyOrder() bool {
	const (
		grpcHost = "202.79.173.41:8443"
	)
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "."+"macaroonfile")
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

	// 连接到grpc服务器
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: grpc.Dial: %v", err)
	}
	// 匿名函数延迟关闭grpc连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close Error: %v", err)
		}
	}(conn)
	// 创建客户端
	client := rfqrpc.NewRfqClient(conn)
	// 构建请求
	request := &rfqrpc.AddAssetBuyOrderRequest{}
	// 得到响应
	response, err := client.AddAssetBuyOrder(context.Background(), request)
	if err != nil {
		log.Printf("rfqrpc  Error: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

func AddAssetSellOffer() bool {
	const (
		grpcHost = "202.79.173.41:8443"
	)
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "."+"macaroonfile")
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

	// 连接到grpc服务器
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: grpc.Dial: %v", err)
	}
	// 匿名函数延迟关闭grpc连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close Error: %v", err)
		}
	}(conn)
	// 创建客户端
	client := rfqrpc.NewRfqClient(conn)
	// 构建请求
	request := &rfqrpc.AddAssetSellOfferRequest{}
	// 得到响应
	response, err := client.AddAssetSellOffer(context.Background(), request)
	if err != nil {
		log.Printf("rfqrpc  Error: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

func QueryRfqAcceptedQuotes() string {
	const (
		grpcHost = "202.79.173.41:8443"
	)
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "."+"macaroonfile")
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

	// 连接到grpc服务器
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: grpc.Dial: %v", err)
	}
	// 匿名函数延迟关闭grpc连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close Error: %v", err)
		}
	}(conn)
	// 创建客户端
	client := rfqrpc.NewRfqClient(conn)
	// 构建请求
	request := &rfqrpc.QueryRfqAcceptedQuotesRequest{}
	// 得到响应
	response, err := client.QueryRfqAcceptedQuotes(context.Background(), request)
	if err != nil {
		log.Printf("rfqrpc QueryRfqAcceptedQuotes Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

func SubscribeRfqEventNtfns() bool {
	const (
		grpcHost = "202.79.173.41:8443"
	)
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "."+"macaroonfile")
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

	// 连接到grpc服务器
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		log.Printf("did not connect: grpc.Dial: %v", err)
	}
	// 匿名函数延迟关闭grpc连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("conn Close Error: %v", err)
		}
	}(conn)
	// 创建客户端
	client := rfqrpc.NewRfqClient(conn)
	// 构建请求
	request := &rfqrpc.SubscribeRfqEventNtfnsRequest{}
	// 得到响应
	stream, err := client.SubscribeRfqEventNtfns(context.Background(), request)
	if err != nil {
		log.Printf("rfqrpc  Error: %v", err)
		return false
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// 流已经关闭，退出循环
				log.Printf("err == io.EOF, err: %v\n", err)
				return false
			}
			log.Printf("stream Recv err: %v\n", err)
			return false
		}
		log.Printf("%v\n", response)
		return true
	}

}
