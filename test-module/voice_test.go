package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	Discord "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1"
	DiscordModule "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/module"
	pb "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// VoiceTestModule demonstrates voice streaming capabilities
type VoiceTestModule struct {
	helper      Discord.Helper
	voiceClient *DiscordModule.VoiceClient
	log         hclog.Logger
}

// NewVoiceTestModule creates a new voice test module
func NewVoiceTestModule(helper Discord.Helper, voiceStream pb.VoiceStreamClient, log hclog.Logger) *VoiceTestModule {
	return &VoiceTestModule{
		helper:      helper,
		voiceClient: DiscordModule.NewVoiceClient(voiceStream, "test-module"),
		log:         log.Named("voice-test"),
	}
}

// TestVoiceSend demonstrates sending voice data
func (v *VoiceTestModule) TestVoiceSend(guildID, channelID string) error {
	v.log.Info("Starting voice send test", "guild_id", guildID, "channel_id", channelID)

	// Join voice channel
	ctx := context.Background()
	err := v.voiceClient.Join(ctx, guildID, channelID, false, false)
	if err != nil {
		v.log.Error("Failed to join voice channel", "error", err)
		return err
	}
	defer v.voiceClient.Leave()

	v.log.Info("Joined voice channel successfully")

	// Wait for stream to be ready
	err = v.voiceClient.WaitForReady(5 * time.Second)
	if err != nil {
		v.log.Error("Stream not ready", "error", err)
		return err
	}

	// Set speaking state
	err = v.voiceClient.Speaking(true)
	if err != nil {
		v.log.Error("Failed to set speaking state", "error", err)
		return err
	}

	v.log.Info("Starting to send voice packets...")

	// Send 100 test packets (2 seconds of audio at 20ms per frame)
	for i := 0; i < 100; i++ {
		// Generate test opus data (in real scenario, this would be actual opus-encoded audio)
		opusData := generateTestOpusFrame(i)

		// Send with priority (higher numbers = higher priority)
		priority := int32(5)
		timeoutMs := int64(5000) // 5 second timeout

		v.voiceClient.Send(opusData, priority, timeoutMs)

		v.log.Debug("Sent voice packet", "packet_num", i+1)

		// Wait 20ms (standard opus frame duration)
		time.Sleep(20 * time.Millisecond)
	}

	v.log.Info("Finished sending voice packets")

	// Stop speaking
	err = v.voiceClient.Speaking(false)
	if err != nil {
		v.log.Error("Failed to unset speaking state", "error", err)
	}

	// Get stats
	packetsSent, packetsRecv, bytesSent, bytesRecv := v.voiceClient.GetStats()
	v.log.Info("Voice client statistics",
		"packets_sent", packetsSent,
		"packets_recv", packetsRecv,
		"bytes_sent", bytesSent,
		"bytes_recv", bytesRecv,
	)

	return nil
}

// TestVoiceReceive demonstrates receiving voice data
func (v *VoiceTestModule) TestVoiceReceive(guildID, channelID string, duration time.Duration) error {
	v.log.Info("Starting voice receive test", "guild_id", guildID, "channel_id", channelID, "duration", duration)

	// Join voice channel
	ctx := context.Background()
	err := v.voiceClient.Join(ctx, guildID, channelID, false, false)
	if err != nil {
		v.log.Error("Failed to join voice channel", "error", err)
		return err
	}
	defer v.voiceClient.Leave()

	v.log.Info("Joined voice channel successfully, listening for voice...")

	// Wait for stream to be ready
	err = v.voiceClient.WaitForReady(5 * time.Second)
	if err != nil {
		v.log.Error("Stream not ready", "error", err)
		return err
	}

	// Listen for voice packets
	timeout := time.After(duration)
	packetCount := 0

	for {
		select {
		case packet := <-v.voiceClient.RecvChan:
			packetCount++
			v.log.Debug("Received voice packet",
				"packet_num", packetCount,
				"ssrc", packet.SSRC,
				"sequence", packet.Sequence,
				"timestamp", packet.Timestamp,
				"data_size", len(packet.OpusData),
			)

		case <-timeout:
			v.log.Info("Voice receive test completed",
				"total_packets", packetCount,
				"duration", duration,
			)
			return nil
		}
	}
}

// TestBidirectional demonstrates both sending and receiving simultaneously
func (v *VoiceTestModule) TestBidirectional(guildID, channelID string, sendDuration time.Duration) error {
	v.log.Info("Starting bidirectional voice test", "guild_id", guildID, "channel_id", channelID)

	// Join voice channel
	ctx := context.Background()
	err := v.voiceClient.Join(ctx, guildID, channelID, false, false)
	if err != nil {
		v.log.Error("Failed to join voice channel", "error", err)
		return err
	}
	defer v.voiceClient.Leave()

	v.log.Info("Joined voice channel, starting bidirectional stream...")

	// Wait for stream to be ready
	err = v.voiceClient.WaitForReady(5 * time.Second)
	if err != nil {
		v.log.Error("Stream not ready", "error", err)
		return err
	}

	// Start receive goroutine
	recvCount := 0
	go func() {
		for packet := range v.voiceClient.RecvChan {
			recvCount++
			if recvCount%10 == 0 {
				v.log.Info("Received voice packets", "count", recvCount)
			}
			v.log.Debug("Received voice packet",
				"ssrc", packet.SSRC,
				"data_size", len(packet.OpusData),
			)
		}
	}()

	// Send voice data
	v.voiceClient.Speaking(true)
	defer v.voiceClient.Speaking(false)

	sendCount := int(sendDuration.Milliseconds() / 20) // 20ms per frame
	for i := 0; i < sendCount; i++ {
		opusData := generateTestOpusFrame(i)
		v.voiceClient.Send(opusData, 5, 5000)

		if (i+1)%50 == 0 {
			v.log.Info("Sent voice packets", "count", i+1)
		}

		time.Sleep(20 * time.Millisecond)
	}

	v.log.Info("Bidirectional test completed",
		"packets_sent", sendCount,
		"packets_received", recvCount,
	)

	return nil
}

// generateTestOpusFrame generates a test opus frame
// In a real implementation, this would be actual opus-encoded audio data
func generateTestOpusFrame(frameNum int) []byte {
	// Generate dummy opus data (60 bytes is a typical small opus frame)
	data := make([]byte, 60)
	// Fill with pattern based on frame number for testing
	for i := range data {
		data[i] = byte((frameNum + i) % 256)
	}
	return data
}

// RunVoiceTests runs all voice tests
func RunVoiceTests(helper Discord.Helper, voiceStream pb.VoiceStreamClient, guildID, channelID string, log hclog.Logger) {
	module := NewVoiceTestModule(helper, voiceStream, log)

	log.Info("=== Starting Voice Tests ===")

	// Test 1: Send only
	log.Info("Test 1: Voice Send Test")
	err := module.TestVoiceSend(guildID, channelID)
	if err != nil {
		log.Error("Voice send test failed", "error", err)
	}
	time.Sleep(2 * time.Second)

	// Test 2: Receive only
	log.Info("Test 2: Voice Receive Test")
	err = module.TestVoiceReceive(guildID, channelID, 10*time.Second)
	if err != nil {
		log.Error("Voice receive test failed", "error", err)
	}
	time.Sleep(2 * time.Second)

	// Test 3: Bidirectional
	log.Info("Test 3: Bidirectional Voice Test")
	err = module.TestBidirectional(guildID, channelID, 5*time.Second)
	if err != nil {
		log.Error("Bidirectional test failed", "error", err)
	}

	log.Info("=== Voice Tests Completed ===")
}

// GetVoiceQueueStatus queries queue status for debugging
func GetVoiceQueueStatus(voiceStream pb.VoiceStreamClient, guildID string, log hclog.Logger) error {
	ctx := context.Background()
	resp, err := voiceStream.GetQueueStatus(ctx, &pb.QueueStatusRequest{
		GuildId: guildID,
	})
	if err != nil {
		return fmt.Errorf("failed to get queue status: %w", err)
	}

	log.Info("Queue Status",
		"queue_length", resp.QueueLength,
		"active_subscribers", resp.ActiveSubscribers,
	)

	for _, stats := range resp.ModuleStats {
		log.Info("Module Stats",
			"module_id", stats.ModuleId,
			"tasks_submitted", stats.TasksSubmitted,
			"tasks_processed", stats.TasksProcessed,
			"bytes_sent", stats.BytesSent,
			"bytes_received", stats.BytesReceived,
		)
	}

	return nil
}
