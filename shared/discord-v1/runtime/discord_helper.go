package runtime

import (
	"github.com/bwmarrin/discordgo"
	shared "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
)

// DiscordHelper implements the Helper interface using an actual Discord session
type DiscordHelper struct {
	session *discordgo.Session
}

// NewDiscordHelper creates a new DiscordHelper with the given Discord session
func NewDiscordHelper(session *discordgo.Session) shared.Helper {
	return &DiscordHelper{
		session: session,
	}
}

// ================================================
// Message operations
// ================================================

func (h *DiscordHelper) ChannelMessageSend(channelID string, content string) (*discordgo.Message, error) {
	return h.session.ChannelMessageSend(channelID, content)
}

func (h *DiscordHelper) ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (*discordgo.Message, error) {
	return h.session.ChannelMessageSendComplex(channelID, data)
}

func (h *DiscordHelper) ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return h.session.ChannelMessageSendEmbed(channelID, embed)
}

func (h *DiscordHelper) ChannelMessageSendEmbeds(channelID string, embeds []*discordgo.MessageEmbed) (*discordgo.Message, error) {
	return h.session.ChannelMessageSendEmbeds(channelID, embeds)
}

func (h *DiscordHelper) ChannelMessageEdit(channelID, messageID, content string) (*discordgo.Message, error) {
	return h.session.ChannelMessageEdit(channelID, messageID, content)
}

func (h *DiscordHelper) ChannelMessageEditComplex(m *discordgo.MessageEdit) (*discordgo.Message, error) {
	return h.session.ChannelMessageEditComplex(m)
}

func (h *DiscordHelper) ChannelMessageDelete(channelID, messageID string) error {
	return h.session.ChannelMessageDelete(channelID, messageID)
}

func (h *DiscordHelper) ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string) ([]*discordgo.Message, error) {
	return h.session.ChannelMessages(channelID, limit, beforeID, afterID, aroundID)
}

func (h *DiscordHelper) ChannelMessage(channelID, messageID string) (*discordgo.Message, error) {
	return h.session.ChannelMessage(channelID, messageID)
}

// ================================================
// Channel operations
// ================================================

func (h *DiscordHelper) Channel(channelID string) (*discordgo.Channel, error) {
	return h.session.Channel(channelID)
}

func (h *DiscordHelper) ChannelEdit(channelID string, data *discordgo.ChannelEdit) (*discordgo.Channel, error) {
	return h.session.ChannelEdit(channelID, data)
}

func (h *DiscordHelper) ChannelDelete(channelID string) (*discordgo.Channel, error) {
	return h.session.ChannelDelete(channelID)
}

func (h *DiscordHelper) ChannelTyping(channelID string) error {
	return h.session.ChannelTyping(channelID)
}

// ================================================
// Guild operations
// ================================================

func (h *DiscordHelper) Guild(guildID string) (*discordgo.Guild, error) {
	return h.session.Guild(guildID)
}

func (h *DiscordHelper) GuildChannels(guildID string) ([]*discordgo.Channel, error) {
	return h.session.GuildChannels(guildID)
}

func (h *DiscordHelper) GuildMembers(guildID string, after string, limit int) ([]*discordgo.Member, error) {
	return h.session.GuildMembers(guildID, after, limit)
}

func (h *DiscordHelper) GuildMember(guildID, userID string) (*discordgo.Member, error) {
	return h.session.GuildMember(guildID, userID)
}

func (h *DiscordHelper) GuildRoles(guildID string) ([]*discordgo.Role, error) {
	return h.session.GuildRoles(guildID)
}

// ================================================
// User operations
// ================================================

func (h *DiscordHelper) User(userID string) (*discordgo.User, error) {
	return h.session.User(userID)
}

func (h *DiscordHelper) UserChannelCreate(recipientID string) (*discordgo.Channel, error) {
	return h.session.UserChannelCreate(recipientID)
}

// ================================================
// Interaction operations
// ================================================

func (h *DiscordHelper) InteractionRespond(interaction *discordgo.Interaction, resp *discordgo.InteractionResponse) error {
	return h.session.InteractionRespond(interaction, resp)
}

func (h *DiscordHelper) InteractionResponseEdit(interaction *discordgo.Interaction, newresp *discordgo.WebhookEdit) (*discordgo.Message, error) {
	return h.session.InteractionResponseEdit(interaction, newresp)
}

func (h *DiscordHelper) InteractionResponseDelete(interaction *discordgo.Interaction) error {
	return h.session.InteractionResponseDelete(interaction)
}

// ================================================
// Application Command operations
// ================================================

func (h *DiscordHelper) ApplicationCommandCreate(appID string, guildID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	return h.session.ApplicationCommandCreate(appID, guildID, cmd)
}

func (h *DiscordHelper) ApplicationCommandEdit(appID, guildID, cmdID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	return h.session.ApplicationCommandEdit(appID, guildID, cmdID, cmd)
}

func (h *DiscordHelper) ApplicationCommandDelete(appID, guildID, cmdID string) error {
	return h.session.ApplicationCommandDelete(appID, guildID, cmdID)
}

func (h *DiscordHelper) ApplicationCommands(appID, guildID string) ([]*discordgo.ApplicationCommand, error) {
	return h.session.ApplicationCommands(appID, guildID)
}

// ================================================
// Reaction operations
// ================================================

func (h *DiscordHelper) MessageReactionAdd(channelID, messageID, emojiID string) error {
	return h.session.MessageReactionAdd(channelID, messageID, emojiID)
}

func (h *DiscordHelper) MessageReactionRemove(channelID, messageID, emojiID, userID string) error {
	return h.session.MessageReactionRemove(channelID, messageID, emojiID, userID)
}

func (h *DiscordHelper) MessageReactionsRemoveAll(channelID, messageID string) error {
	return h.session.MessageReactionsRemoveAll(channelID, messageID)
}

// ================================================
// Thread operations
// ================================================

func (h *DiscordHelper) ThreadStart(channelID, name string, typ discordgo.ChannelType, archiveDuration int) (*discordgo.Channel, error) {
	return h.session.ThreadStart(channelID, name, typ, archiveDuration)
}

func (h *DiscordHelper) ThreadJoin(threadID string) error {
	return h.session.ThreadJoin(threadID)
}

func (h *DiscordHelper) ThreadLeave(threadID string) error {
	return h.session.ThreadLeave(threadID)
}

func (h *DiscordHelper) ThreadMemberAdd(threadID, memberID string) error {
	return h.session.ThreadMemberAdd(threadID, memberID)
}

func (h *DiscordHelper) ThreadMemberRemove(threadID, memberID string) error {
	return h.session.ThreadMemberRemove(threadID, memberID)
}

// ================================================
// Voice operations
// ================================================

func (h *DiscordHelper) VoiceRegions() ([]*discordgo.VoiceRegion, error) {
	return h.session.VoiceRegions()
}

// ================================================
// Webhook operations
// ================================================

func (h *DiscordHelper) WebhookCreate(channelID, name, avatar string) (*discordgo.Webhook, error) {
	return h.session.WebhookCreate(channelID, name, avatar)
}

func (h *DiscordHelper) WebhookExecute(webhookID, token string, wait bool, data *discordgo.WebhookParams) (*discordgo.Message, error) {
	return h.session.WebhookExecute(webhookID, token, wait, data)
}

// ================================================
// Permission operations
// ================================================

func (h *DiscordHelper) UserChannelPermissions(userID, channelID string) (int64, error) {
	return h.session.UserChannelPermissions(userID, channelID)
}

// ================================================
// Utility operations
// ================================================

func (h *DiscordHelper) Gateway() (string, error) {
	return h.session.Gateway()
}

func (h *DiscordHelper) GatewayBot() (*discordgo.GatewayBotResponse, error) {
	return h.session.GatewayBot()
}
