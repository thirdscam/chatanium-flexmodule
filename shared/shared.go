package shared

import (
	"github.com/hashicorp/go-plugin"
	core "github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "FLEXMODULE_PLUGIN",
	MagicCookieValue: "CHATANIUM_FOREVER",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"core-v1": &core.Plugin{},
}
