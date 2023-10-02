// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: honey_badger.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Data_Set_FullMethodName              = "/hb.Data/Set"
	Data_Get_FullMethodName              = "/hb.Data/Get"
	Data_Delete_FullMethodName           = "/hb.Data/Delete"
	Data_DeleteByPrefix_FullMethodName   = "/hb.Data/DeleteByPrefix"
	Data_CreateReadStream_FullMethodName = "/hb.Data/CreateReadStream"
	Data_CreateSendStream_FullMethodName = "/hb.Data/CreateSendStream"
)

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*EmptyResult, error)
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*GetResult, error)
	Delete(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*EmptyResult, error)
	DeleteByPrefix(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*EmptyResult, error)
	CreateReadStream(ctx context.Context, in *ReadStreamReq, opts ...grpc.CallOption) (Data_CreateReadStreamClient, error)
	CreateSendStream(ctx context.Context, opts ...grpc.CallOption) (Data_CreateSendStreamClient, error)
}

type dataClient struct {
	cc grpc.ClientConnInterface
}

func NewDataClient(cc grpc.ClientConnInterface) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*EmptyResult, error) {
	out := new(EmptyResult)
	err := c.cc.Invoke(ctx, Data_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*GetResult, error) {
	out := new(GetResult)
	err := c.cc.Invoke(ctx, Data_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Delete(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*EmptyResult, error) {
	out := new(EmptyResult)
	err := c.cc.Invoke(ctx, Data_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) DeleteByPrefix(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*EmptyResult, error) {
	out := new(EmptyResult)
	err := c.cc.Invoke(ctx, Data_DeleteByPrefix_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) CreateReadStream(ctx context.Context, in *ReadStreamReq, opts ...grpc.CallOption) (Data_CreateReadStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Data_ServiceDesc.Streams[0], Data_CreateReadStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataCreateReadStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Data_CreateReadStreamClient interface {
	Recv() (*DataItem, error)
	grpc.ClientStream
}

type dataCreateReadStreamClient struct {
	grpc.ClientStream
}

func (x *dataCreateReadStreamClient) Recv() (*DataItem, error) {
	m := new(DataItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataClient) CreateSendStream(ctx context.Context, opts ...grpc.CallOption) (Data_CreateSendStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Data_ServiceDesc.Streams[1], Data_CreateSendStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataCreateSendStreamClient{stream}
	return x, nil
}

type Data_CreateSendStreamClient interface {
	Send(*SendStreamReq) error
	CloseAndRecv() (*EmptyResult, error)
	grpc.ClientStream
}

type dataCreateSendStreamClient struct {
	grpc.ClientStream
}

func (x *dataCreateSendStreamClient) Send(m *SendStreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataCreateSendStreamClient) CloseAndRecv() (*EmptyResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServer is the server API for Data service.
// All implementations must embed UnimplementedDataServer
// for forward compatibility
type DataServer interface {
	Set(context.Context, *SetRequest) (*EmptyResult, error)
	Get(context.Context, *KeyRequest) (*GetResult, error)
	Delete(context.Context, *KeyRequest) (*EmptyResult, error)
	DeleteByPrefix(context.Context, *PrefixRequest) (*EmptyResult, error)
	CreateReadStream(*ReadStreamReq, Data_CreateReadStreamServer) error
	CreateSendStream(Data_CreateSendStreamServer) error
	mustEmbedUnimplementedDataServer()
}

// UnimplementedDataServer must be embedded to have forward compatible implementations.
type UnimplementedDataServer struct {
}

func (UnimplementedDataServer) Set(context.Context, *SetRequest) (*EmptyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedDataServer) Get(context.Context, *KeyRequest) (*GetResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDataServer) Delete(context.Context, *KeyRequest) (*EmptyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDataServer) DeleteByPrefix(context.Context, *PrefixRequest) (*EmptyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByPrefix not implemented")
}
func (UnimplementedDataServer) CreateReadStream(*ReadStreamReq, Data_CreateReadStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateReadStream not implemented")
}
func (UnimplementedDataServer) CreateSendStream(Data_CreateSendStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateSendStream not implemented")
}
func (UnimplementedDataServer) mustEmbedUnimplementedDataServer() {}

// UnsafeDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServer will
// result in compilation errors.
type UnsafeDataServer interface {
	mustEmbedUnimplementedDataServer()
}

func RegisterDataServer(s grpc.ServiceRegistrar, srv DataServer) {
	s.RegisterService(&Data_ServiceDesc, srv)
}

func _Data_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Get(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Delete(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_DeleteByPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).DeleteByPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_DeleteByPrefix_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).DeleteByPrefix(ctx, req.(*PrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_CreateReadStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadStreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServer).CreateReadStream(m, &dataCreateReadStreamServer{stream})
}

type Data_CreateReadStreamServer interface {
	Send(*DataItem) error
	grpc.ServerStream
}

type dataCreateReadStreamServer struct {
	grpc.ServerStream
}

func (x *dataCreateReadStreamServer) Send(m *DataItem) error {
	return x.ServerStream.SendMsg(m)
}

func _Data_CreateSendStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServer).CreateSendStream(&dataCreateSendStreamServer{stream})
}

type Data_CreateSendStreamServer interface {
	SendAndClose(*EmptyResult) error
	Recv() (*SendStreamReq, error)
	grpc.ServerStream
}

type dataCreateSendStreamServer struct {
	grpc.ServerStream
}

func (x *dataCreateSendStreamServer) SendAndClose(m *EmptyResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataCreateSendStreamServer) Recv() (*SendStreamReq, error) {
	m := new(SendStreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Data_ServiceDesc is the grpc.ServiceDesc for Data service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Data_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hb.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Data_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Data_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Data_Delete_Handler,
		},
		{
			MethodName: "DeleteByPrefix",
			Handler:    _Data_DeleteByPrefix_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateReadStream",
			Handler:       _Data_CreateReadStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateSendStream",
			Handler:       _Data_CreateSendStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "honey_badger.proto",
}

const (
	Db_Create_FullMethodName = "/hb.Db/Create"
	Db_Drop_FullMethodName   = "/hb.Db/Drop"
)

// DbClient is the client API for Db service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DbClient interface {
	Create(ctx context.Context, in *CreateDbRequest, opts ...grpc.CallOption) (*EmptyResult, error)
	Drop(ctx context.Context, in *DropDbRequest, opts ...grpc.CallOption) (*EmptyResult, error)
}

type dbClient struct {
	cc grpc.ClientConnInterface
}

func NewDbClient(cc grpc.ClientConnInterface) DbClient {
	return &dbClient{cc}
}

func (c *dbClient) Create(ctx context.Context, in *CreateDbRequest, opts ...grpc.CallOption) (*EmptyResult, error) {
	out := new(EmptyResult)
	err := c.cc.Invoke(ctx, Db_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) Drop(ctx context.Context, in *DropDbRequest, opts ...grpc.CallOption) (*EmptyResult, error) {
	out := new(EmptyResult)
	err := c.cc.Invoke(ctx, Db_Drop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DbServer is the server API for Db service.
// All implementations must embed UnimplementedDbServer
// for forward compatibility
type DbServer interface {
	Create(context.Context, *CreateDbRequest) (*EmptyResult, error)
	Drop(context.Context, *DropDbRequest) (*EmptyResult, error)
	mustEmbedUnimplementedDbServer()
}

// UnimplementedDbServer must be embedded to have forward compatible implementations.
type UnimplementedDbServer struct {
}

func (UnimplementedDbServer) Create(context.Context, *CreateDbRequest) (*EmptyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDbServer) Drop(context.Context, *DropDbRequest) (*EmptyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Drop not implemented")
}
func (UnimplementedDbServer) mustEmbedUnimplementedDbServer() {}

// UnsafeDbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DbServer will
// result in compilation errors.
type UnsafeDbServer interface {
	mustEmbedUnimplementedDbServer()
}

func RegisterDbServer(s grpc.ServiceRegistrar, srv DbServer) {
	s.RegisterService(&Db_ServiceDesc, srv)
}

func _Db_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).Create(ctx, req.(*CreateDbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_Drop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropDbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).Drop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_Drop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).Drop(ctx, req.(*DropDbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Db_ServiceDesc is the grpc.ServiceDesc for Db service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Db_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hb.Db",
	HandlerType: (*DbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Db_Create_Handler,
		},
		{
			MethodName: "Drop",
			Handler:    _Db_Drop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "honey_badger.proto",
}

const (
	Sys_Ping_FullMethodName = "/hb.Sys/Ping"
)

// SysClient is the client API for Sys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResult, error)
}

type sysClient struct {
	cc grpc.ClientConnInterface
}

func NewSysClient(cc grpc.ClientConnInterface) SysClient {
	return &sysClient{cc}
}

func (c *sysClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResult, error) {
	out := new(PingResult)
	err := c.cc.Invoke(ctx, Sys_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysServer is the server API for Sys service.
// All implementations must embed UnimplementedSysServer
// for forward compatibility
type SysServer interface {
	Ping(context.Context, *PingRequest) (*PingResult, error)
	mustEmbedUnimplementedSysServer()
}

// UnimplementedSysServer must be embedded to have forward compatible implementations.
type UnimplementedSysServer struct {
}

func (UnimplementedSysServer) Ping(context.Context, *PingRequest) (*PingResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSysServer) mustEmbedUnimplementedSysServer() {}

// UnsafeSysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysServer will
// result in compilation errors.
type UnsafeSysServer interface {
	mustEmbedUnimplementedSysServer()
}

func RegisterSysServer(s grpc.ServiceRegistrar, srv SysServer) {
	s.RegisterService(&Sys_ServiceDesc, srv)
}

func _Sys_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sys_ServiceDesc is the grpc.ServiceDesc for Sys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hb.Sys",
	HandlerType: (*SysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Sys_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "honey_badger.proto",
}