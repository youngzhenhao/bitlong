package api

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"strconv"
	"strings"
)

type SimplifiedAssetsTransfer struct {
	TransferTimestamp  string                  `json:"transfer_timestamp"`
	AnchorTxHash       string                  `json:"anchor_tx_hash"`
	AnchorTxHeightHint int                     `json:"anchor_tx_height_hint"`
	AnchorTxChainFees  string                  `json:"anchor_tx_chain_fees"`
	Inputs             []AssetsTransfersInput  `json:"inputs"`
	Outputs            []AssetsTransfersOutput `json:"outputs"`
}

type AssetsTransfersInput struct {
	AnchorPoint string `json:"anchor_point"`
	AssetID     string `json:"asset_id"`
	Amount      int    `json:"amount"`
	//ScriptKey   string `json:"script_key"`
}

type AssetsTransfersOutputAnchor struct {
	Outpoint string `json:"outpoint"`
	Value    string `json:"value"`
	//TaprootAssetRoot string `json:"taproot_asset_root"`
	//MerkleRoot       string `json:"merkle_root"`
	//TapscriptSibling string `json:"tapscript_sibling"`
	//NumPassiveAssets int    `json:"num_passive_assets"`
}

type AssetsTransfersOutput struct {
	Anchor           AssetsTransfersOutputAnchor
	ScriptKeyIsLocal bool   `json:"script_key_is_local"`
	Amount           string `json:"amount"`
	//SplitCommitRootHash string `json:"split_commit_root_hash"`
	OutputType   string `json:"output_type"`
	AssetVersion string `json:"asset_version"`
}

func SimplifyAssetsTransfer() *[]SimplifiedAssetsTransfer {
	var simpleTransfers []SimplifiedAssetsTransfer
	response, _ := listTransfers()
	for _, transfers := range response.Transfers {
		var inputs []AssetsTransfersInput
		for _, _input := range transfers.Inputs {
			inputs = append(inputs, AssetsTransfersInput{
				AnchorPoint: _input.AnchorPoint,
				AssetID:     hex.EncodeToString(_input.AssetId),
				Amount:      int(_input.Amount),
				//ScriptKey:   hex.EncodeToString(_input.ScriptKey),
			})
		}
		var outputs []AssetsTransfersOutput
		for _, _output := range transfers.Outputs {
			outputs = append(outputs, AssetsTransfersOutput{
				Anchor: AssetsTransfersOutputAnchor{
					Outpoint: _output.Anchor.Outpoint,
					Value:    strconv.FormatInt(_output.Anchor.Value, 10),
					//TaprootAssetRoot: hex.EncodeToString(_output.Anchor.TaprootAssetRoot),
					//MerkleRoot:       hex.EncodeToString(_output.Anchor.MerkleRoot),
					//TapscriptSibling: hex.EncodeToString(_output.Anchor.TapscriptSibling),
					//NumPassiveAssets: int(_output.Anchor.NumPassiveAssets),
				},
				ScriptKeyIsLocal: _output.ScriptKeyIsLocal,
				Amount:           strconv.FormatUint(_output.Amount, 10),
				//SplitCommitRootHash: hex.EncodeToString(_output.SplitCommitRootHash),
				OutputType:   _output.OutputType.String(),
				AssetVersion: _output.AssetVersion.String(),
			})
		}
		simpleTransfers = append(simpleTransfers, SimplifiedAssetsTransfer{
			TransferTimestamp:  strconv.FormatInt(transfers.TransferTimestamp, 10),
			AnchorTxHash:       hex.EncodeToString(transfers.AnchorTxHash),
			AnchorTxHeightHint: int(transfers.AnchorTxHeightHint),
			AnchorTxChainFees:  strconv.FormatInt(transfers.AnchorTxChainFees, 10),
			Inputs:             inputs,
			Outputs:            outputs,
		})
	}
	return &simpleTransfers
}

type SimplifiedAssetsList struct {
	Version      string                 `json:"version"`
	AssetGenesis AssetsListAssetGenesis `json:"asset_genesis"`
	Amount       string                 `json:"amount"`
	LockTime     int                    `json:"lock_time"`
	//RelativeLockTime int    `json:"relative_lock_time"`
	//ScriptVersion    int    `json:"script_version"`
	//ScriptKey        string `json:"script_key"`
	ScriptKeyIsLocal bool `json:"script_key_is_local"`
	//RawGroupKey      string `json:"raw_group_key"`
	//AssetGroup       struct {
	//	RawGroupKey     string `json:"raw_group_key"`
	//	TweakedGroupKey string `json:"tweaked_group_key"`
	//	AssetWitness    string `json:"asset_witness"`
	//} `json:"asset_group"`
	ChainAnchor AssetsListChainAnchor `json:"chain_anchor"`
	//PrevWitnesses []interface{} `json:"prev_witnesses"`
	IsSpent     bool   `json:"is_spent"`
	LeaseOwner  string `json:"lease_owner"`
	LeaseExpiry string `json:"lease_expiry"`
	IsBurn      bool   `json:"is_burn"`
}

type AssetsListAssetGenesis struct {
	GenesisPoint string `json:"genesis_point"`
	Name         string `json:"name"`
	MetaHash     string `json:"meta_hash"`
	AssetID      string `json:"asset_id"`
	AssetType    string `json:"asset_type"`
	OutputIndex  int    `json:"output_index"`
	Version      int    `json:"version"`
}

type AssetsListChainAnchor struct {
	AnchorTx         string `json:"anchor_tx"`
	AnchorBlockHash  string `json:"anchor_block_hash"`
	AnchorOutpoint   string `json:"anchor_outpoint"`
	InternalKey      string `json:"internal_key"`
	MerkleRoot       string `json:"merkle_root"`
	TapscriptSibling string `json:"tapscript_sibling"`
	BlockHeight      int    `json:"block_height"`
}

func SimplifyAssetsList(assets []*taprpc.Asset) *[]SimplifiedAssetsList {
	var simpleAssetsList []SimplifiedAssetsList
	for _, _asset := range assets {
		simpleAssetsList = append(simpleAssetsList, SimplifiedAssetsList{
			Version: _asset.Version.String(),
			AssetGenesis: AssetsListAssetGenesis{
				GenesisPoint: _asset.AssetGenesis.GenesisPoint,
				Name:         _asset.AssetGenesis.Name,
				MetaHash:     hex.EncodeToString(_asset.AssetGenesis.MetaHash),
				AssetID:      hex.EncodeToString(_asset.AssetGenesis.AssetId),
				AssetType:    _asset.AssetGenesis.AssetType.String(),
				OutputIndex:  int(_asset.AssetGenesis.OutputIndex),
				Version:      int(_asset.AssetGenesis.Version),
			},
			Amount:           strconv.FormatUint(_asset.Amount, 10),
			LockTime:         int(_asset.LockTime),
			ScriptKeyIsLocal: _asset.ScriptKeyIsLocal,
			//RawGroupKey:      hex.EncodeToString(_asset.AssetGroup.RawGroupKey),
			ChainAnchor: AssetsListChainAnchor{
				AnchorTx:         hex.EncodeToString(_asset.ChainAnchor.AnchorTx),
				AnchorBlockHash:  _asset.ChainAnchor.AnchorBlockHash,
				AnchorOutpoint:   _asset.ChainAnchor.AnchorOutpoint,
				InternalKey:      hex.EncodeToString(_asset.ChainAnchor.InternalKey),
				MerkleRoot:       hex.EncodeToString(_asset.ChainAnchor.MerkleRoot),
				TapscriptSibling: hex.EncodeToString(_asset.ChainAnchor.TapscriptSibling),
				BlockHeight:      int(_asset.ChainAnchor.BlockHeight),
			},
			IsSpent:     _asset.IsSpent,
			LeaseOwner:  hex.EncodeToString(_asset.LeaseOwner),
			LeaseExpiry: strconv.FormatInt(_asset.LeaseExpiry, 10),
			IsBurn:      _asset.IsBurn,
		})
	}
	return &simpleAssetsList
}

type AssetsBalanceAssetGenesis struct {
	GenesisPoint string `json:"genesis_point"`
	Name         string `json:"name"`
	MetaHash     string `json:"meta_hash"`
	AssetID      string `json:"asset_id"`
	AssetType    string `json:"asset_type"`
	OutputIndex  int    `json:"output_index"`
	Version      int    `json:"version"`
}

type AssetsBalanceGroupBalance struct {
	GroupKey string `json:"group_key"`
	Balance  int    `json:"balance"`
}

// SyncUniverseFullSpecified @dev
func SyncUniverseFullSpecified(universeHost string, id string, proofType string) string {
	if universeHost == "" {
		universeHost = "testnet.universe.lightning.finance:10029"
	}
	var _proofType universerpc.ProofType
	if proofType == "issuance" || proofType == "ISSUANCE" || proofType == "PROOF_TYPE_ISSUANCE" {
		_proofType = universerpc.ProofType_PROOF_TYPE_ISSUANCE
	} else if proofType == "transfer" || proofType == "TRANSFER" || proofType == "PROOF_TYPE_TRANSFER" {
		_proofType = universerpc.ProofType_PROOF_TYPE_TRANSFER
	} else {
		_proofType = universerpc.ProofType_PROOF_TYPE_UNSPECIFIED
	}
	var targets []*universerpc.SyncTarget
	universeID := &universerpc.ID{
		Id: &universerpc.ID_AssetIdStr{
			AssetIdStr: id,
		},
		ProofType: _proofType,
	}
	targets = append(targets, &universerpc.SyncTarget{
		Id: universeID,
	})
	response, err := syncUniverse(universeHost, targets, universerpc.UniverseSyncMode_SYNC_FULL)
	if err != nil {
		return MakeJsonResult(false, err.Error(), "")
	}
	return MakeJsonResult(true, "", response)
}

// SyncAssetIssuance @dev
func SyncAssetIssuance(id string) string {
	return SyncUniverseFullSpecified("", id, universerpc.ProofType_PROOF_TYPE_ISSUANCE.String())
}

// SyncAssetTransfer @dev
func SyncAssetTransfer(id string) string {
	return SyncUniverseFullSpecified("", id, universerpc.ProofType_PROOF_TYPE_TRANSFER.String())
}

// SyncAssetAll
// @dev: 1
func SyncAssetAll(id string) {
	fmt.Println(SyncAssetIssuance(id))
	fmt.Println(SyncAssetTransfer(id))
}

// SyncAssetAllSlice
// @dev
// @note: api.SyncAssetAllSlice(api.GetAllAssetIdByListAll())
func SyncAssetAllSlice(ids []string) {
	if len(ids) == 0 {
		return
	}
	for _, _id := range ids {
		fmt.Println("Sync issuance:", _id, ".", SyncAssetIssuance(_id))
		fmt.Println("Sync transfer:", _id, ".", SyncAssetTransfer(_id))
	}
}

// SyncAssetAllWithAssets @dev
func SyncAssetAllWithAssets(ids ...string) {
	if len(ids) == 0 {
		return
	}
	for _, _id := range ids {
		fmt.Println(SyncAssetIssuance(_id))
		fmt.Println(SyncAssetTransfer(_id))
	}
}

type AssetBalance struct {
	Name      string `json:"name"`
	MetaHash  string `json:"meta_hash"`
	AssetID   string `json:"asset_id"`
	AssetType string `json:"asset_type"`
	Balance   int    `json:"balance"`
}

type AssetGroupBalance struct {
	ID       string `json:"id"`
	Balance  int    `json:"balance"`
	GroupKey string `json:"group_key"`
}

func allAssetBalances() *[]AssetBalance {
	response, _ := listBalances(false, nil, nil)
	var assetBalances []AssetBalance
	for _, v := range response.AssetBalances {
		assetBalances = append(assetBalances, AssetBalance{
			Name:      v.AssetGenesis.Name,
			MetaHash:  hex.EncodeToString(v.AssetGenesis.MetaHash),
			AssetID:   hex.EncodeToString(v.AssetGenesis.AssetId),
			AssetType: v.AssetGenesis.AssetType.String(),
			Balance:   int(v.Balance),
		})
	}
	if len(assetBalances) == 0 {
		return nil
	}
	return &assetBalances
}

// GetAllAssetBalances
// @dev: 0
// @note: Get all balance of assets info
func GetAllAssetBalances() string {
	result := allAssetBalances()
	if result == nil {
		return MakeJsonResult(false, "Null Balances", nil)
	}
	return MakeJsonResult(true, "", result)
}

func allAssetGroupBalances() *[]AssetGroupBalance {
	response, _ := listBalances(false, nil, nil)
	var assetGroupBalances []AssetGroupBalance
	for k, v := range response.AssetGroupBalances {
		assetGroupBalances = append(assetGroupBalances, AssetGroupBalance{
			ID:       k,
			Balance:  int(v.Balance),
			GroupKey: hex.EncodeToString(v.GroupKey),
		})
	}
	if len(assetGroupBalances) == 0 {
		return nil
	}
	return &assetGroupBalances
}

func GetAllAssetGroupBalances() string {
	result := allAssetGroupBalances()
	if result == nil {
		return MakeJsonResult(false, "Null Group Balances", nil)
	}
	return MakeJsonResult(true, "", result)
}

func GetAllAssetId(assetBalance *[]AssetBalance) *[]string {
	if assetBalance == nil {
		return nil
	}
	var ids []string
	for _, v := range *assetBalance {
		ids = append(ids, v.AssetID)
	}
	return &ids
}

// SyncAllAssetsByAssetBalance
// @dev: 2
// @note: Sync all assets of non-zero-balance to public universe
func SyncAllAssetsByAssetBalance() string {
	ids := GetAllAssetId(allAssetBalances())
	if ids != nil {
		SyncAssetAllSlice(*ids)
	}
	return MakeJsonResult(true, "", ids)
}

// GetAllAssetsIdSlice
// @dev: 3
// @note: Get an array including all assets ids
func GetAllAssetsIdSlice() string {
	ids := GetAllAssetId(allAssetBalances())
	return MakeJsonResult(true, "", ids)
}

// assetKeysTransfer
// @dev
func assetKeysTransfer(id string) *[]AssetKey {
	var _proofType universerpc.ProofType
	_proofType = universerpc.ProofType_PROOF_TYPE_TRANSFER
	response, err := assetLeafKeys(id, _proofType)
	if err != nil {
		fmt.Printf("%s universerpc AssetLeafKeys Error: %v\n", GetTimeNow(), err)
		return nil
	}
	if len(response.AssetKeys) == 0 {
		return nil
	}
	return processAssetKey(response)
}

func AssetKeysTransfer(id string) string {
	result := assetKeysTransfer(id)
	if result == nil {
		return MakeJsonResult(false, "Null Asset Keys", nil)
	}
	return MakeJsonResult(true, "", result)
}

// AssetLeavesSpecified
// @dev: Need To Complete
func AssetLeavesSpecified(id string, proofType string) *universerpc.AssetLeafResponse {
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
		fmt.Printf("%s universerpc AssetLeaves Error: %v\n", GetTimeNow(), err)
		return nil
	}
	return response
}

type AssetTransferLeave struct {
	Name string `json:"name"`
	//MetaHash     string `json:"meta_hash"`
	AssetID   string `json:"asset_id"`
	Amount    int    `json:"amount"`
	ScriptKey string `json:"script_key"`
	//PrevWitnesses []struct {
	//	PrevID struct {
	//		AnchorPoint string `json:"anchor_point"`
	//		AssetID     string `json:"asset_id"`
	//		ScriptKey   string `json:"script_key"`
	//	} `json:"prev_id"`
	//	SplitCommitment struct {
	//		RootAsset struct {
	//			AssetGenesis struct {
	//				GenesisPoint string `json:"genesis_point"`
	//				Name         string `json:"name"`
	//				MetaHash     string `json:"meta_hash"`
	//				AssetID      string `json:"asset_id"`
	//			} `json:"asset_genesis"`
	//			Amount        int    `json:"amount"`
	//			ScriptKey     string `json:"script_key"`
	//			PrevWitnesses []struct {
	//				PrevID struct {
	//					AnchorPoint string `json:"anchor_point"`
	//					AssetID     string `json:"asset_id"`
	//					ScriptKey   string `json:"script_key"`
	//				} `json:"prev_id"`
	//				TxWitness []string `json:"tx_witness"`
	//			} `json:"prev_witnesses"`
	//		} `json:"root_asset"`
	//	} `json:"split_commitment"`
	//} `json:"prev_witnesses"`
	Proof string `json:"proof"`
}

func ProcessAssetTransferLeave(response *universerpc.AssetLeafResponse) *[]AssetTransferLeave {
	var assetTransferLeaves []AssetTransferLeave
	for _, leave := range response.Leaves {
		assetTransferLeaves = append(assetTransferLeaves, AssetTransferLeave{
			Name:      leave.Asset.AssetGenesis.Name,
			AssetID:   hex.EncodeToString(leave.Asset.AssetGenesis.AssetId),
			Amount:    int(leave.Asset.Amount),
			ScriptKey: hex.EncodeToString(leave.Asset.ScriptKey),
			Proof:     hex.EncodeToString(leave.Proof),
		})
	}
	return &assetTransferLeaves
}

func AssetLeavesTransfer(id string) string {
	response := AssetLeavesSpecified(id, universerpc.ProofType_PROOF_TYPE_TRANSFER.String())
	if response == nil {
		fmt.Printf("%s universerpc AssetLeaves Error.\n", GetTimeNow())
		return MakeJsonResult(false, errors.New("null asset leaves").Error(), nil)
	}
	assetTransferLeaves := ProcessAssetTransferLeave(response)
	return MakeJsonResult(true, "", assetTransferLeaves)
}

func AssetLeavesTransfer_ONLY_FOR_TEST(id string) *[]AssetTransferLeave {
	response := AssetLeavesSpecified(id, universerpc.ProofType_PROOF_TYPE_TRANSFER.String())
	if response == nil {
		fmt.Printf("%s universerpc AssetLeaves Error.\n", GetTimeNow())
		return nil
	}
	return ProcessAssetTransferLeave(response)
}

func DecodeRawProofByte(rawProof []byte) *taprpc.DecodeProofResponse {
	result, err := decodeProof(rawProof, 0, false, false)
	if err != nil {
		return nil
	}
	return result
}

// DecodeRawProofString
// @dev:
func DecodeRawProofString(proof string) *taprpc.DecodeProofResponse {
	decodeString, err := hex.DecodeString(proof)
	if err != nil {
		return nil
	}
	return DecodeRawProofByte(decodeString)
}

type DecodedProof struct {
	NumberOfProofs  int    `json:"number_of_proofs"`
	Name            string `json:"name"`
	AssetID         string `json:"asset_id"`
	Amount          int    `json:"amount"`
	ScriptKey       string `json:"script_key"`
	AnchorTx        string `json:"anchor_tx"`
	AnchorBlockHash string `json:"anchor_block_hash"`
	AnchorOutpoint  string `json:"anchor_outpoint"`
	InternalKey     string `json:"internal_key"`
	MerkleRoot      string `json:"merkle_root"`
	BlockHeight     int    `json:"block_height"`
}

func ProcessProof(response *taprpc.DecodeProofResponse) *DecodedProof {
	if response == nil {
		return nil
	}
	return &DecodedProof{
		NumberOfProofs:  int(response.DecodedProof.NumberOfProofs),
		Name:            response.DecodedProof.Asset.AssetGenesis.Name,
		AssetID:         hex.EncodeToString(response.DecodedProof.Asset.AssetGenesis.AssetId),
		Amount:          int(response.DecodedProof.Asset.Amount),
		ScriptKey:       hex.EncodeToString(response.DecodedProof.Asset.ScriptKey),
		AnchorTx:        hex.EncodeToString(response.DecodedProof.Asset.ChainAnchor.AnchorTx),
		AnchorBlockHash: response.DecodedProof.Asset.ChainAnchor.AnchorBlockHash,
		AnchorOutpoint:  response.DecodedProof.Asset.ChainAnchor.AnchorOutpoint,
		InternalKey:     hex.EncodeToString(response.DecodedProof.Asset.ChainAnchor.InternalKey),
		MerkleRoot:      hex.EncodeToString(response.DecodedProof.Asset.ChainAnchor.MerkleRoot),
		BlockHeight:     int(response.DecodedProof.Asset.ChainAnchor.BlockHeight),
	}
}

func DecodeRawProof(proof string) string {
	response := DecodeRawProofString(proof)
	if response == nil {
		return MakeJsonResult(false, "null raw proof", nil)
	}
	return MakeJsonResult(true, "", ProcessProof(response))
}

func allAssetList() *taprpc.ListAssetResponse {
	response, err := listAssets(false, true, false)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return response
}

type ListAllAsset struct {
	Version            string `json:"version"`
	GenesisPoint       string `json:"genesis_point"`
	GenesisName        string `json:"genesis_name"`
	GenesisMetaHash    string `json:"genesis_meta_hash"`
	GenesisAssetID     string `json:"genesis_asset_id"`
	GenesisAssetType   string `json:"genesis_asset_type"`
	GenesisOutputIndex int    `json:"genesis_output_index"`
	Amount             string `json:"amount"`
	LockTime           int    `json:"lock_time"`
	RelativeLockTime   int    `json:"relative_lock_time"`
	ScriptVersion      int    `json:"script_version"`
	ScriptKey          string `json:"script_key"`
	ScriptKeyIsLocal   bool   `json:"script_key_is_local"`
	AnchorTx           string `json:"anchor_tx"`
	AnchorBlockHash    string `json:"anchor_block_hash"`
	AnchorOutpoint     string `json:"anchor_outpoint"`
	AnchorInternalKey  string `json:"anchor_internal_key"`
	AnchorBlockHeight  int    `json:"anchor_block_height"`
	IsSpent            bool   `json:"is_spent"`
	LeaseOwner         string `json:"lease_owner"`
	LeaseExpiry        string `json:"lease_expiry"`
	IsBurn             bool   `json:"is_burn"`
}

func ProcessListAllAssets(response *taprpc.ListAssetResponse) *[]ListAllAsset {
	if response == nil || response.Assets == nil || len(response.Assets) == 0 {
		return nil
	}
	var listAllAssets []ListAllAsset
	for _, asset := range response.Assets {
		listAllAssets = append(listAllAssets, ListAllAsset{
			Version:            asset.Version.String(),
			GenesisPoint:       asset.AssetGenesis.GenesisPoint,
			GenesisName:        asset.AssetGenesis.Name,
			GenesisMetaHash:    hex.EncodeToString(asset.AssetGenesis.MetaHash),
			GenesisAssetID:     hex.EncodeToString(asset.AssetGenesis.AssetId),
			GenesisAssetType:   asset.AssetGenesis.AssetType.String(),
			GenesisOutputIndex: int(asset.AssetGenesis.OutputIndex),
			Amount:             strconv.FormatUint(asset.Amount, 10),
			LockTime:           int(asset.LockTime),
			RelativeLockTime:   int(asset.RelativeLockTime),
			ScriptVersion:      int(asset.ScriptVersion),
			ScriptKey:          hex.EncodeToString(asset.ScriptKey),
			ScriptKeyIsLocal:   asset.ScriptKeyIsLocal,
			AnchorTx:           hex.EncodeToString(asset.ChainAnchor.AnchorTx),
			AnchorBlockHash:    asset.ChainAnchor.AnchorBlockHash,
			AnchorOutpoint:     asset.ChainAnchor.AnchorOutpoint,
			AnchorInternalKey:  hex.EncodeToString(asset.ChainAnchor.InternalKey),
			AnchorBlockHeight:  int(asset.ChainAnchor.BlockHeight),
			IsSpent:            asset.IsSpent,
			LeaseOwner:         hex.EncodeToString(asset.LeaseOwner),
			LeaseExpiry:        strconv.FormatInt(asset.LeaseExpiry, 10),
			IsBurn:             asset.IsBurn,
		})
	}
	if len(listAllAssets) == 0 {
		return nil
	}
	return &listAllAssets
}

func GetAllAssetList() string {
	response := allAssetList()
	if response == nil {
		return MakeJsonResult(false, "null asset list", nil)
	}
	return MakeJsonResult(true, "", ProcessListAllAssets(response))
}

type ListAllAssetSimplified struct {
	GenesisName      string `json:"genesis_name"`
	GenesisAssetID   string `json:"genesis_asset_id"`
	GenesisAssetType string `json:"genesis_asset_type"`
	Amount           string `json:"amount"`
	AnchorOutpoint   string `json:"anchor_outpoint"`
	IsSpent          bool   `json:"is_spent"`
}

func ProcessListAllAssetsSimplified(result *[]ListAllAsset) *[]ListAllAssetSimplified {
	if result == nil || len(*result) == 0 {
		return nil
	}
	var listAllAssetsSimplified []ListAllAssetSimplified
	for _, asset := range *result {
		listAllAssetsSimplified = append(listAllAssetsSimplified, ListAllAssetSimplified{
			GenesisName:      asset.GenesisName,
			GenesisAssetID:   asset.GenesisAssetID,
			GenesisAssetType: asset.GenesisAssetType,
			Amount:           asset.Amount,
			AnchorOutpoint:   asset.AnchorOutpoint,
			IsSpent:          asset.IsSpent,
		})
	}
	if len(listAllAssetsSimplified) == 0 {
		return nil
	}
	return &listAllAssetsSimplified
}

// GetAllAssetListSimplified
// @dev:
func GetAllAssetListSimplified() string {
	result := ProcessListAllAssetsSimplified(ProcessListAllAssets(allAssetList()))
	if result == nil {
		return MakeJsonResult(false, "null asset list", nil)
	}
	return MakeJsonResult(true, "", result)
}

func GetAllAssetIdByListAll() []string {
	id := make(map[string]bool)
	var ids []string
	result := ProcessListAllAssetsSimplified(ProcessListAllAssets(allAssetList()))
	//var index int
	if result == nil || len(*result) == 0 {
		return nil
	}
	for _, asset := range *result {
		//index++
		//fmt.Println(index, asset.GenesisAssetID)
		id[asset.GenesisAssetID] = true
	}
	for k, _ := range id {
		ids = append(ids, k)
	}
	if len(ids) == 0 {
		return nil
	}
	//fmt.Println(len(ids))
	return ids
}

// SyncUniverseFullIssuanceByIdSlice
// @dev
// @note: Deprecated
func SyncUniverseFullIssuanceByIdSlice(ids []string) string {
	universeHost := "testnet.universe.lightning.finance:10029"
	var targets []*universerpc.SyncTarget
	for _, id := range ids {
		targets = append(targets, &universerpc.SyncTarget{
			Id: &universerpc.ID{
				Id: &universerpc.ID_AssetIdStr{
					AssetIdStr: id,
				},
				ProofType: universerpc.ProofType_PROOF_TYPE_ISSUANCE,
			},
		})
	}
	response, err := syncUniverse(universeHost, targets, universerpc.UniverseSyncMode_SYNC_FULL)
	if err != nil {
		return MakeJsonResult(false, err.Error(), "")
	}
	return MakeJsonResult(true, "", response)
}

// SyncUniverseFullTransferByIdSlice
// @dev
// @note: Deprecated
func SyncUniverseFullTransferByIdSlice(ids []string) string {
	universeHost := "testnet.universe.lightning.finance:10029"
	var targets []*universerpc.SyncTarget
	for _, id := range ids {
		targets = append(targets, &universerpc.SyncTarget{
			Id: &universerpc.ID{
				Id: &universerpc.ID_AssetIdStr{
					AssetIdStr: id,
				},
				ProofType: universerpc.ProofType_PROOF_TYPE_TRANSFER,
			},
		})
	}
	response, err := syncUniverse(universeHost, targets, universerpc.UniverseSyncMode_SYNC_FULL)
	if err != nil {
		return MakeJsonResult(false, err.Error(), "")
	}
	return MakeJsonResult(true, "", response)
}

// SyncUniverseFullNoSlice
// @dev
// @note: Sync all assets
func SyncUniverseFullNoSlice() string {
	universeHost := "testnet.universe.lightning.finance:10029"
	var targets []*universerpc.SyncTarget

	response, err := syncUniverse(universeHost, targets, universerpc.UniverseSyncMode_SYNC_FULL)
	if err != nil {
		return MakeJsonResult(false, err.Error(), "")
	}
	return MakeJsonResult(true, "", response)
}

type AssetHoldInfo struct {
	Name      string `json:"name"`
	AssetId   string `json:"assetId"`
	Amount    int    `json:"amount"`
	Outpoint  string `json:"outpoint"`
	Address   string `json:"address"`
	ScriptKey string `json:"scriptKey"`
	Proof     string `json:"proof"`
}

// OutpointToAddress
// TODO
// @dev:
func OutpointToAddress(outpoint string) string {
	transaction, index := GetTransactionAndIndexByOutpoint(outpoint)
	_ = transaction + index

	return ""
}

// GetTransactionAndIndexByOutpoint
// @dev: Split outpoint
func GetTransactionAndIndexByOutpoint(outpoint string) (transaction string, index string) {
	result := strings.Split(outpoint, ":")
	return result[0], result[1]
}

// CompareScriptKey
// @dev:
func CompareScriptKey(scriptKey1 string, scriptKey2 string) string {
	if scriptKey1 == scriptKey2 {
		return scriptKey1
	} else if len(scriptKey1) == len(scriptKey2) {
		return ""
	} else if len(scriptKey1) > len(scriptKey2) {
		if scriptKey1 == "0"+scriptKey2 || scriptKey1 == "02"+scriptKey2 {
			return scriptKey1
		} else if scriptKey1 == "2"+scriptKey2 {
			return "02" + scriptKey2
		} else {
			return ""
		}
	} else if len(scriptKey1) < len(scriptKey2) {
		if "0"+scriptKey1 == scriptKey2 || "02"+scriptKey1 == scriptKey2 {
			return scriptKey2
		} else if "2"+scriptKey1 == scriptKey2 {
			return "02" + scriptKey1
		} else {
			return ""
		}
	}
	return ""
}