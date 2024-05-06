package api

import (
	"encoding/hex"
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

func SimplifyAssetsList() *[]SimplifiedAssetsList {
	var simpleAssetsList []SimplifiedAssetsList
	response, _ := listAssets(false, false, false)
	for _, assets := range response.Assets {
		simpleAssetsList = append(simpleAssetsList, SimplifiedAssetsList{
			Version: assets.Version.String(),
			AssetGenesis: AssetsListAssetGenesis{
				GenesisPoint: assets.AssetGenesis.GenesisPoint,
				Name:         assets.AssetGenesis.Name,
				MetaHash:     hex.EncodeToString(assets.AssetGenesis.MetaHash),
				AssetID:      hex.EncodeToString(assets.AssetGenesis.AssetId),
				AssetType:    assets.AssetGenesis.AssetType.String(),
				OutputIndex:  int(assets.AssetGenesis.OutputIndex),
				Version:      int(assets.AssetGenesis.Version),
			},
			Amount:           strconv.FormatUint(assets.Amount, 10),
			LockTime:         int(assets.LockTime),
			ScriptKeyIsLocal: assets.ScriptKeyIsLocal,
			//RawGroupKey:      hex.EncodeToString(assets.AssetGroup.RawGroupKey),
			ChainAnchor: AssetsListChainAnchor{
				AnchorTx:         hex.EncodeToString(assets.ChainAnchor.AnchorTx),
				AnchorBlockHash:  assets.ChainAnchor.AnchorBlockHash,
				AnchorOutpoint:   assets.ChainAnchor.AnchorOutpoint,
				InternalKey:      hex.EncodeToString(assets.ChainAnchor.InternalKey),
				MerkleRoot:       hex.EncodeToString(assets.ChainAnchor.MerkleRoot),
				TapscriptSibling: hex.EncodeToString(assets.ChainAnchor.TapscriptSibling),
				BlockHeight:      int(assets.ChainAnchor.BlockHeight),
			},
			IsSpent:     assets.IsSpent,
			LeaseOwner:  hex.EncodeToString(assets.LeaseOwner),
			LeaseExpiry: strconv.FormatInt(assets.LeaseExpiry, 10),
			IsBurn:      assets.IsBurn,
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

type SimplifiedAssetsBalanceAssetBalances struct {
	AssetBalanceID string                    `json:"asset_balance_id"`
	AssetGenesis   AssetsBalanceAssetGenesis `json:"asset_genesis"`
	Balance        int                       `json:"balance"`
}

type SimplifiedAssetsBalanceAssetGroupBalances struct {
	AssetGroupBalanceID string                    `json:"asset_group_balance_id"`
	GroupBalance        AssetsBalanceGroupBalance `json:"group_balance"`
	GroupKey            string                    `json:"group_key"`
}

type SimplifiedAssetsBalance struct {
	AssetBalances      []SimplifiedAssetsBalanceAssetBalances      `json:"asset_balances"`
	AssetGroupBalances []SimplifiedAssetsBalanceAssetGroupBalances `json:"asset_group_balances"`
}

func SimplifyAssetsBalance() *SimplifiedAssetsBalance {
	var simplifiedBalance SimplifiedAssetsBalance
	response, _ := listBalances(false, nil, nil)
	for k1, v1 := range response.AssetBalances {
		simplifiedBalance.AssetBalances = append(simplifiedBalance.AssetBalances, SimplifiedAssetsBalanceAssetBalances{
			AssetBalanceID: k1,
			AssetGenesis: AssetsBalanceAssetGenesis{
				GenesisPoint: v1.AssetGenesis.GenesisPoint,
				Name:         v1.AssetGenesis.Name,
				MetaHash:     hex.EncodeToString(v1.AssetGenesis.MetaHash),
				AssetID:      hex.EncodeToString(v1.AssetGenesis.AssetId),
				AssetType:    v1.AssetGenesis.AssetType.String(),
				OutputIndex:  int(v1.AssetGenesis.OutputIndex),
				Version:      int(v1.AssetGenesis.Version),
			},
			Balance: int(v1.Balance),
		})
	}
	for k2, v2 := range response.AssetGroupBalances {
		simplifiedBalance.AssetGroupBalances = append(simplifiedBalance.AssetGroupBalances, SimplifiedAssetsBalanceAssetGroupBalances{
			AssetGroupBalanceID: k2,
			GroupBalance: AssetsBalanceGroupBalance{
				GroupKey: hex.EncodeToString(v2.GroupKey),
				Balance:  int(v2.Balance),
			},
		})
	}
	return &simplifiedBalance
}
