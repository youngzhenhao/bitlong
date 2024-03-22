package main

import "github.com/wallet/api"

func main() {

	//api.LndWorkflowUnlockWallet("12345678")

	//print(api.GetWalletBalance())
	//print(api.GetWalletBalance())
	//print(api.ListAssets(false, false, false))
	//print(api.TapGetInfo())
	//api.MintAsset(false, false, "cat667", "12312", false, 500, false, false, "", "", false)
	//api.FinalizeBatch(false, 1000)
	//api.CancelBatch()

	print(api.ListBatches())
}
