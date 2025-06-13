package runtime

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto_common "github.com/thirdscam/chatanium-flexmodule/proto"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/core-v1"
	shared "github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
)

// `runtime/client.go` implements the gRPC client for making calls to the module.
//
// This part works on the runtime-side and is the gRPC client implementation for the module.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.HookClient
}

func (m *GRPCClient) GetManifest() (shared.Manifest, error) {
	// RPC call to the gRPC server on the module-side
	resp, err := m.client.GetManifest(context.Background(), &proto_common.Empty{})
	if err != nil {
		return shared.Manifest{}, err
	}

	// Pass the results received from the module to the runtime
	return shared.Manifest{
		Name:        resp.Name,
		Version:     resp.Version,
		Author:      resp.Author,
		Repository:  resp.Repository,
		Permissions: shared.Permissions(resp.Permissions),
	}, nil
}

func (m *GRPCClient) GetStatus() (shared.Status, error) {
	// RPC call to the gRPC server on the module-side
	resp, err := m.client.GetStatus(context.Background(), &proto_common.Empty{})
	if err != nil {
		return shared.Status{}, err
	}

	// Pass the results received from the module to the runtime
	return shared.Status{
		IsReady: resp.IsReady,
	}, nil
}

func (m *GRPCClient) OnStage(stage string) {
	// RPC call to the gRPC server on the module-side
	m.client.OnStage(context.Background(), &proto.OnStageRequest{Stage: stage})

	// This function (hook) doesn't receive any results from the module, so the void function
}
