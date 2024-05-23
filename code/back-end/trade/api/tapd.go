package api

import (
	"encoding/hex"
	"errors"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/mintrpc"
	"strconv"
	"strings"
	"trade/models"
	"trade/utils"
)

func GetAssetInfo(id string) *models.AssetIssuanceLeaf {
	return assetLeafIssuanceInfo(id)
}

func MintAsset(name string, assetTypeIsCollectible bool, assetMetaData *Meta, amount int, newGroupedAsset bool) string {
	Metastr := assetMetaData.ToJsonStr()
	response, err := mintAsset(false, assetTypeIsCollectible, name, Metastr, false, amount, newGroupedAsset, false, "", "", false)
	if err != nil {
		return utils.MakeJsonResult(false, "mintAsset error. "+err.Error(), "")
	}
	return utils.MakeJsonResult(true, "", response)
}

func FinalizeBatch(feeRate int) string {
	response, err := finalizeBatch(false, feeRate)
	if err != nil {
		return utils.MakeJsonResult(false, err.Error(), nil)
	}
	return utils.MakeJsonResult(true, "", response)
}

func AddGroupAsset(name string, assetTypeIsCollectible bool, assetMetaData *Meta, amount int, groupKey string) string {
	Metastr := assetMetaData.ToJsonStr()
	response, err := mintAsset(false, assetTypeIsCollectible, name, Metastr, false, amount, false, true, groupKey, "", false)
	if err != nil {
		return utils.MakeJsonResult(false, "mintAsset error. "+err.Error(), "")
	}
	return utils.MakeJsonResult(true, "", response)
}

func NewAddr(assetId string, amt int) string {
	response, err := newAddr(assetId, amt)
	if err != nil {
		return utils.MakeJsonResult(false, err.Error(), "")
	}
	return utils.MakeJsonResult(true, "", response)
}

func SendAsset(assetId string, feeRate int) string {
	response, err := sendAsset(assetId, feeRate)
	if err != nil {
		return utils.MakeJsonResult(false, err.Error(), "")
	}
	return utils.MakeJsonResult(true, "", response)
}

func SendAssetBool(assetId string, feeRate int) (bool, error) {
	_, err := sendAsset(assetId, feeRate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DecodeAddr(addr string) string {
	response, err := decodeAddr(addr)
	if err != nil {
		return utils.MakeJsonResult(false, err.Error(), "")
	}
	return utils.MakeJsonResult(true, "", response)
}

func GetDecodedAddrInfo(addr string) (*taprpc.Addr, error) {
	return decodeAddr(addr)
}

func MintAssetAndGetResponse(name string, assetTypeIsCollectible bool, assetMetaData *Meta, amount int, newGroupedAsset bool) (*mintrpc.MintAssetResponse, error) {
	return mintAsset(false, assetTypeIsCollectible, name, assetMetaData.ToJsonStr(), false, amount, newGroupedAsset, false, "", "", false)
}

func FinalizeBatchAndGetResponse(feeRate int) (*mintrpc.FinalizeBatchResponse, error) {
	return finalizeBatch(false, feeRate)
}

func GetListAssetsResponse(withWitness bool, includeSpent bool, includeLeased bool) (*taprpc.ListAssetResponse, error) {
	return listAssets(withWitness, includeSpent, includeLeased)
}

func TransactionAndIndexToOutpoint(transaction string, index int) (outpoint string) {
	return transaction + strconv.Itoa(index)
}

func OutpointToTransactionAndIndex(outpoint string) (transaction string, index string) {
	result := strings.Split(outpoint, ":")
	return result[0], result[1]
}

func BatchTxidAnchorToAssetId(batchTxidAnchor string) (string, error) {
	assets, _ := listAssets(true, true, false)
	for _, asset := range assets.Assets {
		txid, _ := OutpointToTransactionAndIndex(asset.GetChainAnchor().GetAnchorOutpoint())
		if batchTxidAnchor == txid {
			return hex.EncodeToString(asset.GetAssetGenesis().AssetId), nil
		}
	}
	err := errors.New("no asset found for batch txid")
	utils.LogError("", err)
	return "", err
}

func QueryAssetType(assetType int) (string, error) {
	if assetType == 0 {
		return taprpc.AssetType_NORMAL.String(), nil
	} else if assetType == 1 {
		return taprpc.AssetType_COLLECTIBLE.String(), nil
	}
	return "", errors.New("not a valid asset type code")
}
