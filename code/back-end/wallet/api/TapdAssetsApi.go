package api

import (
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/universerpc"
	"strconv"
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

// SyncAssetAllSlice @dev
func SyncAssetAllSlice(ids []string) {
	if len(ids) == 0 {
		return
	}
	for _, _id := range ids {
		fmt.Println(SyncAssetIssuance(_id))
		fmt.Println(SyncAssetTransfer(_id))
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

// TransferLeave @dev
type TransferLeave struct {
	Version      string `json:"version"`
	AssetGenesis struct {
		GenesisPoint string `json:"genesis_point"`
		Name         string `json:"name"`
		MetaHash     string `json:"meta_hash"`
		AssetID      string `json:"asset_id"`
		AssetType    string `json:"asset_type"`
		OutputIndex  int    `json:"output_index"`
		Version      int    `json:"version"`
	} `json:"asset_genesis"`
	Amount           string `json:"amount"`
	LockTime         int    `json:"lock_time"`
	RelativeLockTime int    `json:"relative_lock_time"`
	ScriptVersion    int    `json:"script_version"`
	ScriptKey        string `json:"script_key"`
	ScriptKeyIsLocal bool   `json:"script_key_is_local"`
	//AssetGroup       interface{} `json:"asset_group"`
	//ChainAnchor      interface{} `json:"chain_anchor"`
	//PrevWitnesses    []struct {
	//	PrevID struct {
	//		AnchorPoint string `json:"anchor_point"`
	//		AssetID     string `json:"asset_id"`
	//		ScriptKey   string `json:"script_key"`
	//		Amount      string `json:"amount"`
	//	} `json:"prev_id"`
	//	TxWitness       []interface{} `json:"tx_witness"`
	//	SplitCommitment struct {
	//		RootAsset struct {
	//			Version      string `json:"version"`
	//			AssetGenesis struct {
	//				GenesisPoint string `json:"genesis_point"`
	//				Name         string `json:"name"`
	//				MetaHash     string `json:"meta_hash"`
	//				AssetID      string `json:"asset_id"`
	//				AssetType    string `json:"asset_type"`
	//				OutputIndex  int    `json:"output_index"`
	//				Version      int    `json:"version"`
	//			} `json:"asset_genesis"`
	//			Amount           string      `json:"amount"`
	//			LockTime         int         `json:"lock_time"`
	//			RelativeLockTime int         `json:"relative_lock_time"`
	//			ScriptVersion    int         `json:"script_version"`
	//			ScriptKey        string      `json:"script_key"`
	//			ScriptKeyIsLocal bool        `json:"script_key_is_local"`
	//			AssetGroup       interface{} `json:"asset_group"`
	//			ChainAnchor      interface{} `json:"chain_anchor"`
	//			PrevWitnesses    []struct {
	//				PrevID struct {
	//					AnchorPoint string `json:"anchor_point"`
	//					AssetID     string `json:"asset_id"`
	//					ScriptKey   string `json:"script_key"`
	//					Amount      string `json:"amount"`
	//				} `json:"prev_id"`
	//				TxWitness       []string    `json:"tx_witness"`
	//				SplitCommitment interface{} `json:"split_commitment"`
	//			} `json:"prev_witnesses"`
	//			IsSpent     bool   `json:"is_spent"`
	//			LeaseOwner  string `json:"lease_owner"`
	//			LeaseExpiry string `json:"lease_expiry"`
	//			IsBurn      bool   `json:"is_burn"`
	//		} `json:"root_asset"`
	//	} `json:"split_commitment"`
	//} `json:"prev_witnesses"`
	IsSpent     bool   `json:"is_spent"`
	LeaseOwner  string `json:"lease_owner"`
	LeaseExpiry string `json:"lease_expiry"`
	IsBurn      bool   `json:"is_burn"`
}

// IssuanceLeave @dev
type IssuanceLeave struct {
	Asset struct {
		Version      string `json:"version"`
		AssetGenesis struct {
			GenesisPoint string `json:"genesis_point"`
			Name         string `json:"name"`
			MetaHash     string `json:"meta_hash"`
			AssetID      string `json:"asset_id"`
			AssetType    string `json:"asset_type"`
			OutputIndex  int    `json:"output_index"`
			Version      int    `json:"version"`
		} `json:"asset_genesis"`
		Amount           string      `json:"amount"`
		LockTime         int         `json:"lock_time"`
		RelativeLockTime int         `json:"relative_lock_time"`
		ScriptVersion    int         `json:"script_version"`
		ScriptKey        string      `json:"script_key"`
		ScriptKeyIsLocal bool        `json:"script_key_is_local"`
		AssetGroup       interface{} `json:"asset_group"`
		ChainAnchor      interface{} `json:"chain_anchor"`
		PrevWitnesses    []struct {
			PrevID struct {
				AnchorPoint string `json:"anchor_point"`
				AssetID     string `json:"asset_id"`
				ScriptKey   string `json:"script_key"`
				Amount      string `json:"amount"`
			} `json:"prev_id"`
			TxWitness       []interface{} `json:"tx_witness"`
			SplitCommitment interface{}   `json:"split_commitment"`
		} `json:"prev_witnesses"`
		IsSpent     bool   `json:"is_spent"`
		LeaseOwner  string `json:"lease_owner"`
		LeaseExpiry string `json:"lease_expiry"`
		IsBurn      bool   `json:"is_burn"`
	} `json:"asset"`
	Proof string `json:"proof"`
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

func DecodeRawProofByte(rawProof []byte) *taprpc.DecodeProofResponse {
	result, err := decodeProof(rawProof, 0, false, false)
	if err != nil {
		return nil
	}
	return result
}

// DecodeRawProof
// @dev:
func DecodeRawProof(proof string) {
	decodeString, err := hex.DecodeString(proof)
	if err != nil {
		return
	}
	DecodeRawProofByte(decodeString)
}

// AssetLeavesSpecified
// @dev: Need To Complete
func AssetLeavesSpecified(id string, proofType string) string {
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
		return MakeJsonResult(false, err.Error(), nil)
	}
	if response.Leaves == nil {
		return MakeJsonResult(false, "NOT_FOUND", nil)
	}
	return MakeJsonResult(true, "", response)
}
