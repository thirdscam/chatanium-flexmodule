package core

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/sample/proto/core-v1"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.CoreClient
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

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl ICore

	broker *plugin.GRPCBroker
}

func (m *GRPCServer) GetManifest(ctx context.Context, req *proto.Empty) (*proto.GetManifestResponse, error) {
	manifest, err := m.Impl.GetManifest()
	if err != nil {
		return nil, err
	}

	return &proto.GetManifestResponse{
		Name:        manifest.Name,
		Version:     manifest.Version,
		Author:      manifest.Author,
		Repository:  manifest.Repository,
		Permissions: []string(manifest.Permissions),
	}, nil
}

func (m *GRPCServer) GetStatus(ctx context.Context, req *proto.Empty) (*proto.GetStatusResponse, error) {
	status, err := m.Impl.GetStatus()
	if err != nil {
		return nil, err
	}

	return &proto.GetStatusResponse{
		IsReady: status.IsReady,
	}, nil
}

func (m *GRPCServer) OnStage(ctx context.Context, req *proto.OnStageRequest) (*proto.Empty, error) {
	m.Impl.OnStage(req.Stage)

	return &proto.Empty{}, nil
}
