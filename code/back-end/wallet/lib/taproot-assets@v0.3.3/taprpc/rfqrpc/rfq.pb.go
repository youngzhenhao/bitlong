// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: rfqrpc/rfq.proto

package rfqrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AssetSpecifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//
	//	*AssetSpecifier_AssetId
	//	*AssetSpecifier_AssetIdStr
	//	*AssetSpecifier_GroupKey
	//	*AssetSpecifier_GroupKeyStr
	Id isAssetSpecifier_Id `protobuf_oneof:"id"`
}

func (x *AssetSpecifier) Reset() {
	*x = AssetSpecifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetSpecifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSpecifier) ProtoMessage() {}

func (x *AssetSpecifier) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSpecifier.ProtoReflect.Descriptor instead.
func (*AssetSpecifier) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{0}
}

func (m *AssetSpecifier) GetId() isAssetSpecifier_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *AssetSpecifier) GetAssetId() []byte {
	if x, ok := x.GetId().(*AssetSpecifier_AssetId); ok {
		return x.AssetId
	}
	return nil
}

func (x *AssetSpecifier) GetAssetIdStr() string {
	if x, ok := x.GetId().(*AssetSpecifier_AssetIdStr); ok {
		return x.AssetIdStr
	}
	return ""
}

func (x *AssetSpecifier) GetGroupKey() []byte {
	if x, ok := x.GetId().(*AssetSpecifier_GroupKey); ok {
		return x.GroupKey
	}
	return nil
}

func (x *AssetSpecifier) GetGroupKeyStr() string {
	if x, ok := x.GetId().(*AssetSpecifier_GroupKeyStr); ok {
		return x.GroupKeyStr
	}
	return ""
}

type isAssetSpecifier_Id interface {
	isAssetSpecifier_Id()
}

type AssetSpecifier_AssetId struct {
	// The 32-byte asset ID specified as raw bytes (gRPC only).
	AssetId []byte `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3,oneof"`
}

type AssetSpecifier_AssetIdStr struct {
	// The 32-byte asset ID encoded as a hex string (use this for REST).
	AssetIdStr string `protobuf:"bytes,2,opt,name=asset_id_str,json=assetIdStr,proto3,oneof"`
}

type AssetSpecifier_GroupKey struct {
	// The 32-byte asset group key specified as raw bytes (gRPC only).
	GroupKey []byte `protobuf:"bytes,3,opt,name=group_key,json=groupKey,proto3,oneof"`
}

type AssetSpecifier_GroupKeyStr struct {
	// The 32-byte asset group key encoded as hex string (use this for
	// REST).
	GroupKeyStr string `protobuf:"bytes,4,opt,name=group_key_str,json=groupKeyStr,proto3,oneof"`
}

func (*AssetSpecifier_AssetId) isAssetSpecifier_Id() {}

func (*AssetSpecifier_AssetIdStr) isAssetSpecifier_Id() {}

func (*AssetSpecifier_GroupKey) isAssetSpecifier_Id() {}

func (*AssetSpecifier_GroupKeyStr) isAssetSpecifier_Id() {}

type AddAssetBuyOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// asset_specifier is the subject asset.
	AssetSpecifier *AssetSpecifier `protobuf:"bytes,1,opt,name=asset_specifier,json=assetSpecifier,proto3" json:"asset_specifier,omitempty"`
	// The minimum amount of the asset to buy.
	MinAssetAmount uint64 `protobuf:"varint,2,opt,name=min_asset_amount,json=minAssetAmount,proto3" json:"min_asset_amount,omitempty"`
	// The maximum amount BTC to spend (units: millisats).
	MaxBid uint64 `protobuf:"varint,3,opt,name=max_bid,json=maxBid,proto3" json:"max_bid,omitempty"`
	// The unix timestamp after which the order is no longer valid.
	Expiry uint64 `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// peer_pub_key is an optional field for specifying the public key of the
	// intended recipient peer for the order.
	PeerPubKey []byte `protobuf:"bytes,5,opt,name=peer_pub_key,json=peerPubKey,proto3" json:"peer_pub_key,omitempty"`
}

func (x *AddAssetBuyOrderRequest) Reset() {
	*x = AddAssetBuyOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAssetBuyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAssetBuyOrderRequest) ProtoMessage() {}

func (x *AddAssetBuyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAssetBuyOrderRequest.ProtoReflect.Descriptor instead.
func (*AddAssetBuyOrderRequest) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{1}
}

func (x *AddAssetBuyOrderRequest) GetAssetSpecifier() *AssetSpecifier {
	if x != nil {
		return x.AssetSpecifier
	}
	return nil
}

func (x *AddAssetBuyOrderRequest) GetMinAssetAmount() uint64 {
	if x != nil {
		return x.MinAssetAmount
	}
	return 0
}

func (x *AddAssetBuyOrderRequest) GetMaxBid() uint64 {
	if x != nil {
		return x.MaxBid
	}
	return 0
}

func (x *AddAssetBuyOrderRequest) GetExpiry() uint64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

func (x *AddAssetBuyOrderRequest) GetPeerPubKey() []byte {
	if x != nil {
		return x.PeerPubKey
	}
	return nil
}

type AddAssetBuyOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddAssetBuyOrderResponse) Reset() {
	*x = AddAssetBuyOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAssetBuyOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAssetBuyOrderResponse) ProtoMessage() {}

func (x *AddAssetBuyOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAssetBuyOrderResponse.ProtoReflect.Descriptor instead.
func (*AddAssetBuyOrderResponse) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{2}
}

type AddAssetSellOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// asset_specifier is the subject asset.
	AssetSpecifier *AssetSpecifier `protobuf:"bytes,1,opt,name=asset_specifier,json=assetSpecifier,proto3" json:"asset_specifier,omitempty"`
	// max_units is the maximum amount of the asset to sell.
	MaxUnits uint64 `protobuf:"varint,2,opt,name=max_units,json=maxUnits,proto3" json:"max_units,omitempty"`
}

func (x *AddAssetSellOfferRequest) Reset() {
	*x = AddAssetSellOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAssetSellOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAssetSellOfferRequest) ProtoMessage() {}

func (x *AddAssetSellOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAssetSellOfferRequest.ProtoReflect.Descriptor instead.
func (*AddAssetSellOfferRequest) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{3}
}

func (x *AddAssetSellOfferRequest) GetAssetSpecifier() *AssetSpecifier {
	if x != nil {
		return x.AssetSpecifier
	}
	return nil
}

func (x *AddAssetSellOfferRequest) GetMaxUnits() uint64 {
	if x != nil {
		return x.MaxUnits
	}
	return 0
}

type AddAssetSellOfferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddAssetSellOfferResponse) Reset() {
	*x = AddAssetSellOfferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAssetSellOfferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAssetSellOfferResponse) ProtoMessage() {}

func (x *AddAssetSellOfferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAssetSellOfferResponse.ProtoReflect.Descriptor instead.
func (*AddAssetSellOfferResponse) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{4}
}

type QueryRfqAcceptedQuotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryRfqAcceptedQuotesRequest) Reset() {
	*x = QueryRfqAcceptedQuotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRfqAcceptedQuotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRfqAcceptedQuotesRequest) ProtoMessage() {}

func (x *QueryRfqAcceptedQuotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRfqAcceptedQuotesRequest.ProtoReflect.Descriptor instead.
func (*QueryRfqAcceptedQuotesRequest) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{5}
}

type AcceptedQuote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Quote counterparty peer.
	Peer string `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	// The unique identifier of the quote request.
	Id []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// scid is the short channel ID of the channel over which the payment for
	// the quote should be made.
	Scid uint64 `protobuf:"varint,3,opt,name=scid,proto3" json:"scid,omitempty"`
	// asset_amount is the amount of the subject asset.
	AssetAmount uint64 `protobuf:"varint,4,opt,name=asset_amount,json=assetAmount,proto3" json:"asset_amount,omitempty"`
	// ask_price is the price in millisats for the entire asset amount.
	AskPrice uint64 `protobuf:"varint,5,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	// The unix timestamp after which the quote is no longer valid.
	Expiry uint64 `protobuf:"varint,6,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *AcceptedQuote) Reset() {
	*x = AcceptedQuote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptedQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptedQuote) ProtoMessage() {}

func (x *AcceptedQuote) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptedQuote.ProtoReflect.Descriptor instead.
func (*AcceptedQuote) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{6}
}

func (x *AcceptedQuote) GetPeer() string {
	if x != nil {
		return x.Peer
	}
	return ""
}

func (x *AcceptedQuote) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *AcceptedQuote) GetScid() uint64 {
	if x != nil {
		return x.Scid
	}
	return 0
}

func (x *AcceptedQuote) GetAssetAmount() uint64 {
	if x != nil {
		return x.AssetAmount
	}
	return 0
}

func (x *AcceptedQuote) GetAskPrice() uint64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *AcceptedQuote) GetExpiry() uint64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

type QueryRfqAcceptedQuotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcceptedQuotes []*AcceptedQuote `protobuf:"bytes,1,rep,name=accepted_quotes,json=acceptedQuotes,proto3" json:"accepted_quotes,omitempty"`
}

func (x *QueryRfqAcceptedQuotesResponse) Reset() {
	*x = QueryRfqAcceptedQuotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRfqAcceptedQuotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRfqAcceptedQuotesResponse) ProtoMessage() {}

func (x *QueryRfqAcceptedQuotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRfqAcceptedQuotesResponse.ProtoReflect.Descriptor instead.
func (*QueryRfqAcceptedQuotesResponse) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{7}
}

func (x *QueryRfqAcceptedQuotesResponse) GetAcceptedQuotes() []*AcceptedQuote {
	if x != nil {
		return x.AcceptedQuotes
	}
	return nil
}

type SubscribeRfqEventNtfnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeRfqEventNtfnsRequest) Reset() {
	*x = SubscribeRfqEventNtfnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRfqEventNtfnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRfqEventNtfnsRequest) ProtoMessage() {}

func (x *SubscribeRfqEventNtfnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRfqEventNtfnsRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRfqEventNtfnsRequest) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{8}
}

type IncomingAcceptQuoteEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp.
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The accepted quote.
	AcceptedQuote *AcceptedQuote `protobuf:"bytes,2,opt,name=accepted_quote,json=acceptedQuote,proto3" json:"accepted_quote,omitempty"`
}

func (x *IncomingAcceptQuoteEvent) Reset() {
	*x = IncomingAcceptQuoteEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomingAcceptQuoteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingAcceptQuoteEvent) ProtoMessage() {}

func (x *IncomingAcceptQuoteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingAcceptQuoteEvent.ProtoReflect.Descriptor instead.
func (*IncomingAcceptQuoteEvent) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{9}
}

func (x *IncomingAcceptQuoteEvent) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *IncomingAcceptQuoteEvent) GetAcceptedQuote() *AcceptedQuote {
	if x != nil {
		return x.AcceptedQuote
	}
	return nil
}

type AcceptHtlcEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp.
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// scid is the short channel ID of the channel over which the payment for
	// the quote is made.
	Scid uint64 `protobuf:"varint,2,opt,name=scid,proto3" json:"scid,omitempty"`
}

func (x *AcceptHtlcEvent) Reset() {
	*x = AcceptHtlcEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptHtlcEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptHtlcEvent) ProtoMessage() {}

func (x *AcceptHtlcEvent) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptHtlcEvent.ProtoReflect.Descriptor instead.
func (*AcceptHtlcEvent) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{10}
}

func (x *AcceptHtlcEvent) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *AcceptHtlcEvent) GetScid() uint64 {
	if x != nil {
		return x.Scid
	}
	return 0
}

type RfqEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*RfqEvent_IncomingAcceptQuote
	//	*RfqEvent_AcceptHtlc
	Event isRfqEvent_Event `protobuf_oneof:"event"`
}

func (x *RfqEvent) Reset() {
	*x = RfqEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rfqrpc_rfq_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RfqEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RfqEvent) ProtoMessage() {}

func (x *RfqEvent) ProtoReflect() protoreflect.Message {
	mi := &file_rfqrpc_rfq_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RfqEvent.ProtoReflect.Descriptor instead.
func (*RfqEvent) Descriptor() ([]byte, []int) {
	return file_rfqrpc_rfq_proto_rawDescGZIP(), []int{11}
}

func (m *RfqEvent) GetEvent() isRfqEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *RfqEvent) GetIncomingAcceptQuote() *IncomingAcceptQuoteEvent {
	if x, ok := x.GetEvent().(*RfqEvent_IncomingAcceptQuote); ok {
		return x.IncomingAcceptQuote
	}
	return nil
}

func (x *RfqEvent) GetAcceptHtlc() *AcceptHtlcEvent {
	if x, ok := x.GetEvent().(*RfqEvent_AcceptHtlc); ok {
		return x.AcceptHtlc
	}
	return nil
}

type isRfqEvent_Event interface {
	isRfqEvent_Event()
}

type RfqEvent_IncomingAcceptQuote struct {
	// incoming_accept_quote is an event that is sent when an incoming
	// accept quote message is received.
	IncomingAcceptQuote *IncomingAcceptQuoteEvent `protobuf:"bytes,1,opt,name=incoming_accept_quote,json=incomingAcceptQuote,proto3,oneof"`
}

type RfqEvent_AcceptHtlc struct {
	// accept_htlc is an event that is sent when a HTLC is accepted by the
	// RFQ service.
	AcceptHtlc *AcceptHtlcEvent `protobuf:"bytes,2,opt,name=accept_htlc,json=acceptHtlc,proto3,oneof"`
}

func (*RfqEvent_IncomingAcceptQuote) isRfqEvent_Event() {}

func (*RfqEvent_AcceptHtlc) isRfqEvent_Event() {}

var File_rfqrpc_rfq_proto protoreflect.FileDescriptor

var file_rfqrpc_rfq_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x66, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x53, 0x74, 0x72, 0x12, 0x1d,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a,
	0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79,
	0x53, 0x74, 0x72, 0x42, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x17, 0x41, 0x64,
	0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x50, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42,
	0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x78, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x6c, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x41, 0x64, 0x64,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x66, 0x71, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x63, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0x60, 0x0a, 0x1e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x66, 0x71, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x0e, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x66, 0x71, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x74, 0x66, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x76, 0x0a, 0x18,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3c, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x65, 0x64, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x22, 0x43, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x48, 0x74,
	0x6c, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x63, 0x69, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x08, 0x52, 0x66,
	0x71, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a, 0x15, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x13, 0x69, 0x6e, 0x63, 0x6f, 0x6d,
	0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x3a,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x68, 0x74, 0x6c, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x48, 0x74, 0x6c, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x48, 0x74, 0x6c, 0x63, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x32, 0xf4, 0x02, 0x0a, 0x03, 0x52, 0x66, 0x71, 0x12, 0x55, 0x0a, 0x10, 0x41,
	0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1f, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65,
	0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63,
	0x2e, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x66, 0x71, 0x72,
	0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x6c, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x16,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x66, 0x71, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x66, 0x71, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x66, 0x71, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x66, 0x71, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x74, 0x66, 0x6e, 0x73, 0x12,
	0x25, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x66, 0x71, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x74, 0x66, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x66, 0x71, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x66, 0x71, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69,
	0x6e, 0x67, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x6f, 0x74, 0x2d, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x66, 0x71,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rfqrpc_rfq_proto_rawDescOnce sync.Once
	file_rfqrpc_rfq_proto_rawDescData = file_rfqrpc_rfq_proto_rawDesc
)

func file_rfqrpc_rfq_proto_rawDescGZIP() []byte {
	file_rfqrpc_rfq_proto_rawDescOnce.Do(func() {
		file_rfqrpc_rfq_proto_rawDescData = protoimpl.X.CompressGZIP(file_rfqrpc_rfq_proto_rawDescData)
	})
	return file_rfqrpc_rfq_proto_rawDescData
}

var file_rfqrpc_rfq_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_rfqrpc_rfq_proto_goTypes = []interface{}{
	(*AssetSpecifier)(nil),                 // 0: rfqrpc.AssetSpecifier
	(*AddAssetBuyOrderRequest)(nil),        // 1: rfqrpc.AddAssetBuyOrderRequest
	(*AddAssetBuyOrderResponse)(nil),       // 2: rfqrpc.AddAssetBuyOrderResponse
	(*AddAssetSellOfferRequest)(nil),       // 3: rfqrpc.AddAssetSellOfferRequest
	(*AddAssetSellOfferResponse)(nil),      // 4: rfqrpc.AddAssetSellOfferResponse
	(*QueryRfqAcceptedQuotesRequest)(nil),  // 5: rfqrpc.QueryRfqAcceptedQuotesRequest
	(*AcceptedQuote)(nil),                  // 6: rfqrpc.AcceptedQuote
	(*QueryRfqAcceptedQuotesResponse)(nil), // 7: rfqrpc.QueryRfqAcceptedQuotesResponse
	(*SubscribeRfqEventNtfnsRequest)(nil),  // 8: rfqrpc.SubscribeRfqEventNtfnsRequest
	(*IncomingAcceptQuoteEvent)(nil),       // 9: rfqrpc.IncomingAcceptQuoteEvent
	(*AcceptHtlcEvent)(nil),                // 10: rfqrpc.AcceptHtlcEvent
	(*RfqEvent)(nil),                       // 11: rfqrpc.RfqEvent
}
var file_rfqrpc_rfq_proto_depIdxs = []int32{
	0,  // 0: rfqrpc.AddAssetBuyOrderRequest.asset_specifier:type_name -> rfqrpc.AssetSpecifier
	0,  // 1: rfqrpc.AddAssetSellOfferRequest.asset_specifier:type_name -> rfqrpc.AssetSpecifier
	6,  // 2: rfqrpc.QueryRfqAcceptedQuotesResponse.accepted_quotes:type_name -> rfqrpc.AcceptedQuote
	6,  // 3: rfqrpc.IncomingAcceptQuoteEvent.accepted_quote:type_name -> rfqrpc.AcceptedQuote
	9,  // 4: rfqrpc.RfqEvent.incoming_accept_quote:type_name -> rfqrpc.IncomingAcceptQuoteEvent
	10, // 5: rfqrpc.RfqEvent.accept_htlc:type_name -> rfqrpc.AcceptHtlcEvent
	1,  // 6: rfqrpc.Rfq.AddAssetBuyOrder:input_type -> rfqrpc.AddAssetBuyOrderRequest
	3,  // 7: rfqrpc.Rfq.AddAssetSellOffer:input_type -> rfqrpc.AddAssetSellOfferRequest
	5,  // 8: rfqrpc.Rfq.QueryRfqAcceptedQuotes:input_type -> rfqrpc.QueryRfqAcceptedQuotesRequest
	8,  // 9: rfqrpc.Rfq.SubscribeRfqEventNtfns:input_type -> rfqrpc.SubscribeRfqEventNtfnsRequest
	2,  // 10: rfqrpc.Rfq.AddAssetBuyOrder:output_type -> rfqrpc.AddAssetBuyOrderResponse
	4,  // 11: rfqrpc.Rfq.AddAssetSellOffer:output_type -> rfqrpc.AddAssetSellOfferResponse
	7,  // 12: rfqrpc.Rfq.QueryRfqAcceptedQuotes:output_type -> rfqrpc.QueryRfqAcceptedQuotesResponse
	11, // 13: rfqrpc.Rfq.SubscribeRfqEventNtfns:output_type -> rfqrpc.RfqEvent
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_rfqrpc_rfq_proto_init() }
func file_rfqrpc_rfq_proto_init() {
	if File_rfqrpc_rfq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rfqrpc_rfq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetSpecifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAssetBuyOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAssetBuyOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAssetSellOfferRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAssetSellOfferResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRfqAcceptedQuotesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptedQuote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRfqAcceptedQuotesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRfqEventNtfnsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomingAcceptQuoteEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptHtlcEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rfqrpc_rfq_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RfqEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_rfqrpc_rfq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AssetSpecifier_AssetId)(nil),
		(*AssetSpecifier_AssetIdStr)(nil),
		(*AssetSpecifier_GroupKey)(nil),
		(*AssetSpecifier_GroupKeyStr)(nil),
	}
	file_rfqrpc_rfq_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*RfqEvent_IncomingAcceptQuote)(nil),
		(*RfqEvent_AcceptHtlc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rfqrpc_rfq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rfqrpc_rfq_proto_goTypes,
		DependencyIndexes: file_rfqrpc_rfq_proto_depIdxs,
		MessageInfos:      file_rfqrpc_rfq_proto_msgTypes,
	}.Build()
	File_rfqrpc_rfq_proto = out.File
	file_rfqrpc_rfq_proto_rawDesc = nil
	file_rfqrpc_rfq_proto_goTypes = nil
	file_rfqrpc_rfq_proto_depIdxs = nil
}
