// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: server_streaming/calculator/calculatorpb/calculator.proto

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

type PrimeFactorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PrimeFactorsRequest) Reset() {
	*x = PrimeFactorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_streaming_calculator_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeFactorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeFactorsRequest) ProtoMessage() {}

func (x *PrimeFactorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_streaming_calculator_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeFactorsRequest.ProtoReflect.Descriptor instead.
func (*PrimeFactorsRequest) Descriptor() ([]byte, []int) {
	return file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeFactorsRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type PrimeFactorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PrimeFactorsResponse) Reset() {
	*x = PrimeFactorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_streaming_calculator_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeFactorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeFactorsResponse) ProtoMessage() {}

func (x *PrimeFactorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_streaming_calculator_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeFactorsResponse.ProtoReflect.Descriptor instead.
func (*PrimeFactorsResponse) Descriptor() ([]byte, []int) {
	return file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeFactorsResponse) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_server_streaming_calculator_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_server_streaming_calculator_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x39, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x22, 0x27, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x6d, 0x65,
	0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x22, 0x28, 0x0a, 0x14, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0x6b, 0x0a, 0x11, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x56, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x72, 0x6f, 0x72, 0x73,
	0x12, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescData = file_server_streaming_calculator_calculatorpb_calculator_proto_rawDesc
)

func file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescData)
	})
	return file_server_streaming_calculator_calculatorpb_calculator_proto_rawDescData
}

var file_server_streaming_calculator_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_streaming_calculator_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*PrimeFactorsRequest)(nil),  // 0: calcualtor.PrimeFactorsRequest
	(*PrimeFactorsResponse)(nil), // 1: calcualtor.PrimeFactorsResponse
}
var file_server_streaming_calculator_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calcualtor.CalculatorService.PrimeFactrors:input_type -> calcualtor.PrimeFactorsRequest
	1, // 1: calcualtor.CalculatorService.PrimeFactrors:output_type -> calcualtor.PrimeFactorsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_streaming_calculator_calculatorpb_calculator_proto_init() }
func file_server_streaming_calculator_calculatorpb_calculator_proto_init() {
	if File_server_streaming_calculator_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_streaming_calculator_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeFactorsRequest); i {
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
		file_server_streaming_calculator_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeFactorsResponse); i {
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
			RawDescriptor: file_server_streaming_calculator_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_streaming_calculator_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_server_streaming_calculator_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_server_streaming_calculator_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_server_streaming_calculator_calculatorpb_calculator_proto = out.File
	file_server_streaming_calculator_calculatorpb_calculator_proto_rawDesc = nil
	file_server_streaming_calculator_calculatorpb_calculator_proto_goTypes = nil
	file_server_streaming_calculator_calculatorpb_calculator_proto_depIdxs = nil
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
	PrimeFactrors(ctx context.Context, in *PrimeFactorsRequest, opts ...grpc.CallOption) (CalculatorService_PrimeFactrorsClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) PrimeFactrors(ctx context.Context, in *PrimeFactorsRequest, opts ...grpc.CallOption) (CalculatorService_PrimeFactrorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calcualtor.CalculatorService/PrimeFactrors", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeFactrorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeFactrorsClient interface {
	Recv() (*PrimeFactorsResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeFactrorsClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeFactrorsClient) Recv() (*PrimeFactorsResponse, error) {
	m := new(PrimeFactorsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	PrimeFactrors(*PrimeFactorsRequest, CalculatorService_PrimeFactrorsServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) PrimeFactrors(*PrimeFactorsRequest, CalculatorService_PrimeFactrorsServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeFactrors not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_PrimeFactrors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeFactorsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeFactrors(m, &calculatorServicePrimeFactrorsServer{stream})
}

type CalculatorService_PrimeFactrorsServer interface {
	Send(*PrimeFactorsResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeFactrorsServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeFactrorsServer) Send(m *PrimeFactorsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calcualtor.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeFactrors",
			Handler:       _CalculatorService_PrimeFactrors_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server_streaming/calculator/calculatorpb/calculator.proto",
}
