// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: dataway.proto

package pb

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

// DatawayClient is the client API for Dataway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatawayClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RegisterNewTom(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UUID, error)
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error)
	DeleteSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type datawayClient struct {
	cc grpc.ClientConnInterface
}

func NewDatawayClient(cc grpc.ClientConnInterface) DatawayClient {
	return &datawayClient{cc}
}

func (c *datawayClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Dataway/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datawayClient) RegisterNewTom(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/proto.Dataway/RegisterNewTom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datawayClient) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Subscription, error) {
	out := new(Subscription)
	err := c.cc.Invoke(ctx, "/proto.Dataway/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datawayClient) DeleteSubscription(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Dataway/DeleteSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatawayServer is the server API for Dataway service.
// All implementations must embed UnimplementedDatawayServer
// for forward compatibility
type DatawayServer interface {
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	RegisterNewTom(context.Context, *emptypb.Empty) (*UUID, error)
	Subscribe(context.Context, *Subscription) (*Subscription, error)
	DeleteSubscription(context.Context, *Subscription) (*emptypb.Empty, error)
	mustEmbedUnimplementedDatawayServer()
}

// UnimplementedDatawayServer must be embedded to have forward compatible implementations.
type UnimplementedDatawayServer struct {
}

func (UnimplementedDatawayServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDatawayServer) RegisterNewTom(context.Context, *emptypb.Empty) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNewTom not implemented")
}
func (UnimplementedDatawayServer) Subscribe(context.Context, *Subscription) (*Subscription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedDatawayServer) DeleteSubscription(context.Context, *Subscription) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (UnimplementedDatawayServer) mustEmbedUnimplementedDatawayServer() {}

// UnsafeDatawayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatawayServer will
// result in compilation errors.
type UnsafeDatawayServer interface {
	mustEmbedUnimplementedDatawayServer()
}

func RegisterDatawayServer(s grpc.ServiceRegistrar, srv DatawayServer) {
	s.RegisterService(&Dataway_ServiceDesc, srv)
}

func _Dataway_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatawayServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dataway/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatawayServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataway_RegisterNewTom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatawayServer).RegisterNewTom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dataway/RegisterNewTom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatawayServer).RegisterNewTom(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataway_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatawayServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dataway/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatawayServer).Subscribe(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dataway_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatawayServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dataway/DeleteSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatawayServer).DeleteSubscription(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

// Dataway_ServiceDesc is the grpc.ServiceDesc for Dataway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dataway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Dataway",
	HandlerType: (*DatawayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Dataway_Ping_Handler,
		},
		{
			MethodName: "RegisterNewTom",
			Handler:    _Dataway_RegisterNewTom_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Dataway_Subscribe_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _Dataway_DeleteSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataway.proto",
}
