// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: datanode.proto

package datanode

import (
	context "context"
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

type StatusCode int32

const (
	StatusCode_Success       StatusCode = 0
	StatusCode_NotFound      StatusCode = 1
	StatusCode_InternalError StatusCode = 2
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "Success",
		1: "NotFound",
		2: "InternalError",
	}
	StatusCode_value = map[string]int32{
		"Success":       0,
		"NotFound":      1,
		"InternalError": 2,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_datanode_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_datanode_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{0}
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanReceive bool   `protobuf:"varint,1,opt,name=canReceive,proto3" json:"canReceive,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetCanReceive() bool {
	if x != nil {
		return x.CanReceive
	}
	return false
}

func (x *Ping) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Petition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Petition) Reset() {
	*x = Petition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Petition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Petition) ProtoMessage() {}

func (x *Petition) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Petition.ProtoReflect.Descriptor instead.
func (*Petition) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{1}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  StatusCode `protobuf:"varint,2,opt,name=status,proto3,enum=datanode.StatusCode" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_Success
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ChunkName string `protobuf:"bytes,2,opt,name=chunkName,proto3" json:"chunkName,omitempty"`
	FileName  string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ChunkPos  uint64 `protobuf:"varint,4,opt,name=chunkPos,proto3" json:"chunkPos,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{3}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetChunkName() string {
	if x != nil {
		return x.ChunkName
	}
	return ""
}

func (x *Chunk) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Chunk) GetChunkPos() uint64 {
	if x != nil {
		return x.ChunkPos
	}
	return 0
}

type ChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   []byte     `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ChunkName string     `protobuf:"bytes,2,opt,name=chunkName,proto3" json:"chunkName,omitempty"`
	FileName  string     `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ChunkPos  uint64     `protobuf:"varint,4,opt,name=chunkPos,proto3" json:"chunkPos,omitempty"`
	Status    StatusCode `protobuf:"varint,5,opt,name=status,proto3,enum=datanode.StatusCode" json:"status,omitempty"`
	Message   string     `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChunkResponse) Reset() {
	*x = ChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkResponse) ProtoMessage() {}

func (x *ChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkResponse.ProtoReflect.Descriptor instead.
func (*ChunkResponse) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{4}
}

func (x *ChunkResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ChunkResponse) GetChunkName() string {
	if x != nil {
		return x.ChunkName
	}
	return ""
}

func (x *ChunkResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ChunkResponse) GetChunkPos() uint64 {
	if x != nil {
		return x.ChunkPos
	}
	return 0
}

func (x *ChunkResponse) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_Success
}

func (x *ChunkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	ChunkPos uint64 `protobuf:"varint,2,opt,name=chunkPos,proto3" json:"chunkPos,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datanode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_datanode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_datanode_proto_rawDescGZIP(), []int{5}
}

func (x *FileInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileInfo) GetChunkPos() uint64 {
	if x != nil {
		return x.ChunkPos
	}
	return 0
}

var File_datanode_proto protoreflect.FileDescriptor

var file_datanode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x40, 0x0a, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0a, 0x0a, 0x08,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x77, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x50, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x50, 0x6f, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x50, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x50, 0x6f, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x42, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x50, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x50, 0x6f, 0x73, 0x2a, 0x3a, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x32,
	0xf4, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x0f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x1a, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x0a,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x12, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x0e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x69, 0x6e, 0x67,
	0x1a, 0x0e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x69, 0x6e, 0x67,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_datanode_proto_rawDescOnce sync.Once
	file_datanode_proto_rawDescData = file_datanode_proto_rawDesc
)

func file_datanode_proto_rawDescGZIP() []byte {
	file_datanode_proto_rawDescOnce.Do(func() {
		file_datanode_proto_rawDescData = protoimpl.X.CompressGZIP(file_datanode_proto_rawDescData)
	})
	return file_datanode_proto_rawDescData
}

var file_datanode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_datanode_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_datanode_proto_goTypes = []interface{}{
	(StatusCode)(0),       // 0: datanode.StatusCode
	(*Ping)(nil),          // 1: datanode.ping
	(*Petition)(nil),      // 2: datanode.petition
	(*Response)(nil),      // 3: datanode.response
	(*Chunk)(nil),         // 4: datanode.chunk
	(*ChunkResponse)(nil), // 5: datanode.chunkResponse
	(*FileInfo)(nil),      // 6: datanode.fileInfo
}
var file_datanode_proto_depIdxs = []int32{
	0, // 0: datanode.response.status:type_name -> datanode.StatusCode
	0, // 1: datanode.chunkResponse.status:type_name -> datanode.StatusCode
	4, // 2: datanode.DataNodeHandler.UploadBook:input_type -> datanode.chunk
	4, // 3: datanode.DataNodeHandler.UploadFile:input_type -> datanode.chunk
	6, // 4: datanode.DataNodeHandler.DownloadFile:input_type -> datanode.fileInfo
	1, // 5: datanode.DataNodeHandler.Ping:input_type -> datanode.ping
	3, // 6: datanode.DataNodeHandler.UploadBook:output_type -> datanode.response
	3, // 7: datanode.DataNodeHandler.UploadFile:output_type -> datanode.response
	5, // 8: datanode.DataNodeHandler.DownloadFile:output_type -> datanode.chunkResponse
	1, // 9: datanode.DataNodeHandler.Ping:output_type -> datanode.ping
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_datanode_proto_init() }
func file_datanode_proto_init() {
	if File_datanode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datanode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_datanode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Petition); i {
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
		file_datanode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_datanode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_datanode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkResponse); i {
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
		file_datanode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
			RawDescriptor: file_datanode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datanode_proto_goTypes,
		DependencyIndexes: file_datanode_proto_depIdxs,
		EnumInfos:         file_datanode_proto_enumTypes,
		MessageInfos:      file_datanode_proto_msgTypes,
	}.Build()
	File_datanode_proto = out.File
	file_datanode_proto_rawDesc = nil
	file_datanode_proto_goTypes = nil
	file_datanode_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataNodeHandlerClient is the client API for DataNodeHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataNodeHandlerClient interface {
	UploadBook(ctx context.Context, opts ...grpc.CallOption) (DataNodeHandler_UploadBookClient, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (DataNodeHandler_UploadFileClient, error)
	DownloadFile(ctx context.Context, opts ...grpc.CallOption) (DataNodeHandler_DownloadFileClient, error)
	Ping(ctx context.Context, opts ...grpc.CallOption) (DataNodeHandler_PingClient, error)
}

type dataNodeHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeHandlerClient(cc grpc.ClientConnInterface) DataNodeHandlerClient {
	return &dataNodeHandlerClient{cc}
}

func (c *dataNodeHandlerClient) UploadBook(ctx context.Context, opts ...grpc.CallOption) (DataNodeHandler_UploadBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNodeHandler_serviceDesc.Streams[0], "/datanode.DataNodeHandler/UploadBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeHandlerUploadBookClient{stream}
	return x, nil
}

type DataNodeHandler_UploadBookClient interface {
	Send(*Chunk) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type dataNodeHandlerUploadBookClient struct {
	grpc.ClientStream
}

func (x *dataNodeHandlerUploadBookClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeHandlerUploadBookClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataNodeHandlerClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (DataNodeHandler_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNodeHandler_serviceDesc.Streams[1], "/datanode.DataNodeHandler/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeHandlerUploadFileClient{stream}
	return x, nil
}

type DataNodeHandler_UploadFileClient interface {
	Send(*Chunk) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type dataNodeHandlerUploadFileClient struct {
	grpc.ClientStream
}

func (x *dataNodeHandlerUploadFileClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeHandlerUploadFileClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataNodeHandlerClient) DownloadFile(ctx context.Context, opts ...grpc.CallOption) (DataNodeHandler_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNodeHandler_serviceDesc.Streams[2], "/datanode.DataNodeHandler/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeHandlerDownloadFileClient{stream}
	return x, nil
}

type DataNodeHandler_DownloadFileClient interface {
	Send(*FileInfo) error
	Recv() (*ChunkResponse, error)
	grpc.ClientStream
}

type dataNodeHandlerDownloadFileClient struct {
	grpc.ClientStream
}

func (x *dataNodeHandlerDownloadFileClient) Send(m *FileInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeHandlerDownloadFileClient) Recv() (*ChunkResponse, error) {
	m := new(ChunkResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataNodeHandlerClient) Ping(ctx context.Context, opts ...grpc.CallOption) (DataNodeHandler_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataNodeHandler_serviceDesc.Streams[3], "/datanode.DataNodeHandler/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeHandlerPingClient{stream}
	return x, nil
}

type DataNodeHandler_PingClient interface {
	Send(*Ping) error
	Recv() (*Ping, error)
	grpc.ClientStream
}

type dataNodeHandlerPingClient struct {
	grpc.ClientStream
}

func (x *dataNodeHandlerPingClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeHandlerPingClient) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataNodeHandlerServer is the server API for DataNodeHandler service.
type DataNodeHandlerServer interface {
	UploadBook(DataNodeHandler_UploadBookServer) error
	UploadFile(DataNodeHandler_UploadFileServer) error
	DownloadFile(DataNodeHandler_DownloadFileServer) error
	Ping(DataNodeHandler_PingServer) error
}

// UnimplementedDataNodeHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedDataNodeHandlerServer struct {
}

func (*UnimplementedDataNodeHandlerServer) UploadBook(DataNodeHandler_UploadBookServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBook not implemented")
}
func (*UnimplementedDataNodeHandlerServer) UploadFile(DataNodeHandler_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedDataNodeHandlerServer) DownloadFile(DataNodeHandler_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (*UnimplementedDataNodeHandlerServer) Ping(DataNodeHandler_PingServer) error {
	return status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterDataNodeHandlerServer(s *grpc.Server, srv DataNodeHandlerServer) {
	s.RegisterService(&_DataNodeHandler_serviceDesc, srv)
}

func _DataNodeHandler_UploadBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeHandlerServer).UploadBook(&dataNodeHandlerUploadBookServer{stream})
}

type DataNodeHandler_UploadBookServer interface {
	Send(*Response) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type dataNodeHandlerUploadBookServer struct {
	grpc.ServerStream
}

func (x *dataNodeHandlerUploadBookServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeHandlerUploadBookServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataNodeHandler_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeHandlerServer).UploadFile(&dataNodeHandlerUploadFileServer{stream})
}

type DataNodeHandler_UploadFileServer interface {
	Send(*Response) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type dataNodeHandlerUploadFileServer struct {
	grpc.ServerStream
}

func (x *dataNodeHandlerUploadFileServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeHandlerUploadFileServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataNodeHandler_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeHandlerServer).DownloadFile(&dataNodeHandlerDownloadFileServer{stream})
}

type DataNodeHandler_DownloadFileServer interface {
	Send(*ChunkResponse) error
	Recv() (*FileInfo, error)
	grpc.ServerStream
}

type dataNodeHandlerDownloadFileServer struct {
	grpc.ServerStream
}

func (x *dataNodeHandlerDownloadFileServer) Send(m *ChunkResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeHandlerDownloadFileServer) Recv() (*FileInfo, error) {
	m := new(FileInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataNodeHandler_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeHandlerServer).Ping(&dataNodeHandlerPingServer{stream})
}

type DataNodeHandler_PingServer interface {
	Send(*Ping) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type dataNodeHandlerPingServer struct {
	grpc.ServerStream
}

func (x *dataNodeHandlerPingServer) Send(m *Ping) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeHandlerPingServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataNodeHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datanode.DataNodeHandler",
	HandlerType: (*DataNodeHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadBook",
			Handler:       _DataNodeHandler_UploadBook_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _DataNodeHandler_UploadFile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _DataNodeHandler_DownloadFile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Ping",
			Handler:       _DataNodeHandler_Ping_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "datanode.proto",
}
