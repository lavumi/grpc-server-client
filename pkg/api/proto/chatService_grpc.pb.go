// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: api/proto/chatService.proto

package proto

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

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomServiceClient interface {
	GetRoomList(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomList, error)
	CreateRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error)
	JoinRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error)
	CreateStream(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (RoomService_CreateStreamClient, error)
	BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error)
	LeaveRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) GetRoomList(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomList, error) {
	out := new(RoomList)
	err := c.cc.Invoke(ctx, "/room.RoomService/GetRoomList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error) {
	out := new(RoomResponse)
	err := c.cc.Invoke(ctx, "/room.RoomService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) JoinRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error) {
	out := new(RoomResponse)
	err := c.cc.Invoke(ctx, "/room.RoomService/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) CreateStream(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (RoomService_CreateStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoomService_ServiceDesc.Streams[0], "/room.RoomService/CreateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &roomServiceCreateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoomService_CreateStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type roomServiceCreateStreamClient struct {
	grpc.ClientStream
}

func (x *roomServiceCreateStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *roomServiceClient) BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error) {
	out := new(Close)
	err := c.cc.Invoke(ctx, "/room.RoomService/BroadcastMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) LeaveRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*RoomResponse, error) {
	out := new(RoomResponse)
	err := c.cc.Invoke(ctx, "/room.RoomService/LeaveRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
// All implementations must embed UnimplementedRoomServiceServer
// for forward compatibility
type RoomServiceServer interface {
	GetRoomList(context.Context, *RoomRequest) (*RoomList, error)
	CreateRoom(context.Context, *RoomRequest) (*RoomResponse, error)
	JoinRoom(context.Context, *RoomRequest) (*RoomResponse, error)
	CreateStream(*RoomRequest, RoomService_CreateStreamServer) error
	BroadcastMessage(context.Context, *Message) (*Close, error)
	LeaveRoom(context.Context, *RoomRequest) (*RoomResponse, error)
	mustEmbedUnimplementedRoomServiceServer()
}

// UnimplementedRoomServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (UnimplementedRoomServiceServer) GetRoomList(context.Context, *RoomRequest) (*RoomList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}
func (UnimplementedRoomServiceServer) CreateRoom(context.Context, *RoomRequest) (*RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServiceServer) JoinRoom(context.Context, *RoomRequest) (*RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedRoomServiceServer) CreateStream(*RoomRequest, RoomService_CreateStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateStream not implemented")
}
func (UnimplementedRoomServiceServer) BroadcastMessage(context.Context, *Message) (*Close, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastMessage not implemented")
}
func (UnimplementedRoomServiceServer) LeaveRoom(context.Context, *RoomRequest) (*RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}
func (UnimplementedRoomServiceServer) mustEmbedUnimplementedRoomServiceServer() {}

// UnsafeRoomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServiceServer will
// result in compilation errors.
type UnsafeRoomServiceServer interface {
	mustEmbedUnimplementedRoomServiceServer()
}

func RegisterRoomServiceServer(s grpc.ServiceRegistrar, srv RoomServiceServer) {
	s.RegisterService(&RoomService_ServiceDesc, srv)
}

func _RoomService_GetRoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/GetRoomList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRoomList(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).JoinRoom(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_CreateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoomServiceServer).CreateStream(m, &roomServiceCreateStreamServer{stream})
}

type RoomService_CreateStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type roomServiceCreateStreamServer struct {
	grpc.ServerStream
}

func (x *roomServiceCreateStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _RoomService_BroadcastMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).BroadcastMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/BroadcastMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).BroadcastMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_LeaveRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).LeaveRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/LeaveRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).LeaveRoom(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomService_ServiceDesc is the grpc.ServiceDesc for RoomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "room.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoomList",
			Handler:    _RoomService_GetRoomList_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _RoomService_JoinRoom_Handler,
		},
		{
			MethodName: "BroadcastMessage",
			Handler:    _RoomService_BroadcastMessage_Handler,
		},
		{
			MethodName: "LeaveRoom",
			Handler:    _RoomService_LeaveRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateStream",
			Handler:       _RoomService_CreateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/proto/chatService.proto",
}