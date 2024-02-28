package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneFormRecaptchaV2Resource{}
)

type PingoneFormRecaptchaV2Resource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneFormRecaptchaV2Resource
func FormRecaptchaV2(clientInfo *connector.SDKClientInfo) *PingoneFormRecaptchaV2Resource {
	return &PingoneFormRecaptchaV2Resource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneFormRecaptchaV2Resource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_forms_recaptcha_v2 resources...")

	importBlocks := []connector.ImportBlock{}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "recaptcha_configuration",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingoneFormRecaptchaV2Resource) ResourceType() string {
	return "pingone_forms_recaptcha_v2"
}
