// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.7.1
// source: container/container.proto

package container

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

type InvokeResponse_Code int32

const (
	InvokeResponse_OK            InvokeResponse_Code = 0
	InvokeResponse_NOT_READY     InvokeResponse_Code = 1
	InvokeResponse_FUNC_MISMATCH InvokeResponse_Code = 2
	InvokeResponse_RUNTIME_ERROR InvokeResponse_Code = 3
)

// Enum value maps for InvokeResponse_Code.
var (
	InvokeResponse_Code_name = map[int32]string{
		0: "OK",
		1: "NOT_READY",
		2: "FUNC_MISMATCH",
		3: "RUNTIME_ERROR",
	}
	InvokeResponse_Code_value = map[string]int32{
		"OK":            0,
		"NOT_READY":     1,
		"FUNC_MISMATCH": 2,
		"RUNTIME_ERROR": 3,
	}
)

func (x InvokeResponse_Code) Enum() *InvokeResponse_Code {
	p := new(InvokeResponse_Code)
	*p = x
	return p
}

func (x InvokeResponse_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvokeResponse_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_container_container_proto_enumTypes[0].Descriptor()
}

func (InvokeResponse_Code) Type() protoreflect.EnumType {
	return &file_container_container_proto_enumTypes[0]
}

func (x InvokeResponse_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvokeResponse_Code.Descriptor instead.
func (InvokeResponse_Code) EnumDescriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{1, 0}
}

type SetEnvsResponse_Code int32

const (
	SetEnvsResponse_OK          SetEnvsResponse_Code = 0
	SetEnvsResponse_INVALID_ENV SetEnvsResponse_Code = 1
)

// Enum value maps for SetEnvsResponse_Code.
var (
	SetEnvsResponse_Code_name = map[int32]string{
		0: "OK",
		1: "INVALID_ENV",
	}
	SetEnvsResponse_Code_value = map[string]int32{
		"OK":          0,
		"INVALID_ENV": 1,
	}
)

func (x SetEnvsResponse_Code) Enum() *SetEnvsResponse_Code {
	p := new(SetEnvsResponse_Code)
	*p = x
	return p
}

func (x SetEnvsResponse_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetEnvsResponse_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_container_container_proto_enumTypes[1].Descriptor()
}

func (SetEnvsResponse_Code) Type() protoreflect.EnumType {
	return &file_container_container_proto_enumTypes[1]
}

func (x SetEnvsResponse_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetEnvsResponse_Code.Descriptor instead.
func (SetEnvsResponse_Code) EnumDescriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{3, 0}
}

type LoadCodeResponse_Code int32

const (
	LoadCodeResponse_OK    LoadCodeResponse_Code = 0
	LoadCodeResponse_ERROR LoadCodeResponse_Code = 1
)

// Enum value maps for LoadCodeResponse_Code.
var (
	LoadCodeResponse_Code_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	LoadCodeResponse_Code_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x LoadCodeResponse_Code) Enum() *LoadCodeResponse_Code {
	p := new(LoadCodeResponse_Code)
	*p = x
	return p
}

func (x LoadCodeResponse_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoadCodeResponse_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_container_container_proto_enumTypes[2].Descriptor()
}

func (LoadCodeResponse_Code) Type() protoreflect.EnumType {
	return &file_container_container_proto_enumTypes[2]
}

func (x LoadCodeResponse_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoadCodeResponse_Code.Descriptor instead.
func (LoadCodeResponse_Code) EnumDescriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{5, 0}
}

type StopResponse_Code int32

const (
	StopResponse_OK StopResponse_Code = 0
)

// Enum value maps for StopResponse_Code.
var (
	StopResponse_Code_name = map[int32]string{
		0: "OK",
	}
	StopResponse_Code_value = map[string]int32{
		"OK": 0,
	}
)

func (x StopResponse_Code) Enum() *StopResponse_Code {
	p := new(StopResponse_Code)
	*p = x
	return p
}

func (x StopResponse_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StopResponse_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_container_container_proto_enumTypes[3].Descriptor()
}

func (StopResponse_Code) Type() protoreflect.EnumType {
	return &file_container_container_proto_enumTypes[3]
}

func (x StopResponse_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StopResponse_Code.Descriptor instead.
func (StopResponse_Code) EnumDescriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{7, 0}
}

type InvokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FuncName string `protobuf:"bytes,1,opt,name=funcName,proto3" json:"funcName,omitempty"`
	Payload  []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InvokeRequest) Reset() {
	*x = InvokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeRequest) ProtoMessage() {}

func (x *InvokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_container_container_proto_msgTypes[0]
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
	return file_container_container_proto_rawDescGZIP(), []int{0}
}

func (x *InvokeRequest) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *InvokeRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type InvokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   InvokeResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=container.InvokeResponse_Code" json:"code,omitempty"`
	Output []byte              `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *InvokeResponse) Reset() {
	*x = InvokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_container_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeResponse) ProtoMessage() {}

func (x *InvokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_container_container_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeResponse.ProtoReflect.Descriptor instead.
func (*InvokeResponse) Descriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{1}
}

func (x *InvokeResponse) GetCode() InvokeResponse_Code {
	if x != nil {
		return x.Code
	}
	return InvokeResponse_OK
}

func (x *InvokeResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

type SetEnvsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Env []string `protobuf:"bytes,1,rep,name=env,proto3" json:"env,omitempty"`
}

func (x *SetEnvsRequest) Reset() {
	*x = SetEnvsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_container_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEnvsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEnvsRequest) ProtoMessage() {}

func (x *SetEnvsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_container_container_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEnvsRequest.ProtoReflect.Descriptor instead.
func (*SetEnvsRequest) Descriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{2}
}

func (x *SetEnvsRequest) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

type SetEnvsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code SetEnvsResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=container.SetEnvsResponse_Code" json:"code,omitempty"`
}

func (x *SetEnvsResponse) Reset() {
	*x = SetEnvsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_container_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEnvsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEnvsResponse) ProtoMessage() {}

func (x *SetEnvsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_container_container_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEnvsResponse.ProtoReflect.Descriptor instead.
func (*SetEnvsResponse) Descriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{3}
}

func (x *SetEnvsResponse) GetCode() SetEnvsResponse_Code {
	if x != nil {
		return x.Code
	}
	return SetEnvsResponse_OK
}

type LoadCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FuncName string `protobuf:"bytes,1,opt,name=funcName,proto3" json:"funcName,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *LoadCodeRequest) Reset() {
	*x = LoadCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_container_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadCodeRequest) ProtoMessage() {}

func (x *LoadCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_container_container_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadCodeRequest.ProtoReflect.Descriptor instead.
func (*LoadCodeRequest) Descriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{4}
}

func (x *LoadCodeRequest) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *LoadCodeRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type LoadCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code LoadCodeResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=container.LoadCodeResponse_Code" json:"code,omitempty"`
}

func (x *LoadCodeResponse) Reset() {
	*x = LoadCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_container_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadCodeResponse) ProtoMessage() {}

func (x *LoadCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_container_container_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadCodeResponse.ProtoReflect.Descriptor instead.
func (*LoadCodeResponse) Descriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{5}
}

func (x *LoadCodeResponse) GetCode() LoadCodeResponse_Code {
	if x != nil {
		return x.Code
	}
	return LoadCodeResponse_OK
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_container_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_container_container_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{6}
}

type StopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code StopResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=container.StopResponse_Code" json:"code,omitempty"`
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_container_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_container_container_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_container_container_proto_rawDescGZIP(), []int{7}
}

func (x *StopResponse) GetCode() StopResponse_Code {
	if x != nil {
		return x.Code
	}
	return StopResponse_OK
}

var File_container_container_proto protoreflect.FileDescriptor

var file_container_container_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xa1, 0x01,
	0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x43, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46,
	0x55, 0x4e, 0x43, 0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x52, 0x55, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x22, 0x22, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0x67, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1f, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45, 0x4e, 0x56, 0x10, 0x01, 0x22, 0x3f,
	0x0a, 0x0f, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x63, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x19, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x01, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0e, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x4b, 0x10, 0x00, 0x32, 0x92, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x18, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x73, 0x12,
	0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x45,
	0x6e, 0x76, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x08, 0x4c, 0x6f, 0x61, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x61,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x0a, 0x13, 0x6a, 0x6f,
	0x69, 0x6e, 0x74, 0x66, 0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x42, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x2e, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0xa2, 0x02, 0x03, 0x48, 0x4c, 0x57, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_container_container_proto_rawDescOnce sync.Once
	file_container_container_proto_rawDescData = file_container_container_proto_rawDesc
)

func file_container_container_proto_rawDescGZIP() []byte {
	file_container_container_proto_rawDescOnce.Do(func() {
		file_container_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_container_container_proto_rawDescData)
	})
	return file_container_container_proto_rawDescData
}

var file_container_container_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_container_container_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_container_container_proto_goTypes = []interface{}{
	(InvokeResponse_Code)(0),   // 0: container.InvokeResponse.Code
	(SetEnvsResponse_Code)(0),  // 1: container.SetEnvsResponse.Code
	(LoadCodeResponse_Code)(0), // 2: container.LoadCodeResponse.Code
	(StopResponse_Code)(0),     // 3: container.StopResponse.Code
	(*InvokeRequest)(nil),      // 4: container.InvokeRequest
	(*InvokeResponse)(nil),     // 5: container.InvokeResponse
	(*SetEnvsRequest)(nil),     // 6: container.SetEnvsRequest
	(*SetEnvsResponse)(nil),    // 7: container.SetEnvsResponse
	(*LoadCodeRequest)(nil),    // 8: container.LoadCodeRequest
	(*LoadCodeResponse)(nil),   // 9: container.LoadCodeResponse
	(*StopRequest)(nil),        // 10: container.StopRequest
	(*StopResponse)(nil),       // 11: container.StopResponse
}
var file_container_container_proto_depIdxs = []int32{
	0,  // 0: container.InvokeResponse.code:type_name -> container.InvokeResponse.Code
	1,  // 1: container.SetEnvsResponse.code:type_name -> container.SetEnvsResponse.Code
	2,  // 2: container.LoadCodeResponse.code:type_name -> container.LoadCodeResponse.Code
	3,  // 3: container.StopResponse.code:type_name -> container.StopResponse.Code
	4,  // 4: container.Container.Invoke:input_type -> container.InvokeRequest
	6,  // 5: container.Container.SetEnvs:input_type -> container.SetEnvsRequest
	8,  // 6: container.Container.LoadCode:input_type -> container.LoadCodeRequest
	10, // 7: container.Container.Stop:input_type -> container.StopRequest
	5,  // 8: container.Container.Invoke:output_type -> container.InvokeResponse
	7,  // 9: container.Container.SetEnvs:output_type -> container.SetEnvsResponse
	9,  // 10: container.Container.LoadCode:output_type -> container.LoadCodeResponse
	11, // 11: container.Container.Stop:output_type -> container.StopResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_container_container_proto_init() }
func file_container_container_proto_init() {
	if File_container_container_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_container_container_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_container_container_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeResponse); i {
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
		file_container_container_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEnvsRequest); i {
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
		file_container_container_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEnvsResponse); i {
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
		file_container_container_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadCodeRequest); i {
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
		file_container_container_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadCodeResponse); i {
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
		file_container_container_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_container_container_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResponse); i {
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
			RawDescriptor: file_container_container_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_container_container_proto_goTypes,
		DependencyIndexes: file_container_container_proto_depIdxs,
		EnumInfos:         file_container_container_proto_enumTypes,
		MessageInfos:      file_container_container_proto_msgTypes,
	}.Build()
	File_container_container_proto = out.File
	file_container_container_proto_rawDesc = nil
	file_container_container_proto_goTypes = nil
	file_container_container_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ContainerClient is the client API for Container service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContainerClient interface {
	Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
	SetEnvs(ctx context.Context, in *SetEnvsRequest, opts ...grpc.CallOption) (*SetEnvsResponse, error)
	LoadCode(ctx context.Context, in *LoadCodeRequest, opts ...grpc.CallOption) (*LoadCodeResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type containerClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerClient(cc grpc.ClientConnInterface) ContainerClient {
	return &containerClient{cc}
}

func (c *containerClient) Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := c.cc.Invoke(ctx, "/container.Container/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerClient) SetEnvs(ctx context.Context, in *SetEnvsRequest, opts ...grpc.CallOption) (*SetEnvsResponse, error) {
	out := new(SetEnvsResponse)
	err := c.cc.Invoke(ctx, "/container.Container/SetEnvs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerClient) LoadCode(ctx context.Context, in *LoadCodeRequest, opts ...grpc.CallOption) (*LoadCodeResponse, error) {
	out := new(LoadCodeResponse)
	err := c.cc.Invoke(ctx, "/container.Container/LoadCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/container.Container/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerServer is the server API for Container service.
type ContainerServer interface {
	Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error)
	SetEnvs(context.Context, *SetEnvsRequest) (*SetEnvsResponse, error)
	LoadCode(context.Context, *LoadCodeRequest) (*LoadCodeResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

// UnimplementedContainerServer can be embedded to have forward compatible implementations.
type UnimplementedContainerServer struct {
}

func (*UnimplementedContainerServer) Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedContainerServer) SetEnvs(context.Context, *SetEnvsRequest) (*SetEnvsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEnvs not implemented")
}
func (*UnimplementedContainerServer) LoadCode(context.Context, *LoadCodeRequest) (*LoadCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCode not implemented")
}
func (*UnimplementedContainerServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterContainerServer(s *grpc.Server, srv ContainerServer) {
	s.RegisterService(&_Container_serviceDesc, srv)
}

func _Container_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.Container/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServer).Invoke(ctx, req.(*InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Container_SetEnvs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEnvsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServer).SetEnvs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.Container/SetEnvs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServer).SetEnvs(ctx, req.(*SetEnvsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Container_LoadCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServer).LoadCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.Container/LoadCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServer).LoadCode(ctx, req.(*LoadCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Container_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.Container/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Container_serviceDesc = grpc.ServiceDesc{
	ServiceName: "container.Container",
	HandlerType: (*ContainerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _Container_Invoke_Handler,
		},
		{
			MethodName: "SetEnvs",
			Handler:    _Container_SetEnvs_Handler,
		},
		{
			MethodName: "LoadCode",
			Handler:    _Container_LoadCode_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Container_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "container/container.proto",
}
