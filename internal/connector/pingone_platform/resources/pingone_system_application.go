package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneSystemApplicationResource{}
)

type PingoneSystemApplicationResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneSystemApplicationResource
func SystemApplication(clientInfo *connector.SDKClientInfo) *PingoneSystemApplicationResource {
	return &PingoneSystemApplicationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneSystemApplicationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllApplications"

	embedded, err := GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
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
		default:
			continue
		}

		if appIdOk && appNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *appName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *appId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneSystemApplicationResource) ResourceType() string {
	return "pingone_system_application"
}
