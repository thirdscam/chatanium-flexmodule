package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/thirdscam/chatanium-flexmodule/sample/shared"
)

var PERMISSIONS = shared.Permissions{
	"DISCORD_V1_ON_CREATE_MESSAGE",
	"DISCORD_V1_ON_CREATE_INTERACTION",
	"DISCORD_V1_REQ_VOICE_STATE",
	"DISCORD_V1_CREATE_VOICE_STREAM",
}

var MANIFEST = shared.Manifest{
	Name:        "TestModule",
	Version:     "0.0.1",
	Author:      "thirdscam",
	Repository:  "github:thirdscam/chatanium-flexmodule",
	Permissions: PERMISSIONS,
}

var Ready bool

type Core struct{}

func (m *Core) GetManifest() (shared.Manifest, error) {
	return MANIFEST, nil
}

func (m *Core) GetStatus() (shared.Status, error) {
	return shared.Status{
		IsReady: Ready,
	}, nil
}

func (m *Core) OnStage(stage string) {
	if stage == "RUNTIME_STARTED" {
		Ready = true
	}
}

type Discord struct{}

func (d *Discord) OnCreateMessage(message string) {
	// Do nothing
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"core-v1": &shared.CorePlugin{Impl: &Core{}},
			// "discord-v1": &shared.CorePlugin{Impl: &Discord{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
