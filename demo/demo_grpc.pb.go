// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: demo.proto

package demo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Demo_Create_FullMethodName = "/Demo/Create"
	Demo_Delulu_FullMethodName = "/Demo/Delulu"
)

// DemoClient is the client API for Demo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoClient interface {
	Create(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error)
	Delulu(ctx context.Context, in *DeluluRequest, opts ...grpc.CallOption) (*DeluluRequest, error)
}

type demoClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoClient(cc grpc.ClientConnInterface) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) Create(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DemoResponse)
	err := c.cc.Invoke(ctx, Demo_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) Delulu(ctx context.Context, in *DeluluRequest, opts ...grpc.CallOption) (*DeluluRequest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeluluRequest)
	err := c.cc.Invoke(ctx, Demo_Delulu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServer is the server API for Demo service.
// All implementations must embed UnimplementedDemoServer
// for forward compatibility
type DemoServer interface {
	Create(context.Context, *DemoRequest) (*DemoResponse, error)
	Delulu(context.Context, *DeluluRequest) (*DeluluRequest, error)
	mustEmbedUnimplementedDemoServer()
}

// UnimplementedDemoServer must be embedded to have forward compatible implementations.
type UnimplementedDemoServer struct {
}

func (UnimplementedDemoServer) Create(context.Context, *DemoRequest) (*DemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDemoServer) Delulu(context.Context, *DeluluRequest) (*DeluluRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delulu not implemented")
}
func (UnimplementedDemoServer) mustEmbedUnimplementedDemoServer() {}

// UnsafeDemoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServer will
// result in compilation errors.
type UnsafeDemoServer interface {
	mustEmbedUnimplementedDemoServer()
}

func RegisterDemoServer(s grpc.ServiceRegistrar, srv DemoServer) {
	s.RegisterService(&Demo_ServiceDesc, srv)
}

func _Demo_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Demo_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).Create(ctx, req.(*DemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_Delulu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeluluRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).Delulu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Demo_Delulu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).Delulu(ctx, req.(*DeluluRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Demo_ServiceDesc is the grpc.ServiceDesc for Demo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Demo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Demo_Create_Handler,
		},
		{
			MethodName: "Delulu",
			Handler:    _Demo_Delulu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
