package module

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/core-v1"
	shared "github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
	"google.golang.org/grpc"
)

// FlexModule uses hashicorp/go-plugin:
// So we need to declare a separate Plugin for the runtime and module.
//
// `module/server.go` implements the gRPC server for receiving from the runtime.
//
// `module/client.go` implements the gRPC client for making calls to the runtime.
type Plugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl shared.Hook
}

func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterHookServer(s, &GRPCServer{
		Impl:   p.Impl,
		broker: broker,
	})
	return nil
}

func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, nil
}

var _ plugin.GRPCPlugin = &Plugin{}
