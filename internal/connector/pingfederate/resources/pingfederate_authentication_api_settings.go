package resources

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingFederateAuthenticationApiSettingsResource{}
)

type PingFederateAuthenticationApiSettingsResource struct {
	clientInfo *connector.PingFederateClientInfo
}

// Utility method for creating a PingFederateAuthenticationApiSettingsResource
func AuthenticationApiSettings(clientInfo *connector.PingFederateClientInfo) *PingFederateAuthenticationApiSettingsResource {
	return &PingFederateAuthenticationApiSettingsResource{
		clientInfo: clientInfo,
	}
}

func (r *PingFederateAuthenticationApiSettingsResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	authnApiSettingsId := "authn_api_settings_id"
	authnApiSettingsName := "Authentication API Settings"

	commentData := map[string]string{
		"Resource Type": r.ResourceType(),
		"Authentication API Settings Resource ID":   authnApiSettingsId,
		"Authentication API Settings Resource Name": authnApiSettingsName,
	}

	importBlocks = append(importBlocks, connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       authnApiSettingsName,
		ResourceID:         authnApiSettingsId,
		CommentInformation: common.GenerateCommentInformation(commentData),
	})

	return &importBlocks, nil
}

func (r *PingFederateAuthenticationApiSettingsResource) ResourceType() string {
	return "pingfederate_authentication_api_settings"
}
