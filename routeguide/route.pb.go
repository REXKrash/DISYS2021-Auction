// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: routeguide/route.proto

package routeguide

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

type AskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AskRequest) Reset() {
	*x = AskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskRequest) ProtoMessage() {}

func (x *AskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskRequest.ProtoReflect.Descriptor instead.
func (*AskRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{0}
}

func (x *AskRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type LeaderPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderPort int32 `protobuf:"varint,1,opt,name=leaderPort,proto3" json:"leaderPort,omitempty"`
}

func (x *LeaderPort) Reset() {
	*x = LeaderPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderPort) ProtoMessage() {}

func (x *LeaderPort) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderPort.ProtoReflect.Descriptor instead.
func (*LeaderPort) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{1}
}

func (x *LeaderPort) GetLeaderPort() int32 {
	if x != nil {
		return x.LeaderPort
	}
	return 0
}

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{2}
}

func (x *BidRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *BidRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	CurrentHighestBid int32 `protobuf:"varint,2,opt,name=currentHighestBid,proto3" json:"currentHighestBid,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{3}
}

func (x *BidResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BidResponse) GetCurrentHighestBid() int32 {
	if x != nil {
		return x.CurrentHighestBid
	}
	return 0
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{4}
}

func (x *ResultRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{5}
}

type AuctionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighestBid       int32  `protobuf:"varint,1,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	HighestBidder    string `protobuf:"bytes,2,opt,name=highestBidder,proto3" json:"highestBidder,omitempty"`
	IsAuctionRunning bool   `protobuf:"varint,3,opt,name=isAuctionRunning,proto3" json:"isAuctionRunning,omitempty"`
}

func (x *AuctionData) Reset() {
	*x = AuctionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionData) ProtoMessage() {}

func (x *AuctionData) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionData.ProtoReflect.Descriptor instead.
func (*AuctionData) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{6}
}

func (x *AuctionData) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *AuctionData) GetHighestBidder() string {
	if x != nil {
		return x.HighestBidder
	}
	return ""
}

func (x *AuctionData) GetIsAuctionRunning() bool {
	if x != nil {
		return x.IsAuctionRunning
	}
	return false
}

type LeadershipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentLeader int32 `protobuf:"varint,1,opt,name=currentLeader,proto3" json:"currentLeader,omitempty"`
}

func (x *LeadershipRequest) Reset() {
	*x = LeadershipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeadershipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeadershipRequest) ProtoMessage() {}

func (x *LeadershipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeadershipRequest.ProtoReflect.Descriptor instead.
func (*LeadershipRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{7}
}

func (x *LeadershipRequest) GetCurrentLeader() int32 {
	if x != nil {
		return x.CurrentLeader
	}
	return 0
}

type LeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentLeader int32 `protobuf:"varint,1,opt,name=currentLeader,proto3" json:"currentLeader,omitempty"`
	LowestPort    int32 `protobuf:"varint,2,opt,name=lowestPort,proto3" json:"lowestPort,omitempty"`
}

func (x *LeaderRequest) Reset() {
	*x = LeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderRequest) ProtoMessage() {}

func (x *LeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderRequest.ProtoReflect.Descriptor instead.
func (*LeaderRequest) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{8}
}

func (x *LeaderRequest) GetCurrentLeader() int32 {
	if x != nil {
		return x.CurrentLeader
	}
	return 0
}

func (x *LeaderRequest) GetLowestPort() int32 {
	if x != nil {
		return x.LowestPort
	}
	return 0
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeguide_route_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routeguide_route_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_routeguide_route_proto_rawDescGZIP(), []int{9}
}

func (x *HeartbeatResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_routeguide_route_proto protoreflect.FileDescriptor

var file_routeguide_route_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x0a, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x3c, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a,
	0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x7f, 0x0a,
	0x0b, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x73, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73,
	0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x39,
	0x0a, 0x11, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x0d, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x2b, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xcc, 0x01,
	0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x40, 0x0a, 0x0c, 0x41, 0x73, 0x6b, 0x46, 0x6f, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x41, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x42,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x32, 0xd3, 0x01, 0x0a,
	0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x11, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64,
	0x65, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routeguide_route_proto_rawDescOnce sync.Once
	file_routeguide_route_proto_rawDescData = file_routeguide_route_proto_rawDesc
)

func file_routeguide_route_proto_rawDescGZIP() []byte {
	file_routeguide_route_proto_rawDescOnce.Do(func() {
		file_routeguide_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_routeguide_route_proto_rawDescData)
	})
	return file_routeguide_route_proto_rawDescData
}

var file_routeguide_route_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_routeguide_route_proto_goTypes = []interface{}{
	(*AskRequest)(nil),        // 0: routeguide.AskRequest
	(*LeaderPort)(nil),        // 1: routeguide.LeaderPort
	(*BidRequest)(nil),        // 2: routeguide.BidRequest
	(*BidResponse)(nil),       // 3: routeguide.BidResponse
	(*ResultRequest)(nil),     // 4: routeguide.ResultRequest
	(*Empty)(nil),             // 5: routeguide.Empty
	(*AuctionData)(nil),       // 6: routeguide.AuctionData
	(*LeadershipRequest)(nil), // 7: routeguide.LeadershipRequest
	(*LeaderRequest)(nil),     // 8: routeguide.LeaderRequest
	(*HeartbeatResponse)(nil), // 9: routeguide.HeartbeatResponse
}
var file_routeguide_route_proto_depIdxs = []int32{
	0, // 0: routeguide.AuctionService.AskForLeader:input_type -> routeguide.AskRequest
	2, // 1: routeguide.AuctionService.Bid:input_type -> routeguide.BidRequest
	4, // 2: routeguide.AuctionService.Result:input_type -> routeguide.ResultRequest
	8, // 3: routeguide.TokenService.FindLeaderRequest:input_type -> routeguide.LeaderRequest
	5, // 4: routeguide.TokenService.SendHeartbeat:input_type -> routeguide.Empty
	6, // 5: routeguide.TokenService.ShareData:input_type -> routeguide.AuctionData
	1, // 6: routeguide.AuctionService.AskForLeader:output_type -> routeguide.LeaderPort
	3, // 7: routeguide.AuctionService.Bid:output_type -> routeguide.BidResponse
	6, // 8: routeguide.AuctionService.Result:output_type -> routeguide.AuctionData
	5, // 9: routeguide.TokenService.FindLeaderRequest:output_type -> routeguide.Empty
	9, // 10: routeguide.TokenService.SendHeartbeat:output_type -> routeguide.HeartbeatResponse
	5, // 11: routeguide.TokenService.ShareData:output_type -> routeguide.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_routeguide_route_proto_init() }
func file_routeguide_route_proto_init() {
	if File_routeguide_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routeguide_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskRequest); i {
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
		file_routeguide_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderPort); i {
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
		file_routeguide_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidRequest); i {
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
		file_routeguide_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidResponse); i {
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
		file_routeguide_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultRequest); i {
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
		file_routeguide_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_routeguide_route_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionData); i {
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
		file_routeguide_route_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeadershipRequest); i {
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
		file_routeguide_route_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderRequest); i {
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
		file_routeguide_route_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_routeguide_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_routeguide_route_proto_goTypes,
		DependencyIndexes: file_routeguide_route_proto_depIdxs,
		MessageInfos:      file_routeguide_route_proto_msgTypes,
	}.Build()
	File_routeguide_route_proto = out.File
	file_routeguide_route_proto_rawDesc = nil
	file_routeguide_route_proto_goTypes = nil
	file_routeguide_route_proto_depIdxs = nil
}
