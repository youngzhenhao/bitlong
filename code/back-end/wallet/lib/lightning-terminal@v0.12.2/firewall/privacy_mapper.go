package firewall

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/lightninglabs/lightning-terminal/firewalldb"
	mid "github.com/lightninglabs/lightning-terminal/rpcmiddleware"
	"github.com/lightninglabs/lightning-terminal/session"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/protobuf/proto"
)

const (
	// privacyMapperName is the name of the RequestLogger interceptor.
	privacyMapperName = "lit-privacy-mapper"

	// amountVariation and timeVariation are used to set the randomization
	// of amounts and timestamps that are sent to the autopilot. Changing
	// these values may lead to unintended consequences in the behavior of
	// the autpilot.
	amountVariation = 0.05
	timeVariation   = time.Duration(10) * time.Minute

	// minTimeVariation and maxTimeVariation are the acceptable bounds
	// between which timeVariation can be set.
	minTimeVariation = time.Minute
	maxTimeVariation = time.Duration(24) * time.Hour

	// min and maxChanIDLen are the lengths to consider an int to be a
	// channel id. 13 corresponds to block height 1 and 20 to block height
	// 10_000_000.
	minChanIDLen = 13
	maxChanIDLen = 20

	// pubKeyLen is the length of a node pubkey.
	pubKeyLen = 66
)

var (
	// ErrNotSupportedByPrivacyMapper indicates that the invoked RPC method
	// is not supported by the privacy mapper.
	ErrNotSupportedByPrivacyMapper = errors.New("this RPC call is not " +
		"supported by the privacy mapper interceptor")
)

// A compile-time assertion that PrivacyMapper is a
// rpcmiddleware.RequestInterceptor.
var _ mid.RequestInterceptor = (*PrivacyMapper)(nil)

// PrivacyMapper is a RequestInterceptor that maps any pseudo names in certain
// requests to their real values and vice versa for responses.
type PrivacyMapper struct {
	newDB            firewalldb.NewPrivacyMapDB
	randIntn         func(int) (int, error)
	sessionIDIndexDB session.IDToGroupIndex
}

// NewPrivacyMapper returns a new instance of PrivacyMapper. The randIntn
// function is used to draw randomness for request field obfuscation.
func NewPrivacyMapper(newDB firewalldb.NewPrivacyMapDB,
	randIntn func(int) (int, error),
	sessionIDIndexDB session.IDToGroupIndex) *PrivacyMapper {

	return &PrivacyMapper{
		newDB:            newDB,
		randIntn:         randIntn,
		sessionIDIndexDB: sessionIDIndexDB,
	}
}

// Name returns the name of the interceptor.
func (p *PrivacyMapper) Name() string {
	return privacyMapperName
}

// ReadOnly returns true if this interceptor should be registered in read-only
// mode. In read-only mode no custom caveat name can be specified.
func (p *PrivacyMapper) ReadOnly() bool {
	return false
}

// CustomCaveatName returns the name of the custom caveat that is expected to be
// handled by this interceptor. Cannot be specified in read-only mode.
func (p *PrivacyMapper) CustomCaveatName() string {
	return CondPrivacy
}

// Intercept processes an RPC middleware interception request and returns the
// interception result which either accepts or rejects the intercepted message.
func (p *PrivacyMapper) Intercept(ctx context.Context,
	req *lnrpc.RPCMiddlewareRequest) (*lnrpc.RPCMiddlewareResponse, error) {

	ri, err := NewInfoFromRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error parsing incoming RPC middleware "+
			"interception request: %v", err)
	}

	sessionID, err := session.IDFromMacaroon(ri.Macaroon)
	if err != nil {
		return nil, fmt.Errorf("could not extract ID from macaroon")
	}

	// Get group ID for session ID.
	groupID, err := p.sessionIDIndexDB.GetGroupID(sessionID)
	if err != nil {
		return nil, err
	}

	log.Tracef("PrivacyMapper: Intercepting %v", ri)

	switch r := req.InterceptType.(type) {
	case *lnrpc.RPCMiddlewareRequest_StreamAuth:
		return mid.RPCErr(req, fmt.Errorf("streams unsupported"))

	// Parse incoming requests and act on them.
	case *lnrpc.RPCMiddlewareRequest_Request:
		msg, err := mid.ParseProtobuf(
			r.Request.TypeName, r.Request.Serialized,
		)
		if err != nil {
			return mid.RPCErrString(req, "error parsing proto: %v",
				err)
		}

		replacement, err := p.checkAndReplaceIncomingRequest(
			ctx, r.Request.MethodFullUri, msg, groupID,
		)
		if err != nil {
			return mid.RPCErr(req, err)
		}

		// No error occurred but the response should be replaced with
		// the given custom response. Wrap it in the correct RPC
		// response of the interceptor now.
		if replacement != nil {
			return mid.RPCReplacement(req, replacement)
		}

		// No error and no replacement, just return an empty response of
		// the correct type.
		return mid.RPCOk(req)

	// Parse and possibly manipulate outgoing responses.
	case *lnrpc.RPCMiddlewareRequest_Response:
		if ri.IsError {
			// TODO(elle): should we replace all litd errors with
			// a generic error?
			return mid.RPCOk(req)
		}

		msg, err := mid.ParseProtobuf(
			r.Response.TypeName, r.Response.Serialized,
		)
		if err != nil {
			return mid.RPCErrString(req, "error parsing proto: %v",
				err)
		}

		replacement, err := p.replaceOutgoingResponse(
			ctx, r.Response.MethodFullUri, msg, groupID,
		)
		if err != nil {
			return mid.RPCErr(req, err)
		}

		// No error occurred but the response should be replaced with
		// the given custom response. Wrap it in the correct RPC
		// response of the interceptor now.
		if replacement != nil {
			return mid.RPCReplacement(req, replacement)
		}

		// No error and no replacement, just return an empty response of
		// the correct type.
		return mid.RPCOk(req)

	default:
		return mid.RPCErrString(req, "invalid intercept type: %v", r)
	}
}

// checkAndReplaceIncomingRequest inspects an incoming request and optionally
// modifies some of the request parameters.
func (p *PrivacyMapper) checkAndReplaceIncomingRequest(ctx context.Context,
	uri string, req proto.Message, groupID session.ID) (proto.Message,
	error) {

	db := p.newDB(groupID)

	// If we don't have a handler for the URI, we don't allow the request
	// to go through.
	checker, ok := p.checkers(db)[uri]
	if !ok {
		return nil, ErrNotSupportedByPrivacyMapper
	}

	// This is just a sanity check to make sure the implementation for the
	// checker actually matches the correct request type.
	if !checker.HandlesRequest(req.ProtoReflect().Type()) {
		return nil, fmt.Errorf("invalid implementation, checker for "+
			"URI %s does not accept request of type %v", uri,
			req.ProtoReflect().Type())
	}

	return checker.HandleRequest(ctx, req)
}

// replaceOutgoingResponse inspects the responses before sending them out to the
// client and replaces them if needed.
func (p *PrivacyMapper) replaceOutgoingResponse(ctx context.Context, uri string,
	resp proto.Message, groupID session.ID) (proto.Message, error) {

	db := p.newDB(groupID)

	// If we don't have a handler for the URI, we don't allow the response
	// to go to avoid accidental leaks.
	checker, ok := p.checkers(db)[uri]
	if !ok {
		return nil, ErrNotSupportedByPrivacyMapper
	}

	// This is just a sanity check to make sure the implementation for the
	// checker actually matches the correct response type.
	if !checker.HandlesResponse(resp.ProtoReflect().Type()) {
		return nil, fmt.Errorf("invalid implementation, checker for "+
			"URI %s does not accept response of type %v", uri,
			resp.ProtoReflect().Type())
	}

	return checker.HandleResponse(ctx, resp)
}

func (p *PrivacyMapper) checkers(
	db firewalldb.PrivacyMapDB) map[string]mid.RoundTripChecker {

	return map[string]mid.RoundTripChecker{
		"/lnrpc.Lightning/GetInfo": mid.NewResponseRewriter(
			&lnrpc.GetInfoRequest{}, &lnrpc.GetInfoResponse{},
			handleGetInfoResponse(db), mid.PassThroughErrorHandler,
		),
		"/lnrpc.Lightning/ForwardingHistory": mid.NewResponseRewriter(
			&lnrpc.ForwardingHistoryRequest{},
			&lnrpc.ForwardingHistoryResponse{},
			handleFwdHistoryResponse(db, p.randIntn),
			mid.PassThroughErrorHandler,
		),
		"/lnrpc.Lightning/FeeReport": mid.NewResponseRewriter(
			&lnrpc.FeeReportRequest{}, &lnrpc.FeeReportResponse{},
			handleFeeReportResponse(db),
			mid.PassThroughErrorHandler,
		),
		"/lnrpc.Lightning/ListChannels": mid.NewFullRewriter(
			&lnrpc.ListChannelsRequest{},
			&lnrpc.ListChannelsResponse{},
			handleListChannelsRequest(db),
			handleListChannelsResponse(db, p.randIntn),
			mid.PassThroughErrorHandler,
		),
		"/lnrpc.Lightning/UpdateChannelPolicy": mid.NewFullRewriter(
			&lnrpc.PolicyUpdateRequest{},
			&lnrpc.PolicyUpdateResponse{},
			handleUpdatePolicyRequest(db),
			handleUpdatePolicyResponse(db),
			mid.PassThroughErrorHandler,
		),
	}
}

func handleGetInfoResponse(db firewalldb.PrivacyMapDB) func(ctx context.Context,
	r *lnrpc.GetInfoResponse) (proto.Message, error) {

	return func(ctx context.Context, r *lnrpc.GetInfoResponse) (
		proto.Message, error) {

		var pseudoPubKey string
		err := db.Update(
			func(tx firewalldb.PrivacyMapTx) error {
				var err error
				pseudoPubKey, err = firewalldb.HideString(
					tx, r.IdentityPubkey,
				)
				if err != nil {
					return err
				}

				return nil
			},
		)
		if err != nil {
			return nil, err
		}

		return &lnrpc.GetInfoResponse{
			// We purposefully hide our alias and URIs from the
			// autopilot server.
			Alias:                  "",
			Color:                  "",
			Uris:                   nil,
			Version:                r.Version,
			CommitHash:             r.CommitHash,
			IdentityPubkey:         pseudoPubKey,
			NumPendingChannels:     r.NumPendingChannels,
			NumActiveChannels:      r.NumActiveChannels,
			NumInactiveChannels:    r.NumInactiveChannels,
			NumPeers:               r.NumPeers,
			BlockHeight:            r.BlockHeight,
			BlockHash:              r.BlockHash,
			BestHeaderTimestamp:    r.BestHeaderTimestamp,
			SyncedToChain:          r.SyncedToChain,
			SyncedToGraph:          r.SyncedToGraph,
			Testnet:                r.Testnet,
			Chains:                 r.Chains,
			Features:               r.Features,
			RequireHtlcInterceptor: r.RequireHtlcInterceptor,
		}, nil
	}
}

func handleFwdHistoryResponse(db firewalldb.PrivacyMapDB,
	randIntn func(int) (int, error)) func(ctx context.Context,
	r *lnrpc.ForwardingHistoryResponse) (proto.Message, error) {

	return func(_ context.Context, r *lnrpc.ForwardingHistoryResponse) (
		proto.Message, error) {

		fwdEvents := make(
			[]*lnrpc.ForwardingEvent, len(r.ForwardingEvents),
		)

		err := db.Update(func(tx firewalldb.PrivacyMapTx) error {
			for i, fe := range r.ForwardingEvents {
				// Deterministically hide channel ids.
				chanIn, err := firewalldb.HideUint64(
					tx, fe.ChanIdIn,
				)
				if err != nil {
					return err
				}

				chanOut, err := firewalldb.HideUint64(
					tx, fe.ChanIdOut,
				)
				if err != nil {
					return err
				}

				// We randomize the outgoing amount for privacy.
				amtOutMsat, err := hideAmount(
					randIntn, amountVariation,
					fe.AmtOutMsat,
				)
				if err != nil {
					return err
				}

				// We randomize fees for privacy.
				feeMsat, err := hideAmount(
					randIntn, amountVariation, fe.FeeMsat,
				)
				if err != nil {
					return err
				}

				// Populate other fields in a consistent manner.
				amtInMsat := amtOutMsat + feeMsat
				amtOut := amtOutMsat / 1000
				amtIn := amtInMsat / 1000
				fee := feeMsat / 1000

				// We randomize the forwarding timestamp.
				timestamp, err := hideTimestamp(
					randIntn, timeVariation,
					time.Unix(0, int64(fe.TimestampNs)),
				)
				if err != nil {
					return err
				}

				fwdEvents[i] = &lnrpc.ForwardingEvent{
					ChanIdIn:   chanIn,
					ChanIdOut:  chanOut,
					AmtIn:      amtIn,
					AmtOut:     amtOut,
					Fee:        fee,
					FeeMsat:    feeMsat,
					AmtInMsat:  amtInMsat,
					AmtOutMsat: amtOutMsat,
					TimestampNs: uint64(
						timestamp.UnixNano(),
					),
					Timestamp: uint64(
						timestamp.Unix(),
					),
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}

		return &lnrpc.ForwardingHistoryResponse{
			ForwardingEvents: fwdEvents,
			LastOffsetIndex:  r.LastOffsetIndex,
		}, nil
	}
}

func handleFeeReportResponse(db firewalldb.PrivacyMapDB) func(
	ctx context.Context, r *lnrpc.FeeReportResponse) (proto.Message,
	error) {

	return func(ctx context.Context, r *lnrpc.FeeReportResponse) (
		proto.Message, error) {

		chanFees := make([]*lnrpc.ChannelFeeReport, len(r.ChannelFees))

		err := db.Update(func(tx firewalldb.PrivacyMapTx) error {
			for i, c := range r.ChannelFees {
				chanID, err := firewalldb.HideUint64(
					tx, c.ChanId,
				)
				if err != nil {
					return err
				}

				chanPoint, err := firewalldb.HideChanPointStr(
					tx, c.ChannelPoint,
				)
				if err != nil {
					return err
				}

				chanFees[i] = &lnrpc.ChannelFeeReport{
					ChanId:       chanID,
					ChannelPoint: chanPoint,
					BaseFeeMsat:  c.BaseFeeMsat,
					FeePerMil:    c.FeePerMil,
					FeeRate:      c.FeeRate,
				}
			}

			return nil
		})
		if err != nil {
			return nil, err
		}

		return &lnrpc.FeeReportResponse{
			ChannelFees: chanFees,
			DayFeeSum:   r.DayFeeSum,
			WeekFeeSum:  r.WeekFeeSum,
			MonthFeeSum: r.MonthFeeSum,
		}, nil
	}
}

func handleListChannelsRequest(db firewalldb.PrivacyMapDB) func(
	ctx context.Context, r *lnrpc.ListChannelsRequest) (proto.Message,
	error) {

	return func(ctx context.Context, r *lnrpc.ListChannelsRequest) (
		proto.Message, error) {

		if len(r.Peer) == 0 {
			return nil, nil
		}

		err := db.View(func(tx firewalldb.PrivacyMapTx) error {
			peer, err := firewalldb.RevealBytes(tx, r.Peer)
			if err != nil {
				return err
			}

			r.Peer = peer
			return nil
		})
		if err != nil {
			return nil, err
		}

		return r, nil
	}
}

func handleListChannelsResponse(db firewalldb.PrivacyMapDB,
	randIntn func(int) (int, error)) func(ctx context.Context,
	r *lnrpc.ListChannelsResponse) (proto.Message, error) {

	return func(_ context.Context, r *lnrpc.ListChannelsResponse) (
		proto.Message, error) {

		hideAmount := func(a int64) (int64, error) {
			hiddenAmount, err := hideAmount(
				randIntn, amountVariation, uint64(a),
			)
			if err != nil {
				return 0, err
			}

			return int64(hiddenAmount), nil
		}

		channels := make([]*lnrpc.Channel, len(r.Channels))

		err := db.Update(func(tx firewalldb.PrivacyMapTx) error {
			for i, c := range r.Channels {
				// Deterministically hide the peer pubkey,
				// the channel point, and the channel id.
				remotePub, err := firewalldb.HideString(
					tx, c.RemotePubkey,
				)
				if err != nil {
					return err
				}

				chanPoint, err := firewalldb.HideChanPointStr(
					tx, c.ChannelPoint,
				)
				if err != nil {
					return err
				}

				chanID, err := firewalldb.HideUint64(
					tx, c.ChanId,
				)
				if err != nil {
					return err
				}

				// We hide the initiator.
				initiator, err := hideBool(randIntn)
				if err != nil {
					return err
				}

				// Consider the capacity to be public
				// information. We don't care about reserves, as
				// having some funds as a balance is the normal
				// state over the lifetime of a channel. The
				// balance would be zero only for the initial
				// state as a non-funder.

				// We randomize local/remote balances.
				localBalance, err := hideAmount(c.LocalBalance)
				if err != nil {
					return err
				}

				// We may have a too large value for the local
				// balance, restrict it to the capacity.
				if localBalance > c.Capacity {
					localBalance = c.Capacity
				}

				// We adapt the remote balance accordingly.
				remoteBalance := c.Capacity - localBalance

				// We hide the total sats sent and received.
				satsReceived, err := hideAmount(
					c.TotalSatoshisReceived,
				)
				if err != nil {
					return err
				}

				satsSent, err := hideAmount(
					c.TotalSatoshisSent,
				)
				if err != nil {
					return err
				}

				// We only keep track of the _number_ of
				// unsettled HTLCs.
				pendingHtlcs := make(
					[]*lnrpc.HTLC, len(c.PendingHtlcs),
				)

				// We hide the unsettled balance.
				unsettled, err := hideAmount(c.UnsettledBalance)
				if err != nil {
					return err
				}

				//nolint:lll
				channels[i] = &lnrpc.Channel{
					// Items we adjust.
					RemotePubkey:          remotePub,
					ChannelPoint:          chanPoint,
					ChanId:                chanID,
					Initiator:             initiator,
					LocalBalance:          localBalance,
					RemoteBalance:         remoteBalance,
					TotalSatoshisReceived: satsReceived,
					TotalSatoshisSent:     satsSent,
					UnsettledBalance:      unsettled,
					PendingHtlcs:          pendingHtlcs,

					// Items that we zero out.
					CloseAddress:          "",
					PushAmountSat:         0,
					AliasScids:            nil,
					ZeroConfConfirmedScid: 0,

					// Items we keep as is.
					Active:               c.Active,
					Capacity:             c.Capacity,
					CommitFee:            c.CommitFee,
					CommitWeight:         c.CommitWeight,
					FeePerKw:             c.FeePerKw,
					NumUpdates:           c.NumUpdates,
					CsvDelay:             c.CsvDelay,
					Private:              c.Private,
					ChanStatusFlags:      c.ChanStatusFlags,
					LocalChanReserveSat:  c.LocalChanReserveSat,
					RemoteChanReserveSat: c.RemoteChanReserveSat,
					StaticRemoteKey:      c.StaticRemoteKey,
					CommitmentType:       c.CommitmentType,
					Lifetime:             c.Lifetime,
					Uptime:               c.Uptime,
					ThawHeight:           c.ThawHeight,
					LocalConstraints:     c.LocalConstraints,
					RemoteConstraints:    c.RemoteConstraints,
					ZeroConf:             c.ZeroConf,
				}
			}

			return nil
		})
		if err != nil {
			return nil, err
		}

		return &lnrpc.ListChannelsResponse{
			Channels: channels,
		}, nil
	}
}

func handleUpdatePolicyRequest(db firewalldb.PrivacyMapDB) func(
	ctx context.Context, r *lnrpc.PolicyUpdateRequest) (proto.Message,
	error) {

	return func(ctx context.Context, r *lnrpc.PolicyUpdateRequest) (
		proto.Message, error) {

		chanPoint := r.GetChanPoint()

		// If no channel point is specified then the
		// update request applies globally.
		if chanPoint == nil {
			return nil, nil
		}

		txid, err := lnrpc.GetChanPointFundingTxid(chanPoint)
		if err != nil {
			return nil, err
		}

		index := chanPoint.GetOutputIndex()

		var (
			newTxid  string
			newIndex uint32
		)
		err = db.View(func(tx firewalldb.PrivacyMapTx) error {
			var err error
			newTxid, newIndex, err = firewalldb.RevealChanPoint(
				tx, txid.String(), index,
			)
			return err
		})
		if err != nil {
			return nil, err
		}

		r.Scope = &lnrpc.PolicyUpdateRequest_ChanPoint{
			ChanPoint: &lnrpc.ChannelPoint{
				FundingTxid: &lnrpc.ChannelPoint_FundingTxidStr{
					FundingTxidStr: newTxid,
				},
				OutputIndex: newIndex,
			},
		}

		return r, nil
	}
}

func handleUpdatePolicyResponse(db firewalldb.PrivacyMapDB) func(
	ctx context.Context, r *lnrpc.PolicyUpdateResponse) (proto.Message,
	error) {

	return func(ctx context.Context, r *lnrpc.PolicyUpdateResponse) (
		proto.Message, error) {

		failedUpdates := make(
			[]*lnrpc.FailedUpdate, len(r.FailedUpdates),
		)

		err := db.Update(func(tx firewalldb.PrivacyMapTx) error {
			for i, u := range r.FailedUpdates {
				failedUpdates[i] = &lnrpc.FailedUpdate{
					Reason:      u.Reason,
					UpdateError: u.UpdateError,
				}

				if u.Outpoint == nil {
					continue
				}

				txid, index, err := firewalldb.HideChanPoint(
					tx, u.Outpoint.TxidStr,
					u.Outpoint.OutputIndex,
				)
				if err != nil {
					return err
				}

				failedUpdates[i].Outpoint = &lnrpc.OutPoint{
					TxidBytes:   nil,
					TxidStr:     txid,
					OutputIndex: index,
				}
			}

			return nil
		})
		if err != nil {
			return nil, err
		}

		return &lnrpc.PolicyUpdateResponse{
			FailedUpdates: failedUpdates,
		}, nil
	}
}

// hideAmount symmetrically randomizes an amount around a given relative
// variation interval. relativeVariation should be between 0 and 1.
func hideAmount(randIntn func(n int) (int, error), relativeVariation float64,
	amount uint64) (uint64, error) {

	if relativeVariation < 0 || relativeVariation > 1 {
		return 0, fmt.Errorf("hide amount: relative variation is not "+
			"between allowed bounds of [0, 1], is %v",
			relativeVariation)
	}

	if amount == 0 {
		return 0, nil
	}

	// fuzzInterval is smaller than the amount provided fuzzVariation is
	// between 0 and 1.
	fuzzInterval := uint64(float64(amount) * relativeVariation)

	amountMin := int(amount - fuzzInterval)
	amountMax := int(amount + fuzzInterval)

	randAmount, err := randBetween(randIntn, amountMin, amountMax)
	if err != nil {
		return 0, err
	}

	return uint64(randAmount), nil
}

// hideTimestamp symmetrically randomizes a unix timestamp given an absolute
// variation interval. The random input is expected to be rand.Intn.
func hideTimestamp(randIntn func(n int) (int, error),
	absoluteVariation time.Duration,
	timestamp time.Time) (time.Time, error) {

	if absoluteVariation < minTimeVariation ||
		absoluteVariation > maxTimeVariation {

		return time.Time{}, fmt.Errorf("hide timestamp: absolute time "+
			"variation is out of bounds, have %v",
			absoluteVariation)
	}

	// Don't fuzz meaningless timestamps.
	if timestamp.Add(-absoluteVariation).Unix() < 0 ||
		timestamp.IsZero() {

		return timestamp, nil
	}

	// We vary symmetrically around the provided timestamp.
	timeMin := timestamp.Add(-absoluteVariation)
	timeMax := timestamp.Add(absoluteVariation)

	timeNs, err := randBetween(
		randIntn, int(timeMin.UnixNano()), int(timeMax.UnixNano()),
	)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, int64(timeNs)), nil
}

// randBetween generates a random number between [min, max) given a source of
// randomness.
func randBetween(randIntn func(int) (int, error), min, max int) (int, error) {
	if max < min {
		return 0, fmt.Errorf("min is not allowed to be greater than "+
			"max, (min: %v, max: %v)", min, max)
	}

	// We don't want to pass zero to randIntn to avoid panics.
	if max == min {
		return min, nil
	}

	add, err := randIntn(max - min)
	if err != nil {
		return 0, err
	}

	return min + add, nil
}

// hideBool generates a random bool given a random input.
func hideBool(randIntn func(n int) (int, error)) (bool, error) {
	random, err := randIntn(2)
	if err != nil {
		return false, err
	}

	// For testing we may expect larger random numbers, which we map to
	// true.
	return random >= 1, nil
}

// CryptoRandIntn generates a random number between [0, n).
func CryptoRandIntn(n int) (int, error) {
	if n == 0 {
		return 0, nil
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}

	return int(nBig.Int64()), nil
}

// ObfuscateConfig alters the config string by replacing sensitive data with
// random values and returns new replacement pairs. We only substitute items in
// strings, numbers are left unchanged.
func ObfuscateConfig(db firewalldb.PrivacyMapReader, configB []byte) ([]byte,
	map[string]string, error) {

	if len(configB) == 0 {
		return nil, nil, nil
	}

	// We assume that the config is a json dict.
	var configMap map[string]any
	err := json.Unmarshal(configB, &configMap)
	if err != nil {
		return nil, nil, err
	}

	privMapPairs := make(map[string]string)
	newConfigMap := make(map[string]any)
	for k, v := range configMap {
		// We only substitute items in lists.
		list, ok := v.([]any)
		if !ok {
			newConfigMap[k] = v
			continue
		}

		// We only substitute items in lists of strings.
		stringList := make([]string, len(list))
		anyString := false
		allStrings := true
		for i, item := range list {
			item, ok := item.(string)
			allStrings = allStrings && ok
			anyString = anyString || ok

			if !ok {
				continue
			}

			stringList[i] = item
		}
		if anyString && !allStrings {
			return nil, nil, fmt.Errorf("invalid config, "+
				"expected list of only strings for key %s", k)
		} else if !anyString {
			newConfigMap[k] = v
			continue
		}

		obfuscatedValues := make([]string, len(stringList))
		for i, value := range stringList {
			value := strings.TrimSpace(value)

			// We first check if we have a mapping for this value
			// already.
			obfVal, haveValue := db.GetPseudo(value)
			if haveValue {
				obfuscatedValues[i] = obfVal

				continue
			}

			// We check if we have obfuscated this value already in
			// this run.
			obfVal, haveValue = privMapPairs[value]
			if haveValue {
				obfuscatedValues[i] = obfVal

				continue
			}

			// From here on we create new obfuscated values.
			// Try to replace with a chan point.
			_, _, err := firewalldb.DecodeChannelPoint(value)
			if err == nil {
				obfVal, err = firewalldb.NewPseudoChanPoint()
				if err != nil {
					return nil, nil, err
				}

				obfuscatedValues[i] = obfVal
				privMapPairs[value] = obfVal

				continue
			}

			// If the value is a pubkey, replace it with a random
			// value.
			_, err = hex.DecodeString(value)
			if err == nil && len(value) == pubKeyLen {
				obfVal, err := firewalldb.NewPseudoStr(
					len(value),
				)
				if err != nil {
					return nil, nil, err
				}

				obfuscatedValues[i] = obfVal
				privMapPairs[value] = obfVal

				continue
			}

			// If the value is a channel id, replace it with
			// a random value.
			_, err = strconv.ParseInt(value, 10, 64)
			length := len(value)

			// Channel ids can have different lengths depending on
			// the blockheight, 20 is equivalent to 10E9 blocks.
			if err == nil && minChanIDLen <= length &&
				length <= maxChanIDLen {

				obfVal, err := firewalldb.NewPseudoStr(length)
				if err != nil {
					return nil, nil, err
				}

				obfuscatedValues[i] = obfVal
				privMapPairs[value] = obfVal

				continue
			}

			// If we don't have a replacement for this value, we
			// just leave it as is.
			obfuscatedValues[i] = value
		}

		newConfigMap[k] = obfuscatedValues
	}

	// Marshal the map back into a JSON blob.
	newConfigB, err := json.Marshal(newConfigMap)
	if err != nil {
		return nil, nil, err
	}

	return newConfigB, privMapPairs, nil
}
