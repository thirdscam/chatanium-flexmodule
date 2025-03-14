package shared

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/thirdscam/chatanium-flexmodule/sample/proto"
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
		Name:       resp.Name,
		Backend:    resp.Backend,
		Version:    resp.Version,
		Author:     resp.Author,
		Repository: resp.Repository,
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
		Name:       manifest.Name,
		Backend:    manifest.Backend,
		Version:    manifest.Version,
		Author:     manifest.Author,
		Repository: manifest.Repository,
	}, nil
}

func (m *GRPCServer) OnStage(ctx context.Context, req *proto.OnStageRequest) (*proto.Empty, error) {
	m.Impl.OnStage(req.Stage)

	return &proto.Empty{}, nil
}
