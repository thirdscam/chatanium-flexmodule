package discord

import (
	"github.com/bwmarrin/discordgo"
)

type InitResponse struct {
	// if bot needs to register interactions, write them here
	Interactions []*discordgo.ApplicationCommand
}

type Hook interface {
	OnInit() InitResponse
	OnCreateChatMessage(message *discordgo.Message) error
	OnCreateInteraction(interaction *discordgo.Interaction) error
	OnEvent(event string) error
}

type Helper interface {
	SendMessage(message *discordgo.Message) error
	SendInteractionResponse(interactionID, interactionToken string, response *discordgo.InteractionResponse) error
	// SendMessageToChannel(channelID string, message *discordgo.Message) error
}

// AbstractHooks is a partial implementation of the hook Interface.
//
// It is useful for embedding in a struct that only needs to
// implement a subset of the hook Interface.
type AbstractHooks struct{}

func (u *AbstractHooks) OnInit() InitResponse {
	return InitResponse{}
}

func (u *AbstractHooks) OnCreateChatMessage(m *discordgo.Message) error {
	return nil
}

func (u *AbstractHooks) OnCreateInteraction(i *discordgo.Interaction) error {
	return nil
}

func (u *AbstractHooks) OnEvent(e string) error {
	return nil
}
