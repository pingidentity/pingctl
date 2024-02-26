package resources

import (
	"fmt"

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
func BrandingThemeDefaultResource(clientInfo *connector.SDKClientInfo) *PingoneBrandingThemeDefaultResource {
	return &PingoneBrandingThemeDefaultResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneBrandingThemeDefaultResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_branding_theme_default resources...")

	entityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadBrandingSettings Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned ReadBrandingThemes() entity array is nil.")
		l.Error().Msgf("ReadBrandingThemes Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_branding_theme_default resources via ReadBrandingThemes()")
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned ReadBrandingThemes() embedded data is nil.")
		l.Error().Msgf("ReadBrandingThemes Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_branding_theme_default resources via ReadBrandingThemes()")
	}

	importBlocks := []connector.ImportBlock{}

	for _, theme := range embedded.GetThemes() {
		// Only add an import block for the default/active theme
		themeDefault, themeDefaultOk := theme.GetDefaultOk()
		if themeDefaultOk && *themeDefault {
			themeConfiguration, themeConfigurationOk := theme.GetConfigurationOk()
			var themeName *string
			var themeNameOk = false
			if themeConfigurationOk {
				themeName, themeNameOk = themeConfiguration.GetNameOk()
			}

			if themeConfigurationOk && themeNameOk {
				importBlocks = append(importBlocks, connector.ImportBlock{
					ResourceType: r.ResourceType(),
					ResourceName: fmt.Sprintf("%s_active", *themeName),
					ResourceID:   r.clientInfo.ExportEnvironmentID,
				})
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneBrandingThemeDefaultResource) ResourceType() string {
	return "pingone_branding_theme_default"
}
