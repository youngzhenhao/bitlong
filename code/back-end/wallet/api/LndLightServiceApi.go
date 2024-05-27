package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/wallet/api/connect"
	"io"
	"strconv"
	"strings"
)

// GetWalletBalance
//
//	@Description: WalletBalance returns total unspent outputs(confirmed and unconfirmed),
//	all confirmed unspent outputs and all unconfirmed unspent outputs under control of the wallet.
//	@return string
func getWalletBalance() (*lnrpc.WalletBalanceResponse, error) {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.WalletBalanceRequest{}

	response, err := client.WalletBalance(context.Background(), request)
	fmt.Printf("%s lnrpc WalletBalance response: %v\n", GetTimeNow(), response.String())
	return response, err

}

// GetInfoOfLnd
//
//	@Description: GetInfo returns general information concerning the lightning node including it's identity pubkey, alias,
//	the chains it is connected to, and information concerning the number of open+pending channels.
//	@return string
func getInfoOfLnd() (*lnrpc.GetInfoResponse, error) {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.GetInfoRequest{}
	response, err := client.GetInfo(context.Background(), request)
	return response, err
}

func sendCoins(addr string, amount int64, feeRate uint64, all bool) (*lnrpc.SendCoinsResponse, error) {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()

	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.SendCoinsRequest{
		Addr: addr,
	}
	if feeRate > 0 {
		request.SatPerVbyte = feeRate
	}
	if all {
		request.SendAll = true
	} else {
		request.Amount = amount
	}
	response, err := client.SendCoins(context.Background(), request)
	return response, err
}

func GetWalletBalance() string {
	response, err := getWalletBalance()
	if err != nil {
		fmt.Printf("%s lnrpc WalletBalance err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

func GetInfoOfLnd() string {
	response, err := getInfoOfLnd()
	if err != nil {
		fmt.Printf("%s lnrpc GetInfo err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

// GetIdentityPubkey
//
//	@Description: GetInfo returns general information concerning the lightning node including it's identity pubkey, alias,
//	the chains it is connected to, and information concerning the number of open+pending channels.
//	@return string
func GetIdentityPubkey() string {
	response, err := getInfoOfLnd()
	if err != nil {
		fmt.Printf("%s lnrpc GetInfo.IdentityPubkey err: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.GetIdentityPubkey()
}

// GetNewAddress
//
//	@Description:NewAddress creates a new address under control of the local wallet.
//	@return string
func GetNewAddress() string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.NewAddressRequest{}
	response, err := client.NewAddress(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc NewAddress err: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.Address
}

// AddInvoice
//
//	@Description:AddInvoice attempts to add a new invoice to the invoice database.
//	Any duplicated invoices are rejected, therefore all invoices must have a unique payment preimage.
//	@return string
func AddInvoice(value int64, memo string) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.Invoice{
		Value: value,
		Memo:  memo,
	}
	response, err := client.AddInvoice(context.Background(), request)
	if err != nil {
		fmt.Printf("%s client.AddInvoice :%v\n", GetTimeNow(), err)
		return ""
	}
	return response.GetPaymentRequest()
}

// ListInvoices
//
//	@Description:ListInvoices returns a list of all the invoices currently stored within the database.
//	Any active debug invoices are ignored. It has full support for paginated responses, allowing users to query for specific invoices through their add_index.
//	This can be done by using either the first_index_offset or last_index_offset fields included in the response as the index_offset of the next request.
//	By default, the first 100 invoices created will be returned. Backwards pagination is also supported through the Reversed flag.
//	@return string
func ListInvoices() string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ListInvoiceRequest{}
	response, err := client.ListInvoices(context.Background(), request)
	if err != nil {
		fmt.Printf("%s client.ListInvoice :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	invoices := SimplifyInvoice(response)
	return MakeJsonResult(true, "", invoices)
}

type InvoiceSimplified struct {
	PaymentRequest string `json:"payment_request"`
	Value          int    `json:"value"`
	State          string `json:"state"`
	CreationDate   int    `json:"creation_date"`
}

func SimplifyInvoice(invoice *lnrpc.ListInvoiceResponse) *[]InvoiceSimplified {
	var invoices []InvoiceSimplified
	for _, invoice := range invoice.Invoices {
		invoices = append(invoices, InvoiceSimplified{
			PaymentRequest: invoice.PaymentRequest,
			Value:          int(invoice.Value),
			State:          invoice.State.String(),
			CreationDate:   int(invoice.CreationDate),
		})
	}
	return &invoices
}

type InvoiceAll struct {
	Invoices []struct {
		Memo            string        `json:"memo"`
		RPreimage       string        `json:"r_preimage"`
		RHash           string        `json:"r_hash"`
		Value           string        `json:"value"`
		ValueMsat       string        `json:"value_msat"`
		Settled         bool          `json:"settled"`
		CreationDate    string        `json:"creation_date"`
		SettleDate      string        `json:"settle_date"`
		PaymentRequest  string        `json:"payment_request"`
		DescriptionHash string        `json:"description_hash"`
		Expiry          string        `json:"expiry"`
		FallbackAddr    string        `json:"fallback_addr"`
		CltvExpiry      string        `json:"cltv_expiry"`
		RouteHints      []interface{} `json:"route_hints"`
		Private         bool          `json:"private"`
		AddIndex        string        `json:"add_index"`
		SettleIndex     string        `json:"settle_index"`
		AmtPaid         string        `json:"amt_paid"`
		AmtPaidSat      string        `json:"amt_paid_sat"`
		AmtPaidMsat     string        `json:"amt_paid_msat"`
		State           string        `json:"state"`
		Htlcs           []interface{} `json:"htlcs"`
		Features        struct {
			Num9 struct {
				Name       string `json:"name"`
				IsRequired bool   `json:"is_required"`
				IsKnown    bool   `json:"is_known"`
			} `json:"9"`
			Num14 struct {
				Name       string `json:"name"`
				IsRequired bool   `json:"is_required"`
				IsKnown    bool   `json:"is_known"`
			} `json:"14"`
			Num17 struct {
				Name       string `json:"name"`
				IsRequired bool   `json:"is_required"`
				IsKnown    bool   `json:"is_known"`
			} `json:"17"`
		} `json:"features"`
		IsKeysend       bool   `json:"is_keysend"`
		PaymentAddr     string `json:"payment_addr"`
		IsAmp           bool   `json:"is_amp"`
		AmpInvoiceState struct {
		} `json:"amp_invoice_state"`
	} `json:"invoices"`
	LastIndexOffset  string `json:"last_index_offset"`
	FirstIndexOffset string `json:"first_index_offset"`
}

// LookupInvoice
//
//	@Description:LookupInvoice attempts to look up an invoice according to its payment hash.
//	The passed payment hash must be exactly 32 bytes, if not, an error is returned.
//	@return string
func LookupInvoice(rhash string) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	b_rhash, _ := hex.DecodeString(rhash)
	request := &lnrpc.PaymentHash{
		RHash: b_rhash,
	}
	response, err := client.LookupInvoice(context.Background(), request)
	if err != nil {
		fmt.Printf("%s client.LookupInvoice :%v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// AbandonChannel
//
//	@Description: AbandonChannel removes all channel state from the database except for a close summary.
//	This method can be used to get rid of permanently unusable channels due to bugs fixed in newer versions of lnd.
//	This method can also be used to remove externally funded channels where the funding transaction was never broadcast.
//	Only available for non-externally funded channels in dev build.
//	@return bool
func AbandonChannel() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.AbandonChannelRequest{}
	response, err := client.AbandonChannel(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc AbandonChannel err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// BatchOpenChannel
//
//	@Description: BatchOpenChannel attempts to open multiple single-funded channels in a single transaction in an atomic way.
//	This means either all channel open requests succeed at once or all attempts are aborted if any of them fail.
//	This is the safer variant of using PSBTs to manually fund a batch of channels through the OpenChannel RPC.
//	@return bool
func BatchOpenChannel() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.BatchOpenChannelRequest{}
	response, err := client.BatchOpenChannel(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc BatchOpenChannel err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// ChannelAcceptor
//
//	@Description: ChannelAcceptor dispatches a bi-directional streaming RPC in which OpenChannel requests are sent to the client and the client responds with a boolean that tells LND whether or not to accept the channel.
//	This allows node operators to specify their own criteria for accepting inbound channels through a single persistent connection.
//	@return bool
func ChannelAcceptor() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	stream, err := client.ChannelAcceptor(context.Background())
	if err != nil {
		fmt.Printf("%s lnrpc ChannelAcceptor err: %v\n", GetTimeNow(), err)
		return false
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s err == io.EOF, err: %v\n", GetTimeNow(), err)
				return false
			}
			fmt.Printf("%s stream Recv err: %v\n", GetTimeNow(), err)
			return false
		}
		fmt.Printf("%s %v\n", GetTimeNow(), response)
		return true
	}
}

// ChannelBalance
//
//	@Description: ChannelBalance returns a report on the total funds across all open channels, categorized in local/remote, pending local/remote and unsettled local/remote balances.
//	@return string
func ChannelBalance() string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChannelBalanceRequest{}
	response, err := client.ChannelBalance(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ChannelBalance err: %v\n", GetTimeNow(), err)
		return ""
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return response.String()
}

// CheckMacaroonPermissions
//
//	@Description: CheckMacaroonPermissions checks whether a request follows the constraints imposed on the macaroon and that the macaroon is authorized to follow the provided permissions.
//	@return bool
func CheckMacaroonPermissions() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.CheckMacPermRequest{}
	response, err := client.CheckMacaroonPermissions(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc CheckMacaroonPermissions err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// CloseChannel
//
//	@Description:CloseChannel attempts to close an active channel identified by its channel outpoint (ChannelPoint).
//	The actions of this method can additionally be augmented to attempt a force close after a timeout period in the case of an inactive peer.
//	If a non-force close (cooperative closure) is requested, then the user can specify either a target number of blocks until the closure transaction is confirmed, or a manual fee rate.
//	If neither are specified, then a default lax, block confirmation target is used.
//	@return bool
func CloseChannel(fundingTxidStr string, outputIndex int) bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)

	request := &lnrpc.CloseChannelRequest{
		ChannelPoint: &lnrpc.ChannelPoint{
			FundingTxid: &lnrpc.ChannelPoint_FundingTxidStr{FundingTxidStr: fundingTxidStr},
			OutputIndex: uint32(outputIndex),
		},
	}
	stream, err := client.CloseChannel(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc CloseChannel err: %v\n", GetTimeNow(), err)
		return false
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s err == io.EOF, err: %v\n", GetTimeNow(), err)
				return false
			}
			fmt.Printf("%s stream Recv err: %v\n", GetTimeNow(), err)
			return false
		}
		fmt.Printf("%s %v\n", GetTimeNow(), response)
		return true
	}

}

// ClosedChannels
//
//	@Description: ClosedChannels returns a description of all the closed channels that this node was a participant in.
//	@return string
func ClosedChannels() string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ClosedChannelsRequest{}
	response, err := client.ClosedChannels(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ClosedChannels err: %v\n", GetTimeNow(), err)
		return err.Error()
	}
	return response.String()
}

// ExportAllChannelBackups
//
//	@Description: ExportAllChannelBackups returns static channel backups for all existing channels known to lnd.
//	A set of regular singular static channel backups for each channel are returned.
//	Additionally, a multi-channel backup is returned as well, which contains a single encrypted blob containing the backups of each channel.
//	@return bool
func ExportAllChannelBackups() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChanBackupExportRequest{}
	response, err := client.ExportAllChannelBackups(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ChanBackupExportRequest err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// ExportChannelBackup
//
//	@Description: ExportChannelBackup attempts to return an encrypted static channel backup for the target channel identified by it channel point.
//	The backup is encrypted with a key generated from the aezeed seed of the user.
//	The returned backup can either be restored using the RestoreChannelBackup method once lnd is running, or via the InitWallet and UnlockWallet methods from the WalletUnlocker service.
//	@return bool
func ExportChannelBackup() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ExportChannelBackupRequest{}
	response, err := client.ExportChannelBackup(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ExportChannelBackup err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// GetChanInfo
//
//	@Description:GetChanInfo returns the latest authenticated network announcement for the given channel identified by its channel ID: an 8-byte integer which uniquely identifies the location of transaction's funding output within the blockchain.
//	@return string
func GetChanInfo(chanId string) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	chainIdUint64, err := strconv.ParseUint(chanId, 10, 64)
	if err != nil {
		fmt.Printf("%s string to uint64 err: %v\n", GetTimeNow(), err)
	}
	request := &lnrpc.ChanInfoRequest{
		ChanId: chainIdUint64,
	}
	response, err := client.GetChanInfo(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc GetChanInfo err: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// OpenChannelSync
//
//	@Description:OpenChannelSync is a synchronous version of the OpenChannel RPC call. This call is meant to be consumed by clients to the REST proxy.
//	As with all other sync calls, all byte slices are intended to be populated as hex encoded strings.
//	@return string
func OpenChannelSync(nodePubkey string, localFundingAmount int64) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	_nodePubkeyByteSlice, _ := hex.DecodeString(nodePubkey)
	request := &lnrpc.OpenChannelRequest{
		NodePubkey:         _nodePubkeyByteSlice,
		LocalFundingAmount: localFundingAmount,
	}
	response, err := client.OpenChannelSync(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc OpenChannelSync err: %v\n", GetTimeNow(), err)
		return err.Error()
	}
	//deal with  the byte-reversed hash
	var txBytes []byte
	for i := len(response.GetFundingTxidBytes()) - 1; i >= 0; {
		txBytes = append(txBytes, response.GetFundingTxidBytes()[i])
		i--
	}

	txStr := hex.EncodeToString(txBytes)
	return txStr + ":" + strconv.Itoa(int(response.GetOutputIndex()))
}

// OpenChannel
//
//	@Description:OpenChannel attempts to open a singly funded channel specified in the request to a remote peer.
//	Users are able to specify a target number of blocks that the funding transaction should be confirmed in, or a manual fee rate to us for the funding transaction.
//	If neither are specified, then a lax block confirmation target is used. Each OpenStatusUpdate will return the pending channel ID of the in-progress channel.
//	Depending on the arguments specified in the OpenChannelRequest, this pending channel ID can then be used to manually progress the channel funding flow.
//	@return bool
func OpenChannel(nodePubkey string, localFundingAmount int64) bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	_nodePubkeyByteSlice, _ := hex.DecodeString(nodePubkey)
	request := &lnrpc.OpenChannelRequest{
		NodePubkey:         _nodePubkeyByteSlice,
		LocalFundingAmount: localFundingAmount,
	}
	stream, err := client.OpenChannel(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc OpenChannel err: %v\n", GetTimeNow(), err)
		return false
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s err == io.EOF, err: %v\n", GetTimeNow(), err)
				return false
			}
			fmt.Printf("%s stream Recv err: %v\n", GetTimeNow(), err)
			return false
		}
		fmt.Printf("%s %v\n", GetTimeNow(), response)
		return true
	}
}

// ListChannels
//
//	@Description: ListChannels returns a description of all the open channels that this node is a participant in.
//	@return string
func ListChannels() string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ListChannelsRequest{}
	response, err := client.ListChannels(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ListChannels err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

// PendingChannels
//
//	@Description: PendingChannels returns a list of all the channels that are currently considered "pending".
//	A channel is pending if it has finished the funding workflow and is waiting for confirmations for the funding txn, or is in the process of closure, either initiated cooperatively or non-cooperatively.
//	@return string
func PendingChannels() string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.PendingChannelsRequest{}
	response, err := client.PendingChannels(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc PendingChannels err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

// FindChanPoint
//
//	@Description:FindChanPoint is a helper function that takes a channel point string and returns the channel state
//	@return string
//	ACTIVE: channel is active
//	INACTIVE: channel is inactive
//	PENDING_OPEN: channel is pending open
//	PENDING_CLOSE: channel is pending close
//	CLOSED: channel is closed
//	NO_FIND: channel not found
//	ERR: grpc error
func GetChannelState(chanPoint string) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()

	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ListChannelsRequest{}
	response, err := client.ListChannels(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ListChannels err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}

	var ChannelState string
	for _, channel := range response.Channels {
		if channel.ChannelPoint == chanPoint {
			if channel.Active {
				ChannelState = "ACTIVE"
			} else {
				ChannelState = "INACTIVE"
			}
			return MakeJsonResult(true, "", ChannelState)
		}
	}
	pendrequest := &lnrpc.PendingChannelsRequest{}
	pendingresponse, err := client.PendingChannels(context.Background(), pendrequest)
	if err != nil {
		fmt.Printf("%s lnrpc PendingChannels err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	for _, channel := range pendingresponse.PendingOpenChannels {
		if channel.Channel.ChannelPoint == chanPoint {

			ChannelState = "PENDING_OPEN"
			return MakeJsonResult(true, "", ChannelState)
		}
	}
	for _, channel := range pendingresponse.WaitingCloseChannels {
		if channel.Channel.ChannelPoint == chanPoint {
			ChannelState = "PENDING_CLOSE"
			return MakeJsonResult(true, "", ChannelState)
		}
	}

	closerequest := &lnrpc.ClosedChannelsRequest{}
	closeresponse, err := client.ClosedChannels(context.Background(), closerequest)
	if err != nil {
		fmt.Printf("%s lnrpc ClosedChannels err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	for _, channel := range closeresponse.Channels {
		if channel.ChannelPoint == chanPoint {
			ChannelState = "CLOSED"
			return MakeJsonResult(true, "", ChannelState)
		}
	}

	return MakeJsonResult(false, "NO_FIND_CHANNEL", nil)
}

// GetChanBalance
//
//	@Description:GetChanBalance returns the balance of a channel specified by its channel point.
//	@return string
//	（localBalance:remoteBalance) :local balance and remotebalance of the channel
//	ERR: grpc error
//	NO_FIND: channel not found
func GetChannelInfo(chanPoint string) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()

	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ListChannelsRequest{}
	response, err := client.ListChannels(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ListChannels err: %v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	for _, channel := range response.Channels {
		if channel.ChannelPoint == chanPoint {

			return MakeJsonResult(true, "", channel)
		}
	}
	return MakeJsonResult(false, "NO_FIND_CHANNEL", nil)
}

// RestoreChannelBackups
//
//	@Description:RestoreChannelBackups accepts a set of singular channel backups, or a single encrypted multi-chan backup and attempts to recover any funds remaining within the channel.
//	If we are able to unpack the backup, then the new channel will be shown under listchannels, as well as pending channels.
//	@return bool
func RestoreChannelBackups() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.RestoreChanBackupRequest{}
	response, err := client.RestoreChannelBackups(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc RestoreChannelBackups err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// SubscribeChannelBackups
//
//	@Description:SubscribeChannelBackups allows a client to sub-subscribe to the most up to date information concerning the state of all channel backups.
//	Each time a new channel is added, we return the new set of channels, along with a multi-chan backup containing the backup info for all channels.
//	Each time a channel is closed, we send a new update, which contains new new chan back ups, but the updated set of encrypted multi-chan backups with the closed channel(s) removed.
//	@return bool
func SubscribeChannelBackups() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChannelBackupSubscription{}
	stream, err := client.SubscribeChannelBackups(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc SubscribeChannelBackups err: %v\n", GetTimeNow(), err)
		return false
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s err == io.EOF, err: %v\n", GetTimeNow(), err)
				return false
			}
			fmt.Printf("%s stream Recv err: %v\n", GetTimeNow(), err)
			return false
		}
		fmt.Printf("%s %v\n", GetTimeNow(), response)
		return true
	}

}

// SubscribeChannelEvents
//
//	@Description: SubscribeChannelEvents creates a uni-directional stream from the server to the client in which any updates relevant to the state of the channels are sent over.
//	Events include new active channels, inactive channels, and closed channels.
//	@return bool
func SubscribeChannelEvents() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChannelEventSubscription{}
	stream, err := client.SubscribeChannelEvents(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc  err: %v\n", GetTimeNow(), err)
		return false
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s err == io.EOF, err: %v\n", GetTimeNow(), err)
				return false
			}
			fmt.Printf("%s stream Recv err: %v\n", GetTimeNow(), err)
			return false
		}
		fmt.Printf("%s %v\n", GetTimeNow(), response)
		return true
	}

}

// SubscribeChannelGraph
//
//	@Description: SubscribeChannelGraph launches a streaming RPC that allows the caller to receive notifications upon any changes to the channel graph topology from the point of view of the responding node.
//	Events notified include: new nodes coming online, nodes updating their authenticated attributes, new channels being advertised, updates in the routing policy for a directional channel edge, and when channels are closed on-chain.
//	@return bool
func SubscribeChannelGraph() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.GraphTopologySubscription{}
	stream, err := client.SubscribeChannelGraph(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc SubscribeChannelGraph err: %v\n", GetTimeNow(), err)
		return false
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s err == io.EOF, err: %v\n", GetTimeNow(), err)
				return false
			}
			fmt.Printf("%s stream Recv err: %v\n", GetTimeNow(), err)
			return false
		}
		fmt.Printf("%s %v\n", GetTimeNow(), response)
		return true
	}

}

// UpdateChannelPolicy
//
//	@Description: UpdateChannelPolicy allows the caller to update the fee schedule and channel policies for all channels globally, or a particular channel.
//	@return bool
func UpdateChannelPolicy() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.PolicyUpdateRequest{}
	response, err := client.UpdateChannelPolicy(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc UpdateChannelPolicy err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// VerifyChanBackup
//
//	@Description: VerifyChanBackup allows a caller to verify the integrity of a channel backup snapshot.
//	This method will accept either a packed Single or a packed Multi. Specifying both will result in an error.
//	@return bool
func VerifyChanBackup() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ChanBackupSnapshot{}
	response, err := client.VerifyChanBackup(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc VerifyChanBackup err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// ConnectPeer
//
//	@Description:ConnectPeer attempts to establish a connection to a remote peer. This is at the networking level, and is used for communication between nodes.
//	This is distinct from establishing a channel with a peer.
//	@return bool
func ConnectPeer(pubkey, host string) bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ConnectPeerRequest{
		Addr: &lnrpc.LightningAddress{
			Pubkey: pubkey,
			Host:   host,
		},
	}
	response, err := client.ConnectPeer(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ConnectPeer err: %v\n", GetTimeNow(), err)
		if strings.Contains(err.Error(), "already connected to peer") {
			return true
		} else {
			return false
		}
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

// EstimateFee
//
//	@Description:EstimateFee asks the chain backend to estimate the fee rate and total fees for a transaction that pays to multiple specified outputs.
//	When using REST, the AddrToAmount map type can be set by appending &AddrToAmount[<address>]=<amount_to_send> to the URL.
//	Unfortunately this map type doesn't appear in the REST API documentation because of a bug in the grpc-gateway library.
//	@return string
func EstimateFee(addr string, amount int64) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	addrToAmount := make(map[string]int64)
	addrToAmount[addr] = amount
	request := &lnrpc.EstimateFeeRequest{
		AddrToAmount: addrToAmount,
	}
	response, err := client.EstimateFee(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ConnectPeer err: %v\n", GetTimeNow(), err)
		return ""
	}
	return response.String()
}

// DecodePayReq
//
//	@Description:DecodePayReq takes an encoded payment request string and attempts to decode it, returning a full description of the conditions encoded within the payment request.
//	@return int64
func DecodePayReq(payReq string) int64 {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.PayReqString{
		PayReq: payReq,
	}
	response, err := client.DecodePayReq(context.Background(), request)
	if err != nil {
		fmt.Printf("%s client.DecodePayReq :%v\n", GetTimeNow(), err)
		return -1
	}
	return response.NumSatoshis
}

// SendPaymentSync
//
//	@Description:SendPaymentSync is the synchronous non-streaming version of SendPayment.
//	This RPC is intended to be consumed by clients of the REST proxy. Additionally, this RPC expects the destination's public key and the payment hash (if any) to be encoded as hex strings.
//	@return string
func SendPaymentSync(invoice string) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.SendRequest{
		PaymentRequest: invoice,
	}
	response, err := client.SendPaymentSync(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc SendPaymentSync :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, err.Error(), nil)
	}
	paymentHash := hex.EncodeToString(response.PaymentHash)
	return MakeJsonResult(true, "", paymentHash)
}

func SendPaymentSync0amt(invoice string, amt int64) string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.SendRequest{
		PaymentRequest: invoice,
		Amt:            amt,
	}
	stream, err := client.SendPaymentSync(context.Background(), request)
	if err != nil {
		fmt.Printf("%s client.SendPaymentSync :%v\n", GetTimeNow(), err)
		return "false"
	}
	fmt.Printf(GetTimeNow() + stream.String())
	return hex.EncodeToString(stream.PaymentHash)
}

// SendCoins
//
//	@Description: SendCoins executes a request to send coins to a particular address. Unlike SendMany, this RPC call only allows creating a single output at a time.
//	If neither target_conf, or sat_per_vbyte are set, then the internal wallet will consult its fee model to determine a fee for the default confirmation target.
//	@return string
func SendCoins(addr string, amount int64, feeRate int64, sendAll bool) string {
	response, err := sendCoins(addr, amount, uint64(feeRate), sendAll)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

// jsonaddr :{"bcrt1pq83tk5uu0lpwk2gd7f736ttrmexed8xazfz3jmwj0ml26cwyurast4xk3w":1111,"bcrt1pra9w5dphnx75n0pjzcxlc5e8k9vg9sdupttyr36prn2t6ullr9eq0utvac":2222}
func SendMany(jsonAddr string, feeRate int64) string {
	var addrs []struct {
		Address string `json:"address"`
		Amount  int64  `json:"btcSum"`
	}
	err := json.Unmarshal([]byte(jsonAddr), &addrs)
	if err != nil {
		return MakeJsonResult(false, "Please use the correct json format", nil)
	}
	if len(addrs) == 0 {
		return MakeJsonResult(false, "Please input the correct address and amount", nil)
	}
	addrTo := make(map[string]int64)
	for _, addr := range addrs {
		addrTo[addr.Address] = addr.Amount
	}
	response, err := sendMany(addrTo, uint64(feeRate))
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

func sendMany(addr map[string]int64, feerate uint64) (*lnrpc.SendManyResponse, error) {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()

	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.SendManyRequest{
		AddrToAmount: addr,
	}
	if feerate > 0 {
		request.SatPerVbyte = feerate
	}
	response, err := client.SendMany(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Deprecated: Please Use SendCoins
func SendAllCoins(addr string) string {
	response, err := sendCoins(addr, 0, 0, true)
	if err != nil {
		return MakeJsonResult(false, err.Error(), nil)
	}
	return MakeJsonResult(true, "", response)
}

// LndStopDaemon
//
//	@Description: Stop gracefully shuts down the Pool trader daemon.
//	@return bool
func LndStopDaemon() bool {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()

	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.StopRequest{}
	response, err := client.StopDaemon(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc StopDaemon err: %v\n", GetTimeNow(), err)
		return false
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return true
}

func ListPermissions() string {
	conn, clearUp, err := connect.GetConnection("lnd", false)
	if err != nil {
		fmt.Printf("%s did not connect: %v\n", GetTimeNow(), err)
	}

	defer clearUp()
	client := lnrpc.NewLightningClient(conn)
	request := &lnrpc.ListPermissionsRequest{}
	response, err := client.ListPermissions(context.Background(), request)
	if err != nil {
		fmt.Printf("%s lnrpc ListPermissions err: %v\n", GetTimeNow(), err)
		return err.Error()
	}
	fmt.Printf("%s %v\n", GetTimeNow(), response)
	return response.String()
}
