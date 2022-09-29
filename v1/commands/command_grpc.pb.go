// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: v1/commands/command.proto

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

// CommandClient is the client API for Command service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandClient interface {
	//SNMP Get given the list of OIDs
	Get(ctx context.Context, in *OidList, opts ...grpc.CallOption) (*SnmpPacket, error)
}

type commandClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandClient(cc grpc.ClientConnInterface) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) Get(ctx context.Context, in *OidList, opts ...grpc.CallOption) (*SnmpPacket, error) {
	out := new(SnmpPacket)
	err := c.cc.Invoke(ctx, "/thebinary.snmp.Command/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServer is the server API for Command service.
// All implementations must embed UnimplementedCommandServer
// for forward compatibility
type CommandServer interface {
	//SNMP Get given the list of OIDs
	Get(context.Context, *OidList) (*SnmpPacket, error)
	mustEmbedUnimplementedCommandServer()
}

// UnimplementedCommandServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServer struct {
}

func (UnimplementedCommandServer) Get(context.Context, *OidList) (*SnmpPacket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCommandServer) mustEmbedUnimplementedCommandServer() {}

// UnsafeCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServer will
// result in compilation errors.
type UnsafeCommandServer interface {
	mustEmbedUnimplementedCommandServer()
}

func RegisterCommandServer(s grpc.ServiceRegistrar, srv CommandServer) {
	s.RegisterService(&Command_ServiceDesc, srv)
}

func _Command_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OidList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thebinary.snmp.Command/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Get(ctx, req.(*OidList))
	}
	return interceptor(ctx, in, info, handler)
}

// Command_ServiceDesc is the grpc.ServiceDesc for Command service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Command_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thebinary.snmp.Command",
	HandlerType: (*CommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Command_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/commands/command.proto",
}