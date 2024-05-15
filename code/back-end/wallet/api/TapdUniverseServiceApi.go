package api

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"path/filepath"
)

func AddFederationServer() {}

func assetLeafKeys(id string, proofType universerpc.ProofType) (*universerpc.AssetLeafKeyResponse, error) {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
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
	client := universerpc.NewUniverseClient(conn)
	request := &universerpc.AssetLeafKeysRequest{
		Id: &universerpc.ID{
			Id: &universerpc.ID_AssetIdStr{
				AssetIdStr: id,
			},
			ProofType: proofType,
		},
		//Offset:    0,
		//Limit:     0,
		//Direction: 0,
	}
	response, err := client.AssetLeafKeys(context.Background(), request)
	if err != nil {
		fmt.Printf("%s universerpc Info Error: %v\n", GetTimeNow(), err)
		return nil, err
	}
	return response, nil
}

func AssetLeafKeys(id string, proofType string) string {
	var _proofType universerpc.ProofType
	if proofType == "issuance" || proofType == "ISSUANCE" || proofType == "PROOF_TYPE_ISSUANCE" {
		_proofType = universerpc.ProofType_PROOF_TYPE_ISSUANCE
	} else if proofType == "transfer" || proofType == "TRANSFER" || proofType == "PROOF_TYPE_TRANSFER" {
		_proofType = universerpc.ProofType_PROOF_TYPE_TRANSFER
	} else {
		_proofType = universerpc.ProofType_PROOF_TYPE_UNSPECIFIED
	}
	response, err := assetLeafKeys(id, _proofType)
	if err != nil {
		fmt.Printf("%s universerpc AssetLeafKeys Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	if len(response.AssetKeys) == 0 {
		return MakeJsonResult(false, "Result length is zero.", nil)
	}
	return MakeJsonResult(true, "", processAssetKey(response))
}

type AssetKey struct {
	OpStr          string `json:"op_str"`
	ScriptKeyBytes string `json:"script_key_bytes"`
}

func processAssetKey(response *universerpc.AssetLeafKeyResponse) *[]AssetKey {
	var assetKey []AssetKey
	for _, keys := range response.AssetKeys {
		assetKey = append(assetKey, AssetKey{
			OpStr:          keys.GetOpStr(),
			ScriptKeyBytes: hex.EncodeToString(keys.GetScriptKeyBytes()),
		})
	}
	return &assetKey
}

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
	response, err := assetLeaves(false, id, universerpc.ProofType_PROOF_TYPE_ISSUANCE)
	if err != nil {
		fmt.Printf("%s universerpc AssetLeaves Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	if response.Leaves == nil {
		return MakeJsonResult(false, "NOT_FOUND", nil)
	}
	proof, err := decodeProof(response.Leaves[0].Proof, 0, true, false)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	block, err := GetBlock(proof.DecodedProof.Asset.ChainAnchor.AnchorBlockHash)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	msgBlock := &wire.MsgBlock{}
	blockReader := bytes.NewReader(block.RawBlock)
	err = msgBlock.Deserialize(blockReader)
	timeStamp := msgBlock.Header.Timestamp
	createTime := timeStamp.Unix()

	assetId := hex.EncodeToString(proof.DecodedProof.Asset.AssetGenesis.GetAssetId())
	assetType := proof.DecodedProof.Asset.AssetGenesis.AssetType.String()
	assetIsGroup := false
	if proof.DecodedProof.Asset.AssetGroup != nil {
		assetIsGroup = true
	}
	amount := proof.DecodedProof.Asset.Amount
	var newMeta Meta

	meta := hex.EncodeToString(proof.DecodedProof.MetaReveal.Data)
	newMeta.GetMetaFromStr(meta)

	var assetInfo = struct {
		AssetId      string `json:"assetId"`
		AssetType    string `json:"assetType"`
		AssetIsGroup bool   `json:"assetIsGroup"`
		Amount       uint64 `json:"amount"`
		Meta         string `json:"meta"`
		CreateTime   int64  `json:"createTime"`
	}{
		AssetId:      assetId,
		AssetType:    assetType,
		AssetIsGroup: assetIsGroup,
		Amount:       amount,
		Meta:         newMeta.Description,
		CreateTime:   createTime,
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
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
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
	client := universerpc.NewUniverseClient(conn)
	request := &universerpc.InfoRequest{}
	response, err := client.Info(context.Background(), request)
	if err != nil {
		fmt.Printf("%s universerpc Info Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
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
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
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

func QueryAssetRoots(id string) string {
	response, err := queryAssetRoot(id)
	if err != nil {
		fmt.Printf("%s universerpc AssetRoots Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

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
	return MakeJsonResult(true, "", response)

}

func UniverseStats() {}

func queryAssetRoot(id string) (*universerpc.QueryRootResponse, error) {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
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

	requst := &universerpc.AssetRootQuery{
		Id: &universerpc.ID{
			Id: &universerpc.ID_AssetIdStr{
				AssetIdStr: id,
			},
		},
	}
	client := universerpc.NewUniverseClient(conn)
	response, err := client.QueryAssetRoots(context.Background(), requst)
	return response, err
}

func assetLeaves(isGroup bool, id string, proofType universerpc.ProofType) (*universerpc.AssetLeafResponse, error) {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
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

	request := &universerpc.ID{
		ProofType: proofType,
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
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
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
	id, err := hex.DecodeString(assetId)
	client := universerpc.NewUniverseClient(conn)
	request := &universerpc.AssetStatsQuery{
		AssetIdFilter: id,
	}
	response, err := client.QueryAssetStats(context.Background(), request)
	return response, err
}

func syncUniverse(universeHost string, syncTargets []*universerpc.SyncTarget, syncMode universerpc.UniverseSyncMode) (*universerpc.SyncResponse, error) {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
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
	request := &universerpc.SyncRequest{
		UniverseHost: universeHost,
		SyncMode:     syncMode,
		SyncTargets:  syncTargets,
	}
	client := universerpc.NewUniverseClient(conn)
	response, err := client.SyncUniverse(context.Background(), request)
	return response, err
}
