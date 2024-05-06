package api

import (
	"encoding/json"
	"fmt"
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
type SimplifiedAssetsTransfers struct {
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

// SimplifyAssetsTransfer
// @dev: need to complete
func SimplifyAssetsTransfer() string {
	var respBytes []byte
	if err := json.Unmarshal(TapMarshalRespBytes(ListTransfersEx()), &respBytes); err != nil {
		fmt.Printf("%s GATBM json.Unmarshal :%v\n", GetTimeNow(), err)
		return ""
	}
	return string(respBytes)
}
