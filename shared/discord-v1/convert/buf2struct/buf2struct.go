package buf2struct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/util"
)

func BufUserToStruct(buf *proto.User) *discordgo.User {
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

func BufMemberToStruct(buf *proto.Member) *discordgo.Member {
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

func BufMessageAttachmentToStruct(buf *proto.MessageAttachment) *discordgo.MessageAttachment {
	if buf == nil {
		return nil
	}

	return &discordgo.MessageAttachment{
		ID:          buf.Id,
		URL:         buf.Url,
		ProxyURL:    buf.ProxyUrl,
		Filename:    buf.Filename,
		ContentType: buf.ContentType,
		Width:       int(buf.Width),
		Height:      int(buf.Height),
		Size:        int(buf.Size),
		Ephemeral:   buf.Ephemeral,
	}
}

func BufChannelToStruct(buf *proto.Channel) *discordgo.Channel {
	if buf == nil {
		return nil
	}

	recipients := make([]*discordgo.User, 0)
	for _, recipient := range buf.Recipients {
		recipients = append(recipients, BufUserToStruct(recipient))
	}

	availableTags := make([]discordgo.ForumTag, 0)
	for _, tag := range buf.AvailableTags {
		availableTags = append(availableTags, *BufForumTagToStruct(tag))
	}

	permissionOverwrites := make([]*discordgo.PermissionOverwrite, 0)
	for _, overwrite := range buf.PermissionOverwrites {
		permissionOverwrites = append(permissionOverwrites, &discordgo.PermissionOverwrite{
			ID:    overwrite.Id,
			Type:  discordgo.PermissionOverwriteType(overwrite.Type),
			Deny:  overwrite.Deny,
			Allow: overwrite.Allow,
		})
	}

	sortOrder := discordgo.ForumSortOrderType(buf.DefaultSortOrder)
	defaultSortOrder := &sortOrder
	return &discordgo.Channel{
		ID:                   buf.Id,
		GuildID:              buf.GuildId,
		Name:                 buf.Name,
		Topic:                buf.Topic,
		Type:                 discordgo.ChannelType(buf.Type),
		LastMessageID:        buf.LastMessageId,
		LastPinTimestamp:     util.PbTimestamp2AsTimePtr(buf.LastPinTimestamp),
		MessageCount:         int(buf.MessageCount),
		MemberCount:          int(buf.MemberCount),
		NSFW:                 buf.Nsfw,
		Icon:                 buf.Icon,
		Position:             int(buf.Position),
		Bitrate:              int(buf.Bitrate),
		Recipients:           recipients,
		Messages:             nil, // TODO: Mapping state-enabled fields (#1)
		PermissionOverwrites: permissionOverwrites,
		UserLimit:            int(buf.UserLimit),
		ParentID:             buf.ParentId,
		RateLimitPerUser:     int(buf.RateLimitPerUser),
		OwnerID:              buf.OwnerId,
		ApplicationID:        buf.ApplicationId,
		ThreadMetadata: &discordgo.ThreadMetadata{
			Archived:            buf.ThreadMetadata.Archived,
			AutoArchiveDuration: int(buf.ThreadMetadata.AutoArchiveDuration),
			ArchiveTimestamp:    buf.ThreadMetadata.ArchiveTimestamp.AsTime(),
			Locked:              buf.ThreadMetadata.Locked,
			Invitable:           buf.ThreadMetadata.Invitable,
		},
		Member: &discordgo.ThreadMember{
			ID:            buf.Member.Id,
			UserID:        buf.Member.UserId,
			JoinTimestamp: buf.Member.JoinTimestamp.AsTime(),
			Flags:         int(buf.Member.Flags),
			Member:        BufMemberToStruct(buf.Member.Member),
		},
		Members:       nil, // TODO: Mapping state-enabled fields (#1)
		Flags:         discordgo.ChannelFlags(buf.Flags),
		AvailableTags: availableTags,
		AppliedTags:   buf.AppliedTags,
		DefaultReactionEmoji: discordgo.ForumDefaultReaction{
			EmojiID:   buf.DefaultReactionEmoji.EmojiId,
			EmojiName: buf.DefaultReactionEmoji.EmojiName,
		},
		DefaultThreadRateLimitPerUser: int(buf.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              defaultSortOrder,
		DefaultForumLayout:            discordgo.ForumLayout(buf.DefaultForumLayout),
	}
}

func BufForumTagToStruct(buf *proto.ForumTag) *discordgo.ForumTag {
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

func BufThreadMemberToStruct(buf *proto.ThreadMember) *discordgo.ThreadMember {
	if buf == nil {
		return nil
	}

	return &discordgo.ThreadMember{
		ID:            buf.Id,
		UserID:        buf.UserId,
		JoinTimestamp: buf.JoinTimestamp.AsTime(),
		Flags:         int(buf.Flags),
		Member:        BufMemberToStruct(buf.Member),
	}
}

func BufMessageEmbedToStruct(buf *proto.MessageEmbed) *discordgo.MessageEmbed {
	if buf == nil {
		return nil
	}

	messageEmbedField := make([]*discordgo.MessageEmbedField, 0)
	for _, field := range buf.Fields {
		messageEmbedField = append(messageEmbedField, &discordgo.MessageEmbedField{
			Name:   field.Name,
			Value:  field.Value,
			Inline: field.Inline,
		})
	}

	return &discordgo.MessageEmbed{
		URL:         buf.Url,
		Type:        discordgo.EmbedType(buf.Type),
		Title:       buf.Title,
		Description: buf.Description,
		Timestamp:   buf.Timestamp,
		Color:       int(buf.Color),
		Footer: &discordgo.MessageEmbedFooter{
			Text:         buf.Footer.Text,
			IconURL:      buf.Footer.IconUrl,
			ProxyIconURL: buf.Footer.ProxyIconUrl,
		},
		Image: &discordgo.MessageEmbedImage{
			URL:      buf.Image.Url,
			ProxyURL: buf.Image.ProxyUrl,
			Height:   int(buf.Image.Height),
			Width:    int(buf.Image.Width),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:      buf.Thumbnail.Url,
			ProxyURL: buf.Thumbnail.ProxyUrl,
			Height:   int(buf.Thumbnail.Height),
			Width:    int(buf.Thumbnail.Width),
		},
		Video: &discordgo.MessageEmbedVideo{
			URL:    buf.Video.Url,
			Height: int(buf.Video.Height),
			Width:  int(buf.Video.Width),
		},
		Provider: &discordgo.MessageEmbedProvider{
			URL:  buf.Provider.Url,
			Name: buf.Provider.Name,
		},
		Author: &discordgo.MessageEmbedAuthor{
			URL:          buf.Author.Url,
			Name:         buf.Author.Name,
			IconURL:      buf.Author.IconUrl,
			ProxyIconURL: buf.Author.ProxyIconUrl,
		},
		Fields: messageEmbedField,
	}
}

func BufMessageReactionsToStruct(buf *proto.MessageReactions) *discordgo.MessageReactions {
	if buf == nil {
		return nil
	}

	return &discordgo.MessageReactions{
		Count: int(buf.Count),
		Me:    buf.Me,
		Emoji: BufEmojiToStruct(buf.Emoji),
	}
}

func BufEmojiToStruct(buf *proto.Emoji) *discordgo.Emoji {
	if buf == nil {
		return nil
	}

	return &discordgo.Emoji{
		ID:            buf.Id,
		Name:          buf.Name,
		Roles:         buf.Roles,
		User:          BufUserToStruct(buf.User),
		RequireColons: buf.RequireColons,
		Managed:       buf.Managed,
		Animated:      buf.Animated,
		Available:     buf.Animated,
	}
}

func BufMessageToStruct(buf *proto.Message) *discordgo.Message {
	if buf == nil {
		return nil
	}

	attachments := make([]*discordgo.MessageAttachment, 0)
	for _, attachment := range buf.Attachments {
		attachments = append(attachments, BufMessageAttachmentToStruct(attachment))
	}

	mentionChannels := make([]*discordgo.Channel, 0)
	for _, channel := range buf.MentionChannels {
		mentionChannels = append(mentionChannels, BufChannelToStruct(channel))
	}

	embeds := make([]*discordgo.MessageEmbed, 0)
	for _, embed := range buf.Embeds {
		embeds = append(embeds, BufMessageEmbedToStruct(embed))
	}

	var activity *discordgo.MessageActivity
	if buf.Activity != nil {
		activity = &discordgo.MessageActivity{
			Type:    discordgo.MessageActivityType(buf.Activity.Type),
			PartyID: buf.Activity.PartyId,
		}
	}

	var application *discordgo.MessageApplication
	if buf.Application != nil {
		application = &discordgo.MessageApplication{
			ID:          buf.Application.Id,
			CoverImage:  buf.Application.CoverImage,
			Description: buf.Application.Description,
			Icon:        buf.Application.Icon,
			Name:        buf.Application.Name,
		}
	}

	return &discordgo.Message{
		ID:              buf.Id,
		ChannelID:       buf.ChannelId,
		GuildID:         buf.GuildId,
		Content:         buf.Content,
		Timestamp:       buf.Timestamp.AsTime(),
		EditedTimestamp: util.PbTimestamp2AsTimePtr(buf.EditedTimestamp),
		MentionRoles:    buf.MentionRoles,
		TTS:             buf.Tts,
		MentionEveryone: buf.MentionEveryone,
		Author:          BufUserToStruct(buf.Author),
		Attachments:     attachments,
		Components:      nil, // TODO(discord/buf2struct): Support components
		Embeds:          embeds,
		Reactions:       nil,
		Pinned:          buf.Pinned,
		Type:            discordgo.MessageType(buf.Type),
		WebhookID:       buf.WebhookId,
		Member:          BufMemberToStruct(buf.Member),
		MentionChannels: mentionChannels,
		Activity:        activity,
		Application:     application,
	}
}

func BufInteractionToStruct(buf *proto.Interaction) *discordgo.Interaction {
	if buf == nil {
		return nil
	}

	return &discordgo.Interaction{
		ID:             buf.Id,
		AppID:          buf.AppId,
		Type:           discordgo.InteractionType(buf.Type),
		GuildID:        buf.GuildId,
		ChannelID:      buf.ChannelId,
		Message:        BufMessageToStruct(buf.Message),
		AppPermissions: buf.AppPermissions,
		Member:         BufMemberToStruct(buf.Member),
		User:           BufUserToStruct(buf.User),
		Locale:         "",                                   // TODO(discord/bufstruct): implements Locale
		GuildLocale:    (*discordgo.Locale)(buf.GuildLocale), // TODO(discord/bufstruct): add type guard
		Token:          buf.Token,
		Version:        int(buf.Version),
	}
}

func BufApplicationCmdToStruct(buf *proto.ApplicationCommand) *discordgo.ApplicationCommand {
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
