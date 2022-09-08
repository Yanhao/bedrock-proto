// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetaServiceClient is the client API for MetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaServiceClient interface {
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetShardRoutes(ctx context.Context, in *GetShardRoutesRequest, opts ...grpc.CallOption) (*GetShardRoutesResponse, error)
	ShardInfo(ctx context.Context, in *ShardInfoRequest, opts ...grpc.CallOption) (*ShardInfoResponse, error)
	CreateShard(ctx context.Context, in *CreateShardRequest, opts ...grpc.CallOption) (*CreateShardResponse, error)
	RemoveShard(ctx context.Context, in *RemoveShardRequest, opts ...grpc.CallOption) (*RemoveShardResponse, error)
	// rpc ChangeShardReplicateCount(ChangeShardReplicateCountRequest) returns(ChangeShardReplicateCountResponse);
	GetShardIDByKey(ctx context.Context, in *GetShardIDByKeyRequest, opts ...grpc.CallOption) (*GetShardIDByKeyResponse, error)
	CreateStorage(ctx context.Context, in *CreateStorageRequest, opts ...grpc.CallOption) (*CreateStorageResponse, error)
	DeleteStorage(ctx context.Context, in *DeleteStorageRequest, opts ...grpc.CallOption) (*DeleteStorageResponse, error)
	UndeleteStorage(ctx context.Context, in *UndeleteStorageRequest, opts ...grpc.CallOption) (*UndeleteStorageResponse, error)
	RenameStorage(ctx context.Context, in *RenameStorageRequest, opts ...grpc.CallOption) (*RenameStorageResponse, error)
	ResizeStorage(ctx context.Context, in *ResizeStorageRequest, opts ...grpc.CallOption) (*ResizeStorageResponse, error)
	GetStorages(ctx context.Context, in *GetStoragesRequest, opts ...grpc.CallOption) (*GetStoragesResponse, error)
	AddDataServer(ctx context.Context, in *AddDataServerRequest, opts ...grpc.CallOption) (*AddDataServerResponse, error)
	RemoveDataServer(ctx context.Context, in *RemoveDataServerRequest, opts ...grpc.CallOption) (*RemoveDataServerResponse, error)
	ListDataServer(ctx context.Context, in *ListDataServerRequest, opts ...grpc.CallOption) (*ListDataServerResponse, error)
	UpdateDataServer(ctx context.Context, in *UpdateDataServerRequest, opts ...grpc.CallOption) (*UpdateDataServerResponse, error)
	SyncShardInDataServer(ctx context.Context, opts ...grpc.CallOption) (MetaService_SyncShardInDataServerClient, error)
}

type metaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaServiceClient(cc grpc.ClientConnInterface) MetaServiceClient {
	return &metaServiceClient{cc}
}

func (c *metaServiceClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) GetShardRoutes(ctx context.Context, in *GetShardRoutesRequest, opts ...grpc.CallOption) (*GetShardRoutesResponse, error) {
	out := new(GetShardRoutesResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/GetShardRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) ShardInfo(ctx context.Context, in *ShardInfoRequest, opts ...grpc.CallOption) (*ShardInfoResponse, error) {
	out := new(ShardInfoResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/ShardInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) CreateShard(ctx context.Context, in *CreateShardRequest, opts ...grpc.CallOption) (*CreateShardResponse, error) {
	out := new(CreateShardResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/CreateShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) RemoveShard(ctx context.Context, in *RemoveShardRequest, opts ...grpc.CallOption) (*RemoveShardResponse, error) {
	out := new(RemoveShardResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/RemoveShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) GetShardIDByKey(ctx context.Context, in *GetShardIDByKeyRequest, opts ...grpc.CallOption) (*GetShardIDByKeyResponse, error) {
	out := new(GetShardIDByKeyResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/GetShardIDByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) CreateStorage(ctx context.Context, in *CreateStorageRequest, opts ...grpc.CallOption) (*CreateStorageResponse, error) {
	out := new(CreateStorageResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/CreateStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) DeleteStorage(ctx context.Context, in *DeleteStorageRequest, opts ...grpc.CallOption) (*DeleteStorageResponse, error) {
	out := new(DeleteStorageResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/DeleteStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) UndeleteStorage(ctx context.Context, in *UndeleteStorageRequest, opts ...grpc.CallOption) (*UndeleteStorageResponse, error) {
	out := new(UndeleteStorageResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/UndeleteStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) RenameStorage(ctx context.Context, in *RenameStorageRequest, opts ...grpc.CallOption) (*RenameStorageResponse, error) {
	out := new(RenameStorageResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/RenameStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) ResizeStorage(ctx context.Context, in *ResizeStorageRequest, opts ...grpc.CallOption) (*ResizeStorageResponse, error) {
	out := new(ResizeStorageResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/ResizeStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) GetStorages(ctx context.Context, in *GetStoragesRequest, opts ...grpc.CallOption) (*GetStoragesResponse, error) {
	out := new(GetStoragesResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/GetStorages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) AddDataServer(ctx context.Context, in *AddDataServerRequest, opts ...grpc.CallOption) (*AddDataServerResponse, error) {
	out := new(AddDataServerResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/AddDataServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) RemoveDataServer(ctx context.Context, in *RemoveDataServerRequest, opts ...grpc.CallOption) (*RemoveDataServerResponse, error) {
	out := new(RemoveDataServerResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/RemoveDataServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) ListDataServer(ctx context.Context, in *ListDataServerRequest, opts ...grpc.CallOption) (*ListDataServerResponse, error) {
	out := new(ListDataServerResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/ListDataServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) UpdateDataServer(ctx context.Context, in *UpdateDataServerRequest, opts ...grpc.CallOption) (*UpdateDataServerResponse, error) {
	out := new(UpdateDataServerResponse)
	err := c.cc.Invoke(ctx, "/bedrock.metaserver.MetaService/UpdateDataServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) SyncShardInDataServer(ctx context.Context, opts ...grpc.CallOption) (MetaService_SyncShardInDataServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetaService_ServiceDesc.Streams[0], "/bedrock.metaserver.MetaService/SyncShardInDataServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &metaServiceSyncShardInDataServerClient{stream}
	return x, nil
}

type MetaService_SyncShardInDataServerClient interface {
	Send(*SyncShardInDataServerRequest) error
	CloseAndRecv() (*SyncShardInDataServerResponse, error)
	grpc.ClientStream
}

type metaServiceSyncShardInDataServerClient struct {
	grpc.ClientStream
}

func (x *metaServiceSyncShardInDataServerClient) Send(m *SyncShardInDataServerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metaServiceSyncShardInDataServerClient) CloseAndRecv() (*SyncShardInDataServerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SyncShardInDataServerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetaServiceServer is the server API for MetaService service.
// All implementations must embed UnimplementedMetaServiceServer
// for forward compatibility
type MetaServiceServer interface {
	HeartBeat(context.Context, *HeartBeatRequest) (*emptypb.Empty, error)
	GetShardRoutes(context.Context, *GetShardRoutesRequest) (*GetShardRoutesResponse, error)
	ShardInfo(context.Context, *ShardInfoRequest) (*ShardInfoResponse, error)
	CreateShard(context.Context, *CreateShardRequest) (*CreateShardResponse, error)
	RemoveShard(context.Context, *RemoveShardRequest) (*RemoveShardResponse, error)
	// rpc ChangeShardReplicateCount(ChangeShardReplicateCountRequest) returns(ChangeShardReplicateCountResponse);
	GetShardIDByKey(context.Context, *GetShardIDByKeyRequest) (*GetShardIDByKeyResponse, error)
	CreateStorage(context.Context, *CreateStorageRequest) (*CreateStorageResponse, error)
	DeleteStorage(context.Context, *DeleteStorageRequest) (*DeleteStorageResponse, error)
	UndeleteStorage(context.Context, *UndeleteStorageRequest) (*UndeleteStorageResponse, error)
	RenameStorage(context.Context, *RenameStorageRequest) (*RenameStorageResponse, error)
	ResizeStorage(context.Context, *ResizeStorageRequest) (*ResizeStorageResponse, error)
	GetStorages(context.Context, *GetStoragesRequest) (*GetStoragesResponse, error)
	AddDataServer(context.Context, *AddDataServerRequest) (*AddDataServerResponse, error)
	RemoveDataServer(context.Context, *RemoveDataServerRequest) (*RemoveDataServerResponse, error)
	ListDataServer(context.Context, *ListDataServerRequest) (*ListDataServerResponse, error)
	UpdateDataServer(context.Context, *UpdateDataServerRequest) (*UpdateDataServerResponse, error)
	SyncShardInDataServer(MetaService_SyncShardInDataServerServer) error
	mustEmbedUnimplementedMetaServiceServer()
}

// UnimplementedMetaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetaServiceServer struct {
}

func (UnimplementedMetaServiceServer) HeartBeat(context.Context, *HeartBeatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedMetaServiceServer) GetShardRoutes(context.Context, *GetShardRoutesRequest) (*GetShardRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShardRoutes not implemented")
}
func (UnimplementedMetaServiceServer) ShardInfo(context.Context, *ShardInfoRequest) (*ShardInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShardInfo not implemented")
}
func (UnimplementedMetaServiceServer) CreateShard(context.Context, *CreateShardRequest) (*CreateShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShard not implemented")
}
func (UnimplementedMetaServiceServer) RemoveShard(context.Context, *RemoveShardRequest) (*RemoveShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveShard not implemented")
}
func (UnimplementedMetaServiceServer) GetShardIDByKey(context.Context, *GetShardIDByKeyRequest) (*GetShardIDByKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShardIDByKey not implemented")
}
func (UnimplementedMetaServiceServer) CreateStorage(context.Context, *CreateStorageRequest) (*CreateStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorage not implemented")
}
func (UnimplementedMetaServiceServer) DeleteStorage(context.Context, *DeleteStorageRequest) (*DeleteStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorage not implemented")
}
func (UnimplementedMetaServiceServer) UndeleteStorage(context.Context, *UndeleteStorageRequest) (*UndeleteStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UndeleteStorage not implemented")
}
func (UnimplementedMetaServiceServer) RenameStorage(context.Context, *RenameStorageRequest) (*RenameStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameStorage not implemented")
}
func (UnimplementedMetaServiceServer) ResizeStorage(context.Context, *ResizeStorageRequest) (*ResizeStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeStorage not implemented")
}
func (UnimplementedMetaServiceServer) GetStorages(context.Context, *GetStoragesRequest) (*GetStoragesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorages not implemented")
}
func (UnimplementedMetaServiceServer) AddDataServer(context.Context, *AddDataServerRequest) (*AddDataServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDataServer not implemented")
}
func (UnimplementedMetaServiceServer) RemoveDataServer(context.Context, *RemoveDataServerRequest) (*RemoveDataServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDataServer not implemented")
}
func (UnimplementedMetaServiceServer) ListDataServer(context.Context, *ListDataServerRequest) (*ListDataServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataServer not implemented")
}
func (UnimplementedMetaServiceServer) UpdateDataServer(context.Context, *UpdateDataServerRequest) (*UpdateDataServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDataServer not implemented")
}
func (UnimplementedMetaServiceServer) SyncShardInDataServer(MetaService_SyncShardInDataServerServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncShardInDataServer not implemented")
}
func (UnimplementedMetaServiceServer) mustEmbedUnimplementedMetaServiceServer() {}

// UnsafeMetaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaServiceServer will
// result in compilation errors.
type UnsafeMetaServiceServer interface {
	mustEmbedUnimplementedMetaServiceServer()
}

func RegisterMetaServiceServer(s grpc.ServiceRegistrar, srv MetaServiceServer) {
	s.RegisterService(&MetaService_ServiceDesc, srv)
}

func _MetaService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_GetShardRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShardRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).GetShardRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/GetShardRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).GetShardRoutes(ctx, req.(*GetShardRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_ShardInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).ShardInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/ShardInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).ShardInfo(ctx, req.(*ShardInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_CreateShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).CreateShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/CreateShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).CreateShard(ctx, req.(*CreateShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_RemoveShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).RemoveShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/RemoveShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).RemoveShard(ctx, req.(*RemoveShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_GetShardIDByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShardIDByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).GetShardIDByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/GetShardIDByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).GetShardIDByKey(ctx, req.(*GetShardIDByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_CreateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).CreateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/CreateStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).CreateStorage(ctx, req.(*CreateStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_DeleteStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).DeleteStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/DeleteStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).DeleteStorage(ctx, req.(*DeleteStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_UndeleteStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeleteStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).UndeleteStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/UndeleteStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).UndeleteStorage(ctx, req.(*UndeleteStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_RenameStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).RenameStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/RenameStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).RenameStorage(ctx, req.(*RenameStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_ResizeStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).ResizeStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/ResizeStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).ResizeStorage(ctx, req.(*ResizeStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_GetStorages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoragesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).GetStorages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/GetStorages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).GetStorages(ctx, req.(*GetStoragesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_AddDataServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDataServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).AddDataServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/AddDataServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).AddDataServer(ctx, req.(*AddDataServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_RemoveDataServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDataServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).RemoveDataServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/RemoveDataServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).RemoveDataServer(ctx, req.(*RemoveDataServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_ListDataServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).ListDataServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/ListDataServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).ListDataServer(ctx, req.(*ListDataServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_UpdateDataServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).UpdateDataServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.metaserver.MetaService/UpdateDataServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).UpdateDataServer(ctx, req.(*UpdateDataServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_SyncShardInDataServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetaServiceServer).SyncShardInDataServer(&metaServiceSyncShardInDataServerServer{stream})
}

type MetaService_SyncShardInDataServerServer interface {
	SendAndClose(*SyncShardInDataServerResponse) error
	Recv() (*SyncShardInDataServerRequest, error)
	grpc.ServerStream
}

type metaServiceSyncShardInDataServerServer struct {
	grpc.ServerStream
}

func (x *metaServiceSyncShardInDataServerServer) SendAndClose(m *SyncShardInDataServerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metaServiceSyncShardInDataServerServer) Recv() (*SyncShardInDataServerRequest, error) {
	m := new(SyncShardInDataServerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetaService_ServiceDesc is the grpc.ServiceDesc for MetaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bedrock.metaserver.MetaService",
	HandlerType: (*MetaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _MetaService_HeartBeat_Handler,
		},
		{
			MethodName: "GetShardRoutes",
			Handler:    _MetaService_GetShardRoutes_Handler,
		},
		{
			MethodName: "ShardInfo",
			Handler:    _MetaService_ShardInfo_Handler,
		},
		{
			MethodName: "CreateShard",
			Handler:    _MetaService_CreateShard_Handler,
		},
		{
			MethodName: "RemoveShard",
			Handler:    _MetaService_RemoveShard_Handler,
		},
		{
			MethodName: "GetShardIDByKey",
			Handler:    _MetaService_GetShardIDByKey_Handler,
		},
		{
			MethodName: "CreateStorage",
			Handler:    _MetaService_CreateStorage_Handler,
		},
		{
			MethodName: "DeleteStorage",
			Handler:    _MetaService_DeleteStorage_Handler,
		},
		{
			MethodName: "UndeleteStorage",
			Handler:    _MetaService_UndeleteStorage_Handler,
		},
		{
			MethodName: "RenameStorage",
			Handler:    _MetaService_RenameStorage_Handler,
		},
		{
			MethodName: "ResizeStorage",
			Handler:    _MetaService_ResizeStorage_Handler,
		},
		{
			MethodName: "GetStorages",
			Handler:    _MetaService_GetStorages_Handler,
		},
		{
			MethodName: "AddDataServer",
			Handler:    _MetaService_AddDataServer_Handler,
		},
		{
			MethodName: "RemoveDataServer",
			Handler:    _MetaService_RemoveDataServer_Handler,
		},
		{
			MethodName: "ListDataServer",
			Handler:    _MetaService_ListDataServer_Handler,
		},
		{
			MethodName: "UpdateDataServer",
			Handler:    _MetaService_UpdateDataServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncShardInDataServer",
			Handler:       _MetaService_SyncShardInDataServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "metaserver.proto",
}
