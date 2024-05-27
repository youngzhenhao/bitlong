package api

import (
	"context"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/mintrpc"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"strconv"
	"strings"
	"trade/config"
	"trade/models"
	"trade/utils"
)

func assetLeaves(isGroup bool, id string, proofType universerpc.ProofType) (*universerpc.AssetLeafResponse, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
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

func assetLeavesSpecified(id string, proofType string) *universerpc.AssetLeafResponse {
	var _proofType universerpc.ProofType
	if proofType == "issuance" || proofType == "ISSUANCE" || proofType == "PROOF_TYPE_ISSUANCE" {
		_proofType = universerpc.ProofType_PROOF_TYPE_ISSUANCE
	} else if proofType == "transfer" || proofType == "TRANSFER" || proofType == "PROOF_TYPE_TRANSFER" {
		_proofType = universerpc.ProofType_PROOF_TYPE_TRANSFER
	} else {
		_proofType = universerpc.ProofType_PROOF_TYPE_UNSPECIFIED
	}
	response, err := assetLeaves(false, id, _proofType)
	if err != nil {
		utils.LogError("Universe AssetLeaves Error", err)
		return nil
	}
	return response
}

func processAssetIssuanceLeaf(response *universerpc.AssetLeafResponse) *models.AssetIssuanceLeaf {
	if response == nil {
		return nil
	}
	return &models.AssetIssuanceLeaf{
		Version:            response.Leaves[0].Asset.Version.String(),
		GenesisPoint:       response.Leaves[0].Asset.AssetGenesis.GenesisPoint,
		Name:               response.Leaves[0].Asset.AssetGenesis.Name,
		MetaHash:           hex.EncodeToString(response.Leaves[0].Asset.AssetGenesis.MetaHash),
		AssetID:            hex.EncodeToString(response.Leaves[0].Asset.AssetGenesis.AssetId),
		AssetType:          response.Leaves[0].Asset.AssetGenesis.AssetType.String(),
		GenesisOutputIndex: int(response.Leaves[0].Asset.AssetGenesis.OutputIndex),
		Amount:             int(response.Leaves[0].Asset.Amount),
		LockTime:           int(response.Leaves[0].Asset.LockTime),
		RelativeLockTime:   int(response.Leaves[0].Asset.RelativeLockTime),
		ScriptVersion:      int(response.Leaves[0].Asset.ScriptVersion),
		ScriptKey:          hex.EncodeToString(response.Leaves[0].Asset.ScriptKey),
		ScriptKeyIsLocal:   response.Leaves[0].Asset.ScriptKeyIsLocal,
		IsSpent:            response.Leaves[0].Asset.IsSpent,
		LeaseOwner:         hex.EncodeToString(response.Leaves[0].Asset.LeaseOwner),
		LeaseExpiry:        int(response.Leaves[0].Asset.LeaseExpiry),
		IsBurn:             response.Leaves[0].Asset.IsBurn,
		Proof:              hex.EncodeToString(response.Leaves[0].Proof),
	}
}

func assetLeafIssuanceInfo(id string) *models.AssetIssuanceLeaf {
	response := assetLeavesSpecified(id, universerpc.ProofType_PROOF_TYPE_ISSUANCE.String())
	if response == nil {
		utils.LogInfo("Universe asset leaves issuance error.")
		return nil
	}
	return processAssetIssuanceLeaf(response)
}

func mintAsset(assetVersionIsV1 bool, assetTypeIsCollectible bool, name string, assetMetaData string, AssetMetaTypeIsJsonNotOpaque bool, amount int, newGroupedAsset bool, groupedAsset bool, groupKey string, groupAnchor string, shortResponse bool) (*mintrpc.MintAssetResponse, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
	client := mintrpc.NewMintClient(conn)
	var _assetVersion taprpc.AssetVersion
	if assetVersionIsV1 {
		_assetVersion = taprpc.AssetVersion_ASSET_VERSION_V1
	} else {
		_assetVersion = taprpc.AssetVersion_ASSET_VERSION_V0
	}
	var _assetType taprpc.AssetType
	if assetTypeIsCollectible {
		_assetType = taprpc.AssetType_COLLECTIBLE
	} else {
		_assetType = taprpc.AssetType_NORMAL
	}
	_assetMetaDataByteSlice := []byte(assetMetaData)
	var _assetMetaType taprpc.AssetMetaType
	if AssetMetaTypeIsJsonNotOpaque {
		_assetMetaType = taprpc.AssetMetaType_META_TYPE_JSON
	} else {
		_assetMetaType = taprpc.AssetMetaType_META_TYPE_OPAQUE
	}
	_groupKeyByteSlices := []byte(groupKey)
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
			GroupKey:        _groupKeyByteSlices,
			GroupAnchor:     groupAnchor,
		},
		ShortResponse: shortResponse,
	}
	response, err := client.MintAsset(context.Background(), request)
	if err != nil {
		utils.LogError("mintrpc MintAsset Error", err)
		return nil, err
	}
	return response, nil
}

func finalizeBatch(shortResponse bool, feeRate int) (*mintrpc.FinalizeBatchResponse, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
	client := mintrpc.NewMintClient(conn)
	request := &mintrpc.FinalizeBatchRequest{
		ShortResponse: shortResponse,
		FeeRate:       uint32(feeRate),
	}
	response, err := client.FinalizeBatch(context.Background(), request)
	if err != nil {
		utils.LogError("mintrpc FinalizeBatch Error", err)
		return nil, err
	}
	return response, nil
}

func fetchAssetMeta(isHash bool, data string) (*taprpc.AssetMeta, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
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

func newAddr(assetId string, amt int) (*taprpc.Addr, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
	client := taprpc.NewTaprootAssetsClient(conn)
	_assetIdByteSlice, _ := hex.DecodeString(assetId)
	request := &taprpc.NewAddrRequest{
		AssetId: _assetIdByteSlice,
		Amt:     uint64(amt),
	}
	response, err := client.NewAddr(context.Background(), request)
	if err != nil {
		utils.LogError("taprpc NewAddr Error", err)
		return nil, err
	}
	return response, nil
}

func sendAsset(tapAddrs string, feeRate int) (*taprpc.SendAssetResponse, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
	client := taprpc.NewTaprootAssetsClient(conn)
	addrs := strings.Split(tapAddrs, ",")
	request := &taprpc.SendAssetRequest{
		TapAddrs: addrs,
		FeeRate:  uint32(feeRate),
	}
	response, err := client.SendAsset(context.Background(), request)
	if err != nil {
		utils.LogError("taprpc SendAsset Error", err)
		return nil, err
	}
	return response, nil
}

func sendAssetAddrSlice(addrSlice []string, feeRate int) (*taprpc.SendAssetResponse, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.SendAssetRequest{
		TapAddrs: addrSlice,
		FeeRate:  uint32(feeRate),
	}
	response, err := client.SendAsset(context.Background(), request)
	if err != nil {
		utils.LogError("taprpc SendAsset Error", err)
		return nil, err
	}
	return response, nil
}

func decodeAddr(addr string) (*taprpc.Addr, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.DecodeAddrRequest{
		Addr: addr,
	}
	response, err := client.DecodeAddr(context.Background(), request)
	if err != nil {
		utils.LogError("taprpc DecodeAddr Error", err)
		return nil, err
	}
	return response, nil
}

func listAssets(withWitness, includeSpent, includeLeased bool) (*taprpc.ListAssetResponse, error) {
	grpcHost := config.GetLoadConfig().ApiConfig.Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().ApiConfig.Tapd.Port)
	tlsCertPath := config.GetLoadConfig().ApiConfig.Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().ApiConfig.Tapd.MacaroonPath
	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListAssetRequest{
		WithWitness:   withWitness,
		IncludeSpent:  includeSpent,
		IncludeLeased: includeLeased,
	}
	response, err := client.ListAssets(context.Background(), request)
	return response, err
}
