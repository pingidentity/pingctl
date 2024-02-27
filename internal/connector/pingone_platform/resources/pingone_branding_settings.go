package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneBrandingSettingsResource{}
)

type PingoneBrandingSettingsResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneBrandingSettingsResource
func BrandingSettings(clientInfo *connector.SDKClientInfo) *PingoneBrandingSettingsResource {
	return &PingoneBrandingSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneBrandingSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_branding_settings resources...")

	importBlocks := []connector.ImportBlock{}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "branding",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingoneBrandingSettingsResource) ResourceType() string {
	return "pingone_branding_settings"
}
