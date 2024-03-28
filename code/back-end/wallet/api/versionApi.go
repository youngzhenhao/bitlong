package api

import "github.com/wallet/base"

func GetV() string {
	return base.Configure("tapd")
}
