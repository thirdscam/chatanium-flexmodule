// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/core-v1/hooks.proto

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

const (
	Hooks_GetManifest_FullMethodName = "/core_v1.Hooks/GetManifest"
	Hooks_GetStatus_FullMethodName   = "/core_v1.Hooks/GetStatus"
	Hooks_OnStage_FullMethodName     = "/core_v1.Hooks/OnStage"
)

// HooksClient is the client API for Hooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HooksClient interface {
	GetManifest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetManifestResponse, error)
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStatusResponse, error)
	OnStage(ctx context.Context, in *OnStageRequest, opts ...grpc.CallOption) (*Empty, error)
}

type hooksClient struct {
	cc grpc.ClientConnInterface
}

func NewHooksClient(cc grpc.ClientConnInterface) HooksClient {
	return &hooksClient{cc}
}

func (c *hooksClient) GetManifest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetManifestResponse, error) {
	out := new(GetManifestResponse)
	err := c.cc.Invoke(ctx, Hooks_GetManifest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hooksClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, Hooks_GetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hooksClient) OnStage(ctx context.Context, in *OnStageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Hooks_OnStage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HooksServer is the server API for Hooks service.
// All implementations should embed UnimplementedHooksServer
// for forward compatibility
type HooksServer interface {
	GetManifest(context.Context, *Empty) (*GetManifestResponse, error)
	GetStatus(context.Context, *Empty) (*GetStatusResponse, error)
	OnStage(context.Context, *OnStageRequest) (*Empty, error)
}

// UnimplementedHooksServer should be embedded to have forward compatible implementations.
type UnimplementedHooksServer struct {
}

func (UnimplementedHooksServer) GetManifest(context.Context, *Empty) (*GetManifestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManifest not implemented")
}
func (UnimplementedHooksServer) GetStatus(context.Context, *Empty) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedHooksServer) OnStage(context.Context, *OnStageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnStage not implemented")
}

// UnsafeHooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HooksServer will
// result in compilation errors.
type UnsafeHooksServer interface {
	mustEmbedUnimplementedHooksServer()
}

func RegisterHooksServer(s grpc.ServiceRegistrar, srv HooksServer) {
	s.RegisterService(&Hooks_ServiceDesc, srv)
}

func _Hooks_GetManifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksServer).GetManifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hooks_GetManifest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksServer).GetManifest(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hooks_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hooks_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hooks_OnStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HooksServer).OnStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hooks_OnStage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HooksServer).OnStage(ctx, req.(*OnStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hooks_ServiceDesc is the grpc.ServiceDesc for Hooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core_v1.Hooks",
	HandlerType: (*HooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManifest",
			Handler:    _Hooks_GetManifest_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Hooks_GetStatus_Handler,
		},
		{
			MethodName: "OnStage",
			Handler:    _Hooks_OnStage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/core-v1/hooks.proto",
}
