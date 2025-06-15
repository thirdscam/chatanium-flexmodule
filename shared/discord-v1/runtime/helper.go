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

// HelperServerImpl implements the Helper gRPC server for the runtime-side.
// This server receives calls from the module and provides actual Discord operations.
type HelperServerImpl struct {
	proto.UnimplementedHelperServer
	Impl   shared.Helper // Helper implementation that provides Discord operations
	broker *plugin.GRPCBroker
}

// ================================================
// Message operations (implemented in proto)
// ================================================

// ChannelMessageSend handles sending a simple text message.
func (h *HelperServerImpl) ChannelMessageSend(ctx context.Context, req *proto.ChannelMessageSendRequest) (*proto.ChannelMessageSendResponse, error) {
	message, err := h.Impl.ChannelMessageSend(req.ChannelId, req.Content)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageSendResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageSendComplex handles sending a complex message with attachments, embeds, etc.
func (h *HelperServerImpl) ChannelMessageSendComplex(ctx context.Context, req *proto.ChannelMessageSendComplexRequest) (*proto.ChannelMessageSendComplexResponse, error) {
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
func (h *HelperServerImpl) ChannelMessageSendEmbed(ctx context.Context, req *proto.ChannelMessageSendEmbedRequest) (*proto.ChannelMessageSendEmbedResponse, error) {
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
func (h *HelperServerImpl) ChannelMessageSendEmbeds(ctx context.Context, req *proto.ChannelMessageSendEmbedsRequest) (*proto.ChannelMessageSendEmbedsResponse, error) {
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
func (h *HelperServerImpl) ChannelMessageEdit(ctx context.Context, req *proto.ChannelMessageEditRequest) (*proto.ChannelMessageEditResponse, error) {
	message, err := h.Impl.ChannelMessageEdit(req.ChannelId, req.MessageId, req.Content)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageEditResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ChannelMessageEditComplex handles editing a message with complex data.
func (h *HelperServerImpl) ChannelMessageEditComplex(ctx context.Context, req *proto.ChannelMessageEditComplexRequest) (*proto.ChannelMessageEditComplexResponse, error) {
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
func (h *HelperServerImpl) ChannelMessageDelete(ctx context.Context, req *proto.ChannelMessageDeleteRequest) (*proto_common.Empty, error) {
	err := h.Impl.ChannelMessageDelete(req.ChannelId, req.MessageId)
	if err != nil {
		return nil, err
	}

	return &proto_common.Empty{}, nil
}

// ChannelMessages handles retrieving multiple messages from a channel.
func (h *HelperServerImpl) ChannelMessages(ctx context.Context, req *proto.ChannelMessagesRequest) (*proto.ChannelMessagesResponse, error) {
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
func (h *HelperServerImpl) ChannelMessage(ctx context.Context, req *proto.ChannelMessageRequest) (*proto.ChannelMessageResponse, error) {
	message, err := h.Impl.ChannelMessage(req.ChannelId, req.MessageId)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelMessageResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ================================================
// Channel operations (implemented in proto)
// ================================================

// Channel handles getting a channel by ID.
func (h *HelperServerImpl) Channel(ctx context.Context, req *proto.ChannelRequest) (*proto.ChannelResponse, error) {
	channel, err := h.Impl.Channel(req.ChannelId)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelResponse{
		Channel: struct2buf.Channel(channel),
	}, nil
}

// ChannelEdit handles editing a channel.
func (h *HelperServerImpl) ChannelEdit(ctx context.Context, req *proto.ChannelEditRequest) (*proto.ChannelEditResponse, error) {
	data := buf2struct.ChannelEdit(req.Data)
	channel, err := h.Impl.ChannelEdit(req.ChannelId, data)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelEditResponse{
		Channel: struct2buf.Channel(channel),
	}, nil
}

// ChannelDelete handles deleting a channel.
func (h *HelperServerImpl) ChannelDelete(ctx context.Context, req *proto.ChannelDeleteRequest) (*proto.ChannelDeleteResponse, error) {
	channel, err := h.Impl.ChannelDelete(req.ChannelId)
	if err != nil {
		return nil, err
	}

	return &proto.ChannelDeleteResponse{
		Channel: struct2buf.Channel(channel),
	}, nil
}

// ChannelTyping handles sending typing indicator.
func (h *HelperServerImpl) ChannelTyping(ctx context.Context, req *proto.ChannelTypingRequest) (*proto_common.Empty, error) {
	err := h.Impl.ChannelTyping(req.ChannelId)
	if err != nil {
		return nil, err
	}

	return &proto_common.Empty{}, nil
}

// ================================================
// Guild operations (implemented in proto)
// ================================================

// Guild handles getting a guild by ID.
func (h *HelperServerImpl) Guild(ctx context.Context, req *proto.GuildRequest) (*proto.GuildResponse, error) {
	guild, err := h.Impl.Guild(req.GuildId)
	if err != nil {
		return nil, err
	}

	return &proto.GuildResponse{
		Guild: struct2buf.Guild(guild),
	}, nil
}

// GuildChannels handles getting channels in a guild.
func (h *HelperServerImpl) GuildChannels(ctx context.Context, req *proto.GuildChannelsRequest) (*proto.GuildChannelsResponse, error) {
	channels, err := h.Impl.GuildChannels(req.GuildId)
	if err != nil {
		return nil, err
	}

	var protoChannels []*proto.Channel
	for _, channel := range channels {
		protoChannels = append(protoChannels, struct2buf.Channel(channel))
	}

	return &proto.GuildChannelsResponse{
		Channels: protoChannels,
	}, nil
}

// GuildMembers handles getting members in a guild.
func (h *HelperServerImpl) GuildMembers(ctx context.Context, req *proto.GuildMembersRequest) (*proto.GuildMembersResponse, error) {
	members, err := h.Impl.GuildMembers(req.GuildId, req.After, int(req.Limit))
	if err != nil {
		return nil, err
	}

	var protoMembers []*proto.Member
	for _, member := range members {
		protoMembers = append(protoMembers, struct2buf.Member(member))
	}

	return &proto.GuildMembersResponse{
		Members: protoMembers,
	}, nil
}

// GuildMember handles getting a specific member in a guild.
func (h *HelperServerImpl) GuildMember(ctx context.Context, req *proto.GuildMemberRequest) (*proto.GuildMemberResponse, error) {
	member, err := h.Impl.GuildMember(req.GuildId, req.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.GuildMemberResponse{
		Member: struct2buf.Member(member),
	}, nil
}

// GuildRoles handles getting roles in a guild.
func (h *HelperServerImpl) GuildRoles(ctx context.Context, req *proto.GuildRolesRequest) (*proto.GuildRolesResponse, error) {
	roles, err := h.Impl.GuildRoles(req.GuildId)
	if err != nil {
		return nil, err
	}

	var protoRoles []*proto.Role
	for _, role := range roles {
		protoRoles = append(protoRoles, struct2buf.Role(role))
	}

	return &proto.GuildRolesResponse{
		Roles: protoRoles,
	}, nil
}

// ================================================
// User operations (implemented in proto)
// ================================================

// User handles getting a user by ID.
func (h *HelperServerImpl) User(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	user, err := h.Impl.User(req.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		User: struct2buf.User(user),
	}, nil
}

// UserChannelCreate handles creating a DM channel with a user.
func (h *HelperServerImpl) UserChannelCreate(ctx context.Context, req *proto.UserChannelCreateRequest) (*proto.UserChannelCreateResponse, error) {
	channel, err := h.Impl.UserChannelCreate(req.RecipientId)
	if err != nil {
		return nil, err
	}

	return &proto.UserChannelCreateResponse{
		Channel: struct2buf.Channel(channel),
	}, nil
}

// ================================================
// Interaction operations (implemented in proto)
// ================================================

// InteractionRespond handles responding to an interaction.
func (h *HelperServerImpl) InteractionRespond(ctx context.Context, req *proto.InteractionRespondRequest) (*proto_common.Empty, error) {
	interaction := buf2struct.Interaction(req.Interaction)
	response := buf2struct.InteractionResponse(req.Response)

	err := h.Impl.InteractionRespond(interaction, response)
	if err != nil {
		return nil, err
	}

	return &proto_common.Empty{}, nil
}

// InteractionResponseEdit handles editing an interaction response.
func (h *HelperServerImpl) InteractionResponseEdit(ctx context.Context, req *proto.InteractionResponseEditRequest) (*proto.InteractionResponseEditResponse, error) {
	interaction := buf2struct.Interaction(req.Interaction)
	edit := buf2struct.WebhookEdit(req.WebhookEdit)

	message, err := h.Impl.InteractionResponseEdit(interaction, edit)
	if err != nil {
		return nil, err
	}

	return &proto.InteractionResponseEditResponse{
		Message: struct2buf.Message(message),
	}, nil
}

// ================================================
// Helper Client for runtime (stub implementation for interface compliance)
// ================================================

// HelperClientImpl implements the Helper interface for runtime-side operations.
// This is usually not needed since runtime provides helpers, but included for completeness.
type HelperClientImpl struct {
	broker *plugin.GRPCBroker
	client proto.HelperClient
}

// Minimal implementation for interface compliance - just enough to satisfy shared.Helper interface
func (h *HelperClientImpl) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelMessageSendEmbeds(channelID string, embeds []*discordgo.MessageEmbed) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelMessageEdit(channelID, messageID, content string) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelMessageEditComplex(m *discordgo.MessageEdit) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelMessageDelete(channelID, messageID string) error {
	return nil
}

func (h *HelperClientImpl) ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string) ([]*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelMessage(channelID, messageID string) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) Channel(channelID string) (*discordgo.Channel, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelEdit(channelID string, data *discordgo.ChannelEdit) (*discordgo.Channel, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelDelete(channelID string) (*discordgo.Channel, error) {
	return nil, nil
}

func (h *HelperClientImpl) ChannelTyping(channelID string) error {
	return nil
}

func (h *HelperClientImpl) Guild(guildID string) (*discordgo.Guild, error) {
	return nil, nil
}

func (h *HelperClientImpl) GuildChannels(guildID string) ([]*discordgo.Channel, error) {
	return nil, nil
}

func (h *HelperClientImpl) GuildMembers(guildID string, after string, limit int) ([]*discordgo.Member, error) {
	return nil, nil
}

func (h *HelperClientImpl) GuildMember(guildID, userID string) (*discordgo.Member, error) {
	return nil, nil
}

func (h *HelperClientImpl) GuildRoles(guildID string) ([]*discordgo.Role, error) {
	return nil, nil
}

func (h *HelperClientImpl) User(userID string) (*discordgo.User, error) {
	return nil, nil
}

func (h *HelperClientImpl) UserChannelCreate(recipientID string) (*discordgo.Channel, error) {
	return nil, nil
}

func (h *HelperClientImpl) InteractionRespond(interaction *discordgo.Interaction, resp *discordgo.InteractionResponse) error {
	return nil
}

func (h *HelperClientImpl) InteractionResponseEdit(interaction *discordgo.Interaction, newresp *discordgo.WebhookEdit) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) ApplicationCommandCreate(appID string, guildID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	return nil, nil
}

func (h *HelperClientImpl) ApplicationCommandEdit(appID, guildID, cmdID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	return nil, nil
}

func (h *HelperClientImpl) ApplicationCommandDelete(appID, guildID, cmdID string) error {
	return nil
}

func (h *HelperClientImpl) ApplicationCommands(appID, guildID string) ([]*discordgo.ApplicationCommand, error) {
	return nil, nil
}

func (h *HelperClientImpl) MessageReactionAdd(channelID, messageID, emojiID string) error {
	return nil
}

func (h *HelperClientImpl) MessageReactionRemove(channelID, messageID, emojiID, userID string) error {
	return nil
}

func (h *HelperClientImpl) MessageReactionsRemoveAll(channelID, messageID string) error {
	return nil
}

func (h *HelperClientImpl) ThreadStart(channelID, name string, typ discordgo.ChannelType, archiveDuration int) (*discordgo.Channel, error) {
	return nil, nil
}

func (h *HelperClientImpl) ThreadJoin(threadID string) error {
	return nil
}

func (h *HelperClientImpl) ThreadLeave(threadID string) error {
	return nil
}

func (h *HelperClientImpl) ThreadMemberAdd(threadID, memberID string) error {
	return nil
}

func (h *HelperClientImpl) ThreadMemberRemove(threadID, memberID string) error {
	return nil
}

func (h *HelperClientImpl) VoiceRegions() ([]*discordgo.VoiceRegion, error) {
	return nil, nil
}

func (h *HelperClientImpl) WebhookCreate(channelID, name, avatar string) (*discordgo.Webhook, error) {
	return nil, nil
}

func (h *HelperClientImpl) WebhookExecute(webhookID, token string, wait bool, data *discordgo.WebhookParams) (*discordgo.Message, error) {
	return nil, nil
}

func (h *HelperClientImpl) UserChannelPermissions(userID, channelID string) (int64, error) {
	return 0, nil
}

func (h *HelperClientImpl) Gateway() (string, error) {
	return "", nil
}

func (h *HelperClientImpl) GatewayBot() (*discordgo.GatewayBotResponse, error) {
	return nil, nil
}

// Ensure HelperClientImpl implements the Helper interface
var _ shared.Helper = &HelperClientImpl{}
