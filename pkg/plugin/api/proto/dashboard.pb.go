// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dashboard.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type KeyRequest struct {
	Namespace            string               `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ApiVersion           string               `protobuf:"bytes,2,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind                 string               `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	LabelSelector        *wrappers.BytesValue `protobuf:"bytes,5,opt,name=labelSelector,proto3" json:"labelSelector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KeyRequest) Reset()         { *m = KeyRequest{} }
func (m *KeyRequest) String() string { return proto.CompactTextString(m) }
func (*KeyRequest) ProtoMessage()    {}
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{1}
}

func (m *KeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRequest.Unmarshal(m, b)
}
func (m *KeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRequest.Marshal(b, m, deterministic)
}
func (m *KeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRequest.Merge(m, src)
}
func (m *KeyRequest) XXX_Size() int {
	return xxx_messageInfo_KeyRequest.Size(m)
}
func (m *KeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRequest proto.InternalMessageInfo

func (m *KeyRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *KeyRequest) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *KeyRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *KeyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyRequest) GetLabelSelector() *wrappers.BytesValue {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

type ListResponse struct {
	Objects              [][]byte `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{2}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetObjects() [][]byte {
	if m != nil {
		return m.Objects
	}
	return nil
}

type GetResponse struct {
	Object               []byte   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetObject() []byte {
	if m != nil {
		return m.Object
	}
	return nil
}

type UpdateRequest struct {
	Object               []byte   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetObject() []byte {
	if m != nil {
		return m.Object
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{5}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type PortForwardRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName              string   `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
	ContainerName        string   `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	PortNumber           uint32   `protobuf:"varint,4,opt,name=portNumber,proto3" json:"portNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortForwardRequest) Reset()         { *m = PortForwardRequest{} }
func (m *PortForwardRequest) String() string { return proto.CompactTextString(m) }
func (*PortForwardRequest) ProtoMessage()    {}
func (*PortForwardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{6}
}

func (m *PortForwardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortForwardRequest.Unmarshal(m, b)
}
func (m *PortForwardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortForwardRequest.Marshal(b, m, deterministic)
}
func (m *PortForwardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortForwardRequest.Merge(m, src)
}
func (m *PortForwardRequest) XXX_Size() int {
	return xxx_messageInfo_PortForwardRequest.Size(m)
}
func (m *PortForwardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PortForwardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PortForwardRequest proto.InternalMessageInfo

func (m *PortForwardRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PortForwardRequest) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *PortForwardRequest) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *PortForwardRequest) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

type PortForwardResponse struct {
	PortForwardID        string   `protobuf:"bytes,1,opt,name=portForwardID,proto3" json:"portForwardID,omitempty"`
	PortNumber           uint32   `protobuf:"varint,2,opt,name=portNumber,proto3" json:"portNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortForwardResponse) Reset()         { *m = PortForwardResponse{} }
func (m *PortForwardResponse) String() string { return proto.CompactTextString(m) }
func (*PortForwardResponse) ProtoMessage()    {}
func (*PortForwardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{7}
}

func (m *PortForwardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortForwardResponse.Unmarshal(m, b)
}
func (m *PortForwardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortForwardResponse.Marshal(b, m, deterministic)
}
func (m *PortForwardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortForwardResponse.Merge(m, src)
}
func (m *PortForwardResponse) XXX_Size() int {
	return xxx_messageInfo_PortForwardResponse.Size(m)
}
func (m *PortForwardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PortForwardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PortForwardResponse proto.InternalMessageInfo

func (m *PortForwardResponse) GetPortForwardID() string {
	if m != nil {
		return m.PortForwardID
	}
	return ""
}

func (m *PortForwardResponse) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

type CancelPortForwardRequest struct {
	PortForwardID        string   `protobuf:"bytes,1,opt,name=portForwardID,proto3" json:"portForwardID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelPortForwardRequest) Reset()         { *m = CancelPortForwardRequest{} }
func (m *CancelPortForwardRequest) String() string { return proto.CompactTextString(m) }
func (*CancelPortForwardRequest) ProtoMessage()    {}
func (*CancelPortForwardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{8}
}

func (m *CancelPortForwardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelPortForwardRequest.Unmarshal(m, b)
}
func (m *CancelPortForwardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelPortForwardRequest.Marshal(b, m, deterministic)
}
func (m *CancelPortForwardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelPortForwardRequest.Merge(m, src)
}
func (m *CancelPortForwardRequest) XXX_Size() int {
	return xxx_messageInfo_CancelPortForwardRequest.Size(m)
}
func (m *CancelPortForwardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelPortForwardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelPortForwardRequest proto.InternalMessageInfo

func (m *CancelPortForwardRequest) GetPortForwardID() string {
	if m != nil {
		return m.PortForwardID
	}
	return ""
}

type NamespacesResponse struct {
	Namespaces           []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamespacesResponse) Reset()         { *m = NamespacesResponse{} }
func (m *NamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*NamespacesResponse) ProtoMessage()    {}
func (*NamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b97678da3a35dfb, []int{9}
}

func (m *NamespacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespacesResponse.Unmarshal(m, b)
}
func (m *NamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespacesResponse.Marshal(b, m, deterministic)
}
func (m *NamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacesResponse.Merge(m, src)
}
func (m *NamespacesResponse) XXX_Size() int {
	return xxx_messageInfo_NamespacesResponse.Size(m)
}
func (m *NamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacesResponse proto.InternalMessageInfo

func (m *NamespacesResponse) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*KeyRequest)(nil), "proto.KeyRequest")
	proto.RegisterType((*ListResponse)(nil), "proto.ListResponse")
	proto.RegisterType((*GetResponse)(nil), "proto.GetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "proto.UpdateResponse")
	proto.RegisterType((*PortForwardRequest)(nil), "proto.PortForwardRequest")
	proto.RegisterType((*PortForwardResponse)(nil), "proto.PortForwardResponse")
	proto.RegisterType((*CancelPortForwardRequest)(nil), "proto.CancelPortForwardRequest")
	proto.RegisterType((*NamespacesResponse)(nil), "proto.NamespacesResponse")
}

func init() { proto.RegisterFile("dashboard.proto", fileDescriptor_9b97678da3a35dfb) }

var fileDescriptor_9b97678da3a35dfb = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0xe6, 0x4b, 0x99, 0x24, 0x85, 0x6e, 0x00, 0x19, 0x83, 0x42, 0x64, 0x15, 0x91,
	0x03, 0x72, 0x45, 0x81, 0x03, 0x37, 0x28, 0x21, 0x15, 0x02, 0x45, 0xc8, 0x88, 0x5e, 0x38, 0xad,
	0xed, 0xa1, 0x18, 0x1c, 0xef, 0xb2, 0xbb, 0x51, 0x95, 0xd7, 0xe0, 0xc4, 0x8b, 0xf0, 0x7e, 0x95,
	0x77, 0xd7, 0xb5, 0xdd, 0xa4, 0x52, 0x4f, 0xd9, 0xf9, 0xcf, 0x47, 0x66, 0x7e, 0x33, 0x86, 0x3b,
	0x09, 0x95, 0x3f, 0x23, 0x46, 0x45, 0x12, 0x70, 0xc1, 0x14, 0x23, 0x1d, 0xfd, 0xe3, 0x4d, 0xce,
	0x19, 0x3b, 0xcf, 0xf0, 0x48, 0x5b, 0xd1, 0xfa, 0xc7, 0xd1, 0x85, 0xa0, 0x9c, 0xa3, 0x90, 0x26,
	0xcc, 0xef, 0x41, 0xe7, 0xc3, 0x8a, 0xab, 0x8d, 0xff, 0xdf, 0x01, 0xf8, 0x84, 0x9b, 0x10, 0xff,
	0xac, 0x51, 0x2a, 0xf2, 0x18, 0xfa, 0x39, 0x5d, 0xa1, 0xe4, 0x34, 0x46, 0xd7, 0x99, 0x3a, 0xb3,
	0x7e, 0x58, 0x09, 0x64, 0x02, 0x40, 0x79, 0x7a, 0x86, 0x42, 0xa6, 0x2c, 0x77, 0xf7, 0xb4, 0xbb,
	0xa6, 0x10, 0x02, 0xed, 0xdf, 0x69, 0x9e, 0xb8, 0x2d, 0xed, 0xd1, 0xef, 0x42, 0x2b, 0x0a, 0xb8,
	0x6d, 0xa3, 0x15, 0x6f, 0xf2, 0x0e, 0x46, 0x19, 0x8d, 0x30, 0xfb, 0x8a, 0x19, 0xc6, 0x8a, 0x09,
	0xb7, 0x33, 0x75, 0x66, 0x83, 0xe3, 0x47, 0x81, 0xe9, 0x3a, 0x28, 0xbb, 0x0e, 0x4e, 0x36, 0x0a,
	0xe5, 0x19, 0xcd, 0xd6, 0x18, 0x36, 0x33, 0xfc, 0x19, 0x0c, 0x3f, 0xa7, 0x52, 0x85, 0x28, 0x39,
	0xcb, 0x25, 0x12, 0x17, 0x7a, 0x2c, 0xfa, 0x85, 0xb1, 0x92, 0xae, 0x33, 0x6d, 0xcd, 0x86, 0x61,
	0x69, 0xfa, 0x4f, 0x61, 0x70, 0x8a, 0x55, 0xe0, 0x03, 0xe8, 0x1a, 0x8f, 0x1e, 0x6f, 0x18, 0x5a,
	0xcb, 0x7f, 0x06, 0xa3, 0x6f, 0x3c, 0xa1, 0x0a, 0x4b, 0x14, 0x37, 0x05, 0xde, 0x85, 0xfd, 0x32,
	0xd0, 0x94, 0xf4, 0xff, 0x3a, 0x40, 0xbe, 0x30, 0xa1, 0x16, 0x4c, 0x5c, 0x50, 0x91, 0xdc, 0x8e,
	0xa5, 0x0b, 0x3d, 0xce, 0x92, 0x65, 0x81, 0xc6, 0x80, 0x2c, 0x4d, 0x72, 0x08, 0xa3, 0x98, 0xe5,
	0x8a, 0xa6, 0x39, 0x0a, 0xed, 0x37, 0x38, 0x9b, 0x62, 0xb1, 0x0b, 0xce, 0x84, 0x5a, 0xae, 0x57,
	0x11, 0x0a, 0x4d, 0x77, 0x14, 0xd6, 0x14, 0xff, 0x3b, 0x8c, 0x1b, 0x3d, 0xd9, 0xf1, 0x0f, 0x61,
	0xc4, 0x2b, 0xf9, 0xe3, 0xdc, 0x36, 0xd6, 0x14, 0xaf, 0x15, 0xdf, 0xdb, 0x2a, 0xfe, 0x16, 0xdc,
	0xf7, 0x34, 0x8f, 0x31, 0xdb, 0x31, 0xf6, 0xad, 0xfe, 0xc1, 0x7f, 0x05, 0x64, 0x59, 0xb2, 0x90,
	0x57, 0xdd, 0x4d, 0x00, 0xae, 0x08, 0x99, 0x45, 0xf6, 0xc3, 0x9a, 0x72, 0xfc, 0xaf, 0x05, 0xfd,
	0x79, 0x79, 0xf1, 0x24, 0x80, 0x76, 0x71, 0x03, 0xe4, 0xc0, 0x1c, 0x4c, 0x50, 0xdd, 0xb1, 0x37,
	0xb6, 0x52, 0xe3, 0x46, 0x9e, 0x43, 0xeb, 0x14, 0x77, 0x86, 0x13, 0x2b, 0xd5, 0x0f, 0xe5, 0x35,
	0x74, 0xcd, 0x9e, 0xc9, 0x3d, 0xeb, 0x6d, 0xdc, 0x87, 0x77, 0xff, 0x9a, 0x6a, 0xd3, 0xe6, 0x30,
	0xa8, 0x41, 0x21, 0x0f, 0x6d, 0xd4, 0x36, 0x28, 0xcf, 0xdb, 0xe5, 0xb2, 0x55, 0x4e, 0xe0, 0x60,
	0x0b, 0x30, 0x79, 0x62, 0x13, 0x6e, 0x42, 0xef, 0x0d, 0x6d, 0x80, 0xfe, 0xb4, 0xc9, 0x1b, 0xd8,
	0x2f, 0xc6, 0xaf, 0x30, 0x93, 0x86, 0xdf, 0x2b, 0x5b, 0xdb, 0xb1, 0x87, 0x17, 0x30, 0x5e, 0x30,
	0x11, 0xe3, 0x42, 0xb0, 0x5c, 0x61, 0x9e, 0x58, 0x10, 0xcd, 0xfc, 0x86, 0x15, 0x75, 0xb5, 0xf1,
	0xf2, 0x32, 0x00, 0x00, 0xff, 0xff, 0x15, 0xd4, 0x9d, 0x8f, 0x92, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DashboardClient is the client API for Dashboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DashboardClient interface {
	List(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	PortForward(ctx context.Context, in *PortForwardRequest, opts ...grpc.CallOption) (*PortForwardResponse, error)
	CancelPortForward(ctx context.Context, in *CancelPortForwardRequest, opts ...grpc.CallOption) (*Empty, error)
	ListNamespaces(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NamespacesResponse, error)
	ForceFrontendUpdate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type dashboardClient struct {
	cc *grpc.ClientConn
}

func NewDashboardClient(cc *grpc.ClientConn) DashboardClient {
	return &dashboardClient{cc}
}

func (c *dashboardClient) List(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) PortForward(ctx context.Context, in *PortForwardRequest, opts ...grpc.CallOption) (*PortForwardResponse, error) {
	out := new(PortForwardResponse)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/PortForward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) CancelPortForward(ctx context.Context, in *CancelPortForwardRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/CancelPortForward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) ListNamespaces(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NamespacesResponse, error) {
	out := new(NamespacesResponse)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) ForceFrontendUpdate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/ForceFrontendUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServer is the server API for Dashboard service.
type DashboardServer interface {
	List(context.Context, *KeyRequest) (*ListResponse, error)
	Get(context.Context, *KeyRequest) (*GetResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	PortForward(context.Context, *PortForwardRequest) (*PortForwardResponse, error)
	CancelPortForward(context.Context, *CancelPortForwardRequest) (*Empty, error)
	ListNamespaces(context.Context, *Empty) (*NamespacesResponse, error)
	ForceFrontendUpdate(context.Context, *Empty) (*Empty, error)
}

func RegisterDashboardServer(s *grpc.Server, srv DashboardServer) {
	s.RegisterService(&_Dashboard_serviceDesc, srv)
}

func _Dashboard_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).List(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).Get(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_PortForward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).PortForward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/PortForward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).PortForward(ctx, req.(*PortForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_CancelPortForward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPortForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).CancelPortForward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/CancelPortForward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).CancelPortForward(ctx, req.(*CancelPortForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).ListNamespaces(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_ForceFrontendUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).ForceFrontendUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/ForceFrontendUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).ForceFrontendUpdate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dashboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Dashboard",
	HandlerType: (*DashboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Dashboard_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Dashboard_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Dashboard_Update_Handler,
		},
		{
			MethodName: "PortForward",
			Handler:    _Dashboard_PortForward_Handler,
		},
		{
			MethodName: "CancelPortForward",
			Handler:    _Dashboard_CancelPortForward_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _Dashboard_ListNamespaces_Handler,
		},
		{
			MethodName: "ForceFrontendUpdate",
			Handler:    _Dashboard_ForceFrontendUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dashboard.proto",
}
