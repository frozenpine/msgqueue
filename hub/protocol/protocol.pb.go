// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: protocol.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResumeType int32

const (
	ResumeType_Restart ResumeType = 0
	ResumeType_Resume  ResumeType = 1
	ResumeType_Quick   ResumeType = 2
)

// Enum value maps for ResumeType.
var (
	ResumeType_name = map[int32]string{
		0: "Restart",
		1: "Resume",
		2: "Quick",
	}
	ResumeType_value = map[string]int32{
		"Restart": 0,
		"Resume":  1,
		"Quick":   2,
	}
)

func (x ResumeType) Enum() *ResumeType {
	p := new(ResumeType)
	*p = x
	return p
}

func (x ResumeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResumeType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_proto_enumTypes[0].Descriptor()
}

func (ResumeType) Type() protoreflect.EnumType {
	return &file_protocol_proto_enumTypes[0]
}

func (x ResumeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResumeType.Descriptor instead.
func (ResumeType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

type Topics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Define map[string]string `protobuf:"bytes,1,rep,name=define,proto3" json:"define,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Topics) Reset() {
	*x = Topics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topics) ProtoMessage() {}

func (x *Topics) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topics.ProtoReflect.Descriptor instead.
func (*Topics) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Topics) GetDefine() map[string]string {
	if x != nil {
		return x.Define
	}
	return nil
}

type RspInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorId  int32  `protobuf:"zigzag32,1,opt,name=error_id,json=errorId,proto3" json:"error_id,omitempty"`
	ErrorMsg string `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
}

func (x *RspInfo) Reset() {
	*x = RspInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspInfo) ProtoMessage() {}

func (x *RspInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspInfo.ProtoReflect.Descriptor instead.
func (*RspInfo) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *RspInfo) GetErrorId() int32 {
	if x != nil {
		return x.ErrorId
	}
	return 0
}

func (x *RspInfo) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type ReqSub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic      string     `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Subscriber string     `protobuf:"bytes,2,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
	ResumeType ResumeType `protobuf:"varint,3,opt,name=resume_type,json=resumeType,proto3,enum=protocol.ResumeType" json:"resume_type,omitempty"`
}

func (x *ReqSub) Reset() {
	*x = ReqSub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSub) ProtoMessage() {}

func (x *ReqSub) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSub.ProtoReflect.Descriptor instead.
func (*ReqSub) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *ReqSub) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ReqSub) GetSubscriber() string {
	if x != nil {
		return x.Subscriber
	}
	return ""
}

func (x *ReqSub) GetResumeType() ResumeType {
	if x != nil {
		return x.ResumeType
	}
	return ResumeType_Restart
}

type ReqUnSub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	SubId string `protobuf:"bytes,2,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
}

func (x *ReqUnSub) Reset() {
	*x = ReqUnSub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUnSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUnSub) ProtoMessage() {}

func (x *ReqUnSub) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUnSub.ProtoReflect.Descriptor instead.
func (*ReqUnSub) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *ReqUnSub) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ReqUnSub) GetSubId() string {
	if x != nil {
		return x.SubId
	}
	return ""
}

type RtnData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Seq   uint32 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Len   uint32 `protobuf:"varint,3,opt,name=len,proto3" json:"len,omitempty"`
	Data  []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RtnData) Reset() {
	*x = RtnData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtnData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtnData) ProtoMessage() {}

func (x *RtnData) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtnData.ProtoReflect.Descriptor instead.
func (*RtnData) Descriptor() ([]byte, []int) {
	return file_protocol_proto_rawDescGZIP(), []int{4}
}

func (x *RtnData) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *RtnData) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *RtnData) GetLen() uint32 {
	if x != nil {
		return x.Len
	}
	return 0
}

func (x *RtnData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protocol_proto protoreflect.FileDescriptor

var file_protocol_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x06, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x12, 0x34, 0x0a, 0x06, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x41, 0x0a, 0x07, 0x52, 0x73, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x75, 0x0a, 0x06, 0x52, 0x65, 0x71, 0x53, 0x75, 0x62, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x37, 0x0a, 0x08,
	0x52, 0x65, 0x71, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x15,
	0x0a, 0x06, 0x73, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x75, 0x62, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x07, 0x52, 0x74, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6c, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x30,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x75, 0x69, 0x63, 0x6b, 0x10, 0x02,
	0x32, 0xad, 0x01, 0x0a, 0x0a, 0x48, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x35, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52,
	0x65, 0x71, 0x53, 0x75, 0x62, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x52, 0x74, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x0b, 0x55, 0x6e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x1a, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x73, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_proto_rawDescOnce sync.Once
	file_protocol_proto_rawDescData = file_protocol_proto_rawDesc
)

func file_protocol_proto_rawDescGZIP() []byte {
	file_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_proto_rawDescData)
	})
	return file_protocol_proto_rawDescData
}

var file_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protocol_proto_goTypes = []interface{}{
	(ResumeType)(0),       // 0: protocol.ResumeType
	(*Topics)(nil),        // 1: protocol.Topics
	(*RspInfo)(nil),       // 2: protocol.RspInfo
	(*ReqSub)(nil),        // 3: protocol.ReqSub
	(*ReqUnSub)(nil),      // 4: protocol.ReqUnSub
	(*RtnData)(nil),       // 5: protocol.RtnData
	nil,                   // 6: protocol.Topics.DefineEntry
	(*emptypb.Empty)(nil), // 7: google.protobuf.Empty
}
var file_protocol_proto_depIdxs = []int32{
	6, // 0: protocol.Topics.define:type_name -> protocol.Topics.DefineEntry
	0, // 1: protocol.ReqSub.resume_type:type_name -> protocol.ResumeType
	7, // 2: protocol.HubService.GetTopics:input_type -> google.protobuf.Empty
	3, // 3: protocol.HubService.Subscribe:input_type -> protocol.ReqSub
	4, // 4: protocol.HubService.UnSubscribe:input_type -> protocol.ReqUnSub
	1, // 5: protocol.HubService.GetTopics:output_type -> protocol.Topics
	5, // 6: protocol.HubService.Subscribe:output_type -> protocol.RtnData
	2, // 7: protocol.HubService.UnSubscribe:output_type -> protocol.RspInfo
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_proto_init() }
func file_protocol_proto_init() {
	if File_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topics); i {
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
		file_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspInfo); i {
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
		file_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSub); i {
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
		file_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUnSub); i {
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
		file_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RtnData); i {
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
			RawDescriptor: file_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_proto_depIdxs,
		EnumInfos:         file_protocol_proto_enumTypes,
		MessageInfos:      file_protocol_proto_msgTypes,
	}.Build()
	File_protocol_proto = out.File
	file_protocol_proto_rawDesc = nil
	file_protocol_proto_goTypes = nil
	file_protocol_proto_depIdxs = nil
}