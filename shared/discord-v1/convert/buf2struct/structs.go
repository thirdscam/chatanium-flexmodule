package buf2struct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/util"
)

func User(buf *proto.User) *discordgo.User {
	if buf == nil {
		return nil
	}

	return &discordgo.User{
		ID:            buf.Id,
		Username:      buf.Username,
		Discriminator: buf.Discriminator,
		Avatar:        buf.Avatar,
		Bot:           buf.Bot,
		PublicFlags:   discordgo.UserFlags(buf.PublicFlagsValue),
		Locale:        buf.Locale,
		Verified:      buf.Verified,
		Email:         buf.Email,
		Flags:         int(buf.Flags),
	}
}

func Member(buf *proto.Member) *discordgo.Member {
	if buf == nil {
		return nil
	}

	return &discordgo.Member{
		GuildID:                    buf.GuildId,
		JoinedAt:                   buf.JoinedAt.AsTime(),
		Nick:                       buf.Nick,
		Deaf:                       buf.Deaf,
		Mute:                       buf.Mute,
		Avatar:                     buf.Avatar,
		Roles:                      buf.Roles,
		PremiumSince:               util.PbTimestamp2AsTimePtr(buf.PremiumSince),
		Flags:                      discordgo.MemberFlags(buf.Flags),
		Pending:                    buf.Pending,
		Permissions:                buf.Permissions,
		CommunicationDisabledUntil: util.PbTimestamp2AsTimePtr(buf.CommunicationDisabledUntil),
	}
}

func ForumTag(buf *proto.ForumTag) *discordgo.ForumTag {
	if buf == nil {
		return nil
	}

	return &discordgo.ForumTag{
		ID:        buf.Id,
		Name:      buf.Name,
		EmojiID:   buf.EmojiId,
		EmojiName: buf.EmojiName,
		Moderated: buf.Moderated,
	}
}

func ThreadMember(buf *proto.ThreadMember) *discordgo.ThreadMember {
	if buf == nil {
		return nil
	}

	return &discordgo.ThreadMember{
		ID:            buf.Id,
		UserID:        buf.UserId,
		JoinTimestamp: buf.JoinTimestamp.AsTime(),
		Flags:         int(buf.Flags),
		Member:        Member(buf.Member),
	}
}

func MessageReactions(buf *proto.MessageReactions) *discordgo.MessageReactions {
	if buf == nil {
		return nil
	}

	return &discordgo.MessageReactions{
		Count: int(buf.Count),
		Me:    buf.Me,
		Emoji: Emoji(buf.Emoji),
	}
}

func Emoji(buf *proto.Emoji) *discordgo.Emoji {
	if buf == nil {
		return nil
	}

	return &discordgo.Emoji{
		ID:            buf.Id,
		Name:          buf.Name,
		Roles:         buf.Roles,
		User:          User(buf.User),
		RequireColons: buf.RequireColons,
		Managed:       buf.Managed,
		Animated:      buf.Animated,
		Available:     buf.Animated,
	}
}

func Interaction(buf *proto.Interaction) *discordgo.Interaction {
	if buf == nil {
		return nil
	}

	return &discordgo.Interaction{
		ID:             buf.Id,
		AppID:          buf.AppId,
		Type:           discordgo.InteractionType(buf.Type),
		GuildID:        buf.GuildId,
		ChannelID:      buf.ChannelId,
		Message:        Message(buf.Message),
		AppPermissions: buf.AppPermissions,
		Member:         Member(buf.Member),
		User:           User(buf.User),
		Locale:         "",                                   // TODO(discord/bufstruct): implements Locale
		GuildLocale:    (*discordgo.Locale)(buf.GuildLocale), // TODO(discord/bufstruct): add type guard
		Token:          buf.Token,
		Version:        int(buf.Version),
	}
}

func ApplicationCommand(buf *proto.ApplicationCommand) *discordgo.ApplicationCommand {
	if buf == nil {
		return nil
	}

	return &discordgo.ApplicationCommand{
		ID:                buf.Id,
		ApplicationID:     buf.ApplicationId,
		GuildID:           buf.GuildId,
		Version:           buf.Version,
		Type:              discordgo.ApplicationCommandType(buf.Type),
		Name:              buf.Name,
		NameLocalizations: nil, // TODO(discord/bufstruct): implements NameLocalizations
		// DefaultPermission:        buf.DefaultPermission, // Deprecated
		DefaultMemberPermissions: &buf.DefaultMemberPermissions,
		// DMPermission:             buf.DmPermission, // Deprecated
		NSFW:                     buf.Nsfw,
		Description:              buf.Description,
		DescriptionLocalizations: nil, // TODO(discord/bufstruct): implements DescriptionLocalizations
		Options:                  nil, // TODO(discord/bufstruct): implements Options
	}
}
