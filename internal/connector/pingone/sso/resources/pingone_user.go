package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneUserResource{}
)

type PingoneUserResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneUserResource
func User(clientInfo *connector.SDKClientInfo) *PingoneUserResource {
	return &PingoneUserResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneUserResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.UsersApi.ReadAllUsers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllUsers"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, user := range embedded.GetUsers() {
		userId, userIdOk := user.GetIdOk()
		username, usernameOk := user.GetUsernameOk()
		if userIdOk && usernameOk {
			commentData := map[string]string{
				"Resource Type":         r.ResourceType(),
				"Username":              *username,
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"User ID":               *userId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *username,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *userId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneUserResource) ResourceType() string {
	return "pingone_user"
}
