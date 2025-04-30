package discord

import (
	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.HookClient
}
