package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	broker "github.com/thirdscam/chatanium-flexmodule/shared"

	Core "github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
	CorePlugin "github.com/thirdscam/chatanium-flexmodule/shared/core-v1/module"

	Discord "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
	DiscordPlugin "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/module"
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
	// You don't necessarily need to use a variable named ready for your Core implementation like this.
	// However, we recommend that you implement some other logic that allows you to control the IsReady field of Core.Status.
	ready bool
}

// GetManifest returns the manifest of the plugin.
//
// Identify modules at runtime, and proactively check for required permissions.
//
// Note: out-of-permission features may be ignored based on runtime preferences.
func (m *core) GetManifest() (Core.Manifest, error) {
	return MANIFEST, nil
}

// GetStatus returns the status of the plugin.
//
// If IsReady is not true, no hooks will be called at runtime for the backend extension.
func (m *core) GetStatus() (Core.Status, error) {
	return Core.Status{
		IsReady: m.ready,
	}, nil
}

// OnStage is a Hook that signals that the runtime has entered a particular lifecycle stage.
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
	// If you don't want to use all the hooks, you can embed this struct, which is an empty set
	// of hooks, and use it similarly to implementing an abstract class.
	Discord.AbstractHooks
	helper Discord.Helper // Store helper instance for use in other methods
}

func (u *discord) OnInit(h Discord.Helper) Discord.InitResponse {
	// Store helper for later use
	u.helper = h

	return Discord.InitResponse{
		Interactions: []*discordgo.ApplicationCommand{
			{
				Name:        "test",
				Description: "Test command",
			},
			{
				Name:        "hello",
				Description: "Say hello",
			},
		},
	}
}

func (u *discord) OnCreateChatMessage(m *discordgo.Message) error {
	log.Debug("MESSAGE_CREATE", "message", hclog.Fmt("%+v", m))

	// Don't respond to bot messages to avoid loops
	if m.Author.Bot {
		return nil
	}

	// Example: respond to "ping" with "pong"
	if m.Content == "ping" && u.helper != nil {
		_, err := u.helper.ChannelMessageSend(m.ChannelID, "pong! 🏓")
		if err != nil {
			log.Error("Failed to send pong message", "error", err)
			return err
		}
	}

	// Example: respond to "hello plugin" with an embed
	if m.Content == "hello plugin" && u.helper != nil {
		embed := &discordgo.MessageEmbed{
			Title:       "Hello from Plugin!",
			Description: "This is a response from the test plugin using helper functions.",
			Color:       0x00ff00,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Plugin Status",
					Value:  "✅ Working correctly",
					Inline: true,
				},
				{
					Name:   "Helper Functions",
					Value:  "✅ Available",
					Inline: true,
				},
			},
		}

		_, err := u.helper.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			log.Error("Failed to send embed message", "error", err)
			return err
		}
	}

	return nil
}

func (u *discord) OnCreateInteraction(i *discordgo.Interaction) error {
	log.Debug("INTERACTION_CREATE", "interaction", hclog.Fmt("%+v", i))

	if i.Type == discordgo.InteractionApplicationCommand {
		log.Debug("INTERACTION_CREATE", "i.Type", i.Type, "i.Data", i.Data)
		cmdData := i.ApplicationCommandData()

		switch cmdData.Name {
		case "test":
			log.Debug("INTERACTION_CREATE > test")
			// Respond to the test command
			if u.helper != nil {
				err := u.helper.InteractionRespond(i, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "This is a test response from the plugin! 🎉",
					},
				})
				if err != nil {
					log.Error("Failed to respond to interaction", "error", err)
					return err
				}
			}
		case "hello":
			log.Debug("INTERACTION_CREATE > hello")
			// Send a hello message
			if u.helper != nil {
				err := u.helper.InteractionRespond(i, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello! 👋 This message was sent using the helper function from the plugin!",
						Embeds: []*discordgo.MessageEmbed{
							{
								Title:       "Plugin Helper Test",
								Description: "This embed was created using the Discord helper functions!",
								Color:       0x00ff00,
							},
						},
					},
				})
				if err != nil {
					log.Error("Failed to respond to interaction", "error", err)
					return err
				}
			}
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
		"core-v1":    &CorePlugin.Plugin{Impl: &core{}},
		"discord-v1": &DiscordPlugin.Plugin{Impl: &discord{}},
	})
}
