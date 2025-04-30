package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
	plugin "github.com/hashicorp/go-plugin"
	proto_common "github.com/thirdscam/chatanium-flexmodule/proto"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Hook

	broker *plugin.GRPCBroker
}

func (m *GRPCServer) OnInit(ctx context.Context, req *proto_common.Empty) (*proto_common.Empty, error) {
	m.Impl.OnInit()

	return &proto_common.Empty{}, nil
}

func (m *GRPCServer) OnCreateChatMessage(ctx context.Context, req *proto.Message) (*proto_common.Empty, error) {
	attachments := make([]discordgo.MessageAttachment, 0)
	for _, attachment := range req.Attachments {
		attachments = append(attachments, discordgo.MessageAttachment{
			URL: attachment.Url,
			ProxyURL: attachment.ProxyUrl,
		})
	}

	mentionChannels := make([]discordgo.ChannelMention, 0)
	for _, mentionChannel := range req.MentionChannels {
		mentionChannels = append(mentionChannels, discordgo.ChannelMention{
			ID: mentionChannel.Id,
			GuildID: mentionChannel.GuildId,
			Type: mentionChannel.Type,
		})
	}

	m.Impl.OnCreateChatMessage(discordgo.Message{
		ID:        req.Id,
		ChannelID: req.ChannelId,
		GuildID:   req.GuildId,
		Content:   req.Content,
		Timestamp: req.Timestamp.AsTime(),
		EditedTimestamp: req.EditedTimestamp.AsTime(),
		MentionRoles: req.MentionRoles,
		TTS: req.Tts,
		MentionEveryone: req.MentionEveryone,
		Author: &discordgo.User{
			ID: req.Author.Id,
			Username: req.Author.Username,
			Discriminator: req.Author.Discriminator,
			Avatar: req.Author.Avatar,
			Bot: req.Author.Bot,
			PublicFlags: req.Author.PublicFlags,
			BotPublicFlags: req.Author.BotPublicFlags,
			Locale: req.Author.Locale,
			Verified: req.Author.Verified,
			Email: req.Author.Email,
			Flags: req.Author.Flags,
		},
		Attachments: attachments,
		Components: req.Components,
		Embeds: req.Embeds,
		Reactions: req.Reactions,
		Pinned: req.Pinned,
		Type: req.Type,
		WebhookID: req.WebhookId,
		Member: &discordgo.Member{
			GuildID: req.GuildId,
			JoinedAt: req.Member.JoinedAt.AsTime(),
			Nick: req.Member.Nick,
			Deaf: req.Member.Deaf,
			Mute: req.Member.Mute,
			Avatar: req.Member.Avatar,
			Roles: req.Member.Roles,
			PremiumSince: req.Member.PremiumSince.AsTime(),
			Flags: req.Member.Flags,
			Pending: req.Member.Pending,
			Permissions: req.Member.Permissions,
			CommunicationDisabledUntil: req.Member.CommunicationDisabledUntil.AsTime(),
		},
		MentionChannels: mentionChannels,
		Activity: &discordgo.Activity{
			Name: req.Activity.Name,
			Type: req.Activity.Type,
			CreatedAt: req.Activity.CreatedAt.AsTime(),
		},
		Channel: &discordgo.Channel{
			ID: req.Activity.Channel.Id,
		},
	})

	return &proto_common.Empty{}, nil
}

func (m *GRPCServer) OnCreateInteraction(ctx context.Context, req *proto.Interaction) (*proto_common.Empty, error) {
	m.Impl.OnCreateInteraction(discordgo.Interaction{
		ID:    req.Id,
		AppID: req.AppId,
		Type:  discordgo.InteractionType(req.Type),
		// Data: 
		GuildID: req.GuildId,
		ChannelID: req.ChannelId,
		Message: &discordgo.Message{
			ID: req.Message.Id,
			ChannelID: req.Message.ChannelId,
			GuildID: req.Message.GuildId,
			Content: req.Message.Content,
			Timestamp: req.Message.Timestamp.AsTime(),
		},
		Messsage: &discordgo.Message{
			ID: req.Message.Id,
			ChannelID: req.Message.ChannelId,
			GuildID: req.Message.GuildId,
			Content: req.Message.Content,
			Timestamp: req.Message.Timestamp.AsTime(),
		}
	})

	return &proto_common.Empty{}, nil
}
