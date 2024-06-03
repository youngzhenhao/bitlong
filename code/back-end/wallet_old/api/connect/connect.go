package connect

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
	"time"
)

type rpccfg struct {
	grpcHost     string
	tlsCertPath  string
	macaroonPath string
}

type ConnCfg struct {
	isInit  bool
	Lndcfg  rpccfg
	Tapdcfg rpccfg
	LitdCfg rpccfg
}

var connCfg = ConnCfg{
	isInit: false,
}

func LoadConfig() {
	connCfg.Lndcfg.grpcHost = base.QueryConfigByKey("lndhost")
	connCfg.Lndcfg.tlsCertPath = filepath.Join(base.Configure("lnd"), "tls.cert")
	connCfg.Lndcfg.macaroonPath = filepath.Join(base.Configure("lnd"), "data", "chain", "bitcoin", base.NetWork, "admin.macaroon")

	connCfg.Tapdcfg.grpcHost = base.QueryConfigByKey("taproothost")
	connCfg.Tapdcfg.tlsCertPath = filepath.Join(base.Configure("lit"), "tls.cert")
	connCfg.Tapdcfg.macaroonPath = filepath.Join(base.Configure("tapd"), "data", base.NetWork, "admin.macaroon")

	connCfg.LitdCfg.grpcHost = base.QueryConfigByKey("litdhost")
	connCfg.LitdCfg.tlsCertPath = filepath.Join(base.Configure("lit"), "tls.cert")
	connCfg.LitdCfg.macaroonPath = filepath.Join(base.Configure("lit"), base.NetWork, "lit.macaroon")
	log.Printf(" grpc config: %s", connCfg)
}

func GetConnection(grpcTarget string, isNoMacaroon bool) (*grpc.ClientConn, func(), error) {
	if !connCfg.isInit {
		LoadConfig()
		connCfg.isInit = true
	}
	cfg := rpccfg{}
	//select grpc config by grpcTarget
	switch grpcTarget {
	case "lnd":
		cfg = connCfg.Lndcfg
	case "tapd":
		cfg = connCfg.Tapdcfg
	case "litd":
		cfg = connCfg.LitdCfg
	default:
		return nil, nil, fmt.Errorf("grpcTarget not found")
	}
	// get tls cert and macaroon
	creds := NewTlsCert(cfg.tlsCertPath)
	var (
		conn *grpc.ClientConn
		err  error
	)
	// if isNoMacaroon is true, then we don't need to add macaroon to grpc connection
	if isNoMacaroon {
		conn, err = grpc.Dial(cfg.grpcHost, grpc.WithTransportCredentials(creds))
	} else {
		macaroon := GetMacaroon(cfg.macaroonPath)
		conn, err = grpc.Dial(cfg.grpcHost, grpc.WithTransportCredentials(creds),
			grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	}

	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	cleanUp := func() {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}
	return conn, cleanUp, err
}

type MacaroonCredential struct {
	macaroon string
}

func NewMacaroonCredential(macaroon string) *MacaroonCredential {
	return &MacaroonCredential{macaroon: macaroon}
}

func (c *MacaroonCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"macaroon": c.macaroon}, nil
}

func (c *MacaroonCredential) RequireTransportSecurity() bool {
	return true
}

func NewTlsCert(tlsCertPath string) credentials.TransportCredentials {
	cert, err := os.ReadFile(tlsCertPath)
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
	return creds
}

func GetMacaroon(macaroonPath string) string {
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	return macaroon
}

func GetTimeNow() string {
	return time.Now().Format("2006/01/02 15:04:05")
}
