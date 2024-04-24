package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
)

func AddFederationServer() {}

func AssetLeafKeys() {}

func AssetLeaves(id string) string {
	response, err := assetLeaves(false, id, universerpc.ProofType_PROOF_TYPE_ISSUANCE)
	if err != nil {
		fmt.Printf("%s universerpc AssetLeaves Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}

	if response.Leaves == nil {
		return MakeJsonResult(false, "NOT_FOUND", nil)
	}

	return MakeJsonResult(true, "", response)
}

func GetAssetInfo(id string) string {
	//从宇宙中读取发行记录
	response, err := assetLeaves(false, id, universerpc.ProofType_PROOF_TYPE_ISSUANCE)
	if err != nil {
		fmt.Printf("%s universerpc AssetLeaves Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	if response.Leaves == nil {
		return MakeJsonResult(false, "NOT_FOUND", nil)
	}
	//解析证明文件
	proof, err := decodeProof(response.Leaves[0].Proof, 0, true, false)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	// 获取时间戳
	block, err := getblock(proof.DecodedProof.Asset.ChainAnchor.AnchorBlockHash)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	msgBlock := &wire.MsgBlock{}
	blockReader := bytes.NewReader(block.RawBlock)
	err = msgBlock.Deserialize(blockReader)
	timeStamp := msgBlock.Header.Timestamp

	// 转换为Unix时间戳
	createTime := timeStamp.Unix()

	var assetInfo = struct {
		Asset      *taprpc.Asset `json:"asset"`
		Meta       string        `json:"meta"`
		CreateTime int64         `json:"createTime"`
	}{
		Asset:      proof.DecodedProof.Asset,
		Meta:       hex.EncodeToString(proof.DecodedProof.MetaReveal.Data),
		CreateTime: createTime,
	}
	return MakeJsonResult(true, "", assetInfo)
}

func AssetRoots() {}

func DeleteAssetRoot() {}

func DeleteFederationServer() {}

// UniverseInfo
//
//	@Description: Info returns a set of information about the current state of the Universe.
//	@return string
func UniverseInfo() string {
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
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := universerpc.NewUniverseClient(conn)
	request := &universerpc.InfoRequest{}
	response, err := client.Info(context.Background(), request)
	if err != nil {
		fmt.Printf("%s universerpc Info Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

func InsertProof() {}

// ListFederationServers
//
//	@Description: ListFederationServers lists the set of servers that make up the federation of the local Universe server.
//	This servers are used to push out new proofs, and also periodically call sync new proofs from the remote server.
//	@return string
func ListFederationServers() string {
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
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := universerpc.NewUniverseClient(conn)
	request := &universerpc.ListFederationServersRequest{}
	response, err := client.ListFederationServers(context.Background(), request)
	if err != nil {
		fmt.Printf("%s universerpc ListFederationServers Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

func MultiverseRoot() {}

func QueryAssetRoots() {}

func QueryAssetStats(assetId string) string {
	response, err := queryAssetStats(assetId)
	if err != nil {
		fmt.Printf("%s universerpc QueryAssetStats Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), "")
	}
	return MakeJsonResult(true, "", response)
}

func QueryEvents() {}

func QueryFederationSyncConfig() {}

func QueryProof() {}

func SetFederationSyncConfig() {}

func SyncUniverse(universeHost string, asset_id string) string {
	var targets []*universerpc.SyncTarget
	universeID := &universerpc.ID{
		Id: &universerpc.ID_AssetIdStr{
			AssetIdStr: asset_id,
		},
		ProofType: 1,
	}
	if universeID != nil {
		targets = append(targets, &universerpc.SyncTarget{
			Id: universeID,
		})
	}
	if universeHost == "" {
		universeHost = "testnet.universe.lightning.finance:10029"
	}
	response, err := syncUniverse(universeHost, targets, 0)
	if err != nil {
		return MakeJsonResult(false, err.Error(), "")
	}
	if response.SyncedUniverses == nil {
		return MakeJsonResult(false, "NOT_NEW_DATA", "")
	}
	return MakeJsonResult(true, "", response)

}

func syncUniverse(universeHost string, syncTargets []*universerpc.SyncTarget, syncMode universerpc.UniverseSyncMode) (*universerpc.SyncResponse, error) {

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
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	request := &universerpc.SyncRequest{
		UniverseHost: universeHost,
		SyncMode:     syncMode,
		SyncTargets:  syncTargets,
	}
	client := universerpc.NewUniverseClient(conn)
	response, err := client.SyncUniverse(context.Background(), request)
	return response, err
}

func UniverseStats() {}

func assetLeaves(isGroup bool, id string, prooftype universerpc.ProofType) (*universerpc.AssetLeafResponse, error) {
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
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)

	request := &universerpc.ID{
		ProofType: prooftype,
	}

	if isGroup {
		groupKey := &universerpc.ID_GroupKeyStr{
			GroupKeyStr: id,
		}
		request.Id = groupKey
	} else {
		AssetId := &universerpc.ID_AssetIdStr{
			AssetIdStr: id,
		}
		request.Id = AssetId
	}

	client := universerpc.NewUniverseClient(conn)
	response, err := client.AssetLeaves(context.Background(), request)
	return response, err
}

func queryAssetStats(assetId string) (*universerpc.UniverseAssetStats, error) {
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
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	id, err := hex.DecodeString(assetId)
	client := universerpc.NewUniverseClient(conn)
	request := &universerpc.AssetStatsQuery{
		AssetIdFilter: id,
	}
	response, err := client.QueryAssetStats(context.Background(), request)
	return response, err
}
