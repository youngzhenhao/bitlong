package api

type Account struct {
	Name              string `json:"name"`
	Type              string `json:"address_type"`
	ExtendedPublicKey string `json:"extended_public_key"`
	DerivationPath    string `json:"derivation_path"`
}

func GetAllAccounts() []Account {
	var accs []Account
	response, err := listAccounts()
	if err != nil {
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
	if addressType == "NESTED_PUBKEY_HASH" {
		addressType = "HYBRID_NESTED_WITNESS_PUBKEY_HASH"
	}
	for _, acc := range accs {
		if acc.Type == addressType {
			return acc.DerivationPath
		}
	}
	return ""
}
