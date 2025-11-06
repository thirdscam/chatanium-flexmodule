package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc/metadata"

	broker "github.com/thirdscam/chatanium-flexmodule/shared"

	Core "github.com/thirdscam/chatanium-flexmodule/shared/core-v1"
	CorePlugin "github.com/thirdscam/chatanium-flexmodule/shared/core-v1/module"

	Discord "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
	DiscordPlugin "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/module"
	pb "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

const (
	TARGET_GUILD_ID   = "919823370600742942"
	TARGET_CHANNEL_ID = "1434561578560131245"
	TEMP_DIR          = "./temp_audio"
)

var PERMISSIONS = Core.Permissions{
	"DISCORD_V1_ON_CREATE_MESSAGE",
	"DISCORD_V1_REQ_VOICE_STATE",
	"DISCORD_V1_CREATE_VOICE_STREAM",
}

var MANIFEST = Core.Manifest{
	Name:        "VoicePlayerModule",
	Version:     "0.0.1",
	Author:      "thirdscam",
	Repository:  "github:thirdscam/chatanium-flexmodule",
	Permissions: PERMISSIONS,
}

var log hclog.Logger

type core struct {
	ready bool
}

func (m *core) GetManifest() (Core.Manifest, error) {
	return MANIFEST, nil
}

func (m *core) GetStatus() (Core.Status, error) {
	return Core.Status{
		IsReady: m.ready,
	}, nil
}

func (m *core) OnStage(stage string) {
	log.Debug("OnStage", "stage", stage)
	switch stage {
	case "MODULE_INIT":
		m.ready = true
		// Create temp directory for audio files
		os.MkdirAll(TEMP_DIR, 0755)
	case "MODULE_SHUTDOWN":
		// Cleanup temp directory
		os.RemoveAll(TEMP_DIR)
	}
}

type voicePlayer struct {
	Discord.AbstractHooks
	helper      Discord.Helper
	voiceStream pb.VoiceStreamClient
}

func (vp *voicePlayer) OnInit(h Discord.Helper) Discord.InitResponse {
	vp.helper = h
	log.Info("Voice Player Module initialized")

	return Discord.InitResponse{
		Interactions: []*discordgo.ApplicationCommand{},
	}
}

// SetVoiceStream implements VoiceStreamAware interface
// This is called automatically by the plugin system to provide VoiceStream client
func (vp *voicePlayer) SetVoiceStream(stream pb.VoiceStreamClient) {
	vp.voiceStream = stream
	log.Debug("VoiceStream client set successfully")
}

func (vp *voicePlayer) OnCreateChatMessage(m *discordgo.Message) error {
	// Ignore messages with nil Author (can happen with some Discord events)
	if m == nil || m.Author == nil {
		return nil
	}

	// Ignore bot messages
	if m.Author.Bot {
		return nil
	}

	// Check if message has audio attachments
	if len(m.Attachments) == 0 {
		return nil
	}

	for _, attachment := range m.Attachments {
		// Check if it's an audio file
		if isAudioFile(attachment.Filename) {
			log.Info("Audio file detected",
				"filename", attachment.Filename,
				"url", attachment.URL,
				"size", attachment.Size,
			)

			// Send acknowledgment
			vp.helper.ChannelMessageSend(m.ChannelID,
				fmt.Sprintf("🎵 음성 파일을 감지했습니다: `%s`\n음성 채널에 접속하여 재생합니다...", attachment.Filename))

			// Play audio in background
			go vp.playAudioFile(attachment, m.ChannelID)
		}
	}

	return nil
}

func (vp *voicePlayer) playAudioFile(attachment *discordgo.MessageAttachment, replyChannelID string) {
	// Download audio file
	localPath := filepath.Join(TEMP_DIR, attachment.Filename)
	err := downloadFile(attachment.URL, localPath)
	if err != nil {
		log.Error("Failed to download audio file", "error", err)
		vp.helper.ChannelMessageSend(replyChannelID, "❌ 음성 파일 다운로드 실패: "+err.Error())
		return
	}
	defer os.Remove(localPath)

	log.Info("Audio file downloaded", "path", localPath)

	// Convert to DCA (Discord Compatible Audio)
	dcaPath := localPath + ".dca"
	err = convertToDCA(localPath, dcaPath)
	if err != nil {
		log.Error("Failed to convert to DCA", "error", err)
		vp.helper.ChannelMessageSend(replyChannelID, "❌ 음성 변환 실패: "+err.Error())
		return
	}
	defer os.Remove(dcaPath)

	log.Info("Audio file converted to DCA", "path", dcaPath)

	// Create voice client
	voiceClient := DiscordPlugin.NewVoiceClient(vp.voiceStream, "voice-player")

	// Join voice channel
	ctx := getContextWithModuleID("voice-player")
	err = voiceClient.Join(ctx, TARGET_GUILD_ID, TARGET_CHANNEL_ID, false, false)
	if err != nil {
		log.Error("Failed to join voice channel", "error", err)
		vp.helper.ChannelMessageSend(replyChannelID, "❌ 음성 채널 접속 실패: "+err.Error())
		return
	}
	defer voiceClient.Leave()

	log.Info("Joined voice channel", "guild", TARGET_GUILD_ID, "channel", TARGET_CHANNEL_ID)
	vp.helper.ChannelMessageSend(replyChannelID, "✅ 음성 채널에 접속했습니다. 재생을 시작합니다...")

	// Wait for connection to be ready
	err = voiceClient.WaitForReady(5 * time.Second)
	if err != nil {
		log.Error("Voice connection not ready", "error", err)
		vp.helper.ChannelMessageSend(replyChannelID, "❌ 음성 연결 실패: "+err.Error())
		return
	}

	// Set speaking state
	err = voiceClient.Speaking(true)
	if err != nil {
		log.Error("Failed to set speaking state", "error", err)
	}

	// Play DCA file
	err = playDCA(voiceClient, dcaPath)
	if err != nil {
		log.Error("Failed to play audio", "error", err)
		vp.helper.ChannelMessageSend(replyChannelID, "❌ 재생 실패: "+err.Error())
		return
	}

	// Stop speaking
	voiceClient.Speaking(false)

	// Get stats
	packetsSent, _, bytesSent, _ := voiceClient.GetStats()
	log.Info("Playback completed",
		"packets_sent", packetsSent,
		"bytes_sent", bytesSent,
	)

	vp.helper.ChannelMessageSend(replyChannelID,
		fmt.Sprintf("✅ 재생 완료!\n📊 전송: %d 패킷, %d 바이트", packetsSent, bytesSent))
}

func isAudioFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	audioExtensions := []string{".mp3", ".wav", ".ogg", ".flac", ".m4a", ".aac", ".opus", ".webm"}

	for _, audioExt := range audioExtensions {
		if ext == audioExt {
			return true
		}
	}
	return false
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func getContextWithModuleID(moduleID string) context.Context {
	ctx := context.Background()
	md := metadata.New(map[string]string{
		"module-id": moduleID,
	})
	return metadata.NewOutgoingContext(ctx, md)
}

func main() {
	log = hclog.New(&hclog.LoggerOptions{
		Name:       "VoicePlayerModule",
		Level:      hclog.LevelFromString("DEBUG"),
		JSONFormat: true,
		Output:     os.Stderr,
	})

	broker.ServeToRuntime(map[string]plugin.Plugin{
		"core-v1":    &CorePlugin.Plugin{Impl: &core{}},
		"discord-v1": &DiscordPlugin.Plugin{Impl: &voicePlayer{}},
	})
}
