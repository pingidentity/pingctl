package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationResourceGrantResource{}
)

type PingoneApplicationResourceGrantResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneApplicationResourceGrantResource
func ApplicationResourceGrant(clientInfo *connector.SDKClientInfo) *PingoneApplicationResourceGrantResource {
	return &PingoneApplicationResourceGrantResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationResourceGrantResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteApplicationsFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiApplicationFunctionName := "ReadAllApplications"

	embedded, err := common.GetManagementEmbedded(apiExecuteApplicationsFunc, apiApplicationFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, app := range embedded.GetApplications() {
		var (
			appId     *string
			appIdOk   bool
			appName   *string
			appNameOk bool
		)

		switch {
		case app.ApplicationPingOnePortal != nil:
			appId, appIdOk = app.ApplicationPingOnePortal.GetIdOk()
			appName, appNameOk = app.ApplicationPingOnePortal.GetNameOk()
		case app.ApplicationPingOneSelfService != nil:
			appId, appIdOk = app.ApplicationPingOneSelfService.GetIdOk()
			appName, appNameOk = app.ApplicationPingOneSelfService.GetNameOk()
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
		case app.ApplicationWSFED != nil:
			appId, appIdOk = app.ApplicationWSFED.GetIdOk()
			appName, appNameOk = app.ApplicationWSFED.GetNameOk()
		default:
			continue
		}

		if appIdOk && appNameOk {
			apiExecutePoliciesFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationResourceGrantsApi.ReadAllApplicationGrants(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiApplicationGrantFunctionName := "ReadAllApplicationGrants"

			applicationEmbedded, err := common.GetManagementEmbedded(apiExecutePoliciesFunc, apiApplicationGrantFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, grant := range applicationEmbedded.GetGrants() {
				grantId, grantIdOk := grant.GetIdOk()
				if grantIdOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *appName, *grantId),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *grantId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationResourceGrantResource) ResourceType() string {
	return "pingone_application_resource_grant"
}
