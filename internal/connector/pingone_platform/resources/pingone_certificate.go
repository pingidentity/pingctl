package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
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

	l.Debug().Msgf("Fetching all pingone_certificate resources...")

	entityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.CertificateManagementApi.GetCertificates(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("GetCertificates Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned GetCertificates() entity array is nil.")
		l.Error().Msgf("GetCertificates Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_certificate resources via GetCertificates()")
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned GetCertificates() embedded data is nil.")
		l.Error().Msgf("GetCertificates Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_certificate resources via GetCertificates()")
	}

	importBlocks := []connector.ImportBlock{}

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
