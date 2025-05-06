package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	broker "github.com/thirdscam/chatanium-flexmodule/shared"
	Core "github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
	Discord "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
)

var PERMISSIONS = Core.Permissions{
	"DISCORD_V1_ON_CREATE_MESSAGE",
	"DISCORD_V1_ON_CREATE_INTERACTION",
	"DISCORD_V1_REQ_VOICE_STATE",
	"DISCORD_V1_CREATE_VOICE_STREAM",
}

var MANIFEST = Core.Manifest{
	Name:        "TestModule",
	Version:     "0.0.1",
	Author:      "thirdscam",
	Repository:  "github:thirdscam/chatanium-flexmodule",
	Permissions: PERMISSIONS,
}

var log hclog.Logger

type core struct {
	ready bool
}

// GetManifest returns the manifest of the plugin.
//
// Identify modules at runtime, and proactively check for
// required permissions.
//
// Note that out-of-permission features may be ignored
// based on runtime preferences.
func (m *core) GetManifest() (Core.Manifest, error) {
	return MANIFEST, nil
}

// GetStatus returns the status of the plugin.
//
// If IsReady is not true, no hooks will be called at
// runtime for the backend extension.
func (m *core) GetStatus() (Core.Status, error) {
	return Core.Status{
		IsReady: m.ready,
	}, nil
}

// OnStage is a Hook that signals that the runtime has
// entered a particular lifecycle stage.
func (m *core) OnStage(stage string) {
	log.Debug("OnStage", "stage", stage)
	switch stage {
	case "MODULE_INIT":
		m.ready = true
	case "MODULE_START":
		// do something
	case "MODULE_SHUTDOWN":
		// do something
	default:
	}
}

type discord struct {
	// If you don't want to use all the hooks,
	// you can embed this struct, which is an empty set
	// of hooks, and use it similarly to implementing
	// an abstract class.
	Discord.AbstractHooks
}

func (u *discord) OnInit() Discord.InitResponse {
	return Discord.InitResponse{
		Interactions: []*discordgo.ApplicationCommand{
			{
				Name:        "test",
				Description: "Test command",
			},
		},
	}
}

func (u *discord) OnCreateChatMessage(m *discordgo.Message) error {
	log.Debug("MESSAGE_CREATE", "message", hclog.Fmt("%+v", m))
	return nil
}

func (u *discord) OnCreateInteraction(i *discordgo.Interaction) error {
	log.Debug("INTERACTION_CREATE", "interaction", hclog.Fmt("%+v", i))

	if i.Type == discordgo.InteractionApplicationCommand {
		log.Debug("INTERACTION_CREATE", "i.Type", i.Type, "i.Data", i.Data)
		if i.ApplicationCommandData().Name == "test" {
			log.Debug("INTERACTION_CREATE > test")
			// err := broker.SendInteractionResponse(i.ID, i.Token, &discordgo.InteractionResponse{
			// 	Type: discordgo.InteractionResponseChannelMessageWithSource,
			// 	Data: &discordgo.InteractionResponseData{
			// 		Content: "Hello, world!",
			// 	},
			// })
			// if err != nil {
			// 	log.Error("Error sending interaction response", "error", err.Error())
			// 	return err
			// }
		}
	}
	return nil
}

func main() {
	log = hclog.New(&hclog.LoggerOptions{
		Name:       "TestModule",
		Level:      hclog.LevelFromString("DEBUG"),
		JSONFormat: true,
	})

	broker.ServeToRuntime(map[string]plugin.Plugin{
		"core-v1":    &Core.Plugin{Impl: &core{}},
		"discord-v1": &Discord.Plugin{Impl: &discord{}},
	})
}
