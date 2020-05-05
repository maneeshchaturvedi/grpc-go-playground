// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: bidirectional_streaming/calculator/calculatorpb/calculator.proto

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

type FindMaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
}

func (x *FindMaximumRequest) Reset() {
	*x = FindMaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumRequest) ProtoMessage() {}

func (x *FindMaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_msgTypes[0]
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
	return file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *FindMaximumRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int64 `protobuf:"varint,1,opt,name=Maximum,proto3" json:"Maximum,omitempty"`
}

func (x *FindMaximumResponse) Reset() {
	*x = FindMaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumResponse) ProtoMessage() {}

func (x *FindMaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_msgTypes[1]
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
	return file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *FindMaximumResponse) GetMaximum() int64 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_bidirectional_streaming_calculator_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x40, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70,
	0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x2c,
	0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x13,
	0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x32, 0x6a, 0x0a,
	0x12, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6e, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x2f, 0x62, 0x69,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescData = file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDesc
)

func file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescData)
	})
	return file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDescData
}

var file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*FindMaximumRequest)(nil),  // 0: calculator.FindMaximumRequest
	(*FindMaximumResponse)(nil), // 1: calculator.FindMaximumResponse
}
var file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.FindMaximumService.FindMaximun:input_type -> calculator.FindMaximumRequest
	1, // 1: calculator.FindMaximumService.FindMaximun:output_type -> calculator.FindMaximumResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_init() }
func file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_init() {
	if File_bidirectional_streaming_calculator_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_bidirectional_streaming_calculator_calculatorpb_calculator_proto = out.File
	file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_rawDesc = nil
	file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_goTypes = nil
	file_bidirectional_streaming_calculator_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FindMaximumServiceClient is the client API for FindMaximumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FindMaximumServiceClient interface {
	FindMaximun(ctx context.Context, opts ...grpc.CallOption) (FindMaximumService_FindMaximunClient, error)
}

type findMaximumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFindMaximumServiceClient(cc grpc.ClientConnInterface) FindMaximumServiceClient {
	return &findMaximumServiceClient{cc}
}

func (c *findMaximumServiceClient) FindMaximun(ctx context.Context, opts ...grpc.CallOption) (FindMaximumService_FindMaximunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FindMaximumService_serviceDesc.Streams[0], "/calculator.FindMaximumService/FindMaximun", opts...)
	if err != nil {
		return nil, err
	}
	x := &findMaximumServiceFindMaximunClient{stream}
	return x, nil
}

type FindMaximumService_FindMaximunClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type findMaximumServiceFindMaximunClient struct {
	grpc.ClientStream
}

func (x *findMaximumServiceFindMaximunClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *findMaximumServiceFindMaximunClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FindMaximumServiceServer is the server API for FindMaximumService service.
type FindMaximumServiceServer interface {
	FindMaximun(FindMaximumService_FindMaximunServer) error
}

// UnimplementedFindMaximumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFindMaximumServiceServer struct {
}

func (*UnimplementedFindMaximumServiceServer) FindMaximun(FindMaximumService_FindMaximunServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximun not implemented")
}

func RegisterFindMaximumServiceServer(s *grpc.Server, srv FindMaximumServiceServer) {
	s.RegisterService(&_FindMaximumService_serviceDesc, srv)
}

func _FindMaximumService_FindMaximun_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FindMaximumServiceServer).FindMaximun(&findMaximumServiceFindMaximunServer{stream})
}

type FindMaximumService_FindMaximunServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type findMaximumServiceFindMaximunServer struct {
	grpc.ServerStream
}

func (x *findMaximumServiceFindMaximunServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *findMaximumServiceFindMaximunServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FindMaximumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.FindMaximumService",
	HandlerType: (*FindMaximumServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaximun",
			Handler:       _FindMaximumService_FindMaximun_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bidirectional_streaming/calculator/calculatorpb/calculator.proto",
}
