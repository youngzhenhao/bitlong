package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
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
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)

	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf(GetTimeNow() + "Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
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
	return MakeJsonResult(true, "", response.Address)
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
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)

	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf(GetTimeNow() + "Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
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
	return MakeJsonResult(true, "", response.Address)
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
	macaroonBytes, err := os.ReadFile(macaroonPath)
	if err != nil {
		panic(err)
	}
	macaroon := hex.EncodeToString(macaroonBytes)

	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Printf(GetTimeNow() + "Failed to append cert")
	}

	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
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
	return MakeJsonResult(true, "", response.Address)
}

// StoreAddr
// @Description: Store Addr after being chosen.
// @param address
// @param balance
// @param _type
// @return string
func StoreAddr(address string, balance int, _type string) string {
	_ = InitAddrDB()
	path := filepath.Join(base.QueryConfigByKey("dirpath"), "addr.db")
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
		Address: address,
		Balance: balance,
		Type:    _type,
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
	path := filepath.Join(base.QueryConfigByKey("dirpath"), "addr.db")
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
	path := filepath.Join(base.QueryConfigByKey("dirpath"), "addr.db")
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
	path := filepath.Join(base.QueryConfigByKey("dirpath"), "addr.db")
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
		_addresses := accWithAddr.Addresses
		_addressType := accWithAddr.AddressType
		for _, _addr := range _addresses {
			if _addr.Balance != 0 {
				addrs = append(addrs, Addr{
					Address: _addr.Address,
					Balance: int(_addr.Balance),
					Type:    _addressType.String(),
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
		_addresses := accWithAddr.Addresses
		_addressType := accWithAddr.AddressType
		for _, _addr := range _addresses {
			if _addr.Balance != 0 {
				_ad := _addr.Address
				_ba := int(_addr.Balance)
				_ty := _addressType.String()
				var result JsonResult
				_re := StoreAddr(_ad, _ba, _ty)
				err := json.Unmarshal([]byte(_re), &result)
				if err != nil {
					return MakeJsonResult(false, "Store address Unmarshal fail. "+err.Error(), "")
				}
				if !result.Success {
					return MakeJsonResult(false, "Store address result false", "")
				}
				addresses = append(addresses, _ad)
			}
		}
	}
	return MakeJsonResult(true, "", addresses)
}
