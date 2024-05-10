package api

type FairLaunch struct {
	ID           string
	AssetID      string
	Name         string
	Amount       int
	Reserved     int
	MintQuantity int
	StartTime    int
	EndTime      int
	Minted       *[]MintedInfo
}

type MintedInfo struct {
	EncodedAddr      string
	AssetID          string
	AssetType        string
	Amount           int
	ScriptKey        string
	InternalKey      string
	TaprootOutputKey string
	ProofCourierAddr string
	AssetVersion     string
	MintTime         int
	Outpoint         string
	Address          string
}
