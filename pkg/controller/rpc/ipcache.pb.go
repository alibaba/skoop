// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: ipcache.proto

package rpc

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

type AddressFamily int32

const (
	AddressFamily_AF_INET  AddressFamily = 0
	AddressFamily_AF_INET6 AddressFamily = 1
)

// Enum value maps for AddressFamily.
var (
	AddressFamily_name = map[int32]string{
		0: "AF_INET",
		1: "AF_INET6",
	}
	AddressFamily_value = map[string]int32{
		"AF_INET":  0,
		"AF_INET6": 1,
	}
)

func (x AddressFamily) Enum() *AddressFamily {
	p := new(AddressFamily)
	*p = x
	return p
}

func (x AddressFamily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressFamily) Descriptor() protoreflect.EnumDescriptor {
	return file_ipcache_proto_enumTypes[0].Descriptor()
}

func (AddressFamily) Type() protoreflect.EnumType {
	return &file_ipcache_proto_enumTypes[0]
}

func (x AddressFamily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressFamily.Descriptor instead.
func (AddressFamily) EnumDescriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{0}
}

type ValueType int32

const (
	ValueType_Pod  ValueType = 0
	ValueType_Node ValueType = 1
)

// Enum value maps for ValueType.
var (
	ValueType_name = map[int32]string{
		0: "Pod",
		1: "Node",
	}
	ValueType_value = map[string]int32{
		"Pod":  0,
		"Node": 1,
	}
)

func (x ValueType) Enum() *ValueType {
	p := new(ValueType)
	*p = x
	return p
}

func (x ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_ipcache_proto_enumTypes[1].Descriptor()
}

func (ValueType) Type() protoreflect.EnumType {
	return &file_ipcache_proto_enumTypes[1]
}

func (x ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValueType.Descriptor instead.
func (ValueType) EnumDescriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{1}
}

type OpCode int32

const (
	OpCode_Set OpCode = 0
	OpCode_Del OpCode = 1
)

// Enum value maps for OpCode.
var (
	OpCode_name = map[int32]string{
		0: "Set",
		1: "Del",
	}
	OpCode_value = map[string]int32{
		"Set": 0,
		"Del": 1,
	}
)

func (x OpCode) Enum() *OpCode {
	p := new(OpCode)
	*p = x
	return p
}

func (x OpCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ipcache_proto_enumTypes[2].Descriptor()
}

func (OpCode) Type() protoreflect.EnumType {
	return &file_ipcache_proto_enumTypes[2]
}

func (x OpCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpCode.Descriptor instead.
func (OpCode) EnumDescriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{2}
}

type PodMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PodMeta) Reset() {
	*x = PodMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipcache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodMeta) ProtoMessage() {}

func (x *PodMeta) ProtoReflect() protoreflect.Message {
	mi := &file_ipcache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodMeta.ProtoReflect.Descriptor instead.
func (*PodMeta) Descriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{0}
}

func (x *PodMeta) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NodeMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NodeMeta) Reset() {
	*x = NodeMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipcache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMeta) ProtoMessage() {}

func (x *NodeMeta) ProtoReflect() protoreflect.Message {
	mi := &file_ipcache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMeta.ProtoReflect.Descriptor instead.
func (*NodeMeta) Descriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{1}
}

func (x *NodeMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CacheEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP   string    `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	Type ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=controller_rpc.ValueType" json:"type,omitempty"`
	// Types that are assignable to Meta:
	//
	//	*CacheEntry_Pod
	//	*CacheEntry_Node
	Meta isCacheEntry_Meta `protobuf_oneof:"meta"`
}

func (x *CacheEntry) Reset() {
	*x = CacheEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipcache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheEntry) ProtoMessage() {}

func (x *CacheEntry) ProtoReflect() protoreflect.Message {
	mi := &file_ipcache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheEntry.ProtoReflect.Descriptor instead.
func (*CacheEntry) Descriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{2}
}

func (x *CacheEntry) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *CacheEntry) GetType() ValueType {
	if x != nil {
		return x.Type
	}
	return ValueType_Pod
}

func (m *CacheEntry) GetMeta() isCacheEntry_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (x *CacheEntry) GetPod() *PodMeta {
	if x, ok := x.GetMeta().(*CacheEntry_Pod); ok {
		return x.Pod
	}
	return nil
}

func (x *CacheEntry) GetNode() *NodeMeta {
	if x, ok := x.GetMeta().(*CacheEntry_Node); ok {
		return x.Node
	}
	return nil
}

type isCacheEntry_Meta interface {
	isCacheEntry_Meta()
}

type CacheEntry_Pod struct {
	Pod *PodMeta `protobuf:"bytes,3,opt,name=pod,proto3,oneof"`
}

type CacheEntry_Node struct {
	Node *NodeMeta `protobuf:"bytes,4,opt,name=node,proto3,oneof"`
}

func (*CacheEntry_Pod) isCacheEntry_Meta() {}

func (*CacheEntry_Node) isCacheEntry_Meta() {}

type ListCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCacheRequest) Reset() {
	*x = ListCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipcache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCacheRequest) ProtoMessage() {}

func (x *ListCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ipcache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCacheRequest.ProtoReflect.Descriptor instead.
func (*ListCacheRequest) Descriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{3}
}

type ListCacheResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period   string        `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	Revision uint64        `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	Entries  []*CacheEntry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ListCacheResponse) Reset() {
	*x = ListCacheResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipcache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCacheResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCacheResponse) ProtoMessage() {}

func (x *ListCacheResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ipcache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCacheResponse.ProtoReflect.Descriptor instead.
func (*ListCacheResponse) Descriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{4}
}

func (x *ListCacheResponse) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *ListCacheResponse) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *ListCacheResponse) GetEntries() []*CacheEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type WatchCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period   string `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	Revision uint64 `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *WatchCacheRequest) Reset() {
	*x = WatchCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipcache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchCacheRequest) ProtoMessage() {}

func (x *WatchCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ipcache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchCacheRequest.ProtoReflect.Descriptor instead.
func (*WatchCacheRequest) Descriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{5}
}

func (x *WatchCacheRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *WatchCacheRequest) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

type WatchCacheResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision uint64      `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	Opcode   OpCode      `protobuf:"varint,2,opt,name=opcode,proto3,enum=controller_rpc.OpCode" json:"opcode,omitempty"`
	Entry    *CacheEntry `protobuf:"bytes,3,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *WatchCacheResponse) Reset() {
	*x = WatchCacheResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipcache_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchCacheResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchCacheResponse) ProtoMessage() {}

func (x *WatchCacheResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ipcache_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchCacheResponse.ProtoReflect.Descriptor instead.
func (*WatchCacheResponse) Descriptor() ([]byte, []int) {
	return file_ipcache_proto_rawDescGZIP(), []int{6}
}

func (x *WatchCacheResponse) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *WatchCacheResponse) GetOpcode() OpCode {
	if x != nil {
		return x.Opcode
	}
	return OpCode_Set
}

func (x *WatchCacheResponse) GetEntry() *CacheEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

var File_ipcache_proto protoreflect.FileDescriptor

var file_ipcache_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x70, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x22,
	0x3b, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb0, 0x01, 0x0a,
	0x0a, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x2d, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x70, 0x6f,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22,
	0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x7d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x47, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x01, 0x0a, 0x12,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x4f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x30,
	0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x2a, 0x2a, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x46, 0x5f, 0x49, 0x4e, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x46, 0x5f, 0x49, 0x4e, 0x45, 0x54, 0x36, 0x10, 0x01, 0x2a, 0x1e, 0x0a, 0x09,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x6f, 0x64,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x10, 0x01, 0x2a, 0x1a, 0x0a, 0x06,
	0x4f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x10, 0x01, 0x32, 0xb9, 0x01, 0x0a, 0x0e, 0x49, 0x50, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a,
	0x0a, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ipcache_proto_rawDescOnce sync.Once
	file_ipcache_proto_rawDescData = file_ipcache_proto_rawDesc
)

func file_ipcache_proto_rawDescGZIP() []byte {
	file_ipcache_proto_rawDescOnce.Do(func() {
		file_ipcache_proto_rawDescData = protoimpl.X.CompressGZIP(file_ipcache_proto_rawDescData)
	})
	return file_ipcache_proto_rawDescData
}

var file_ipcache_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_ipcache_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ipcache_proto_goTypes = []interface{}{
	(AddressFamily)(0),         // 0: controller_rpc.AddressFamily
	(ValueType)(0),             // 1: controller_rpc.ValueType
	(OpCode)(0),                // 2: controller_rpc.OpCode
	(*PodMeta)(nil),            // 3: controller_rpc.PodMeta
	(*NodeMeta)(nil),           // 4: controller_rpc.NodeMeta
	(*CacheEntry)(nil),         // 5: controller_rpc.CacheEntry
	(*ListCacheRequest)(nil),   // 6: controller_rpc.ListCacheRequest
	(*ListCacheResponse)(nil),  // 7: controller_rpc.ListCacheResponse
	(*WatchCacheRequest)(nil),  // 8: controller_rpc.WatchCacheRequest
	(*WatchCacheResponse)(nil), // 9: controller_rpc.WatchCacheResponse
}
var file_ipcache_proto_depIdxs = []int32{
	1, // 0: controller_rpc.CacheEntry.type:type_name -> controller_rpc.ValueType
	3, // 1: controller_rpc.CacheEntry.pod:type_name -> controller_rpc.PodMeta
	4, // 2: controller_rpc.CacheEntry.node:type_name -> controller_rpc.NodeMeta
	5, // 3: controller_rpc.ListCacheResponse.entries:type_name -> controller_rpc.CacheEntry
	2, // 4: controller_rpc.WatchCacheResponse.opcode:type_name -> controller_rpc.OpCode
	5, // 5: controller_rpc.WatchCacheResponse.entry:type_name -> controller_rpc.CacheEntry
	6, // 6: controller_rpc.IPCacheService.ListCache:input_type -> controller_rpc.ListCacheRequest
	8, // 7: controller_rpc.IPCacheService.WatchCache:input_type -> controller_rpc.WatchCacheRequest
	7, // 8: controller_rpc.IPCacheService.ListCache:output_type -> controller_rpc.ListCacheResponse
	9, // 9: controller_rpc.IPCacheService.WatchCache:output_type -> controller_rpc.WatchCacheResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ipcache_proto_init() }
func file_ipcache_proto_init() {
	if File_ipcache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ipcache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodMeta); i {
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
		file_ipcache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMeta); i {
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
		file_ipcache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheEntry); i {
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
		file_ipcache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCacheRequest); i {
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
		file_ipcache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCacheResponse); i {
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
		file_ipcache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchCacheRequest); i {
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
		file_ipcache_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchCacheResponse); i {
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
	file_ipcache_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CacheEntry_Pod)(nil),
		(*CacheEntry_Node)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ipcache_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ipcache_proto_goTypes,
		DependencyIndexes: file_ipcache_proto_depIdxs,
		EnumInfos:         file_ipcache_proto_enumTypes,
		MessageInfos:      file_ipcache_proto_msgTypes,
	}.Build()
	File_ipcache_proto = out.File
	file_ipcache_proto_rawDesc = nil
	file_ipcache_proto_goTypes = nil
	file_ipcache_proto_depIdxs = nil
}
