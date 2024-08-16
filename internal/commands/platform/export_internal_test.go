package platform_internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

// Test RunInternalExport function
func TestRunInternalExport(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Create a directory in the temp directory
	outputDir := t.TempDir()

	// Run the export command
	err := RunInternalExport(context.Background(), "v1.2.3", outputDir, connector.ENUMEXPORTFORMAT_HCL, true, customtypes.NewMultiService(), false, false)
	testutils.CheckExpectedError(t, err, nil)

	// Check if there are terraform files in the export directory
	files, err := os.ReadDir(outputDir)
	if err != nil {
		t.Fatalf("os.ReadDir() error = %v", err)
	}

	// Check the number of files in the directory
	if len(files) == 0 {
		t.Errorf("RunInternalExport() num files = %v, want non-zero", len(files))
	}

	// Check the file type is .tf
	re := regexp.MustCompile(`^.*\.tf$`)
	for _, file := range files {
		if file.IsDir() {
			t.Errorf("RunInternalExport() file = %v, want file", file)
		}

		if !re.MatchString(file.Name()) {
			t.Errorf("RunInternalExport() file = %v, want .tf file", file.Name())
		}
	}
}

// Test RunInternalExport function fails on invalid output directory
func TestRunInternalExport_invalidOutputDir(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^failed to create 'platform export' output directory '/invalid': mkdir /invalid:.*$`
	err := RunInternalExport(context.Background(), "v1.2.3", "/invalid", connector.ENUMEXPORTFORMAT_HCL, true, customtypes.NewMultiService(), false, false)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalExport function fails on invalid export format
func TestRunInternalExport_invalidExportFormat(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Create a directory in the temp directory
	outputDir := t.TempDir()

	expectedErrorPattern := `^failed to export '.*' service: unrecognized export format "invalid". Must be one of: \[.*\]$`
	err := RunInternalExport(context.Background(), "v1.2.3", outputDir, "invalid", true, customtypes.NewMultiService(), false, false)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalExport function fails nil context
func TestRunInternalExport_nilContext(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Create a directory in the temp directory
	outputDir := t.TempDir()

	expectedErrorPattern := `^failed to run 'platform export' command. context is nil$`
	// nolint:staticcheck // ignore SA1012 this is a test
	err := RunInternalExport(nil, "v1.2.3", outputDir, connector.ENUMEXPORTFORMAT_HCL, true, customtypes.NewMultiService(), false, false) //lint:ignore SA1012 this is a test
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalExport function fails when not overwriting a populated output directory
func TestRunInternalExport_notOverwriting(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Create a directory in the temp directory
	outputDir := t.TempDir()

	// Create a file in the new directory
	file := outputDir + "/file"
	if _, err := os.Create(file); err != nil {
		t.Fatalf("os.Create() error = %v", err)
	}

	expectedErrorPattern := `'platform export' output directory '.*' is not empty. Use --overwrite to overwrite existing export data$`
	err := RunInternalExport(context.Background(), "v1.2.3", outputDir, connector.ENUMEXPORTFORMAT_HCL, false, customtypes.NewMultiService(), false, false)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalExport function succeeds on nil multiService
func TestRunInternalExport_nilMultiService(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Create a directory in the temp directory
	outputDir := t.TempDir()

	err := RunInternalExport(context.Background(), "v1.2.3", outputDir, connector.ENUMEXPORTFORMAT_HCL, true, nil, false, false)
	testutils.CheckExpectedError(t, err, nil)

	// Check if there are terraform files in the export directory
	files, err := os.ReadDir(outputDir)
	if err != nil {
		t.Fatalf("os.ReadDir() error = %v", err)
	}

	// Check the number of files in the directory
	if len(files) != 0 {
		t.Errorf("RunInternalExport() num files = %v, want 0", len(files))
	}
}

// Test initPingFederateServices function
func Test_initPingFederateServices(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test the function
	err := initPingFederateServices(context.Background(), "v1.2.3", false, false)
	testutils.CheckExpectedError(t, err, nil)

	// Check the API client is not nil
	if pingfederateApiClient == nil {
		t.Errorf("initPingFederateServices() apiClient = %v, want non-nil", pingfederateApiClient)
	}

	// Check the context is not nil
	if pingfederateContext == nil {
		t.Errorf("initPingFederateServices() context = %v, want non-nil", pingfederateContext)
	}
}

// Test initPingFederateServices function with basic auth used
func Test_initPingFederateServices_basicAuth(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test the function
	err := initPingFederateServices(context.Background(), "v1.2.3", true, false)
	testutils.CheckExpectedError(t, err, nil)

	// Check the API client is not nil
	if pingfederateApiClient == nil {
		t.Fatalf("initPingFederateServices() apiClient = %v, want non-nil", pingfederateApiClient)
	}

	// Check the context is not nil
	if pingfederateContext == nil {
		t.Fatalf("initPingFederateServices() context = %v, want non-nil", pingfederateContext)
	}

	// Check context has basic auth value
	basicAuthInfo := pingfederateContext.Value(pingfederateGoClient.ContextBasicAuth)
	if basicAuthInfo == nil {
		t.Fatalf("initPingFederateServices() basicAuthInfo = %v, want non-nil", basicAuthInfo)
	}

	// Check auth username and password are not empty
	basicAuth, ok := basicAuthInfo.(pingfederateGoClient.BasicAuth)
	if !ok {
		t.Fatalf("initPingFederateServices() basicAuth = %v, want PF go-client BasicAuth struct.", basicAuth)
	}
	if basicAuth.UserName == "" || basicAuth.Password == "" {
		t.Errorf("initPingFederateServices() basicAuth = %v, want non-empty username and password", basicAuth)
	}

	if basicAuth.UserName != os.Getenv(profiles.PingFederateUsernameOption.EnvVar) {
		t.Errorf("initPingFederateServices() basicAuth.UserName = %v, want %v", basicAuth.UserName, os.Getenv(profiles.PingFederateUsernameOption.EnvVar))
	}

	if basicAuth.Password != os.Getenv(profiles.PingFederatePasswordOption.EnvVar) {
		t.Errorf("initPingFederateServices() basicAuth.Password = %v, want %v", basicAuth.Password, os.Getenv(profiles.PingFederatePasswordOption.EnvVar))
	}
}

// Test initPingFederateServices function with client credentials used
func Test_initPingFederateServices_clientCredentials(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test the function
	err := initPingFederateServices(context.Background(), "v1.2.3", false, false)
	testutils.CheckExpectedError(t, err, nil)

	// Check the API client is not nil
	if pingfederateApiClient == nil {
		t.Fatalf("initPingFederateServices() apiClient = %v, want non-nil", pingfederateApiClient)
	}

	// Check the context is not nil
	if pingfederateContext == nil {
		t.Fatalf("initPingFederateServices() context = %v, want non-nil", pingfederateContext)
	}

	// Check context has basic auth value
	ccAuthInfo := pingfederateContext.Value(pingfederateGoClient.ContextOAuth2)
	if ccAuthInfo == nil {
		t.Fatalf("initPingFederateServices() ccAuthInfo = %v, want non-nil", ccAuthInfo)
	}

	ccAuth, ok := ccAuthInfo.(pingfederateGoClient.OAuthValues)
	if !ok {
		t.Fatalf("initPingFederateServices() ccAuth = %v, want PF go-client OAuthValues struct.", ccAuth)
	}

	// Check client ID and secret are not empty
	if ccAuth.ClientId == "" || ccAuth.ClientSecret == "" {
		t.Errorf("initPingFederateServices() ccAuth = %v, want non-empty client ID and secret", ccAuth)
	}

	if ccAuth.ClientId != os.Getenv(profiles.PingFederateClientIDOption.EnvVar) {
		t.Errorf("initPingFederateServices() ccAuth.ClientId = %v, want %v", ccAuth.ClientId, os.Getenv(profiles.PingFederateClientIDOption.EnvVar))
	}

	if ccAuth.ClientSecret != os.Getenv(profiles.PingFederateClientSecretOption.EnvVar) {
		t.Errorf("initPingFederateServices() ccAuth.ClientSecret = %v, want %v", ccAuth.ClientSecret, os.Getenv(profiles.PingFederateClientSecretOption.EnvVar))
	}
}

// Test initPingFederateServices function fails when no auth method is provided
func Test_initPingFederateServices_noAuth(t *testing.T) {
	testutils_viper.InitVipersCustomFile(t, fmt.Sprintf(`activeProfile: default
default:
    description: "default description"
    pingctl:
        color: true
        outputFormat: text
    pingone:
        export:
            environmentid: ""
        region: ""
        worker:
            clientid: ""
            clientsecret: ""
            environmentid: ""
    pingfederate:
        accesstokenauth:
            accesstoken: ""
        adminapipath: ""
        basicauth:
            password: ""
            username: ""
        clientcredentialsauth:
            clientid: ""
            clientsecret: ""
            scopes: ""
            tokenurl: ""
        httpshost: "%s"`,
		os.Getenv(profiles.PingFederateHttpsHostOption.EnvVar)))

	expectedErrorPattern := `^failed to initialize PingFederate API client\. none of the following sets of authentication configuration values are set: OAuth2 client credentials \(client ID, client secret, token URL\), Access token, or Basic Authentication credentials \(username, password\)\. configure these properties via parameter flags, environment variables, or the tool\'s configuration file \(default: \$HOME/\.pingctl/config\.yaml\)$`
	err := initPingFederateServices(context.Background(), "v1.2.3", false, false)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingFederateServices function fails when no https host is provided
func Test_initPingFederateServices_noHttpsHost(t *testing.T) {
	testutils_viper.InitVipersCustomFile(t, `activeProfile: default
default:
    description: "default description"
    pingctl:
        color: true
        outputFormat: text
    pingone:
        export:
            environmentid: ""
        region: ""
        worker:
            clientid: ""
            clientsecret: ""
            environmentid: ""
    pingfederate:
        accesstokenauth:
            accesstoken: ""
        adminapipath: ""
        basicauth:
            password: ""
            username: ""
        clientcredentialsauth:
            clientid: ""
            clientsecret: ""
            scopes: ""
            tokenurl: ""
        httpshost: ""`)

	expectedErrorPattern := `^failed to initialize pingfederate API client\. the pingfederate https host configuration value is not set: configure this property via parameter flags, environment variables, or the tool\'s configuration file \(default: \$HOME/\.pingctl/config\.yaml\)$`
	err := initPingFederateServices(context.Background(), "v1.2.3", false, false)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingOneServices function
func Test_initPingOneServices(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test the function
	err := initPingOneServices(context.Background(), "v1.2.3")
	testutils.CheckExpectedError(t, err, nil)

	// Check the API client is not nil
	if pingoneApiClient == nil {
		t.Errorf("initPingOneServices() apiClient = %v, want non-nil", pingoneApiClient)
	}

	// Check the API client ID is not empty
	if pingoneApiClientId == "" {
		t.Errorf("initPingOneServices() apiClientId = '%s', want non-empty", pingoneApiClientId)
	}

	// Check the API export environment ID is not empty
	if pingoneExportEnvID == "" {
		t.Errorf("initPingOneServices() exportEnvID = '%s', want non-empty", pingoneExportEnvID)
	}

	// Check the context is not nil
	if pingoneContext == nil {
		t.Errorf("initPingOneServices() context = %v, want non-nil", pingoneContext)
	}

	// Check the API client ID is a valid UUID
	if _, err := uuid.ParseUUID(pingoneApiClientId); err != nil {
		t.Errorf("initPingOneServices() api clientId = '%s', want valid UUID", pingoneApiClientId)
	}

	// Check the API export environment ID is a valid UUID
	if _, err := uuid.ParseUUID(pingoneExportEnvID); err != nil {
		t.Errorf("initPingOneServices() exportEnvID = '%s', want valid UUID", pingoneExportEnvID)
	}
}

// Test initPingOneServices function fails with no region
func Test_initPingOneServices_noPingOneRegion(t *testing.T) {
	testutils_viper.InitVipersCustomFile(t, `activeProfile: default
default:
    description: "default description"
    pingctl:
        color: true
        outputFormat: text
    pingone:
        export:
            environmentid: ""
        region: ""
        worker:
            clientid: ""
            clientsecret: ""
            environmentid: ""
    pingfederate:
        accesstokenauth:
            accesstoken: ""
        adminapipath: ""
        basicauth:
            password: ""
            username: ""
        clientcredentialsauth:
            clientid: ""
            clientsecret: ""
            scopes: ""
            tokenurl: ""
        httpshost: ""`)

	expectedErrorPattern := `^failed to initialize pingone API client\. unrecognized pingone region: ''. Must be one of:.*$`
	err := initPingOneServices(context.Background(), "v1.2.3")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingOneServices function fails with no client credentials
func Test_initPingOneServices_noPingOneClientCredentials(t *testing.T) {
	testutils_viper.InitVipersCustomFile(t, fmt.Sprintf(`activeProfile: default
default:
    description: "default description"
    pingctl:
        color: true
        outputFormat: text
    pingone:
        export:
            environmentid: ""
        region: "%s"
        worker:
            clientid: ""
            clientsecret: ""
            environmentid: ""
    pingfederate:
        accesstokenauth:
            accesstoken: ""
        adminapipath: ""
        basicauth:
            password: ""
            username: ""
        clientcredentialsauth:
            clientid: ""
            clientsecret: ""
            scopes: ""
            tokenurl: ""
        httpshost: ""`, os.Getenv(profiles.PingOneRegionOption.EnvVar)))

	expectedErrorPattern := `^failed to initialize pingone API client\. one of worker client ID, worker client secret, and/or worker environment ID is empty\. configure these properties via parameter flags, environment variables, or the tool's configuration file \(default: \$HOME/\.pingctl/config\.yaml\)$`
	err := initPingOneServices(context.Background(), "v1.2.3")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingFederateApiClient function
func Test_initPingFederateApiClient(t *testing.T) {
	testutils_viper.InitVipers(t)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, //#nosec G402 -- This is a test
		},
	}

	// Test the function
	err := initPingFederateApiClient(tr, "v1.2.3")
	testutils.CheckExpectedError(t, err, nil)

	// Check the API client is not nil
	if pingfederateApiClient == nil {
		t.Errorf("initPingFederateApiClient() apiClient = %v, want non-nil", pingfederateApiClient)
	}

	// Check the API client has a valid http client
	if pingfederateApiClient.GetConfig().HTTPClient == nil {
		t.Errorf("initPingFederateApiClient() httpClient = %v, want non-nil", pingfederateApiClient.GetConfig().HTTPClient)
	}

	// Check the API client has a valid http transport
	if pingfederateApiClient.GetConfig().HTTPClient.Transport == nil {
		t.Errorf("initPingFederateApiClient() httpTransport = %v, want non-nil", pingfederateApiClient.GetConfig().HTTPClient.Transport)
	}

	// Check the API client has a valid default "X-Xsrf-Header" header
	if pingfederateApiClient.GetConfig().DefaultHeader["X-Xsrf-Header"] != "PingFederate" {
		t.Errorf("initPingFederateApiClient() defaultHeader = %v, want 'PingFederate'", pingfederateApiClient.GetConfig().DefaultHeader["X-Xsrf-Header"])
	}

	// Check the API client Servers
	if len(pingfederateApiClient.GetConfig().Servers) == 0 {
		t.Errorf("initPingFederateApiClient() servers = %v, want non-empty", pingfederateApiClient.GetConfig().Servers)
	}

	if pingfederateApiClient.GetConfig().Servers[0].URL != os.Getenv(profiles.PingFederateHttpsHostOption.EnvVar)+os.Getenv(profiles.PingFederateAdminApiPathOption.EnvVar) {
		t.Errorf("initPingFederateApiClient() server URL = %v, want %v", pingfederateApiClient.GetConfig().Servers[0].URL, os.Getenv(profiles.PingFederateHttpsHostOption.EnvVar)+os.Getenv(profiles.PingFederateAdminApiPathOption.EnvVar))
	}
}

// Test initPingFederateApiClient function fails on missing https host
func Test_initPingFederateApiClient_noHttpsHost(t *testing.T) {
	testutils_viper.InitVipersCustomFile(t, `activeProfile: default
default:
    description: "default description"
    pingctl:
        color: true
        outputFormat: text
    pingone:
        export:
            environmentid: ""
        region: ""
        worker:
            clientid: ""
            clientsecret: ""
            environmentid: ""
    pingfederate:
        accesstokenauth:
            accesstoken: ""
        adminapipath: ""
        basicauth:
            password: ""
            username: ""
        clientcredentialsauth:
            clientid: ""
            clientsecret: ""
            scopes: ""
            tokenurl: ""
        httpshost: ""`)

	expectedErrorPattern := `^failed to initialize pingfederate API client\. the pingfederate https host configuration value is not set: configure this property via parameter flags, environment variables, or the tool's configuration file \(default: \$HOME/\.pingctl/config\.yaml\)$`
	err := initPingFederateApiClient(&http.Transport{}, "v1.2.3")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingFederateApiClient function fails on nil transport
func Test_initPingFederateApiClient_nilTransport(t *testing.T) {
	expectedErrorPattern := `^failed to initialize pingfederate API client\. http transport is nil$`
	err := initPingFederateApiClient(nil, "v1.2.3")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingOneApiClient function
func Test_initPingOneApiClient(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test the function
	err := initPingOneApiClient(context.Background(), "v1.2.3")
	testutils.CheckExpectedError(t, err, nil)

	// Check the API client is not nil
	if pingoneApiClient == nil {
		t.Errorf("initPingOneApiClient() apiClient = %v, want non-nil", pingoneApiClient)
	}

	if pingoneApiClient.ManagementAPIClient.GetConfig().HTTPClient == nil {
		t.Errorf("initPingOneApiClient() httpClient = %v, want non-nil", pingoneApiClient.ManagementAPIClient.GetConfig().HTTPClient)
	}

	// Check the API client id is not empty
	if pingoneApiClientId == "" {
		t.Errorf("initPingOneApiClient() api client id = '%s', want non-empty", pingoneApiClientId)
	}
}

// Test fixEmptyOutputDirVar function
func Test_fixEmptyOutputDirVar(t *testing.T) {
	// Test the function
	outputDir, err := fixEmptyOutputDirVar("")
	if err != nil {
		t.Errorf("fixEmptyOutputDirVar() error = %v", err)
	}
	if outputDir == "" {
		t.Errorf("fixEmptyOutputDirVar() outputDir = %v, want non-empty", outputDir)
	}
}

// Test fixEmptyOutputDirVar function with non-empty output directory
func Test_fixEmptyOutputDirVar_nonEmpty(t *testing.T) {
	// Test the function
	outputDir, err := fixEmptyOutputDirVar("test")
	if err != nil {
		t.Errorf("fixEmptyOutputDirVar() error = %v", err)
	}
	if outputDir != "test" {
		t.Errorf("fixEmptyOutputDirVar() outputDir = %v, want 'test'", outputDir)
	}
}

// Test createOrValidateOutputDir function
// - Empty directory that exists is valid, and should not return an error
func Test_createOrValidateOutputDir_emptyDir(t *testing.T) {
	// Create a directory in the temp directory
	outputDir := t.TempDir()

	// Test the function
	err := createOrValidateOutputDir(outputDir, false)
	testutils.CheckExpectedError(t, err, nil)

	err = createOrValidateOutputDir(outputDir, true)
	testutils.CheckExpectedError(t, err, nil)
}

// Test createOrValidateOutputDir function
// - Create an empty directory that does exist
// - Add a file to the directory
// - Validate that the function returns an error with overwrite set to false
func Test_createOrValidateOutputDir_existingDir(t *testing.T) {
	// Create a directory in the temp directory
	outputDir := t.TempDir()

	// Create a file in the new directory
	file := outputDir + "/file"
	if _, err := os.Create(file); err != nil {
		t.Fatalf("os.Create() error = %v", err)
	}

	// Test the function
	expectedErrorPattern := `^'platform export' output directory '.*' is not empty. Use --overwrite to overwrite existing export data$`
	err := createOrValidateOutputDir(outputDir, false)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)

	err = createOrValidateOutputDir(outputDir, true)
	testutils.CheckExpectedError(t, err, nil)
}

// Test createOrValidateOutputDir function
// - Provide function with a directory that does not exist
// - Validate that the function creates the directory
func Test_createOrValidateOutputDir_nonExistingDir(t *testing.T) {
	// Create a directory in the temp directory
	outputDir := t.TempDir() + "/new"

	// Test the function
	err := createOrValidateOutputDir(outputDir, false)
	testutils.CheckExpectedError(t, err, nil)

	// Check if the directory was created
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		t.Errorf("createOrValidateOutputDir() directory = %v, want exists", outputDir)
	}

	err = createOrValidateOutputDir(outputDir, true)
	testutils.CheckExpectedError(t, err, nil)
}

// Test getPingOneExportEnvID function
func Test_getPingOneExportEnvID(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test the function
	err := getPingOneExportEnvID()
	testutils.CheckExpectedError(t, err, nil)

	// Check the export environment ID is not empty
	if pingoneExportEnvID == "" {
		t.Errorf("getPingOneExportEnvID() pingoneExportEnvID = '%s', want non-empty", pingoneExportEnvID)
	}
}

// Test getPingOneExportEnvID function fails with no export environment ID
func Test_getPingOneExportEnvID_noExportEnvID(t *testing.T) {
	testutils_viper.InitVipersCustomFile(t, `activeProfile: default
default:
    description: "default description"
    pingctl:
        color: true
        outputFormat: text
    pingone:
        export:
            environmentid: ""
        region: ""
        worker:
            clientid: ""
            clientsecret: ""
            environmentid: ""
    pingfederate:
        accesstokenauth:
            accesstoken: ""
        adminapipath: ""
        basicauth:
            password: ""
            username: ""
        clientcredentialsauth:
            clientid: ""
            clientsecret: ""
            scopes: ""
            tokenurl: ""
        httpshost: ""`)

	expectedErrorPattern := `^failed to determine pingone export environment ID$`
	err := getPingOneExportEnvID()
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test validatePingOneExportEnvID function
func Test_validatePingOneExportEnvID(t *testing.T) {
	testutils_viper.InitVipers(t)

	// init the api client
	err := initPingOneApiClient(context.Background(), "v1.2.3")
	if err != nil {
		t.Fatalf("initPingOneApiClient() error = %v", err)
	}

	// Test the function
	err = validatePingOneExportEnvID(context.Background())
	testutils.CheckExpectedError(t, err, nil)
}

// Test validatePingOneExportEnvID function fails with invalid export environment ID
func Test_validatePingOneExportEnvID_invalidExportEnvID(t *testing.T) {
	testutils_viper.InitVipers(t)

	// init the api client
	err := initPingOneApiClient(context.Background(), "v1.2.3")
	if err != nil {
		t.Fatalf("initPingOneApiClient() error = %v", err)
	}

	// Set the export environment ID to an invalid value
	pingoneExportEnvID, err = uuid.GenerateUUID()
	if err != nil {
		t.Fatalf("uuid.GenerateUUID() error = %v", err)
	}

	expectedErrorPattern := `(?s)^.* Request for resource '.*' was not successful\..*Response Code: 404 Not Found.*Response Body: {{.*"id" : ".*",.*"code" : "NOT_FOUND",.*"message" : "Unable to find environment with ID: '.*'".*}}.*Error: 404 Not Found$`
	err = validatePingOneExportEnvID(context.Background())
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)

	// Reset the export environment ID
	pingoneExportEnvID = ""
}

// Test validatePingOneExportEnvID function fails with nil context
func Test_validatePingOneExportEnvID_nilContext(t *testing.T) {
	testutils_viper.InitVipers(t)

	if err := initPingOneServices(context.Background(), "v1.2.3"); err != nil {
		t.Fatalf("initPingOneServices() error = %v", err)
	}

	pingoneContext = nil

	expectedErrorPattern := `^failed to validate pingone environment ID '.*'. context is nil$`
	// nolint:staticcheck // ignore SA1012 this is a test
	err := validatePingOneExportEnvID(nil) //lint:ignore SA1012 this is a test
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test validatePingOneExportEnvID function fails with nil api client
func Test_validatePingOneExportEnvID_nilApiClient(t *testing.T) {
	testutils_viper.InitVipers(t)

	if err := initPingOneServices(context.Background(), "v1.2.3"); err != nil {
		t.Fatalf("initPingOneServices() error = %v", err)
	}

	pingoneApiClient = nil

	expectedErrorPattern := `^failed to validate pingone environment ID '.*'. apiClient is nil$`
	err := validatePingOneExportEnvID(context.Background())
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test getExportableConnectors function
func Test_getExportableConnectors(t *testing.T) {
	// Test the function
	connectors := getExportableConnectors(customtypes.NewMultiService())
	if connectors == nil {
		t.Errorf("getExportableConnectors() connectors = %v, want non-nil", connectors)
	}

	// Check the number of connectors
	if len(*connectors) == 0 {
		t.Errorf("getExportableConnectors() num connectors = %v, want non-zero", len(*connectors))
	}

	// Check the connectors are not nil
	for _, connector := range *connectors {
		if connector == nil {
			t.Errorf("getExportableConnectors() connector = %v, want non-nil", connector)
		}
	}

	if len(customtypes.MultiServiceValidValues()) != len(*connectors) {
		t.Errorf("getExportableConnectors() num connectors = %v, want %v", len(*connectors), len(customtypes.MultiServiceValidValues()))
	}
}

// Test getExportableConnectors function with nil multiService
func Test_getExportableConnectors_nilMultiService(t *testing.T) {
	// Test the function
	connectors := getExportableConnectors(nil)
	if len(*connectors) != 0 {
		t.Errorf("getExportableConnectors() num connectors = %v, want 0", len(*connectors))
	}
}

// Test getExportableConnectors function with one service
func Test_getExportableConnectors_oneService(t *testing.T) {
	ms := customtypes.NewMultiService()
	if err := ms.Set(customtypes.ENUM_SERVICE_PINGFEDERATE); err != nil {
		t.Fatalf("ms.Set() error = %v", err)
	}

	// Test the function
	connectors := getExportableConnectors(ms)
	if len(*connectors) != 1 {
		t.Errorf("getExportableConnectors() num connectors = %v, want 1", len(*connectors))
	}

	// Make sure the connector is not nil
	if (*connectors)[0] == nil {
		t.Errorf("getExportableConnectors() connector = %v, want non-nil", (*connectors)[0])
	}

	// Make sure the connector is the correct type
	_, ok := (*connectors)[0].(*pingfederate.PingfederateConnector)
	if !ok {
		t.Errorf("getExportableConnectors() connector = %v, want PingfederateConnector", (*connectors)[0])
	}
}

// Test exportConnectors function
func Test_exportConnectors(t *testing.T) {
	testutils_viper.InitVipers(t)

	// init the pingfederate services
	err := initPingFederateServices(context.Background(), "v1.2.3", false, false)
	if err != nil {
		t.Fatalf("initPingFederateServices() error = %v", err)
	}

	ms := customtypes.NewMultiService()
	if err := ms.Set(customtypes.ENUM_SERVICE_PINGFEDERATE); err != nil {
		t.Fatalf("ms.Set() error = %v", err)
	}

	exportableConnectors := getExportableConnectors(ms)

	// Test the function
	err = exportConnectors(exportableConnectors, connector.ENUMEXPORTFORMAT_HCL, t.TempDir(), true)
	testutils.CheckExpectedError(t, err, nil)
}

// Test exportConnectors function fails with nil connectors
func Test_exportConnectors_nilConnectors(t *testing.T) {
	// Test the function
	expectedErrorPattern := `^failed to export services. exportable connectors list is nil$`
	err := exportConnectors(nil, connector.ENUMEXPORTFORMAT_HCL, t.TempDir(), true)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test exportConnectors function fails with invalid export format
func Test_exportConnectors_invalidExportFormat(t *testing.T) {
	testutils_viper.InitVipers(t)

	// init the pingfederate services
	err := initPingFederateServices(context.Background(), "v1.2.3", false, false)
	if err != nil {
		t.Fatalf("initPingFederateServices() error = %v", err)
	}

	ms := customtypes.NewMultiService()
	if err := ms.Set(customtypes.ENUM_SERVICE_PINGFEDERATE); err != nil {
		t.Fatalf("ms.Set() error = %v", err)
	}

	exportableConnectors := getExportableConnectors(ms)

	// Test the function
	expectedErrorPattern := `^failed to export '.*' service: unrecognized export format "invalid". Must be one of: \[.*\]$`
	err = exportConnectors(exportableConnectors, "invalid", t.TempDir(), true)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test exportConnectors function fails with invalid output directory
func Test_exportConnectors_invalidOutputDir(t *testing.T) {
	testutils_viper.InitVipers(t)

	// init the pingfederate services
	err := initPingFederateServices(context.Background(), "v1.2.3", false, false)
	if err != nil {
		t.Fatalf("initPingFederateServices() error = %v", err)
	}

	ms := customtypes.NewMultiService()
	if err := ms.Set(customtypes.ENUM_SERVICE_PINGFEDERATE); err != nil {
		t.Fatalf("ms.Set() error = %v", err)
	}

	exportableConnectors := getExportableConnectors(ms)

	// Test the function
	expectedErrorPattern := `^failed to export '.*' service: failed to create export file ".*"\. err: open /invalid/.*\.tf: no such file or directory$`
	err = exportConnectors(exportableConnectors, connector.ENUMEXPORTFORMAT_HCL, "/invalid", true)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
