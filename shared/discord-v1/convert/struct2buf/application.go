package struct2buf

import (
	"github.com/bwmarrin/discordgo"
	proto "github.com/thirdscam/chatanium-flexmodule/proto/discord-v1"
)

// Application converts discordgo.Application to proto.Application
func Application(s *discordgo.Application) *proto.Application {
	if s == nil {
		return nil
	}

	integrationTypesConfig := make(map[int32]*proto.ApplicationIntegrationTypeConfig)
	for k, v := range s.IntegrationTypesConfig {
		integrationTypesConfig[int32(k)] = ApplicationIntegrationTypeConfig(v)
	}

	return &proto.Application{
		Id:                     s.ID,
		Name:                   s.Name,
		Icon:                   s.Icon,
		Description:            s.Description,
		RpcOrigins:             s.RPCOrigins,
		BotPublic:              s.BotPublic,
		BotRequireCodeGrant:    s.BotRequireCodeGrant,
		TermsOfServiceUrl:      s.TermsOfServiceURL,
		PrivacyPolicyUrl:       s.PrivacyProxyURL,
		Owner:                  User(s.Owner),
		Summary:                s.Summary,
		VerifyKey:              s.VerifyKey,
		Team:                   Team(s.Team),
		GuildId:                s.GuildID,
		PrimarySkuId:           s.PrimarySKUID,
		Slug:                   s.Slug,
		CoverImage:             s.CoverImage,
		Flags:                  int32(s.Flags),
		IntegrationTypesConfig: integrationTypesConfig,
	}
}

// ApplicationIntegrationTypeConfig converts discordgo.ApplicationIntegrationTypeConfig to proto.ApplicationIntegrationTypeConfig
func ApplicationIntegrationTypeConfig(s *discordgo.ApplicationIntegrationTypeConfig) *proto.ApplicationIntegrationTypeConfig {
	if s == nil {
		return nil
	}

	return &proto.ApplicationIntegrationTypeConfig{
		Oauth2InstallParams: ApplicationInstallParams(s.OAuth2InstallParams),
	}
}

// ApplicationInstallParams converts discordgo.ApplicationInstallParams to proto.ApplicationInstallParams
func ApplicationInstallParams(s *discordgo.ApplicationInstallParams) *proto.ApplicationInstallParams {
	if s == nil {
		return nil
	}

	return &proto.ApplicationInstallParams{
		Scopes:      s.Scopes,
		Permissions: int64(s.Permissions),
	}
}

// ApplicationRoleConnectionMetadata converts discordgo.ApplicationRoleConnectionMetadata to proto.ApplicationRoleConnectionMetadata
func ApplicationRoleConnectionMetadata(s *discordgo.ApplicationRoleConnectionMetadata) *proto.ApplicationRoleConnectionMetadata {
	if s == nil {
		return nil
	}

	nameLocalizations := make(map[string]string)
	for k, v := range s.NameLocalizations {
		nameLocalizations[string(k)] = v
	}

	descriptionLocalizations := make(map[string]string)
	for k, v := range s.DescriptionLocalizations {
		descriptionLocalizations[string(k)] = v
	}

	return &proto.ApplicationRoleConnectionMetadata{
		Type:                     int32(s.Type),
		Key:                      s.Key,
		Name:                     s.Name,
		NameLocalizations:        nameLocalizations,
		Description:              s.Description,
		DescriptionLocalizations: descriptionLocalizations,
	}
}

// ApplicationRoleConnection converts discordgo.ApplicationRoleConnection to proto.ApplicationRoleConnection
func ApplicationRoleConnection(s *discordgo.ApplicationRoleConnection) *proto.ApplicationRoleConnection {
	if s == nil {
		return nil
	}

	return &proto.ApplicationRoleConnection{
		PlatformName:     s.PlatformName,
		PlatformUsername: s.PlatformUsername,
		Metadata:         s.Metadata,
	}
}

// Team converts discordgo.Team to proto.Team
func Team(s *discordgo.Team) *proto.Team {
	if s == nil {
		return nil
	}

	members := make([]*proto.TeamMember, len(s.Members))
	for i, m := range s.Members {
		members[i] = TeamMember(m)
	}

	return &proto.Team{
		Id:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Icon:        s.Icon,
		OwnerUserId: s.OwnerID,
		Members:     members,
	}
}

// TeamMember converts discordgo.TeamMember to proto.TeamMember
func TeamMember(s *discordgo.TeamMember) *proto.TeamMember {
	if s == nil {
		return nil
	}

	return &proto.TeamMember{
		User:            User(s.User),
		TeamId:          s.TeamID,
		MembershipState: proto.MembershipState(s.MembershipState),
		Permissions:     s.Permissions,
	}
}

// UserConnection converts discordgo.UserConnection to proto.UserConnection
func UserConnection(s *discordgo.UserConnection) *proto.UserConnection {
	if s == nil {
		return nil
	}

	integrations := make([]*proto.Integration, len(s.Integrations))
	for i, integ := range s.Integrations {
		integrations[i] = Integration(integ)
	}

	return &proto.UserConnection{
		Id:           s.ID,
		Name:         s.Name,
		Type:         s.Type,
		Revoked:      s.Revoked,
		Integrations: integrations,
	}
}

// Integration converts discordgo.Integration to proto.Integration
func Integration(s *discordgo.Integration) *proto.Integration {
	if s == nil {
		return nil
	}

	return &proto.Integration{
		Id:                s.ID,
		Name:              s.Name,
		Type:              s.Type,
		Enabled:           s.Enabled,
		Syncing:           s.Syncing,
		RoleId:            s.RoleID,
		EnableEmoticons:   s.EnableEmoticons,
		ExpireBehavior:    int32(s.ExpireBehavior),
		ExpireGracePeriod: int32(s.ExpireGracePeriod),
		User:              User(s.User),
		Account:           IntegrationAccount(&s.Account),
		SyncedAt:          Timestamp(s.SyncedAt),
	}
}

// IntegrationAccount converts discordgo.IntegrationAccount to proto.IntegrationAccount
func IntegrationAccount(s *discordgo.IntegrationAccount) *proto.IntegrationAccount {
	if s == nil {
		return nil
	}

	return &proto.IntegrationAccount{
		Id:   s.ID,
		Name: s.Name,
	}
}
