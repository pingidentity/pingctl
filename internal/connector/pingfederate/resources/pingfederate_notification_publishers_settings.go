package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateNotificationPublishersSettingsResource{}
)

type PingFederateNotificationPublishersSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateNotificationPublishersSettingsResource
func NotificationPublishersSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateNotificationPublishersSettingsResource {
	return &PingFederateNotificationPublishersSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateNotificationPublishersSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	notificationPublishersSettingsId := "notification_publishers_settings_singleton_id"
	notificationPublishersSettingsName := "Notification Publishers Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Singleton ID":  common.SINGLETON_ID_COMMENT_DATA,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       notificationPublishersSettingsName,
		ResourceID:         notificationPublishersSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateNotificationPublishersSettingsResource) ResourceType() string {
	return "pingfederate_notification_publishers_settings"
}
