// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.9.0
// source: peers/peerpb/peer.proto

package peerpb

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

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peers_peerpb_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_peers_peerpb_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_peers_peerpb_peer_proto_rawDescGZIP(), []int{0}
}

func (x *GetReq) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GetReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peers_peerpb_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_peers_peerpb_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_peers_peerpb_peer_proto_rawDescGZIP(), []int{1}
}

func (x *GetResp) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_peers_peerpb_peer_proto protoreflect.FileDescriptor

var file_peers_peerpb_peer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x70,
	0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x22, 0x30, 0x0a, 0x06, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x1f, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x39, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12,
	0x0d, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0f,
	0x5a, 0x0d, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peers_peerpb_peer_proto_rawDescOnce sync.Once
	file_peers_peerpb_peer_proto_rawDescData = file_peers_peerpb_peer_proto_rawDesc
)

func file_peers_peerpb_peer_proto_rawDescGZIP() []byte {
	file_peers_peerpb_peer_proto_rawDescOnce.Do(func() {
		file_peers_peerpb_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_peers_peerpb_peer_proto_rawDescData)
	})
	return file_peers_peerpb_peer_proto_rawDescData
}

var file_peers_peerpb_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_peers_peerpb_peer_proto_goTypes = []interface{}{
	(*GetReq)(nil),  // 0: peers.getReq
	(*GetResp)(nil), // 1: peers.getResp
}
var file_peers_peerpb_peer_proto_depIdxs = []int32{
	0, // 0: peers.access_cache.GetCache:input_type -> peers.getReq
	1, // 1: peers.access_cache.GetCache:output_type -> peers.getResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_peers_peerpb_peer_proto_init() }
func file_peers_peerpb_peer_proto_init() {
	if File_peers_peerpb_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peers_peerpb_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_peers_peerpb_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
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
			RawDescriptor: file_peers_peerpb_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peers_peerpb_peer_proto_goTypes,
		DependencyIndexes: file_peers_peerpb_peer_proto_depIdxs,
		MessageInfos:      file_peers_peerpb_peer_proto_msgTypes,
	}.Build()
	File_peers_peerpb_peer_proto = out.File
	file_peers_peerpb_peer_proto_rawDesc = nil
	file_peers_peerpb_peer_proto_goTypes = nil
	file_peers_peerpb_peer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccessCacheClient is the client API for AccessCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessCacheClient interface {
	GetCache(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type accessCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessCacheClient(cc grpc.ClientConnInterface) AccessCacheClient {
	return &accessCacheClient{cc}
}

func (c *accessCacheClient) GetCache(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/peers.access_cache/GetCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessCacheServer is the server API for AccessCache service.
type AccessCacheServer interface {
	GetCache(context.Context, *GetReq) (*GetResp, error)
}

// UnimplementedAccessCacheServer can be embedded to have forward compatible implementations.
type UnimplementedAccessCacheServer struct {
}

func (*UnimplementedAccessCacheServer) GetCache(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCache not implemented")
}

func RegisterAccessCacheServer(s *grpc.Server, srv AccessCacheServer) {
	s.RegisterService(&_AccessCache_serviceDesc, srv)
}

func _AccessCache_GetCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessCacheServer).GetCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peers.access_cache/GetCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessCacheServer).GetCache(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessCache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peers.access_cache",
	HandlerType: (*AccessCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCache",
			Handler:    _AccessCache_GetCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peers/peerpb/peer.proto",
}