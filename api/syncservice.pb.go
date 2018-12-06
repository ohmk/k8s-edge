// Code generated by protoc-gen-go. DO NOT EDIT.
// source: syncservice.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetEdgeNodeRequest struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEdgeNodeRequest) Reset()         { *m = GetEdgeNodeRequest{} }
func (m *GetEdgeNodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetEdgeNodeRequest) ProtoMessage()    {}
func (*GetEdgeNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37be25d20501358f, []int{0}
}

func (m *GetEdgeNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEdgeNodeRequest.Unmarshal(m, b)
}
func (m *GetEdgeNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEdgeNodeRequest.Marshal(b, m, deterministic)
}
func (m *GetEdgeNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEdgeNodeRequest.Merge(m, src)
}
func (m *GetEdgeNodeRequest) XXX_Size() int {
	return xxx_messageInfo_GetEdgeNodeRequest.Size(m)
}
func (m *GetEdgeNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEdgeNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEdgeNodeRequest proto.InternalMessageInfo

func (m *GetEdgeNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type GetEdgeNodeReply struct {
	EdgeNode             *EdgeNode `protobuf:"bytes,1,opt,name=edge_node,json=edgeNode,proto3" json:"edge_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetEdgeNodeReply) Reset()         { *m = GetEdgeNodeReply{} }
func (m *GetEdgeNodeReply) String() string { return proto.CompactTextString(m) }
func (*GetEdgeNodeReply) ProtoMessage()    {}
func (*GetEdgeNodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_37be25d20501358f, []int{1}
}

func (m *GetEdgeNodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEdgeNodeReply.Unmarshal(m, b)
}
func (m *GetEdgeNodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEdgeNodeReply.Marshal(b, m, deterministic)
}
func (m *GetEdgeNodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEdgeNodeReply.Merge(m, src)
}
func (m *GetEdgeNodeReply) XXX_Size() int {
	return xxx_messageInfo_GetEdgeNodeReply.Size(m)
}
func (m *GetEdgeNodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEdgeNodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetEdgeNodeReply proto.InternalMessageInfo

func (m *GetEdgeNodeReply) GetEdgeNode() *EdgeNode {
	if m != nil {
		return m.EdgeNode
	}
	return nil
}

type EdgeNode struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Synced               bool     `protobuf:"varint,2,opt,name=synced,proto3" json:"synced,omitempty"`
	Pods                 []*Pod   `protobuf:"bytes,3,rep,name=pods,proto3" json:"pods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeNode) Reset()         { *m = EdgeNode{} }
func (m *EdgeNode) String() string { return proto.CompactTextString(m) }
func (*EdgeNode) ProtoMessage()    {}
func (*EdgeNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_37be25d20501358f, []int{2}
}

func (m *EdgeNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeNode.Unmarshal(m, b)
}
func (m *EdgeNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeNode.Marshal(b, m, deterministic)
}
func (m *EdgeNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeNode.Merge(m, src)
}
func (m *EdgeNode) XXX_Size() int {
	return xxx_messageInfo_EdgeNode.Size(m)
}
func (m *EdgeNode) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeNode.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeNode proto.InternalMessageInfo

func (m *EdgeNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EdgeNode) GetSynced() bool {
	if m != nil {
		return m.Synced
	}
	return false
}

func (m *EdgeNode) GetPods() []*Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type Pod struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pod) Reset()         { *m = Pod{} }
func (m *Pod) String() string { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()    {}
func (*Pod) Descriptor() ([]byte, []int) {
	return fileDescriptor_37be25d20501358f, []int{3}
}

func (m *Pod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pod.Unmarshal(m, b)
}
func (m *Pod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pod.Marshal(b, m, deterministic)
}
func (m *Pod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pod.Merge(m, src)
}
func (m *Pod) XXX_Size() int {
	return xxx_messageInfo_Pod.Size(m)
}
func (m *Pod) XXX_DiscardUnknown() {
	xxx_messageInfo_Pod.DiscardUnknown(m)
}

var xxx_messageInfo_Pod proto.InternalMessageInfo

func (m *Pod) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pod) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*GetEdgeNodeRequest)(nil), "api.GetEdgeNodeRequest")
	proto.RegisterType((*GetEdgeNodeReply)(nil), "api.GetEdgeNodeReply")
	proto.RegisterType((*EdgeNode)(nil), "api.EdgeNode")
	proto.RegisterType((*Pod)(nil), "api.Pod")
}

func init() { proto.RegisterFile("syncservice.proto", fileDescriptor_37be25d20501358f) }

var fileDescriptor_37be25d20501358f = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xa9, 0x99, 0x23, 0xbd, 0x55, 0x99, 0x17, 0xff, 0x14, 0xf5, 0xa1, 0xf4, 0xa9, 0x08,
	0x16, 0xac, 0xcf, 0x0a, 0x3e, 0x88, 0xf8, 0x32, 0x4b, 0xe6, 0xfb, 0x88, 0xcd, 0x65, 0x14, 0xbb,
	0x26, 0xae, 0xdd, 0x20, 0xdf, 0x5e, 0x92, 0x55, 0x99, 0xec, 0xed, 0xdc, 0x93, 0x7b, 0x7e, 0x49,
	0x0e, 0x9c, 0x76, 0xb6, 0xad, 0x3a, 0x5a, 0x6d, 0xea, 0x8a, 0x72, 0xb3, 0xd2, 0xbd, 0x46, 0x26,
	0x4d, 0x9d, 0xde, 0x03, 0xbe, 0x52, 0xff, 0xa2, 0x16, 0x34, 0xd5, 0x8a, 0x04, 0x7d, 0xaf, 0xa9,
	0xeb, 0xf1, 0x1a, 0xc2, 0x56, 0x2b, 0x9a, 0xb7, 0x72, 0x49, 0x71, 0x90, 0x04, 0x59, 0x28, 0xb8,
	0x33, 0xa6, 0x72, 0x49, 0xe9, 0x13, 0x4c, 0xfe, 0x45, 0x4c, 0x63, 0xf1, 0x16, 0x42, 0x52, 0x0b,
	0x9a, 0xbb, 0x25, 0x1f, 0x88, 0x8a, 0xe3, 0x5c, 0x9a, 0x3a, 0xff, 0x5b, 0xe3, 0x34, 0xa8, 0xf4,
	0x03, 0xf8, 0xaf, 0x8b, 0x08, 0xa3, 0x9d, 0x3b, 0xbc, 0xc6, 0x0b, 0x18, 0xbb, 0xc7, 0x92, 0x8a,
	0x0f, 0x92, 0x20, 0xe3, 0x62, 0x98, 0xf0, 0x06, 0x46, 0x46, 0xab, 0x2e, 0x66, 0x09, 0xcb, 0xa2,
	0x82, 0x7b, 0x7c, 0xa9, 0x95, 0xf0, 0x6e, 0x7a, 0x07, 0xac, 0xd4, 0x0a, 0x27, 0xc0, 0xbe, 0xc8,
	0x0e, 0x3c, 0x27, 0xf1, 0x0c, 0x0e, 0x37, 0xb2, 0x59, 0x93, 0xa7, 0x1d, 0x89, 0xed, 0x50, 0xbc,
	0xc3, 0xc9, 0xcc, 0xb6, 0xd5, 0x6c, 0xdb, 0xc8, 0x73, 0xf9, 0x86, 0x8f, 0x10, 0xed, 0x7c, 0x0b,
	0x2f, 0x3d, 0x7f, 0xbf, 0x9b, 0xab, 0xf3, 0xfd, 0x03, 0xd3, 0xd8, 0xcf, 0xb1, 0x2f, 0xf5, 0xe1,
	0x27, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x61, 0x99, 0x73, 0x69, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncServiceAPIClient is the client API for SyncServiceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncServiceAPIClient interface {
	// rpc RegisterEdgeNode;
	GetEdgeNode(ctx context.Context, in *GetEdgeNodeRequest, opts ...grpc.CallOption) (*GetEdgeNodeReply, error)
}

type syncServiceAPIClient struct {
	cc *grpc.ClientConn
}

func NewSyncServiceAPIClient(cc *grpc.ClientConn) SyncServiceAPIClient {
	return &syncServiceAPIClient{cc}
}

func (c *syncServiceAPIClient) GetEdgeNode(ctx context.Context, in *GetEdgeNodeRequest, opts ...grpc.CallOption) (*GetEdgeNodeReply, error) {
	out := new(GetEdgeNodeReply)
	err := c.cc.Invoke(ctx, "/api.SyncServiceAPI/GetEdgeNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServiceAPIServer is the server API for SyncServiceAPI service.
type SyncServiceAPIServer interface {
	// rpc RegisterEdgeNode;
	GetEdgeNode(context.Context, *GetEdgeNodeRequest) (*GetEdgeNodeReply, error)
}

func RegisterSyncServiceAPIServer(s *grpc.Server, srv SyncServiceAPIServer) {
	s.RegisterService(&_SyncServiceAPI_serviceDesc, srv)
}

func _SyncServiceAPI_GetEdgeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEdgeNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceAPIServer).GetEdgeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncServiceAPI/GetEdgeNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceAPIServer).GetEdgeNode(ctx, req.(*GetEdgeNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncServiceAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SyncServiceAPI",
	HandlerType: (*SyncServiceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEdgeNode",
			Handler:    _SyncServiceAPI_GetEdgeNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncservice.proto",
}
