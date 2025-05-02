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

func Channel(buf *proto.Channel) *discordgo.Channel {
	if buf == nil {
		return nil
	}

	recipients := make([]*discordgo.User, 0)
	for _, recipient := range buf.Recipients {
		recipients = append(recipients, User(recipient))
	}

	availableTags := make([]discordgo.ForumTag, 0)
	for _, tag := range buf.AvailableTags {
		availableTags = append(availableTags, *ForumTag(tag))
	}

	permissionOverwrites := make([]*discordgo.PermissionOverwrite, 0)
	for _, overwrite := range buf.PermissionOverwrites {
		permissionOverwrites = append(permissionOverwrites, PermissionOverwrite(overwrite))
	}

	sortOrder := discordgo.ForumSortOrderType(buf.DefaultSortOrder)
	defaultSortOrder := &sortOrder
	return &discordgo.Channel{
		ID:                            buf.Id,
		GuildID:                       buf.GuildId,
		Name:                          buf.Name,
		Topic:                         buf.Topic,
		Type:                          discordgo.ChannelType(buf.Type),
		LastMessageID:                 buf.LastMessageId,
		LastPinTimestamp:              util.PbTimestamp2AsTimePtr(buf.LastPinTimestamp),
		MessageCount:                  int(buf.MessageCount),
		MemberCount:                   int(buf.MemberCount),
		NSFW:                          buf.Nsfw,
		Icon:                          buf.Icon,
		Position:                      int(buf.Position),
		Bitrate:                       int(buf.Bitrate),
		Recipients:                    recipients,
		Messages:                      nil, // TODO: Mapping state-enabled fields (#1)
		PermissionOverwrites:          permissionOverwrites,
		UserLimit:                     int(buf.UserLimit),
		ParentID:                      buf.ParentId,
		RateLimitPerUser:              int(buf.RateLimitPerUser),
		OwnerID:                       buf.OwnerId,
		ApplicationID:                 buf.ApplicationId,
		ThreadMetadata:                ThreadMetadata(buf.ThreadMetadata),
		Member:                        ThreadMember(buf.Member),
		Members:                       nil, // TODO: Mapping state-enabled fields (#1)
		Flags:                         discordgo.ChannelFlags(buf.Flags),
		AvailableTags:                 availableTags,
		AppliedTags:                   buf.AppliedTags,
		DefaultReactionEmoji:          *ForumDefaultReaction(buf.DefaultReactionEmoji),
		DefaultThreadRateLimitPerUser: int(buf.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              defaultSortOrder,
		DefaultForumLayout:            discordgo.ForumLayout(buf.DefaultForumLayout),
	}
}

func ForumDefaultReaction(buf *proto.ForumDefaultReaction) *discordgo.ForumDefaultReaction {
	if buf == nil {
		return nil
	}

	return &discordgo.ForumDefaultReaction{
		EmojiID:   buf.EmojiId,
		EmojiName: buf.EmojiName,
	}
}

func PermissionOverwrite(buf *proto.PermissionOverwrite) *discordgo.PermissionOverwrite {
	if buf == nil {
		return nil
	}

	return &discordgo.PermissionOverwrite{
		ID:    buf.Id,
		Type:  discordgo.PermissionOverwriteType(buf.Type),
		Deny:  buf.Deny,
		Allow: buf.Allow,
	}
}

func ThreadMetadata(buf *proto.ThreadMetadata) *discordgo.ThreadMetadata {
	if buf == nil {
		return nil
	}

	return &discordgo.ThreadMetadata{
		Archived:            buf.Archived,
		AutoArchiveDuration: int(buf.AutoArchiveDuration),
		ArchiveTimestamp:    buf.ArchiveTimestamp.AsTime(),
		Locked:              buf.Locked,
		Invitable:           buf.Invitable,
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

func Role(buf *proto.Role) *discordgo.Role {
	if buf == nil {
		return nil
	}

	return &discordgo.Role{
		ID:           buf.Id,
		Name:         buf.Name,
		Managed:      buf.Managed,
		Mentionable:  buf.Mentionable,
		Hoist:        buf.Hoist,
		Color:        int(buf.Color),
		Position:     int(buf.Position),
		Permissions:  buf.Permissions,
		Icon:         buf.Icon,
		UnicodeEmoji: buf.UnicodeEmoji,
		Flags:        discordgo.RoleFlags(buf.Flags),
	}
}
