package discord

import plugin "github.com/hashicorp/go-plugin"

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Interface

	broker *plugin.GRPCBroker
}
