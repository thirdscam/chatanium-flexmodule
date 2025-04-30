package bufstruct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func StructToBufUser(s *discordgo.User) *proto.User {
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

func StructToBufMember(s *discordgo.Member) *proto.Member {
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

func StructToBufMessageAttachment(s *discordgo.MessageAttachment) *proto.MessageAttachment {
	if s == nil {
		return nil
	}

	return &proto.MessageAttachment{
		Id:          s.ID,
		Url:         s.URL,
		ProxyUrl:    s.ProxyURL,
		Filename:    s.Filename,
		ContentType: s.ContentType,
		Width:       int32(s.Width),
		Height:      int32(s.Height),
		Size:        int32(s.Size),
		Ephemeral:   s.Ephemeral,
	}
}

func StructToBufChannel(s *discordgo.Channel) *proto.Channel {
	if s == nil {
		return nil
	}

	recipients := make([]*proto.User, 0, len(s.Recipients))
	for _, recipient := range s.Recipients {
		recipients = append(recipients, StructToBufUser(recipient))
	}

	availableTags := make([]*proto.ForumTag, 0, len(s.AvailableTags))
	for _, tag := range s.AvailableTags {
		availableTags = append(availableTags, StructToBufForumTag(&tag))
	}

	permissionOverwrites := make([]*proto.PermissionOverwrite, 0, len(s.PermissionOverwrites))
	for _, overwrite := range s.PermissionOverwrites {
		permissionOverwrites = append(permissionOverwrites, &proto.PermissionOverwrite{
			Id:    overwrite.ID,
			Type:  int32(overwrite.Type),
			Deny:  overwrite.Deny,
			Allow: overwrite.Allow,
		})
	}

	var threadMetadata *proto.ThreadMetadata
	if s.ThreadMetadata != nil {
		threadMetadata = &proto.ThreadMetadata{
			Archived:            s.ThreadMetadata.Archived,
			AutoArchiveDuration: int32(s.ThreadMetadata.AutoArchiveDuration),
			ArchiveTimestamp:    timestamppb.New(s.ThreadMetadata.ArchiveTimestamp),
			Locked:              s.ThreadMetadata.Locked,
			Invitable:           s.ThreadMetadata.Invitable,
		}
	}

	var member *proto.ThreadMember
	if s.Member != nil {
		member = &proto.ThreadMember{
			Id:            s.Member.ID,
			UserId:        s.Member.UserID,
			JoinTimestamp: timestamppb.New(s.Member.JoinTimestamp),
			Flags:         int32(s.Member.Flags),
			Member:        StructToBufMember(s.Member.Member),
		}
	}

	var defaultReactionEmoji *proto.ForumDefaultReaction
	if s.DefaultReactionEmoji.EmojiID != "" || s.DefaultReactionEmoji.EmojiName != "" {
		defaultReactionEmoji = &proto.ForumDefaultReaction{
			EmojiId:   s.DefaultReactionEmoji.EmojiID,
			EmojiName: s.DefaultReactionEmoji.EmojiName,
		}
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
		ThreadMetadata:                threadMetadata,
		Member:                        member,
		Flags:                         int32(s.Flags),
		AvailableTags:                 availableTags,
		AppliedTags:                   s.AppliedTags,
		DefaultReactionEmoji:          defaultReactionEmoji,
		DefaultThreadRateLimitPerUser: int32(s.DefaultThreadRateLimitPerUser),
		DefaultSortOrder:              defaultSortOrder,
		DefaultForumLayout:            int32(s.DefaultForumLayout),
	}
}

func StructToBufForumTag(s *discordgo.ForumTag) *proto.ForumTag {
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

func StructToBufThreadMember(s *discordgo.ThreadMember) *proto.ThreadMember {
	if s == nil {
		return nil
	}

	return &proto.ThreadMember{
		Id:            s.ID,
		UserId:        s.UserID,
		JoinTimestamp: timestamppb.New(s.JoinTimestamp),
		Flags:         int32(s.Flags),
		Member:        StructToBufMember(s.Member),
	}
}

func StructToBufMessageEmbed(s *discordgo.MessageEmbed) *proto.MessageEmbed {
	if s == nil {
		return nil
	}

	fields := make([]*proto.MessageEmbedField, 0, len(s.Fields))
	for _, field := range s.Fields {
		fields = append(fields, &proto.MessageEmbedField{
			Name:   field.Name,
			Value:  field.Value,
			Inline: field.Inline,
		})
	}

	var footer *proto.MessageEmbedFooter
	if s.Footer != nil {
		footer = &proto.MessageEmbedFooter{
			Text:         s.Footer.Text,
			IconUrl:      s.Footer.IconURL,
			ProxyIconUrl: s.Footer.ProxyIconURL,
		}
	}

	var image *proto.MessageEmbedImage
	if s.Image != nil {
		image = &proto.MessageEmbedImage{
			Url:      s.Image.URL,
			ProxyUrl: s.Image.ProxyURL,
			Height:   int32(s.Image.Height),
			Width:    int32(s.Image.Width),
		}
	}

	var thumbnail *proto.MessageEmbedThumbnail
	if s.Thumbnail != nil {
		thumbnail = &proto.MessageEmbedThumbnail{
			Url:      s.Thumbnail.URL,
			ProxyUrl: s.Thumbnail.ProxyURL,
			Height:   int32(s.Thumbnail.Height),
			Width:    int32(s.Thumbnail.Width),
		}
	}

	var video *proto.MessageEmbedVideo
	if s.Video != nil {
		video = &proto.MessageEmbedVideo{
			Url:    s.Video.URL,
			Height: int32(s.Video.Height),
			Width:  int32(s.Video.Width),
		}
	}

	var provider *proto.MessageEmbedProvider
	if s.Provider != nil {
		provider = &proto.MessageEmbedProvider{
			Url:  s.Provider.URL,
			Name: s.Provider.Name,
		}
	}

	var author *proto.MessageEmbedAuthor
	if s.Author != nil {
		author = &proto.MessageEmbedAuthor{
			Url:          s.Author.URL,
			Name:         s.Author.Name,
			IconUrl:      s.Author.IconURL,
			ProxyIconUrl: s.Author.ProxyIconURL,
		}
	}

	var embedType proto.EmbedType
	switch s.Type {
	case discordgo.EmbedTypeRich:
		embedType = 0
	case discordgo.EmbedTypeImage:
		embedType = 1
	case discordgo.EmbedTypeVideo:
		embedType = 2
	case discordgo.EmbedTypeGifv:
		embedType = 3
	case discordgo.EmbedTypeArticle:
		embedType = 4
	case discordgo.EmbedTypeLink:
		embedType = 5
	default:
		panic("StructToBufMessageEmbed > unknown embed type")
	}

	return &proto.MessageEmbed{
		Url:         s.URL,
		Type:        embedType,
		Title:       s.Title,
		Description: s.Description,
		Timestamp:   s.Timestamp,
		Color:       int32(s.Color),
		Footer:      footer,
		Image:       image,
		Thumbnail:   thumbnail,
		Video:       video,
		Provider:    provider,
		Author:      author,
		Fields:      fields,
	}
}

func StructToBufMessageReactions(s *discordgo.MessageReactions) *proto.MessageReactions {
	if s == nil {
		return nil
	}

	return &proto.MessageReactions{
		Count: int32(s.Count),
		Me:    s.Me,
		Emoji: StructToBufEmoji(s.Emoji),
	}
}

func StructToBufEmoji(s *discordgo.Emoji) *proto.Emoji {
	if s == nil {
		return nil
	}

	return &proto.Emoji{
		Id:            s.ID,
		Name:          s.Name,
		Roles:         s.Roles,
		User:          StructToBufUser(s.User),
		RequireColons: s.RequireColons,
		Managed:       s.Managed,
		Animated:      s.Animated,
		Available:     s.Available,
	}
}

func StructToBufMessage(s *discordgo.Message) *proto.Message {
	if s == nil {
		return nil
	}

	attachments := make([]*proto.MessageAttachment, 0, len(s.Attachments))
	for _, att := range s.Attachments {
		attachments = append(attachments, StructToBufMessageAttachment(att))
	}

	mentionChannels := make([]*proto.Channel, 0, len(s.MentionChannels))
	for _, channel := range s.MentionChannels {
		mentionChannels = append(mentionChannels, StructToBufChannel(channel))
	}

	embeds := make([]*proto.MessageEmbed, 0, len(s.Embeds))
	for _, embed := range s.Embeds {
		embeds = append(embeds, StructToBufMessageEmbed(embed))
	}

	var activity *proto.MessageActivity
	if s.Activity != nil {
		activity = &proto.MessageActivity{
			Type:    proto.MessageActivityType(s.Activity.Type),
			PartyId: s.Activity.PartyID,
		}
	}

	var application *proto.MessageApplication
	if s.Application != nil {
		application = &proto.MessageApplication{
			Id:          s.Application.ID,
			CoverImage:  s.Application.CoverImage,
			Description: s.Application.Description,
			Icon:        s.Application.Icon,
			Name:        s.Application.Name,
		}
	}

	return &proto.Message{
		Id:              s.ID,
		ChannelId:       s.ChannelID,
		GuildId:         s.GuildID,
		Content:         s.Content,
		Timestamp:       timestamppb.New(s.Timestamp),
		EditedTimestamp: util.AsTimePtrToPbTimestamp(s.EditedTimestamp),
		MentionRoles:    s.MentionRoles,
		Tts:             s.TTS,
		MentionEveryone: s.MentionEveryone,
		Author:          StructToBufUser(s.Author),
		Attachments:     attachments,
		Embeds:          embeds,
		Pinned:          s.Pinned,
		Type:            proto.MessageType(s.Type),
		WebhookId:       s.WebhookID,
		Member:          StructToBufMember(s.Member),
		MentionChannels: mentionChannels,
		Activity:        activity,
		Application:     application,
	}
}

// func StructToBufInteraction(s *discordgo.Interaction) *proto.Interaction {
// 	if s == nil {
// 		return nil
// 	}

// 	return &proto.Interaction{
// 		Id:        s.ID,
// 		AppId:     s.AppID,
// 		Type:      proto.InteractionType(s.Type),
// 		GuildId:   s.GuildID,
// 		ChannelId: s.ChannelID,
// 		Message:   StructToBufMessage(s.Message),
// 		User:      StructToBufUser(s.User),
// 	}
// }

func StructToBufApplicationCmd(s *discordgo.ApplicationCommand) *proto.ApplicationCommand {
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
		panic("StructToBufApplicationCmd > unknown command type")
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
		DefaultMemberPermissions: *s.DefaultMemberPermissions,
		// DmPermission:             s.DMPermission, // Deprecated
		Nsfw:                     s.NSFW,
		Description:              s.Description,
		DescriptionLocalizations: nil, // TODO(discord/bufstruct): implements DescriptionLocalizations
		Options:                  nil, // TODO(discord/bufstruct): implements Options
	}
}

func StructToBufApplicationCmds(s []*discordgo.ApplicationCommand) []*proto.ApplicationCommand {
	if s == nil {
		return nil
	}

	bufs := make([]*proto.ApplicationCommand, len(s))
	for i, v := range s {
		bufs[i] = StructToBufApplicationCmd(v)
	}
	return bufs
}
