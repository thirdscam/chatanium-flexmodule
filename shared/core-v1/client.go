package core

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/core-v1"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.HookClient
}

func (m *GRPCClient) GetManifest() (Manifest, error) {
	resp, err := m.client.GetManifest(context.Background(), &proto.Empty{})
	if err != nil {
		return Manifest{}, err
	}

	return Manifest{
		Name:        resp.Name,
		Version:     resp.Version,
		Author:      resp.Author,
		Repository:  resp.Repository,
		Permissions: Permissions(resp.Permissions),
	}, nil
}

func (m *GRPCClient) GetStatus() (Status, error) {
	resp, err := m.client.GetStatus(context.Background(), &proto.Empty{})
	if err != nil {
		return Status{}, err
	}

	return Status{
		IsReady: resp.IsReady,
	}, nil
}

func (m *GRPCClient) OnStage(stage string) {
	m.client.OnStage(context.Background(), &proto.OnStageRequest{Stage: stage})
}
