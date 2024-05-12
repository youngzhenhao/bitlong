package api

import (
	"AssetsTrade/models"
)

func GetAssetInfo(id string) *models.AssetIssuanceLeaf {
	return assetLeafIssuanceInfo(id)
}
