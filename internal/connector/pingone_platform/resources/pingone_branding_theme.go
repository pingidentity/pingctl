package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneBrandingThemeResource{}
)

type PingoneBrandingThemeResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneBrandingThemeResource
func BrandingTheme(clientInfo *connector.SDKClientInfo) *PingoneBrandingThemeResource {
	return &PingoneBrandingThemeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneBrandingThemeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadBrandingThemes"

	embedded, err := GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, theme := range embedded.GetThemes() {
		themeId, themeIdOk := theme.GetIdOk()
		themeConfiguration, themeConfigurationOk := theme.GetConfigurationOk()
		var themeName *string
		var themeNameOk = false
		if themeConfigurationOk {
			themeName, themeNameOk = themeConfiguration.GetNameOk()
		}

		if themeIdOk && themeConfigurationOk && themeNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *themeName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *themeId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneBrandingThemeResource) ResourceType() string {
	return "pingone_branding_theme"
}
