package resources

import (
	"fmt"

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

// Utility method for creating a PingoneAgreementLocalizationRevisionResource
func BrandingSettingsResource(clientInfo *connector.SDKClientInfo) *PingoneBrandingSettingsResource {
	return &PingoneBrandingSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneBrandingSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_branding_settings resources...")

	brandingSettings, response, err := r.clientInfo.ApiClient.ManagementAPIClient.BrandingSettingsApi.ReadBrandingSettings(r.clientInfo.Context, r.clientInfo.EnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadBrandingSettings Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if brandingSettings == nil {
		l.Error().Msgf("Returned ReadBrandingSettings() settings are nil.")
		l.Error().Msgf("ReadBrandingSettings Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_branding_settings resources via ReadBrandingSettings()")
	}

	importBlocks := []connector.ImportBlock{}

	// Oddly, environment is nil in the returned brandingSettigs, but id seems to be set to the env ID
	environmentId, environmentIdOk := brandingSettings.GetIdOk()

	if environmentIdOk {
		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType: r.ResourceType(),
			ResourceName: "branding",
			ResourceID:   *environmentId,
		})
	}

	return &importBlocks, nil
}

func (r *PingoneBrandingSettingsResource) ResourceType() string {
	return "pingone_branding_settings"
}
