package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc/tapdevrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
)

// ImportProof
//
//	@Description: 尝试将证明文件导入守护进程。如果成功，将在磁盘上插入一个新资产，可使用指定的目标脚本密钥和内部密钥进行花费
//	@param proofFile
//	@param genesisPoint
//	@return *tapdevrpc.ImportProofResponse
//
// func ImportProof(proofFile, genesisPoint string) *tapdevrpc.ImportProofResponse {
func ImportProof(proofFile, genesisPoint string) bool {
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
	client := tapdevrpc.NewTapDevClient(conn)
	_proofFileByteSlice, _ := hex.DecodeString(proofFile)
	// 构建请求
	request := &tapdevrpc.ImportProofRequest{
		ProofFile:    _proofFileByteSlice,
		GenesisPoint: genesisPoint,
	}
	// 得到响应
	response, err := client.ImportProof(context.Background(), request)
	if err != nil {
		log.Printf("tapdevrpc QueryRfqAcceptedQuotes Error: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}
