package api

import (
	"fmt"
	"strings"
)

type Account struct {
	Name              string `json:"name"`
	Type              string `json:"address_type"`
	ExtendedPublicKey string `json:"extended_public_key"`
	DerivationPath    string `json:"derivation_path"`
}

func GetAllAccountsString() string {
	accs := GetAllAccounts()
	if accs == nil {
		return MakeJsonResult(false, "get all accounts fail.", "")
	}
	return MakeJsonResult(true, "", accs)
}

func GetAllAccounts() []Account {
	var accs []Account
	response, err := listAccounts()
	if err != nil {
		fmt.Printf("%s listAccounts fail. %v\n", GetTimeNow(), err)
		return nil
	}
	for _, v := range response.Accounts {
		accs = append(accs, Account{
			v.Name,
			v.AddressType.String(),
			v.ExtendedPublicKey,
			v.DerivationPath,
		})
	}
	return accs
}

func GetPathByAddressType(addressType string) string {
	accs := GetAllAccounts()
	addressType = strings.ToUpper(addressType)
	if addressType == "NESTED_PUBKEY_HASH" {
		addressType = "HYBRID_NESTED_WITNESS_PUBKEY_HASH"
	}
	for _, acc := range accs {
		if acc.Type == addressType {
			return MakeJsonResult(true, "", acc.DerivationPath)
		}
	}
	fmt.Printf("%s %v is not a valid address type.\n", GetTimeNow(), addressType)
	return MakeJsonResult(false, "can't find path by given address type.", "")
}
