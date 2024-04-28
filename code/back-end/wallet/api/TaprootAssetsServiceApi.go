package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
	"strings"
)

func AddrReceives() {

}

func BurnAsset() {

}

func DebugLevel() {

}

func DecodeAddr(addr string) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.DecodeAddrRequest{
		Addr: addr,
	}
	response, err := client.DecodeAddr(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc ListAssets Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

func DecodeProof() {

}

func ExportProof() {

}

// GetInfoOfTap
//
//	@Description: GetInfo returns the information for the node.
//	@return string
func GetInfoOfTap() string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.GetInfoRequest{}
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc GetInfo Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// ListAssets
//
//	@Description: ListAssets lists the set of assets owned by the target daemon.
//	@return string
func ListAssets(withWitness, includeSpent, includeLeased bool) string {
	response, err := listAssets(withWitness, includeSpent, includeLeased)
	if err != nil {
		fmt.Printf("%s taprpc ListAssets Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

func ListSimpleAssets(withWitness, includeSpent, includeLeased bool) string {
	response, err := listAssets(withWitness, includeSpent, includeLeased)
	if err != nil {
		fmt.Printf("%s taprpc ListAssets Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	var (
		simpleAssets []struct {
			AssetId string `json:"asset_id"`
			Name    string `json:"name"`
			Amount  uint64 `json:"amount"`
			Type    string `json:"type"`
		}
	)
	for _, asset := range response.Assets {
		set := true
		for index, s := range simpleAssets {
			if s.AssetId == hex.EncodeToString(asset.AssetGenesis.GetAssetId()) {
				simpleAssets[index].Amount = asset.Amount + s.Amount
				set = false
				break
			}
		}
		if !set {
			continue
		}
		simpleAssets = append(simpleAssets, struct {
			AssetId string `json:"asset_id"`
			Name    string `json:"name"`
			Amount  uint64 `json:"amount"`
			Type    string `json:"type"`
		}{
			AssetId: hex.EncodeToString(asset.AssetGenesis.GetAssetId()),
			Name:    asset.AssetGenesis.Name,
			Amount:  asset.Amount,
			Type:    asset.AssetGenesis.AssetType.String(),
		})
	}

	return MakeJsonResult(true, "", simpleAssets)
}

func FindAssetByAssetName(assetName string) string {
	//data := taprpc.ListAssetResponse{}

	var response = struct {
		Success bool                     `json:"success"`
		Error   string                   `json:"error"`
		Data    taprpc.ListAssetResponse `json:"data"`
	}{}
	list := ListAssets(false, false, false)
	err := json.Unmarshal([]byte(list), &response)

	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	if !response.Success {
		return MakeJsonResult(false, response.Error, nil)
	}
	var assets []*taprpc.Asset
	for _, asset := range response.Data.Assets {
		//if hex.EncodeToString(asset.AssetGenesis.GetAssetId()) == assetName {
		if asset.AssetGenesis.Name == assetName {
			assets = append(assets, asset)
		}
	}
	if len(assets) == 0 {
		return MakeJsonResult(false, "NOT_FOUND", nil)
	}
	return MakeJsonResult(true, "", assets)
}

func ListBalances() string {
	response, err := listBalances(false, nil, nil)
	if err != nil {
		MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

// ListGroups
//
//	@Description: ListGroups lists the asset groups known to the target daemon, and the assets held in each group.
//	@return string
func ListGroups() string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListGroupsRequest{}
	response, err := client.ListGroups(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc ListGroups Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// ListTransfers
//
//	@Description: ListTransfers lists outbound asset transfers tracked by the target daemon.
//	@return string
func ListTransfers() string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListTransfersRequest{}
	response, err := client.ListTransfers(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc ListTransfers Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// ListUtxos
//
//	@Description: ListUtxos lists the UTXOs managed by the target daemon, and the assets they hold.
//	@return string
func ListUtxos(includeLeased bool) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListUtxosRequest{
		IncludeLeased: includeLeased,
	}
	response, err := client.ListUtxos(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc ListUtxos Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// NewAddr
//
//	@Description:NewAddr makes a new address from the set of request params.
//	@return bool
func NewAddr(assetId string, amt int) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	_assetIdByteSlice, _ := hex.DecodeString(assetId)
	request := &taprpc.NewAddrRequest{
		AssetId: _assetIdByteSlice,
		Amt:     uint64(amt),
	}
	response, err := client.NewAddr(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc NewAddr Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

func QueryAddrs() {

}

// SendAsset
//
//	@Description:SendAsset uses one or multiple passed Taproot Asset address(es) to attempt to complete an asset send.
//	The method returns information w.r.t the on chain send, as well as the proof file information the receiver needs to fully receive the asset.
//	@return bool
//
// skipped function SendAsset with unsupported parameter or return types
func SendAsset(tapAddrs string, feeRate int) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	addrs := strings.Split(tapAddrs, ",")
	request := &taprpc.SendAssetRequest{
		TapAddrs: addrs,
		FeeRate:  uint32(feeRate),
	}
	response, err := client.SendAsset(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc SendAsset Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

func SubscribeReceiveAssetEventNtfns() {

}

func SubscribeSendAssetEventNtfns() {

}

func VerifyProof() {

}

// TapStopDaemon
//
//	@Description: StopDaemon will send a shutdown request to the interrupt handler, triggering a graceful shutdown of the daemon.
//	@return bool
func TapStopDaemon() bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.StopRequest{}
	response, err := client.StopDaemon(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc ListGroups Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

func decodeProof(proof []byte, depth uint32, withMetaReveal bool, withPrevWitnesses bool) (*taprpc.DecodeProofResponse, error) {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.DecodeProofRequest{
		RawProof:          proof,
		ProofAtDepth:      depth,
		WithMetaReveal:    withMetaReveal,
		WithPrevWitnesses: withPrevWitnesses,
	}
	response, err := client.DecodeProof(context.Background(), request)
	return response, err
}

func fetchAssetMeta(isHash bool, data string) (*taprpc.AssetMeta, error) {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)

	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.FetchAssetMetaRequest{}
	if isHash {
		request.Asset = &taprpc.FetchAssetMetaRequest_MetaHashStr{
			MetaHashStr: data,
		}
	} else {
		request.Asset = &taprpc.FetchAssetMetaRequest_AssetIdStr{
			AssetIdStr: data,
		}
	}
	response, err := client.FetchAssetMeta(context.Background(), request)
	return response, err
}

// ListBalances
//
//	@Description: ListBalances lists asset balances
//	@return string
func listBalances(useGroupKey bool, assetFilter, groupKeyFilter []byte) (*taprpc.ListBalancesResponse, error) {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListBalancesRequest{
		AssetFilter:    assetFilter,
		GroupKeyFilter: groupKeyFilter,
	}
	if useGroupKey {
		request.GroupBy = &taprpc.ListBalancesRequest_GroupKey{GroupKey: true}
	} else {
		request.GroupBy = &taprpc.ListBalancesRequest_AssetId{AssetId: true}
	}
	response, err := client.ListBalances(context.Background(), request)
	return response, err
}

func listAssets(withWitness, includeSpent, includeLeased bool) (*taprpc.ListAssetResponse, error) {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListAssetRequest{
		WithWitness:   withWitness,
		IncludeSpent:  includeSpent,
		IncludeLeased: includeLeased,
	}
	response, err := client.ListAssets(context.Background(), request)
	return response, err
}
