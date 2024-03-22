package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc/assetwalletrpc"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
)

// AnchorVirtualPsbts
//
//	@Description: AnchorVirtualPsbts merges and then commits multiple virtual transactions in a single BTC level anchor transaction.
//
// This RPC should be used if the BTC level anchor transaction of the assets to be spent are encumbered by a normal key and don't require any special spending conditions.
// For any custom spending conditions on the BTC level, the two RPCs CommitVirtualPsbts and PublishAndLogTransfer should be used instead (which in combination do the same as this RPC but allow for more flexibility).
// @return bool
//
// skipped function AnchorVirtualPsbts with unsupported parameter or return types
func AnchorVirtualPsbts(virtualPsbts []string) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_virtualPsbts := make([][]byte, 0)
	for _, i := range virtualPsbts {
		str, _ := hex.DecodeString(i)
		_virtualPsbts = append(_virtualPsbts, str)
	}
	request := &assetwalletrpc.AnchorVirtualPsbtsRequest{
		VirtualPsbts: _virtualPsbts,
	}
	response, err := client.AnchorVirtualPsbts(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc AnchorVirtualPsbts Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// FundVirtualPsbt
//
//	@Description:FundVirtualPsbt selects inputs from the available asset commitments to fund a virtual transaction matching the template.
//	@return bool
//
// skipped function FundVirtualPsbt with unsupported parameter or return types
func FundVirtualPsbt(isPsbtNotRaw bool, psbt ...string) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	request := &assetwalletrpc.FundVirtualPsbtRequest{}
	if isPsbtNotRaw {
		_psbtByteSlice, _ := hex.DecodeString(psbt[0])
		request.Template = &assetwalletrpc.FundVirtualPsbtRequest_Psbt{Psbt: _psbtByteSlice}
	} else {
		request.Template = &assetwalletrpc.FundVirtualPsbtRequest_Raw{
			Raw: &assetwalletrpc.TxTemplate{
				Recipients: nil,
			}}
	}

	response, err := client.FundVirtualPsbt(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc FundVirtualPsbt Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// NextInternalKey
//
//	@Description:NextInternalKey derives the next internal key for the given key family and stores it as an internal key in the database to make sure it is identified as a local key later on when importing proofs.
//	While an internal key can also be used as the internal key of a script key, it is recommended to use the NextScriptKey RPC instead, to make sure the tweaked Taproot output key is also recognized as a local key.
//	@return string
func NextInternalKey(keyFamily int) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	request := &assetwalletrpc.NextInternalKeyRequest{
		KeyFamily: uint32(keyFamily),
	}
	response, err := client.NextInternalKey(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc NextInternalKey Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// NextScriptKey
//
//	@Description:NextInternalKey derives the next internal key for the given key family and stores it as an internal key in the database to make sure it is identified as a local key later on when importing proofs.
//	While an internal key can also be used as the internal key of a script key, it is recommended to use the NextScriptKey RPC instead, to make sure the tweaked Taproot output key is also recognized as a local key.
//	@return string
func NextScriptKey(keyFamily int) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	request := &assetwalletrpc.NextScriptKeyRequest{
		KeyFamily: uint32(keyFamily),
	}
	response, err := client.NextScriptKey(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc NextScriptKey Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// ProveAssetOwnership
//
//	@Description:ProveAssetOwnership creates an ownership proof embedded in an asset transition proof.
//	That ownership proof is a signed virtual transaction spending the asset with a valid witness to prove the prover owns the keys that can spend the asset.
//	@return bool
func ProveAssetOwnership(assetId, scriptKey string) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_assetIdByteSlice, _ := hex.DecodeString(assetId)
	_scriptKeyByteSlice, _ := hex.DecodeString(scriptKey)
	request := &assetwalletrpc.ProveAssetOwnershipRequest{
		AssetId:   _assetIdByteSlice,
		ScriptKey: _scriptKeyByteSlice,
		Outpoint:  nil,
	}
	response, err := client.ProveAssetOwnership(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc ProveAssetOwnership Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// QueryInternalKey
//
//	@Description:QueryInternalKey returns the key descriptor for the given internal key.
//	@param internalKey
//	@return string
func QueryInternalKey(internalKey string) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_internalKeyByteSlice, _ := hex.DecodeString(internalKey)

	request := &assetwalletrpc.QueryInternalKeyRequest{
		InternalKey: _internalKeyByteSlice,
	}
	response, err := client.QueryInternalKey(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc QueryInternalKey Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// QueryScriptKey
//
//	@Description:QueryScriptKey returns the full script key descriptor for the given tweaked script key.
//	@return string
func QueryScriptKey(tweakedScriptKey string) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_tweakedScriptKeyByteSlice, _ := hex.DecodeString(tweakedScriptKey)
	request := &assetwalletrpc.QueryScriptKeyRequest{
		TweakedScriptKey: _tweakedScriptKeyByteSlice,
	}
	response, err := client.QueryScriptKey(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc QueryScriptKey Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// RemoveUTXOLease
//
//	@Description:RemoveUTXOLease removes the lease/lock/reservation of the given managed UTXO.
//	@return bool
func RemoveUTXOLease() bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	request := &assetwalletrpc.RemoveUTXOLeaseRequest{
		Outpoint: nil,
	}
	response, err := client.RemoveUTXOLease(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc RemoveUTXOLease Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// SignVirtualPsbt
//
//	@Description:SignVirtualPsbt signs the inputs of a virtual transaction and prepares the commitments of the inputs and outputs.
//	@return bool
func SignVirtualPsbt(fundedPsbt string) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_fundedPsbtByteSlice, _ := hex.DecodeString(fundedPsbt)
	request := &assetwalletrpc.SignVirtualPsbtRequest{
		FundedPsbt: _fundedPsbtByteSlice,
	}
	response, err := client.SignVirtualPsbt(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc SignVirtualPsbt Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// VerifyAssetOwnership
//
//	@Description:VerifyAssetOwnership verifies the asset ownership proof embedded in the given transition proof of an asset and returns true if the proof is valid.
//	@return bool
func VerifyAssetOwnership(proofWithWitness string) bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("tapd"), "tls.cert")
	newFilePath := filepath.Join(filepath.Join(base.Configure("tapd"), "data"), "testnet")
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
		fmt.Printf("%s Failed to append cert\n", GetTimeNow())
	}
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    certPool,
	}
	creds := credentials.NewTLS(config)

	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: grpc.Dial: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close Error: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := assetwalletrpc.NewAssetWalletClient(conn)
	_proofWithWitnessByteSlice, _ := hex.DecodeString(proofWithWitness)
	request := &assetwalletrpc.VerifyAssetOwnershipRequest{
		ProofWithWitness: _proofWithWitnessByteSlice,
	}
	response, err := client.VerifyAssetOwnership(context.Background(), request)
	if err != nil {
		fmt.Printf("%s assetwalletrpc VerifyAssetOwnership Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}
