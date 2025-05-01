package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

func ApplicationCmd(s *discordgo.ApplicationCommand) *proto.ApplicationCommand {
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

	return &proto.ApplicationCommand{
		Id:                s.ID,
		ApplicationId:     s.ApplicationID,
		GuildId:           s.GuildID,
		Version:           s.Version,
		Type:              cmdType,
		Name:              s.Name,
		NameLocalizations: nil, // TODO(discord/bufstruct): implements NameLocalizations
		// DefaultPermission:        s.DefaultPermission, // Deprecated
		DefaultMemberPermissions: defaultMemberPermissions,
		// DmPermission:             s.DMPermission, // Deprecated
		Nsfw:                     s.NSFW,
		Description:              s.Description,
		DescriptionLocalizations: nil, // TODO(discord/bufstruct): implements DescriptionLocalizations
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
		User:                         User(s.User),
		Data:                         nil, // TODO(discord/bufstruct): implements Data
		AppPermissions:               s.AppPermissions,
		Member:                       Member(s.Member),
		Locale:                       string(s.Locale),
		GuildLocale:                  guildLocale,
		Context:                      0,
		AuthorizingIntegrationOwners: nil,
		Token:                        "",
		Version:                      0,
		Entitlements:                 nil,
	}

	if s.Data != nil {
		switch s.Data.Type() {
		case discordgo.InteractionApplicationCommand:
			data := s.Data.(discordgo.ApplicationCommandInteractionData)
			result.Data = &proto.Interaction_ApplicationCommandData{
				ApplicationCommandData: &proto.ApplicationCommandInteractionData{
					Id:          data.ID,
					Name:        data.Name,
					CommandType: proto.ApplicationCommandType(data.CommandType),
				},
			}
		}
	}

	return result
}

func AppCmdInteractionResolved() {
}
