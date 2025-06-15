package buf2struct

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/util"
)

func ApplicationCommand(buf *proto.ApplicationCommand) *discordgo.ApplicationCommand {
	if buf == nil {
		return nil
	}

	nameLocalizations := make(map[discordgo.Locale]string)
	if buf.NameLocalizations != nil {
		for k, v := range buf.NameLocalizations {
			nameLocalizations[util.StringToLocale(k)] = v
		}
	}

	descriptionLocalizations := make(map[discordgo.Locale]string)
	if buf.DescriptionLocalizations != nil {
		for k, v := range buf.DescriptionLocalizations {
			descriptionLocalizations[util.StringToLocale(k)] = v
		}
	}

	return &discordgo.ApplicationCommand{
		ID:                buf.Id,
		ApplicationID:     buf.ApplicationId,
		GuildID:           buf.GuildId,
		Version:           buf.Version,
		Type:              discordgo.ApplicationCommandType(buf.Type),
		Name:              buf.Name,
		NameLocalizations: &nameLocalizations,
		// DefaultPermission:        buf.DefaultPermission, // Deprecated
		DefaultMemberPermissions: &buf.DefaultMemberPermissions,
		// DMPermission:             buf.DmPermission, // Deprecated
		NSFW:                     buf.Nsfw,
		Description:              buf.Description,
		DescriptionLocalizations: &descriptionLocalizations,
		Options:                  nil, // TODO(discord/bufstruct): implements Options
	}
}

func Interaction(buf *proto.Interaction) *discordgo.Interaction {
	if buf == nil {
		return nil
	}

	var guildLocale *discordgo.Locale
	if buf.GuildLocale != nil {
		gl := util.StringToLocale(*buf.GuildLocale)
		guildLocale = &gl
	}

	interaction := &discordgo.Interaction{
		ID:             buf.Id,
		AppID:          buf.AppId,
		Type:           discordgo.InteractionType(buf.Type),
		GuildID:        buf.GuildId,
		ChannelID:      buf.ChannelId,
		Message:        Message(buf.Message),
		AppPermissions: buf.AppPermissions,
		Member:         Member(buf.Member),
		User:           User(buf.User),
		Locale:         util.StringToLocale(buf.Locale),
		GuildLocale:    guildLocale,
		// Context:                      0,   // TODO(discord/bufstruct): Preparing for discordgo updates (^v0.28.1)
		// AuthorizingIntegrationOwners: nil, // TODO(discord/bufstruct): Preparing for discordgo updates (^v0.28.1)
		Token:   buf.Token,
		Version: int(buf.Version),
		// Entitlements:                 nil, // TODO(discord/bufstruct): Preparing for discordgo updates (^v0.28.1)
	}

	switch buf.Data.(type) {
	case *proto.Interaction_ApplicationCommandData:
		appCmdData := buf.Data.(*proto.Interaction_ApplicationCommandData).ApplicationCommandData
		interaction.Data = *ApplicationCommandInteractionData(appCmdData)
	case *proto.Interaction_MessageComponentData:
		msgCompData := buf.Data.(*proto.Interaction_MessageComponentData).MessageComponentData
		interaction.Data = *MessageComponentInteractionData(msgCompData)
	case *proto.Interaction_ModalSubmitData:
		modalSubmitData := buf.Data.(*proto.Interaction_ModalSubmitData).ModalSubmitData
		interaction.Data = *ModalSubmitInteractionData(modalSubmitData)
	default:
		fmt.Println("Interaction > Unknown interaction data type")
	}

	return interaction
}

func ApplicationCommandInteractionData(buf *proto.ApplicationCommandInteractionData) *discordgo.ApplicationCommandInteractionData {
	if buf == nil {
		return nil
	}

	return &discordgo.ApplicationCommandInteractionData{
		ID:          buf.Id,
		Name:        buf.Name,
		CommandType: discordgo.ApplicationCommandType(buf.CommandType),
		Resolved:    ApplicationCommandInteractionDataResolved(buf.Resolved),
		Options:     nil,
		TargetID:    buf.TargetId,
	}
}

func ApplicationCommandInteractionDataOption(buf *proto.ApplicationCommandInteractionDataOption) *discordgo.ApplicationCommandInteractionDataOption {
	if buf == nil {
		return nil
	}

	options := make([]*discordgo.ApplicationCommandInteractionDataOption, 0)
	for _, v := range buf.Options {
		options = append(options, ApplicationCommandInteractionDataOption(v))
	}

	return &discordgo.ApplicationCommandInteractionDataOption{
		Name:    buf.Name,
		Type:    discordgo.ApplicationCommandOptionType(buf.Type),
		Value:   buf.Value,
		Options: options,
		Focused: buf.Focused,
	}
}

func ApplicationCommandInteractionDataResolved(buf *proto.ApplicationCommandInteractionDataResolved) *discordgo.ApplicationCommandInteractionDataResolved {
	if buf == nil {
		return nil
	}

	users := make(map[string]*discordgo.User)
	for _, v := range buf.Users {
		users[v.Id] = User(v)
	}

	members := make(map[string]*discordgo.Member)
	for _, v := range buf.Members {
		members[v.User.Id] = Member(v)
	}

	roles := make(map[string]*discordgo.Role)
	for _, v := range buf.Roles {
		roles[v.Id] = Role(v)
	}

	channels := make(map[string]*discordgo.Channel)
	for _, v := range buf.Channels {
		channels[v.Id] = Channel(v)
	}

	return &discordgo.ApplicationCommandInteractionDataResolved{
		Users:    users,
		Members:  members,
		Roles:    roles,
		Channels: channels,
	}
}

func MessageComponentInteractionData(buf *proto.MessageComponentInteractionData) *discordgo.MessageComponentInteractionData {
	if buf == nil {
		return nil
	}

	return &discordgo.MessageComponentInteractionData{
		CustomID:      buf.CustomId,
		ComponentType: discordgo.ComponentType(buf.ComponentType),
		Resolved:      *MessageComponentInteractionDataResolved(buf.Resolved),
		Values:        buf.Values,
	}
}

func MessageComponentInteractionDataResolved(buf *proto.MessageComponentInteractionDataResolved) *discordgo.MessageComponentInteractionDataResolved {
	if buf == nil {
		return nil
	}

	users := make(map[string]*discordgo.User)
	for k, v := range buf.Users {
		users[k] = User(v)
	}

	members := make(map[string]*discordgo.Member)
	for k, v := range buf.Members {
		members[k] = Member(v)
	}

	roles := make(map[string]*discordgo.Role)
	for k, v := range buf.Roles {
		roles[k] = Role(v)
	}

	channels := make(map[string]*discordgo.Channel)
	for k, v := range buf.Channels {
		channels[k] = Channel(v)
	}

	return &discordgo.MessageComponentInteractionDataResolved{
		Users:    users,
		Members:  members,
		Roles:    roles,
		Channels: channels,
	}
}

func ModalSubmitInteractionData(buf *proto.ModalSubmitInteractionData) *discordgo.ModalSubmitInteractionData {
	if buf == nil {
		return nil
	}

	return &discordgo.ModalSubmitInteractionData{
		CustomID:   buf.CustomId,
		Components: nil, // TODO(discord/bufstruct): implements Components
	}
}

// InteractionResponse converts proto InteractionResponse to discordgo InteractionResponse
func InteractionResponse(buf *proto.InteractionResponse) *discordgo.InteractionResponse {
	if buf == nil {
		return nil
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseType(buf.Type),
		Data: nil, // TODO: implement InteractionResponseData conversion
	}
}

// WebhookEdit converts proto WebhookEdit to discordgo WebhookEdit
func WebhookEdit(buf *proto.WebhookEdit) *discordgo.WebhookEdit {
	if buf == nil {
		return nil
	}

	// Basic implementation - more fields can be added as needed
	edit := &discordgo.WebhookEdit{}
	if buf.Content != nil {
		edit.Content = buf.Content
	}
	return edit
}
