package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"github.com/thirdscam/chatanium-flexmodule/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Message(s *discordgo.Message) *proto.Message {
	if s == nil {
		return nil
	}

	attachments := make([]*proto.MessageAttachment, 0, len(s.Attachments))
	for _, att := range s.Attachments {
		attachments = append(attachments, MessageAttachment(att))
	}

	mentionChannels := make([]*proto.Channel, 0, len(s.MentionChannels))
	for _, channel := range s.MentionChannels {
		mentionChannels = append(mentionChannels, Channel(channel))
	}

	embeds := make([]*proto.MessageEmbed, 0, len(s.Embeds))
	for _, embed := range s.Embeds {
		embeds = append(embeds, MessageEmbed(embed))
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
		Author:          User(s.Author),
		Attachments:     attachments,
		Embeds:          embeds,
		Pinned:          s.Pinned,
		Type:            proto.MessageType(s.Type),
		WebhookId:       s.WebhookID,
		Member:          Member(s.Member),
		MentionChannels: mentionChannels,
		Activity:        activity,
		Application:     application,
	}
}

func MessageAttachment(s *discordgo.MessageAttachment) *proto.MessageAttachment {
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

func MessageEmbed(s *discordgo.MessageEmbed) *proto.MessageEmbed {
	if s == nil {
		return nil
	}

	fields := make([]*proto.MessageEmbedField, 0, len(s.Fields))
	for _, field := range s.Fields {
		fields = append(fields, MessageEmbedField(field))
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
		panic("StructToBuf::MessageEmbed > unknown embed type")
	}

	return &proto.MessageEmbed{
		Url:         s.URL,
		Type:        embedType,
		Title:       s.Title,
		Description: s.Description,
		Timestamp:   s.Timestamp,
		Color:       int32(s.Color),
		Footer:      MessageEmbedFooter(s.Footer),
		Image:       MessageEmbedImage(s.Image),
		Thumbnail:   MessageEmbedThumbnail(s.Thumbnail),
		Video:       MessageEmbedVideo(s.Video),
		Provider:    MessageEmbedProvider(s.Provider),
		Author:      MessageEmbedAuthor(s.Author),
		Fields:      fields,
	}
}

func MessageEmbedField(s *discordgo.MessageEmbedField) *proto.MessageEmbedField {
	if s == nil {
		return nil
	}

	return &proto.MessageEmbedField{
		Name:   s.Name,
		Value:  s.Value,
		Inline: s.Inline,
	}
}

func MessageEmbedFooter(s *discordgo.MessageEmbedFooter) *proto.MessageEmbedFooter {
	if s == nil {
		return nil
	}

	return &proto.MessageEmbedFooter{
		Text:         s.Text,
		IconUrl:      s.IconURL,
		ProxyIconUrl: s.ProxyIconURL,
	}
}

func MessageEmbedImage(s *discordgo.MessageEmbedImage) *proto.MessageEmbedImage {
	if s == nil {
		return nil
	}

	return &proto.MessageEmbedImage{
		Url:      s.URL,
		ProxyUrl: s.ProxyURL,
		Height:   int32(s.Height),
		Width:    int32(s.Width),
	}
}

func MessageEmbedThumbnail(s *discordgo.MessageEmbedThumbnail) *proto.MessageEmbedThumbnail {
	if s == nil {
		return nil
	}

	return &proto.MessageEmbedThumbnail{
		Url:      s.URL,
		ProxyUrl: s.ProxyURL,
		Height:   int32(s.Height),
	}
}

func MessageEmbedVideo(s *discordgo.MessageEmbedVideo) *proto.MessageEmbedVideo {
	if s == nil {
		return nil
	}

	return &proto.MessageEmbedVideo{
		Url:    s.URL,
		Height: int32(s.Height),
		Width:  int32(s.Width),
	}
}

func MessageEmbedProvider(s *discordgo.MessageEmbedProvider) *proto.MessageEmbedProvider {
	if s == nil {
		return nil
	}

	return &proto.MessageEmbedProvider{
		Name: s.Name,
		Url:  s.URL,
	}
}

func MessageEmbedAuthor(s *discordgo.MessageEmbedAuthor) *proto.MessageEmbedAuthor {
	if s == nil {
		return nil
	}

	return &proto.MessageEmbedAuthor{
		Name:         s.Name,
		Url:          s.URL,
		IconUrl:      s.IconURL,
		ProxyIconUrl: s.ProxyIconURL,
	}
}

func MessageReactions(s *discordgo.MessageReactions) *proto.MessageReactions {
	if s == nil {
		return nil
	}

	return &proto.MessageReactions{
		Count: int32(s.Count),
		Me:    s.Me,
		Emoji: Emoji(s.Emoji),
	}
}
