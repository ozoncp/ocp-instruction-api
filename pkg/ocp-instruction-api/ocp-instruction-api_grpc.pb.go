// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_instruction_api

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

// OcpInstructionClient is the client API for OcpInstruction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpInstructionClient interface {
	CreateV1(ctx context.Context, in *CreateV1Request, opts ...grpc.CallOption) (*CreateV1Response, error)
	CreateMultiV1(ctx context.Context, in *CreateMultiV1Request, opts ...grpc.CallOption) (*CreateMultiV1Response, error)
	DescribeV1(ctx context.Context, in *DescribeV1Request, opts ...grpc.CallOption) (*DescribeV1Response, error)
	ListV1(ctx context.Context, in *ListV1Request, opts ...grpc.CallOption) (*ListV1Response, error)
	RemoveV1(ctx context.Context, in *RemoveV1Request, opts ...grpc.CallOption) (*RemoveV1Response, error)
	UpdateV1(ctx context.Context, in *UpdateV1Request, opts ...grpc.CallOption) (*UpdateV1Response, error)
}

type ocpInstructionClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpInstructionClient(cc grpc.ClientConnInterface) OcpInstructionClient {
	return &ocpInstructionClient{cc}
}

func (c *ocpInstructionClient) CreateV1(ctx context.Context, in *CreateV1Request, opts ...grpc.CallOption) (*CreateV1Response, error) {
	out := new(CreateV1Response)
	err := c.cc.Invoke(ctx, "/ozoncp.ocp_instruction_api.OcpInstruction/CreateV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpInstructionClient) CreateMultiV1(ctx context.Context, in *CreateMultiV1Request, opts ...grpc.CallOption) (*CreateMultiV1Response, error) {
	out := new(CreateMultiV1Response)
	err := c.cc.Invoke(ctx, "/ozoncp.ocp_instruction_api.OcpInstruction/CreateMultiV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpInstructionClient) DescribeV1(ctx context.Context, in *DescribeV1Request, opts ...grpc.CallOption) (*DescribeV1Response, error) {
	out := new(DescribeV1Response)
	err := c.cc.Invoke(ctx, "/ozoncp.ocp_instruction_api.OcpInstruction/DescribeV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpInstructionClient) ListV1(ctx context.Context, in *ListV1Request, opts ...grpc.CallOption) (*ListV1Response, error) {
	out := new(ListV1Response)
	err := c.cc.Invoke(ctx, "/ozoncp.ocp_instruction_api.OcpInstruction/ListV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpInstructionClient) RemoveV1(ctx context.Context, in *RemoveV1Request, opts ...grpc.CallOption) (*RemoveV1Response, error) {
	out := new(RemoveV1Response)
	err := c.cc.Invoke(ctx, "/ozoncp.ocp_instruction_api.OcpInstruction/RemoveV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpInstructionClient) UpdateV1(ctx context.Context, in *UpdateV1Request, opts ...grpc.CallOption) (*UpdateV1Response, error) {
	out := new(UpdateV1Response)
	err := c.cc.Invoke(ctx, "/ozoncp.ocp_instruction_api.OcpInstruction/UpdateV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpInstructionServer is the server API for OcpInstruction service.
// All implementations must embed UnimplementedOcpInstructionServer
// for forward compatibility
type OcpInstructionServer interface {
	CreateV1(context.Context, *CreateV1Request) (*CreateV1Response, error)
	CreateMultiV1(context.Context, *CreateMultiV1Request) (*CreateMultiV1Response, error)
	DescribeV1(context.Context, *DescribeV1Request) (*DescribeV1Response, error)
	ListV1(context.Context, *ListV1Request) (*ListV1Response, error)
	RemoveV1(context.Context, *RemoveV1Request) (*RemoveV1Response, error)
	UpdateV1(context.Context, *UpdateV1Request) (*UpdateV1Response, error)
	mustEmbedUnimplementedOcpInstructionServer()
}

// UnimplementedOcpInstructionServer must be embedded to have forward compatible implementations.
type UnimplementedOcpInstructionServer struct {
}

func (UnimplementedOcpInstructionServer) CreateV1(context.Context, *CreateV1Request) (*CreateV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateV1 not implemented")
}
func (UnimplementedOcpInstructionServer) CreateMultiV1(context.Context, *CreateMultiV1Request) (*CreateMultiV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMultiV1 not implemented")
}
func (UnimplementedOcpInstructionServer) DescribeV1(context.Context, *DescribeV1Request) (*DescribeV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeV1 not implemented")
}
func (UnimplementedOcpInstructionServer) ListV1(context.Context, *ListV1Request) (*ListV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListV1 not implemented")
}
func (UnimplementedOcpInstructionServer) RemoveV1(context.Context, *RemoveV1Request) (*RemoveV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveV1 not implemented")
}
func (UnimplementedOcpInstructionServer) UpdateV1(context.Context, *UpdateV1Request) (*UpdateV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateV1 not implemented")
}
func (UnimplementedOcpInstructionServer) mustEmbedUnimplementedOcpInstructionServer() {}

// UnsafeOcpInstructionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpInstructionServer will
// result in compilation errors.
type UnsafeOcpInstructionServer interface {
	mustEmbedUnimplementedOcpInstructionServer()
}

func RegisterOcpInstructionServer(s grpc.ServiceRegistrar, srv OcpInstructionServer) {
	s.RegisterService(&OcpInstruction_ServiceDesc, srv)
}

func _OcpInstruction_CreateV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpInstructionServer).CreateV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozoncp.ocp_instruction_api.OcpInstruction/CreateV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpInstructionServer).CreateV1(ctx, req.(*CreateV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpInstruction_CreateMultiV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMultiV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpInstructionServer).CreateMultiV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozoncp.ocp_instruction_api.OcpInstruction/CreateMultiV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpInstructionServer).CreateMultiV1(ctx, req.(*CreateMultiV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpInstruction_DescribeV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpInstructionServer).DescribeV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozoncp.ocp_instruction_api.OcpInstruction/DescribeV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpInstructionServer).DescribeV1(ctx, req.(*DescribeV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpInstruction_ListV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpInstructionServer).ListV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozoncp.ocp_instruction_api.OcpInstruction/ListV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpInstructionServer).ListV1(ctx, req.(*ListV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpInstruction_RemoveV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpInstructionServer).RemoveV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozoncp.ocp_instruction_api.OcpInstruction/RemoveV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpInstructionServer).RemoveV1(ctx, req.(*RemoveV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpInstruction_UpdateV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpInstructionServer).UpdateV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozoncp.ocp_instruction_api.OcpInstruction/UpdateV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpInstructionServer).UpdateV1(ctx, req.(*UpdateV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpInstruction_ServiceDesc is the grpc.ServiceDesc for OcpInstruction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpInstruction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozoncp.ocp_instruction_api.OcpInstruction",
	HandlerType: (*OcpInstructionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateV1",
			Handler:    _OcpInstruction_CreateV1_Handler,
		},
		{
			MethodName: "CreateMultiV1",
			Handler:    _OcpInstruction_CreateMultiV1_Handler,
		},
		{
			MethodName: "DescribeV1",
			Handler:    _OcpInstruction_DescribeV1_Handler,
		},
		{
			MethodName: "ListV1",
			Handler:    _OcpInstruction_ListV1_Handler,
		},
		{
			MethodName: "RemoveV1",
			Handler:    _OcpInstruction_RemoveV1_Handler,
		},
		{
			MethodName: "UpdateV1",
			Handler:    _OcpInstruction_UpdateV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocp-instruction-api.proto",
}
