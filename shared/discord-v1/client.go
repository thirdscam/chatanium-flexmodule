package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
	plugin "github.com/hashicorp/go-plugin"
	proto_common "github.com/thirdscam/chatanium-flexmodule/proto"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/buf2struct"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/struct2buf"
)

// GRPCClient is an implementation of Hook that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.HookClient
}

// ================================================
// Runtime -> Module Server (Hook)
// ================================================

// OnInit calls the OnInit RPC method and returns the initialization response.
func (m *GRPCClient) OnInit() InitResponse {
	resp, err := m.client.OnInit(context.Background(), &proto_common.Empty{})
	if err != nil {
		return InitResponse{}
	}

	// Convert the interactions to discordgo.ApplicationCommand
	cmds := make([]*discordgo.ApplicationCommand, 0)
	for _, interaction := range resp.Interactions {
		// Convert the interaction to a discordgo.ApplicationCommand
		// and add it to the list of interactions.
		cmds = append(cmds, buf2struct.ApplicationCommand(interaction))
	}

	return InitResponse{
		Interactions: cmds,
	}
}

// OnCreateChatMessage sends a message to the plugin via RPC.
func (m *GRPCClient) OnCreateChatMessage(message *discordgo.Message) error {
	_, err := m.client.OnCreateMessage(
		context.Background(),
		struct2buf.Message(message),
	)
	return err
}

// OnCreateInteraction sends an interaction to the plugin via RPC.
func (m *GRPCClient) OnCreateInteraction(interaction *discordgo.Interaction) error {
	_, err := m.client.OnCreateInteraction(
		context.Background(),
		struct2buf.Interaction(interaction),
	)
	return err
}

// OnEvent sends an event to the plugin via RPC.
func (m *GRPCClient) OnEvent(event string) error {
	_, err := m.client.OnEvent(
		context.Background(),
		&proto.OnEventRequest{Event: event},
	)
	return err
}
