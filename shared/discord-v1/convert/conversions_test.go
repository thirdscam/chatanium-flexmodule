package convert_test

import (
	"os"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/buf2struct"
	"github.com/thirdscam/chatanium-flexmodule/shared/discord-v1/convert/struct2buf"
)

// TestActivity tests Activity conversion
func TestActivity(t *testing.T) {
	originalActivity := &discordgo.Activity{
		Name:          "Test Game",
		Type:          discordgo.ActivityTypeGame,
		URL:           "https://test.com",
		CreatedAt:     time.Now(),
		ApplicationID: "123456",
		State:         "Playing",
		Details:       "In match",
		Timestamps: discordgo.TimeStamps{
			StartTimestamp: 1234567890,
			EndTimestamp:   1234567900,
		},
		Emoji: discordgo.Emoji{
			ID:       "emoji123",
			Name:     "smile",
			Animated: true,
		},
		Party: discordgo.Party{
			ID:   "party123",
			Size: []int{2, 4},
		},
		Assets: discordgo.Assets{
			LargeImageID: "large_img",
			SmallImageID: "small_img",
			LargeText:    "Large",
			SmallText:    "Small",
		},
		Secrets: discordgo.Secrets{
			Join:     "join_secret",
			Spectate: "spectate_secret",
			Match:    "match_secret",
		},
		Instance: true,
		Flags:    1,
	}

	// Convert to proto
	protoActivity := struct2buf.Activity(originalActivity)
	if protoActivity == nil {
		t.Fatal("struct2buf.Activity returned nil")
	}

	// Convert back to struct
	convertedActivity := buf2struct.Activity(protoActivity)
	if convertedActivity == nil {
		t.Fatal("buf2struct.Activity returned nil")
	}

	// Verify key fields
	if convertedActivity.Name != originalActivity.Name {
		t.Errorf("Name mismatch: got %s, want %s", convertedActivity.Name, originalActivity.Name)
	}
	if convertedActivity.State != originalActivity.State {
		t.Errorf("State mismatch: got %s, want %s", convertedActivity.State, originalActivity.State)
	}
	if convertedActivity.Instance != originalActivity.Instance {
		t.Errorf("Instance mismatch: got %v, want %v", convertedActivity.Instance, originalActivity.Instance)
	}
}

// TestUser tests User conversion
func TestUser(t *testing.T) {
	originalUser := &discordgo.User{
		ID:            "123456789",
		Username:      "testuser",
		Discriminator: "0001",
		Avatar:        "avatar_hash",
		Bot:           false,
		PublicFlags:   discordgo.UserFlags(1),
		Locale:        "en-US",
		Verified:      true,
		Email:         "test@example.com",
		Flags:         2,
	}

	// Convert to proto
	protoUser := struct2buf.User(originalUser)
	if protoUser == nil {
		t.Fatal("struct2buf.User returned nil")
	}

	// Convert back to struct
	convertedUser := buf2struct.User(protoUser)
	if convertedUser == nil {
		t.Fatal("buf2struct.User returned nil")
	}

	// Verify fields
	if convertedUser.ID != originalUser.ID {
		t.Errorf("ID mismatch: got %s, want %s", convertedUser.ID, originalUser.ID)
	}
	if convertedUser.Username != originalUser.Username {
		t.Errorf("Username mismatch: got %s, want %s", convertedUser.Username, originalUser.Username)
	}
	if convertedUser.Bot != originalUser.Bot {
		t.Errorf("Bot mismatch: got %v, want %v", convertedUser.Bot, originalUser.Bot)
	}
	if convertedUser.Verified != originalUser.Verified {
		t.Errorf("Verified mismatch: got %v, want %v", convertedUser.Verified, originalUser.Verified)
	}
}

// TestEmoji tests Emoji conversion
func TestEmoji(t *testing.T) {
	originalEmoji := &discordgo.Emoji{
		ID:            "emoji123",
		Name:          "test_emoji",
		Roles:         []string{"role1", "role2"},
		RequireColons: true,
		Managed:       false,
		Animated:      true,
		Available:     true,
	}

	// Convert to proto
	protoEmoji := struct2buf.Emoji(originalEmoji)
	if protoEmoji == nil {
		t.Fatal("struct2buf.Emoji returned nil")
	}

	// Convert back to struct
	convertedEmoji := buf2struct.Emoji(protoEmoji)
	if convertedEmoji == nil {
		t.Fatal("buf2struct.Emoji returned nil")
	}

	// Verify fields
	if convertedEmoji.ID != originalEmoji.ID {
		t.Errorf("ID mismatch: got %s, want %s", convertedEmoji.ID, originalEmoji.ID)
	}
	if convertedEmoji.Name != originalEmoji.Name {
		t.Errorf("Name mismatch: got %s, want %s", convertedEmoji.Name, originalEmoji.Name)
	}
	if convertedEmoji.Animated != originalEmoji.Animated {
		t.Errorf("Animated mismatch: got %v, want %v", convertedEmoji.Animated, originalEmoji.Animated)
	}
}

// TestPoll tests Poll conversion
func TestPoll(t *testing.T) {
	originalPoll := &discordgo.Poll{
		Question: discordgo.PollMedia{
			Text: "What is your favorite color?",
		},
		Answers: []discordgo.PollAnswer{
			{
				AnswerID: 1,
				Media: &discordgo.PollMedia{
					Text: "Red",
				},
			},
			{
				AnswerID: 2,
				Media: &discordgo.PollMedia{
					Text: "Blue",
				},
			},
		},
		AllowMultiselect: false,
		LayoutType:       1,
		Duration:         24,
	}

	// Convert to proto
	protoPoll := struct2buf.Poll(originalPoll)
	if protoPoll == nil {
		t.Fatal("struct2buf.Poll returned nil")
	}

	// Convert back to struct
	convertedPoll := buf2struct.Poll(protoPoll)
	if convertedPoll == nil {
		t.Fatal("buf2struct.Poll returned nil")
	}

	// Verify fields
	if convertedPoll.Question.Text != originalPoll.Question.Text {
		t.Errorf("Question text mismatch: got %s, want %s", convertedPoll.Question.Text, originalPoll.Question.Text)
	}
	if len(convertedPoll.Answers) != len(originalPoll.Answers) {
		t.Errorf("Answers length mismatch: got %d, want %d", len(convertedPoll.Answers), len(originalPoll.Answers))
	}
	if convertedPoll.AllowMultiselect != originalPoll.AllowMultiselect {
		t.Errorf("AllowMultiselect mismatch: got %v, want %v", convertedPoll.AllowMultiselect, originalPoll.AllowMultiselect)
	}
}

// TestGatewayBotResponse tests GatewayBotResponse conversion
func TestGatewayBotResponse(t *testing.T) {
	originalResponse := &discordgo.GatewayBotResponse{
		URL:    "wss://gateway.discord.gg",
		Shards: 1,
		SessionStartLimit: discordgo.SessionInformation{
			Total:          1000,
			Remaining:      999,
			ResetAfter:     86400000,
			MaxConcurrency: 1,
		},
	}

	// Convert to proto
	protoResponse := struct2buf.GatewayBotResponse(originalResponse)
	if protoResponse == nil {
		t.Fatal("struct2buf.GatewayBotResponse returned nil")
	}

	// Convert back to struct
	convertedResponse := buf2struct.GatewayBotResponse(protoResponse)
	if convertedResponse == nil {
		t.Fatal("buf2struct.GatewayBotResponse returned nil")
	}

	// Verify fields
	if convertedResponse.URL != originalResponse.URL {
		t.Errorf("URL mismatch: got %s, want %s", convertedResponse.URL, originalResponse.URL)
	}
	if convertedResponse.Shards != originalResponse.Shards {
		t.Errorf("Shards mismatch: got %d, want %d", convertedResponse.Shards, originalResponse.Shards)
	}
	if convertedResponse.SessionStartLimit.Total != originalResponse.SessionStartLimit.Total {
		t.Errorf("SessionStartLimit.Total mismatch: got %d, want %d",
			convertedResponse.SessionStartLimit.Total, originalResponse.SessionStartLimit.Total)
	}
}

// TestApplication tests Application conversion
func TestApplication(t *testing.T) {
	originalApp := &discordgo.Application{
		ID:                  "app123",
		Name:                "Test App",
		Icon:                "icon_hash",
		Description:         "A test application",
		RPCOrigins:          []string{"https://test.com"},
		BotPublic:           true,
		BotRequireCodeGrant: false,
		TermsOfServiceURL:   "https://test.com/tos",
		PrivacyProxyURL:     "https://test.com/privacy",
		Summary:             "Test summary",
		VerifyKey:           "verify_key",
		GuildID:             "guild123",
		PrimarySKUID:        "sku123",
		Slug:                "test-app",
		CoverImage:          "cover_hash",
		Flags:               1,
	}

	// Convert to proto
	protoApp := struct2buf.Application(originalApp)
	if protoApp == nil {
		t.Fatal("struct2buf.Application returned nil")
	}

	// Convert back to struct
	convertedApp := buf2struct.Application(protoApp)
	if convertedApp == nil {
		t.Fatal("buf2struct.Application returned nil")
	}

	// Verify fields
	if convertedApp.ID != originalApp.ID {
		t.Errorf("ID mismatch: got %s, want %s", convertedApp.ID, originalApp.ID)
	}
	if convertedApp.Name != originalApp.Name {
		t.Errorf("Name mismatch: got %s, want %s", convertedApp.Name, originalApp.Name)
	}
	if convertedApp.BotPublic != originalApp.BotPublic {
		t.Errorf("BotPublic mismatch: got %v, want %v", convertedApp.BotPublic, originalApp.BotPublic)
	}
}

// TestNilSafety tests that nil values are handled properly
func TestNilSafety(t *testing.T) {
	// Test nil inputs
	if struct2buf.Activity(nil) != nil {
		t.Error("struct2buf.Activity should return nil for nil input")
	}
	if struct2buf.User(nil) != nil {
		t.Error("struct2buf.User should return nil for nil input")
	}
	if struct2buf.Emoji(nil) != nil {
		t.Error("struct2buf.Emoji should return nil for nil input")
	}
	if struct2buf.Poll(nil) != nil {
		t.Error("struct2buf.Poll should return nil for nil input")
	}

	// Test nil outputs
	if buf2struct.Activity(nil) != nil {
		t.Error("buf2struct.Activity should return nil for nil input")
	}
	if buf2struct.User(nil) != nil {
		t.Error("buf2struct.User should return nil for nil input")
	}
	if buf2struct.Emoji(nil) != nil {
		t.Error("buf2struct.Emoji should return nil for nil input")
	}
	if buf2struct.Poll(nil) != nil {
		t.Error("buf2struct.Poll should return nil for nil input")
	}
}

// TestSessionConversion tests Session conversion (requires Discord token)
func TestSessionConversion(t *testing.T) {
	// Skip if no Discord token available
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		t.Skip("Skipping session conversion test: DISCORD_TOKEN not set")
	}

	// Create a basic session struct (not a real Discord session)
	originalSession := &discordgo.Session{
		Token:                  token,
		MFA:                    false,
		Debug:                  false,
		LogLevel:               0,
		ShouldReconnectOnError: true,
		Compress:               true,
		ShardID:                0,
		ShardCount:             1,
		StateEnabled:           true,
		MaxRestRetries:         3,
		UserAgent:              "DiscordBot",
	}

	// Convert to proto
	protoSession := struct2buf.Session(originalSession)
	if protoSession == nil {
		t.Fatal("struct2buf.Session returned nil")
	}

	// Convert back to struct
	convertedSession := buf2struct.Session(protoSession)
	if convertedSession == nil {
		t.Fatal("buf2struct.Session returned nil")
	}

	// Verify fields
	if convertedSession.Token != originalSession.Token {
		t.Error("Token mismatch")
	}
	if convertedSession.ShardID != originalSession.ShardID {
		t.Errorf("ShardID mismatch: got %d, want %d", convertedSession.ShardID, originalSession.ShardID)
	}
	if convertedSession.Compress != originalSession.Compress {
		t.Errorf("Compress mismatch: got %v, want %v", convertedSession.Compress, originalSession.Compress)
	}
}
