package api

type IssuanceHistoryInfo struct {
	AssetName    string `json:"asset_name"`
	AssetID      string `json:"asset_id"`
	AssetType    int    `json:"asset_type"`
	IssuanceTime int    `json:"issuance_time"`
	State        int    `json:"state"`
}

func GetIssuanceTransactionByteSize() int {
	// TODO: need to complete
	return 170
}

func GetMintTransactionByteSize() int {
	// TODO: need to complete
	return 170
}

// TODO: Assemble local and server asset issuance data
// @dev: Use new makeJsonResult

// http://127.0.0.1:8080/v1/fair_launch/query/own_set
