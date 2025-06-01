package shared

import (
	"github.com/hashicorp/go-plugin"
	core_module "github.com/thirdscam/chatanium-flexmodule/shared/core-v1/module"
	core_runtime "github.com/thirdscam/chatanium-flexmodule/shared/core-v1/runtime"

	discord "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "FLEXMODULE_PLUGIN",
	MagicCookieValue: "CHATANIUM_FOREVER",
}

// ModulePluginMap is the map of plugins for the runtime.
// This map is used at runtime, so it's not needed in the module implementation.
var RuntimePluginMap = map[string]plugin.Plugin{
	"core-v1":    &core_runtime.Plugin{},
	"discord-v1": &discord.Plugin{},
}

// ModulePluginMap is the map of plugins we can dispense.
var ModulePluginMap = map[string]plugin.Plugin{
	"core-v1":    &core_module.Plugin{},
	"discord-v1": &discord.Plugin{},
}

func ServeToRuntime(plugins map[string]plugin.Plugin) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins:         plugins,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
