package pingoneplatformresources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneNotificationSettingsEmailResource{}
)

type PingoneNotificationSettingsEmailResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneNotificationSettingsEmailResource
func NotificationSettingsEmail(clientInfo *connector.SDKClientInfo) *PingoneNotificationSettingsEmailResource {
	return &PingoneNotificationSettingsEmailResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneNotificationSettingsEmailResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "pingone_notification_settings_email",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingoneNotificationSettingsEmailResource) ResourceType() string {
	return "pingone_notification_settings_email"
}
