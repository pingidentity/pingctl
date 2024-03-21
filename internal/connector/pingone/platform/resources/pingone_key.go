package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneKeyResource{}
)

type PingoneKeyResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneKeyResource
func Key(clientInfo *connector.SDKClientInfo) *PingoneKeyResource {
	return &PingoneKeyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneKeyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetKeys(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "GetKeys"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, key := range embedded.GetKeys() {
		keyId, keyIdOk := key.GetIdOk()
		keyName, keyNameOk := key.GetNameOk()

		if keyIdOk && keyNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *keyName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *keyId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneKeyResource) ResourceType() string {
	return "pingone_key"
}
