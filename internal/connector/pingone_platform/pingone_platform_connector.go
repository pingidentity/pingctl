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
	ServiceName = "pingone-platform"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingonePlatformConnector{}
	_ connector.Authenticatable = &PingonePlatformConnector{}
)

type PingonePlatformConnector struct {
	context       context.Context
	apiClient     *sdk.Client
	environmentID string
}

// Utility method for creating a PingonePlatformConnector
func Connector(ctx context.Context, apiClient *sdk.Client, environmentID string) *PingonePlatformConnector {
	return &PingonePlatformConnector{
		context:       ctx,
		apiClient:     apiClient,
		environmentID: environmentID,
	}
}

func (c *PingonePlatformConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne Platform Resources...")

	hclImportBlockTemplate, err := template.New("HCLImportBlock").Parse(
		`import {
	to = {{.ResourceType}}.{{.ResourceName}}
	id = "{{.ResourceID}}"
}
`,
	)
	if err != nil {
		return err
	}

	exportableResources := []connector.ExportableResource{
		resources.AgreementResource(c.context, c.apiClient, c.environmentID),
	}

	for _, exportableResource := range exportableResources {
		importBlocks, err := exportableResource.ExportAll()
		if err != nil {
			return err
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
			return err
		}

		for _, importBlock := range *importBlocks {
			switch format {
			case connector.ENUMEXPORTFORMAT_HCL:
				err := hclImportBlockTemplate.Execute(outputFile, importBlock)
				if err != nil {
					return err
				}
				// default:
				// Note that this default case is already handled in export.go, and should never be called.
			}
		}
	}

	return nil
}

func (c *PingonePlatformConnector) ConnectorServiceName() string {
	return ServiceName
}

func (c *PingonePlatformConnector) Login() error {
	return nil
}

func (c *PingonePlatformConnector) Logout() error {
	return nil
}
