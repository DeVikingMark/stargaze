// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/cron/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("stargaze/cron/query.proto", fileDescriptor_a5073a528095074b) }

var fileDescriptor_a5073a528095074b = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xb1, 0x4a, 0xc0, 0x30,
	0x14, 0x45, 0xdb, 0x41, 0x85, 0x4e, 0x22, 0x2e, 0x16, 0xc9, 0x07, 0x08, 0xe6, 0xd1, 0xba, 0x74,
	0x76, 0x76, 0x71, 0x75, 0x7b, 0x09, 0x8f, 0x18, 0x68, 0xf3, 0x62, 0x92, 0x56, 0xeb, 0x57, 0xf8,
	0x59, 0x8e, 0x1d, 0x1d, 0xa5, 0xfd, 0x11, 0x69, 0xa3, 0x6e, 0x17, 0xce, 0xe1, 0xc2, 0xa9, 0xae,
	0x62, 0xc2, 0x60, 0xf0, 0x9d, 0x40, 0x07, 0x76, 0xf0, 0x32, 0x52, 0x98, 0xa5, 0x0f, 0x9c, 0xf8,
	0xe2, 0xfc, 0x0f, 0xc9, 0xa9, 0x93, 0x3b, 0xad, 0x2f, 0x0d, 0x1b, 0x3e, 0x20, 0xec, 0x2b, 0x7b,
	0xf5, 0xb5, 0x61, 0x36, 0x3d, 0x01, 0x7a, 0x0b, 0xe8, 0x1c, 0x27, 0x4c, 0x96, 0x5d, 0xfc, 0xa5,
	0x37, 0x9a, 0xe3, 0xc0, 0x11, 0x14, 0x46, 0xca, 0xf7, 0x30, 0x35, 0x8a, 0x12, 0x36, 0xe0, 0xd1,
	0x58, 0x77, 0xc8, 0xd9, 0x6d, 0xcf, 0xaa, 0x93, 0xc7, 0xdd, 0xb8, 0x7f, 0xf8, 0x5c, 0x45, 0xb9,
	0xac, 0xa2, 0xfc, 0x5e, 0x45, 0xf9, 0xb1, 0x89, 0x62, 0xd9, 0x44, 0xf1, 0xb5, 0x89, 0xe2, 0xa9,
	0x35, 0x36, 0x3d, 0x8f, 0x4a, 0x6a, 0x1e, 0xc0, 0x8f, 0xaa, 0xb7, 0xfa, 0x16, 0x5f, 0x29, 0xf2,
	0x40, 0xf0, 0x5f, 0x32, 0x75, 0xf0, 0x96, 0x73, 0xd2, 0xec, 0x29, 0xaa, 0xd3, 0xe3, 0xfd, 0xee,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x11, 0x52, 0xe3, 0xec, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stargaze.v8.cron.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "stargaze/cron/query.proto",
}
