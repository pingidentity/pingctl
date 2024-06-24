package mfa

import (
	"context"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingctl/internal/logger"
)

const (
	serviceName = "pingone-mfa"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingoneMFAConnector{}
	_ connector.Authenticatable = &PingoneMFAConnector{}
)

type PingoneMFAConnector struct {
	clientInfo connector.SDKClientInfo
}

// Utility method for creating a PingoneMFAConnector
func MFAConnector(ctx context.Context, apiClient *sdk.Client, apiClientId *string, exportEnvironmentID string) *PingoneMFAConnector {
	return &PingoneMFAConnector{
		clientInfo: connector.SDKClientInfo{
			Context:             ctx,
			ApiClient:           apiClient,
			ApiClientId:         apiClientId,
			ExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingoneMFAConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne MFA Resources...")

	exportableResources := []connector.ExportableResource{
		resources.MFAApplicationPushCredential(&c.clientInfo),
		resources.MFAFido2Policy(&c.clientInfo),
		resources.MFAPolicy(&c.clientInfo),
		resources.MFASettings(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, c.ConnectorServiceName(), overwriteExport)
}

func (c *PingoneMFAConnector) ExportSingle(format, outputDir string, overwriteExport bool, resource connector.ExportableResource) error {
	return common.WriteFiles([]connector.ExportableResource{resource}, format, outputDir, c.ConnectorServiceName(), overwriteExport)
}

func (c *PingoneMFAConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingoneMFAConnector) Login() error {
	return nil
}

func (c *PingoneMFAConnector) Logout() error {
	return nil
}
