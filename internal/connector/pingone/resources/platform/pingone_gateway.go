package platform

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	pingoneresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneGatewayResource{}
)

type PingoneGatewayResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneGatewayResource
func Gateway(clientInfo *connector.SDKClientInfo) *PingoneGatewayResource {
	return &PingoneGatewayResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneGatewayResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllGateways"

	embedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, gatewayInner := range embedded.GetGateways() {
		var (
			gatewayId     *string
			gatewayName   *string
			gatewayIdOk   bool
			gatewayNameOk bool
		)

		switch {
		case gatewayInner.Gateway != nil:
			gatewayId, gatewayIdOk = gatewayInner.Gateway.GetIdOk()
			gatewayName, gatewayNameOk = gatewayInner.Gateway.GetNameOk()
		case gatewayInner.GatewayTypeLDAP != nil:
			gatewayId, gatewayIdOk = gatewayInner.GatewayTypeLDAP.GetIdOk()
			gatewayName, gatewayNameOk = gatewayInner.GatewayTypeLDAP.GetNameOk()
		case gatewayInner.GatewayTypeRADIUS != nil:
			gatewayId, gatewayIdOk = gatewayInner.GatewayTypeRADIUS.GetIdOk()
			gatewayName, gatewayNameOk = gatewayInner.GatewayTypeRADIUS.GetNameOk()
		default:
			continue
		}

		if gatewayIdOk && gatewayNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *gatewayName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *gatewayId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneGatewayResource) ResourceType() string {
	return "pingone_gateway"
}
