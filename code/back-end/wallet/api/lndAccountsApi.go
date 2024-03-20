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
	"log"
	"os"
	"path/filepath"
)

func ListAddress() string {
	grpcHost := base.QueryConfigByKey("lndhost")
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
	client := walletrpc.NewWalletKitClient(conn)
	request := &walletrpc.ListAddressesRequest{}
	response, err := client.ListAddresses(context.Background(), request)
	if err != nil {
		log.Printf("did not connect: %v", err)
		return ""
	}
	log.Printf("%v\n", response)
	return response.String()
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

// ListAccounts
//
//	@Description: 默认情况下会检索属于该钱包的所有账户。可以提供名称和密钥范围过滤器来过滤所有钱包账户，并只返回符合以下条件的账户
//	@return *walletrpc.ListAccountsResponse
//
// func ListAccounts() *walletrpc.ListAccountsResponse {
func ListAccounts() string {
	grpcHost := base.QueryConfigByKey("lndhost")
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
	client := walletrpc.NewWalletKitClient(conn)
	request := &walletrpc.ListAccountsRequest{}
	response, err := client.ListAccounts(context.Background(), request)
	if err != nil {
		log.Printf("watchtowerrpc ListAccounts err: %v", err)
		return ""
	}
	//log.Printf("%v", response)
	return response.String()
}

// ListLeases
//
//	@Description: 列出当前锁定的所有实用程序
//	@return *walletrpc.ListLeasesResponse
//
// func ListLeases() *walletrpc.ListLeasesResponse {
func ListLeases() string {
	grpcHost := base.QueryConfigByKey("lndhost")
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
	client := walletrpc.NewWalletKitClient(conn)
	request := &walletrpc.ListLeasesRequest{}
	response, err := client.ListLeases(context.Background(), request)
	if err != nil {
		log.Printf("watchtowerrpc ListLeases err: %v", err)
		return ""
	}
	//log.Printf("%v", response)
	return response.String()
}

// ListSweeps
//
//	@Description: 返回我们的节点产生的扫描事务列表。请注意，这些清扫事务可能尚未得到确认，因为我们是通过广播而非确认来记录清扫事务的
//	@return *walletrpc.ListSweepsResponse
//
// func ListSweeps() *walletrpc.ListSweepsResponse {
func ListSweeps() string {
	grpcHost := base.QueryConfigByKey("lndhost")
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
	client := walletrpc.NewWalletKitClient(conn)
	request := &walletrpc.ListSweepsRequest{}
	response, err := client.ListSweeps(context.Background(), request)
	if err != nil {
		log.Printf("watchtowerrpc ListSweeps err: %v", err)
		return ""
	}
	//log.Printf("%v", response)
	return response.String()
}

// ListUnspent
//
//	 @Description: 返回钱包可使用的所有 utxos 的列表，确认次数介于指定的最小值和最大值之间。
//		默认情况下，会列出所有 utxos。要只列出未确认的 utxos，请将 unconfirmed_only 设为 true
//	 @return *walletrpc.ListUnspentResponse
//
// func ListUnspent() *walletrpc.ListUnspentResponse {
func ListUnspent() string {
	grpcHost := base.QueryConfigByKey("lndhost")
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
	client := walletrpc.NewWalletKitClient(conn)
	request := &walletrpc.ListUnspentRequest{}
	response, err := client.ListUnspent(context.Background(), request)
	if err != nil {
		log.Printf("watchtowerrpc ListUnspent err: %v", err)
		return ""
	}
	//log.Printf("%v", response)
	return response.String()
}

// NextAddr
//
//	@Description: 返回钱包中下一个未使用的地址
//	@return *walletrpc.AddrResponse
//
// func NextAddr() *walletrpc.AddrResponse {
func NextAddr() string {
	grpcHost := base.QueryConfigByKey("lndhost")
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
	client := walletrpc.NewWalletKitClient(conn)
	request := &walletrpc.AddrRequest{}
	response, err := client.NextAddr(context.Background(), request)
	if err != nil {
		log.Printf("watchtowerrpc NextAddr err: %v", err)
		return ""
	}
	//log.Printf("%v", response)
	return response.String()
}
