package api

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/lightninglabs/taproot-assets/taprpc"
	"github.com/lightninglabs/taproot-assets/taprpc/mintrpc"
	"github.com/vincent-petithory/dataurl"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
	"strings"
)

// CancelBatch
//
//	@Description: CancelBatch will attempt to cancel the current pending batch.
//	@return bool
func CancelBatch() bool {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
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
	client := mintrpc.NewMintClient(conn)
	request := &mintrpc.CancelBatchRequest{}
	response, err := client.CancelBatch(context.Background(), request)
	if err != nil {
		fmt.Printf("%s mintrpc CancelBatch Error: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// finalizeBatch
//
//	@Description: FinalizeBatch will attempt to finalize the current pending batch.
//	@param shortResponse
//	@param feeRate
//	@return string
func finalizeBatch(shortResponse bool, feeRate int) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
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
	client := mintrpc.NewMintClient(conn)
	request := &mintrpc.FinalizeBatchRequest{
		ShortResponse: shortResponse,
		FeeRate:       uint32(feeRate),
	}
	response, err := client.FinalizeBatch(context.Background(), request)
	if err != nil {
		fmt.Printf("%s mintrpc FinalizeBatch Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

// FinalizeBatch
//
//	@Description: Wraps the finalizeBatch. FinalizeBatch will attempt to finalize the current pending batch.
//	@param shortResponse
//	@param feeRate
//	@return bool
func FinalizeBatch(feeRate int) string {
	return finalizeBatch(false, feeRate)
}

// ListBatches
//
//	@Description: ListBatches lists the set of batches submitted to the daemon, including pending and cancelled batches.
//	@return string
func ListBatches() string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
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
	client := mintrpc.NewMintClient(conn)
	request := &mintrpc.ListBatchRequest{}
	response, err := client.ListBatches(context.Background(), request)
	if err != nil {
		fmt.Printf("%s mintrpc ListBatches Error: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// mintAsset
//
//	@Description:MintAsset will attempt to mint the set of assets (async by default to ensure proper batching) specified in the request.
//	The pending batch is returned that shows the other pending assets that are part of the next batch. This call will block until the operation succeeds (asset is staged in the batch) or fails.
//	@param assetVersionIsV1
//	@param assetTypeIsCollectible
//	@param name
//	@param assetMetaData
//	@param AssetMetaTypeIsJsonNotOpaque
//	@param amount
//	@param newGroupedAsset
//	@param groupedAsset
//	@param groupKey
//	@param groupAnchor
//	@param shortResponse
//	@return bool
func mintAsset(assetVersionIsV1 bool, assetTypeIsCollectible bool, name string, assetMetaData string, AssetMetaTypeIsJsonNotOpaque bool, amount int, newGroupedAsset bool, groupedAsset bool, groupKey string, groupAnchor string, shortResponse bool) string {
	grpcHost := base.QueryConfigByKey("taproothost")
	tlsCertPath := filepath.Join(base.Configure("lit"), "tls.cert")
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
	client := mintrpc.NewMintClient(conn)
	var _assetVersion taprpc.AssetVersion
	if assetVersionIsV1 {
		_assetVersion = taprpc.AssetVersion_ASSET_VERSION_V1
	} else {
		_assetVersion = taprpc.AssetVersion_ASSET_VERSION_V0
	}
	var _assetType taprpc.AssetType
	if assetTypeIsCollectible {
		_assetType = taprpc.AssetType_COLLECTIBLE
	} else {
		_assetType = taprpc.AssetType_NORMAL
	}
	_assetMetaDataByteSlice := []byte(assetMetaData)
	var _assetMetaType taprpc.AssetMetaType
	if AssetMetaTypeIsJsonNotOpaque {
		_assetMetaType = taprpc.AssetMetaType_META_TYPE_JSON
	} else {
		_assetMetaType = taprpc.AssetMetaType_META_TYPE_OPAQUE
	}
	_groupKeyByteSlices := []byte(groupKey)

	request := &mintrpc.MintAssetRequest{
		Asset: &mintrpc.MintAsset{
			AssetVersion: _assetVersion,
			AssetType:    _assetType,
			Name:         name,
			AssetMeta: &taprpc.AssetMeta{
				Data: _assetMetaDataByteSlice,
				Type: _assetMetaType,
			},
			Amount:          uint64(amount),
			NewGroupedAsset: newGroupedAsset,
			GroupedAsset:    groupedAsset,
			GroupKey:        _groupKeyByteSlices,
			GroupAnchor:     groupAnchor,
		},
		ShortResponse: shortResponse,
	}
	response, err := client.MintAsset(context.Background(), request)
	if err != nil {
		fmt.Printf("%s mintrpc MintAsset Error: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

// MintAsset
//
//	@Description: Wraps the mintAsset, omitting most of the parameters and making them default values.
//	MintAsset will attempt to mint the set of assets (async by default to ensure proper batching) specified in the request.
//	The pending batch is returned that shows the other pending assets that are part of the next batch.
//	This call will block until the operation succeeds (asset is staged in the batch) or fails.
//	@param name
//	@param assetMetaData
//	@param amount
//	@return bool
//func MintAsset(name string, assetMetaData string, amount int) string {
//	return mintAsset(false, false, name, assetMetaData, false, amount, false, false, "", "", false)
//}

func MintAsset(name string, assetTypeIsCollectible bool, assetMetaData string, amount int, newGroupedAsset bool) string {
	return mintAsset(false, assetTypeIsCollectible, name, assetMetaData, false, amount, newGroupedAsset, false, "", "", false)
}

func AddGroupAsset(name string, assetTypeIsCollectible bool, assetMetaData string, amount int, groupKey string) string {
	return mintAsset(false, assetTypeIsCollectible, name, assetMetaData, false, amount, false, true, groupKey, "", false)

}

type NewMeta struct {
	Acronym     string `json:"acronym,omitempty"`
	Description string `json:"description,omitempty"`
	Image_Data  string `json:"image_data,omitempty"`
}

const (
	OPEN_IMAGE_FILE_ERROR  = "OPEN_IMAGE_FILE_ERROR"
	WRITE_IMAGE_FILE_ERROR = "WRITE_IMAGE_FILE_ERROR"
	Data_not_urldata       = "Data_not_urldata"
	DATA_NOT_IMAGE         = "DATA_NOT_IMAGE"
)

func CreateNewMeta(acronym string, description string, imagefile string) string {
	meta := NewMeta{
		Acronym:     acronym,
		Description: description,
	}
	if imagefile != "" {
		image, err := os.ReadFile(imagefile)
		if err != nil {
			fmt.Println("open image file is error:", err)
			return OPEN_IMAGE_FILE_ERROR
		}
		imageStr := dataurl.EncodeBytes(image)
		meta.Image_Data = imageStr
	}
	metastr, _ := json.Marshal(meta)
	return string(metastr)
}

func CheckMetaStandard(Meta string) bool {
	temp := &NewMeta{}
	err := json.Unmarshal([]byte(Meta), temp)
	if err != nil {
		return false
	}
	return true
}

func DecodeBase64ForImage(image string, dir string, name string) string {
	dataUrl, err := dataurl.DecodeString(image)
	if err != nil {
		return Data_not_urldata
	}
	ContentType := dataUrl.MediaType.ContentType()
	datatype := strings.Split(ContentType, "/")
	if datatype[0] != "image" {
		fmt.Println("is not image dataurl")
		return DATA_NOT_IMAGE
	}
	formatName := strings.Split(name, ".")
	file := filepath.Join(dir, formatName[0]+"."+datatype[1])
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("create new image error:", err)
		return OPEN_IMAGE_FILE_ERROR
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	_, err = f.Write(dataUrl.Data)
	if err != nil {
		fmt.Println("Write data fail:", err)
		return WRITE_IMAGE_FILE_ERROR
	}
	return file
}
