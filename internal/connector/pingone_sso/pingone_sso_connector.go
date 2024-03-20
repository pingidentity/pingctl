package pingone_sso

import (
	"context"
	"fmt"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/connector/pingone_sso/resources"
	"github.com/pingidentity/pingctl/internal/logger"
)

const (
	serviceName = "pingone-sso"
)

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &PingoneSSOConnector{}
	_ connector.Authenticatable = &PingoneSSOConnector{}
)

type PingoneSSOConnector struct {
	clientInfo connector.SDKClientInfo
}

// Utility method for creating a PingoneSSOConnector
func SSOConnector(ctx context.Context, apiClient *sdk.Client, exportEnvironmentID string) *PingoneSSOConnector {
	return &PingoneSSOConnector{
		clientInfo: connector.SDKClientInfo{
			Context:             ctx,
			ApiClient:           apiClient,
			ExportEnvironmentID: exportEnvironmentID,
		},
	}
}

func (c *PingoneSSOConnector) Export(format, outputDir string, overwriteExport bool) error {
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

	l.Debug().Msgf("Exporting all PingOne SSO Resources...")

	exportableResources := []connector.ExportableResource{
		resources.Application(&c.clientInfo),
		resources.ApplicationAttributeMapping(&c.clientInfo),
		resources.ApplicationFlowPolicyAssignment(&c.clientInfo),
		resources.ApplicationResourceGrant(&c.clientInfo),
		resources.ApplicationSignOnPolicyAssignment(&c.clientInfo),
		resources.Group(&c.clientInfo),
		resources.GroupNesting(&c.clientInfo),
		resources.GroupRoleAssignment(&c.clientInfo),
		resources.IdentityProvider(&c.clientInfo),
		resources.IdentityProviderAttribute(&c.clientInfo),
		resources.PasswordPolicy(&c.clientInfo),
		resources.Population(&c.clientInfo),
		resources.PopulationDefault(&c.clientInfo),
		resources.Resource(&c.clientInfo),
		resources.ResourceScope(&c.clientInfo),
		resources.SchemaAttribute(&c.clientInfo),
		resources.SignOnPolicy(&c.clientInfo),
		resources.User(&c.clientInfo),
		resources.UserGroupAssignment(&c.clientInfo),
	}

	return common.WriteFiles(exportableResources, l, format, outputDir, overwriteExport)
}

func (c *PingoneSSOConnector) ConnectorServiceName() string {
	return serviceName
}

func (c *PingoneSSOConnector) Login() error {
	return nil
}

func (c *PingoneSSOConnector) Logout() error {
	return nil
}
