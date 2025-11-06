package buf2struct

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// Application converts proto.Application to discordgo.Application
func Application(buf *proto.Application) *discordgo.Application {
	if buf == nil {
		return nil
	}

	integrationTypesConfig := make(map[discordgo.ApplicationIntegrationType]*discordgo.ApplicationIntegrationTypeConfig)
	for k, v := range buf.IntegrationTypesConfig {
		integrationTypesConfig[discordgo.ApplicationIntegrationType(k)] = ApplicationIntegrationTypeConfig(v)
	}

	return &discordgo.Application{
		ID:                     buf.Id,
		Name:                   buf.Name,
		Icon:                   buf.Icon,
		Description:            buf.Description,
		RPCOrigins:             buf.RpcOrigins,
		BotPublic:              buf.BotPublic,
		BotRequireCodeGrant:    buf.BotRequireCodeGrant,
		TermsOfServiceURL:      buf.TermsOfServiceUrl,
		PrivacyProxyURL:        buf.PrivacyPolicyUrl,
		Owner:                  User(buf.Owner),
		Summary:                buf.Summary,
		VerifyKey:              buf.VerifyKey,
		Team:                   Team(buf.Team),
		GuildID:                buf.GuildId,
		PrimarySKUID:           buf.PrimarySkuId,
		Slug:                   buf.Slug,
		CoverImage:             buf.CoverImage,
		Flags:                  int(buf.Flags),
		IntegrationTypesConfig: integrationTypesConfig,
	}
}

// ApplicationIntegrationTypeConfig converts proto.ApplicationIntegrationTypeConfig to discordgo.ApplicationIntegrationTypeConfig
func ApplicationIntegrationTypeConfig(buf *proto.ApplicationIntegrationTypeConfig) *discordgo.ApplicationIntegrationTypeConfig {
	if buf == nil {
		return nil
	}

	return &discordgo.ApplicationIntegrationTypeConfig{
		OAuth2InstallParams: ApplicationInstallParams(buf.Oauth2InstallParams),
	}
}

// ApplicationInstallParams converts proto.ApplicationInstallParams to discordgo.ApplicationInstallParams
func ApplicationInstallParams(buf *proto.ApplicationInstallParams) *discordgo.ApplicationInstallParams {
	if buf == nil {
		return nil
	}

	return &discordgo.ApplicationInstallParams{
		Scopes:      buf.Scopes,
		Permissions: int64(buf.Permissions),
	}
}

// ApplicationRoleConnectionMetadata converts proto.ApplicationRoleConnectionMetadata to discordgo.ApplicationRoleConnectionMetadata
func ApplicationRoleConnectionMetadata(buf *proto.ApplicationRoleConnectionMetadata) *discordgo.ApplicationRoleConnectionMetadata {
	if buf == nil {
		return nil
	}

	nameLocalizations := make(map[discordgo.Locale]string)
	for k, v := range buf.NameLocalizations {
		nameLocalizations[discordgo.Locale(k)] = v
	}

	descriptionLocalizations := make(map[discordgo.Locale]string)
	for k, v := range buf.DescriptionLocalizations {
		descriptionLocalizations[discordgo.Locale(k)] = v
	}

	return &discordgo.ApplicationRoleConnectionMetadata{
		Type:                     discordgo.ApplicationRoleConnectionMetadataType(buf.Type),
		Key:                      buf.Key,
		Name:                     buf.Name,
		NameLocalizations:        nameLocalizations,
		Description:              buf.Description,
		DescriptionLocalizations: descriptionLocalizations,
	}
}

// ApplicationRoleConnection converts proto.ApplicationRoleConnection to discordgo.ApplicationRoleConnection
func ApplicationRoleConnection(buf *proto.ApplicationRoleConnection) *discordgo.ApplicationRoleConnection {
	if buf == nil {
		return nil
	}

	return &discordgo.ApplicationRoleConnection{
		PlatformName:     buf.PlatformName,
		PlatformUsername: buf.PlatformUsername,
		Metadata:         buf.Metadata,
	}
}

// Team converts proto.Team to discordgo.Team
func Team(buf *proto.Team) *discordgo.Team {
	if buf == nil {
		return nil
	}

	members := make([]*discordgo.TeamMember, len(buf.Members))
	for i, m := range buf.Members {
		members[i] = TeamMember(m)
	}

	return &discordgo.Team{
		ID:          buf.Id,
		Name:        buf.Name,
		Description: buf.Description,
		Icon:        buf.Icon,
		OwnerID:     buf.OwnerUserId,
		Members:     members,
	}
}

// TeamMember converts proto.TeamMember to discordgo.TeamMember
func TeamMember(buf *proto.TeamMember) *discordgo.TeamMember {
	if buf == nil {
		return nil
	}

	return &discordgo.TeamMember{
		User:            User(buf.User),
		TeamID:          buf.TeamId,
		MembershipState: discordgo.MembershipState(buf.MembershipState),
		Permissions:     buf.Permissions,
	}
}

// UserConnection converts proto.UserConnection to discordgo.UserConnection
func UserConnection(buf *proto.UserConnection) *discordgo.UserConnection {
	if buf == nil {
		return nil
	}

	integrations := make([]*discordgo.Integration, len(buf.Integrations))
	for i, integ := range buf.Integrations {
		integrations[i] = Integration(integ)
	}

	return &discordgo.UserConnection{
		ID:           buf.Id,
		Name:         buf.Name,
		Type:         buf.Type,
		Revoked:      buf.Revoked,
		Integrations: integrations,
	}
}

// Integration converts proto.Integration to discordgo.Integration
func Integration(buf *proto.Integration) *discordgo.Integration {
	if buf == nil {
		return nil
	}

	return &discordgo.Integration{
		ID:                buf.Id,
		Name:              buf.Name,
		Type:              buf.Type,
		Enabled:           buf.Enabled,
		Syncing:           buf.Syncing,
		RoleID:            buf.RoleId,
		EnableEmoticons:   buf.EnableEmoticons,
		ExpireBehavior:    discordgo.ExpireBehavior(buf.ExpireBehavior),
		ExpireGracePeriod: int(buf.ExpireGracePeriod),
		User:              User(buf.User),
		Account:           *IntegrationAccount(buf.Account),
		SyncedAt:          TimestampValue(buf.SyncedAt),
	}
}

// IntegrationAccount converts proto.IntegrationAccount to discordgo.IntegrationAccount
func IntegrationAccount(buf *proto.IntegrationAccount) *discordgo.IntegrationAccount {
	if buf == nil {
		return &discordgo.IntegrationAccount{}
	}

	return &discordgo.IntegrationAccount{
		ID:   buf.Id,
		Name: buf.Name,
	}
}
