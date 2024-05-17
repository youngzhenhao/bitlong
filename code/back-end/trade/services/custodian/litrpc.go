package custodian

import (
	"context"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"strconv"
	"trade/config"
	"trade/utils"
)

func accountCreate(balance uint64, expirationDate int64, label string) (string, string) {
	litdconf := config.GetConfig().ApiConfig.Litd

	grpcHost := litdconf.Host + ":" + strconv.Itoa(litdconf.Port)
	tlsCertPath := litdconf.TlsCertPath
	macaroonPath := litdconf.MacaroonPath

	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()

	request := &litrpc.CreateAccountRequest{
		AccountBalance: balance,
		ExpirationDate: expirationDate,
		Label:          label,
	}
	client := litrpc.NewAccountsClient(conn)
	response, err := client.CreateAccount(context.Background(), request)
	if err != nil {
		return "Error: " + err.Error(), ""
	}
	return response.String(), response.Account.Id
}

func accountInfo(id string) string {
	litdconf := config.GetConfig().ApiConfig.Litd

	grpcHost := litdconf.Host + ":" + strconv.Itoa(litdconf.Port)
	tlsCertPath := litdconf.TlsCertPath
	macaroonPath := litdconf.MacaroonPath

	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()

	request := &litrpc.AccountInfoRequest{
		Id: id,
	}
	client := litrpc.NewAccountsClient(conn)
	response, err := client.AccountInfo(context.Background(), request)
	if err != nil {
		return "Error: " + err.Error()
	}
	return response.String()
}

func accountRemove(id string) string {
	litdconf := config.GetConfig().ApiConfig.Litd

	grpcHost := litdconf.Host + ":" + strconv.Itoa(litdconf.Port)
	tlsCertPath := litdconf.TlsCertPath
	macaroonPath := litdconf.MacaroonPath

	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()

	request := &litrpc.RemoveAccountRequest{
		Id: id,
	}
	client := litrpc.NewAccountsClient(conn)
	response, err := client.RemoveAccount(context.Background(), request)
	if err != nil {
		return "Error: " + err.Error()
	}
	return response.String()

}

func accountUpdate(id string, balance int64, expirationDate int64) string {
	litdconf := config.GetConfig().ApiConfig.Litd

	grpcHost := litdconf.Host + ":" + strconv.Itoa(litdconf.Port)
	tlsCertPath := litdconf.TlsCertPath
	macaroonPath := litdconf.MacaroonPath

	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()

	request := &litrpc.UpdateAccountRequest{
		Id:             id,
		AccountBalance: balance,
		ExpirationDate: expirationDate,
	}
	client := litrpc.NewAccountsClient(conn)
	response, err := client.UpdateAccount(context.Background(), request)
	if err != nil {
		return "Error: " + err.Error()
	}
	return response.String()
}

func acountList() string {
	litdconf := config.GetConfig().ApiConfig.Litd

	grpcHost := litdconf.Host + ":" + strconv.Itoa(litdconf.Port)
	tlsCertPath := litdconf.TlsCertPath
	macaroonPath := litdconf.MacaroonPath

	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()

	request := &litrpc.ListAccountsRequest{}
	client := litrpc.NewAccountsClient(conn)
	response, err := client.ListAccounts(context.Background(), request)
	if err != nil {
		return "Error: " + err.Error()
	}
	return response.String()
}

func getLitdInfo() string {
	litdconf := config.GetConfig().ApiConfig.Litd

	grpcHost := litdconf.Host + ":" + strconv.Itoa(litdconf.Port)
	tlsCertPath := litdconf.TlsCertPath
	macaroonPath := litdconf.MacaroonPath

	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()

	request := &litrpc.GetInfoRequest{}

	client := litrpc.NewProxyClient(conn)
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		return "Error: " + err.Error()
	}
	return response.String()
}

func getStatus() string {
	litdconf := config.GetConfig().ApiConfig.Litd

	grpcHost := litdconf.Host + ":" + strconv.Itoa(litdconf.Port)
	tlsCertPath := litdconf.TlsCertPath
	macaroonPath := litdconf.MacaroonPath

	conn, connClose := utils.GetConn(grpcHost, tlsCertPath, macaroonPath)
	defer connClose()

	request := &litrpc.SubServerStatusReq{}
	client := litrpc.NewStatusClient(conn)
	response, err := client.SubServerStatus(context.Background(), request)
	if err != nil {
		return "Error: " + err.Error()
	}
	return response.String()
}
