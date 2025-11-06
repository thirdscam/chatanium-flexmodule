package buf2struct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// VoiceRegion converts proto.VoiceRegion to discordgo.VoiceRegion
// Simplified version without deprecated fields
func VoiceRegion(buf *proto.VoiceRegion) *discordgo.VoiceRegion {
	if buf == nil {
		return nil
	}

	return &discordgo.VoiceRegion{
		ID:   buf.Id,
		Name: buf.Name,
	}
}

// GatewayBotResponse converts proto.GatewayBotResponse to discordgo.GatewayBotResponse
func GatewayBotResponse(buf *proto.GatewayBotResponse) *discordgo.GatewayBotResponse {
	if buf == nil {
		return nil
	}

	return &discordgo.GatewayBotResponse{
		URL:    buf.Url,
		Shards: int(buf.Shards),
	}
}
