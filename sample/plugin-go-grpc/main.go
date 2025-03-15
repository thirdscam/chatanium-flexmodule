package main

import (
	"github.com/hashicorp/go-plugin"
	shared "github.com/thirdscam/chatanium-flexmodule/sample/shared"
	core "github.com/thirdscam/chatanium-flexmodule/sample/shared/core-v1"
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

func (m *Core) GetStatus() (core.Status, error) {
	return core.Status{
		IsReady: m.ready,
	}, nil
}

func (m *Core) OnStage(stage string) {
	if stage == "RUNTIME_STARTED" {
		m.ready = true
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
			// if want to implement more plugins, you can add it here!
			"core-v1": &core.Plugin{Impl: &Core{}},
			// "discord-v1": &shared.CorePlugin{Impl: &Discord{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
