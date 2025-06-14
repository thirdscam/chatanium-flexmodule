package runtime

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

// GRPCServer implements the Helper gRPC server for the runtime-side.
// This server receives calls from the module and provides actual Discord operations.
type GRPCServer struct {
	Impl   shared.Helper // Helper implementation that provides Discord operations
	broker *plugin.GRPCBroker
}

// ================================================
// Message operations
// ================================================

// ChannelMessageSend handles sending a simple text message.
func (h *GRPCServer) ChannelMessageSend(ctx context.Context, req *proto.ChannelMessageSendRequest) (*proto.ChannelMessageSendResponse, error) {
	message, err := h.Impl.ChannelMessageSend(req.ChannelId, req.Content)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageSendComplex handles sending a complex message with attachments, embeds, etc.
func (h *GRPCServer) ChannelMessageSendComplex(ctx context.Context, req *proto.ChannelMessageSendComplexRequest) (*proto.ChannelMessageSendComplexResponse, error) {
	data := buf2struct.MessageSend(req.Data)
	message, err := h.Impl.ChannelMessageSendComplex(req.ChannelId, data)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendComplexResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageSendEmbed handles sending a message with a single embed.
func (h *GRPCServer) ChannelMessageSendEmbed(ctx context.Context, req *proto.ChannelMessageSendEmbedRequest) (*proto.ChannelMessageSendEmbedResponse, error) {
	embed := buf2struct.MessageEmbed(req.Embed)
	message, err := h.Impl.ChannelMessageSendEmbed(req.ChannelId, embed)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendEmbedResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageSendEmbeds handles sending a message with multiple embeds.
func (h *GRPCServer) ChannelMessageSendEmbeds(ctx context.Context, req *proto.ChannelMessageSendEmbedsRequest) (*proto.ChannelMessageSendEmbedsResponse, error) {
	embeds := make([]*discordgo.MessageEmbed, 0, len(req.Embeds))
	for _, embed := range req.Embeds {
		embeds = append(embeds, buf2struct.MessageEmbed(embed))
	}

	message, err := h.Impl.ChannelMessageSendEmbeds(req.ChannelId, embeds)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendEmbedsResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageEdit handles editing a message with simple text content.
func (h *GRPCServer) ChannelMessageEdit(ctx context.Context, req *proto.ChannelMessageEditRequest) (*proto.ChannelMessageEditResponse, error) {
	message, err := h.Impl.ChannelMessageEdit(req.ChannelId, req.MessageId, req.Content)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageEditResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageEditComplex handles editing a message with complex data.
func (h *GRPCServer) ChannelMessageEditComplex(ctx context.Context, req *proto.ChannelMessageEditComplexRequest) (*proto.ChannelMessageEditComplexResponse, error) {
	messageEdit := buf2struct.MessageEdit(req.MessageEdit)
	message, err := h.Impl.ChannelMessageEditComplex(messageEdit)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageEditComplexResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageDelete handles deleting a message from a channel.
func (h *GRPCServer) ChannelMessageDelete(ctx context.Context, req *proto.ChannelMessageDeleteRequest) (*proto_common.Empty, error) {
	err := h.Impl.ChannelMessageDelete(req.ChannelId, req.MessageId)
	if err != nil {
		return nil, err
	}

	return &proto_common.Empty{}, nil
}

// ChannelMessages handles retrieving multiple messages from a channel.
func (h *GRPCServer) ChannelMessages(ctx context.Context, req *proto.ChannelMessagesRequest) (*proto.ChannelMessagesResponse, error) {
	messages, err := h.Impl.ChannelMessages(req.ChannelId, int(req.Limit), req.BeforeId, req.AfterId, req.AroundId)
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
func (h *GRPCServer) ChannelMessage(ctx context.Context, req *proto.ChannelMessageRequest) (*proto.ChannelMessageResponse, error) {
	message, err := h.Impl.ChannelMessage(req.ChannelId, req.MessageId)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageResponse{
		Message: struct2buf.Message(message),
	}, nil
}
