package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"os"
	"path/filepath"
)

func GetNewAddress() bool {
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
	// 创建 WalletUnlocker 客户端
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.NewAddressRequest{}
	response, err := client.NewAddress(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc NewAddress err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

func GetWalletBalance() string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.WalletBalanceRequest{}
	response, err := client.WalletBalance(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc WalletBalance err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// AddInvoice
//
//	@Description: 试图在发票数据库中添加新发票。任何重复的发票都会被拒绝，因此所有发票都必须有唯一的付款预图像
//	@param value
//	@return string
//
// func AddInvoice(value int64) *lnrpc.AddInvoiceResponse {
func AddInvoice(value int64, memo string) string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.Invoice{
		Value: value,
		Memo:  memo,
	}
	response, err := client.AddInvoice(context.Background(), request)
	if err != nil {
		log.Printf("client.AddInvoice :%v", err)
		return ""
	}
	return response.String()
}

// ListInvoices
//
//	@Description: 返回数据库中当前存储的所有发票的列表。任何活动的调试发票都会被忽略。
//	它完全支持分页响应，允许用户通过 add_index 查询特定发票。
//	可以使用响应中包含的 first_index_offset 或 last_index_offset 字段作为下一个请求的 index_offset。默认情况下，
//	将返回创建的前 100 张发票。还可通过 Reversed 标志支持向后分页
//	@return string
func ListInvoices() string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ListInvoiceRequest{}
	response, err := client.ListInvoices(context.Background(), request)
	if err != nil {
		log.Printf("client.ListInvoice :%v", err)
		return ""
	}
	return response.String()
}

func LookupInvoice(rhash []byte) string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.PaymentHash{
		RHash: rhash,
	}
	response, err := client.LookupInvoice(context.Background(), request)
	if err != nil {
		log.Printf("client.ListInvoice :%v", err)
		return ""
	}
	return response.String()
}

// func AbandonChannel() *lnrpc.AbandonChannelResponse {
func AbandonChannel() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.AbandonChannelRequest{}
	response, err := client.AbandonChannel(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc AbandonChannel err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// func BatchOpenChannel() *lnrpc.BatchOpenChannelResponse {
func BatchOpenChannel() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.BatchOpenChannelRequest{}
	response, err := client.BatchOpenChannel(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc BatchOpenChannel err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// func ChannelAcceptor() *lnrpc.Lightning_ChannelAcceptorClient {
func ChannelAcceptor() bool {
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
	client := lnrpc.NewLightningClient(conn)
	//request := &lnrpc.ChannelAcceptRequest{}
	//response, err := client.ChannelAcceptor(context.Background(), request)
	stream, err := client.ChannelAcceptor(context.Background())
	if err != nil {
		log.Printf("lnrpc ChannelAcceptor err: %v", err)
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

// ChannelBalance
//
//	@Description: 返回所有开放渠道的资金总额报告，按本地/远程、待结算的本地/远程和未结算的本地/远程余额分类
//	@return bool
//
// func ChannelBalance() *lnrpc.ChannelBalanceResponse {
func ChannelBalance() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChannelBalanceRequest{}
	response, err := client.ChannelBalance(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ChannelBalance err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// func CheckMacaroonPermissions() *lnrpc.CheckMacPermResponse {
func CheckMacaroonPermissions() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.CheckMacPermRequest{}
	response, err := client.CheckMacaroonPermissions(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc CheckMacaroonPermissions err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// CloseChannel
//
//	 @Description: 试图关闭一个由其通道输出点（ChannelPoint）标识的活动通道。
//		该方法的操作还可以在超时后尝试强制关闭不活动的对等设备。如果请求非强制关闭（合作关闭），
//		用户可以指定关闭交易确认前的目标区块数或手动费率。
//		如果两者都未指定，则使用默认的宽松区块确认目标。
//	 @param fundingTxidStr
//	 @param outputIndex
func CloseChannel(fundingTxidStr string, outputIndex int) bool {
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
	client := lnrpc.NewLightningClient(conn)

	request := &lnrpc.CloseChannelRequest{
		ChannelPoint: &lnrpc.ChannelPoint{
			FundingTxid: &lnrpc.ChannelPoint_FundingTxidStr{FundingTxidStr: fundingTxidStr},
			OutputIndex: uint32(outputIndex),
		},
	}
	stream, err := client.CloseChannel(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc CloseChannel err: %v", err)
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

// ClosedChannels
//
//	@Description: 返回该节点参与的所有封闭通道的描述
//	@return *lnrpc.ClosedChannelsResponse
//
// func ClosedChannels() *lnrpc.ClosedChannelsResponse {
func ClosedChannels() string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ClosedChannelsRequest{}
	response, err := client.ClosedChannels(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ClosedChannels err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// func ExportAllChannelBackups() *lnrpc.ChanBackupSnapshot {
func ExportAllChannelBackups() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChanBackupExportRequest{}
	response, err := client.ExportAllChannelBackups(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ChanBackupExportRequest err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// func ExportChannelBackup() *lnrpc.ChannelBackup {
func ExportChannelBackup() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ExportChannelBackupRequest{}
	response, err := client.ExportChannelBackup(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ExportChannelBackup err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// GetChanInfo
//
//	@Description: 返回指定通道的最新认证网络公告，该通道由通道 ID 标识：一个 8 字节整数，用于唯一标识区块链中交易资金输出的位置
//	@param chanId
//	@return *lnrpc.ChannelEdge
//
// func GetChanInfo(chanId uint64) *lnrpc.ChannelEdge {
func GetChanInfo(chanId int) string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChanInfoRequest{
		ChanId: uint64(chanId),
	}
	response, err := client.GetChanInfo(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc GetChanInfo err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// ListChannels
//
//	@Description: 返回该节点参与的所有开放通道的描述。
//	@return *lnrpc.ListChannelsResponse
//
// func ListChannels() *lnrpc.ListChannelsResponse {
func ListChannels() string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ListChannelsRequest{}
	response, err := client.ListChannels(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ListChannels err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// func OpenChannelSync() *lnrpc.ChannelPoint {
func OpenChannelSync() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.OpenChannelRequest{}
	response, err := client.OpenChannelSync(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc OpenChannelSync err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// OpenChannel
//
//	@Description: 会尝试向远程对等方打开请求中指定的单一注资通道。
//	用户可以指定确认资金交易的目标区块数，或为资金交易手动设定费率。
//	如果两者都未指定，则使用宽松的区块确认目标。每次 OpenStatusUpdate 都会返回正在进行的通道的待定通道 ID。
//	根据 OpenChannelRequest 中指定的参数，该待定通道 ID 可用于手动推进通道资金流。
//	@param nodePubkey
//	@param localFundingAmount
//
// func OpenChannel(nodePubkey string, localFundingAmount int64) *lnrpc.OpenStatusUpdate {
func OpenChannel(nodePubkey string, localFundingAmount int64) bool {
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
	client := lnrpc.NewLightningClient(conn)
	_nodePubkeyByteSlice, _ := hex.DecodeString(nodePubkey)
	request := &lnrpc.OpenChannelRequest{
		NodePubkey:         _nodePubkeyByteSlice,
		LocalFundingAmount: localFundingAmount,
	}
	stream, err := client.OpenChannel(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc OpenChannel err: %v", err)
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

// PendingChannels
//
//	 @Description: 返回当前被视为 "待定 "的所有通道的列表。
//		如果一个通道已完成筹资工作流程，正在等待筹资 txn 的确认，
//		或正在关闭（合作或非合作启动），则该通道为 "待定 "通道。
//	 @return *lnrpc.PendingChannelsResponse
//
// func PendingChannels() *lnrpc.PendingChannelsResponse {
func PendingChannels() string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.PendingChannelsRequest{}
	response, err := client.PendingChannels(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc PendingChannels err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// func RestoreChannelBackups() *lnrpc.RestoreBackupResponse {
func RestoreChannelBackups() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.RestoreChanBackupRequest{}
	response, err := client.RestoreChannelBackups(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc RestoreChannelBackups err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// func SubscribeChannelBackups() *lnrpc.Lightning_SubscribeChannelBackupsClient {
func SubscribeChannelBackups() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChannelBackupSubscription{}
	stream, err := client.SubscribeChannelBackups(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc SubscribeChannelBackups err: %v", err)
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

// func SubscribeChannelEvents() *lnrpc.Lightning_SubscribeChannelEventsClient {
func SubscribeChannelEvents() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChannelEventSubscription{}
	stream, err := client.SubscribeChannelEvents(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc  err: %v", err)
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

// func SubscribeChannelGraph() *lnrpc.Lightning_SubscribeChannelGraphClient {
func SubscribeChannelGraph() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.GraphTopologySubscription{}
	stream, err := client.SubscribeChannelGraph(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc SubscribeChannelGraph err: %v", err)
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

// func UpdateChannelPolicy() *lnrpc.PolicyUpdateResponse {
func UpdateChannelPolicy() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.PolicyUpdateRequest{}
	response, err := client.UpdateChannelPolicy(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc UpdateChannelPolicy err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// func VerifyChanBackup() *lnrpc.VerifyChanBackupResponse {
func VerifyChanBackup() bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChanBackupSnapshot{}
	response, err := client.VerifyChanBackup(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc VerifyChanBackup err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// ConnectPeer
//
//	@Description: 试图建立与远程对等节点的连接。这属于网络层面，用于节点之间的通信。这有别于与对等节点建立通道。
//	@param pubkey
//	@param host
//	@return *lnrpc.ConnectPeerResponse
//
// func ConnectPeer(pubkey, host string) *lnrpc.ConnectPeerResponse {
func ConnectPeer(pubkey, host string) bool {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ConnectPeerRequest{
		Addr: &lnrpc.LightningAddress{
			Pubkey: pubkey,
			Host:   host,
		},
	}
	response, err := client.ConnectPeer(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ConnectPeer err: %v", err)
		return false
	}
	log.Printf("%v\n", response)
	return true
}

// func EstimateFee(addr string, amount int64) *lnrpc.EstimateFeeResponse {
func EstimateFee(addr string, amount int64) string {
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
	client := lnrpc.NewLightningClient(conn)
	addrToAmount := make(map[string]int64)
	addrToAmount[addr] = amount
	request := &lnrpc.EstimateFeeRequest{
		AddrToAmount: addrToAmount,
	}
	response, err := client.EstimateFee(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ConnectPeer err: %v", err)
		return ""
	}
	//log.Printf("%v\n", response)
	return response.String()
}

// func DecodePayReq(pay_req string) *lnrpc.PayReq {
func DecodePayReq(pay_req string) int64 {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.PayReqString{
		PayReq: pay_req,
	}
	response, err := client.DecodePayReq(context.Background(), request)
	if err != nil {
		log.Printf("client.AddInvoice :%v", err)
		return -1
	}
	//log.Printf("%v\n", response)
	return response.NumSatoshis
}

func SendPaymentSync(invoice string) string {
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
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.SendRequest{
		PaymentRequest: invoice,
		//Amt:            amt,
	}
	stream, err := client.SendPaymentSync(context.Background(), request)
	if err != nil {
		log.Printf("client.SendPaymentSync :%v", err)
		return "false"
	}
	log.Printf(stream.String())
	return hex.EncodeToString(stream.PaymentHash)
}

func SendCoins(addr string, amount int64) string {
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

	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.SendCoinsRequest{
		Addr:   addr,
		Amount: amount,
	}
	response, err := client.SendCoins(context.Background(), request)
	if err != nil {
		log.Printf("lnrpc ConnectPeer err: %v", err)
		return "false"
	}
	log.Printf("%v\n", response)
	return response.Txid
}
