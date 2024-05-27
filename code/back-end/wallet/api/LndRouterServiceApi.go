package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/wallet/api/connect"
	"github.com/wallet/base"
	"google.golang.org/grpc"
	"io"
	"os"

	"path/filepath"
)

// SendPaymentV2
//
//	@Description: SendPaymentV2 attempts to route a payment described by the passed PaymentRequest to the final destination.
//	The call returns a stream of payment updates. When using this RPC, make sure to set a fee limit, as the default routing fee limit is 0 sats.
//	Without a non-zero fee limit only routes without fees will be attempted which often fails with FAILURE_REASON_NO_ROUTE.
//	@return string
func SendPaymentV2(invoice string, feelimit int64) string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := connect.NewTlsCert(tlsCertPath)
	macaroon := connect.GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(connect.NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := routerrpc.NewRouterClient(conn)
	request := &routerrpc.SendPaymentRequest{
		PaymentRequest: invoice,
		FeeLimitSat:    feelimit,
		TimeoutSeconds: 60,
	}
	stream, err := client.SendPaymentV2(context.Background(), request)
	if err != nil {
		fmt.Printf("%s routerrpc SendPaymentV2 :%v\n", GetTimeNow(), err)
		return "false"
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s err == io.EOF, err: %v\n", GetTimeNow(), err)
				return "false"
			}
			fmt.Printf("%s stream Recv err: %v\n", GetTimeNow(), err)
			return "false"
		}
		fmt.Printf("%s %v\n", GetTimeNow(), response)
		return response.PaymentHash
	}
}

// TrackPaymentV2
//
//	@Description: TrackPaymentV2 returns an update stream for the payment identified by the payment hash.
//	@return string
func TrackPaymentV2(payhash string) string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := connect.NewTlsCert(tlsCertPath)
	macaroon := connect.GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(connect.NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := routerrpc.NewRouterClient(conn)
	_payhashByteSlice, _ := hex.DecodeString(payhash)
	request := &routerrpc.TrackPaymentRequest{
		PaymentHash: _payhashByteSlice,
	}
	stream, err := client.TrackPaymentV2(context.Background(), request)

	if err != nil {
		fmt.Printf("%s client.SendPaymentV2 :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s err == io.EOF, err: %v\n", GetTimeNow(), err)
				return MakeJsonResult(false, err.Error(), nil)
			}
			fmt.Printf("%s stream Recv err: %v\n", GetTimeNow(), err)
			return MakeJsonResult(false, err.Error(), nil)
		}
		fmt.Printf("%s %v\n", GetTimeNow(), response)
		status := response.Status.String()
		return MakeJsonResult(true, "", status)
	}
}

// SendToRouteV2
//
//	@Description:SendToRouteV2 attempts to make a payment via the specified route.
//	This method differs from SendPayment in that it allows users to specify a full route manually.
//	This can be used for things like rebalancing, and atomic swaps.
//	@param route
//	skipped function SendToRouteV2 with unsupported parameter or return types
func SendToRouteV2(payhash []byte, route *lnrpc.Route) {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := connect.NewTlsCert(tlsCertPath)
	macaroon := connect.GetMacaroon(macaroonPath)
	cert, err := os.ReadFile(tlsCertPath)
	if err != nil {
		fmt.Printf("%s Failed to read cert file: %s", GetTimeNow(), err)
	}
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(connect.NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := routerrpc.NewRouterClient(conn)
	request := &routerrpc.SendToRouteRequest{
		PaymentHash: cert,
		Route:       route,
	}
	response, err := client.SendToRouteV2(context.Background(), request)
	if err != nil {
		fmt.Printf("%s routerrpc SendToRouteV2 :%v\n", GetTimeNow(), err)
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
}

// EstimateRouteFee
//
//	@Description: EstimateRouteFee allows callers to obtain a lower bound w.r.t how much it may cost to send an HTLC to the target end destination.
//	@return string
func EstimateRouteFee(dest string, amtsat int64) string {
	grpcHost := base.QueryConfigByKey("lndhost")
	tlsCertPath := filepath.Join(base.Configure("lnd"), "tls.cert")
	newFilePath := filepath.Join(base.Configure("lnd"), "."+"macaroonfile")
	macaroonPath := filepath.Join(newFilePath, "admin.macaroon")
	creds := connect.NewTlsCert(tlsCertPath)
	macaroon := connect.GetMacaroon(macaroonPath)
	conn, err := grpc.Dial(grpcHost, grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(connect.NewMacaroonCredential(macaroon)))
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("%s conn Close err: %v\n", GetTimeNow(), err)
		}
	}(conn)
	client := routerrpc.NewRouterClient(conn)

	bDest, _ := hex.DecodeString(dest)
	request := &routerrpc.RouteFeeRequest{
		Dest:   bDest,
		AmtSat: amtsat,
	}
	response, err := client.EstimateRouteFee(context.Background(), request)
	if err != nil {
		fmt.Printf("%s routerrpc EstimateRouteFee :%v\n", GetTimeNow(), err)
	}
	fmt.Printf("%s  %v\n", GetTimeNow(), response.RoutingFeeMsat)
	return response.String()
}
