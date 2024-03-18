package platform

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	pingoneresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneNotificationPolicyResource{}
)

type PingoneNotificationPolicyResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneNotificationPolicyResource
func NotificationPolicy(clientInfo *connector.SDKClientInfo) *PingoneNotificationPolicyResource {
	return &PingoneNotificationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneNotificationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsPoliciesApi.ReadAllNotificationsPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllNotificationsPolicies"

	embedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, notificationPolicy := range embedded.GetNotificationsPolicies() {
		notificationPolicyId, notificationPolicyIdOk := notificationPolicy.GetIdOk()
		notificationPolicyName, notificationPolicyNameOk := notificationPolicy.GetNameOk()

		if notificationPolicyIdOk && notificationPolicyNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *notificationPolicyName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *notificationPolicyId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneNotificationPolicyResource) ResourceType() string {
	return "pingone_notification_policy"
}
