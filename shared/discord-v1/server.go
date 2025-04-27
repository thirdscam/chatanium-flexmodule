package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
	plugin "github.com/hashicorp/go-plugin"
	proto_common "github.com/thirdscam/chatanium-flexmodule/proto"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Hook

	broker *plugin.GRPCBroker
}

func (m *GRPCServer) OnInit(ctx context.Context, req *proto_common.Empty) (*proto_common.Empty, error) {
	m.Impl.OnInit()

	return &proto_common.Empty{}, nil
}

func (m *GRPCServer) OnCreateChatMessage(ctx context.Context, req *proto.ChatMessage) (*proto_common.Empty, error) {
	m.Impl.OnCreateChatMessage(discordgo.Message{
		ID:        req.Id,
		ChannelID: req.ChannelId,
		GuildID:   req.GuildId,
		Content:   req.Content,
		Timestamp: req.Timestamp.AsTime(),
	})

	return &proto_common.Empty{}, nil
}

func (m *GRPCServer) OnCreateInteraction(ctx context.Context, req *proto.OnCreateInteractionRequest) (*proto_common.Empty, error) {
	m.Impl.OnCreateInteraction(discordgo.Interaction{
		ID:        req.Id,
		GuildID:   req.GuildId,
		ChannelID: req.ChannelId,
		Message: &discordgo.Message{}
	})

	return &proto_common.Empty{}, nil
}
