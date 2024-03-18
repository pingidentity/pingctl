package pingonessoresources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingonePopulationResource{}
)

type PingonePopulationResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingonePopulationResource
func Population(clientInfo *connector.SDKClientInfo) *PingonePopulationResource {
	return &PingonePopulationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingonePopulationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.PopulationsApi.ReadAllPopulations(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllPopulations"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, population := range embedded.GetPopulations() {
		populationId, populationIdOk := population.GetIdOk()
		populationName, populationNameOk := population.GetNameOk()

		if populationIdOk && populationNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *populationName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *populationId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingonePopulationResource) ResourceType() string {
	return "pingone_population"
}
