package module

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
// `module/server.go` implements the gRPC server for receiving from the runtime.
//
// `module/client.go` implements the gRPC client for making calls to the runtime.
type Plugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl   shared.Hook   // Hook implementation (module side)
	Helper shared.Helper // Helper service (provided by runtime, but can be nil)
}

func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	// Register Hook server (receives calls from runtime)
	proto.RegisterHookServer(s, &GRPCServer{
		Impl:   p.Impl,
		broker: broker,
		helper: p.Helper, // Pass helper service to hook server
	})

	return nil
}

func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	// Create Helper client to call runtime's helper service
	helperClient := &HelperClientImpl{
		client: proto.NewHelperClient(c),
		broker: broker,
	}

	// Create Hook client to call module's hook functions from runtime
	hookClient := &HookClient{
		client: proto.NewHookClient(c),
		broker: broker,
	}

	// Return both clients as a combined interface
	return &ModuleClients{
		Helper: helperClient,
		Hook:   hookClient,
	}, nil
}

// ModuleClients wraps both Helper and Hook clients
type ModuleClients struct {
	Helper shared.Helper
	Hook   shared.Hook
}

var _ plugin.GRPCPlugin = &Plugin{}
