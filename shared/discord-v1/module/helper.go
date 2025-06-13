package module

import (
	"context"

	"github.com/bwmarrin/discordgo"
	plugin "github.com/hashicorp/go-plugin"
	proto_common "github.com/thirdscam/chatanium-flexmodule/proto"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	shared "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/buf2struct"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/struct2buf"
)

// HelperServer implements the Helper gRPC server for the module-side.
// This server receives calls from the runtime and delegates them to the actual Helper implementation.
type HelperServer struct {
	Helper shared.Helper // Helper implementation that provides Discord operations
	broker *plugin.GRPCBroker
}

// NewHelperServer creates a new HelperServer instance
func NewHelperServer(helper shared.Helper, broker *plugin.GRPCBroker) *HelperServer {
	return &HelperServer{
		Helper: helper,
		broker: broker,
	}
}

// ================================================
// Message operations
// ================================================

// ChannelMessageSend handles sending a simple text message.
func (h *HelperServer) ChannelMessageSend(ctx context.Context, req *proto.ChannelMessageSendRequest) (*proto.ChannelMessageSendResponse, error) {
	message, err := h.Helper.ChannelMessageSend(req.ChannelId, req.Content)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageSendComplex handles sending a complex message with attachments, embeds, etc.
func (h *HelperServer) ChannelMessageSendComplex(ctx context.Context, req *proto.ChannelMessageSendComplexRequest) (*proto.ChannelMessageSendComplexResponse, error) {
	data := buf2struct.MessageSend(req.Data)
	message, err := h.Helper.ChannelMessageSendComplex(req.ChannelId, data)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendComplexResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageSendEmbed handles sending a message with a single embed.
func (h *HelperServer) ChannelMessageSendEmbed(ctx context.Context, req *proto.ChannelMessageSendEmbedRequest) (*proto.ChannelMessageSendEmbedResponse, error) {
	embed := buf2struct.MessageEmbed(req.Embed)
	message, err := h.Helper.ChannelMessageSendEmbed(req.ChannelId, embed)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendEmbedResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageSendEmbeds handles sending a message with multiple embeds.
func (h *HelperServer) ChannelMessageSendEmbeds(ctx context.Context, req *proto.ChannelMessageSendEmbedsRequest) (*proto.ChannelMessageSendEmbedsResponse, error) {
	embeds := make([]*discordgo.MessageEmbed, 0, len(req.Embeds))
	for _, embed := range req.Embeds {
		embeds = append(embeds, buf2struct.MessageEmbed(embed))
	}

	message, err := h.Helper.ChannelMessageSendEmbeds(req.ChannelId, embeds)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendEmbedsResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageEdit handles editing a message with simple text content.
func (h *HelperServer) ChannelMessageEdit(ctx context.Context, req *proto.ChannelMessageEditRequest) (*proto.ChannelMessageEditResponse, error) {
	message, err := h.Helper.ChannelMessageEdit(req.ChannelId, req.MessageId, req.Content)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageEditResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageEditComplex handles editing a message with complex data.
func (h *HelperServer) ChannelMessageEditComplex(ctx context.Context, req *proto.ChannelMessageEditComplexRequest) (*proto.ChannelMessageEditComplexResponse, error) {
	messageEdit := buf2struct.MessageEdit(req.MessageEdit)
	message, err := h.Helper.ChannelMessageEditComplex(messageEdit)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageEditComplexResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageDelete handles deleting a message from a channel.
func (h *HelperServer) ChannelMessageDelete(ctx context.Context, req *proto.ChannelMessageDeleteRequest) (*proto_common.Empty, error) {
	err := h.Helper.ChannelMessageDelete(req.ChannelId, req.MessageId)
	if err != nil {
		return nil, err
	}

	return &proto_common.Empty{}, nil
}

// ChannelMessages handles retrieving multiple messages from a channel.
func (h *HelperServer) ChannelMessages(ctx context.Context, req *proto.ChannelMessagesRequest) (*proto.ChannelMessagesResponse, error) {
	messages, err := h.Helper.ChannelMessages(req.ChannelId, int(req.Limit), req.BeforeId, req.AfterId, req.AroundId)
	if err != nil {
		return nil, err
	}

	protoMessages := make([]*proto.Message, 0, len(messages))
	for _, msg := range messages {
		protoMessages = append(protoMessages, struct2buf.Message(msg))
	}

	return &proto.ChannelMessagesResponse{
		Messages: protoMessages,
	}, nil
}

// ChannelMessage handles retrieving a single message from a channel.
func (h *HelperServer) ChannelMessage(ctx context.Context, req *proto.ChannelMessageRequest) (*proto.ChannelMessageResponse, error) {
	message, err := h.Helper.ChannelMessage(req.ChannelId, req.MessageId)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageResponse{
		Message: struct2buf.Message(message),
	}, nil
}
