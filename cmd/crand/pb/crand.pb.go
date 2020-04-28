// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: crand.proto

package pb

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

type RandomnessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *RandomnessRequest) Reset() {
	*x = RandomnessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomnessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomnessRequest) ProtoMessage() {}

func (x *RandomnessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomnessRequest.ProtoReflect.Descriptor instead.
func (*RandomnessRequest) Descriptor() ([]byte, []int) {
	return file_crand_proto_rawDescGZIP(), []int{0}
}

func (x *RandomnessRequest) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

// The response message containing randomness
type RandomnessReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Randomness []byte `protobuf:"bytes,1,opt,name=randomness,proto3" json:"randomness,omitempty"`
	Round      uint64 `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *RandomnessReply) Reset() {
	*x = RandomnessReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomnessReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomnessReply) ProtoMessage() {}

func (x *RandomnessReply) ProtoReflect() protoreflect.Message {
	mi := &file_crand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomnessReply.ProtoReflect.Descriptor instead.
func (*RandomnessReply) Descriptor() ([]byte, []int) {
	return file_crand_proto_rawDescGZIP(), []int{1}
}

func (x *RandomnessReply) GetRandomness() []byte {
	if x != nil {
		return x.Randomness
	}
	return nil
}

func (x *RandomnessReply) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_crand_proto_rawDescGZIP(), []int{2}
}

type InfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pubkey    []byte `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	GenesisTs int64  `protobuf:"varint,2,opt,name=genesis_ts,json=genesisTs,proto3" json:"genesis_ts,omitempty"`
	Round     int64  `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *InfoReply) Reset() {
	*x = InfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReply) ProtoMessage() {}

func (x *InfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_crand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReply.ProtoReflect.Descriptor instead.
func (*InfoReply) Descriptor() ([]byte, []int) {
	return file_crand_proto_rawDescGZIP(), []int{3}
}

func (x *InfoReply) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *InfoReply) GetGenesisTs() int64 {
	if x != nil {
		return x.GenesisTs
	}
	return 0
}

func (x *InfoReply) GetRound() int64 {
	if x != nil {
		return x.Round
	}
	return 0
}

var File_crand_proto protoreflect.FileDescriptor

var file_crand_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x72, 0x61, 0x6e, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x22,
	0x47, 0x0a, 0x0f, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x32, 0x7f, 0x0a, 0x05, 0x43, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x72,
	0x61, 0x6e, 0x64, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x63, 0x72, 0x61,
	0x6e, 0x64, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x63, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x72, 0x61, 0x6e,
	0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crand_proto_rawDescOnce sync.Once
	file_crand_proto_rawDescData = file_crand_proto_rawDesc
)

func file_crand_proto_rawDescGZIP() []byte {
	file_crand_proto_rawDescOnce.Do(func() {
		file_crand_proto_rawDescData = protoimpl.X.CompressGZIP(file_crand_proto_rawDescData)
	})
	return file_crand_proto_rawDescData
}

var file_crand_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crand_proto_goTypes = []interface{}{
	(*RandomnessRequest)(nil), // 0: crand.RandomnessRequest
	(*RandomnessReply)(nil),   // 1: crand.RandomnessReply
	(*InfoRequest)(nil),       // 2: crand.InfoRequest
	(*InfoReply)(nil),         // 3: crand.InfoReply
}
var file_crand_proto_depIdxs = []int32{
	0, // 0: crand.Crand.GetRandomness:input_type -> crand.RandomnessRequest
	2, // 1: crand.Crand.GetInfo:input_type -> crand.InfoRequest
	1, // 2: crand.Crand.GetRandomness:output_type -> crand.RandomnessReply
	3, // 3: crand.Crand.GetInfo:output_type -> crand.InfoReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crand_proto_init() }
func file_crand_proto_init() {
	if File_crand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomnessRequest); i {
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
		file_crand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomnessReply); i {
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
		file_crand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_crand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoReply); i {
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
			RawDescriptor: file_crand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crand_proto_goTypes,
		DependencyIndexes: file_crand_proto_depIdxs,
		MessageInfos:      file_crand_proto_msgTypes,
	}.Build()
	File_crand_proto = out.File
	file_crand_proto_rawDesc = nil
	file_crand_proto_goTypes = nil
	file_crand_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrandClient is the client API for Crand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrandClient interface {
	GetRandomness(ctx context.Context, in *RandomnessRequest, opts ...grpc.CallOption) (*RandomnessReply, error)
	GetInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
}

type crandClient struct {
	cc grpc.ClientConnInterface
}

func NewCrandClient(cc grpc.ClientConnInterface) CrandClient {
	return &crandClient{cc}
}

func (c *crandClient) GetRandomness(ctx context.Context, in *RandomnessRequest, opts ...grpc.CallOption) (*RandomnessReply, error) {
	out := new(RandomnessReply)
	err := c.cc.Invoke(ctx, "/crand.Crand/GetRandomness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crandClient) GetInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := c.cc.Invoke(ctx, "/crand.Crand/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrandServer is the server API for Crand service.
type CrandServer interface {
	GetRandomness(context.Context, *RandomnessRequest) (*RandomnessReply, error)
	GetInfo(context.Context, *InfoRequest) (*InfoReply, error)
}

// UnimplementedCrandServer can be embedded to have forward compatible implementations.
type UnimplementedCrandServer struct {
}

func (*UnimplementedCrandServer) GetRandomness(context.Context, *RandomnessRequest) (*RandomnessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomness not implemented")
}
func (*UnimplementedCrandServer) GetInfo(context.Context, *InfoRequest) (*InfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}

func RegisterCrandServer(s *grpc.Server, srv CrandServer) {
	s.RegisterService(&_Crand_serviceDesc, srv)
}

func _Crand_GetRandomness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomnessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrandServer).GetRandomness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crand.Crand/GetRandomness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrandServer).GetRandomness(ctx, req.(*RandomnessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crand_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrandServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crand.Crand/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrandServer).GetInfo(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crand.Crand",
	HandlerType: (*CrandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandomness",
			Handler:    _Crand_GetRandomness_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Crand_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crand.proto",
}