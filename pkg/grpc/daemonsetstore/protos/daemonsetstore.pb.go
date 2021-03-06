// Code generated by protoc-gen-go.
// source: pkg/grpc/daemonsetstore/protos/daemonsetstore.proto
// DO NOT EDIT!

/*
Package daemonsetstore is a generated protocol buffer package.

It is generated from these files:
	pkg/grpc/daemonsetstore/protos/daemonsetstore.proto

It has these top-level messages:
	DaemonSet
	ListDaemonSetsRequest
	ListDaemonSetsResponse
	DisableDaemonSetRequest
	DisableDaemonSetResponse
	WatchDaemonSetsRequest
	WatchDaemonSetsResponse
*/
package daemonsetstore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// models fields/DaemonSet
type DaemonSet struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Disabled     bool   `protobuf:"varint,2,opt,name=disabled" json:"disabled,omitempty"`
	Manifest     string `protobuf:"bytes,3,opt,name=manifest" json:"manifest,omitempty"`
	MinHealth    int64  `protobuf:"varint,4,opt,name=min_health,json=minHealth" json:"min_health,omitempty"`
	Name         string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	NodeSelector string `protobuf:"bytes,6,opt,name=node_selector,json=nodeSelector" json:"node_selector,omitempty"`
	PodId        string `protobuf:"bytes,7,opt,name=pod_id,json=podId" json:"pod_id,omitempty"`
	// expressed in nanoseconds (matches time.Duration)
	Timeout int64 `protobuf:"varint,8,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *DaemonSet) Reset()                    { *m = DaemonSet{} }
func (m *DaemonSet) String() string            { return proto.CompactTextString(m) }
func (*DaemonSet) ProtoMessage()               {}
func (*DaemonSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DaemonSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DaemonSet) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *DaemonSet) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *DaemonSet) GetMinHealth() int64 {
	if m != nil {
		return m.MinHealth
	}
	return 0
}

func (m *DaemonSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DaemonSet) GetNodeSelector() string {
	if m != nil {
		return m.NodeSelector
	}
	return ""
}

func (m *DaemonSet) GetPodId() string {
	if m != nil {
		return m.PodId
	}
	return ""
}

func (m *DaemonSet) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type ListDaemonSetsRequest struct {
}

func (m *ListDaemonSetsRequest) Reset()                    { *m = ListDaemonSetsRequest{} }
func (m *ListDaemonSetsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDaemonSetsRequest) ProtoMessage()               {}
func (*ListDaemonSetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListDaemonSetsResponse struct {
	DaemonSets []*DaemonSet `protobuf:"bytes,1,rep,name=daemon_sets,json=daemonSets" json:"daemon_sets,omitempty"`
}

func (m *ListDaemonSetsResponse) Reset()                    { *m = ListDaemonSetsResponse{} }
func (m *ListDaemonSetsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDaemonSetsResponse) ProtoMessage()               {}
func (*ListDaemonSetsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListDaemonSetsResponse) GetDaemonSets() []*DaemonSet {
	if m != nil {
		return m.DaemonSets
	}
	return nil
}

type DisableDaemonSetRequest struct {
	DaemonSetId string `protobuf:"bytes,1,opt,name=daemon_set_id,json=daemonSetId" json:"daemon_set_id,omitempty"`
}

func (m *DisableDaemonSetRequest) Reset()                    { *m = DisableDaemonSetRequest{} }
func (m *DisableDaemonSetRequest) String() string            { return proto.CompactTextString(m) }
func (*DisableDaemonSetRequest) ProtoMessage()               {}
func (*DisableDaemonSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DisableDaemonSetRequest) GetDaemonSetId() string {
	if m != nil {
		return m.DaemonSetId
	}
	return ""
}

type DisableDaemonSetResponse struct {
	DaemonSet *DaemonSet `protobuf:"bytes,1,opt,name=daemon_set,json=daemonSet" json:"daemon_set,omitempty"`
}

func (m *DisableDaemonSetResponse) Reset()                    { *m = DisableDaemonSetResponse{} }
func (m *DisableDaemonSetResponse) String() string            { return proto.CompactTextString(m) }
func (*DisableDaemonSetResponse) ProtoMessage()               {}
func (*DisableDaemonSetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DisableDaemonSetResponse) GetDaemonSet() *DaemonSet {
	if m != nil {
		return m.DaemonSet
	}
	return nil
}

type WatchDaemonSetsRequest struct {
}

func (m *WatchDaemonSetsRequest) Reset()                    { *m = WatchDaemonSetsRequest{} }
func (m *WatchDaemonSetsRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchDaemonSetsRequest) ProtoMessage()               {}
func (*WatchDaemonSetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// models dsstore.WatchedDaemonSets
type WatchDaemonSetsResponse struct {
	Created []*DaemonSet `protobuf:"bytes,1,rep,name=created" json:"created,omitempty"`
	Updated []*DaemonSet `protobuf:"bytes,2,rep,name=updated" json:"updated,omitempty"`
	Deleted []*DaemonSet `protobuf:"bytes,3,rep,name=deleted" json:"deleted,omitempty"`
	Error   string       `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *WatchDaemonSetsResponse) Reset()                    { *m = WatchDaemonSetsResponse{} }
func (m *WatchDaemonSetsResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchDaemonSetsResponse) ProtoMessage()               {}
func (*WatchDaemonSetsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WatchDaemonSetsResponse) GetCreated() []*DaemonSet {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *WatchDaemonSetsResponse) GetUpdated() []*DaemonSet {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *WatchDaemonSetsResponse) GetDeleted() []*DaemonSet {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func (m *WatchDaemonSetsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*DaemonSet)(nil), "daemonsetstore.DaemonSet")
	proto.RegisterType((*ListDaemonSetsRequest)(nil), "daemonsetstore.ListDaemonSetsRequest")
	proto.RegisterType((*ListDaemonSetsResponse)(nil), "daemonsetstore.ListDaemonSetsResponse")
	proto.RegisterType((*DisableDaemonSetRequest)(nil), "daemonsetstore.DisableDaemonSetRequest")
	proto.RegisterType((*DisableDaemonSetResponse)(nil), "daemonsetstore.DisableDaemonSetResponse")
	proto.RegisterType((*WatchDaemonSetsRequest)(nil), "daemonsetstore.WatchDaemonSetsRequest")
	proto.RegisterType((*WatchDaemonSetsResponse)(nil), "daemonsetstore.WatchDaemonSetsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for P2DaemonSetStore service

type P2DaemonSetStoreClient interface {
	ListDaemonSets(ctx context.Context, in *ListDaemonSetsRequest, opts ...grpc.CallOption) (*ListDaemonSetsResponse, error)
	DisableDaemonSet(ctx context.Context, in *DisableDaemonSetRequest, opts ...grpc.CallOption) (*DisableDaemonSetResponse, error)
	WatchDaemonSets(ctx context.Context, in *WatchDaemonSetsRequest, opts ...grpc.CallOption) (P2DaemonSetStore_WatchDaemonSetsClient, error)
}

type p2DaemonSetStoreClient struct {
	cc *grpc.ClientConn
}

func NewP2DaemonSetStoreClient(cc *grpc.ClientConn) P2DaemonSetStoreClient {
	return &p2DaemonSetStoreClient{cc}
}

func (c *p2DaemonSetStoreClient) ListDaemonSets(ctx context.Context, in *ListDaemonSetsRequest, opts ...grpc.CallOption) (*ListDaemonSetsResponse, error) {
	out := new(ListDaemonSetsResponse)
	err := grpc.Invoke(ctx, "/daemonsetstore.P2DaemonSetStore/ListDaemonSets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2DaemonSetStoreClient) DisableDaemonSet(ctx context.Context, in *DisableDaemonSetRequest, opts ...grpc.CallOption) (*DisableDaemonSetResponse, error) {
	out := new(DisableDaemonSetResponse)
	err := grpc.Invoke(ctx, "/daemonsetstore.P2DaemonSetStore/DisableDaemonSet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2DaemonSetStoreClient) WatchDaemonSets(ctx context.Context, in *WatchDaemonSetsRequest, opts ...grpc.CallOption) (P2DaemonSetStore_WatchDaemonSetsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_P2DaemonSetStore_serviceDesc.Streams[0], c.cc, "/daemonsetstore.P2DaemonSetStore/WatchDaemonSets", opts...)
	if err != nil {
		return nil, err
	}
	x := &p2DaemonSetStoreWatchDaemonSetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type P2DaemonSetStore_WatchDaemonSetsClient interface {
	Recv() (*WatchDaemonSetsResponse, error)
	grpc.ClientStream
}

type p2DaemonSetStoreWatchDaemonSetsClient struct {
	grpc.ClientStream
}

func (x *p2DaemonSetStoreWatchDaemonSetsClient) Recv() (*WatchDaemonSetsResponse, error) {
	m := new(WatchDaemonSetsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for P2DaemonSetStore service

type P2DaemonSetStoreServer interface {
	ListDaemonSets(context.Context, *ListDaemonSetsRequest) (*ListDaemonSetsResponse, error)
	DisableDaemonSet(context.Context, *DisableDaemonSetRequest) (*DisableDaemonSetResponse, error)
	WatchDaemonSets(*WatchDaemonSetsRequest, P2DaemonSetStore_WatchDaemonSetsServer) error
}

func RegisterP2DaemonSetStoreServer(s *grpc.Server, srv P2DaemonSetStoreServer) {
	s.RegisterService(&_P2DaemonSetStore_serviceDesc, srv)
}

func _P2DaemonSetStore_ListDaemonSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDaemonSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2DaemonSetStoreServer).ListDaemonSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemonsetstore.P2DaemonSetStore/ListDaemonSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2DaemonSetStoreServer).ListDaemonSets(ctx, req.(*ListDaemonSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2DaemonSetStore_DisableDaemonSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableDaemonSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2DaemonSetStoreServer).DisableDaemonSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemonsetstore.P2DaemonSetStore/DisableDaemonSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2DaemonSetStoreServer).DisableDaemonSet(ctx, req.(*DisableDaemonSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2DaemonSetStore_WatchDaemonSets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchDaemonSetsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(P2DaemonSetStoreServer).WatchDaemonSets(m, &p2DaemonSetStoreWatchDaemonSetsServer{stream})
}

type P2DaemonSetStore_WatchDaemonSetsServer interface {
	Send(*WatchDaemonSetsResponse) error
	grpc.ServerStream
}

type p2DaemonSetStoreWatchDaemonSetsServer struct {
	grpc.ServerStream
}

func (x *p2DaemonSetStoreWatchDaemonSetsServer) Send(m *WatchDaemonSetsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _P2DaemonSetStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "daemonsetstore.P2DaemonSetStore",
	HandlerType: (*P2DaemonSetStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDaemonSets",
			Handler:    _P2DaemonSetStore_ListDaemonSets_Handler,
		},
		{
			MethodName: "DisableDaemonSet",
			Handler:    _P2DaemonSetStore_DisableDaemonSet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchDaemonSets",
			Handler:       _P2DaemonSetStore_WatchDaemonSets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/grpc/daemonsetstore/protos/daemonsetstore.proto",
}

func init() {
	proto.RegisterFile("pkg/grpc/daemonsetstore/protos/daemonsetstore.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0xd3, 0xfc, 0xf1, 0x84, 0x86, 0x6a, 0x44, 0x9b, 0x25, 0x12, 0x52, 0x64, 0x04,
	0xf5, 0xa9, 0x41, 0xc9, 0x05, 0x21, 0x71, 0xeb, 0x81, 0x4a, 0x1c, 0x90, 0x53, 0x89, 0xa3, 0xe5,
	0x66, 0xa7, 0xc9, 0x0a, 0xdb, 0x6b, 0xbc, 0x9b, 0x57, 0xe4, 0xc8, 0x13, 0xf0, 0x30, 0xc8, 0x6b,
	0x67, 0xa3, 0xd8, 0x11, 0xbe, 0x65, 0xbe, 0xf9, 0x7e, 0x3b, 0xeb, 0x6f, 0x56, 0x81, 0x55, 0xfe,
	0x73, 0xbb, 0xd8, 0x16, 0xf9, 0x66, 0xc1, 0x63, 0x4a, 0x65, 0xa6, 0x48, 0x2b, 0x2d, 0x0b, 0x5a,
	0xe4, 0x85, 0xd4, 0x52, 0x35, 0xd4, 0x3b, 0xa3, 0xe2, 0xe4, 0x54, 0xf5, 0xff, 0x3a, 0xe0, 0xdd,
	0x1b, 0x69, 0x4d, 0x1a, 0x27, 0xe0, 0x0a, 0xce, 0x9c, 0xb9, 0x13, 0x78, 0xa1, 0x2b, 0x38, 0xce,
	0x60, 0xc4, 0x85, 0x8a, 0x9f, 0x12, 0xe2, 0xcc, 0x9d, 0x3b, 0xc1, 0x28, 0xb4, 0x75, 0xd9, 0x4b,
	0xe3, 0x4c, 0x3c, 0x93, 0xd2, 0xac, 0x67, 0x08, 0x5b, 0xe3, 0x5b, 0x80, 0x54, 0x64, 0xd1, 0x8e,
	0xe2, 0x44, 0xef, 0xd8, 0xc5, 0xdc, 0x09, 0x7a, 0xa1, 0x97, 0x8a, 0xec, 0xab, 0x11, 0x10, 0xe1,
	0x22, 0x8b, 0x53, 0x62, 0x7d, 0x83, 0x99, 0xdf, 0xf8, 0x0e, 0x2e, 0x33, 0xc9, 0x29, 0x52, 0x94,
	0xd0, 0x46, 0xcb, 0x82, 0x0d, 0x4c, 0xf3, 0x65, 0x29, 0xae, 0x6b, 0x0d, 0xaf, 0x61, 0x90, 0x4b,
	0x1e, 0x09, 0xce, 0x86, 0xa6, 0xdb, 0xcf, 0x25, 0x7f, 0xe0, 0xc8, 0x60, 0xa8, 0x45, 0x4a, 0x72,
	0xaf, 0xd9, 0xc8, 0xcc, 0x3a, 0x94, 0xfe, 0x14, 0xae, 0xbf, 0x09, 0xa5, 0xed, 0x17, 0xaa, 0x90,
	0x7e, 0xed, 0x49, 0x69, 0xff, 0x11, 0x6e, 0x9a, 0x0d, 0x95, 0x97, 0xb9, 0xe0, 0x67, 0x18, 0x57,
	0x19, 0x45, 0x65, 0x48, 0xcc, 0x99, 0xf7, 0x82, 0xf1, 0xf2, 0xcd, 0x5d, 0x23, 0x4d, 0x0b, 0x86,
	0xc0, 0xed, 0x19, 0xfe, 0x17, 0x98, 0xde, 0x57, 0xf9, 0x1c, 0xfb, 0xd5, 0x40, 0xf4, 0xe1, 0xf2,
	0x78, 0x6c, 0x64, 0x53, 0x1e, 0x5b, 0xfa, 0x81, 0xfb, 0x8f, 0xc0, 0xda, 0x78, 0x7d, 0xad, 0x4f,
	0x00, 0x47, 0xde, 0xc0, 0xff, 0xbd, 0x95, 0x67, 0xcf, 0xf5, 0x19, 0xdc, 0xfc, 0x88, 0xf5, 0x66,
	0xd7, 0x0e, 0xe1, 0x8f, 0x03, 0xd3, 0x56, 0xab, 0x9e, 0xb7, 0x82, 0xe1, 0xa6, 0xa0, 0x58, 0x13,
	0xef, 0x8e, 0xe0, 0xe0, 0x2c, 0xa1, 0x7d, 0xce, 0x0d, 0xe4, 0x76, 0x42, 0xb5, 0xb3, 0x84, 0x38,
	0x25, 0x54, 0x42, 0xbd, 0x4e, 0xa8, 0x76, 0xe2, 0x6b, 0xe8, 0x53, 0x51, 0xc8, 0xc2, 0x3c, 0x2e,
	0x2f, 0xac, 0x8a, 0xe5, 0x6f, 0x17, 0xae, 0xbe, 0x2f, 0xad, 0x7d, 0x5d, 0xd2, 0x18, 0xc3, 0xe4,
	0x74, 0xd5, 0xf8, 0xbe, 0x39, 0xe0, 0xec, 0x1b, 0x99, 0x7d, 0xe8, 0xb2, 0x55, 0x51, 0xf9, 0x2f,
	0x70, 0x0b, 0x57, 0xcd, 0xc5, 0xe1, 0x6d, 0xeb, 0x2b, 0xce, 0xbf, 0x8c, 0x59, 0xd0, 0x6d, 0xb4,
	0x83, 0x9e, 0xe1, 0x55, 0x63, 0x61, 0xd8, 0xba, 0xe5, 0xf9, 0x65, 0xcf, 0x6e, 0x3b, 0x7d, 0x87,
	0x29, 0x1f, 0x9d, 0xa7, 0x81, 0xf9, 0xb7, 0x58, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x24,
	0x1b, 0x55, 0x64, 0x04, 0x00, 0x00,
}
