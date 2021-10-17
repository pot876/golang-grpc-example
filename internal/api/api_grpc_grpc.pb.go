// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// FiboClient is the client API for Fibo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FiboClient interface {
	GetFiboNumbers(ctx context.Context, in *FiboRequest, opts ...grpc.CallOption) (*FiboReply, error)
}

type fiboClient struct {
	cc grpc.ClientConnInterface
}

func NewFiboClient(cc grpc.ClientConnInterface) FiboClient {
	return &fiboClient{cc}
}

func (c *fiboClient) GetFiboNumbers(ctx context.Context, in *FiboRequest, opts ...grpc.CallOption) (*FiboReply, error) {
	out := new(FiboReply)
	err := c.cc.Invoke(ctx, "/Fibo/GetFiboNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FiboServer is the server API for Fibo service.
// All implementations must embed UnimplementedFiboServer
// for forward compatibility
type FiboServer interface {
	GetFiboNumbers(context.Context, *FiboRequest) (*FiboReply, error)
	mustEmbedUnimplementedFiboServer()
}

// UnimplementedFiboServer must be embedded to have forward compatible implementations.
type UnimplementedFiboServer struct {
}

func (UnimplementedFiboServer) GetFiboNumbers(context.Context, *FiboRequest) (*FiboReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiboNumbers not implemented")
}
func (UnimplementedFiboServer) mustEmbedUnimplementedFiboServer() {}

// UnsafeFiboServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FiboServer will
// result in compilation errors.
type UnsafeFiboServer interface {
	mustEmbedUnimplementedFiboServer()
}

func RegisterFiboServer(s grpc.ServiceRegistrar, srv FiboServer) {
	s.RegisterService(&Fibo_ServiceDesc, srv)
}

func _Fibo_GetFiboNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FiboRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiboServer).GetFiboNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Fibo/GetFiboNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiboServer).GetFiboNumbers(ctx, req.(*FiboRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fibo_ServiceDesc is the grpc.ServiceDesc for Fibo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fibo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Fibo",
	HandlerType: (*FiboServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFiboNumbers",
			Handler:    _Fibo_GetFiboNumbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_grpc.proto",
}
