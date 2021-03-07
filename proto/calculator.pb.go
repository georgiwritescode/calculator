// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: proto/calculator.proto

package calculatorpb

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

type Numbers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First  int32 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Second int32 `protobuf:"varint,2,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *Numbers) Reset() {
	*x = Numbers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Numbers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Numbers) ProtoMessage() {}

func (x *Numbers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Numbers.ProtoReflect.Descriptor instead.
func (*Numbers) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Numbers) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *Numbers) GetSecond() int32 {
	if x != nil {
		return x.Second
	}
	return 0
}

type CalculateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numbers *Numbers `protobuf:"bytes,1,opt,name=numbers,proto3" json:"numbers,omitempty"`
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateRequest) GetNumbers() *Numbers {
	if x != nil {
		return x.Numbers
	}
	return nil
}

type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *CalculateResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type PrimeNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeNumber int32 `protobuf:"varint,1,opt,name=prime_number,json=primeNumber,proto3" json:"prime_number,omitempty"`
}

func (x *PrimeNumber) Reset() {
	*x = PrimeNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumber) ProtoMessage() {}

func (x *PrimeNumber) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumber.ProtoReflect.Descriptor instead.
func (*PrimeNumber) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeNumber) GetPrimeNumber() int32 {
	if x != nil {
		return x.PrimeNumber
	}
	return 0
}

type PrimeNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeNumber *PrimeNumber `protobuf:"bytes,1,opt,name=prime_number,json=primeNumber,proto3" json:"prime_number,omitempty"`
}

func (x *PrimeNumberRequest) Reset() {
	*x = PrimeNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberRequest) ProtoMessage() {}

func (x *PrimeNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *PrimeNumberRequest) GetPrimeNumber() *PrimeNumber {
	if x != nil {
		return x.PrimeNumber
	}
	return nil
}

type PrimeNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeNumber int32 `protobuf:"varint,1,opt,name=prime_number,json=primeNumber,proto3" json:"prime_number,omitempty"`
}

func (x *PrimeNumberResponse) Reset() {
	*x = PrimeNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberResponse) ProtoMessage() {}

func (x *PrimeNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *PrimeNumberResponse) GetPrimeNumber() int32 {
	if x != nil {
		return x.PrimeNumber
	}
	return 0
}

type ComputeAverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ComputeAverageRequest) Reset() {
	*x = ComputeAverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageRequest) ProtoMessage() {}

func (x *ComputeAverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageRequest.ProtoReflect.Descriptor instead.
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *ComputeAverageRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float32 `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *ComputeAverageResponse) Reset() {
	*x = ComputeAverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageResponse) ProtoMessage() {}

func (x *ComputeAverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageResponse.ProtoReflect.Descriptor instead.
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *ComputeAverageResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

type FindMaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FindMaximumRequest) Reset() {
	*x = FindMaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumRequest) ProtoMessage() {}

func (x *FindMaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumRequest.ProtoReflect.Descriptor instead.
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{8}
}

func (x *FindMaximumRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentMax int32 `protobuf:"varint,1,opt,name=current_max,json=currentMax,proto3" json:"current_max,omitempty"`
}

func (x *FindMaximumResponse) Reset() {
	*x = FindMaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumResponse) ProtoMessage() {}

func (x *FindMaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumResponse.ProtoReflect.Descriptor instead.
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{9}
}

func (x *FindMaximumResponse) GetCurrentMax() int32 {
	if x != nil {
		return x.CurrentMax
	}
	return 0
}

var File_proto_calculator_proto protoreflect.FileDescriptor

var file_proto_calculator_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0x37, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x40, 0x0a, 0x10,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x25,
	0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x30, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x6d,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x78, 0x32, 0xeb, 0x02,
	0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a,
	0x18, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x20,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x52, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_calculator_proto_rawDescOnce sync.Once
	file_proto_calculator_proto_rawDescData = file_proto_calculator_proto_rawDesc
)

func file_proto_calculator_proto_rawDescGZIP() []byte {
	file_proto_calculator_proto_rawDescOnce.Do(func() {
		file_proto_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_calculator_proto_rawDescData)
	})
	return file_proto_calculator_proto_rawDescData
}

var file_proto_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_calculator_proto_goTypes = []interface{}{
	(*Numbers)(nil),                // 0: calculate.Numbers
	(*CalculateRequest)(nil),       // 1: calculate.CalculateRequest
	(*CalculateResponse)(nil),      // 2: calculate.CalculateResponse
	(*PrimeNumber)(nil),            // 3: calculate.PrimeNumber
	(*PrimeNumberRequest)(nil),     // 4: calculate.PrimeNumberRequest
	(*PrimeNumberResponse)(nil),    // 5: calculate.PrimeNumberResponse
	(*ComputeAverageRequest)(nil),  // 6: calculate.ComputeAverageRequest
	(*ComputeAverageResponse)(nil), // 7: calculate.ComputeAverageResponse
	(*FindMaximumRequest)(nil),     // 8: calculate.FindMaximumRequest
	(*FindMaximumResponse)(nil),    // 9: calculate.FindMaximumResponse
}
var file_proto_calculator_proto_depIdxs = []int32{
	0, // 0: calculate.CalculateRequest.numbers:type_name -> calculate.Numbers
	3, // 1: calculate.PrimeNumberRequest.prime_number:type_name -> calculate.PrimeNumber
	1, // 2: calculate.CalculatorService.Calculate:input_type -> calculate.CalculateRequest
	4, // 3: calculate.CalculatorService.PrimeNumberDecomposition:input_type -> calculate.PrimeNumberRequest
	6, // 4: calculate.CalculatorService.ComputeAverage:input_type -> calculate.ComputeAverageRequest
	8, // 5: calculate.CalculatorService.FindMaximum:input_type -> calculate.FindMaximumRequest
	2, // 6: calculate.CalculatorService.Calculate:output_type -> calculate.CalculateResponse
	5, // 7: calculate.CalculatorService.PrimeNumberDecomposition:output_type -> calculate.PrimeNumberResponse
	7, // 8: calculate.CalculatorService.ComputeAverage:output_type -> calculate.ComputeAverageResponse
	9, // 9: calculate.CalculatorService.FindMaximum:output_type -> calculate.FindMaximumResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_calculator_proto_init() }
func file_proto_calculator_proto_init() {
	if File_proto_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Numbers); i {
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
		file_proto_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRequest); i {
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
		file_proto_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse); i {
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
		file_proto_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumber); i {
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
		file_proto_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberRequest); i {
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
		file_proto_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberResponse); i {
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
		file_proto_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageRequest); i {
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
		file_proto_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageResponse); i {
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
		file_proto_calculator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumRequest); i {
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
		file_proto_calculator_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumResponse); i {
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
			RawDescriptor: file_proto_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_calculator_proto_goTypes,
		DependencyIndexes: file_proto_calculator_proto_depIdxs,
		MessageInfos:      file_proto_calculator_proto_msgTypes,
	}.Build()
	File_proto_calculator_proto = out.File
	file_proto_calculator_proto_rawDesc = nil
	file_proto_calculator_proto_goTypes = nil
	file_proto_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/calculate.CalculatorService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculate.CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculate.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculate.CalculatorService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaximumClient{stream}
	return x, nil
}

type CalculatorService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
	PrimeNumberDecomposition(*PrimeNumberRequest, CalculatorService_PrimeNumberDecompositionServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
	FindMaximum(CalculatorService_FindMaximumServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeNumberDecomposition(*PrimeNumberRequest, CalculatorService_PrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecomposition not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMaximum(CalculatorService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculate.CalculatorService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMaximum(&calculatorServiceFindMaximumServer{stream})
}

type CalculatorService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculate.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculatorService_Calculate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalculatorService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/calculator.proto",
}
