package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// VoiceRegion converts discordgo.VoiceRegion to proto.VoiceRegion
// Simplified version without deprecated fields
func VoiceRegion(s *discordgo.VoiceRegion) *proto.VoiceRegion {
	if s == nil {
		return nil
	}

	return &proto.VoiceRegion{
		Id:   s.ID,
		Name: s.Name,
	}
}

// GatewayBotResponse converts discordgo.GatewayBotResponse to proto.GatewayBotResponse
func GatewayBotResponse(s *discordgo.GatewayBotResponse) *proto.GatewayBotResponse {
	if s == nil {
		return nil
	}

	return &proto.GatewayBotResponse{
		Url:    s.URL,
		Shards: int32(s.Shards),
	}
}
