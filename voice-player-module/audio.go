package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jonas747/dca"
	DiscordModule "github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/module"
)

// convertToDCA converts an audio file to DCA format using the dca library
func convertToDCA(inputPath string, outputPath string) error {
	// Open input file
	inFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inFile.Close()

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	// Configure encoding options
	options := &dca.EncodeOptions{
		Volume:          256,     // Default volume
		Channels:        2,       // Stereo
		FrameRate:       48000,   // 48kHz
		FrameDuration:   20,      // 20ms frames
		Bitrate:         128,     // 128kbps
		Application:     dca.AudioApplicationAudio,
		CompressionLevel: 10,     // Best quality
		PacketLoss:      1,       // Assume some packet loss
		BufferedFrames:  100,     // Buffer 100 frames
		VBR:             true,    // Variable bitrate
	}

	// Create encoder session
	encodeSession, err := dca.EncodeMem(inFile, options)
	if err != nil {
		return fmt.Errorf("failed to create encode session: %w", err)
	}
	defer encodeSession.Cleanup()

	// Write DCA data to output file
	_, err = io.Copy(outFile, encodeSession)
	if err != nil {
		return fmt.Errorf("failed to write DCA data: %w", err)
	}

	// Check for encoding errors
	if err := encodeSession.Error(); err != nil {
		return fmt.Errorf("encoding error: %w", err)
	}

	log.Info("Successfully converted to DCA",
		"input", inputPath,
		"output", outputPath,
	)

	return nil
}

// playDCA plays a DCA file through the voice client
func playDCA(vc *DiscordModule.VoiceClient, dcaPath string) error {
	file, err := os.Open(dcaPath)
	if err != nil {
		return fmt.Errorf("failed to open DCA file: %w", err)
	}
	defer file.Close()

	// Create DCA decoder
	decoder := dca.NewDecoder(file)

	// Read and send opus frames
	frameCount := 0
	for {
		// Read next opus frame
		frame, err := decoder.OpusFrame()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read opus frame: %w", err)
		}

		// Send to voice channel
		vc.Send(frame, 5, 5000)
		frameCount++

		// Wait 20ms (standard opus frame duration)
		time.Sleep(20 * time.Millisecond)

		if frameCount%100 == 0 {
			log.Debug("Sending audio frames", "count", frameCount)
		}
	}

	log.Info("Finished playing audio", "total_frames", frameCount)
	return nil
}
