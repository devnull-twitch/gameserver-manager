// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// GameserverManagerClient is the client API for GameserverManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameserverManagerClient interface {
	GetGameserver(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type gameserverManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewGameserverManagerClient(cc grpc.ClientConnInterface) GameserverManagerClient {
	return &gameserverManagerClient{cc}
}

func (c *gameserverManagerClient) GetGameserver(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/gameservermanager.GameserverManager/GetGameserver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameserverManagerServer is the server API for GameserverManager service.
// All implementations must embed UnimplementedGameserverManagerServer
// for forward compatibility
type GameserverManagerServer interface {
	GetGameserver(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedGameserverManagerServer()
}

// UnimplementedGameserverManagerServer must be embedded to have forward compatible implementations.
type UnimplementedGameserverManagerServer struct {
}

func (UnimplementedGameserverManagerServer) GetGameserver(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameserver not implemented")
}
func (UnimplementedGameserverManagerServer) mustEmbedUnimplementedGameserverManagerServer() {}

// UnsafeGameserverManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameserverManagerServer will
// result in compilation errors.
type UnsafeGameserverManagerServer interface {
	mustEmbedUnimplementedGameserverManagerServer()
}

func RegisterGameserverManagerServer(s grpc.ServiceRegistrar, srv GameserverManagerServer) {
	s.RegisterService(&GameserverManager_ServiceDesc, srv)
}

func _GameserverManager_GetGameserver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameserverManagerServer).GetGameserver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gameservermanager.GameserverManager/GetGameserver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameserverManagerServer).GetGameserver(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameserverManager_ServiceDesc is the grpc.ServiceDesc for GameserverManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameserverManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gameservermanager.GameserverManager",
	HandlerType: (*GameserverManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGameserver",
			Handler:    _GameserverManager_GetGameserver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gameserver.proto",
}