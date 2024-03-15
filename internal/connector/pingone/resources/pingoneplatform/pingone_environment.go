package pingoneplatformresources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneCustomDomainResource{}
)

type PingoneEnvironmentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneEnvironmentResource
func Environment(clientInfo *connector.SDKClientInfo) *PingoneEnvironmentResource {
	return &PingoneEnvironmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneEnvironmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "export_environment",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingoneEnvironmentResource) ResourceType() string {
	return "pingone_environment"
}
