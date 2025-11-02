package buf2struct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// Poll converts proto.Poll to discordgo.Poll
func Poll(buf *proto.Poll) *discordgo.Poll {
	if buf == nil {
		return nil
	}

	answers := make([]discordgo.PollAnswer, len(buf.Answers))
	for i, a := range buf.Answers {
		answers[i] = *PollAnswer(a)
	}

	return &discordgo.Poll{
		Question:         *PollMedia(buf.Question),
		Answers:          answers,
		AllowMultiselect: buf.AllowMultiselect,
		LayoutType:       discordgo.PollLayoutType(buf.LayoutType),
		Duration:         int(buf.Duration),
		Results:          PollResults(buf.Results),
		Expiry:           Timestamp(buf.Expiry),
	}
}

// PollMedia converts proto.PollMedia to discordgo.PollMedia
func PollMedia(buf *proto.PollMedia) *discordgo.PollMedia {
	if buf == nil {
		return &discordgo.PollMedia{}
	}

	return &discordgo.PollMedia{
		Text:  buf.Text,
		Emoji: ComponentEmoji(buf.Emoji),
	}
}

// PollAnswer converts proto.PollAnswer to discordgo.PollAnswer
func PollAnswer(buf *proto.PollAnswer) *discordgo.PollAnswer {
	if buf == nil {
		return &discordgo.PollAnswer{}
	}

	return &discordgo.PollAnswer{
		AnswerID: int(buf.AnswerId),
		Media:    PollMedia(buf.Media),
	}
}

// PollAnswerCount converts proto.PollAnswerCount to discordgo.PollAnswerCount
func PollAnswerCount(buf *proto.PollAnswerCount) *discordgo.PollAnswerCount {
	if buf == nil {
		return &discordgo.PollAnswerCount{}
	}

	return &discordgo.PollAnswerCount{
		ID:      int(buf.Id),
		Count:   int(buf.Count),
		MeVoted: buf.MeVoted,
	}
}

// PollResults converts proto.PollResults to discordgo.PollResults
func PollResults(buf *proto.PollResults) *discordgo.PollResults {
	if buf == nil {
		return nil
	}

	answerCounts := make([]*discordgo.PollAnswerCount, len(buf.AnswerCounts))
	for i, ac := range buf.AnswerCounts {
		answerCounts[i] = PollAnswerCount(ac)
	}

	return &discordgo.PollResults{
		Finalized:    buf.Finalized,
		AnswerCounts: answerCounts,
	}
}

// ComponentEmoji converts proto.ComponentEmoji to discordgo.ComponentEmoji
func ComponentEmoji(buf *proto.ComponentEmoji) *discordgo.ComponentEmoji {
	if buf == nil {
		return &discordgo.ComponentEmoji{}
	}

	return &discordgo.ComponentEmoji{
		Name:     buf.Name,
		ID:       buf.Id,
		Animated: buf.Animated,
	}
}

// Entitlement converts proto.Entitlement to discordgo.Entitlement
func Entitlement(buf *proto.Entitlement) *discordgo.Entitlement {
	if buf == nil {
		return nil
	}

	consumed := buf.Consumed

	return &discordgo.Entitlement{
		ID:             buf.Id,
		SKUID:          buf.SkuId,
		ApplicationID:  buf.ApplicationId,
		UserID:         buf.UserId,
		Type:           discordgo.EntitlementType(buf.Type),
		Deleted:        buf.Deleted,
		StartsAt:       Timestamp(buf.StartsAt),
		EndsAt:         Timestamp(buf.EndsAt),
		GuildID:        buf.GuildId,
		Consumed:       &consumed,
		SubscriptionID: buf.SubscriptionId,
	}
}

// EntitlementTest converts proto.EntitlementTest to discordgo.EntitlementTest
func EntitlementTest(buf *proto.EntitlementTest) *discordgo.EntitlementTest {
	if buf == nil {
		return nil
	}

	return &discordgo.EntitlementTest{
		SKUID:     buf.SkuId,
		OwnerID:   buf.OwnerId,
		OwnerType: discordgo.EntitlementOwnerType(buf.OwnerType),
	}
}

// EntitlementFilterOptions converts proto.EntitlementFilterOptions to discordgo.EntitlementFilterOptions
func EntitlementFilterOptions(buf *proto.EntitlementFilterOptions) *discordgo.EntitlementFilterOptions {
	if buf == nil {
		return nil
	}

	return &discordgo.EntitlementFilterOptions{
		UserID:       buf.UserId,
		SkuIDs:       buf.SkuIds,
		Before:       Timestamp(buf.Before),
		After:        Timestamp(buf.After),
		Limit:        int(buf.Limit),
		GuildID:      buf.GuildId,
		ExcludeEnded: buf.ExcludeEnded,
	}
}

// SKU converts proto.SKU to discordgo.SKU
func SKU(buf *proto.SKU) *discordgo.SKU {
	if buf == nil {
		return nil
	}

	return &discordgo.SKU{
		ID:            buf.Id,
		Type:          discordgo.SKUType(buf.Type),
		ApplicationID: buf.ApplicationId,
		Name:          buf.Name,
		Slug:          buf.Slug,
		Flags:         discordgo.SKUFlags(buf.Flags),
	}
}

// Subscription converts proto.Subscription to discordgo.Subscription
func Subscription(buf *proto.Subscription) *discordgo.Subscription {
	if buf == nil {
		return nil
	}

	return &discordgo.Subscription{
		ID:                 buf.Id,
		UserID:             buf.UserId,
		SKUIDs:             buf.SkuIds,
		EntitlementIDs:     buf.EntitlementIds,
		RenewalSKUIDs:      buf.RenewalSkuIds,
		CurrentPeriodStart: TimestampValue(buf.CurrentPeriodStart),
		CurrentPeriodEnd:   TimestampValue(buf.CurrentPeriodEnd),
		Status:             discordgo.SubscriptionStatus(buf.Status),
		CanceledAt:         Timestamp(buf.CanceledAt),
		Country:            buf.Country,
	}
}
