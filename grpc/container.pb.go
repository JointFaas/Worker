// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.7.1
// source: container.proto

package grpc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type InvokeReply_Code int32

const (
	InvokeReply_OK    InvokeReply_Code = 0
	InvokeReply_ERROR InvokeReply_Code = 1
)

// Enum value maps for InvokeReply_Code.
var (
	InvokeReply_Code_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	InvokeReply_Code_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x InvokeReply_Code) Enum() *InvokeReply_Code {
	p := new(InvokeReply_Code)
	*p = x
	return p
}

func (x InvokeReply_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvokeReply_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_container_proto_enumTypes[0].Descriptor()
}

func (InvokeReply_Code) Type() protoreflect.EnumType {
	return &file_container_proto_enumTypes[0]
}

func (x InvokeReply_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvokeReply_Code.Descriptor instead.
func (InvokeReply_Code) EnumDescriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{3, 0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*ConnectRequest_Register
	//	*ConnectRequest_InvokeRequest
	//	*ConnectRequest_InvokeReply
	Action isConnectRequest_Action `protobuf_oneof:"Action"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_container_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{0}
}

func (m *ConnectRequest) GetAction() isConnectRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *ConnectRequest) GetRegister() *RegisterRequest {
	if x, ok := x.GetAction().(*ConnectRequest_Register); ok {
		return x.Register
	}
	return nil
}

func (x *ConnectRequest) GetInvokeRequest() *InvokeRequest {
	if x, ok := x.GetAction().(*ConnectRequest_InvokeRequest); ok {
		return x.InvokeRequest
	}
	return nil
}

func (x *ConnectRequest) GetInvokeReply() *InvokeReply {
	if x, ok := x.GetAction().(*ConnectRequest_InvokeReply); ok {
		return x.InvokeReply
	}
	return nil
}

type isConnectRequest_Action interface {
	isConnectRequest_Action()
}

type ConnectRequest_Register struct {
	Register *RegisterRequest `protobuf:"bytes,1,opt,name=register,proto3,oneof"`
}

type ConnectRequest_InvokeRequest struct {
	InvokeRequest *InvokeRequest `protobuf:"bytes,2,opt,name=InvokeRequest,proto3,oneof"`
}

type ConnectRequest_InvokeReply struct {
	InvokeReply *InvokeReply `protobuf:"bytes,3,opt,name=InvokeReply,proto3,oneof"`
}

func (*ConnectRequest_Register) isConnectRequest_Action() {}

func (*ConnectRequest_InvokeRequest) isConnectRequest_Action() {}

func (*ConnectRequest_InvokeReply) isConnectRequest_Action() {}

type ConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *ConnectReply) Reset() {
	*x = ConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReply) ProtoMessage() {}

func (x *ConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_container_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReply.ProtoReflect.Descriptor instead.
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConnectReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type InvokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InvokeRequest) Reset() {
	*x = InvokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeRequest) ProtoMessage() {}

func (x *InvokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_container_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeRequest.ProtoReflect.Descriptor instead.
func (*InvokeRequest) Descriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{2}
}

func (x *InvokeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InvokeRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type InvokeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Output []byte `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *InvokeReply) Reset() {
	*x = InvokeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeReply) ProtoMessage() {}

func (x *InvokeReply) ProtoReflect() protoreflect.Message {
	mi := &file_container_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeReply.ProtoReflect.Descriptor instead.
func (*InvokeReply) Descriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{3}
}

func (x *InvokeReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InvokeReply) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FuncName string `protobuf:"bytes,1,opt,name=funcName,proto3" json:"funcName,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_container_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterRequest) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

var File_container_proto protoreflect.FileDescriptor

var file_container_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0xc3, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x3b, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0b,
	0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x34, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x22, 0x39, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x50,
	0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x19, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01,
	0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0x4c, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x12, 0x39, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_container_proto_rawDescOnce sync.Once
	file_container_proto_rawDescData = file_container_proto_rawDesc
)

func file_container_proto_rawDescGZIP() []byte {
	file_container_proto_rawDescOnce.Do(func() {
		file_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_container_proto_rawDescData)
	})
	return file_container_proto_rawDescData
}

var file_container_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_container_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_container_proto_goTypes = []interface{}{
	(InvokeReply_Code)(0),   // 0: grpc.InvokeReply.Code
	(*ConnectRequest)(nil),  // 1: grpc.ConnectRequest
	(*ConnectReply)(nil),    // 2: grpc.ConnectReply
	(*InvokeRequest)(nil),   // 3: grpc.InvokeRequest
	(*InvokeReply)(nil),     // 4: grpc.InvokeReply
	(*RegisterRequest)(nil), // 5: grpc.RegisterRequest
}
var file_container_proto_depIdxs = []int32{
	5, // 0: grpc.ConnectRequest.register:type_name -> grpc.RegisterRequest
	3, // 1: grpc.ConnectRequest.InvokeRequest:type_name -> grpc.InvokeRequest
	4, // 2: grpc.ConnectRequest.InvokeReply:type_name -> grpc.InvokeReply
	1, // 3: grpc.WorkerContainer.Connect:input_type -> grpc.ConnectRequest
	2, // 4: grpc.WorkerContainer.Connect:output_type -> grpc.ConnectReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_container_proto_init() }
func file_container_proto_init() {
	if File_container_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_container_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_container_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReply); i {
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
		file_container_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeRequest); i {
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
		file_container_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeReply); i {
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
		file_container_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
	file_container_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ConnectRequest_Register)(nil),
		(*ConnectRequest_InvokeRequest)(nil),
		(*ConnectRequest_InvokeReply)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_container_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_container_proto_goTypes,
		DependencyIndexes: file_container_proto_depIdxs,
		EnumInfos:         file_container_proto_enumTypes,
		MessageInfos:      file_container_proto_msgTypes,
	}.Build()
	File_container_proto = out.File
	file_container_proto_rawDesc = nil
	file_container_proto_goTypes = nil
	file_container_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkerContainerClient is the client API for WorkerContainer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerContainerClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (WorkerContainer_ConnectClient, error)
}

type workerContainerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerContainerClient(cc grpc.ClientConnInterface) WorkerContainerClient {
	return &workerContainerClient{cc}
}

func (c *workerContainerClient) Connect(ctx context.Context, opts ...grpc.CallOption) (WorkerContainer_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WorkerContainer_serviceDesc.Streams[0], "/grpc.WorkerContainer/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerContainerConnectClient{stream}
	return x, nil
}

type WorkerContainer_ConnectClient interface {
	Send(*ConnectRequest) error
	Recv() (*ConnectReply, error)
	grpc.ClientStream
}

type workerContainerConnectClient struct {
	grpc.ClientStream
}

func (x *workerContainerConnectClient) Send(m *ConnectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerContainerConnectClient) Recv() (*ConnectReply, error) {
	m := new(ConnectReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WorkerContainerServer is the server API for WorkerContainer service.
type WorkerContainerServer interface {
	Connect(WorkerContainer_ConnectServer) error
}

// UnimplementedWorkerContainerServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerContainerServer struct {
}

func (*UnimplementedWorkerContainerServer) Connect(WorkerContainer_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterWorkerContainerServer(s *grpc.Server, srv WorkerContainerServer) {
	s.RegisterService(&_WorkerContainer_serviceDesc, srv)
}

func _WorkerContainer_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerContainerServer).Connect(&workerContainerConnectServer{stream})
}

type WorkerContainer_ConnectServer interface {
	Send(*ConnectReply) error
	Recv() (*ConnectRequest, error)
	grpc.ServerStream
}

type workerContainerConnectServer struct {
	grpc.ServerStream
}

func (x *workerContainerConnectServer) Send(m *ConnectReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerContainerConnectServer) Recv() (*ConnectRequest, error) {
	m := new(ConnectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _WorkerContainer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.WorkerContainer",
	HandlerType: (*WorkerContainerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _WorkerContainer_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "container.proto",
}
