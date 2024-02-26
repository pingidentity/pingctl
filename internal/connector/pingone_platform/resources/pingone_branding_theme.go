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
func BrandingThemeResource(clientInfo *connector.SDKClientInfo) *PingoneBrandingThemeResource {
	return &PingoneBrandingThemeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneBrandingThemeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_branding_theme resources...")

	entityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.BrandingThemesApi.ReadBrandingThemes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadBrandingSettings Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned ReadBrandingThemes() entity array is nil.")
		l.Error().Msgf("ReadBrandingThemes Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_branding_theme resources via ReadBrandingThemes()")
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned ReadBrandingThemes() embedded data is nil.")
		l.Error().Msgf("ReadBrandingThemes Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_branding_theme resources via ReadBrandingThemes()")
	}

	importBlocks := []connector.ImportBlock{}

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
