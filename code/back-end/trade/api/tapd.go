package api

import (
	"trade/models"
)

func GetAssetInfo(id string) *models.AssetIssuanceLeaf {
	return assetLeafIssuanceInfo(id)
}
