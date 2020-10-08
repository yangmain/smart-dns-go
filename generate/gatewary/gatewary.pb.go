// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: gatewary.proto

package gatewary

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

type DnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // dns报文
	Ip      string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`           // 对方IP
}

func (x *DnsRequest) Reset() {
	*x = DnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsRequest) ProtoMessage() {}

func (x *DnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatewary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsRequest.ProtoReflect.Descriptor instead.
func (*DnsRequest) Descriptor() ([]byte, []int) {
	return file_gatewary_proto_rawDescGZIP(), []int{0}
}

func (x *DnsRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DnsRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type DnsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // response dns报文
}

func (x *DnsResponse) Reset() {
	*x = DnsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatewary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsResponse) ProtoMessage() {}

func (x *DnsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatewary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsResponse.ProtoReflect.Descriptor instead.
func (*DnsResponse) Descriptor() ([]byte, []int) {
	return file_gatewary_proto_rawDescGZIP(), []int{1}
}

func (x *DnsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_gatewary_proto protoreflect.FileDescriptor

var file_gatewary_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x72, 0x61, 0x79, 0x22, 0x36, 0x0a, 0x0a, 0x44, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x22, 0x27, 0x0a, 0x0b, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4b, 0x0a, 0x0f, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x44, 0x4e, 0x53, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x14, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x72, 0x61, 0x79, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x72, 0x61, 0x79, 0x2e, 0x44, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gatewary_proto_rawDescOnce sync.Once
	file_gatewary_proto_rawDescData = file_gatewary_proto_rawDesc
)

func file_gatewary_proto_rawDescGZIP() []byte {
	file_gatewary_proto_rawDescOnce.Do(func() {
		file_gatewary_proto_rawDescData = protoimpl.X.CompressGZIP(file_gatewary_proto_rawDescData)
	})
	return file_gatewary_proto_rawDescData
}

var file_gatewary_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gatewary_proto_goTypes = []interface{}{
	(*DnsRequest)(nil),  // 0: gatewray.DnsRequest
	(*DnsResponse)(nil), // 1: gatewray.DnsResponse
}
var file_gatewary_proto_depIdxs = []int32{
	0, // 0: gatewray.GatewaryService.DNSLookup:input_type -> gatewray.DnsRequest
	1, // 1: gatewray.GatewaryService.DNSLookup:output_type -> gatewray.DnsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gatewary_proto_init() }
func file_gatewary_proto_init() {
	if File_gatewary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gatewary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsRequest); i {
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
		file_gatewary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsResponse); i {
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
			RawDescriptor: file_gatewary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gatewary_proto_goTypes,
		DependencyIndexes: file_gatewary_proto_depIdxs,
		MessageInfos:      file_gatewary_proto_msgTypes,
	}.Build()
	File_gatewary_proto = out.File
	file_gatewary_proto_rawDesc = nil
	file_gatewary_proto_goTypes = nil
	file_gatewary_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewaryServiceClient is the client API for GatewaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewaryServiceClient interface {
	DNSLookup(ctx context.Context, in *DnsRequest, opts ...grpc.CallOption) (*DnsResponse, error)
}

type gatewaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewaryServiceClient(cc grpc.ClientConnInterface) GatewaryServiceClient {
	return &gatewaryServiceClient{cc}
}

func (c *gatewaryServiceClient) DNSLookup(ctx context.Context, in *DnsRequest, opts ...grpc.CallOption) (*DnsResponse, error) {
	out := new(DnsResponse)
	err := c.cc.Invoke(ctx, "/gatewray.GatewaryService/DNSLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewaryServiceServer is the server API for GatewaryService service.
type GatewaryServiceServer interface {
	DNSLookup(context.Context, *DnsRequest) (*DnsResponse, error)
}

// UnimplementedGatewaryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewaryServiceServer struct {
}

func (*UnimplementedGatewaryServiceServer) DNSLookup(context.Context, *DnsRequest) (*DnsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DNSLookup not implemented")
}

func RegisterGatewaryServiceServer(s *grpc.Server, srv GatewaryServiceServer) {
	s.RegisterService(&_GatewaryService_serviceDesc, srv)
}

func _GatewaryService_DNSLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaryServiceServer).DNSLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatewray.GatewaryService/DNSLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaryServiceServer).DNSLookup(ctx, req.(*DnsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewaryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gatewray.GatewaryService",
	HandlerType: (*GatewaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DNSLookup",
			Handler:    _GatewaryService_DNSLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gatewary.proto",
}