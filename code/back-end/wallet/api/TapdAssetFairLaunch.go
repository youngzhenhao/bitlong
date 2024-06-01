package api

import (
	"errors"
	"github.com/wallet/models"
)

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

func ProcessOwnSetFairLaunchResponseToIssuanceHistoryInfo(fairLaunchInfos *[]models.FairLaunchInfo) (*[]IssuanceHistoryInfo, error) {
	var err error
	var issuanceHistoryInfos []IssuanceHistoryInfo
	if fairLaunchInfos == nil || len(*fairLaunchInfos) == 0 {
		err = errors.New("fairLaunchInfos is null")
		LogError("", err)
		return nil, err
	}
	for _, fairLaunchInfo := range *fairLaunchInfos {
		issuanceHistoryInfos = append(issuanceHistoryInfos, IssuanceHistoryInfo{
			AssetName:    fairLaunchInfo.Name,
			AssetID:      fairLaunchInfo.AssetID,
			AssetType:    int(fairLaunchInfo.AssetType),
			IssuanceTime: fairLaunchInfo.IssuanceTime,
			State:        int(fairLaunchInfo.State),
		})
	}
	return &issuanceHistoryInfos, nil
}
