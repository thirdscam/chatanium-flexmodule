package discord

import (
	"context"

	plugin "github.com/hashicorp/go-plugin"
	proto_common "github.com/thirdscam/chatanium-flexmodule/proto"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/buf2struct"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/struct2buf"
)

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Hook

	broker *plugin.GRPCBroker
}

// OnInit is called when the discord plugin is initialized.
func (m *GRPCServer) OnInit(ctx context.Context, req *proto_common.Empty) (*proto.InitResponse, error) {
	resp := m.Impl.OnInit()

	interactions := make([]*proto.ApplicationCommand, 0)
	for _, v := range resp.Interactions {
		interactions = append(interactions, struct2buf.ApplicationCommmand(v))
	}

	return &proto.InitResponse{
		Interactions: interactions,
	}, nil
}

// OnCreateChatMessage is called when a message is created from the runtime.
func (m *GRPCServer) OnCreateMessage(ctx context.Context, req *proto.Message) (*proto_common.Empty, error) {
	// Convert the protobuf message to a discordgo.Message struct
	m.Impl.OnCreateChatMessage(buf2struct.Message(req))

	// Hook function is not required to return anything to the client (runtime)
	return &proto_common.Empty{}, nil
}

// OnCreateInteraction is called when an interaction is created from the runtime.
func (m *GRPCServer) OnCreateInteraction(ctx context.Context, req *proto.Interaction) (*proto_common.Empty, error) {
	// Convert the protobuf message to a discordgo.Interaction struct
	m.Impl.OnCreateInteraction(buf2struct.Interaction(req))

	// Hook function is not required to return anything to the client (runtime)
	return &proto_common.Empty{}, nil
}

// OnEvent is called when an (discord) event is created from the runtime.
func (m *GRPCServer) OnEvent(ctx context.Context, req *proto.OnEventRequest) (*proto_common.Empty, error) {
	m.Impl.OnEvent(req.Event)

	// Hook function is not required to return anything to the client (runtime)
	return &proto_common.Empty{}, nil
}
