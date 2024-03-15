package pingoneplatformresources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	pingoneresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneWebhookResource{}
)

type PingoneWebhookResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneWebhookResource
func Webhook(clientInfo *connector.SDKClientInfo) *PingoneWebhookResource {
	return &PingoneWebhookResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneWebhookResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.SubscriptionsWebhooksApi.ReadAllSubscriptions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllSubscriptions"

	usersEmbedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, subscription := range usersEmbedded.GetSubscriptions() {
		subscriptionId, subscriptionIdOk := subscription.GetIdOk()
		subscriptionName, subscriptionNameOk := subscription.GetNameOk()

		if subscriptionIdOk && subscriptionNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *subscriptionName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *subscriptionId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneWebhookResource) ResourceType() string {
	return "pingone_webhook"
}
