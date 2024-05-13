package api

import (
	"context"
	"encoding/hex"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"google.golang.org/grpc"
	"strconv"
	"trade/config"
	"trade/models"
	"trade/utils"
)

func assetLeaves(isGroup bool, id string, proofType universerpc.ProofType) (*universerpc.AssetLeafResponse, error) {
	grpcHost := config.GetLoadConfig().Tapd.Host + ":" + strconv.Itoa(config.GetLoadConfig().Tapd.Port)
	tlsCertPath := config.GetLoadConfig().Tapd.TlsCertPath
	macaroonPath := config.GetLoadConfig().Tapd.MacaroonPath
	creds := utils.NewTlsCert(tlsCertPath)
	macaroon := utils.GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(utils.NewMacaroonCredential(macaroon)))
	if err != nil {
		utils.LogError("did not connect: grpc.Dial", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		utils.LogError("conn Close Error", err)
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
