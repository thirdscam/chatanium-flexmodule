package runtime

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	shared "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
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

	Helper shared.Helper // The implementation of the Helper interface (runtime side)
	Hook   shared.Hook   // Hook client to call module's hook functions
}

func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	// Register Helper server (runtime provides helper services to modules)
	proto.RegisterHelperServer(s, &HelperServerImpl{
		Impl:   p.Helper,
		broker: broker,
	})
	return nil
}

func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	// Start a broker server for Helper service so module can call runtime's helpers
	var helperServerID uint32 = 1 // Use a fixed ID for the helper server
	go broker.AcceptAndServe(helperServerID, func(opts []grpc.ServerOption) *grpc.Server {
		s := grpc.NewServer(opts...)
		proto.RegisterHelperServer(s, &HelperServerImpl{
			Impl:   p.Helper,
			broker: broker,
		})
		return s
	})

	// Create Hook client to call module's hook functions
	hookClient := &HookClient{
		client:         proto.NewHookClient(c),
		broker:         broker,
		helperServerID: helperServerID, // Pass server ID to hook client
	}

	// Create Helper client (for completeness, though runtime usually provides helpers)
	helperClient := &HelperClientImpl{
		client: proto.NewHelperClient(c),
		broker: broker,
	}

	// Return both clients as a combined interface
	return &RuntimeClients{
		Hook:   hookClient,
		Helper: helperClient,
	}, nil
}

// RuntimeClients wraps both Hook and Helper clients for runtime
type RuntimeClients struct {
	Hook   shared.Hook
	Helper shared.Helper
}

// GetHook returns the Hook client
func (r *RuntimeClients) GetHook() shared.Hook {
	return r.Hook
}

// GetHelper returns the Helper client
func (r *RuntimeClients) GetHelper() shared.Helper {
	return r.Helper
}

var (
	_ plugin.GRPCPlugin     = &Plugin{}
	_ shared.RuntimeClients = &RuntimeClients{}
)
