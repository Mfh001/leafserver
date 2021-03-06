// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: login.proto

package msg

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ResultCode int32

const (
	ResultCode_NoError           ResultCode = 0
	ResultCode_AgentNotFound     ResultCode = 101
	ResultCode_AgentAddressError ResultCode = 102
	ResultCode_GameNotFound      ResultCode = 103
)

// Enum value maps for ResultCode.
var (
	ResultCode_name = map[int32]string{
		0:   "NoError",
		101: "AgentNotFound",
		102: "AgentAddressError",
		103: "GameNotFound",
	}
	ResultCode_value = map[string]int32{
		"NoError":           0,
		"AgentNotFound":     101,
		"AgentAddressError": 102,
		"GameNotFound":      103,
	}
)

func (x ResultCode) Enum() *ResultCode {
	p := new(ResultCode)
	*p = x
	return p
}

func (x ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_login_proto_enumTypes[0].Descriptor()
}

func (ResultCode) Type() protoreflect.EnumType {
	return &file_login_proto_enumTypes[0]
}

func (x ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultCode.Descriptor instead.
func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

type LoginREQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  string `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Platform string `protobuf:"bytes,2,opt,name=Platform,proto3" json:"Platform,omitempty"`
	UID      string `protobuf:"bytes,3,opt,name=UID,proto3" json:"UID,omitempty"`
}

func (x *LoginREQ) Reset() {
	*x = LoginREQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginREQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginREQ) ProtoMessage() {}

func (x *LoginREQ) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginREQ.ProtoReflect.Descriptor instead.
func (*LoginREQ) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginREQ) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *LoginREQ) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *LoginREQ) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

type LoginACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    ResultCode  `protobuf:"varint,1,opt,name=Result,proto3,enum=msg.ResultCode" json:"Result,omitempty"`
	Server    *ServerInfo `protobuf:"bytes,2,opt,name=Server,proto3" json:"Server,omitempty"`
	GameToken string      `protobuf:"bytes,3,opt,name=GameToken,proto3" json:"GameToken,omitempty"`
	GameSvcID string      `protobuf:"bytes,4,opt,name=GameSvcID,proto3" json:"GameSvcID,omitempty"` // 选中的一台game服务器ID
}

func (x *LoginACK) Reset() {
	*x = LoginACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginACK) ProtoMessage() {}

func (x *LoginACK) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginACK.ProtoReflect.Descriptor instead.
func (*LoginACK) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginACK) GetResult() ResultCode {
	if x != nil {
		return x.Result
	}
	return ResultCode_NoError
}

func (x *LoginACK) GetServer() *ServerInfo {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *LoginACK) GetGameToken() string {
	if x != nil {
		return x.GameToken
	}
	return ""
}

func (x *LoginACK) GetGameSvcID() string {
	if x != nil {
		return x.GameSvcID
	}
	return ""
}

type PingREQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingREQ) Reset() {
	*x = PingREQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingREQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingREQ) ProtoMessage() {}

func (x *PingREQ) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingREQ.ProtoReflect.Descriptor instead.
func (*PingREQ) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

type PingACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingACK) Reset() {
	*x = PingACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingACK) ProtoMessage() {}

func (x *PingACK) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingACK.ProtoReflect.Descriptor instead.
func (*PingACK) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{3}
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP   string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{4}
}

func (x *ServerInfo) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *ServerInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type VerifyREQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameToken string `protobuf:"bytes,1,opt,name=GameToken,proto3" json:"GameToken,omitempty"`
	GameSvcID string `protobuf:"bytes,2,opt,name=GameSvcID,proto3" json:"GameSvcID,omitempty"`
}

func (x *VerifyREQ) Reset() {
	*x = VerifyREQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyREQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyREQ) ProtoMessage() {}

func (x *VerifyREQ) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyREQ.ProtoReflect.Descriptor instead.
func (*VerifyREQ) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyREQ) GetGameToken() string {
	if x != nil {
		return x.GameToken
	}
	return ""
}

func (x *VerifyREQ) GetGameSvcID() string {
	if x != nil {
		return x.GameSvcID
	}
	return ""
}

type VerifyACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result ResultCode `protobuf:"varint,1,opt,name=Result,proto3,enum=msg.ResultCode" json:"Result,omitempty"`
}

func (x *VerifyACK) Reset() {
	*x = VerifyACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyACK) ProtoMessage() {}

func (x *VerifyACK) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyACK.ProtoReflect.Descriptor instead.
func (*VerifyACK) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyACK) GetResult() ResultCode {
	if x != nil {
		return x.Result
	}
	return ResultCode_NoError
}

type ChatREQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *ChatREQ) Reset() {
	*x = ChatREQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatREQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatREQ) ProtoMessage() {}

func (x *ChatREQ) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatREQ.ProtoReflect.Descriptor instead.
func (*ChatREQ) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{7}
}

func (x *ChatREQ) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ChatACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *ChatACK) Reset() {
	*x = ChatACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatACK) ProtoMessage() {}

func (x *ChatACK) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatACK.ProtoReflect.Descriptor instead.
func (*ChatACK) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{8}
}

func (x *ChatACK) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type TestACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy string `protobuf:"bytes,1,opt,name=Dummy,proto3" json:"Dummy,omitempty"`
}

func (x *TestACK) Reset() {
	*x = TestACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestACK) ProtoMessage() {}

func (x *TestACK) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestACK.ProtoReflect.Descriptor instead.
func (*TestACK) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{9}
}

func (x *TestACK) GetDummy() string {
	if x != nil {
		return x.Dummy
	}
	return ""
}

type TestEchoACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg   string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Value int32  `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *TestEchoACK) Reset() {
	*x = TestEchoACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestEchoACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestEchoACK) ProtoMessage() {}

func (x *TestEchoACK) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestEchoACK.ProtoReflect.Descriptor instead.
func (*TestEchoACK) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{10}
}

func (x *TestEchoACK) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TestEchoACK) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x52, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x45, 0x51, 0x12, 0x18,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x22, 0x98, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x41, 0x43, 0x4b, 0x12, 0x27, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x06,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x76, 0x63, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x76, 0x63, 0x49,
	0x44, 0x22, 0x09, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x45, 0x51, 0x22, 0x09, 0x0a, 0x07,
	0x50, 0x69, 0x6e, 0x67, 0x41, 0x43, 0x4b, 0x22, 0x30, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x47, 0x0a, 0x09, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x52, 0x45, 0x51, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x76, 0x63, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x76, 0x63,
	0x49, 0x44, 0x22, 0x34, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x43, 0x4b, 0x12,
	0x27, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x45, 0x51, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x23, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x74, 0x41, 0x43, 0x4b, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x1f, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x41, 0x43, 0x4b, 0x12, 0x14, 0x0a,
	0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x44, 0x75,
	0x6d, 0x6d, 0x79, 0x22, 0x35, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x41,
	0x43, 0x4b, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x55, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x6f, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4e, 0x6f,
	0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x66, 0x12,
	0x10, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_login_proto_goTypes = []interface{}{
	(ResultCode)(0),     // 0: msg.ResultCode
	(*LoginREQ)(nil),    // 1: msg.LoginREQ
	(*LoginACK)(nil),    // 2: msg.LoginACK
	(*PingREQ)(nil),     // 3: msg.PingREQ
	(*PingACK)(nil),     // 4: msg.PingACK
	(*ServerInfo)(nil),  // 5: msg.ServerInfo
	(*VerifyREQ)(nil),   // 6: msg.VerifyREQ
	(*VerifyACK)(nil),   // 7: msg.VerifyACK
	(*ChatREQ)(nil),     // 8: msg.ChatREQ
	(*ChatACK)(nil),     // 9: msg.ChatACK
	(*TestACK)(nil),     // 10: msg.TestACK
	(*TestEchoACK)(nil), // 11: msg.TestEchoACK
}
var file_login_proto_depIdxs = []int32{
	0, // 0: msg.LoginACK.Result:type_name -> msg.ResultCode
	5, // 1: msg.LoginACK.Server:type_name -> msg.ServerInfo
	0, // 2: msg.VerifyACK.Result:type_name -> msg.ResultCode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginREQ); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginACK); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingREQ); i {
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
		file_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingACK); i {
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
		file_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyREQ); i {
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
		file_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyACK); i {
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
		file_login_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatREQ); i {
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
		file_login_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatACK); i {
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
		file_login_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestACK); i {
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
		file_login_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestEchoACK); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		EnumInfos:         file_login_proto_enumTypes,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
