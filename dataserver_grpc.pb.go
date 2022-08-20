// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dataserver

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

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	CreateShard(ctx context.Context, in *CreateShardRequest, opts ...grpc.CallOption) (*CreateShardResponse, error)
	DeleteShard(ctx context.Context, in *DeleteShardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ShardInfo(ctx context.Context, in *ShardInfoRequest, opts ...grpc.CallOption) (*ShardInfoResponse, error)
	SplitShard(ctx context.Context, in *SplitShardRequest, opts ...grpc.CallOption) (*SplitShardResponse, error)
	MergeShard(ctx context.Context, in *MergeShardRequest, opts ...grpc.CallOption) (*MergeShardResponse, error)
	// rpc PullShardData(PullShardDataRequest) returns (PullShardDataResponse);
	// rpc AddShardReplica(AddShardReplicaRequest) returns (AddShardReplicaResponse);
	// rpc DeleteShardReplica(DeleteShardReplicaRequest) returns (DeleteShardReplicaResponse);
	TransferShardLeader(ctx context.Context, in *TransferShardLeaderRequest, opts ...grpc.CallOption) (*TransferShardLeaderResponse, error)
	ShardRead(ctx context.Context, in *ShardReadRequest, opts ...grpc.CallOption) (*ShardReadResponse, error)
	ShardWrite(ctx context.Context, in *ShardWriteRequest, opts ...grpc.CallOption) (*ShardWriteResponse, error)
	ShardAppendLog(ctx context.Context, in *ShardAppendLogRequest, opts ...grpc.CallOption) (*ShardAppendLogResponse, error)
	ShardInstallSnapshot(ctx context.Context, opts ...grpc.CallOption) (DataService_ShardInstallSnapshotClient, error)
	MigrateShard(ctx context.Context, opts ...grpc.CallOption) (DataService_MigrateShardClient, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) CreateShard(ctx context.Context, in *CreateShardRequest, opts ...grpc.CallOption) (*CreateShardResponse, error) {
	out := new(CreateShardResponse)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/CreateShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) DeleteShard(ctx context.Context, in *DeleteShardRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/DeleteShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) ShardInfo(ctx context.Context, in *ShardInfoRequest, opts ...grpc.CallOption) (*ShardInfoResponse, error) {
	out := new(ShardInfoResponse)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/ShardInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) SplitShard(ctx context.Context, in *SplitShardRequest, opts ...grpc.CallOption) (*SplitShardResponse, error) {
	out := new(SplitShardResponse)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/SplitShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) MergeShard(ctx context.Context, in *MergeShardRequest, opts ...grpc.CallOption) (*MergeShardResponse, error) {
	out := new(MergeShardResponse)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/MergeShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) TransferShardLeader(ctx context.Context, in *TransferShardLeaderRequest, opts ...grpc.CallOption) (*TransferShardLeaderResponse, error) {
	out := new(TransferShardLeaderResponse)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/TransferShardLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) ShardRead(ctx context.Context, in *ShardReadRequest, opts ...grpc.CallOption) (*ShardReadResponse, error) {
	out := new(ShardReadResponse)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/ShardRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) ShardWrite(ctx context.Context, in *ShardWriteRequest, opts ...grpc.CallOption) (*ShardWriteResponse, error) {
	out := new(ShardWriteResponse)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/ShardWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) ShardAppendLog(ctx context.Context, in *ShardAppendLogRequest, opts ...grpc.CallOption) (*ShardAppendLogResponse, error) {
	out := new(ShardAppendLogResponse)
	err := c.cc.Invoke(ctx, "/bedrock.dataserver.DataService/ShardAppendLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) ShardInstallSnapshot(ctx context.Context, opts ...grpc.CallOption) (DataService_ShardInstallSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[0], "/bedrock.dataserver.DataService/ShardInstallSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceShardInstallSnapshotClient{stream}
	return x, nil
}

type DataService_ShardInstallSnapshotClient interface {
	Send(*ShardInstallSnapshotRequest) error
	CloseAndRecv() (*ShardInstallSnapshotResponse, error)
	grpc.ClientStream
}

type dataServiceShardInstallSnapshotClient struct {
	grpc.ClientStream
}

func (x *dataServiceShardInstallSnapshotClient) Send(m *ShardInstallSnapshotRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataServiceShardInstallSnapshotClient) CloseAndRecv() (*ShardInstallSnapshotResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ShardInstallSnapshotResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) MigrateShard(ctx context.Context, opts ...grpc.CallOption) (DataService_MigrateShardClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[1], "/bedrock.dataserver.DataService/MigrateShard", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceMigrateShardClient{stream}
	return x, nil
}

type DataService_MigrateShardClient interface {
	Send(*MigrateShardRequest) error
	CloseAndRecv() (*MigrateShardResponse, error)
	grpc.ClientStream
}

type dataServiceMigrateShardClient struct {
	grpc.ClientStream
}

func (x *dataServiceMigrateShardClient) Send(m *MigrateShardRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataServiceMigrateShardClient) CloseAndRecv() (*MigrateShardResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MigrateShardResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	CreateShard(context.Context, *CreateShardRequest) (*CreateShardResponse, error)
	DeleteShard(context.Context, *DeleteShardRequest) (*emptypb.Empty, error)
	ShardInfo(context.Context, *ShardInfoRequest) (*ShardInfoResponse, error)
	SplitShard(context.Context, *SplitShardRequest) (*SplitShardResponse, error)
	MergeShard(context.Context, *MergeShardRequest) (*MergeShardResponse, error)
	// rpc PullShardData(PullShardDataRequest) returns (PullShardDataResponse);
	// rpc AddShardReplica(AddShardReplicaRequest) returns (AddShardReplicaResponse);
	// rpc DeleteShardReplica(DeleteShardReplicaRequest) returns (DeleteShardReplicaResponse);
	TransferShardLeader(context.Context, *TransferShardLeaderRequest) (*TransferShardLeaderResponse, error)
	ShardRead(context.Context, *ShardReadRequest) (*ShardReadResponse, error)
	ShardWrite(context.Context, *ShardWriteRequest) (*ShardWriteResponse, error)
	ShardAppendLog(context.Context, *ShardAppendLogRequest) (*ShardAppendLogResponse, error)
	ShardInstallSnapshot(DataService_ShardInstallSnapshotServer) error
	MigrateShard(DataService_MigrateShardServer) error
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) CreateShard(context.Context, *CreateShardRequest) (*CreateShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShard not implemented")
}
func (UnimplementedDataServiceServer) DeleteShard(context.Context, *DeleteShardRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShard not implemented")
}
func (UnimplementedDataServiceServer) ShardInfo(context.Context, *ShardInfoRequest) (*ShardInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShardInfo not implemented")
}
func (UnimplementedDataServiceServer) SplitShard(context.Context, *SplitShardRequest) (*SplitShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SplitShard not implemented")
}
func (UnimplementedDataServiceServer) MergeShard(context.Context, *MergeShardRequest) (*MergeShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeShard not implemented")
}
func (UnimplementedDataServiceServer) TransferShardLeader(context.Context, *TransferShardLeaderRequest) (*TransferShardLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferShardLeader not implemented")
}
func (UnimplementedDataServiceServer) ShardRead(context.Context, *ShardReadRequest) (*ShardReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShardRead not implemented")
}
func (UnimplementedDataServiceServer) ShardWrite(context.Context, *ShardWriteRequest) (*ShardWriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShardWrite not implemented")
}
func (UnimplementedDataServiceServer) ShardAppendLog(context.Context, *ShardAppendLogRequest) (*ShardAppendLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShardAppendLog not implemented")
}
func (UnimplementedDataServiceServer) ShardInstallSnapshot(DataService_ShardInstallSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method ShardInstallSnapshot not implemented")
}
func (UnimplementedDataServiceServer) MigrateShard(DataService_MigrateShardServer) error {
	return status.Errorf(codes.Unimplemented, "method MigrateShard not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_CreateShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).CreateShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/CreateShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).CreateShard(ctx, req.(*CreateShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_DeleteShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).DeleteShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/DeleteShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).DeleteShard(ctx, req.(*DeleteShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_ShardInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).ShardInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/ShardInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).ShardInfo(ctx, req.(*ShardInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_SplitShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SplitShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).SplitShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/SplitShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).SplitShard(ctx, req.(*SplitShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_MergeShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).MergeShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/MergeShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).MergeShard(ctx, req.(*MergeShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_TransferShardLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferShardLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).TransferShardLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/TransferShardLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).TransferShardLeader(ctx, req.(*TransferShardLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_ShardRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).ShardRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/ShardRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).ShardRead(ctx, req.(*ShardReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_ShardWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).ShardWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/ShardWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).ShardWrite(ctx, req.(*ShardWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_ShardAppendLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardAppendLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).ShardAppendLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bedrock.dataserver.DataService/ShardAppendLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).ShardAppendLog(ctx, req.(*ShardAppendLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_ShardInstallSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServiceServer).ShardInstallSnapshot(&dataServiceShardInstallSnapshotServer{stream})
}

type DataService_ShardInstallSnapshotServer interface {
	SendAndClose(*ShardInstallSnapshotResponse) error
	Recv() (*ShardInstallSnapshotRequest, error)
	grpc.ServerStream
}

type dataServiceShardInstallSnapshotServer struct {
	grpc.ServerStream
}

func (x *dataServiceShardInstallSnapshotServer) SendAndClose(m *ShardInstallSnapshotResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataServiceShardInstallSnapshotServer) Recv() (*ShardInstallSnapshotRequest, error) {
	m := new(ShardInstallSnapshotRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataService_MigrateShard_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServiceServer).MigrateShard(&dataServiceMigrateShardServer{stream})
}

type DataService_MigrateShardServer interface {
	SendAndClose(*MigrateShardResponse) error
	Recv() (*MigrateShardRequest, error)
	grpc.ServerStream
}

type dataServiceMigrateShardServer struct {
	grpc.ServerStream
}

func (x *dataServiceMigrateShardServer) SendAndClose(m *MigrateShardResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataServiceMigrateShardServer) Recv() (*MigrateShardRequest, error) {
	m := new(MigrateShardRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bedrock.dataserver.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShard",
			Handler:    _DataService_CreateShard_Handler,
		},
		{
			MethodName: "DeleteShard",
			Handler:    _DataService_DeleteShard_Handler,
		},
		{
			MethodName: "ShardInfo",
			Handler:    _DataService_ShardInfo_Handler,
		},
		{
			MethodName: "SplitShard",
			Handler:    _DataService_SplitShard_Handler,
		},
		{
			MethodName: "MergeShard",
			Handler:    _DataService_MergeShard_Handler,
		},
		{
			MethodName: "TransferShardLeader",
			Handler:    _DataService_TransferShardLeader_Handler,
		},
		{
			MethodName: "ShardRead",
			Handler:    _DataService_ShardRead_Handler,
		},
		{
			MethodName: "ShardWrite",
			Handler:    _DataService_ShardWrite_Handler,
		},
		{
			MethodName: "ShardAppendLog",
			Handler:    _DataService_ShardAppendLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShardInstallSnapshot",
			Handler:       _DataService_ShardInstallSnapshot_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "MigrateShard",
			Handler:       _DataService_MigrateShard_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "dataserver.proto",
}
