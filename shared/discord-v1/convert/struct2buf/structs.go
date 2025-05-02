package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func User(s *discordgo.User) *proto.User {
	if s == nil {
		return nil
	}

	return &proto.User{
		Id:               s.ID,
		Username:         s.Username,
		Discriminator:    s.Discriminator,
		Avatar:           s.Avatar,
		Bot:              s.Bot,
		PublicFlagsValue: int32(s.PublicFlags),
		Locale:           s.Locale,
		Verified:         s.Verified,
		Email:            s.Email,
		Flags:            int32(s.Flags),
	}
}

func Member(s *discordgo.Member) *proto.Member {
	if s == nil {
		return nil
	}

	return &proto.Member{
		GuildId:                    s.GuildID,
		JoinedAt:                   timestamppb.New(s.JoinedAt),
		Nick:                       s.Nick,
		Deaf:                       s.Deaf,
		Mute:                       s.Mute,
		Avatar:                     s.Avatar,
		Roles:                      s.Roles,
		PremiumSince:               util.AsTimePtrToPbTimestamp(s.PremiumSince),
		Flags:                      int32(s.Flags),
		Pending:                    s.Pending,
		Permissions:                s.Permissions,
		CommunicationDisabledUntil: util.AsTimePtrToPbTimestamp(s.CommunicationDisabledUntil),
	}
}

func Channel(s *discordgo.Channel) *proto.Channel {
	if s == nil {
		return nil
	}

	recipients := make([]*proto.User, 0, len(s.Recipients))
	for _, recipient := range s.Recipients {
		recipients = append(recipients, User(recipient))
	}

	availableTags := make([]*proto.ForumTag, 0, len(s.AvailableTags))
	for _, tag := range s.AvailableTags {
		availableTags = append(availableTags, ForumTag(&tag))
	}

	permissionOverwrites := make([]*proto.PermissionOverwrite, 0, len(s.PermissionOverwrites))
	for _, overwrite := range s.PermissionOverwrites {
		permissionOverwrites = append(permissionOverwrites, PermissionOverwrite(overwrite))
	}

	var defaultSortOrder int32
	if s.DefaultSortOrder != nil {
		defaultSortOrder = int32(*s.DefaultSortOrder)
	}

	return &proto.Channel{
		Id:                            s.ID,
		GuildId:                       s.GuildID,
		Name:                          s.Name,
		Topic:                         s.Topic,
		Type:                          int32(s.Type),
		LastMessageId:                 s.LastMessageID,
		LastPinTimestamp:              util.AsTimePtrToPbTimestamp(s.LastPinTimestamp),
		MessageCount:                  int32(s.MessageCount),
		MemberCount:                   int32(s.MemberCount),
		Nsfw:                          s.NSFW,
		Icon:                          s.Icon,
		Position:                      int32(s.Position),
		Bitrate:                       int32(s.Bitrate),
		Recipients:                    recipients,
		PermissionOverwrites:          permissionOverwrites,
		UserLimit:                     int32(s.UserLimit),
		ParentId:                      s.ParentID,
		RateLimitPerUser:              int32(s.RateLimitPerUser),
		OwnerId:                       s.OwnerID,
		ApplicationId:                 s.ApplicationID,
		ThreadMetadata:                ThreadMetadata(s.ThreadMetadata),
		Member:                        ThreadMember(s.Member),
		Flags:                         int32(s.Flags),
		AvailableTags:                 availableTags,
		AppliedTags:                   s.AppliedTags,
		DefaultReactionEmoji:          ForumDefaultReactionNonPtr(s.DefaultReactionEmoji),
		DefaultThreadRateLimitPerUser: int32(s.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              defaultSortOrder,
		DefaultForumLayout:            int32(s.DefaultForumLayout),
	}
}

func ForumDefaultReaction(s *discordgo.ForumDefaultReaction) *proto.ForumDefaultReaction {
	if s == nil {
		return nil
	}

	return &proto.ForumDefaultReaction{
		EmojiId:   s.EmojiID,
		EmojiName: s.EmojiName,
	}
}

func PermissionOverwrite(s *discordgo.PermissionOverwrite) *proto.PermissionOverwrite {
	if s == nil {
		return nil
	}

	return &proto.PermissionOverwrite{
		Id:    s.ID,
		Type:  int32(s.Type),
		Deny:  s.Deny,
		Allow: s.Allow,
	}
}

func ThreadMetadata(s *discordgo.ThreadMetadata) *proto.ThreadMetadata {
	if s == nil {
		return nil
	}

	return &proto.ThreadMetadata{
		Archived:            s.Archived,
		AutoArchiveDuration: int32(s.AutoArchiveDuration),
		ArchiveTimestamp:    timestamppb.New(s.ArchiveTimestamp),
		Locked:              s.Locked,
		Invitable:           s.Invitable,
	}
}

func ForumTag(s *discordgo.ForumTag) *proto.ForumTag {
	if s == nil {
		return nil
	}

	return &proto.ForumTag{
		Id:        s.ID,
		Name:      s.Name,
		EmojiId:   s.EmojiID,
		EmojiName: s.EmojiName,
		Moderated: s.Moderated,
	}
}

func Emoji(s *discordgo.Emoji) *proto.Emoji {
	if s == nil {
		return nil
	}

	return &proto.Emoji{
		Id:            s.ID,
		Name:          s.Name,
		Roles:         s.Roles,
		User:          User(s.User),
		RequireColons: s.RequireColons,
		Managed:       s.Managed,
		Animated:      s.Animated,
		Available:     s.Available,
	}
}

func ThreadMember(s *discordgo.ThreadMember) *proto.ThreadMember {
	if s == nil {
		return nil
	}

	return &proto.ThreadMember{
		Id:            s.ID,
		UserId:        s.UserID,
		JoinTimestamp: timestamppb.New(s.JoinTimestamp),
		Flags:         int32(s.Flags),
		Member:        Member(s.Member),
	}
}

func Role(s *discordgo.Role) *proto.Role {
	if s == nil {
		return nil
	}

	return &proto.Role{
		Id:           s.ID,
		Name:         s.Name,
		Managed:      s.Managed,
		Mentionable:  s.Mentionable,
		Hoist:        s.Hoist,
		Color:        int32(s.Color),
		Position:     int32(s.Position),
		Permissions:  s.Permissions,
		Icon:         s.Icon,
		UnicodeEmoji: s.UnicodeEmoji,
		Flags:        int32(s.Flags),
	}
}
