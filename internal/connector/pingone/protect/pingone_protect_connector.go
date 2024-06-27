package protect

import (
	"context"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/connector/pingone/protect/resources"
	"github.com/pingidentity/pingctl/internal/logger"
)

const (
	serviceName = "pingone-protect"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingoneProtectConnector{}
	_ connector.Authenticatable = &PingoneProtectConnector{}
)

type PingoneProtectConnector struct {
	clientInfo connector.SDKClientInfo
}

// Utility method for creating a PingoneProtectConnector
func ProtectConnector(ctx context.Context, apiClient *sdk.Client, apiClientId *string, exportEnvironmentID string) *PingoneProtectConnector {
	return &PingoneProtectConnector{
		clientInfo: connector.SDKClientInfo{
			Context:             ctx,
			ApiClient:           apiClient,
			ApiClientId:         apiClientId,
			ExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingoneProtectConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all PingOne MFA Resources...")

	exportableResources := []connector.ExportableResource{
		resources.RiskPolicy(&c.clientInfo),
		resources.RiskPredictor(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, overwriteExport)
}

func (c *PingoneProtectConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingoneProtectConnector) Login() error {
	return nil
}

func (c *PingoneProtectConnector) Logout() error {
	return nil
}
