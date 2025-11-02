package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// Poll converts discordgo.Poll to proto.Poll
func Poll(s *discordgo.Poll) *proto.Poll {
	if s == nil {
		return nil
	}

	answers := make([]*proto.PollAnswer, len(s.Answers))
	for i, a := range s.Answers {
		answers[i] = PollAnswer(&a)
	}

	return &proto.Poll{
		Question:         PollMedia(&s.Question),
		Answers:          answers,
		AllowMultiselect: s.AllowMultiselect,
		LayoutType:       int32(s.LayoutType),
		Duration:         int32(s.Duration),
		Results:          PollResults(s.Results),
		Expiry:           TimestampPtr(s.Expiry),
	}
}

// PollMedia converts discordgo.PollMedia to proto.PollMedia
func PollMedia(s *discordgo.PollMedia) *proto.PollMedia {
	if s == nil {
		return nil
	}

	return &proto.PollMedia{
		Text:  s.Text,
		Emoji: ComponentEmoji(s.Emoji),
	}
}

// PollAnswer converts discordgo.PollAnswer to proto.PollAnswer
func PollAnswer(s *discordgo.PollAnswer) *proto.PollAnswer {
	if s == nil {
		return nil
	}

	return &proto.PollAnswer{
		AnswerId: int32(s.AnswerID),
		Media:    PollMedia(s.Media),
	}
}

// PollAnswerCount converts discordgo.PollAnswerCount to proto.PollAnswerCount
func PollAnswerCount(s *discordgo.PollAnswerCount) *proto.PollAnswerCount {
	if s == nil {
		return nil
	}

	return &proto.PollAnswerCount{
		Id:      int32(s.ID),
		Count:   int32(s.Count),
		MeVoted: s.MeVoted,
	}
}

// PollResults converts discordgo.PollResults to proto.PollResults
func PollResults(s *discordgo.PollResults) *proto.PollResults {
	if s == nil {
		return nil
	}

	answerCounts := make([]*proto.PollAnswerCount, len(s.AnswerCounts))
	for i, ac := range s.AnswerCounts {
		answerCounts[i] = PollAnswerCount(ac)
	}

	return &proto.PollResults{
		Finalized:    s.Finalized,
		AnswerCounts: answerCounts,
	}
}

// ComponentEmoji converts discordgo.ComponentEmoji to proto.ComponentEmoji
func ComponentEmoji(s *discordgo.ComponentEmoji) *proto.ComponentEmoji {
	if s == nil {
		return nil
	}

	return &proto.ComponentEmoji{
		Name:     s.Name,
		Id:       s.ID,
		Animated: s.Animated,
	}
}

// Entitlement converts discordgo.Entitlement to proto.Entitlement
func Entitlement(s *discordgo.Entitlement) *proto.Entitlement {
	if s == nil {
		return nil
	}

	var consumed bool
	if s.Consumed != nil {
		consumed = *s.Consumed
	}

	return &proto.Entitlement{
		Id:             s.ID,
		SkuId:          s.SKUID,
		ApplicationId:  s.ApplicationID,
		UserId:         s.UserID,
		Type:           int32(s.Type),
		Deleted:        s.Deleted,
		StartsAt:       TimestampPtr(s.StartsAt),
		EndsAt:         TimestampPtr(s.EndsAt),
		GuildId:        s.GuildID,
		Consumed:       consumed,
		SubscriptionId: s.SubscriptionID,
	}
}

// EntitlementTest converts discordgo.EntitlementTest to proto.EntitlementTest
func EntitlementTest(s *discordgo.EntitlementTest) *proto.EntitlementTest {
	if s == nil {
		return nil
	}

	return &proto.EntitlementTest{
		SkuId:     s.SKUID,
		OwnerId:   s.OwnerID,
		OwnerType: int32(s.OwnerType),
	}
}

// EntitlementFilterOptions converts discordgo.EntitlementFilterOptions to proto.EntitlementFilterOptions
func EntitlementFilterOptions(s *discordgo.EntitlementFilterOptions) *proto.EntitlementFilterOptions {
	if s == nil {
		return nil
	}

	return &proto.EntitlementFilterOptions{
		UserId:        s.UserID,
		SkuIds:        s.SkuIDs,
		Before:        TimestampPtr(s.Before),
		After:         TimestampPtr(s.After),
		Limit:         int32(s.Limit),
		GuildId:       s.GuildID,
		ExcludeEnded:  s.ExcludeEnded,
	}
}

// SKU converts discordgo.SKU to proto.SKU
func SKU(s *discordgo.SKU) *proto.SKU {
	if s == nil {
		return nil
	}

	return &proto.SKU{
		Id:            s.ID,
		Type:          int32(s.Type),
		ApplicationId: s.ApplicationID,
		Name:          s.Name,
		Slug:          s.Slug,
		Flags:         int32(s.Flags),
	}
}

// Subscription converts discordgo.Subscription to proto.Subscription
func Subscription(s *discordgo.Subscription) *proto.Subscription {
	if s == nil {
		return nil
	}

	return &proto.Subscription{
		Id:                 s.ID,
		UserId:             s.UserID,
		SkuIds:             s.SKUIDs,
		EntitlementIds:     s.EntitlementIDs,
		RenewalSkuIds:      s.RenewalSKUIDs,
		CurrentPeriodStart: Timestamp(s.CurrentPeriodStart),
		CurrentPeriodEnd:   Timestamp(s.CurrentPeriodEnd),
		Status:             int32(s.Status),
		CanceledAt:         TimestampPtr(s.CanceledAt),
		Country:            s.Country,
	}
}
