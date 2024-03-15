package pingoneplatformresources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	pingoneresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
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

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadAllTemplates(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllTemplates"

	embedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, template := range embedded.GetTemplates() {
		templateId, templateIdOk := template.GetIdOk()
		templateDisplayName, templateDisplayNameOk := template.GetDisplayNameOk()

		if templateIdOk && templateDisplayNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *templateDisplayName,
				ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *templateDisplayName, *templateId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneNotificationSettingsTemplateContentResource) ResourceType() string {
	return "pingone_notification_settings_template_content"
}
