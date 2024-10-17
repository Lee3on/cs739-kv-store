// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: proto/raft.proto

package raft

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

// RaftServiceClient is the client API for RaftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftServiceClient interface {
	// Sends a raft message to another node
	SendRaftMessage(ctx context.Context, in *RaftMessage, opts ...grpc.CallOption) (*RaftResponse, error)
	// Streaming version to handle multiple messages
	StreamRaftMessages(ctx context.Context, opts ...grpc.CallOption) (RaftService_StreamRaftMessagesClient, error)
}

type raftServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftServiceClient(cc grpc.ClientConnInterface) RaftServiceClient {
	return &raftServiceClient{cc}
}

func (c *raftServiceClient) SendRaftMessage(ctx context.Context, in *RaftMessage, opts ...grpc.CallOption) (*RaftResponse, error) {
	out := new(RaftResponse)
	err := c.cc.Invoke(ctx, "/raft.RaftService/SendRaftMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) StreamRaftMessages(ctx context.Context, opts ...grpc.CallOption) (RaftService_StreamRaftMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RaftService_ServiceDesc.Streams[0], "/raft.RaftService/StreamRaftMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftServiceStreamRaftMessagesClient{stream}
	return x, nil
}

type RaftService_StreamRaftMessagesClient interface {
	Send(*RaftMessage) error
	Recv() (*RaftResponse, error)
	grpc.ClientStream
}

type raftServiceStreamRaftMessagesClient struct {
	grpc.ClientStream
}

func (x *raftServiceStreamRaftMessagesClient) Send(m *RaftMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *raftServiceStreamRaftMessagesClient) Recv() (*RaftResponse, error) {
	m := new(RaftResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RaftServiceServer is the server API for RaftService service.
// All implementations must embed UnimplementedRaftServiceServer
// for forward compatibility
type RaftServiceServer interface {
	// Sends a raft message to another node
	SendRaftMessage(context.Context, *RaftMessage) (*RaftResponse, error)
	// Streaming version to handle multiple messages
	StreamRaftMessages(RaftService_StreamRaftMessagesServer) error
	mustEmbedUnimplementedRaftServiceServer()
}

// UnimplementedRaftServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRaftServiceServer struct {
}

func (UnimplementedRaftServiceServer) SendRaftMessage(context.Context, *RaftMessage) (*RaftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRaftMessage not implemented")
}
func (UnimplementedRaftServiceServer) StreamRaftMessages(RaftService_StreamRaftMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRaftMessages not implemented")
}
func (UnimplementedRaftServiceServer) mustEmbedUnimplementedRaftServiceServer() {}

// UnsafeRaftServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftServiceServer will
// result in compilation errors.
type UnsafeRaftServiceServer interface {
	mustEmbedUnimplementedRaftServiceServer()
}

func RegisterRaftServiceServer(s grpc.ServiceRegistrar, srv RaftServiceServer) {
	s.RegisterService(&RaftService_ServiceDesc, srv)
}

func _RaftService_SendRaftMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).SendRaftMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftService/SendRaftMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).SendRaftMessage(ctx, req.(*RaftMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_StreamRaftMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RaftServiceServer).StreamRaftMessages(&raftServiceStreamRaftMessagesServer{stream})
}

type RaftService_StreamRaftMessagesServer interface {
	Send(*RaftResponse) error
	Recv() (*RaftMessage, error)
	grpc.ServerStream
}

type raftServiceStreamRaftMessagesServer struct {
	grpc.ServerStream
}

func (x *raftServiceStreamRaftMessagesServer) Send(m *RaftResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *raftServiceStreamRaftMessagesServer) Recv() (*RaftMessage, error) {
	m := new(RaftMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RaftService_ServiceDesc is the grpc.ServiceDesc for RaftService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raft.RaftService",
	HandlerType: (*RaftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRaftMessage",
			Handler:    _RaftService_SendRaftMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRaftMessages",
			Handler:       _RaftService_StreamRaftMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/raft.proto",
}