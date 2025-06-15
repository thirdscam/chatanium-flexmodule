package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/thirdscam/chatanium-flexmodule/shared"
	"github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dgSession *discordgo.Session
	log       hclog.Logger
)

var GUILD_ID string

func main() {
	log = hclog.New(&hclog.LoggerOptions{
		Name:                 "Runtime",
		Level:                hclog.LevelFromString("DEBUG"),
		Color:                hclog.AutoColor,
		ColorHeaderAndFields: true,
	})

	godotenv.Load("./private.env")
	GUILD_ID = os.Getenv("GUILD_ID")
	if GUILD_ID == "" {
		log.Error("GUILD_ID is not set")
		os.Exit(1)
	}

	// We don't want to see the plugin logs.
	// log.SetOutput(io.Discard)

	log.Debug("starting up", "path", os.Getenv("PLUGIN_PATH"))

	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Error("Error creating Discord session", "error", err.Error())
		os.Exit(1)
	}
	dgSession = session

	dgSession.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info("Logged in", "username", s.State.User.Username, "discriminator", s.State.User.Discriminator)
	})

	dgSession.Open()

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.RuntimePluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("PLUGIN_PATH")),
		Logger:          log.ResetNamed("Module").Named("TestModule"),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Error("Error creating gRPC Client", "error", err.Error())
		os.Exit(1)
	}

	RunCoreV1(rpcClient)
	RunDiscordV1(rpcClient)

	dgSession.Close()
}

func RunCoreV1(client plugin.ClientProtocol) {
	// Request the plugin
	raw, err := client.Dispense("core-v1")
	if err != nil {
		log.Error("Core", "error", err.Error())
		os.Exit(1)
	}

	// Getting the plugin symbol
	hook, ok := raw.(core.Hook)
	if !ok {
		log.Error("Core", "error", "Plugin has no 'core-v1' plugin symbol")
		os.Exit(1)
	}

	manifest, err := hook.GetManifest()
	if err != nil {
		log.Error("Core", "error", err.Error())
		os.Exit(1)
	}
	log.Debug("Core", "manifest", hclog.Fmt("%+v", manifest))

	status, err := hook.GetStatus()
	if err != nil {
		log.Error("Core", "error", err.Error())
		os.Exit(1)
	}
	log.Debug("Core", "status", hclog.Fmt("%+v", status))

	hook.OnStage("MODULE_INIT")
	log.Debug("Core", "stage", "MODULE_INIT")

	status, err = hook.GetStatus()
	if err != nil {
		log.Error("Core", "error", err.Error())
		os.Exit(1)
	}
	log.Debug("Core", "status", hclog.Fmt("%+v", status))
}

func RunDiscordV1(client plugin.ClientProtocol) {
	// Request the plugin
	raw, err := client.Dispense("discord-v1")
	if err != nil {
		log.Error("Discord", "error", err.Error())
		os.Exit(1)
	}

	// Getting the plugin symbol
	hook, ok := raw.(discord.Hook)
	if !ok {
		log.Error("Discord", "error", "Plugin has no 'discord-v1' plugin symbol")
		os.Exit(1)
	}

	resp := hook.OnInit(nil)
	log.Debug("Discord", "initresp", hclog.Fmt("%+v", resp))
	if len(resp.Interactions) != 0 {
		for _, i := range resp.Interactions {
			log.Debug("Discord", "interaction", hclog.Fmt("%+v", i))
			_, err := dgSession.ApplicationCommandCreate(dgSession.State.User.ID, GUILD_ID, i)
			if err != nil {
				log.Error("Discord", "error", err.Error())
				os.Exit(1)
			}
		}
	}

	dgSession.AddHandler(func(s *discordgo.Session, i *discordgo.MessageCreate) {
		log.Debug("Discord", "type", "MESSAGE_CREATE", "message", hclog.Fmt("%+v", i.Message))
		hook.OnCreateChatMessage(i.Message)
	})

	dgSession.AddHandler(func(s *discordgo.Session, i *discordgo.MessageDelete) {
		log.Debug("Discord", "type", "MESSAGE_CREATE", "message", hclog.Fmt("%+v", i.Message))
		hook.OnCreateChatMessage(i.Message)
	})

	dgSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		log.Debug("Discord", "type", "INTERACTION_CREATE", "interaction", hclog.Fmt("%+v", i.Interaction))
		hook.OnCreateInteraction(i.Interaction)
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info("Ready to serve. (press Ctrl+C to exit)")
	<-stop
	fmt.Println()
	log.Info("Shutting down...")
}
