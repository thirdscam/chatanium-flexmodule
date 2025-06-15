package module

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	plugin "github.com/hashicorp/go-plugin"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	shared "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/buf2struct"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/struct2buf"
)

// HelperClientImpl implements the Helper interface for module-side operations.
// This client communicates with the runtime server to perform Discord operations.
type HelperClientImpl struct {
	broker *plugin.GRPCBroker
	client proto.HelperClient
}

// ================================================
// Message operations (implemented in proto)
// ================================================

// ChannelMessageSend sends a simple text message to a channel.
func (h *HelperClientImpl) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
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
func (h *HelperClientImpl) ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (*discordgo.Message, error) {
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
func (h *HelperClientImpl) ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
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
func (h *HelperClientImpl) ChannelMessageSendEmbeds(channelID string, embeds []*discordgo.MessageEmbed) (*discordgo.Message, error) {
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
func (h *HelperClientImpl) ChannelMessageEdit(channelID, messageID, content string) (*discordgo.Message, error) {
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
func (h *HelperClientImpl) ChannelMessageEditComplex(m *discordgo.MessageEdit) (*discordgo.Message, error) {
	resp, err := h.client.ChannelMessageEditComplex(context.Background(), &proto.ChannelMessageEditComplexRequest{
		MessageEdit: struct2buf.MessageEdit(m),
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}

// ChannelMessageDelete deletes a message from a channel.
func (h *HelperClientImpl) ChannelMessageDelete(channelID, messageID string) error {
	_, err := h.client.ChannelMessageDelete(context.Background(), &proto.ChannelMessageDeleteRequest{
		ChannelId: channelID,
		MessageId: messageID,
	})
	return err
}

// ChannelMessages retrieves multiple messages from a channel.
func (h *HelperClientImpl) ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string) ([]*discordgo.Message, error) {
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
func (h *HelperClientImpl) ChannelMessage(channelID, messageID string) (*discordgo.Message, error) {
	resp, err := h.client.ChannelMessage(context.Background(), &proto.ChannelMessageRequest{
		ChannelId: channelID,
		MessageId: messageID,
	})
	if err != nil {
		return nil, err
	}
	return buf2struct.Message(resp.Message), nil
}

// ================================================
// Channel operations (not implemented in proto - stub implementations)
// ================================================

// Channel retrieves information about a channel.
func (h *HelperClientImpl) Channel(channelID string) (*discordgo.Channel, error) {
	return nil, fmt.Errorf("Channel operation not implemented in gRPC proto")
}

// ChannelEdit modifies a channel's properties.
func (h *HelperClientImpl) ChannelEdit(channelID string, data *discordgo.ChannelEdit) (*discordgo.Channel, error) {
	return nil, fmt.Errorf("ChannelEdit operation not implemented in gRPC proto")
}

// ChannelDelete deletes a channel.
func (h *HelperClientImpl) ChannelDelete(channelID string) (*discordgo.Channel, error) {
	return nil, fmt.Errorf("ChannelDelete operation not implemented in gRPC proto")
}

// ChannelTyping triggers typing indicator in a channel.
func (h *HelperClientImpl) ChannelTyping(channelID string) error {
	return fmt.Errorf("ChannelTyping operation not implemented in gRPC proto")
}

// ================================================
// Operations not implemented in proto - stub implementations
// ================================================

// Guild retrieves information about a guild.
func (h *HelperClientImpl) Guild(guildID string) (*discordgo.Guild, error) {
	return nil, fmt.Errorf("Guild operation not implemented in gRPC proto")
}

// GuildChannels retrieves all channels in a guild.
func (h *HelperClientImpl) GuildChannels(guildID string) ([]*discordgo.Channel, error) {
	return nil, fmt.Errorf("GuildChannels operation not implemented in gRPC proto")
}

// GuildMembers retrieves guild members.
func (h *HelperClientImpl) GuildMembers(guildID string, after string, limit int) ([]*discordgo.Member, error) {
	return nil, fmt.Errorf("GuildMembers operation not implemented in gRPC proto")
}

// GuildMember retrieves a specific guild member.
func (h *HelperClientImpl) GuildMember(guildID, userID string) (*discordgo.Member, error) {
	return nil, fmt.Errorf("GuildMember operation not implemented in gRPC proto")
}

// GuildRoles retrieves all roles in a guild.
func (h *HelperClientImpl) GuildRoles(guildID string) ([]*discordgo.Role, error) {
	return nil, fmt.Errorf("GuildRoles operation not implemented in gRPC proto")
}

// User retrieves information about a user.
func (h *HelperClientImpl) User(userID string) (*discordgo.User, error) {
	return nil, fmt.Errorf("User operation not implemented in gRPC proto")
}

// UserChannelCreate creates a DM channel with a user.
func (h *HelperClientImpl) UserChannelCreate(recipientID string) (*discordgo.Channel, error) {
	return nil, fmt.Errorf("UserChannelCreate operation not implemented in gRPC proto")
}

// InteractionRespond responds to an interaction.
func (h *HelperClientImpl) InteractionRespond(interaction *discordgo.Interaction, resp *discordgo.InteractionResponse) error {
	return fmt.Errorf("InteractionRespond operation not implemented in gRPC proto")
}

// InteractionResponseEdit edits an interaction response.
func (h *HelperClientImpl) InteractionResponseEdit(interaction *discordgo.Interaction, newresp *discordgo.WebhookEdit) (*discordgo.Message, error) {
	return nil, fmt.Errorf("InteractionResponseEdit operation not implemented in gRPC proto")
}

// ApplicationCommandCreate creates a new application command.
func (h *HelperClientImpl) ApplicationCommandCreate(appID string, guildID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	return nil, fmt.Errorf("ApplicationCommandCreate operation not implemented in gRPC proto")
}

// ApplicationCommandEdit edits an existing application command.
func (h *HelperClientImpl) ApplicationCommandEdit(appID, guildID, cmdID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	return nil, fmt.Errorf("ApplicationCommandEdit operation not implemented in gRPC proto")
}

// ApplicationCommandDelete deletes an application command.
func (h *HelperClientImpl) ApplicationCommandDelete(appID, guildID, cmdID string) error {
	return fmt.Errorf("ApplicationCommandDelete operation not implemented in gRPC proto")
}

// ApplicationCommands retrieves all application commands.
func (h *HelperClientImpl) ApplicationCommands(appID, guildID string) ([]*discordgo.ApplicationCommand, error) {
	return nil, fmt.Errorf("ApplicationCommands operation not implemented in gRPC proto")
}

// MessageReactionAdd adds a reaction to a message.
func (h *HelperClientImpl) MessageReactionAdd(channelID, messageID, emojiID string) error {
	return fmt.Errorf("MessageReactionAdd operation not implemented in gRPC proto")
}

// MessageReactionRemove removes a reaction from a message.
func (h *HelperClientImpl) MessageReactionRemove(channelID, messageID, emojiID, userID string) error {
	return fmt.Errorf("MessageReactionRemove operation not implemented in gRPC proto")
}

// MessageReactionsRemoveAll removes all reactions from a message.
func (h *HelperClientImpl) MessageReactionsRemoveAll(channelID, messageID string) error {
	return fmt.Errorf("MessageReactionsRemoveAll operation not implemented in gRPC proto")
}

// ThreadStart creates a new thread.
func (h *HelperClientImpl) ThreadStart(channelID, name string, typ discordgo.ChannelType, archiveDuration int) (*discordgo.Channel, error) {
	return nil, fmt.Errorf("ThreadStart operation not implemented in gRPC proto")
}

// ThreadJoin joins a thread.
func (h *HelperClientImpl) ThreadJoin(threadID string) error {
	return fmt.Errorf("ThreadJoin operation not implemented in gRPC proto")
}

// ThreadLeave leaves a thread.
func (h *HelperClientImpl) ThreadLeave(threadID string) error {
	return fmt.Errorf("ThreadLeave operation not implemented in gRPC proto")
}

// ThreadMemberAdd adds a member to a thread.
func (h *HelperClientImpl) ThreadMemberAdd(threadID, memberID string) error {
	return fmt.Errorf("ThreadMemberAdd operation not implemented in gRPC proto")
}

// ThreadMemberRemove removes a member from a thread.
func (h *HelperClientImpl) ThreadMemberRemove(threadID, memberID string) error {
	return fmt.Errorf("ThreadMemberRemove operation not implemented in gRPC proto")
}

// VoiceRegions retrieves available voice regions.
func (h *HelperClientImpl) VoiceRegions() ([]*discordgo.VoiceRegion, error) {
	return nil, fmt.Errorf("VoiceRegions operation not implemented in gRPC proto")
}

// WebhookCreate creates a new webhook.
func (h *HelperClientImpl) WebhookCreate(channelID, name, avatar string) (*discordgo.Webhook, error) {
	return nil, fmt.Errorf("WebhookCreate operation not implemented in gRPC proto")
}

// WebhookExecute executes a webhook.
func (h *HelperClientImpl) WebhookExecute(webhookID, token string, wait bool, data *discordgo.WebhookParams) (*discordgo.Message, error) {
	return nil, fmt.Errorf("WebhookExecute operation not implemented in gRPC proto")
}

// UserChannelPermissions retrieves user permissions for a channel.
func (h *HelperClientImpl) UserChannelPermissions(userID, channelID string) (int64, error) {
	return 0, fmt.Errorf("UserChannelPermissions operation not implemented in gRPC proto")
}

// Gateway retrieves gateway URL.
func (h *HelperClientImpl) Gateway() (string, error) {
	return "", fmt.Errorf("Gateway operation not implemented in gRPC proto")
}

// GatewayBot retrieves gateway bot information.
func (h *HelperClientImpl) GatewayBot() (*discordgo.GatewayBotResponse, error) {
	return nil, fmt.Errorf("GatewayBot operation not implemented in gRPC proto")
}

// Ensure HelperClientImpl implements the Helper interface
var _ shared.Helper = &HelperClientImpl{}
