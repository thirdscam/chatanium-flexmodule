package core

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/core-v1"
)

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl   Interface
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
