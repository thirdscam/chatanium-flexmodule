// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package shared contains shared data between the host and plugins.
package shared

import (
	"context"

	"google.golang.org/grpc"

	"github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/sample/proto"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "FLEXMODULE_PLUGIN",
	MagicCookieValue: "CHATANIUM_FOREVER",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"core-v1": &CorePlugin{},
}

type Manifest struct {
	Name        string
	Version     string
	Author      string
	Repository  string
	Permissions Permissions
}

type Status struct {
	IsReady bool
}

type Permissions []string

// KV is the interface that we're exposing as a plugin.
type ICore interface {
	GetManifest() (Manifest, error)
	GetStatus() (Status, error)
	OnStage(stage string)
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type CorePlugin struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl ICore
}

func (p *CorePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterCoreServer(s, &GRPCServer{
		Impl:   p.Impl,
		broker: broker,
	})
	return nil
}

func (p *CorePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewCoreClient(c), broker: broker}, nil
}

var _ plugin.GRPCPlugin = &CorePlugin{}
