package platform

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneBrandingThemeDefaultResource{}
)

type PingoneBrandingThemeDefaultResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneBrandingThemeDefaultResource
func BrandingThemeDefault(clientInfo *connector.SDKClientInfo) *PingoneBrandingThemeDefaultResource {
	return &PingoneBrandingThemeDefaultResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneBrandingThemeDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType: r.ResourceType(),
		ResourceName: "active_theme",
		ResourceID:   r.clientInfo.ExportEnvironmentID,
	})

	return &importBlocks, nil
}

func (r *PingoneBrandingThemeDefaultResource) ResourceType() string {
	return "pingone_branding_theme_default"
}
