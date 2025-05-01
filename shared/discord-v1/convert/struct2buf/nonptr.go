package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

func ForumDefaultReactionNonPtr(s discordgo.ForumDefaultReaction) *proto.ForumDefaultReaction {
	return &proto.ForumDefaultReaction{
		EmojiId:   s.EmojiID,
		EmojiName: s.EmojiName,
	}
}
