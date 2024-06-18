package platform_internal

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/go-uuid"
	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
	"github.com/spf13/viper"
)

// Helper test function to get an API client
func getApiClient(t *testing.T) *sdk.Client {
	t.Helper()

	// Set viper configuration needed to initialize the API client
	viper.Set("pingone.worker.clientid", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID"))
	viper.Set("pingone.worker.clientsecret", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_SECRET"))
	viper.Set("pingone.worker.environmentid", os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"))
	viper.Set("pingone.region", os.Getenv("PINGCTL_PINGONE_REGION"))

	// Initialize the API client
	apiClient, apiClientId, err := initApiClient(context.Background(), "v1.2.3")
	if err != nil {
		t.Errorf("initApiClient() error = %v", err)
	}

	// Check API client is not nil
	if apiClient == nil {
		t.Errorf("initApiClient() apiClient = %v, want non-nil", apiClient)
	}

	// Check API client ID is not empty
	if apiClientId == "" {
		t.Errorf("initApiClient() apiClientId = '%s', want non-empty", apiClientId)
	}

	// Check api client id is a valid UUID
	if _, err := uuid.ParseUUID(apiClientId); err != nil {
		t.Errorf("initApiClient() apiClientId = '%s', want valid UUID", apiClientId)
	}

	return apiClient
}

// Test initApiClient function
func Test_initApiClient(t *testing.T) {
	// Test the function
	getApiClient(t)
}

// Test initApiClient function fails on incomplete configuration
func Test_initApiClient_incompleteConfig(t *testing.T) {
	// Set incomplete viper configuration
	viper.Set("pingone.worker.clientid", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID"))
	viper.Set("pingone.worker.clientsecret", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_SECRET"))
	viper.Set("pingone.worker.environmentid", os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"))
	viper.Set("pingone.region", "")

	expectedErrorPattern := `^failed to initialize pingone API client\. unrecognized pingone region: ''\. Must be one of: [A-Za-z\s,]+$`
	_, _, err := initApiClient(context.Background(), "v1.2.3")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test initApiClient function fails on invalid region configuration
func Test_initApiClient_invalidRegionConfig(t *testing.T) {
	// Set invalid viper configuration
	viper.Set("pingone.worker.clientid", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID"))
	viper.Set("pingone.worker.clientsecret", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_SECRET"))
	viper.Set("pingone.worker.environmentid", os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"))
	viper.Set("pingone.region", "invalid")

	expectedErrorPattern := `^failed to initialize pingone API client\. unrecognized pingone region: 'invalid'\. Must be one of: [A-Za-z\s,]+$`
	_, _, err := initApiClient(context.Background(), "v1.2.3")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test initApiClient function fails on client ID configuration
func Test_initApiClient_invalidClientIdConfig(t *testing.T) {
	// Set invalid viper configuration
	viper.Set("pingone.worker.clientid", "12345678-1234-1234-1234-123456789012")
	viper.Set("pingone.worker.clientsecret", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_SECRET"))
	viper.Set("pingone.worker.environmentid", os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"))
	viper.Set("pingone.region", os.Getenv("PINGCTL_PINGONE_REGION"))

	expectedErrorPattern := `^failed to initialize pingone API client\.\s+oauth2: "invalid_client" "Request denied: Invalid client credentials \(Correlation ID: [0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}\)"\s+configuration values used for client initialization:\s+worker client ID - [0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}\s+worker environment ID - [0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}\s+pingone region - [A-Za-z]+\s+worker client secret - .+$`
	_, _, err := initApiClient(context.Background(), "v1.2.3")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test fixEmptyOutputDirVar function with outputDir non-empty
func Test_fixEmptyOutputDirVar_WithOutputDir(t *testing.T) {
	oldOutputDir := os.TempDir()

	expectedErrorPattern := "" // No error expected
	newOutputDir, err := fixEmptyOutputDirVar(oldOutputDir)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	if newOutputDir != oldOutputDir {
		t.Errorf("fixEmptyOutputDirVar() newOutputDir = '%s', want '%s'", newOutputDir, oldOutputDir)
	}
}

// Test fixEmptyOutputDirVar function with outputDir empty
func Test_fixEmptyOutputDirVar_WithoutOutputDir(t *testing.T) {
	expectedErrorPattern := "" // No error expected
	newOutputDir, err := fixEmptyOutputDirVar("")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	if newOutputDir == "" {
		t.Errorf("fixEmptyOutputDirVar() newOutputDir = '%s', want non-empty", newOutputDir)
	}
}

// Test createOrValidateOutputDir function
// - Empty directory that exists is valid, and should not return an error
func Test_createOrValidateOutputDir(t *testing.T) {
	// Create a directory in the temp directory
	outputDir := os.TempDir() + "/pingctlTestCreateOrValidateOutputDir"

	if err := os.Mkdir(outputDir, 0755); err != nil {
		t.Fatalf("os.Mkdir() error = %v", err)
	}

	expectedErrorPattern := "" // No error expected
	err := createOrValidateOutputDir(outputDir, false)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	// Remove the new directory
	if err := os.RemoveAll(outputDir); err != nil {
		t.Fatalf("os.RemoveAll() error = %v", err)
	}
}

// Test createOrValidateOutputDir function
// - Create an empty directory that does exist
// - Add a file to the directory
// - Validate that the function returns an error with overwrite set to false
func Test_createOrValidateOutputDir_WithFile(t *testing.T) {
	// Create a directory in the temp directory
	outputDir := os.TempDir() + "/pingctlTestCreateOrValidateOutputDirWithFile"

	if err := os.Mkdir(outputDir, 0755); err != nil {
		t.Fatalf("os.Mkdir() error = %v", err)
	}

	// Create a file in the new directory
	file := outputDir + "/file"
	if _, err := os.Create(file); err != nil {
		t.Fatalf("os.Create() error = %v", err)
	}

	expectedErrorPattern := `^'platform export' output directory '[\/A-Za-z0-9_-]+' is not empty\. Use --overwrite to overwrite existing export data$`
	err := createOrValidateOutputDir(outputDir, false)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	// Remove the new directory
	if err := os.RemoveAll(outputDir); err != nil {
		t.Fatalf("os.RemoveAll() error = %v", err)
	}
}

// Test createOrValidateOutputDir function
// - Create an empty directory that does exist
// - Add a file to the directory
// - Validate that the function does not return an error with overwrite set to true
func Test_createOrValidateOutputDir_WithFile_Overwrite(t *testing.T) {
	// Create a directory in the temp directory
	outputDir := os.TempDir() + "/pingctlTestCreateOrValidateOutputDirWithFileOverwrite"

	if err := os.Mkdir(outputDir, 0755); err != nil {
		t.Fatalf("os.Mkdir() error = %v", err)
	}

	// Create a file in the new directory
	file := outputDir + "/file"
	if _, err := os.Create(file); err != nil {
		t.Fatalf("os.Create() error = %v", err)
	}

	expectedErrorPattern := "" // No error expected
	err := createOrValidateOutputDir(outputDir, true)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	// Remove the new directory
	if err := os.RemoveAll(outputDir); err != nil {
		t.Fatalf("os.RemoveAll() error = %v", err)
	}
}

// Test createOrValidateOutputDir function
// - Provide function with a directory that does not exist
// - Validate that the function creates the directory
func Test_createOrValidateOutputDir_WithoutDir(t *testing.T) {
	// Create a directory in the temp directory
	outputDir := os.TempDir() + "/pingctlTestCreateOrValidateOutputDirWithoutDir"

	expectedErrorPattern := "" // No error expected
	err := createOrValidateOutputDir(outputDir, false)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	// Validate the directory was created
	if _, err := os.Stat(outputDir); err != nil {
		t.Errorf("os.Stat() error = %v", err)
	}

	// Remove the new directory
	if err := os.RemoveAll(outputDir); err != nil {
		t.Fatalf("os.RemoveAll() error = %v", err)
	}
}

// Test getExportEnvID function
func Test_getExportEnvID(t *testing.T) {
	// Set viper configuration needed to get the export environment ID
	oldExportEnvID := "12345678-1234-1234-1234-123456789012"
	viper.Set("pingone.export.environmentid", oldExportEnvID)

	expectedErrorPattern := "" // No error expected
	newExportEnvID, err := getExportEnvID()
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	// Check envID is not empty
	if newExportEnvID == "" {
		t.Errorf("getExportEnvID() envID = '%s', want non-empty", newExportEnvID)
	}

	// Check envID is the same as the one set in viper
	if newExportEnvID != oldExportEnvID {
		t.Errorf("getExportEnvID() envID = '%s', want '%s'", newExportEnvID, oldExportEnvID)
	}
}

// Test getExportEnvID function fails on missing configuration
func Test_getExportEnvID_missingConfig(t *testing.T) {
	// Clear viper configuration
	viper.Set("pingone.export.environmentid", "")
	viper.Set("pingone.worker.environmentid", "")

	expectedErrorPattern := `^failed to determine export environment ID$`
	_, err := getExportEnvID()
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test getExportEnvID function defaults to worker environment ID
// when export environment ID is not set
func Test_getExportEnvID_defaultToWorkerEnvID(t *testing.T) {
	// Set viper configuration needed to get the export environment ID
	oldWorkerEnvID := "12345678-1234-1234-1234-123456789012"
	viper.Set("pingone.worker.environmentid", oldWorkerEnvID)

	expectedErrorPattern := "" // No error expected
	newExportEnvID, err := getExportEnvID()
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	// Check envID is not empty
	if newExportEnvID == "" {
		t.Errorf("getExportEnvID() envID = '%s', want non-empty", newExportEnvID)
	}

	// Check envID is the same as the one set in viper
	if newExportEnvID != oldWorkerEnvID {
		t.Errorf("getExportEnvID() envID = '%s', want '%s'", newExportEnvID, oldWorkerEnvID)
	}
}

// Test validateExportEnvID function
// - initialize the API client
// - get the export environment ID from client environment id
// - validate the export environment ID
func Test_validateExportEnvID(t *testing.T) {
	// Get apiClient from helper function
	apiClient := getApiClient(t)

	expectedErrorPattern := "" // No error expected
	err := validateExportEnvID(context.Background(), os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"), apiClient)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test validateExportEnvID function fails on invalid export environment ID
func Test_validateExportEnvID_invalidEnvID(t *testing.T) {
	// Get apiClient from helper function
	apiClient := getApiClient(t)

	expectedErrorPattern := `^ReadOneEnvironment Request for resource 'pingone_environment' was not successful\.\s+Response Code: 404 Not Found\s+Response Body: {{\s+"id" : "[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}",\s+"code" : "NOT_FOUND",\s+"message" : "Unable to find environment with ID: '[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}'"\s+}}\s+Error: 404 Not Found$`
	err := validateExportEnvID(context.Background(), "12345678-1234-1234-1234-123456789012", apiClient)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test validateExportEnvID function fails on nil context
func Test_validateExportEnvID_nilContext(t *testing.T) {
	// Get apiClient from helper function
	apiClient := getApiClient(t)

	expectedErrorPattern := `^failed to validate environment ID '[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}'. context is nil$`
	// nolint:staticcheck // ignore SA1012 this is a test
	err := validateExportEnvID(nil, os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"), apiClient) //lint:ignore SA1012 this is a test
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test validateExportEnvID function fails on nil API client
func Test_validateExportEnvID_nilApiClient(t *testing.T) {
	expectedErrorPattern := `^failed to validate environment ID '[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}'. apiClient is nil$`
	err := validateExportEnvID(context.Background(), os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"), nil)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test getExportableConnectors function
func Test_getExportableConnectors(t *testing.T) {
	// Get apiClient from helper function
	apiClient := getApiClient(t)

	// Get the API clientID from env var
	apiClientId := os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID")
	exportEnvID := os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")

	// Initialize multiService with all services
	multiService := customtypes.NewMultiService()
	numServices := len(*multiService.GetServices())

	expectedErrorPattern := "" // No error expected
	exportableConnectors := getExportableConnectors(exportEnvID, apiClientId, context.Background(), multiService, apiClient)
	testutils_helpers.CheckExpectedError(t, nil, expectedErrorPattern)

	// Check the number of exportable connectors
	if len(*exportableConnectors) == 0 {
		t.Errorf("getExportableConnectors() exportableConnectors = %v, want non-empty", exportableConnectors)
	}

	// Check the number of exportable connectors
	if len(*exportableConnectors) != numServices {
		t.Errorf("getExportableConnectors() num exportableConnectors = %v, want %v", len(*exportableConnectors), numServices)
	}
}

// Test getExportableConnectors function returns no exportable connectors
// when no services are provided
func Test_getExportableConnectors_noServices(t *testing.T) {
	// Get apiClient from helper function
	apiClient := getApiClient(t)

	// Get the API clientID from env var
	apiClientId := os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID")
	exportEnvID := os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")

	expectedErrorPattern := "" // No error expected
	exportableConnectors := getExportableConnectors(exportEnvID, apiClientId, context.Background(), nil, apiClient)
	testutils_helpers.CheckExpectedError(t, nil, expectedErrorPattern)

	// Check the number of exportable connectors
	if len(*exportableConnectors) != 0 {
		t.Errorf("getExportableConnectors() num exportableConnectors = %v, want 0", len(*exportableConnectors))
	}
}

// Test getExportableConnectors function returns only one exportable connector
// when only one correct service is provided
func Test_getExportableConnectors_oneService(t *testing.T) {
	// Get apiClient from helper function
	apiClient := getApiClient(t)

	// Get the API clientID from env var
	apiClientId := os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID")
	exportEnvID := os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")

	// Initialize multiService with one service
	multiService := customtypes.NewMultiService()
	if err := multiService.Set("pingone-sso"); err != nil {
		t.Errorf("multiService.Set() error = %v", err)
	}

	expectedErrorPattern := "" // No error expected
	exportableConnectors := getExportableConnectors(exportEnvID, apiClientId, context.Background(), multiService, apiClient)
	testutils_helpers.CheckExpectedError(t, nil, expectedErrorPattern)

	// Check the number of exportable connectors
	if len(*exportableConnectors) != 1 {
		t.Errorf("getExportableConnectors() num exportableConnectors = %v, want 1", len(*exportableConnectors))
	}

	// Check connector is not nil
	connector := (*exportableConnectors)[0]
	if connector == nil {
		t.Errorf("getExportableConnectors() connector = %v, want non-nil", connector)
	}

	// Check connector is of type sso.PingoneSSOConnector
	if connector.ConnectorServiceName() != "pingone-sso" {
		t.Errorf("getExportableConnectors() connector = %v, want 'pingone-sso'", connector.ConnectorServiceName())
	}
}

// Test exportConnectors function
func Test_exportConnectors(t *testing.T) {
	// Get apiClient from helper function
	apiClient := getApiClient(t)

	// Get the API clientID from env var
	apiClientId := os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID")
	exportEnvID := os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")

	exportableConnectors := []connector.Exportable{
		mfa.MFAConnector(context.Background(), apiClient, &apiClientId, exportEnvID),
	}

	// Create a directory in the temp directory
	outputDir := os.TempDir() + "/pingctlTestExportConnectors"
	if err := os.Mkdir(outputDir, 0755); err != nil {
		t.Fatalf("os.Mkdir() error = %v", err)
	}

	expectedErrorPattern := "" // No error expected
	err := exportConnectors(&exportableConnectors, connector.ENUMEXPORTFORMAT_HCL, outputDir, false)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	// MFA connector has 4 resources
	// Check the number of files in the directory
	files, err := os.ReadDir(outputDir)
	if err != nil {
		t.Fatalf("os.ReadDir() error = %v", err)
	}
	if len(files) != 4 {
		t.Errorf("exportConnectors() num files = %v, want 4", len(files))
	}

	// Empty the directory
	if err := os.RemoveAll(outputDir); err != nil {
		t.Fatalf("os.RemoveAll() error = %v", err)
	}
}

// Test exportConnectors function fails on invalid output directory
func Test_exportConnectors_invalidOutputDir(t *testing.T) {
	// Get apiClient from helper function
	apiClient := getApiClient(t)

	// Get the API clientID from env var
	apiClientId := os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID")
	exportEnvID := os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")

	exportableConnectors := []connector.Exportable{
		mfa.MFAConnector(context.Background(), apiClient, &apiClientId, exportEnvID),
	}

	expectedErrorPattern := `^failed to export 'pingone-mfa' service: failed to create export file "/invalid/[a-z_]+\.tf"\. err: open /invalid/[a-z_]+\.tf: no such file or directory$`
	err := exportConnectors(&exportableConnectors, connector.ENUMEXPORTFORMAT_HCL, "/invalid", false)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test exportConnectors function fails on nil exportable connectors
func Test_exportConnectors_nilExportableConnectors(t *testing.T) {
	// Create a directory in the temp directory
	outputDir := os.TempDir() + "/pingctlTestExportConnectorsNilExportableConnectors"
	if err := os.Mkdir(outputDir, 0755); err != nil {
		t.Fatalf("os.Mkdir() error = %v", err)
	}

	expectedErrorPattern := `^failed to export services\. exportable connectors list is nil$`
	err := exportConnectors(nil, connector.ENUMEXPORTFORMAT_HCL, outputDir, false)
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	// Empty the directory
	if err := os.RemoveAll(outputDir); err != nil {
		t.Fatalf("os.RemoveAll() error = %v", err)
	}
}
