package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"path/filepath"
	"strings"
	"time"
)

// GetNewAddress_P2TR
// @dev: Get a p2tr address
// @note: TAPROOT_PUBKEY
// @Description:NewAddress creates a new address under control of the local wallet.
// @return string
func GetNewAddress_P2TR() string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.NewAddressRequest{
		Type: lnrpc.AddressType_TAPROOT_PUBKEY,
	}
	response, err := client.NewAddress(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc NewAddress err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "AddressType_TAPROOT_PUBKEY error", "")
	}
	return MakeJsonResult(true, "", Addr{
		Name:           "default",
		Address:        response.Address,
		Balance:        0,
		AddressType:    lnrpc.AddressType_TAPROOT_PUBKEY.String(),
		DerivationPath: AddressTypeToDerivationPath(lnrpc.AddressType_TAPROOT_PUBKEY.String()),
		IsInternal:     false,
	})
}

// GetNewAddress_P2WKH
// @dev: Get a p2wkh address
// @note: WITNESS_PUBKEY_HASH
// @Description:NewAddress creates a new address under control of the local wallet.
// @return string
func GetNewAddress_P2WKH() string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.NewAddressRequest{
		Type: lnrpc.AddressType_WITNESS_PUBKEY_HASH,
	}
	response, err := client.NewAddress(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc NewAddress err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "AddressType_WITNESS_PUBKEY_HASH error", "")
	}
	return MakeJsonResult(true, "", Addr{
		Name:           "default",
		Address:        response.Address,
		Balance:        0,
		AddressType:    lnrpc.AddressType_WITNESS_PUBKEY_HASH.String(),
		DerivationPath: AddressTypeToDerivationPath(lnrpc.AddressType_WITNESS_PUBKEY_HASH.String()),
		IsInternal:     false,
	})
}

// GetNewAddress_NP2WKH
// @dev: Get a np2wkh address
// @note: HYBRID_NESTED_WITNESS_PUBKEY_HASH
// @Description: NewAddress creates a new address under control of the local wallet.
// @return string
func GetNewAddress_NP2WKH() string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := NewTlsCert(tlsCertPath)
	macaroon := GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.NewAddressRequest{
		Type: lnrpc.AddressType_NESTED_PUBKEY_HASH,
	}
	response, err := client.NewAddress(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc NewAddress err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "AddressType_NESTED_PUBKEY_HASH error", "")
	}
	return MakeJsonResult(true, "", Addr{
		Name:           "default",
		Address:        response.Address,
		Balance:        0,
		AddressType:    lnrpc.AddressType_NESTED_PUBKEY_HASH.String(),
		DerivationPath: AddressTypeToDerivationPath(lnrpc.AddressType_NESTED_PUBKEY_HASH.String()),
		IsInternal:     false,
	})
}

// StoreAddr
// @Description: Store Addr after being chosen.
// @param address
// @param balance
// @param _type
// @return string
func StoreAddr(name string, address string, balance int, addressType string, derivationPath string, isInternal bool) string {
	_ = InitAddrDB()
	path := filepath.Join(base.QueryConfigByKey("dirpath"), "phone.db")
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &AddrStore{DB: db}
	err = s.CreateOrUpdateAddr("addresses", &Addr{
		Name:           name,
		Address:        address,
		Balance:        balance,
		AddressType:    addressType,
		DerivationPath: derivationPath,
		IsInternal:     isInternal,
	})
	if err != nil {
		return MakeJsonResult(false, "Store address fail", "")
	}
	return MakeJsonResult(true, "", address)
}

// RemoveAddr
// @Description: Remove a addr in all addresses
// @param address
// @return string
func RemoveAddr(address string) string {
	_ = InitAddrDB()
	path := filepath.Join(base.QueryConfigByKey("dirpath"), "phone.db")
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &AddrStore{DB: db}
	_, err = s.ReadAddr("addresses", address)
	if err != nil {
		return MakeJsonResult(false, "No such address available for deletion. Read Addr fail.", "")
	}
	err = s.DeleteAddr("addresses", address)
	if err != nil {
		return MakeJsonResult(false, "Delete Addr fail. "+err.Error(), "")
	}
	return MakeJsonResult(true, "", address)
}

// QueryAddr
// @Description: Query Addr in all addresses
// @param address
// @return string
func QueryAddr(address string) string {
	_ = InitAddrDB()
	path := filepath.Join(base.QueryConfigByKey("dirpath"), "phone.db")
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &AddrStore{DB: db}
	addr, err := s.ReadAddr("addresses", address)
	if err != nil {
		return MakeJsonResult(false, "No such address, read Addr fail.", "")
	}
	return MakeJsonResult(true, "", addr)
}

// QueryAllAddr
// @Description: get a json list of addresses
// @return string
func QueryAllAddr() string {
	_ = InitAddrDB()
	path := filepath.Join(base.QueryConfigByKey("dirpath"), "phone.db")
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("%s bolt.Open :%v\n", GetTimeNow(), err)
	}
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("%s db.Close :%v\n", GetTimeNow(), err)
		}
	}(db)
	s := &AddrStore{DB: db}
	addresses, err := s.AllAddresses("addresses")
	if err != nil || len(addresses) == 0 {
		return MakeJsonResult(false, "Addresses is NULL or read fail.", "")
	}
	return MakeJsonResult(true, "", addresses)
}

//	 QueryAddresses
//	 @Description:  Use listAddress to query the non-zero balance address, exported.
//					List of non-zero balance addresses constitutes the Total balance.
//	 @return string
func GetNonZeroBalanceAddresses() string {
	listAddrResp, err := listAddress()
	if err != nil {
		return MakeJsonResult(false, "Query addresses fail. "+err.Error(), "")
	}
	var addrs []Addr
	listAddrs := listAddrResp.GetAccountWithAddresses()
	if len(listAddrs) == 0 {
		return MakeJsonResult(false, "Queried non-zero balance addresses NULL.", "")
	}
	for _, accWithAddr := range listAddrs {
		addresses := accWithAddr.Addresses
		for _, address := range addresses {
			if address.Balance != 0 {
				addrs = append(addrs, Addr{
					Name:           accWithAddr.Name,
					Address:        address.Address,
					Balance:        int(address.Balance),
					AddressType:    accWithAddr.AddressType.String(),
					DerivationPath: accWithAddr.DerivationPath,
					IsInternal:     address.IsInternal,
				})
			}
		}
	}
	return MakeJsonResult(true, "", addrs)
}

// UpdateAllAddressesByGNZBA
// @Description: Update all addresses by query non zero balance addresses
// @return string
func UpdateAllAddressesByGNZBA() string {
	listAddrResp, err := listAddress()
	if err != nil {
		return MakeJsonResult(false, "Query addresses fail. "+err.Error(), "")
	}
	var addresses []string
	listAddrs := listAddrResp.GetAccountWithAddresses()
	if len(listAddrs) == 0 {
		return MakeJsonResult(false, "Queried non-zero balance addresses NULL.", "")
	}
	for _, accWithAddr := range listAddrs {
		if accWithAddr.Name != "default" {
			continue
		}
		_addresses := accWithAddr.Addresses
		for _, _address := range _addresses {
			if _address.Balance != 0 && !_address.IsInternal {
				var result JsonResult
				_re := StoreAddr(accWithAddr.Name, _address.Address, int(_address.Balance), accWithAddr.AddressType.String(), accWithAddr.DerivationPath, _address.IsInternal)
				err := json.Unmarshal([]byte(_re), &result)
				if err != nil {
					return MakeJsonResult(false, "Store address Unmarshal fail. "+err.Error(), "")
				}
				if !result.Success {
					return MakeJsonResult(false, "Store address result false", "")
				}
				addresses = append(addresses, _address.Address)
			}
		}
	}
	return MakeJsonResult(true, "", addresses)
}

// @dev: Acc

type Account struct {
	Name              string `json:"name"`
	AddressType       string `json:"address_type"`
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

// AddressTypeToDerivationPath
// @dev: NOT STANDARD RESULT RETURN
func AddressTypeToDerivationPath(addressType string) string {
	accs := GetAllAccounts()
	addressType = strings.ToUpper(addressType)
	if addressType == "NESTED_PUBKEY_HASH" {
		addressType = "HYBRID_NESTED_WITNESS_PUBKEY_HASH"
	}
	for _, acc := range accs {
		if acc.AddressType == addressType {
			return acc.DerivationPath
		}
	}
	fmt.Printf("%s %v is not a valid address type.\n", GetTimeNow(), addressType)
	return ""
}

func GetPathByAddressType(addressType string) string {
	accs := GetAllAccounts()
	if addressType == "NESTED_PUBKEY_HASH" {
		addressType = "HYBRID_NESTED_WITNESS_PUBKEY_HASH"
	}
	for _, acc := range accs {
		if acc.AddressType == addressType {
			return MakeJsonResult(true, "", acc.DerivationPath)
		}
	}
	fmt.Printf("%s %v is not a valid address type.\n", GetTimeNow(), addressType)
	return MakeJsonResult(false, "can't find path by given address type.", "")
}
