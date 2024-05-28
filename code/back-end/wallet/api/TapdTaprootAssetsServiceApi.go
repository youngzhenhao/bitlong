package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/wallet/api/connect"
	"github.com/wallet/api/rpcclient"
	"strconv"
)

func AddrReceives() string {
	response, err := addrReceives()
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

func addrReceives() (*taprpc.AddrReceivesResponse, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.AddrReceivesRequest{}
	response, err := client.AddrReceives(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc DebugLevel Error: %v\n", GetTimeNow(), err)
		return nil, err
	}
	return response, nil
}

func BurnAsset() {

}

func DebugLevel() {

}

func DecodeAddr(addr string) string {
	response, err := rpcclient.DecodeAddr(addr)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	// make result struct
	var result = struct {
		Encoded          string `json:"encoded"`
		AssetId          string `json:"asset_id"`
		AssetType        int    `json:"asset_type"`
		Amount           int    `json:"amount"`
		GroupKey         string `json:"group_key"`
		ScriptKey        string `json:"script_key"`
		InternalKey      string `json:"internal_key"`
		TapscriptSibling string `json:"tapscript_sibling"`
		TaprootOutputKey string `json:"taproot_output_key"`
		ProofCourierAddr string `json:"proof_courier_addr"`
		AssetVersion     int    `json:"asset_version"`
	}{
		Encoded:          response.Encoded,
		AssetId:          hex.EncodeToString(response.AssetId),
		AssetType:        int(response.AssetType),
		Amount:           int(response.Amount),
		GroupKey:         hex.EncodeToString(response.GroupKey),
		ScriptKey:        hex.EncodeToString(response.ScriptKey),
		InternalKey:      hex.EncodeToString(response.InternalKey),
		TapscriptSibling: hex.EncodeToString(response.TapscriptSibling),
		TaprootOutputKey: hex.EncodeToString(response.TaprootOutputKey),
		ProofCourierAddr: response.ProofCourierAddr,
		AssetVersion:     int(response.AssetVersion),
	}
	return MakeJsonResult(true, "", result)
}

func DecodeProof(rawProof string) {

}

func ExportProof() {

}

func FetchAssetMeta(isHash bool, data string) string {
	response, err := fetchAssetMeta(isHash, data)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", string(response.Data))
}

// GetInfoOfTap
//
//	@Description: GetInfo returns the information for the node.
//	@return string
func GetInfoOfTap() string {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.GetInfoRequest{}
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc GetInfo Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// ListAssets
//
//	@Description: ListAssets lists the set of assets owned by the target daemon.
//	@return string
func ListAssets(withWitness, includeSpent, includeLeased bool) string {
	response, err := listAssets(withWitness, includeSpent, includeLeased)
	if err != nil {
		fmt.Printf("%s taprpc ListAssets Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

func ListSimpleAssets(withWitness, includeSpent, includeLeased bool) string {
	response, err := listAssets(withWitness, includeSpent, includeLeased)
	if err != nil {
		fmt.Printf("%s taprpc ListAssets Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	var (
		simpleAssets []struct {
			AssetId string `json:"asset_id"`
			Name    string `json:"name"`
			Amount  uint64 `json:"amount"`
			Type    string `json:"type"`
		}
	)
	for _, asset := range response.Assets {
		set := true
		for index, s := range simpleAssets {
			if s.AssetId == hex.EncodeToString(asset.AssetGenesis.GetAssetId()) {
				simpleAssets[index].Amount = asset.Amount + s.Amount
				set = false
				break
			}
		}
		if !set {
			continue
		}
		simpleAssets = append(simpleAssets, struct {
			AssetId string `json:"asset_id"`
			Name    string `json:"name"`
			Amount  uint64 `json:"amount"`
			Type    string `json:"type"`
		}{
			AssetId: hex.EncodeToString(asset.AssetGenesis.GetAssetId()),
			Name:    asset.AssetGenesis.Name,
			Amount:  asset.Amount,
			Type:    asset.AssetGenesis.AssetType.String(),
		})
	}

	return MakeJsonResult(true, "", simpleAssets)
}

func FindAssetByAssetName(assetName string) string {
	var response = struct {
		Success bool                     `json:"success"`
		Error   string                   `json:"error"`
		Data    taprpc.ListAssetResponse `json:"data"`
	}{}
	list := ListAssets(false, false, false)
	err := json.Unmarshal([]byte(list), &response)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	if response.Success == false {
		return MakeJsonResult(false, response.Error, nil)
	}
	var assets []*taprpc.Asset
	for _, asset := range response.Data.Assets {
		//if hex.EncodeToString(asset.AssetGenesis.GetAssetId()) == assetName {
		if asset.AssetGenesis.Name == assetName {
			assets = append(assets, asset)
		}
	}
	if len(assets) == 0 {
		return MakeJsonResult(false, "NOT_FOUND", nil)
	}
	return MakeJsonResult(true, "", assets)
}

// ListGroups
//
//	@Description: ListGroups lists the asset groups known to the target daemon, and the assets held in each group.
//	@return string
func ListGroups() string {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListGroupsRequest{}
	response, err := client.ListGroups(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc ListGroups Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// ListTransfers
//
//	@Description: ListTransfers lists outbound asset transfers tracked by the target daemon.
//	@return string
func ListTransfers() string {
	response, err := listTransfers()
	if err != nil {
		fmt.Printf("%s taprpc ListTransfers Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

func listTransfers() (*taprpc.ListTransfersResponse, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListTransfersRequest{}
	response, err := client.ListTransfers(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc ListTransfers Error: %v\n", GetTimeNow(), err)
		return nil, err
	}
	return response, err
}

// ListUtxos
//
//	@Description: ListUtxos lists the UTXOs managed by the target daemon, and the assets they hold.
//	@return string
func ListUtxos(includeLeased bool) string {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListUtxosRequest{
		IncludeLeased: includeLeased,
	}
	response, err := client.ListUtxos(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc ListUtxos Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// NewAddr
//
//	@Description:NewAddr makes a new address from the set of request params.
//	@return string
func NewAddr(assetId string, amt int) string {
	response, err := newAddr(assetId, amt)
	if err != nil {
		return MakeJsonResult(false, err.Error(), "")
	}
	return MakeJsonResult(true, "", response)
}

func newAddr(assetId string, amt int) (*taprpc.Addr, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	_assetIdByteSlice, _ := hex.DecodeString(assetId)
	request := &taprpc.NewAddrRequest{
		AssetId: _assetIdByteSlice,
		Amt:     uint64(amt),
	}
	response, err := client.NewAddr(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc NewAddr Error: %v\n", GetTimeNow(), err)
		return nil, err
	}
	return response, nil
}

func QueryAddrs() {

}

// jsonAddrs : ["addrs1","addrs2",...]
func SendAssets(jsonAddrs string, feeRate int64) string {
	var addrs []string
	err := json.Unmarshal([]byte(jsonAddrs), &addrs)
	if err != nil {
		fmt.Printf("%s json.Unmarshal Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "Please use the correct json format", "")
	}
	response, err := sendAssets(addrs, uint32(feeRate))
	if err != nil {
		return MakeJsonResult(false, err.Error(), "")
	}
	return MakeJsonResult(true, "", response)
}

// SendAsset
// @Description:SendAsset uses one or multiple passed Taproot Asset address(es) to attempt to complete an asset send.
// The method returns information w.r.t the on chain send, as well as the proof file information the receiver needs to fully receive the asset.
// @return string
// skipped function SendAsset with unsupported parameter or return types
func sendAssets(addrs []string, feeRate uint32) (*taprpc.SendAssetResponse, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)

	request := &taprpc.SendAssetRequest{
		TapAddrs: addrs,
	}
	if feeRate > 0 {
		request.FeeRate = feeRate
	}
	response, err := client.SendAsset(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc SendAsset Error: %v\n", GetTimeNow(), err)
		return nil, err
	}
	return response, nil
}

func SubscribeReceiveAssetEventNtfns() {

}

func SubscribeSendAssetEventNtfns() {

}

func VerifyProof() {

}

// TapStopDaemon
//
//	@Description: StopDaemon will send a shutdown request to the interrupt handler, triggering a graceful shutdown of the daemon.
//	@return bool
func TapStopDaemon() bool {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.StopRequest{}
	_, err = client.StopDaemon(context.Background(), request)
	if err != nil {
		fmt.Printf("%s taprpc TapStopDaemon Error: %v\n", GetTimeNow(), err)
		return false
	}
	return true
}

func decodeProof(proof []byte, depth uint32, withMetaReveal bool, withPrevWitnesses bool) (*taprpc.DecodeProofResponse, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.DecodeProofRequest{
		RawProof:          proof,
		ProofAtDepth:      depth,
		WithMetaReveal:    withMetaReveal,
		WithPrevWitnesses: withPrevWitnesses,
	}
	response, err := client.DecodeProof(context.Background(), request)
	return response, err
}

func fetchAssetMeta(isHash bool, data string) (*taprpc.AssetMeta, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()

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

func listBalances(useGroupKey bool, assetFilter, groupKeyFilter []byte) (*taprpc.ListBalancesResponse, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListBalancesRequest{
		AssetFilter:    assetFilter,
		GroupKeyFilter: groupKeyFilter,
	}
	if useGroupKey {
		request.GroupBy = &taprpc.ListBalancesRequest_GroupKey{GroupKey: true}
	} else {
		request.GroupBy = &taprpc.ListBalancesRequest_AssetId{AssetId: true}
	}
	response, err := client.ListBalances(context.Background(), request)
	return response, err
}

type ListAssetBalanceInfo struct {
	GenesisPoint string `json:"genesis_point"`
	Name         string `json:"name"`
	MetaHash     string `json:"meta_hash"`
	AssetID      string `json:"asset_id"`
	AssetType    string `json:"asset_type"`
	OutputIndex  int    `json:"output_index"`
	Version      int    `json:"version"`
	Balance      string `json:"balance"`
}

func ProcessListBalancesResponse(response *taprpc.ListBalancesResponse) *[]ListAssetBalanceInfo {
	var listAssetBalanceInfos []ListAssetBalanceInfo
	for _, balance := range response.AssetBalances {
		listAssetBalanceInfos = append(listAssetBalanceInfos, ListAssetBalanceInfo{
			GenesisPoint: balance.AssetGenesis.GenesisPoint,
			Name:         balance.AssetGenesis.Name,
			MetaHash:     hex.EncodeToString(balance.AssetGenesis.MetaHash),
			AssetID:      hex.EncodeToString(balance.AssetGenesis.AssetId),
			AssetType:    balance.AssetGenesis.AssetType.String(),
			OutputIndex:  int(balance.AssetGenesis.OutputIndex),
			Version:      int(balance.AssetGenesis.Version),
			Balance:      strconv.FormatUint(balance.Balance, 10),
		})
	}
	return &listAssetBalanceInfos
}

func ListBalances() string {
	response, err := listBalances(false, nil, nil)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", ProcessListBalancesResponse(response))
}

func listAssets(withWitness, includeSpent, includeLeased bool) (*taprpc.ListAssetResponse, error) {
	conn, clearUp, err := connect.GetConnection("tapd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer clearUp()
	client := taprpc.NewTaprootAssetsClient(conn)
	request := &taprpc.ListAssetRequest{
		WithWitness:   withWitness,
		IncludeSpent:  includeSpent,
		IncludeLeased: includeLeased,
	}
	response, err := client.ListAssets(context.Background(), request)
	return response, err
}
