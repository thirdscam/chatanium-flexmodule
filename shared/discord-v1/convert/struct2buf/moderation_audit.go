package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// AutoModerationRule converts discordgo.AutoModerationRule to proto.AutoModerationRule
func AutoModerationRule(s *discordgo.AutoModerationRule) *proto.AutoModerationRule {
	if s == nil {
		return nil
	}

	actions := make([]*proto.AutoModerationAction, len(s.Actions))
	for i, a := range s.Actions {
		actions[i] = AutoModerationAction(&a)
	}

	var enabled bool
	if s.Enabled != nil {
		enabled = *s.Enabled
	}

	var exemptRoles []string
	if s.ExemptRoles != nil {
		exemptRoles = *s.ExemptRoles
	}

	var exemptChannels []string
	if s.ExemptChannels != nil {
		exemptChannels = *s.ExemptChannels
	}

	return &proto.AutoModerationRule{
		Id:              s.ID,
		GuildId:         s.GuildID,
		Name:            s.Name,
		CreatorId:       s.CreatorID,
		EventType:       int32(s.EventType),
		TriggerType:     int32(s.TriggerType),
		TriggerMetadata: AutoModerationTriggerMetadata(s.TriggerMetadata),
		Actions:         actions,
		Enabled:         enabled,
		ExemptRoles:     exemptRoles,
		ExemptChannels:  exemptChannels,
	}
}

// AutoModerationTriggerMetadata converts discordgo.AutoModerationTriggerMetadata to proto.AutoModerationTriggerMetadata
func AutoModerationTriggerMetadata(s *discordgo.AutoModerationTriggerMetadata) *proto.AutoModerationTriggerMetadata {
	if s == nil {
		return nil
	}

	presets := make([]uint32, len(s.Presets))
	for i, p := range s.Presets {
		presets[i] = uint32(p)
	}

	var allowList []string
	if s.AllowList != nil {
		allowList = *s.AllowList
	}

	return &proto.AutoModerationTriggerMetadata{
		KeywordFilter:      s.KeywordFilter,
		RegexPatterns:      s.RegexPatterns,
		Presets:            presets,
		AllowList:          allowList,
		MentionTotalLimit:  int32(s.MentionTotalLimit),
	}
}

// AutoModerationAction converts discordgo.AutoModerationAction to proto.AutoModerationAction
func AutoModerationAction(s *discordgo.AutoModerationAction) *proto.AutoModerationAction {
	if s == nil {
		return nil
	}

	return &proto.AutoModerationAction{
		Type:     int32(s.Type),
		Metadata: AutoModerationActionMetadata(s.Metadata),
	}
}

// AutoModerationActionMetadata converts discordgo.AutoModerationActionMetadata to proto.AutoModerationActionMetadata
func AutoModerationActionMetadata(s *discordgo.AutoModerationActionMetadata) *proto.AutoModerationActionMetadata {
	if s == nil {
		return nil
	}

	return &proto.AutoModerationActionMetadata{
		ChannelId:     s.ChannelID,
		Duration:      int32(s.Duration),
		CustomMessage: s.CustomMessage,
	}
}

// GuildAuditLog converts discordgo.GuildAuditLog to proto.GuildAuditLog
func GuildAuditLog(s *discordgo.GuildAuditLog) *proto.GuildAuditLog {
	if s == nil {
		return nil
	}

	webhooks := make([]*proto.Webhook, len(s.Webhooks))
	for i, w := range s.Webhooks {
		webhooks[i] = Webhook(w)
	}

	users := make([]*proto.User, len(s.Users))
	for i, u := range s.Users {
		users[i] = User(u)
	}

	entries := make([]*proto.AuditLogEntry, len(s.AuditLogEntries))
	for i, e := range s.AuditLogEntries {
		entries[i] = AuditLogEntry(e)
	}

	integrations := make([]*proto.Integration, len(s.Integrations))
	for i, integ := range s.Integrations {
		integrations[i] = Integration(integ)
	}

	return &proto.GuildAuditLog{
		Webhooks:         webhooks,
		Users:            users,
		AuditLogEntries:  entries,
		Integrations:     integrations,
	}
}

// AuditLogEntry converts discordgo.AuditLogEntry to proto.AuditLogEntry
func AuditLogEntry(s *discordgo.AuditLogEntry) *proto.AuditLogEntry {
	if s == nil {
		return nil
	}

	changes := make([]*proto.AuditLogChange, len(s.Changes))
	for i, c := range s.Changes {
		changes[i] = AuditLogChange(c)
	}

	var actionType int32
	if s.ActionType != nil {
		actionType = int32(*s.ActionType)
	}

	return &proto.AuditLogEntry{
		TargetId:   s.TargetID,
		Changes:    changes,
		UserId:     s.UserID,
		Id:         s.ID,
		ActionType: actionType,
		Options:    AuditLogOptions(s.Options),
		Reason:     s.Reason,
	}
}

// AuditLogChange converts discordgo.AuditLogChange to proto.AuditLogChange
func AuditLogChange(s *discordgo.AuditLogChange) *proto.AuditLogChange {
	if s == nil {
		return nil
	}

	// Convert interface{} to bytes - simplified version
	var newValue []byte
	var oldValue []byte

	// Note: In a real implementation, you'd need proper serialization
	// This is a placeholder that would need actual JSON marshaling

	var key string
	if s.Key != nil {
		key = string(*s.Key)
	}

	return &proto.AuditLogChange{
		NewValue: newValue,
		OldValue: oldValue,
		Key:      key,
	}
}

// AuditLogOptions converts discordgo.AuditLogOptions to proto.AuditLogOptions
func AuditLogOptions(s *discordgo.AuditLogOptions) *proto.AuditLogOptions {
	if s == nil {
		return nil
	}

	var typeStr string
	if s.Type != nil {
		typeStr = string(*s.Type)
	}

	return &proto.AuditLogOptions{
		DeleteMemberDays:              s.DeleteMemberDays,
		MembersRemoved:                s.MembersRemoved,
		ChannelId:                     s.ChannelID,
		MessageId:                     s.MessageID,
		Count:                         s.Count,
		Id:                            s.ID,
		Type:                          typeStr,
		RoleName:                      s.RoleName,
		ApplicationId:                 s.ApplicationID,
		AutoModerationRuleName:        s.AutoModerationRuleName,
		AutoModerationRuleTriggerType: s.AutoModerationRuleTriggerType,
		IntegrationType:               s.IntegrationType,
	}
}
