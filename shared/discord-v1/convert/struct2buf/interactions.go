package struct2buf

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// The channel parameters of the interaction, etc. require a discordgo.Session.
//
// Therefore, need to assign it for use in some conversion functions,
// and some object conversion functions may not be available if a session is not provided.
var DiscordSession *discordgo.Session

func ApplicationCommmand(s *discordgo.ApplicationCommand) *proto.ApplicationCommand {
	if s == nil {
		return nil
	}

	var cmdType proto.ApplicationCommandType
	switch s.Type {
	case discordgo.ChatApplicationCommand:
		cmdType = 1
	case discordgo.UserApplicationCommand:
		cmdType = 2
	case discordgo.MessageApplicationCommand:
		cmdType = 3
	default:
		cmdType = 0
	}

	var defaultMemberPermissions int64
	if s.DefaultMemberPermissions != nil {
		defaultMemberPermissions = *s.DefaultMemberPermissions
	}

	nameLocalizations := make(map[string]string)
	if s.NameLocalizations != nil {
		for k, v := range *s.NameLocalizations {
			nameLocalizations[string(k)] = v
		}
	}

	descriptionLocalizations := make(map[string]string)
	if s.DescriptionLocalizations != nil {
		for k, v := range *s.DescriptionLocalizations {
			descriptionLocalizations[string(k)] = v
		}
	}

	return &proto.ApplicationCommand{
		Id:                s.ID,
		ApplicationId:     s.ApplicationID,
		GuildId:           s.GuildID,
		Version:           s.Version,
		Type:              cmdType,
		Name:              s.Name,
		NameLocalizations: nameLocalizations,
		// DefaultPermission:        s.DefaultPermission, // Deprecated
		DefaultMemberPermissions: defaultMemberPermissions,
		// DmPermission:             s.DMPermission, // Deprecated
		Nsfw:                     s.NSFW,
		Description:              s.Description,
		DescriptionLocalizations: descriptionLocalizations,
		Options:                  nil, // TODO(discord/bufstruct): implements Options
	}
}

func Interaction(s *discordgo.Interaction) *proto.Interaction {
	if s == nil {
		return nil
	}

	var guildLocale *string
	if s.GuildLocale != nil {
		gl := string(*s.GuildLocale)
		guildLocale = &gl
	}

	result := &proto.Interaction{
		Id:                           s.ID,
		AppId:                        s.AppID,
		Type:                         proto.InteractionType(s.Type),
		GuildId:                      s.GuildID,
		ChannelId:                    s.ChannelID,
		Message:                      Message(s.Message),
		AppPermissions:               s.AppPermissions,
		Member:                       Member(s.Member),
		User:                         User(s.User),
		Locale:                       string(s.Locale),
		GuildLocale:                  guildLocale,
		Context:                      0,   // TODO(discord/bufstruct): Preparing for discordgo updates (^v0.28.1)
		AuthorizingIntegrationOwners: nil, // TODO(discord/bufstruct): Preparing for discordgo updates (^v0.28.1)
		Token:                        s.Token,
		Version:                      int32(s.Version),
		Entitlements:                 nil, // TODO(discord/bufstruct): Preparing for discordgo updates (^v0.28.1)
	}

	if s.Data != nil {
		switch s.Data.Type() {
		case discordgo.InteractionApplicationCommand, discordgo.InteractionApplicationCommandAutocomplete:
			data := s.Data.(discordgo.ApplicationCommandInteractionData)
			options := make([]*proto.ApplicationCommandInteractionDataOption, 0)
			for _, v := range data.Options {
				options = append(options, ApplicationCommandInteractionDataOption(v))
			}

			result.Data = &proto.Interaction_ApplicationCommandData{
				ApplicationCommandData: &proto.ApplicationCommandInteractionData{
					Id:          data.ID,
					Name:        data.Name,
					CommandType: proto.ApplicationCommandType(data.CommandType),
					Resolved:    ApplicationCommandInteractionDataResolved(data.Resolved),
					Options:     options,
					TargetId:    data.TargetID,
				},
			}
		case discordgo.InteractionMessageComponent:
			data := s.Data.(discordgo.MessageComponentInteractionData)
			result.Data = &proto.Interaction_MessageComponentData{
				MessageComponentData: &proto.MessageComponentInteractionData{
					CustomId:      data.CustomID,
					ComponentType: proto.ComponentType(data.ComponentType),
					Resolved:      MessageComponentInteractionDataResolved(&data.Resolved),
					Values:        data.Values,
				},
			}
		case discordgo.InteractionModalSubmit:
			data := s.Data.(discordgo.ModalSubmitInteractionData)
			result.Data = &proto.Interaction_ModalSubmitData{
				ModalSubmitData: &proto.ModalSubmitInteractionData{
					CustomId:   data.CustomID,
					Components: nil, // TODO(discord/bufstruct): implements MessageComponent
				},
			}

		}
	}

	return result
}

func ApplicationCommandInteractionData(s *discordgo.ApplicationCommandInteractionData) *proto.ApplicationCommandInteractionData {
	if s == nil {
		return nil
	}

	return &proto.ApplicationCommandInteractionData{
		Id:          s.ID,
		Name:        s.Name,
		CommandType: proto.ApplicationCommandType(s.CommandType),
		Resolved:    ApplicationCommandInteractionDataResolved(s.Resolved),
		TargetId:    s.TargetID,
	}
}

func MessageComponentInteractionDataResolved(s *discordgo.MessageComponentInteractionDataResolved) *proto.MessageComponentInteractionDataResolved {
	if s == nil {
		return nil
	}

	users := make(map[string]*proto.User)
	for k, v := range s.Users {
		users[k] = User(v)
	}

	members := make(map[string]*proto.Member)
	for k, v := range s.Members {
		members[k] = Member(v)
	}

	roles := make(map[string]*proto.Role)
	for k, v := range s.Roles {
		roles[k] = Role(v)
	}

	channels := make(map[string]*proto.Channel)
	for k, v := range s.Channels {
		channels[k] = Channel(v)
	}

	return &proto.MessageComponentInteractionDataResolved{
		Users:    users,
		Members:  members,
		Roles:    roles,
		Channels: channels,
	}
}

func ApplicationCommandInteractionDataResolved(s *discordgo.ApplicationCommandInteractionDataResolved) *proto.ApplicationCommandInteractionDataResolved {
	if s == nil {
		return nil
	}

	users := make(map[string]*proto.User)
	for k, v := range s.Users {
		users[k] = User(v)
	}

	roles := make(map[string]*proto.Role)
	for k, v := range s.Roles {
		roles[k] = Role(v)
	}

	members := make(map[string]*proto.Member)
	for k, v := range s.Members {
		members[k] = Member(v)
	}

	channels := make(map[string]*proto.Channel)
	for k, v := range s.Channels {
		channels[k] = Channel(v)
	}

	messages := make(map[string]*proto.Message)
	for k, v := range s.Messages {
		messages[k] = Message(v)
	}

	attachments := make(map[string]*proto.MessageAttachment)
	for k, v := range s.Attachments {
		attachments[k] = MessageAttachment(v)
	}

	return &proto.ApplicationCommandInteractionDataResolved{
		Users:       users,
		Members:     members,
		Roles:       roles,
		Channels:    channels,
		Messages:    messages,
		Attachments: attachments,
	}
}

func ApplicationCommandInteractionDataOption(s *discordgo.ApplicationCommandInteractionDataOption) *proto.ApplicationCommandInteractionDataOption {
	if s == nil {
		return nil
	}

	if DiscordSession == nil {
		fmt.Println("ApplicationCommandInteractionDataOption > DiscordSession is nil")
		return nil
	}

	options := make([]*proto.ApplicationCommandInteractionDataOption, 0)
	for _, v := range s.Options {
		options = append(options, ApplicationCommandInteractionDataOption(v))
	}

	result := &proto.ApplicationCommandInteractionDataOption{
		Name:    s.Name,
		Type:    proto.ApplicationCommandOptionType(s.Type),
		Options: options, // TODO(struct2buf): Consider same-type recursion issues
		Focused: s.Focused,
	}

	if s.Value != nil {
		switch s.Type {
		case discordgo.ApplicationCommandOptionInteger:
			result.Value = &proto.ApplicationCommandInteractionDataOption_IntegerValue{
				IntegerValue: s.IntValue(),
			}
		case discordgo.ApplicationCommandOptionNumber:
			result.Value = &proto.ApplicationCommandInteractionDataOption_NumberValue{
				NumberValue: s.FloatValue(),
			}
		case discordgo.ApplicationCommandOptionString:
			result.Value = &proto.ApplicationCommandInteractionDataOption_StringValue{
				StringValue: s.StringValue(),
			}
		case discordgo.ApplicationCommandOptionBoolean:
			result.Value = &proto.ApplicationCommandInteractionDataOption_BooleanValue{
				BooleanValue: s.BoolValue(),
			}
		case discordgo.ApplicationCommandOptionChannel:
			channel := s.ChannelValue(DiscordSession)
			if channel == nil {
				fmt.Println("ApplicationCommandInteractionDataOption > discordgo.Channel is nil, filling channel_id field with empty string")
				result.Value = &proto.ApplicationCommandInteractionDataOption_ChannelValueId{
					ChannelValueId: "",
				}
			} else {
				result.Value = &proto.ApplicationCommandInteractionDataOption_ChannelValueId{
					ChannelValueId: channel.ID,
				}
			}
		}
	}

	return result
}
