package runtime

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"google.golang.org/grpc"
)

// FlexModule uses hashicorp/go-plugin:
// So we need to declare a separate Plugin for the runtime and module.
//
// `runtime/server.go` implements the gRPC server for receiving from the module.
//
// `runtime/client.go` implements the gRPC client for making calls to the module.
type Plugin struct {
	plugin.NetRPCUnsupportedPlugin
}

func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return nil
}

func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewHookClient(c), broker: broker}, nil
}

var _ plugin.GRPCPlugin = &Plugin{}
