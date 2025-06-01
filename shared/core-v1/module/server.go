package module

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto_common "github.com/thirdscam/chatanium-flexmodule/proto"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/core-v1"
	shared "github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
)

// This part works on the module-side and is the gRPC server implementation for the runtime.
type GRPCServer struct {
	Impl   shared.Hook // Hook functions implemented by module developers
	broker *plugin.GRPCBroker
}

func (m *GRPCServer) GetManifest(ctx context.Context, req *proto_common.Empty) (*proto.GetManifestResponse, error) {
	// Receive results from modules
	manifest, err := m.Impl.GetManifest()
	if err != nil {
		return nil, err
	}

	// Serve manifests to the runtime
	return &proto.GetManifestResponse{
		Name:        manifest.Name,
		Version:     manifest.Version,
		Author:      manifest.Author,
		Repository:  manifest.Repository,
		Permissions: []string(manifest.Permissions),
	}, nil
}

func (m *GRPCServer) GetStatus(ctx context.Context, req *proto_common.Empty) (*proto.GetStatusResponse, error) {
	// Receive results from modules
	status, err := m.Impl.GetStatus()
	if err != nil {
		return nil, err
	}

	// Serve status to the runtime
	return &proto.GetStatusResponse{
		IsReady: status.IsReady,
	}, nil
}

func (m *GRPCServer) OnStage(ctx context.Context, req *proto.OnStageRequest) (*proto_common.Empty, error) {
	m.Impl.OnStage(req.Stage)

	return &proto_common.Empty{}, nil
}
