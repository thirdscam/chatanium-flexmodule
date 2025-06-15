package module

import (
	"context"

	"github.com/bwmarrin/discordgo"
	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/buf2struct"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/struct2buf"
)

// HelperClient implements the Helper interface for module-side operations.
// This client communicates with the runtime server to perform Discord operations.
type HelperClient struct {
	broker *plugin.GRPCBroker
	client proto.HelperClient
}

// ================================================
// Message operations
// ================================================

// ChannelMessageSend sends a simple text message to a channel.
func (h *HelperClient) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	resp, err := h.client.ChannelMessageSend(context.Background(), &proto.ChannelMessageSendRequest{
		ChannelId: channelID,
		Content:   content,
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}

// ChannelMessageSendComplex sends a complex message with attachments, embeds, etc.
func (h *HelperClient) ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (*discordgo.Message, error) {
	resp, err := h.client.ChannelMessageSendComplex(context.Background(), &proto.ChannelMessageSendComplexRequest{
		ChannelId: channelID,
		Data:      struct2buf.MessageSend(data),
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}

// ChannelMessageSendEmbed sends a message with a single embed.
func (h *HelperClient) ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	resp, err := h.client.ChannelMessageSendEmbed(context.Background(), &proto.ChannelMessageSendEmbedRequest{
		ChannelId: channelID,
		Embed:     struct2buf.MessageEmbed(embed),
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}

// ChannelMessageSendEmbeds sends a message with multiple embeds.
func (h *HelperClient) ChannelMessageSendEmbeds(channelID string, embeds []*discordgo.MessageEmbed) (*discordgo.Message, error) {
	protoEmbeds := make([]*proto.MessageEmbed, 0, len(embeds))
	for _, embed := range embeds {
		protoEmbeds = append(protoEmbeds, struct2buf.MessageEmbed(embed))
	}

	resp, err := h.client.ChannelMessageSendEmbeds(context.Background(), &proto.ChannelMessageSendEmbedsRequest{
		ChannelId: channelID,
		Embeds:    protoEmbeds,
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}

// ChannelMessageEdit edits a message with simple text content.
func (h *HelperClient) ChannelMessageEdit(channelID, messageID, content string) (*discordgo.Message, error) {
	resp, err := h.client.ChannelMessageEdit(context.Background(), &proto.ChannelMessageEditRequest{
		ChannelId: channelID,
		MessageId: messageID,
		Content:   content,
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}

// ChannelMessageEditComplex edits a message with complex data.
func (h *HelperClient) ChannelMessageEditComplex(m *discordgo.MessageEdit) (*discordgo.Message, error) {
	resp, err := h.client.ChannelMessageEditComplex(context.Background(), &proto.ChannelMessageEditComplexRequest{
		MessageEdit: struct2buf.MessageEdit(m),
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}

// ChannelMessageDelete deletes a message from a channel.
func (h *HelperClient) ChannelMessageDelete(channelID, messageID string) error {
	_, err := h.client.ChannelMessageDelete(context.Background(), &proto.ChannelMessageDeleteRequest{
		ChannelId: channelID,
		MessageId: messageID,
	})
	return err
}

// ChannelMessages retrieves multiple messages from a channel.
func (h *HelperClient) ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string) ([]*discordgo.Message, error) {
	resp, err := h.client.ChannelMessages(context.Background(), &proto.ChannelMessagesRequest{
		ChannelId: channelID,
		Limit:     int32(limit),
		BeforeId:  beforeID,
		AfterId:   afterID,
		AroundId:  aroundID,
	})
	if err != nil {
		return nil, err
	}

	messages := make([]*discordgo.Message, 0, len(resp.Messages))
	for _, msg := range resp.Messages {
		messages = append(messages, buf2struct.Message(msg))
	}
	return messages, nil
}

// ChannelMessage retrieves a single message from a channel.
func (h *HelperClient) ChannelMessage(channelID, messageID string) (*discordgo.Message, error) {
	resp, err := h.client.ChannelMessage(context.Background(), &proto.ChannelMessageRequest{
		ChannelId: channelID,
		MessageId: messageID,
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}
