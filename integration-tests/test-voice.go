package main

import (
	"context"
	"os"
	"os/exec"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/thirdscam/chatanium-flexmodule/shared"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
	discordRuntime "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/runtime"
	pb "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
	"google.golang.org/grpc/metadata"
)

func main() {
	log := hclog.New(&hclog.LoggerOptions{
		Name:                 "VoiceTest",
		Level:                hclog.LevelFromString("DEBUG"),
		Color:                hclog.AutoColor,
		ColorHeaderAndFields: true,
	})

	godotenv.Load("./private.env")
	guildID := os.Getenv("GUILD_ID")
	voiceChannelID := os.Getenv("VOICE_CHANNEL_ID")

	if guildID == "" {
		log.Error("GUILD_ID is not set in private.env")
		os.Exit(1)
	}

	if voiceChannelID == "" {
		log.Error("VOICE_CHANNEL_ID is not set in private.env - please add a voice channel ID for testing")
		os.Exit(1)
	}

	// Create Discord session
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Error("Error creating Discord session", "error", err.Error())
		os.Exit(1)
	}

	err = session.Open()
	if err != nil {
		log.Error("Error opening Discord session", "error", err.Error())
		os.Exit(1)
	}
	defer session.Close()

	log.Info("Discord session opened", "username", session.State.User.Username)

	// Create Discord helper and Voice helper
	discordHelper := discordRuntime.NewDiscordHelper(session)
	voiceHelper := discordRuntime.NewVoiceHelper(session, log)

	// Create runtime plugin map
	runtimePluginMap := shared.CreateRuntimePluginMap(discordHelper, voiceHelper)

	// Launch plugin
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         runtimePluginMap,
		Cmd:             exec.Command(os.Getenv("PLUGIN_PATH")),
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

	// Get discord-v1 plugin
	raw, err := rpcClient.Dispense("discord-v1")
	if err != nil {
		log.Error("Discord dispense error", "error", err.Error())
		os.Exit(1)
	}

	runtimeClients, ok := raw.(discord.RuntimeClients)
	if !ok {
		log.Error("Plugin has no 'discord-v1' plugin symbol")
		os.Exit(1)
	}

	// Get VoiceStream client
	voiceStream := runtimeClients.GetVoiceStream()

	log.Info("=== Starting Voice Streaming Integration Tests ===")

	// Run voice tests
	runVoiceIntegrationTests(voiceStream, guildID, voiceChannelID, log)

	log.Info("=== Voice Tests Completed Successfully ===")
}

func runVoiceIntegrationTests(voiceStream pb.VoiceStreamClient, guildID, channelID string, log hclog.Logger) {
	// Test 1: VoiceJoin and VoiceLeave
	log.Info("Test 1: VoiceJoin and VoiceLeave")
	testVoiceJoinLeave(voiceStream, guildID, channelID, log)

	// Test 2: Queue status query
	log.Info("Test 2: Queue Status Query")
	testQueueStatus(voiceStream, guildID, log)

	// Test 3: Voice send test (would require actual voice test module)
	log.Info("Test 3: Voice Send/Receive (requires test module implementation)")
	log.Warn("Voice send/receive test should be run with test-module voice_test.go")
}

func testVoiceJoinLeave(voiceStream pb.VoiceStreamClient, guildID, channelID string, log hclog.Logger) {
	ctx := getContextWithModuleID("test-integration")

	// Join voice channel
	resp, err := voiceStream.VoiceJoin(ctx, &pb.VoiceJoinRequest{
		GuildId:   guildID,
		ChannelId: channelID,
		Mute:      false,
		Deaf:      false,
	})
	if err != nil {
		log.Error("VoiceJoin failed", "error", err)
		panic(err)
	}

	log.Info("✅ VoiceJoin succeeded",
		"connection_id", resp.ConnectionId,
		"ready", resp.Ready,
		"queue_length", resp.QueueLength,
		"active_subscribers", resp.ActiveSubscribers,
	)

	// Leave voice channel
	_, err = voiceStream.VoiceLeave(ctx, &pb.VoiceLeaveRequest{
		ConnectionId: resp.ConnectionId,
	})
	if err != nil {
		log.Error("VoiceLeave failed", "error", err)
		panic(err)
	}

	log.Info("✅ VoiceLeave succeeded")
}

func testQueueStatus(voiceStream pb.VoiceStreamClient, guildID string, log hclog.Logger) {
	ctx := getContextWithModuleID("test-integration")

	resp, err := voiceStream.GetQueueStatus(ctx, &pb.QueueStatusRequest{
		GuildId: guildID,
	})

	// It's okay if there's no active session
	if err != nil {
		log.Info("No active voice session (expected if not joined)", "error", err)
		return
	}

	log.Info("✅ GetQueueStatus succeeded",
		"queue_length", resp.QueueLength,
		"active_subscribers", resp.ActiveSubscribers,
		"module_count", len(resp.ModuleStats),
	)

	for _, stats := range resp.ModuleStats {
		log.Info("Module stats",
			"module_id", stats.ModuleId,
			"tasks_submitted", stats.TasksSubmitted,
			"tasks_processed", stats.TasksProcessed,
			"bytes_sent", stats.BytesSent,
			"bytes_received", stats.BytesReceived,
		)
	}
}

func getContextWithModuleID(moduleID string) context.Context {
	ctx := context.Background()
	md := metadata.New(map[string]string{
		"module-id": moduleID,
	})
	return metadata.NewOutgoingContext(ctx, md)
}
