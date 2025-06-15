package module

import (
	"context"

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
	resp, err := h.client.Channel(context.Background(), &proto.ChannelRequest{
		ChannelId: channelID,
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Channel(resp.Channel), nil
}

// ChannelEdit modifies a channel's properties.
func (h *HelperClientImpl) ChannelEdit(channelID string, data *discordgo.ChannelEdit) (*discordgo.Channel, error) {
	resp, err := h.client.ChannelEdit(context.Background(), &proto.ChannelEditRequest{
		ChannelId: channelID,
		Data:      struct2buf.ChannelEdit(data),
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Channel(resp.Channel), nil
}

// ChannelDelete deletes a channel.
func (h *HelperClientImpl) ChannelDelete(channelID string) (*discordgo.Channel, error) {
	resp, err := h.client.ChannelDelete(context.Background(), &proto.ChannelDeleteRequest{
		ChannelId: channelID,
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Channel(resp.Channel), nil
}

// ChannelTyping triggers typing indicator in a channel.
func (h *HelperClientImpl) ChannelTyping(channelID string) error {
	_, err := h.client.ChannelTyping(context.Background(), &proto.ChannelTypingRequest{
		ChannelId: channelID,
	})
	return err
}

// Guild retrieves information about a guild.
func (h *HelperClientImpl) Guild(guildID string) (*discordgo.Guild, error) {
	resp, err := h.client.Guild(context.Background(), &proto.GuildRequest{
		GuildId: guildID,
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Guild(resp.Guild), nil
}

// GuildChannels retrieves all channels in a guild.
func (h *HelperClientImpl) GuildChannels(guildID string) ([]*discordgo.Channel, error) {
	resp, err := h.client.GuildChannels(context.Background(), &proto.GuildChannelsRequest{
		GuildId: guildID,
	})
	if err != nil {
		return nil, err
	}

	var channels []*discordgo.Channel
	for _, protoChannel := range resp.Channels {
		channels = append(channels, buf2struct.Channel(protoChannel))
	}

	return channels, nil
}

// GuildMembers retrieves guild members.
func (h *HelperClientImpl) GuildMembers(guildID string, after string, limit int) ([]*discordgo.Member, error) {
	resp, err := h.client.GuildMembers(context.Background(), &proto.GuildMembersRequest{
		GuildId: guildID,
		After:   after,
		Limit:   int32(limit),
	})
	if err != nil {
		return nil, err
	}

	var members []*discordgo.Member
	for _, protoMember := range resp.Members {
		members = append(members, buf2struct.Member(protoMember))
	}

	return members, nil
}

// GuildMember retrieves a specific guild member.
func (h *HelperClientImpl) GuildMember(guildID, userID string) (*discordgo.Member, error) {
	resp, err := h.client.GuildMember(context.Background(), &proto.GuildMemberRequest{
		GuildId: guildID,
		UserId:  userID,
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Member(resp.Member), nil
}

// GuildRoles retrieves all roles in a guild.
func (h *HelperClientImpl) GuildRoles(guildID string) ([]*discordgo.Role, error) {
	resp, err := h.client.GuildRoles(context.Background(), &proto.GuildRolesRequest{
		GuildId: guildID,
	})
	if err != nil {
		return nil, err
	}

	var roles []*discordgo.Role
	for _, protoRole := range resp.Roles {
		roles = append(roles, buf2struct.Role(protoRole))
	}

	return roles, nil
}

// User retrieves information about a user.
func (h *HelperClientImpl) User(userID string) (*discordgo.User, error) {
	resp, err := h.client.User(context.Background(), &proto.UserRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.User(resp.User), nil
}

// UserChannelCreate creates a DM channel with a user.
func (h *HelperClientImpl) UserChannelCreate(recipientID string) (*discordgo.Channel, error) {
	resp, err := h.client.UserChannelCreate(context.Background(), &proto.UserChannelCreateRequest{
		RecipientId: recipientID,
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Channel(resp.Channel), nil
}

// InteractionRespond responds to an interaction.
func (h *HelperClientImpl) InteractionRespond(interaction *discordgo.Interaction, resp *discordgo.InteractionResponse) error {
	_, err := h.client.InteractionRespond(context.Background(), &proto.InteractionRespondRequest{
		Interaction: struct2buf.Interaction(interaction),
		Response:    struct2buf.InteractionResponse(resp),
	})
	return err
}

// InteractionResponseEdit edits an interaction response.
func (h *HelperClientImpl) InteractionResponseEdit(interaction *discordgo.Interaction, newresp *discordgo.WebhookEdit) (*discordgo.Message, error) {
	resp, err := h.client.InteractionResponseEdit(context.Background(), &proto.InteractionResponseEditRequest{
		Interaction: struct2buf.Interaction(interaction),
		WebhookEdit: struct2buf.WebhookEdit(newresp),
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Message(resp.Message), nil
}

// ApplicationCommandCreate creates a new application command.
func (h *HelperClientImpl) ApplicationCommandCreate(appID string, guildID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	resp, err := h.client.ApplicationCommandCreate(context.Background(), &proto.ApplicationCommandCreateRequest{
		AppId:   appID,
		GuildId: guildID,
		Command: struct2buf.ApplicationCommand(cmd),
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.ApplicationCommand(resp.Command), nil
}

// ApplicationCommandEdit edits an existing application command.
func (h *HelperClientImpl) ApplicationCommandEdit(appID, guildID, cmdID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	resp, err := h.client.ApplicationCommandEdit(context.Background(), &proto.ApplicationCommandEditRequest{
		AppId:   appID,
		GuildId: guildID,
		CmdId:   cmdID,
		Command: struct2buf.ApplicationCommand(cmd),
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.ApplicationCommand(resp.Command), nil
}

// ApplicationCommandDelete deletes an application command.
func (h *HelperClientImpl) ApplicationCommandDelete(appID, guildID, cmdID string) error {
	_, err := h.client.ApplicationCommandDelete(context.Background(), &proto.ApplicationCommandDeleteRequest{
		AppId:   appID,
		GuildId: guildID,
		CmdId:   cmdID,
	})
	return err
}

// ApplicationCommands retrieves all application commands.
func (h *HelperClientImpl) ApplicationCommands(appID, guildID string) ([]*discordgo.ApplicationCommand, error) {
	resp, err := h.client.ApplicationCommands(context.Background(), &proto.ApplicationCommandsRequest{
		AppId:   appID,
		GuildId: guildID,
	})
	if err != nil {
		return nil, err
	}

	var commands []*discordgo.ApplicationCommand
	for _, protoCmd := range resp.Commands {
		commands = append(commands, buf2struct.ApplicationCommand(protoCmd))
	}

	return commands, nil
}

// MessageReactionAdd adds a reaction to a message.
func (h *HelperClientImpl) MessageReactionAdd(channelID, messageID, emojiID string) error {
	_, err := h.client.MessageReactionAdd(context.Background(), &proto.MessageReactionAddRequest{
		ChannelId: channelID,
		MessageId: messageID,
		EmojiId:   emojiID,
	})
	return err
}

// MessageReactionRemove removes a reaction from a message.
func (h *HelperClientImpl) MessageReactionRemove(channelID, messageID, emojiID, userID string) error {
	_, err := h.client.MessageReactionRemove(context.Background(), &proto.MessageReactionRemoveRequest{
		ChannelId: channelID,
		MessageId: messageID,
		EmojiId:   emojiID,
		UserId:    userID,
	})
	return err
}

// MessageReactionsRemoveAll removes all reactions from a message.
func (h *HelperClientImpl) MessageReactionsRemoveAll(channelID, messageID string) error {
	_, err := h.client.MessageReactionsRemoveAll(context.Background(), &proto.MessageReactionsRemoveAllRequest{
		ChannelId: channelID,
		MessageId: messageID,
	})
	return err
}

// ThreadStart creates a new thread.
func (h *HelperClientImpl) ThreadStart(channelID, name string, typ discordgo.ChannelType, archiveDuration int) (*discordgo.Channel, error) {
	resp, err := h.client.ThreadStart(context.Background(), &proto.ThreadStartRequest{
		ChannelId:       channelID,
		Name:            name,
		Type:            int32(typ),
		ArchiveDuration: int32(archiveDuration),
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Channel(resp.Channel), nil
}

// ThreadJoin joins a thread.
func (h *HelperClientImpl) ThreadJoin(threadID string) error {
	_, err := h.client.ThreadJoin(context.Background(), &proto.ThreadJoinRequest{
		ThreadId: threadID,
	})
	return err
}

// ThreadLeave leaves a thread.
func (h *HelperClientImpl) ThreadLeave(threadID string) error {
	_, err := h.client.ThreadLeave(context.Background(), &proto.ThreadLeaveRequest{
		ThreadId: threadID,
	})
	return err
}

// ThreadMemberAdd adds a member to a thread.
func (h *HelperClientImpl) ThreadMemberAdd(threadID, memberID string) error {
	_, err := h.client.ThreadMemberAdd(context.Background(), &proto.ThreadMemberAddRequest{
		ThreadId: threadID,
		MemberId: memberID,
	})
	return err
}

// ThreadMemberRemove removes a member from a thread.
func (h *HelperClientImpl) ThreadMemberRemove(threadID, memberID string) error {
	_, err := h.client.ThreadMemberRemove(context.Background(), &proto.ThreadMemberRemoveRequest{
		ThreadId: threadID,
		MemberId: memberID,
	})
	return err
}

// VoiceRegions retrieves available voice regions.
func (h *HelperClientImpl) VoiceRegions() ([]*discordgo.VoiceRegion, error) {
	resp, err := h.client.VoiceRegions(context.Background(), &proto.VoiceRegionsRequest{})
	if err != nil {
		return nil, err
	}

	var regions []*discordgo.VoiceRegion
	for _, protoRegion := range resp.Regions {
		regions = append(regions, buf2struct.VoiceRegion(protoRegion))
	}

	return regions, nil
}

// WebhookCreate creates a new webhook.
func (h *HelperClientImpl) WebhookCreate(channelID, name, avatar string) (*discordgo.Webhook, error) {
	resp, err := h.client.WebhookCreate(context.Background(), &proto.WebhookCreateRequest{
		ChannelId: channelID,
		Name:      name,
		Avatar:    avatar,
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Webhook(resp.Webhook), nil
}

// WebhookExecute executes a webhook.
func (h *HelperClientImpl) WebhookExecute(webhookID, token string, wait bool, data *discordgo.WebhookParams) (*discordgo.Message, error) {
	resp, err := h.client.WebhookExecute(context.Background(), &proto.WebhookExecuteRequest{
		WebhookId: webhookID,
		Token:     token,
		Wait:      wait,
		Data:      struct2buf.WebhookParams(data),
	})
	if err != nil {
		return nil, err
	}

	return buf2struct.Message(resp.Message), nil
}

// UserChannelPermissions retrieves user permissions for a channel.
func (h *HelperClientImpl) UserChannelPermissions(userID, channelID string) (int64, error) {
	resp, err := h.client.UserChannelPermissions(context.Background(), &proto.UserChannelPermissionsRequest{
		UserId:    userID,
		ChannelId: channelID,
	})
	if err != nil {
		return 0, err
	}

	return resp.Permissions, nil
}

// Gateway retrieves gateway URL.
func (h *HelperClientImpl) Gateway() (string, error) {
	resp, err := h.client.Gateway(context.Background(), &proto.GatewayRequest{})
	if err != nil {
		return "", err
	}

	return resp.Url, nil
}

// GatewayBot retrieves gateway bot information.
func (h *HelperClientImpl) GatewayBot() (*discordgo.GatewayBotResponse, error) {
	resp, err := h.client.GatewayBot(context.Background(), &proto.GatewayBotRequest{})
	if err != nil {
		return nil, err
	}

	return buf2struct.GatewayBotResponse(resp), nil
}

// Ensure HelperClientImpl implements the Helper interface
var _ shared.Helper = &HelperClientImpl{}
