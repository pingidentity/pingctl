package pingfederate

import (
	"context"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingctl/internal/logger"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

const (
	serviceName = "pingfederate"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingfederateConnector{}
	_ connector.Authenticatable = &PingfederateConnector{}
)

type PingfederateConnector struct {
	clientInfo connector.PingFederateClientInfo
}

// Utility method for creating a PingfederateConnector
func PFConnector(ctx context.Context, apiClient *pingfederateGoClient.APIClient) *PingfederateConnector {
	return &PingfederateConnector{
		clientInfo: connector.PingFederateClientInfo{
			ApiClient: apiClient,
			Context:   ctx,
		},
	}
}

func (c *PingfederateConnector) Export(format, outputDir string, overwriteExport bool) error {
	l := logger.Get()

	l.Debug().Msgf("Exporting all Pingfederate Resources...")

	exportableResources := []connector.ExportableResource{
		resources.AuthenticationApiApplication(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, format, outputDir, overwriteExport)
}

func (c *PingfederateConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingfederateConnector) Login() error {
	return nil
}

func (c *PingfederateConnector) Logout() error {
	return nil
}