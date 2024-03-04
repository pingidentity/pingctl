package pingone_platform

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone_platform/resources"
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
func Connector(ctx context.Context, apiClient *sdk.Client, exportEnvironmentID string) *PingonePlatformConnector {
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

	hclImportBlockTemplate, err := template.New("HCLImportBlock").Parse(connector.HCLImportBlockTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse HCL import block template. err: %s", err.Error())
	}

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
		resources.Gateway(&c.clientInfo),
		resources.GatewayCredential(&c.clientInfo),
		resources.GatewayRoleAssignment(&c.clientInfo),
		resources.IdentityPropagationPlan(&c.clientInfo),
		resources.Key(&c.clientInfo),
		resources.KeyRotationPolicy(&c.clientInfo),
		resources.Language(&c.clientInfo),
		resources.NotificationPolicy(&c.clientInfo),
		resources.NotificationSettings(&c.clientInfo),
		resources.PhoneDeliverySettings(&c.clientInfo),
	}

	for _, exportableResource := range exportableResources {
		importBlocks, err := exportableResource.ExportAll()
		if err != nil {
			return fmt.Errorf("failed to export resource %s. err: %s", exportableResource.ResourceType(), err.Error())
		}

		if len(*importBlocks) == 0 {
			// No resources exported. Avoid creating an empty import.tf file
			l.Debug().Msgf("Nothing exported for resource %s. Skipping import file generation...", exportableResource.ResourceType())
			continue
		}

		l.Debug().Msgf("Generating import file for %s resource...", exportableResource.ResourceType())

		outputFileName := fmt.Sprintf("%s.tf", exportableResource.ResourceType())
		outputFilePath := filepath.Join(outputDir, filepath.Base(outputFileName))

		// Check to see if outputFile already exists.
		// If so, default behavior is to exit and not overwrite.
		// This can be changed with the --overwrite export parameter
		_, err = os.Stat(outputFilePath)
		if err == nil && !overwriteExport {
			return fmt.Errorf("generated import file for %q already exists. Use --overwrite to overwrite existing export data", outputFileName)
		}

		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return fmt.Errorf("failed to create export file %q. err: %s", outputFilePath, err.Error())
		}
		defer outputFile.Close()

		for _, importBlock := range *importBlocks {
			// Sanitize import block "to". Make lowercase, remove special chars, convert space to underscore
			importBlock.Sanitize()

			switch format {
			case connector.ENUMEXPORTFORMAT_HCL:
				err := hclImportBlockTemplate.Execute(outputFile, importBlock)
				if err != nil {
					return fmt.Errorf("failed to write import block template to file %q. err: %s", outputFilePath, err.Error())
				}
				// default:
				// Note that this default case is already handled in export.go, and should never be called.
			}
		}
	}

	return nil
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
