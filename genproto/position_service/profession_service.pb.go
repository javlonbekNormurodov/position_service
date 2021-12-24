// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: profession_service.proto

package position_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_profession_service_proto protoreflect.FileDescriptor

var file_profession_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdd, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x66, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x21, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_profession_service_proto_goTypes = []interface{}{
	(*CreateProfession)(nil),         // 0: genproto.CreateProfession
	(*ProfessionId)(nil),             // 1: genproto.ProfessionId
	(*GetAllProfessionRequest)(nil),  // 2: genproto.GetAllProfessionRequest
	(*Profession)(nil),               // 3: genproto.Profession
	(*GetAllProfessionResponse)(nil), // 4: genproto.GetAllProfessionResponse
}
var file_profession_service_proto_depIdxs = []int32{
	0, // 0: genproto.ProfessionService.Create:input_type -> genproto.CreateProfession
	1, // 1: genproto.ProfessionService.Get:input_type -> genproto.ProfessionId
	2, // 2: genproto.ProfessionService.GetAll:input_type -> genproto.GetAllProfessionRequest
	1, // 3: genproto.ProfessionService.Create:output_type -> genproto.ProfessionId
	3, // 4: genproto.ProfessionService.Get:output_type -> genproto.Profession
	4, // 5: genproto.ProfessionService.GetAll:output_type -> genproto.GetAllProfessionResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_profession_service_proto_init() }
func file_profession_service_proto_init() {
	if File_profession_service_proto != nil {
		return
	}
	file_profession_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profession_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profession_service_proto_goTypes,
		DependencyIndexes: file_profession_service_proto_depIdxs,
	}.Build()
	File_profession_service_proto = out.File
	file_profession_service_proto_rawDesc = nil
	file_profession_service_proto_goTypes = nil
	file_profession_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProfessionServiceClient is the client API for ProfessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfessionServiceClient interface {
	Create(ctx context.Context, in *CreateProfession, opts ...grpc.CallOption) (*ProfessionId, error)
	Get(ctx context.Context, in *ProfessionId, opts ...grpc.CallOption) (*Profession, error)
	GetAll(ctx context.Context, in *GetAllProfessionRequest, opts ...grpc.CallOption) (*GetAllProfessionResponse, error)
}

type professionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfessionServiceClient(cc grpc.ClientConnInterface) ProfessionServiceClient {
	return &professionServiceClient{cc}
}

func (c *professionServiceClient) Create(ctx context.Context, in *CreateProfession, opts ...grpc.CallOption) (*ProfessionId, error) {
	out := new(ProfessionId)
	err := c.cc.Invoke(ctx, "/genproto.ProfessionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionServiceClient) Get(ctx context.Context, in *ProfessionId, opts ...grpc.CallOption) (*Profession, error) {
	out := new(Profession)
	err := c.cc.Invoke(ctx, "/genproto.ProfessionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionServiceClient) GetAll(ctx context.Context, in *GetAllProfessionRequest, opts ...grpc.CallOption) (*GetAllProfessionResponse, error) {
	out := new(GetAllProfessionResponse)
	err := c.cc.Invoke(ctx, "/genproto.ProfessionService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfessionServiceServer is the server API for ProfessionService service.
type ProfessionServiceServer interface {
	Create(context.Context, *CreateProfession) (*ProfessionId, error)
	Get(context.Context, *ProfessionId) (*Profession, error)
	GetAll(context.Context, *GetAllProfessionRequest) (*GetAllProfessionResponse, error)
}

// UnimplementedProfessionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfessionServiceServer struct {
}

func (*UnimplementedProfessionServiceServer) Create(context.Context, *CreateProfession) (*ProfessionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedProfessionServiceServer) Get(context.Context, *ProfessionId) (*Profession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedProfessionServiceServer) GetAll(context.Context, *GetAllProfessionRequest) (*GetAllProfessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterProfessionServiceServer(s *grpc.Server, srv ProfessionServiceServer) {
	s.RegisterService(&_ProfessionService_serviceDesc, srv)
}

func _ProfessionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ProfessionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).Create(ctx, req.(*CreateProfession))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfessionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ProfessionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).Get(ctx, req.(*ProfessionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProfessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ProfessionService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionServiceServer).GetAll(ctx, req.(*GetAllProfessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ProfessionService",
	HandlerType: (*ProfessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProfessionService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProfessionService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ProfessionService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profession_service.proto",
}