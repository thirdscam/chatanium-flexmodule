// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/discord-v1/helpers.proto

package proto

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

const (
	Helpers_ResponseMessage_FullMethodName     = "/discord_v1.Helpers/ResponseMessage"
	Helpers_ResponseInteraction_FullMethodName = "/discord_v1.Helpers/ResponseInteraction"
	Helpers_EditMessage_FullMethodName         = "/discord_v1.Helpers/EditMessage"
	Helpers_EditInteraction_FullMethodName     = "/discord_v1.Helpers/EditInteraction"
)

// HelpersClient is the client API for Helpers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelpersClient interface {
	ResponseMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ChatIdResponse, error)
	ResponseInteraction(ctx context.Context, in *ResponseInteractionRequest, opts ...grpc.CallOption) (*ChatIdResponse, error)
	EditMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EditInteraction(ctx context.Context, in *EditInteractionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type helpersClient struct {
	cc grpc.ClientConnInterface
}

func NewHelpersClient(cc grpc.ClientConnInterface) HelpersClient {
	return &helpersClient{cc}
}

func (c *helpersClient) ResponseMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ChatIdResponse, error) {
	out := new(ChatIdResponse)
	err := c.cc.Invoke(ctx, Helpers_ResponseMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpersClient) ResponseInteraction(ctx context.Context, in *ResponseInteractionRequest, opts ...grpc.CallOption) (*ChatIdResponse, error) {
	out := new(ChatIdResponse)
	err := c.cc.Invoke(ctx, Helpers_ResponseInteraction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpersClient) EditMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Helpers_EditMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpersClient) EditInteraction(ctx context.Context, in *EditInteractionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Helpers_EditInteraction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelpersServer is the server API for Helpers service.
// All implementations should embed UnimplementedHelpersServer
// for forward compatibility
type HelpersServer interface {
	ResponseMessage(context.Context, *ChatMessage) (*ChatIdResponse, error)
	ResponseInteraction(context.Context, *ResponseInteractionRequest) (*ChatIdResponse, error)
	EditMessage(context.Context, *ChatMessage) (*emptypb.Empty, error)
	EditInteraction(context.Context, *EditInteractionRequest) (*emptypb.Empty, error)
}

// UnimplementedHelpersServer should be embedded to have forward compatible implementations.
type UnimplementedHelpersServer struct {
}

func (UnimplementedHelpersServer) ResponseMessage(context.Context, *ChatMessage) (*ChatIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResponseMessage not implemented")
}
func (UnimplementedHelpersServer) ResponseInteraction(context.Context, *ResponseInteractionRequest) (*ChatIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResponseInteraction not implemented")
}
func (UnimplementedHelpersServer) EditMessage(context.Context, *ChatMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMessage not implemented")
}
func (UnimplementedHelpersServer) EditInteraction(context.Context, *EditInteractionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditInteraction not implemented")
}

// UnsafeHelpersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelpersServer will
// result in compilation errors.
type UnsafeHelpersServer interface {
	mustEmbedUnimplementedHelpersServer()
}

func RegisterHelpersServer(s grpc.ServiceRegistrar, srv HelpersServer) {
	s.RegisterService(&Helpers_ServiceDesc, srv)
}

func _Helpers_ResponseMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpersServer).ResponseMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Helpers_ResponseMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpersServer).ResponseMessage(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpers_ResponseInteraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResponseInteractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpersServer).ResponseInteraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Helpers_ResponseInteraction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpersServer).ResponseInteraction(ctx, req.(*ResponseInteractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpers_EditMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpersServer).EditMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Helpers_EditMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpersServer).EditMessage(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helpers_EditInteraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditInteractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpersServer).EditInteraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Helpers_EditInteraction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpersServer).EditInteraction(ctx, req.(*EditInteractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Helpers_ServiceDesc is the grpc.ServiceDesc for Helpers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Helpers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discord_v1.Helpers",
	HandlerType: (*HelpersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResponseMessage",
			Handler:    _Helpers_ResponseMessage_Handler,
		},
		{
			MethodName: "ResponseInteraction",
			Handler:    _Helpers_ResponseInteraction_Handler,
		},
		{
			MethodName: "EditMessage",
			Handler:    _Helpers_EditMessage_Handler,
		},
		{
			MethodName: "EditInteraction",
			Handler:    _Helpers_EditInteraction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/discord-v1/helpers.proto",
}
