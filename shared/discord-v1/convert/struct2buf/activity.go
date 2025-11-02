package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Activity converts discordgo.Activity to proto.Activity
func Activity(s *discordgo.Activity) *proto.Activity {
	if s == nil {
		return nil
	}

	return &proto.Activity{
		Name:          s.Name,
		Type:          int32(s.Type),
		Url:           s.URL,
		CreatedAt:     timestamppb.New(s.CreatedAt),
		ApplicationId: s.ApplicationID,
		State:         s.State,
		Details:       s.Details,
		Timestamps:    TimeStamps(&s.Timestamps),
		Emoji:         Emoji(&s.Emoji),
		Party:         Party(&s.Party),
		Assets:        Assets(&s.Assets),
		Secrets:       Secrets(&s.Secrets),
		Instance:      s.Instance,
		Flags:         int32(s.Flags),
	}
}

// TimeStamps converts discordgo.TimeStamps to proto.TimeStamps
func TimeStamps(s *discordgo.TimeStamps) *proto.TimeStamps {
	if s == nil {
		return nil
	}

	return &proto.TimeStamps{
		EndTimestamp:   s.EndTimestamp,
		StartTimestamp: s.StartTimestamp,
	}
}

// Assets converts discordgo.Assets to proto.Assets
func Assets(s *discordgo.Assets) *proto.Assets {
	if s == nil {
		return nil
	}

	return &proto.Assets{
		LargeImageId: s.LargeImageID,
		SmallImageId: s.SmallImageID,
		LargeText:    s.LargeText,
		SmallText:    s.SmallText,
	}
}

// Party converts discordgo.Party to proto.Party
func Party(s *discordgo.Party) *proto.Party {
	if s == nil {
		return nil
	}

	size := make([]int32, len(s.Size))
	for i, v := range s.Size {
		size[i] = int32(v)
	}

	return &proto.Party{
		Id:   s.ID,
		Size: size,
	}
}

// Secrets converts discordgo.Secrets to proto.Secrets
func Secrets(s *discordgo.Secrets) *proto.Secrets {
	if s == nil {
		return nil
	}

	return &proto.Secrets{
		Join:     s.Join,
		Spectate: s.Spectate,
		Match:    s.Match,
	}
}

// GatewayStatusUpdate converts discordgo.GatewayStatusUpdate to proto.GatewayStatusUpdate
func GatewayStatusUpdate(s *discordgo.GatewayStatusUpdate) *proto.GatewayStatusUpdate {
	if s == nil {
		return nil
	}

	return &proto.GatewayStatusUpdate{
		Since:  int32(s.Since),
		Game:   Activity(&s.Game),
		Status: s.Status,
		Afk:    s.AFK,
	}
}

// Identify converts discordgo.Identify to proto.Identify
func Identify(s *discordgo.Identify) *proto.Identify {
	if s == nil {
		return nil
	}

	var shard []int32
	if s.Shard != nil {
		shard = []int32{int32(s.Shard[0]), int32(s.Shard[1])}
	}

	return &proto.Identify{
		Token:          s.Token,
		Properties:     IdentifyProperties(&s.Properties),
		Compress:       s.Compress,
		LargeThreshold: int32(s.LargeThreshold),
		Shard:          shard,
		Presence:       GatewayStatusUpdate(&s.Presence),
		Intents:        int32(s.Intents),
	}
}

// IdentifyProperties converts discordgo.IdentifyProperties to proto.IdentifyProperties
func IdentifyProperties(s *discordgo.IdentifyProperties) *proto.IdentifyProperties {
	if s == nil {
		return nil
	}

	return &proto.IdentifyProperties{
		Os:              s.OS,
		Browser:         s.Browser,
		Device:          s.Device,
		Referer:         s.Referer,
		ReferringDomain: s.ReferringDomain,
	}
}

// Session converts discordgo.Session to proto.Session
func Session(s *discordgo.Session) *proto.Session {
	if s == nil {
		return nil
	}

	return &proto.Session{
		Token:                                s.Token,
		Mfa:                                  s.MFA,
		Debug:                                s.Debug,
		LogLevel:                             int32(s.LogLevel),
		ShouldReconnectOnError:               s.ShouldReconnectOnError,
		ShouldReconnectVoiceOnSessionError:   s.ShouldReconnectVoiceOnSessionError,
		ShouldRetryOnRateLimit:               s.ShouldRetryOnRateLimit,
		Identify:                             Identify(&s.Identify),
		Compress:                             s.Compress,
		ShardId:                              int32(s.ShardID),
		ShardCount:                           int32(s.ShardCount),
		StateEnabled:                         s.StateEnabled,
		SyncEvents:                           s.SyncEvents,
		DataReady:                            s.DataReady,
		MaxRestRetries:                       int32(s.MaxRestRetries),
		VoiceReady:                           s.VoiceReady,
		UdpReady:                             s.UDPReady,
		State:                                State(s.State),
		UserAgent:                            s.UserAgent,
		LastHeartbeatAck:                     timestamppb.New(s.LastHeartbeatAck),
		LastHeartbeatSent:                    timestamppb.New(s.LastHeartbeatSent),
	}
}

// State converts discordgo.State to proto.State
// Note: This is a simplified conversion as proto.State doesn't include all State fields
func State(s *discordgo.State) *proto.State {
	if s == nil {
		return nil
	}

	return &proto.State{
		Ready:              false, // Simplified: proto has bool, discordgo has embedded Ready struct
		MaxMessageCount:    int32(s.MaxMessageCount),
		TrackChannels:      s.TrackChannels,
		TrackThreads:       s.TrackThreads,
		TrackEmojis:        s.TrackEmojis,
		TrackMembers:       s.TrackMembers,
		TrackThreadMembers: s.TrackThreadMembers,
		TrackRoles:         s.TrackRoles,
		TrackVoice:         s.TrackVoice,
	}
}
