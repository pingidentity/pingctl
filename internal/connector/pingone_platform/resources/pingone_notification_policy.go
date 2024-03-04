package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
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

	l.Debug().Msgf("Fetching all pingone_agreement_enable resources...")

	agreementImportBlocks, err := Agreement(r.clientInfo).ExportAll()
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_enable resources...")

	for _, importBlock := range *agreementImportBlocks {
		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType: r.ResourceType(),
			ResourceName: fmt.Sprintf("%s_enable", importBlock.ResourceName),
			ResourceID:   importBlock.ResourceID,
		})
	}

	return &importBlocks, nil
}

func (r *PingoneNotificationPolicyResource) ResourceType() string {
	return "pingone_notification_policy"
}
