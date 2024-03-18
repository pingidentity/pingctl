package platform

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneNotificationSettingsResource{}
)

type PingoneNotificationSettingsResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneNotificationSettingsResource
func NotificationSettings(clientInfo *connector.SDKClientInfo) *PingoneNotificationSettingsResource {
	return &PingoneNotificationSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneNotificationSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "notification_settings",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingoneNotificationSettingsResource) ResourceType() string {
	return "pingone_notification_settings"
}
