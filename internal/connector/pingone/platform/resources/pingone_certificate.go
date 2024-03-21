package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneCertificateResource{}
)

type PingoneCertificateResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneCertificateResource
func Certificate(clientInfo *connector.SDKClientInfo) *PingoneCertificateResource {
	return &PingoneCertificateResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneCertificateResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetCertificates(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "GetCertificates"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, certificate := range embedded.GetCertificates() {
		certificateName, certificateNameOk := certificate.GetNameOk()
		certificateId, certificateIdOk := certificate.GetIdOk()

		if certificateNameOk && certificateIdOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *certificateName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *certificateId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneCertificateResource) ResourceType() string {
	return "pingone_certificate"
}
