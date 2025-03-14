// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/thirdscam/chatanium-flexmodule/sample/shared"
)

type Core struct{}

func (m *Core) GetManifest() (shared.Manifest, error) {
	return shared.Manifest{
		Name:       "TestModule",
		Backend:    "discord",
		Version:    "0.0.1",
		Author:     "thirdscam",
		Repository: "github:thirdscam/chatanium-flexmodule",
	}, nil
}

func (m *Core) OnStage(stage string) {
	// Do nothing
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"Core": &shared.CorePlugin{Impl: &Core{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
