package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/assetwalletrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
)

// AnchorVirtualPsbts NOT COMPLETED
func AnchorVirtualPsbts(virtualPsbts [][]byte) *taprpc.SendAssetResponse {
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	// 构建请求
	request := &assetwalletrpc.AnchorVirtualPsbtsRequest{
		VirtualPsbts: virtualPsbts,
	}
	// 得到响应
	response, err := client.AnchorVirtualPsbts(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc AnchorVirtualPsbts Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// FundVirtualPsbt NOT COMPLETED
func FundVirtualPsbt(isPsbtNotRaw bool, psbt ...string) *assetwalletrpc.FundVirtualPsbtResponse { // 读取参数
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	// 构建请求
	request := &assetwalletrpc.FundVirtualPsbtRequest{}
	if isPsbtNotRaw {
		_psbtByteSlice, _ := hex.DecodeString(psbt[0])
		request.Template = &assetwalletrpc.FundVirtualPsbtRequest_Psbt{Psbt: _psbtByteSlice}
	} else {
		request.Template = &assetwalletrpc.FundVirtualPsbtRequest_Raw{
			Raw: &assetwalletrpc.TxTemplate{
				//Inputs:	[]*assetwalletrpc.PrevId{},
				Recipients: nil,
			}}
	}

	// 得到响应
	response, err := client.FundVirtualPsbt(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc FundVirtualPsbt Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// NextInternalKey
//
//	 @Description: 获取给定密钥族的下一个内部密钥，并将其作为内部密钥存储在数据库中，
//		以确保以后导入证明时能将其识别为本地密钥。虽然内部密钥也可用作脚本密钥的内部密钥，
//		但建议使用 NextScriptKey RPC，以确保经过调整的 Taproot 输出密钥也能被识别为本地密钥。
//	 @param keyFamily
//	 @return *assetwalletrpc.NextInternalKeyResponse
func NextInternalKey(keyFamily uint32) *assetwalletrpc.NextInternalKeyResponse { // 读取参数
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	// 构建请求
	request := &assetwalletrpc.NextInternalKeyRequest{
		KeyFamily: keyFamily,
	}
	// 得到响应
	response, err := client.NextInternalKey(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc NextInternalKey Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// NextScriptKey
//
//	@Description: 导出下一个脚本密钥（及其相应的内部密钥），并将它们都存储在数据库中，以确保以后导入证明时能将它们识别为本地密钥
//	@param keyFamily
//	@return *assetwalletrpc.NextScriptKeyResponse
func NextScriptKey(keyFamily uint32) *assetwalletrpc.NextScriptKeyResponse { // 读取参数
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	// 构建请求
	request := &assetwalletrpc.NextScriptKeyRequest{
		KeyFamily: keyFamily,
	}
	// 得到响应
	response, err := client.NextScriptKey(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc NextScriptKey Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// ProveAssetOwnership
//
//	 @Description: 在资产转换证明中嵌入所有权证明。
//		该所有权证明是一个签名的虚拟交易，由一个有效的见证人花费该资产，以证明证明者拥有可以花费该资产的密钥
//	 @param assetId
//	 @param scriptKey
//	 @return *assetwalletrpc.ProveAssetOwnershipResponse
func ProveAssetOwnership(assetId, scriptKey string) *assetwalletrpc.ProveAssetOwnershipResponse {
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_assetIdByteSlice, _ := hex.DecodeString(assetId)
	_scriptKeyByteSlice, _ := hex.DecodeString(scriptKey)
	// 构建请求
	request := &assetwalletrpc.ProveAssetOwnershipRequest{
		AssetId:   _assetIdByteSlice,
		ScriptKey: _scriptKeyByteSlice,
		Outpoint:  nil,
	}
	// 得到响应
	response, err := client.ProveAssetOwnership(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc ProveAssetOwnership Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// QueryInternalKey
//
//	 @Description: 返回给定内部密钥的密钥描述符
//	 @param internalKey
//	 @return *assetwalletrpc.QueryInternalKeyResponse
//		rpc error: code = Unknown desc = unknown request
func QueryInternalKey(internalKey string) *assetwalletrpc.QueryInternalKeyResponse {
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_internalKeyByteSlice, _ := hex.DecodeString(internalKey)

	// 构建请求
	request := &assetwalletrpc.QueryInternalKeyRequest{
		InternalKey: _internalKeyByteSlice,
	}
	// 得到响应
	response, err := client.QueryInternalKey(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc QueryInternalKey Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// QueryScriptKey
//
//	@Description: 返回给定调整后脚本密钥的完整脚本密钥描述符
//	@param tweakedScriptKey
//	@return *assetwalletrpc.QueryScriptKeyResponse
func QueryScriptKey(tweakedScriptKey string) *assetwalletrpc.QueryScriptKeyResponse {
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_tweakedScriptKeyByteSlice, _ := hex.DecodeString(tweakedScriptKey)
	// 构建请求
	request := &assetwalletrpc.QueryScriptKeyRequest{
		TweakedScriptKey: _tweakedScriptKeyByteSlice,
	}
	// 得到响应
	response, err := client.QueryScriptKey(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc QueryScriptKey Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// RemoveUTXOLease
//
//	@Description: 移除给定受管 UTXO 的租用/锁定/保留
//	@return *assetwalletrpc.RemoveUTXOLeaseResponse
func RemoveUTXOLease() *assetwalletrpc.RemoveUTXOLeaseResponse {
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	// 构建请求
	request := &assetwalletrpc.RemoveUTXOLeaseRequest{
		Outpoint: nil,
	}
	// 得到响应
	response, err := client.RemoveUTXOLease(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc RemoveUTXOLease Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// SignVirtualPsbt
//
//	@Description: 签署虚拟交易的输入，并准备输入和输出的承诺
//	@param fundedPsbt
//	@return *assetwalletrpc.SignVirtualPsbtResponse
func SignVirtualPsbt(fundedPsbt string) *assetwalletrpc.SignVirtualPsbtResponse {
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_fundedPsbtByteSlice, _ := hex.DecodeString(fundedPsbt)
	// 构建请求
	request := &assetwalletrpc.SignVirtualPsbtRequest{
		FundedPsbt: _fundedPsbtByteSlice,
	}
	// 得到响应
	response, err := client.SignVirtualPsbt(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc SignVirtualPsbt Error: %v", err)
		return nil
	}
	// 处理结果
	return response
}

// VerifyAssetOwnership
//
//	@Description: 验证给定资产过渡证明中嵌入的资产所有权证明，
//	如果证明有效，则返回 true
//	validProof := response.ValidProof
//	@return *assetwalletrpc.VerifyAssetOwnershipResponse
func VerifyAssetOwnership(proofWithWitness string) *assetwalletrpc.VerifyAssetOwnershipResponse {
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
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_proofWithWitnessByteSlice, _ := hex.DecodeString(proofWithWitness)
	// 构建请求
	request := &assetwalletrpc.VerifyAssetOwnershipRequest{
		ProofWithWitness: _proofWithWitnessByteSlice,
	}
	// 得到响应
	response, err := client.VerifyAssetOwnership(context.Background(), request)
	if err != nil {
		log.Printf("assetwalletrpc VerifyAssetOwnership Error: %v", err)
		return nil
	}
	// bool
	//response.ValidProof
	// 处理结果
	return response
}
