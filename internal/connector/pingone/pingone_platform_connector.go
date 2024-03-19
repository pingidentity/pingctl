package pingone

import (
	"context"
	"fmt"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	connectorcommon "github.com/pingidentity/pingctl/internal/connector/common"
	platformresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/platform"
	ssoresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/sso"
	"github.com/pingidentity/pingctl/internal/logger"
)

const (
	serviceName = "pingone-platform"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingonePlatformConnector{}
	_ connector.Authenticatable = &PingonePlatformConnector{}
)

type PingonePlatformConnector struct {
	clientInfo connector.SDKClientInfo
}

// Utility method for creating a PingonePlatformConnector
func PlatformConnector(ctx context.Context, apiClient *sdk.Client, exportEnvironmentID string) *PingonePlatformConnector {
	return &PingonePlatformConnector{
		clientInfo: connector.SDKClientInfo{
			Context:             ctx,
			ApiClient:           apiClient,
			ExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingonePlatformConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Validating export environment ID...")

	environment, response, err := c.clientInfo.ApiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(c.clientInfo.Context, c.clientInfo.ExportEnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadOneEnvironment Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return err
	}

	if environment == nil {
		l.Error().Msgf("Returned ReadOneEnvironment() environment is nil.")
		l.Error().Msgf("ReadOneEnvironment Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		if response.StatusCode == 404 {
			return fmt.Errorf("failed to fetch environment. the provided environment id %q does not exist", c.clientInfo.ExportEnvironmentID)
		} else {
			return fmt.Errorf("failed to fetch environment %q via ReadOneEnvironment()", c.clientInfo.ExportEnvironmentID)
		}
	}

	l.Debug().Msgf("Exporting all PingOne Platform Resources...")

	exportableResources := []connector.ExportableResource{
		platformresources.Agreement(&c.clientInfo),
		platformresources.AgreementEnable(&c.clientInfo),
		platformresources.AgreementLocalization(&c.clientInfo),
		platformresources.AgreementLocalizationEnable(&c.clientInfo),
		platformresources.AgreementLocalizationRevision(&c.clientInfo),
		platformresources.BrandingSettings(&c.clientInfo),
		platformresources.BrandingTheme(&c.clientInfo),
		platformresources.BrandingThemeDefault(&c.clientInfo),
		platformresources.Certificate(&c.clientInfo),
		platformresources.CustomDomain(&c.clientInfo),
		platformresources.Environment(&c.clientInfo),
		platformresources.Form(&c.clientInfo),
		platformresources.FormRecaptchaV2(&c.clientInfo),
		platformresources.Gateway(&c.clientInfo),
		platformresources.GatewayCredential(&c.clientInfo),
		platformresources.GatewayRoleAssignment(&c.clientInfo),
		platformresources.IdentityPropagationPlan(&c.clientInfo),
		platformresources.Key(&c.clientInfo),
		platformresources.KeyRotationPolicy(&c.clientInfo),
		platformresources.Language(&c.clientInfo),
		platformresources.NotificationPolicy(&c.clientInfo),
		platformresources.NotificationSettings(&c.clientInfo),
		platformresources.NotificationSettingsEmail(&c.clientInfo),
		platformresources.NotificationSettingsTemplateContent(&c.clientInfo),
		platformresources.PhoneDeliverySettings(&c.clientInfo),
		platformresources.RoleAssignmentUser(&c.clientInfo),
		platformresources.SystemApplication(&c.clientInfo),
		platformresources.TrustedEmailAddress(&c.clientInfo),
		platformresources.TrustedEmailDomain(&c.clientInfo),
		platformresources.Webhook(&c.clientInfo),
		ssoresources.Application(&c.clientInfo),
		ssoresources.ApplicationAttributeMapping(&c.clientInfo),
		ssoresources.ApplicationFlowPolicyAssignment(&c.clientInfo),
		ssoresources.ApplicationResourceGrant(&c.clientInfo),
		ssoresources.Group(&c.clientInfo),
		ssoresources.GroupNesting(&c.clientInfo),
		ssoresources.IdentityProvider(&c.clientInfo),
		ssoresources.IdentityProviderAttribute(&c.clientInfo),
		ssoresources.PasswordPolicy(&c.clientInfo),
		ssoresources.Population(&c.clientInfo),
		ssoresources.PopulationDefault(&c.clientInfo),
		ssoresources.Resource(&c.clientInfo),
		ssoresources.SignOnPolicy(&c.clientInfo),
		ssoresources.User(&c.clientInfo),
		ssoresources.UserGroupAssignment(&c.clientInfo),
	}

	return connectorcommon.WriteFiles(exportableResources, l, format, outputDir, overwriteExport)
}

func (c *PingonePlatformConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingonePlatformConnector) Login() error {
	return nil
}

func (c *PingonePlatformConnector) Logout() error {
	return nil
}
