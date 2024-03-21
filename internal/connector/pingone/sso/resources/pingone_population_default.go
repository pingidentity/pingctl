package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingonePopulationDefaultDefaultResource{}
)

type PingonePopulationDefaultDefaultResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingonePopulationDefaultDefaultResource
func PopulationDefault(clientInfo *connector.SDKClientInfo) *PingonePopulationDefaultDefaultResource {
	return &PingonePopulationDefaultDefaultResource{
		clientInfo: clientInfo,
	}
}

func (r *PingonePopulationDefaultDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "population_default",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingonePopulationDefaultDefaultResource) ResourceType() string {
	return "pingone_population_default"
}
