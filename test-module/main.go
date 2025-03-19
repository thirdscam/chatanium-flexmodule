package main

import (
	"github.com/hashicorp/go-plugin"
	shared "github.com/thirdscam/chatanium-flexmodule/shared"
	core "github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
	discord "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
)

var PERMISSIONS = core.Permissions{
	"DISCORD_V1_ON_CREATE_MESSAGE",
	"DISCORD_V1_ON_CREATE_INTERACTION",
	"DISCORD_V1_REQ_VOICE_STATE",
	"DISCORD_V1_CREATE_VOICE_STREAM",
}

var MANIFEST = core.Manifest{
	Name:        "TestModule",
	Version:     "0.0.1",
	Author:      "thirdscam",
	Repository:  "github:thirdscam/chatanium-flexmodule",
	Permissions: PERMISSIONS,
}

type Core struct {
	ready bool
}

func (m *Core) GetManifest() (core.Manifest, error) {
	return MANIFEST, nil
}

// GetStatus returns the status of the plugin.
func (m *Core) GetStatus() (core.Status, error) {
	return core.Status{
		IsReady: m.ready,
	}, nil
}

// OnStage is called when the plugin is in a certain stage.
func (m *Core) OnStage(stage string) {
	if stage == "RUNTIME_STARTED" {
		m.ready = true
	}
}

type Discord struct {
	discord.UsePartialHooks
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			// if want to implement more plugins, you can add it here!
			"core-v1":    &core.Plugin{Impl: &Core{}},
			"discord-v1": &discord.Plugin{Impl: &Discord{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
