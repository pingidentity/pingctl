package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneFormResource{}
)

type PingoneFormResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneFormResource
func Form(clientInfo *connector.SDKClientInfo) *PingoneFormResource {
	return &PingoneFormResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneFormResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_form resources...")

	entityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.FormManagementApi.ReadAllForms(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadAllForms Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned ReadAllForms() entity array is nil.")
		l.Error().Msgf("ReadAllForms Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_form resources via ReadAllForms()")
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned ReadAllForms() embedded data is nil.")
		l.Error().Msgf("ReadAllForms Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_form resources via ReadAllForms()")
	}

	importBlocks := []connector.ImportBlock{}

	for _, form := range embedded.GetForms() {
		formId, formIdOk := form.GetIdOk()
		formName, formNameOk := form.GetNameOk()

		if formIdOk && formNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *formName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *formId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneFormResource) ResourceType() string {
	return "pingone_form"
}
