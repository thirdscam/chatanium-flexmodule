package buf2struct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// Activity converts proto.Activity to discordgo.Activity
func Activity(buf *proto.Activity) *discordgo.Activity {
	if buf == nil {
		return nil
	}

	return &discordgo.Activity{
		Name:          buf.Name,
		Type:          discordgo.ActivityType(buf.Type),
		URL:           buf.Url,
		CreatedAt:     buf.CreatedAt.AsTime(),
		ApplicationID: buf.ApplicationId,
		State:         buf.State,
		Details:       buf.Details,
		Timestamps:    *TimeStamps(buf.Timestamps),
		Emoji:         *Emoji(buf.Emoji),
		Party:         *Party(buf.Party),
		Assets:        *Assets(buf.Assets),
		Secrets:       *Secrets(buf.Secrets),
		Instance:      buf.Instance,
		Flags:         int(buf.Flags),
	}
}

// TimeStamps converts proto.TimeStamps to discordgo.TimeStamps
func TimeStamps(buf *proto.TimeStamps) *discordgo.TimeStamps {
	if buf == nil {
		return &discordgo.TimeStamps{}
	}

	return &discordgo.TimeStamps{
		EndTimestamp:   buf.EndTimestamp,
		StartTimestamp: buf.StartTimestamp,
	}
}

// Assets converts proto.Assets to discordgo.Assets
func Assets(buf *proto.Assets) *discordgo.Assets {
	if buf == nil {
		return &discordgo.Assets{}
	}

	return &discordgo.Assets{
		LargeImageID: buf.LargeImageId,
		SmallImageID: buf.SmallImageId,
		LargeText:    buf.LargeText,
		SmallText:    buf.SmallText,
	}
}

// Party converts proto.Party to discordgo.Party
func Party(buf *proto.Party) *discordgo.Party {
	if buf == nil {
		return &discordgo.Party{}
	}

	size := make([]int, len(buf.Size))
	for i, v := range buf.Size {
		size[i] = int(v)
	}

	return &discordgo.Party{
		ID:   buf.Id,
		Size: size,
	}
}

// Secrets converts proto.Secrets to discordgo.Secrets
func Secrets(buf *proto.Secrets) *discordgo.Secrets {
	if buf == nil {
		return &discordgo.Secrets{}
	}

	return &discordgo.Secrets{
		Join:     buf.Join,
		Spectate: buf.Spectate,
		Match:    buf.Match,
	}
}

// GatewayStatusUpdate converts proto.GatewayStatusUpdate to discordgo.GatewayStatusUpdate
func GatewayStatusUpdate(buf *proto.GatewayStatusUpdate) *discordgo.GatewayStatusUpdate {
	if buf == nil {
		return nil
	}

	var game discordgo.Activity
	if buf.Game != nil {
		game = *Activity(buf.Game)
	}

	return &discordgo.GatewayStatusUpdate{
		Since:  int(buf.Since),
		Game:   game,
		Status: buf.Status,
		AFK:    buf.Afk,
	}
}

// Identify converts proto.Identify to discordgo.Identify
func Identify(buf *proto.Identify) *discordgo.Identify {
	if buf == nil {
		return nil
	}

	var shard *[2]int
	if len(buf.Shard) >= 2 {
		shard = &[2]int{int(buf.Shard[0]), int(buf.Shard[1])}
	}

	var presence discordgo.GatewayStatusUpdate
	if buf.Presence != nil {
		presence = *GatewayStatusUpdate(buf.Presence)
	}

	return &discordgo.Identify{
		Token:          buf.Token,
		Properties:     *IdentifyProperties(buf.Properties),
		Compress:       buf.Compress,
		LargeThreshold: int(buf.LargeThreshold),
		Shard:          shard,
		Presence:       presence,
		Intents:        discordgo.Intent(buf.Intents),
	}
}

// IdentifyProperties converts proto.IdentifyProperties to discordgo.IdentifyProperties
func IdentifyProperties(buf *proto.IdentifyProperties) *discordgo.IdentifyProperties {
	if buf == nil {
		return &discordgo.IdentifyProperties{}
	}

	return &discordgo.IdentifyProperties{
		OS:              buf.Os,
		Browser:         buf.Browser,
		Device:          buf.Device,
		Referer:         buf.Referer,
		ReferringDomain: buf.ReferringDomain,
	}
}

// Session converts proto.Session to discordgo.Session
func Session(buf *proto.Session) *discordgo.Session {
	if buf == nil {
		return nil
	}

	var identify discordgo.Identify
	if buf.Identify != nil {
		identify = *Identify(buf.Identify)
	}

	return &discordgo.Session{
		Token:                                buf.Token,
		MFA:                                  buf.Mfa,
		Debug:                                buf.Debug,
		LogLevel:                             int(buf.LogLevel),
		ShouldReconnectOnError:               buf.ShouldReconnectOnError,
		ShouldReconnectVoiceOnSessionError:   buf.ShouldReconnectVoiceOnSessionError,
		ShouldRetryOnRateLimit:               buf.ShouldRetryOnRateLimit,
		Identify:                             identify,
		Compress:                             buf.Compress,
		ShardID:                              int(buf.ShardId),
		ShardCount:                           int(buf.ShardCount),
		StateEnabled:                         buf.StateEnabled,
		SyncEvents:                           buf.SyncEvents,
		DataReady:                            buf.DataReady,
		MaxRestRetries:                       int(buf.MaxRestRetries),
		VoiceReady:                           buf.VoiceReady,
		UDPReady:                             buf.UdpReady,
		State:                                State(buf.State),
		UserAgent:                            buf.UserAgent,
		LastHeartbeatAck:                     buf.LastHeartbeatAck.AsTime(),
		LastHeartbeatSent:                    buf.LastHeartbeatSent.AsTime(),
	}
}

// State converts proto.State to discordgo.State
// Note: This is a simplified conversion as proto.State doesn't include all State fields
func State(buf *proto.State) *discordgo.State {
	if buf == nil {
		return nil
	}

	return &discordgo.State{
		MaxMessageCount:    int(buf.MaxMessageCount),
		TrackChannels:      buf.TrackChannels,
		TrackThreads:       buf.TrackThreads,
		TrackEmojis:        buf.TrackEmojis,
		TrackMembers:       buf.TrackMembers,
		TrackThreadMembers: buf.TrackThreadMembers,
		TrackRoles:         buf.TrackRoles,
		TrackVoice:         buf.TrackVoice,
	}
}
