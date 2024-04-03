package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneResourceAttributeResource{}
)

type PingoneResourceAttributeResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneResourceAttributeResource
func ResourceAttribute(clientInfo *connector.SDKClientInfo) *PingoneResourceAttributeResource {
	return &PingoneResourceAttributeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceAttributeResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllResources"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, resource := range embedded.GetResources() {
		resourceId, resourceIdOk := resource.GetIdOk()
		resourceName, resourceNameOk := resource.GetNameOk()

		if resourceIdOk && resourceNameOk {
			apiExecuteFunc = r.clientInfo.ApiClient.ManagementAPIClient.ResourceAttributesApi.ReadAllResourceAttributes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute
			apiFunctionName = "ReadAllResourceAttributes"

			embedded, err = common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, resourceAttribute := range embedded.GetAttributes() {
				if resourceAttribute.ResourceAttribute == nil {
					continue
				}

				resourceAttributeId, resourceAttributeIdOk := resourceAttribute.ResourceAttribute.GetIdOk()
				resourceAttributeName, resourceAttributeNameOk := resourceAttribute.ResourceAttribute.GetNameOk()

				if resourceAttributeIdOk && resourceAttributeNameOk {
					commentData := map[string]string{
						"Resource Type":           r.ResourceType(),
						"Resource Name":           *resourceName,
						"Resource Attribute Name": *resourceAttributeName,
						"Export Environment ID":   r.clientInfo.ExportEnvironmentID,
						"Resource ID":             *resourceId,
						"Resource Attribute ID":   *resourceAttributeId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *resourceName, *resourceAttributeName),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *resourceId, *resourceAttributeId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceAttributeResource) ResourceType() string {
	return "pingone_resource_attribute"
}
