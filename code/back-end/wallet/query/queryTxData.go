package query

import (
	"github.com/wallet/config"
	"github.com/wallet/other"
	"strconv"
)

func TxListReqData(data *other.TxListData, address string) *other.TransactionList {
	res := &other.TransactionList{}
	ret := make([]*other.TransactionBase, 0)
	for _, value := range data.Data {
		regData := other.TransactionBase{}
		regData.Chain = "bitcoin"
		regData.TxHash = value.Id
		regData.Time = strconv.Itoa(value.Date * 1000)
		regData.Confirm = strconv.Itoa(value.Confirmations)
		if value.Status == "completed" {
			regData.Status = config.TransactionStatusOk
		} else {
			regData.Status = config.TransactionStatusFaild
		}
		regData.TokenName = "bitcoin"
		regData.Height = value.BlockNumber.String()
		other.ParseEvent(value.Events, &regData, address, "bitcoin")
		ret = append(ret, &regData)
	}
	res.PageToken = data.Meta.Paging.NextPageToken
	res.List = ret
	return res
}
