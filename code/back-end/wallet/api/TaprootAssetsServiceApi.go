package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
)

func AddrReceives() {

}

func BurnAsset() {

}

func DebugLevel() {

}

func DecodeAddr() {

}

func DecodeProof() {

}

func ExportProof() {

}

func FetchAssetMeta() {

}

// TapGetInfo
//
//	@Description: 返回节点的信息
//	@return *taprpc.GetInfoResponse
//
// func TapGetInfo() *taprpc.GetInfoResponse {
func TapGetInfo() string {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.GetInfoRequest{}
	// 得到响应
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		log.Printf("taprpc GetInfo Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// ListAssets
//
//	@Description: 列出目标守护程序拥有的资产集，返回的嵌套结构体中的[]byte需要使用get函数单独处理，如hex.EncodeToString
//	请注意，IncludeSpent和IncludeLeased不能同时指定为true
//	@return *taprpc.ListAssetResponse
//
// func ListAssets(withWitness, includeSpent, includeLeased bool) *taprpc.ListAssetResponse {
func ListAssets(withWitness, includeSpent, includeLeased bool) string {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListAssetRequest{
		WithWitness:   withWitness,
		IncludeSpent:  includeSpent,
		IncludeLeased: includeLeased,
	}
	// 得到响应
	response, err := client.ListAssets(context.Background(), request)
	if err != nil {
		log.Printf("taprpc ListAssets Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// ListBalances
//
//	@Description: 列出资产余额
//	参数为true按照AssetsId排序，false则按照GroupKey排序，资产和组键过滤器未设置
//	@return *taprpc.ListBalancesResponse
//
// func ListBalances(isListAssetIdNotGroupKey bool) *taprpc.ListBalancesResponse {
func ListBalances(isListAssetIdNotGroupKey bool) string {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListBalancesRequest{
		// 不设置过滤器
		AssetFilter:    nil,
		GroupKeyFilter: nil,
	}
	// 根据给定参数修改请求结构体的排序方式
	if isListAssetIdNotGroupKey {
		request.GroupBy = &taprpc.ListBalancesRequest_AssetId{AssetId: true}
	} else {
		request.GroupBy = &taprpc.ListBalancesRequest_GroupKey{GroupKey: true}
	}
	// 得到响应
	response, err := client.ListBalances(context.Background(), request)
	if err != nil {
		log.Printf("taprpc ListBalances Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// ListGroups
//
//	@Description: 列出目标守护程序已知的资产组，以及每个组中持有的资产
//	@return *taprpc.ListGroupsResponse
//
// func ListGroups() *taprpc.ListGroupsResponse {
func ListGroups() string {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListGroupsRequest{}
	// 得到响应
	response, err := client.ListGroups(context.Background(), request)
	if err != nil {
		log.Printf("taprpc ListGroups Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// ListTransfers
//
//	@Description: 列出目标守护程序跟踪的出站资产转移
//	@return *taprpc.ListTransfersResponse
//
// func ListTransfers() *taprpc.ListTransfersResponse {
func ListTransfers() string {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListTransfersRequest{}
	// 得到响应
	response, err := client.ListTransfers(context.Background(), request)
	if err != nil {
		log.Printf("taprpc ListTransfers Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// ListUtxos
//
//	@Description: 列出目标守护进程管理的 UTXO 及其持有的资产
//	@param includeLeased
//	@return *taprpc.ListUtxosResponse
//
// func ListUtxos(includeLeased bool) *taprpc.ListUtxosResponse {
func ListUtxos(includeLeased bool) string {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.ListUtxosRequest{
		IncludeLeased: includeLeased,
	}
	// 得到响应
	response, err := client.ListUtxos(context.Background(), request)
	if err != nil {
		log.Printf("taprpc ListUtxos Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// func NewAddr(assetId string, amt uint64) *taprpc.Addr {
func NewAddr(assetId string, amt int) bool {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	_assetIdByteSlice, _ := hex.DecodeString(assetId)
	// 构建请求
	request := &taprpc.NewAddrRequest{
		AssetId: _assetIdByteSlice,
		Amt:     uint64(amt),
	}
	// 得到响应
	response, err := client.NewAddr(context.Background(), request)
	if err != nil {
		log.Printf("taprpc NewAddr Error: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

func QueryAddrs() {

}

// func SendAsset(tapAddrs []string, feeRate uint32) *taprpc.SendAssetResponse {
func SendAsset(tapAddrs []string, feeRate int) bool {
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
	client := taprpc.NewTaprootAssetsClient(conn)
	// 构建请求
	request := &taprpc.SendAssetRequest{
		TapAddrs: tapAddrs,
		FeeRate:  uint32(feeRate),
	}
	// 得到响应
	response, err := client.SendAsset(context.Background(), request)
	if err != nil {
		log.Printf("taprpc SendAsset Error: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

func StopDaemon() {

}

func SubscribeReceiveAssetEventNtfns() {

}

func SubscribeSendAssetEventNtfns() {

}

func VerifyProof() {

}
