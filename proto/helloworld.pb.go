// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/helloworld.proto

package proto

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

type HelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloWorldRequest) Reset() {
	*x = HelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldRequest) ProtoMessage() {}

func (x *HelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldRequest.ProtoReflect.Descriptor instead.
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloWorldResponse) Reset() {
	*x = HelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldResponse) ProtoMessage() {}

func (x *HelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldResponse.ProtoReflect.Descriptor instead.
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_proto_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_helloworld_proto protoreflect.FileDescriptor

var file_proto_helloworld_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2e, 0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x4e, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x12, 0x12, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_helloworld_proto_rawDescOnce sync.Once
	file_proto_helloworld_proto_rawDescData = file_proto_helloworld_proto_rawDesc
)

func file_proto_helloworld_proto_rawDescGZIP() []byte {
	file_proto_helloworld_proto_rawDescOnce.Do(func() {
		file_proto_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_helloworld_proto_rawDescData)
	})
	return file_proto_helloworld_proto_rawDescData
}

var file_proto_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_helloworld_proto_goTypes = []interface{}{
	(*HelloWorldRequest)(nil),  // 0: HelloWorldRequest
	(*HelloWorldResponse)(nil), // 1: HelloWorldResponse
}
var file_proto_helloworld_proto_depIdxs = []int32{
	0, // 0: HelloWorldService.HelloWorld:input_type -> HelloWorldRequest
	1, // 1: HelloWorldService.HelloWorld:output_type -> HelloWorldResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_helloworld_proto_init() }
func file_proto_helloworld_proto_init() {
	if File_proto_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldRequest); i {
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
		file_proto_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldResponse); i {
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
			RawDescriptor: file_proto_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_helloworld_proto_goTypes,
		DependencyIndexes: file_proto_helloworld_proto_depIdxs,
		MessageInfos:      file_proto_helloworld_proto_msgTypes,
	}.Build()
	File_proto_helloworld_proto = out.File
	file_proto_helloworld_proto_rawDesc = nil
	file_proto_helloworld_proto_goTypes = nil
	file_proto_helloworld_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloWorldServiceClient is the client API for HelloWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldServiceClient interface {
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (HelloWorldService_HelloWorldClient, error)
}

type helloWorldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldServiceClient(cc grpc.ClientConnInterface) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (HelloWorldService_HelloWorldClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloWorldService_serviceDesc.Streams[0], "/HelloWorldService/HelloWorld", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldServiceHelloWorldClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloWorldService_HelloWorldClient interface {
	Recv() (*HelloWorldResponse, error)
	grpc.ClientStream
}

type helloWorldServiceHelloWorldClient struct {
	grpc.ClientStream
}

func (x *helloWorldServiceHelloWorldClient) Recv() (*HelloWorldResponse, error) {
	m := new(HelloWorldResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloWorldServiceServer is the server API for HelloWorldService service.
type HelloWorldServiceServer interface {
	HelloWorld(*HelloWorldRequest, HelloWorldService_HelloWorldServer) error
}

// UnimplementedHelloWorldServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServiceServer struct {
}

func (*UnimplementedHelloWorldServiceServer) HelloWorld(*HelloWorldRequest, HelloWorldService_HelloWorldServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}

func RegisterHelloWorldServiceServer(s *grpc.Server, srv HelloWorldServiceServer) {
	s.RegisterService(&_HelloWorldService_serviceDesc, srv)
}

func _HelloWorldService_HelloWorld_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloWorldRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloWorldServiceServer).HelloWorld(m, &helloWorldServiceHelloWorldServer{stream})
}

type HelloWorldService_HelloWorldServer interface {
	Send(*HelloWorldResponse) error
	grpc.ServerStream
}

type helloWorldServiceHelloWorldServer struct {
	grpc.ServerStream
}

func (x *helloWorldServiceHelloWorldServer) Send(m *HelloWorldResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _HelloWorldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloWorld",
			Handler:       _HelloWorldService_HelloWorld_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/helloworld.proto",
}
