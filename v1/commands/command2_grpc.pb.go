// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: v1/commands/command2.proto

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

// GetCommandClient is the client API for GetCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetCommandClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type getCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewGetCommandClient(cc grpc.ClientConnInterface) GetCommandClient {
	return &getCommandClient{cc}
}

func (c *getCommandClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/command.GetCommand/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetCommandServer is the server API for GetCommand service.
// All implementations must embed UnimplementedGetCommandServer
// for forward compatibility
type GetCommandServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedGetCommandServer()
}

// UnimplementedGetCommandServer must be embedded to have forward compatible implementations.
type UnimplementedGetCommandServer struct {
}

func (UnimplementedGetCommandServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGetCommandServer) mustEmbedUnimplementedGetCommandServer() {}

// UnsafeGetCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetCommandServer will
// result in compilation errors.
type UnsafeGetCommandServer interface {
	mustEmbedUnimplementedGetCommandServer()
}

func RegisterGetCommandServer(s grpc.ServiceRegistrar, srv GetCommandServer) {
	s.RegisterService(&GetCommand_ServiceDesc, srv)
}

func _GetCommand_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetCommandServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/command.GetCommand/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetCommandServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetCommand_ServiceDesc is the grpc.ServiceDesc for GetCommand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetCommand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "command.GetCommand",
	HandlerType: (*GetCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GetCommand_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/commands/command2.proto",
}
