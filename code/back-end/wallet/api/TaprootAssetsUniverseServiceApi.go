package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
)

func AddFederationServer() {}

func AssetLeafKeys() {}

func AssetLeaves() {}

func AssetRoots() {}

func DeleteAssetRoot() {}

func DeleteFederationServer() {}

// Info
//
//	@Description: 返回一组关于宇宙当前状态的信息
//	@return *universerpc.InfoResponse
//
// func UniverseInfo() *universerpc.InfoResponse {
func UniverseInfo() string {
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
	client := universerpc.NewUniverseClient(conn)
	// 构建请求
	request := &universerpc.InfoRequest{}
	// 得到响应
	response, err := client.Info(context.Background(), request)
	if err != nil {
		log.Printf("universerpc Info Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

func InsertProof() {}

// ListFederationServers
//
//	@Description: 列出了组成本地 Universe 服务器联盟的服务器集。这些服务器用于推送新的证明，并定期从远程服务器同步调用新的证明
//	@return *universerpc.ListFederationServersResponse
//
// func ListFederationServers() *universerpc.ListFederationServersResponse {
func ListFederationServers() string {
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
	client := universerpc.NewUniverseClient(conn)
	// 构建请求
	request := &universerpc.ListFederationServersRequest{}
	// 得到响应
	response, err := client.ListFederationServers(context.Background(), request)
	if err != nil {
		log.Printf("universerpc ListFederationServers Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

func MultiverseRoot() {}

func QueryAssetRoots() {}

func QueryAssetStats() {}

func QueryEvents() {}

func QueryFederationSyncConfig() {}

func QueryProof() {}

func SetFederationSyncConfig() {}

func SyncUniverse() {}

func UniverseStats() {}
