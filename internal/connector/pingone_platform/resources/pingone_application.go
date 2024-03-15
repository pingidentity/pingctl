package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneCustomDomainResource{}
)

type PingoneApplicationResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneApplicationResource
func Application(clientInfo *connector.SDKClientInfo) *PingoneApplicationResource {
	return &PingoneApplicationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "export_application",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingoneApplicationResource) ResourceType() string {
	return "pingone_application"
}
