package buf2struct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// AutoModerationRule converts proto.AutoModerationRule to discordgo.AutoModerationRule
func AutoModerationRule(buf *proto.AutoModerationRule) *discordgo.AutoModerationRule {
	if buf == nil {
		return nil
	}

	actions := make([]discordgo.AutoModerationAction, len(buf.Actions))
	for i, a := range buf.Actions {
		actions[i] = *AutoModerationAction(a)
	}

	enabled := buf.Enabled
	exemptRoles := buf.ExemptRoles
	exemptChannels := buf.ExemptChannels

	return &discordgo.AutoModerationRule{
		ID:              buf.Id,
		GuildID:         buf.GuildId,
		Name:            buf.Name,
		CreatorID:       buf.CreatorId,
		EventType:       discordgo.AutoModerationRuleEventType(buf.EventType),
		TriggerType:     discordgo.AutoModerationRuleTriggerType(buf.TriggerType),
		TriggerMetadata: AutoModerationTriggerMetadata(buf.TriggerMetadata),
		Actions:         actions,
		Enabled:         &enabled,
		ExemptRoles:     &exemptRoles,
		ExemptChannels:  &exemptChannels,
	}
}

// AutoModerationTriggerMetadata converts proto.AutoModerationTriggerMetadata to discordgo.AutoModerationTriggerMetadata
func AutoModerationTriggerMetadata(buf *proto.AutoModerationTriggerMetadata) *discordgo.AutoModerationTriggerMetadata {
	if buf == nil {
		return nil
	}

	presets := make([]discordgo.AutoModerationKeywordPreset, len(buf.Presets))
	for i, p := range buf.Presets {
		presets[i] = discordgo.AutoModerationKeywordPreset(p)
	}

	var allowList *[]string
	if len(buf.AllowList) > 0 {
		allowList = &buf.AllowList
	}

	return &discordgo.AutoModerationTriggerMetadata{
		KeywordFilter:     buf.KeywordFilter,
		RegexPatterns:     buf.RegexPatterns,
		Presets:           presets,
		AllowList:         allowList,
		MentionTotalLimit: int(buf.MentionTotalLimit),
	}
}

// AutoModerationAction converts proto.AutoModerationAction to discordgo.AutoModerationAction
func AutoModerationAction(buf *proto.AutoModerationAction) *discordgo.AutoModerationAction {
	if buf == nil {
		return &discordgo.AutoModerationAction{}
	}

	return &discordgo.AutoModerationAction{
		Type:     discordgo.AutoModerationActionType(buf.Type),
		Metadata: AutoModerationActionMetadata(buf.Metadata),
	}
}

// AutoModerationActionMetadata converts proto.AutoModerationActionMetadata to discordgo.AutoModerationActionMetadata
func AutoModerationActionMetadata(buf *proto.AutoModerationActionMetadata) *discordgo.AutoModerationActionMetadata {
	if buf == nil {
		return &discordgo.AutoModerationActionMetadata{}
	}

	return &discordgo.AutoModerationActionMetadata{
		ChannelID:     buf.ChannelId,
		Duration:      int(buf.Duration),
		CustomMessage: buf.CustomMessage,
	}
}

// GuildAuditLog converts proto.GuildAuditLog to discordgo.GuildAuditLog
func GuildAuditLog(buf *proto.GuildAuditLog) *discordgo.GuildAuditLog {
	if buf == nil {
		return nil
	}

	webhooks := make([]*discordgo.Webhook, len(buf.Webhooks))
	for i, w := range buf.Webhooks {
		webhooks[i] = Webhook(w)
	}

	users := make([]*discordgo.User, len(buf.Users))
	for i, u := range buf.Users {
		users[i] = User(u)
	}

	entries := make([]*discordgo.AuditLogEntry, len(buf.AuditLogEntries))
	for i, e := range buf.AuditLogEntries {
		entries[i] = AuditLogEntry(e)
	}

	integrations := make([]*discordgo.Integration, len(buf.Integrations))
	for i, integ := range buf.Integrations {
		integrations[i] = Integration(integ)
	}

	return &discordgo.GuildAuditLog{
		Webhooks:         webhooks,
		Users:            users,
		AuditLogEntries:  entries,
		Integrations:     integrations,
	}
}

// AuditLogEntry converts proto.AuditLogEntry to discordgo.AuditLogEntry
func AuditLogEntry(buf *proto.AuditLogEntry) *discordgo.AuditLogEntry {
	if buf == nil {
		return nil
	}

	changes := make([]*discordgo.AuditLogChange, len(buf.Changes))
	for i, c := range buf.Changes {
		changes[i] = AuditLogChange(c)
	}

	actionType := discordgo.AuditLogAction(buf.ActionType)

	return &discordgo.AuditLogEntry{
		TargetID:   buf.TargetId,
		Changes:    changes,
		UserID:     buf.UserId,
		ID:         buf.Id,
		ActionType: &actionType,
		Options:    AuditLogOptions(buf.Options),
		Reason:     buf.Reason,
	}
}

// AuditLogChange converts proto.AuditLogChange to discordgo.AuditLogChange
func AuditLogChange(buf *proto.AuditLogChange) *discordgo.AuditLogChange {
	if buf == nil {
		return nil
	}

	// Convert bytes to interface{} - simplified version
	// In real implementation, you'd need proper deserialization
	var newValue interface{}
	var oldValue interface{}

	key := discordgo.AuditLogChangeKey(buf.Key)

	return &discordgo.AuditLogChange{
		NewValue: newValue,
		OldValue: oldValue,
		Key:      &key,
	}
}

// AuditLogOptions converts proto.AuditLogOptions to discordgo.AuditLogOptions
func AuditLogOptions(buf *proto.AuditLogOptions) *discordgo.AuditLogOptions {
	if buf == nil {
		return nil
	}

	typeVal := discordgo.AuditLogOptionsType(buf.Type)

	return &discordgo.AuditLogOptions{
		DeleteMemberDays:              buf.DeleteMemberDays,
		MembersRemoved:                buf.MembersRemoved,
		ChannelID:                     buf.ChannelId,
		MessageID:                     buf.MessageId,
		Count:                         buf.Count,
		ID:                            buf.Id,
		Type:                          &typeVal,
		RoleName:                      buf.RoleName,
		ApplicationID:                 buf.ApplicationId,
		AutoModerationRuleName:        buf.AutoModerationRuleName,
		AutoModerationRuleTriggerType: buf.AutoModerationRuleTriggerType,
		IntegrationType:               buf.IntegrationType,
	}
}
