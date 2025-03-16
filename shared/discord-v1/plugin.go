package discord

import (
	plugin "github.com/hashicorp/go-plugin"
)

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type Plugin struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Interface
}

// func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
// 	proto.RegisterHookServer(s, &GRPCServer{
// 		Impl:   p.Impl,
// 		broker: broker,
// 	})
// 	return nil
// }

// func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
// 	return &GRPCClient{client: proto.NewHookClient(c), broker: broker}, nil
// }

// var _ plugin.GRPCPlugin = &Plugin{}
