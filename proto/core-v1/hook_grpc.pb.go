// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: core-v1/hook.proto

package core_v1

import (
	context "context"
	proto "github.com/thirdscam/chatanium-flexmodule/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Hook_GetManifest_FullMethodName = "/core_v1.Hook/GetManifest"
	Hook_GetStatus_FullMethodName   = "/core_v1.Hook/GetStatus"
	Hook_OnStage_FullMethodName     = "/core_v1.Hook/OnStage"
)

// HookClient is the client API for Hook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HookClient interface {
	GetManifest(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*GetManifestResponse, error)
	GetStatus(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*GetStatusResponse, error)
	OnStage(ctx context.Context, in *OnStageRequest, opts ...grpc.CallOption) (*proto.Empty, error)
}

type hookClient struct {
	cc grpc.ClientConnInterface
}

func NewHookClient(cc grpc.ClientConnInterface) HookClient {
	return &hookClient{cc}
}

func (c *hookClient) GetManifest(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*GetManifestResponse, error) {
	out := new(GetManifestResponse)
	err := c.cc.Invoke(ctx, Hook_GetManifest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookClient) GetStatus(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, Hook_GetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hookClient) OnStage(ctx context.Context, in *OnStageRequest, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, Hook_OnStage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HookServer is the server API for Hook service.
// All implementations should embed UnimplementedHookServer
// for forward compatibility
type HookServer interface {
	GetManifest(context.Context, *proto.Empty) (*GetManifestResponse, error)
	GetStatus(context.Context, *proto.Empty) (*GetStatusResponse, error)
	OnStage(context.Context, *OnStageRequest) (*proto.Empty, error)
}

// UnimplementedHookServer should be embedded to have forward compatible implementations.
type UnimplementedHookServer struct {
}

func (UnimplementedHookServer) GetManifest(context.Context, *proto.Empty) (*GetManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManifest not implemented")
}
func (UnimplementedHookServer) GetStatus(context.Context, *proto.Empty) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedHookServer) OnStage(context.Context, *OnStageRequest) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnStage not implemented")
}

// UnsafeHookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HookServer will
// result in compilation errors.
type UnsafeHookServer interface {
	mustEmbedUnimplementedHookServer()
}

func RegisterHookServer(s grpc.ServiceRegistrar, srv HookServer) {
	s.RegisterService(&Hook_ServiceDesc, srv)
}

func _Hook_GetManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServer).GetManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hook_GetManifest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServer).GetManifest(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hook_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hook_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServer).GetStatus(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hook_OnStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServer).OnStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hook_OnStage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServer).OnStage(ctx, req.(*OnStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hook_ServiceDesc is the grpc.ServiceDesc for Hook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core_v1.Hook",
	HandlerType: (*HookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManifest",
			Handler:    _Hook_GetManifest_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Hook_GetStatus_Handler,
		},
		{
			MethodName: "OnStage",
			Handler:    _Hook_OnStage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core-v1/hook.proto",
}
