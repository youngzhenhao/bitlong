package models

type FairLaunchInfo struct {
	ID           string        `json:"id"`
	AssetID      string        `json:"asset_id"`
	Name         string        `json:"name"`
	Amount       int           `json:"amount"`
	Reserved     int           `json:"reserved"`
	MintQuantity int           `json:"mint_quantity"`
	StartTime    int           `json:"start_time"`
	EndTime      int           `json:"end_time"`
	Minted       *[]MintedInfo `json:"minted"`
}

type MintedInfo struct {
	EncodedAddr      string `json:"encoded_addr"`
	AssetID          string `json:"asset_id"`
	AssetType        string `json:"asset_type"`
	Amount           int    `json:"amount"`
	ScriptKey        string `json:"script_key"`
	InternalKey      string `json:"internal_key"`
	TaprootOutputKey string `json:"taproot_output_key"`
	ProofCourierAddr string `json:"proof_courier_addr"`
	AssetVersion     string `json:"asset_version"`
	MintTime         int    `json:"mint_time"`
	Outpoint         string `json:"outpoint"`
	Address          string `json:"address"`
}
