package platform

import (
	"context"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
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
func PlatformConnector(ctx context.Context, apiClient *sdk.Client, apiClientId *string, exportEnvironmentID string) *PingonePlatformConnector {
	return &PingonePlatformConnector{
		clientInfo: connector.SDKClientInfo{
			Context:             ctx,
			ApiClient:           apiClient,
			ApiClientId:         apiClientId,
			ExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingonePlatformConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne Platform Resources...")

	exportableResources := []connector.ExportableResource{
		resources.Agreement(&c.clientInfo),
		resources.AgreementEnable(&c.clientInfo),
		resources.AgreementLocalization(&c.clientInfo),
		resources.AgreementLocalizationEnable(&c.clientInfo),
		resources.AgreementLocalizationRevision(&c.clientInfo),
		resources.BrandingSettings(&c.clientInfo),
		resources.BrandingTheme(&c.clientInfo),
		resources.BrandingThemeDefault(&c.clientInfo),
		resources.Certificate(&c.clientInfo),
		resources.CustomDomain(&c.clientInfo),
		resources.Environment(&c.clientInfo),
		resources.Form(&c.clientInfo),
		resources.FormRecaptchaV2(&c.clientInfo),
		// resources.Gateway(&c.clientInfo),
		resources.GatewayCredential(&c.clientInfo),
		resources.GatewayRoleAssignment(&c.clientInfo),
		resources.IdentityPropagationPlan(&c.clientInfo),
		resources.Key(&c.clientInfo),
		resources.KeyRotationPolicy(&c.clientInfo),
		resources.Language(&c.clientInfo),
		resources.LanguageUpdate(&c.clientInfo),
		resources.NotificationPolicy(&c.clientInfo),
		resources.NotificationSettings(&c.clientInfo),
		resources.NotificationSettingsEmail(&c.clientInfo),
		resources.NotificationTemplateContent(&c.clientInfo),
		resources.PhoneDeliverySettings(&c.clientInfo),
		resources.RoleAssignmentUser(&c.clientInfo),
		resources.SystemApplication(&c.clientInfo),
		resources.TrustedEmailAddress(&c.clientInfo),
		resources.TrustedEmailDomain(&c.clientInfo),
		resources.Webhook(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, c.ConnectorServiceName(), overwriteExport)
}

func (c *PingonePlatformConnector) ExportSingle(format, outputDir string, overwriteExport bool, resource connector.ExportableResource) error {
	return common.WriteFiles([]connector.ExportableResource{resource}, format, outputDir, c.ConnectorServiceName(), overwriteExport)
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
