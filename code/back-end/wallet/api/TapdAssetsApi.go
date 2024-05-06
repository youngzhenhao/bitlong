package api

import (
	"encoding/hex"
	"strconv"
)

type AssetsTransfers struct {
	Transfers []struct {
		TransferTimestamp  string `json:"transfer_timestamp"`
		AnchorTxHash       string `json:"anchor_tx_hash"`
		AnchorTxHeightHint int    `json:"anchor_tx_height_hint"`
		AnchorTxChainFees  string `json:"anchor_tx_chain_fees"`
		Inputs             []struct {
			AnchorPoint string `json:"anchor_point"`
			AssetID     string `json:"asset_id"`
			ScriptKey   string `json:"script_key"`
			Amount      string `json:"amount"`
		} `json:"inputs"`
		Outputs []struct {
			Anchor struct {
				Outpoint         string `json:"outpoint"`
				Value            string `json:"value"`
				InternalKey      string `json:"internal_key"`
				TaprootAssetRoot string `json:"taproot_asset_root"`
				MerkleRoot       string `json:"merkle_root"`
				TapscriptSibling string `json:"tapscript_sibling"`
				NumPassiveAssets int    `json:"num_passive_assets"`
			} `json:"anchor"`
			ScriptKey           string `json:"script_key"`
			ScriptKeyIsLocal    bool   `json:"script_key_is_local"`
			Amount              string `json:"amount"`
			NewProofBlob        string `json:"new_proof_blob"`
			SplitCommitRootHash string `json:"split_commit_root_hash"`
			OutputType          string `json:"output_type"`
			AssetVersion        string `json:"asset_version"`
		} `json:"outputs"`
	} `json:"transfers"`
}

// SimplifiedAssetsTransfers
// @dev: need to complete
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
}

type AssetsTransfersOutputAnchor struct {
	Outpoint         string `json:"outpoint"`
	Value            string `json:"value"`
	TaprootAssetRoot string `json:"taproot_asset_root"`
	MerkleRoot       string `json:"merkle_root"`
	TapscriptSibling string `json:"tapscript_sibling"`
	NumPassiveAssets int    `json:"num_passive_assets"`
}

type AssetsTransfersOutput struct {
	Anchor              AssetsTransfersOutputAnchor
	ScriptKeyIsLocal    bool   `json:"script_key_is_local"`
	Amount              string `json:"amount"`
	SplitCommitRootHash string `json:"split_commit_root_hash"`
	OutputType          string `json:"output_type"`
	AssetVersion        string `json:"asset_version"`
}

// SimplifyAssetsTransfer
// @dev: need to complete
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
			})
		}
		var outputs []AssetsTransfersOutput
		for _, _output := range transfers.Outputs {
			outputs = append(outputs, AssetsTransfersOutput{
				Anchor: AssetsTransfersOutputAnchor{
					Outpoint:         _output.Anchor.Outpoint,
					Value:            strconv.FormatInt(_output.Anchor.Value, 10),
					TaprootAssetRoot: hex.EncodeToString(_output.Anchor.TaprootAssetRoot),
					MerkleRoot:       hex.EncodeToString(_output.Anchor.MerkleRoot),
					TapscriptSibling: hex.EncodeToString(_output.Anchor.TapscriptSibling),
					NumPassiveAssets: int(_output.Anchor.NumPassiveAssets),
				},
				ScriptKeyIsLocal:    _output.ScriptKeyIsLocal,
				Amount:              strconv.FormatUint(_output.Amount, 10),
				SplitCommitRootHash: hex.EncodeToString(_output.SplitCommitRootHash),
				OutputType:          _output.OutputType.String(),
				AssetVersion:        _output.AssetVersion.String(),
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
