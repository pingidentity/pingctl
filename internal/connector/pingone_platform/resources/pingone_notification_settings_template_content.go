package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneNotificationSettingsTemplateContentResource{}
)

type PingoneNotificationSettingsTemplateContentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneNotificationSettingsTemplateContentResource
func NotificationSettingsTemplateContent(clientInfo *connector.SDKClientInfo) *PingoneNotificationSettingsTemplateContentResource {
	return &PingoneNotificationSettingsTemplateContentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneNotificationSettingsTemplateContentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "notification_settings_template_content",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingoneNotificationSettingsTemplateContentResource) ResourceType() string {
	return "pingone_notification_settings_template_content"
}
