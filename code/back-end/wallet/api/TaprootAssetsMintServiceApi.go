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
	"log"
	"os"
	"path/filepath"
)

// InitMint SAME AS MintAsset, HAVE NOT DELETED
func InitMint() bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "data/testnet")
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
			log.Fatalf("conn Close err: %v", err)
		}
	}(conn)
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
	if err != nil {
		log.Printf("mintrpc MintAsset err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// FinalizeMint SAME AS FinalizeBatch, HAVE NOT DELETED
func FinalizeMint() bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "data/testnet")
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
			log.Fatalf("conn Close err: %v", err)
		}
	}(conn)
	client := mintrpc.NewMintClient(conn)
	request := &mintrpc.FinalizeBatchRequest{}
	response, err := client.FinalizeBatch(context.Background(), request)
	if err != nil {
		log.Printf("mintrpc FinalizeBatch err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// GetTapRootAddr SAME AS NewAddr IN TaprootAssetsServiceApi, HAVE NOT DELETED
func GetTapRootAddr(assetId string, amt int) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "data/testnet")
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
			log.Fatalf("conn Close err: %v", err)
		}
	}(conn)
	_assetIdBtyeSlice, err := hex.DecodeString(assetId)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.NewAddrRequest{
		AssetId: _assetIdBtyeSlice,
		Amt:     uint64(amt),
	}
	response, err := client.NewAddr(context.Background(), request)
	if err != nil {
		log.Printf("taprpc NewAddr err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// SendAssets SAME AS SendAsset IN TaprootAssetsServiceApi, HAVE NOT DELETED
func SendAssets() bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("tapd"), "data/testnet")
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
			log.Fatalf("conn Close err: %v", err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.SendAssetRequest{
		TapAddrs: []string{"taptb1qqqsqqspqqzzqsczw2wfw56q2d5stk87yzw9vpnzpchaqrqp3y0umaexckc75lhwq5ssxcfat33js52y49pgqzlald48wltsmkse4k5cadnxgug0d7zmecmzqcssy3jra2y3kwd93v399hmpwttke8v6sg09evsvmjdfqmvj06k7335mpqss8jrgusekgysyszexkzlg2lhmmractu9eu8e7qhwx99j7xeennmzhpgqsxrpkw4hxjan9wfek2unsvvaz7tm5v4ehgmn9wsh82mnfwejhyum99ekxjemgw3hxjmn89enxjmnpde3k2w33xqcrywgj07486"},
		FeeRate:  0,
	}
	response, err := client.SendAsset(context.Background(), request)
	if err != nil {
		log.Printf("taprpc SendAsset err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// CancelBatch
//
//	@Description: 将尝试取消当前待处理批次
//	@return *mintrpc.CancelBatchResponse
//
// func CancelBatch() *mintrpc.CancelBatchResponse {
func CancelBatch() bool {
	grpcHost := base.QueryConfigByKey("taproothost")
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
	client := mintrpc.NewMintClient(conn)
	// 构建请求
	request := &mintrpc.CancelBatchRequest{}
	// 得到响应
	response, err := client.CancelBatch(context.Background(), request)
	if err != nil {
		log.Printf("mintrpc CancelBatch Error: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// FinalizeBatch
//
//	@Description: 将尝试完成当前待处理的批次
//	@return *mintrpc.FinalizeBatchResponse
//
// func FinalizeBatch(shortResponse bool, feeRate uint32) *mintrpc.FinalizeBatchResponse {
func FinalizeBatch(shortResponse bool, feeRate int) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
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
	client := mintrpc.NewMintClient(conn)
	// 构建请求
	request := &mintrpc.FinalizeBatchRequest{
		ShortResponse: shortResponse,
		FeeRate:       uint32(feeRate),
	}
	// 得到响应
	response, err := client.FinalizeBatch(context.Background(), request)
	if err != nil {
		log.Printf("mintrpc FinalizeBatch Error: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// ListBatches
//
//	 @Description: 列出提交到守护程序的批处理集，包括待处理和已取消的批处理
//		过滤器设置为空
//	 @return *mintrpc.ListBatchResponse
//
// func ListBatches() *mintrpc.ListBatchResponse {
func ListBatches() string {
	grpcHost := base.QueryConfigByKey("taproothost")
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
	client := mintrpc.NewMintClient(conn)
	// 构建请求
	request := &mintrpc.ListBatchRequest{}
	// 得到响应
	response, err := client.ListBatches(context.Background(), request)
	if err != nil {
		log.Printf("mintrpc ListBatches Error: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// MintAsset
//
//	@Description:  将尝试铸模请求中指定的资产集（默认为异步，以确保正确的批处理）。
//	返回的待处理批次将显示属于下一批次的其他待处理资产。此调用将阻塞，直至操作成功（资产已在批次中分期）或失败
//	@return *mintrpc.MintAssetResponse
//
// func MintAsset(assetVersionIsV1 bool, assetTypeIsCollectible bool, name string, assetMetaData string, AssetMetaTypeIsJsonNotOpaque bool, amount uint64, newGroupedAsset bool, groupedAsset bool, groupKey string, groupAnchor string, shortResponse bool) *mintrpc.MintAssetResponse {
func MintAsset(assetVersionIsV1 bool, assetTypeIsCollectible bool, name string, assetMetaData string, AssetMetaTypeIsJsonNotOpaque bool, amount int, newGroupedAsset bool, groupedAsset bool, groupKey string, groupAnchor string, shortResponse bool) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
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
	client := mintrpc.NewMintClient(conn)
	// 指定资产版本是否为V1，默认V0
	var _assetVersion taprpc.AssetVersion
	if assetVersionIsV1 {
		_assetVersion = taprpc.AssetVersion_ASSET_VERSION_V1
	} else {
		_assetVersion = taprpc.AssetVersion_ASSET_VERSION_V0
	}
	// 指定资产类型是否为收藏品，默认正常
	var _assetType taprpc.AssetType
	if assetTypeIsCollectible {
		_assetType = taprpc.AssetType_COLLECTIBLE
	} else {
		_assetType = taprpc.AssetType_NORMAL
	}
	// AssetMeta Data
	_assetMetaDataByteSlice, _ := hex.DecodeString(assetMetaData)
	// 指定资产元数据类型
	var _assetMetaType taprpc.AssetMetaType
	if AssetMetaTypeIsJsonNotOpaque {
		_assetMetaType = taprpc.AssetMetaType_META_TYPE_JSON
	} else {
		_assetMetaType = taprpc.AssetMetaType_META_TYPE_OPAQUE
	}
	// GroupKey
	_groupKeyByteSlice, _ := hex.DecodeString(groupKey)
	// 构建请求
	request := &mintrpc.MintAssetRequest{
		Asset: &mintrpc.MintAsset{
			AssetVersion: _assetVersion,
			AssetType:    _assetType,
			Name:         name,
			AssetMeta: &taprpc.AssetMeta{
				Data: _assetMetaDataByteSlice,
				Type: _assetMetaType,
			},
			Amount:          uint64(amount),
			NewGroupedAsset: newGroupedAsset,
			GroupedAsset:    groupedAsset,
			GroupKey:        _groupKeyByteSlice,
			GroupAnchor:     groupAnchor,
		},
		ShortResponse: shortResponse,
	}
	// 得到响应
	response, err := client.MintAsset(context.Background(), request)
	if err != nil {
		log.Printf("mintrpc MintAsset Error: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}
