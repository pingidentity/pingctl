package testutils

import (
	"context"
	"os"
	"sync"
	"testing"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
)

var (
	envIdOnce      sync.Once
	apiClientOnce  sync.Once
	sdkClientInfo  *connector.SDKClientInfo
	environementId string
)

func GetEnvironmentID() string {
	envIdOnce.Do(func() {
		environementId = os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")
	})

	return environementId
}

// Utility method to print log file if present.
func PrintLogs(t *testing.T) {
	t.Helper()

	logContent, err := os.ReadFile(os.Getenv("PINGCTL_LOG_PATH"))
	if err == nil {
		t.Logf("Captured Logs: %s", string(logContent[:]))
	}
}

// Utility method to initialize a PingOne SDK client for testing
func GetPingOneSDKClientInfo(t *testing.T) *connector.SDKClientInfo {
	t.Helper()

	apiClientOnce.Do(func() {
		// Grab environment vars for initializing the API client.
		// These are set in GitHub Actions.
		clientID := os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID")
		clientSecret := os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_SECRET")
		environmentId := GetEnvironmentID()
		region := os.Getenv("PINGCTL_PINGONE_REGION")

		if clientID == "" || clientSecret == "" || environmentId == "" || region == "" {
			t.Fatalf("Unable to retrieve env var value for one or more of clientID, clientSecret, environmentID, region.")
		}

		apiConfig := &sdk.Config{
			ClientID:      &clientID,
			ClientSecret:  &clientSecret,
			EnvironmentID: &environmentId,
			Region:        region,
		}

		// Make empty context for testing
		ctx := context.Background()

		// Initialize the API client
		client, err := apiConfig.APIClient(ctx)
		if err != nil {
			t.Fatalf(err.Error())
		}

		sdkClientInfo = &connector.SDKClientInfo{
			Context:             ctx,
			ApiClient:           client,
			ExportEnvironmentID: environmentId,
		}
	})

	return sdkClientInfo
}

func ValidateImportBlocks(t *testing.T, resource connector.ExportableResource, expectedImportBlocks *[]connector.ImportBlock) {
	t.Helper()

	importBlocks, err := resource.ExportAll()
	if err != nil {
		t.Fatalf("Failed to export %s: %s", resource.ResourceType(), err.Error())
	}

	// Check number of export blocks
	expectedNumberOfBlocks := len(*expectedImportBlocks)
	actualNumberOfBlocks := len(*importBlocks)
	if actualNumberOfBlocks != expectedNumberOfBlocks {
		t.Fatalf("Expected %d import blocks, got %d", expectedNumberOfBlocks, actualNumberOfBlocks)
	}

	// Make sure the importblocks match the expected import blocks
	for index, importBlock := range *importBlocks {
		expectedImportBlock := (*expectedImportBlocks)[index]

		if !importBlock.Equals(expectedImportBlock) {
			t.Errorf("Expected import block \n%s\n Got import block \n%s", expectedImportBlock.String(), importBlock.String())
		}
	}

	// Make sure the resource name and id in each import block is unique across all import blocks
	resourceNames := map[string]bool{}
	resourceIDs := map[string]bool{}
	for _, importBlock := range *importBlocks {
		if resourceNames[importBlock.ResourceName] {
			t.Errorf("Resource name %s is not unique", importBlock.ResourceName)
		}
		resourceNames[importBlock.ResourceName] = true

		if resourceIDs[importBlock.ResourceID] {
			t.Errorf("Resource ID %s is not unique", importBlock.ResourceID)
		}
		resourceIDs[importBlock.ResourceID] = true
	}
}
