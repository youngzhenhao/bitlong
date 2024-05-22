package api

import (
	"trade/models"
	"trade/utils"
)

func GetAssetInfo(id string) *models.AssetIssuanceLeaf {
	return assetLeafIssuanceInfo(id)
}

func MintAsset(name string, assetTypeIsCollectible bool, assetMetaData *Meta, amount int, newGroupedAsset bool) string {
	Metastr := assetMetaData.ToJsonStr()
	return mintAsset(false, assetTypeIsCollectible, name, Metastr, false, amount, newGroupedAsset, false, "", "", false)
}

func FinalizeBatch(feeRate int) string {
	return finalizeBatch(false, feeRate)
}

func AddGroupAsset(name string, assetTypeIsCollectible bool, assetMetaData *Meta, amount int, groupKey string) string {
	Metastr := assetMetaData.ToJsonStr()
	return mintAsset(false, assetTypeIsCollectible, name, Metastr, false, amount, false, true, groupKey, "", false)
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
