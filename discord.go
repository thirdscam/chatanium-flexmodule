package main

import "github.com/bwmarrin/discordgo"

// test discord implementation
// this implementation only use for PoC implementation
type DiscordHelper struct {
	session *discordgo.Session
}

func (d *DiscordHelper) Channel(channelID string) (discordgo.Channel, error) {
	st, err := d.session.Channel(channelID)
	if err != nil {
		return discordgo.Channel{}, err
	}

	return st, nil
}
