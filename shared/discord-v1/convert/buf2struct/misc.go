package buf2struct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// GatewayBotResponse converts proto.GatewayBotResponse to discordgo.GatewayBotResponse
func GatewayBotResponse(buf *proto.GatewayBotResponse) *discordgo.GatewayBotResponse {
	if buf == nil {
		return nil
	}

	return &discordgo.GatewayBotResponse{
		URL:               buf.Url,
		Shards:            int(buf.Shards),
		SessionStartLimit: *SessionInformation(buf.SessionStartLimit),
	}
}

// SessionInformation converts proto.SessionInformation to discordgo.SessionInformation
func SessionInformation(buf *proto.SessionInformation) *discordgo.SessionInformation {
	if buf == nil {
		return &discordgo.SessionInformation{}
	}

	return &discordgo.SessionInformation{
		Total:          int(buf.Total),
		Remaining:      int(buf.Remaining),
		ResetAfter:     int(buf.ResetAfter),
		MaxConcurrency: int(buf.MaxConcurrency),
	}
}

// APIErrorMessage converts proto.APIErrorMessage to discordgo.APIErrorMessage
func APIErrorMessage(buf *proto.APIErrorMessage) *discordgo.APIErrorMessage {
	if buf == nil {
		return nil
	}

	return &discordgo.APIErrorMessage{
		Code:    int(buf.Code),
		Message: buf.Message,
	}
}

// TooManyRequests converts proto.TooManyRequests to discordgo.TooManyRequests
func TooManyRequests(buf *proto.TooManyRequests) *discordgo.TooManyRequests {
	if buf == nil {
		return nil
	}

	return &discordgo.TooManyRequests{
		Bucket:     buf.Bucket,
		Message:    buf.Message,
		RetryAfter: buf.RetryAfter.AsDuration(),
	}
}

// ReadState converts proto.ReadState to discordgo.ReadState
func ReadState(buf *proto.ReadState) *discordgo.ReadState {
	if buf == nil {
		return nil
	}

	return &discordgo.ReadState{
		MentionCount:  int(buf.MentionCount),
		LastMessageID: buf.LastMessageId,
		ID:            buf.Id,
	}
}

// GuildRole converts proto.GuildRole to discordgo.GuildRole
func GuildRole(buf *proto.GuildRole) *discordgo.GuildRole {
	if buf == nil {
		return nil
	}

	return &discordgo.GuildRole{
		Role:    Role(buf.Role),
		GuildID: buf.GuildId,
	}
}

// GuildBan converts proto.GuildBan to discordgo.GuildBan
func GuildBan(buf *proto.GuildBan) *discordgo.GuildBan {
	if buf == nil {
		return nil
	}

	return &discordgo.GuildBan{
		Reason: buf.Reason,
		User:   User(buf.User),
	}
}

// File converts proto.File to discordgo.File
func File(buf *proto.File) *discordgo.File {
	if buf == nil {
		return nil
	}

	return &discordgo.File{
		Name:        buf.Name,
		ContentType: buf.ContentType,
	}
}

// MessageSnapshot converts proto.MessageSnapshot to discordgo.MessageSnapshot
func MessageSnapshot(buf *proto.MessageSnapshot) *discordgo.MessageSnapshot {
	if buf == nil {
		return nil
	}

	return &discordgo.MessageSnapshot{
		Message: Message(buf.Message),
	}
}

// ThreadStart converts proto.ThreadStart to discordgo.ThreadStart
func ThreadStart(buf *proto.ThreadStart) *discordgo.ThreadStart {
	if buf == nil {
		return nil
	}

	return &discordgo.ThreadStart{
		Name:                buf.Name,
		AutoArchiveDuration: int(buf.AutoArchiveDuration),
		Type:                discordgo.ChannelType(buf.Type),
		Invitable:           buf.Invitable,
		RateLimitPerUser:    int(buf.RateLimitPerUser),
		AppliedTags:         buf.AppliedTags,
	}
}

// ThreadsList converts proto.ThreadsList to discordgo.ThreadsList
func ThreadsList(buf *proto.ThreadsList) *discordgo.ThreadsList {
	if buf == nil {
		return nil
	}

	threads := make([]*discordgo.Channel, len(buf.Threads))
	for i, t := range buf.Threads {
		threads[i] = Channel(t)
	}

	members := make([]*discordgo.ThreadMember, len(buf.Members))
	for i, m := range buf.Members {
		members[i] = ThreadMember(m)
	}

	return &discordgo.ThreadsList{
		Threads: threads,
		Members: members,
		HasMore: buf.HasMore,
	}
}

// AddedThreadMember converts proto.AddedThreadMember to discordgo.AddedThreadMember
func AddedThreadMember(buf *proto.AddedThreadMember) *discordgo.AddedThreadMember {
	if buf == nil {
		return nil
	}

	return &discordgo.AddedThreadMember{
		ThreadMember: ThreadMember(buf.ThreadMember),
		Member:       Member(buf.Member),
		Presence:     Presence(buf.Presence),
	}
}

// Invite converts proto.Invite to discordgo.Invite
func Invite(buf *proto.Invite) *discordgo.Invite {
	if buf == nil {
		return nil
	}

	return &discordgo.Invite{
		Guild:                    Guild(buf.Guild),
		Channel:                  Channel(buf.Channel),
		Inviter:                  User(buf.Inviter),
		Code:                     buf.Code,
		CreatedAt:                TimestampValue(buf.CreatedAt),
		MaxAge:                   int(buf.MaxAge),
		Uses:                     int(buf.Uses),
		MaxUses:                  int(buf.MaxUses),
		Revoked:                  buf.Revoked,
		Temporary:                buf.Temporary,
		Unique:                   buf.Unique,
		TargetUser:               User(buf.TargetUser),
		TargetType:               discordgo.InviteTargetType(buf.TargetType),
		TargetApplication:        Application(buf.TargetApplication),
		ApproximatePresenceCount: int(buf.ApproximatePresenceCount),
		ApproximateMemberCount:   int(buf.ApproximateMemberCount),
		ExpiresAt:                Timestamp(buf.ExpiresAt),
	}
}

// VoiceRegion converts proto.VoiceRegion to discordgo.VoiceRegion
func VoiceRegion(buf *proto.VoiceRegion) *discordgo.VoiceRegion {
	if buf == nil {
		return nil
	}

	return &discordgo.VoiceRegion{
		ID:         buf.Id,
		Name:       buf.Name,
		Optimal:    buf.Optimal,
		Deprecated: buf.Deprecated,
		Custom:     buf.Custom,
	}
}

// EmojiParams converts proto.EmojiParams to discordgo.EmojiParams
func EmojiParams(buf *proto.EmojiParams) *discordgo.EmojiParams {
	if buf == nil {
		return nil
	}

	return &discordgo.EmojiParams{
		Name:  buf.Name,
		Image: buf.Image,
		Roles: buf.Roles,
	}
}
