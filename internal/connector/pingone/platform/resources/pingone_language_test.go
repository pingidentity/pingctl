package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestLanguageExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.Language(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_language",
			ResourceName: "Chinese",
			ResourceID:   fmt.Sprintf("%s/d4004466-e900-4951-a22b-c230bc28afa0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Czech",
			ResourceID:   fmt.Sprintf("%s/a9b99bbe-5b87-425f-91b3-324e16ea95c5", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Dutch",
			ResourceID:   fmt.Sprintf("%s/ec5dd62a-336c-4e42-b48a-302e76892913", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "English",
			ResourceID:   fmt.Sprintf("%s/88c78fb2-9d74-41e3-a1d8-a9f729a2b463", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "French",
			ResourceID:   fmt.Sprintf("%s/3f8a2e14-0ace-41db-a92d-74b3b7913ffe", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "French (Canada)",
			ResourceID:   fmt.Sprintf("%s/9c6cacb4-4fe7-48bc-bb7f-91acc38baa3d", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "German",
			ResourceID:   fmt.Sprintf("%s/0909ab3f-5793-44b8-94ca-b7e6d0d61e49", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Hungarian",
			ResourceID:   fmt.Sprintf("%s/ea6410f8-daa5-47c4-917c-ce2169b99109", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Italian",
			ResourceID:   fmt.Sprintf("%s/327ba0cb-4902-4162-9ad9-cb2242ae62fd", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Japanese",
			ResourceID:   fmt.Sprintf("%s/3d14f035-455d-4417-9859-a1965c07ad11", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Korean",
			ResourceID:   fmt.Sprintf("%s/4c782975-bbf0-43d8-8ee4-2e790e21ce82", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Polish",
			ResourceID:   fmt.Sprintf("%s/cf552c83-45e4-4796-82a7-bee215bfec72", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Portuguese",
			ResourceID:   fmt.Sprintf("%s/2c41000f-60c6-4a3c-a738-373a74099486", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Russian",
			ResourceID:   fmt.Sprintf("%s/d39432a3-b6b8-41f9-bbee-568089c1d25b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Spanish",
			ResourceID:   fmt.Sprintf("%s/719d5c7c-de5d-4f83-a954-4c2607726c98", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Thai",
			ResourceID:   fmt.Sprintf("%s/6b52ccdd-0ad9-4e55-88d4-cba22ac3b95c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_language",
			ResourceName: "Turkish",
			ResourceID:   fmt.Sprintf("%s/353a7b78-d189-4125-987b-efcb0a4352ba", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
