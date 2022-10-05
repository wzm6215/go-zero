// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: systemuserget.proto

package systemuserget

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuserget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_systemuserget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_systemuserget_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=nick_name,json=nick_name,proto3" json:"nick_name,omitempty"`
	Avatar   string `protobuf:"bytes,1,opt,name=avatar,json=avatar,proto3" json:"avatar,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuserget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_systemuserget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_systemuserget_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

var File_systemuserget_proto protoreflect.FileDescriptor

var file_systemuserget_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x67, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x75, 0x73, 0x65,
	0x72, 0x67, 0x65, 0x74, 0x22, 0x19, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x27, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x53, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x67, 0x65, 0x74, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0d, 0x67,
	0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x75, 0x73, 0x65,
	0x72, 0x67, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_systemuserget_proto_rawDescOnce sync.Once
	file_systemuserget_proto_rawDescData = file_systemuserget_proto_rawDesc
)

func file_systemuserget_proto_rawDescGZIP() []byte {
	file_systemuserget_proto_rawDescOnce.Do(func() {
		file_systemuserget_proto_rawDescData = protoimpl.X.CompressGZIP(file_systemuserget_proto_rawDescData)
	})
	return file_systemuserget_proto_rawDescData
}

var file_systemuserget_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_systemuserget_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: systemuserget.Request
	(*Response)(nil), // 1: systemuserget.Response
}
var file_systemuserget_proto_depIdxs = []int32{
	0, // 0: systemuserget.Systemusergeter.getSystemuser:input_type -> systemuserget.Request
	1, // 1: systemuserget.Systemusergeter.getSystemuser:output_type -> systemuserget.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_systemuserget_proto_init() }
func file_systemuserget_proto_init() {
	if File_systemuserget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_systemuserget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_systemuserget_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_systemuserget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_systemuserget_proto_goTypes,
		DependencyIndexes: file_systemuserget_proto_depIdxs,
		MessageInfos:      file_systemuserget_proto_msgTypes,
	}.Build()
	File_systemuserget_proto = out.File
	file_systemuserget_proto_rawDesc = nil
	file_systemuserget_proto_goTypes = nil
	file_systemuserget_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SystemusergeterClient is the client API for Systemusergeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SystemusergeterClient interface {
	GetSystemuser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type systemusergeterClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemusergeterClient(cc grpc.ClientConnInterface) SystemusergeterClient {
	return &systemusergeterClient{cc}
}

func (c *systemusergeterClient) GetSystemuser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/systemuserget.Systemusergeter/getSystemuser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemusergeterServer is the server API for Systemusergeter service.
type SystemusergeterServer interface {
	GetSystemuser(context.Context, *Request) (*Response, error)
}

// UnimplementedSystemusergeterServer can be embedded to have forward compatible implementations.
type UnimplementedSystemusergeterServer struct {
}

func (*UnimplementedSystemusergeterServer) GetSystemuser(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemuser not implemented")
}

func RegisterSystemusergeterServer(s *grpc.Server, srv SystemusergeterServer) {
	s.RegisterService(&_Systemusergeter_serviceDesc, srv)
}

func _Systemusergeter_GetSystemuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemusergeterServer).GetSystemuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systemuserget.Systemusergeter/GetSystemuser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemusergeterServer).GetSystemuser(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Systemusergeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "systemuserget.Systemusergeter",
	HandlerType: (*SystemusergeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSystemuser",
			Handler:    _Systemusergeter_GetSystemuser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "systemuserget.proto",
}