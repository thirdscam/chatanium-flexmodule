package discord

import (
	"github.com/bwmarrin/discordgo"
)

type InitResponse struct {
	// if bot needs to register interactions, write them here
	Interactions []*discordgo.ApplicationCommand
}

type Hook interface {
	OnInit(h Helper) InitResponse
	OnCreateChatMessage(message *discordgo.Message) error
	OnCreateInteraction(interaction *discordgo.Interaction) error
	OnEvent(event string) error
}

// RuntimeClients represents the client interfaces for runtime
type RuntimeClients interface {
	GetHook() Hook
	GetHelper() Helper
}

type Helper interface {
	// Message operations
	ChannelMessageSend(channelID string, content string) (*discordgo.Message, error)
	ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (*discordgo.Message, error)
	ChannelMessageSendEmbed(channelID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error)
	ChannelMessageSendEmbeds(channelID string, embeds []*discordgo.MessageEmbed) (*discordgo.Message, error)
	ChannelMessageEdit(channelID, messageID, content string) (*discordgo.Message, error)
	ChannelMessageEditComplex(m *discordgo.MessageEdit) (*discordgo.Message, error)
	ChannelMessageDelete(channelID, messageID string) error
	ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string) ([]*discordgo.Message, error)
	ChannelMessage(channelID, messageID string) (*discordgo.Message, error)

	// Channel operations
	Channel(channelID string) (*discordgo.Channel, error)
	ChannelEdit(channelID string, data *discordgo.ChannelEdit) (*discordgo.Channel, error)
	ChannelDelete(channelID string) (*discordgo.Channel, error)
	ChannelTyping(channelID string) error

	// Guild operations
	Guild(guildID string) (*discordgo.Guild, error)
	GuildChannels(guildID string) ([]*discordgo.Channel, error)
	GuildMembers(guildID string, after string, limit int) ([]*discordgo.Member, error)
	GuildMember(guildID, userID string) (*discordgo.Member, error)
	GuildRoles(guildID string) ([]*discordgo.Role, error)

	// User operations
	User(userID string) (*discordgo.User, error)
	UserChannelCreate(recipientID string) (*discordgo.Channel, error)

	// Interaction operations
	InteractionRespond(interaction *discordgo.Interaction, resp *discordgo.InteractionResponse) error
	InteractionResponseEdit(interaction *discordgo.Interaction, newresp *discordgo.WebhookEdit) (*discordgo.Message, error)

	// Application Command operations
	ApplicationCommandCreate(appID string, guildID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error)
	ApplicationCommandEdit(appID, guildID, cmdID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error)
	ApplicationCommandDelete(appID, guildID, cmdID string) error
	ApplicationCommands(appID, guildID string) ([]*discordgo.ApplicationCommand, error)

	// Reaction operations
	MessageReactionAdd(channelID, messageID, emojiID string) error
	MessageReactionRemove(channelID, messageID, emojiID, userID string) error
	MessageReactionsRemoveAll(channelID, messageID string) error

	// Thread operations
	ThreadStart(channelID, name string, typ discordgo.ChannelType, archiveDuration int) (*discordgo.Channel, error)
	ThreadJoin(threadID string) error
	ThreadLeave(threadID string) error
	ThreadMemberAdd(threadID, memberID string) error
	ThreadMemberRemove(threadID, memberID string) error

	// Voice operations
	VoiceRegions() ([]*discordgo.VoiceRegion, error)

	// Webhook operations
	WebhookCreate(channelID, name, avatar string) (*discordgo.Webhook, error)
	WebhookExecute(webhookID, token string, wait bool, data *discordgo.WebhookParams) (*discordgo.Message, error)

	// Permission operations
	UserChannelPermissions(userID, channelID string) (int64, error)

	// Utility operations
	Gateway() (string, error)
	GatewayBot() (*discordgo.GatewayBotResponse, error)
}

// AbstractHooks is a partial implementation of the hook Interface.
//
// It is useful for embedding in a struct that only needs to
// implement a subset of the hook Interface.
type AbstractHooks struct{}

func (u *AbstractHooks) OnInit(h Helper) InitResponse {
	return InitResponse{}
}

func (u *AbstractHooks) OnCreateChatMessage(m *discordgo.Message) error {
	return nil
}

func (u *AbstractHooks) OnCreateInteraction(i *discordgo.Interaction) error {
	return nil
}

func (u *AbstractHooks) OnEvent(e string) error {
	return nil
}
