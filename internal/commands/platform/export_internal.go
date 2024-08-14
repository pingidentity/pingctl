package platform_internal

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	pingoneGoClient "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate"
	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform"
	"github.com/pingidentity/pingctl/internal/connector/pingone/protect"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
	"github.com/rs/zerolog"
)

var (
	pingfederateApiClient *pingfederateGoClient.APIClient
	pingfederateContext   context.Context

	pingoneApiClient   *pingoneGoClient.Client
	pingoneApiClientId string
	pingoneExportEnvID string
	pingoneContext     context.Context
)

func RunInternalExport(ctx context.Context, commandVersion string, outputDir, exportFormat string, overwriteExport bool, multiService *customtypes.MultiService, basicAuthFlagsUsed, AccessTokenAuthFlagsUsed bool) (err error) {
	if ctx == nil {
		return fmt.Errorf("failed to run 'platform export' command. context is nil")
	}

	if multiService.ContainsPingOneService() {
		if err = initPingOneServices(ctx, commandVersion); err != nil {
			return err
		}
	}

	if multiService.ContainsPingFederateService() {
		if err = initPingFederateServices(ctx, basicAuthFlagsUsed, AccessTokenAuthFlagsUsed); err != nil {
			return err
		}
	}

	outputDir, err = fixEmptyOutputDirVar(outputDir)
	if err != nil {
		return err
	}

	if err := createOrValidateOutputDir(outputDir, overwriteExport); err != nil {
		return err
	}

	exportableConnectors := getExportableConnectors(multiService)

	if err := exportConnectors(exportableConnectors, exportFormat, outputDir, overwriteExport); err != nil {
		return err
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Export to directory '%s' complete.", outputDir),
		Result:  output.ENUM_RESULT_SUCCESS,
	})
	return nil
}

func initPingFederateServices(ctx context.Context, basicAuthFlagsUsed, accessTokenAuthFlagsUsed bool) (err error) {
	// Get all the PingFederate configuration values
	profileViper := profiles.GetProfileViper()
	pfClientID := profileViper.GetString(profiles.PingFederateClientIDOption.ViperKey)
	pfClientSecret := profileViper.GetString(profiles.PingFederateClientSecretOption.ViperKey)
	pfTokenUrl := profileViper.GetString(profiles.PingFederateTokenURLOption.ViperKey)
	pfScopes := profileViper.GetStringSlice(profiles.PingFederateScopesOption.ViperKey)
	pfAccessToken := profileViper.GetString(profiles.PingFederateAccessTokenOption.ViperKey)
	pfUsername := profileViper.GetString(profiles.PingFederateUsernameOption.ViperKey)
	pfPassword := profileViper.GetString(profiles.PingFederatePasswordOption.ViperKey)
	pfInsecureTrustAllTLS := profileViper.GetBool(profiles.PingFederateInsecureTrustAllTLSOption.ViperKey)
	caCertPemFiles := profileViper.GetStringSlice(profiles.PingFederateCACertificatePemFilesOption.ViperKey)

	caCertPool := x509.NewCertPool()
	for _, caCertPemFile := range caCertPemFiles {
		caCertPemFile := filepath.Clean(caCertPemFile)
		caCert, err := os.ReadFile(caCertPemFile)
		if err != nil {
			return fmt.Errorf("failed to read CA certificate PEM file '%s': %s", caCertPemFile, err.Error())
		}

		ok := caCertPool.AppendCertsFromPEM(caCert)
		if !ok {
			return fmt.Errorf("failed to parse CA certificate PEM file '%s' to certificate pool", caCertPemFile)
		}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: pfInsecureTrustAllTLS, //#nosec G402 -- This is defined by the user (default false), and warned as inappropriate in production.
			RootCAs:            caCertPool,
		},
	}

	if err = initPingFederateApiClient(tr); err != nil {
		return err
	}

	switch {
	case basicAuthFlagsUsed && pfUsername != "" && pfPassword != "":
		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextBasicAuth, pingfederateGoClient.BasicAuth{
			UserName: pfUsername,
			Password: pfPassword,
		})
	case accessTokenAuthFlagsUsed && pfAccessToken != "":
		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextAccessToken, pfAccessToken)
	case pfClientID != "" && pfClientSecret != "" && pfTokenUrl != "":
		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextOAuth2, pingfederateGoClient.OAuthValues{
			Transport:    tr,
			TokenUrl:     pfTokenUrl,
			ClientId:     pfClientID,
			ClientSecret: pfClientSecret,
			Scopes:       pfScopes,
		})
	case pfAccessToken != "":
		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextAccessToken, pfAccessToken)
	case pfUsername != "" && pfPassword != "":
		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextBasicAuth, pingfederateGoClient.BasicAuth{
			UserName: pfUsername,
			Password: pfPassword,
		})
	default:
		return fmt.Errorf(`failed to initialize PingFederate API client. none of the following sets of authentication configuration values are set: OAuth2 client credentials (client ID, client secret, token URL), Access token, or Basic Authentication credentials (username, password). configure these properties via parameter flags, environment variables, or the tool's configuration file (default: $HOME/.pingctl/config.yaml)`)
	}

	return nil
}

func initPingOneServices(ctx context.Context, cmdVersion string) (err error) {
	if err = initPingOneApiClient(ctx, cmdVersion); err != nil {
		return err
	}

	if err = getPingOneExportEnvID(); err != nil {
		return err
	}

	if err := validatePingOneExportEnvID(ctx); err != nil {
		return err
	}

	pingoneContext = ctx

	return nil
}

func initPingFederateApiClient(tr *http.Transport) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Initializing PingFederate API client.")

	if tr == nil {
		return fmt.Errorf("failed to initialize pingfederate API client. http transport is nil")
	}

	profileViper := profiles.GetProfileViper()
	httpsHost := profileViper.GetString(profiles.PingFederateHttpsHostOption.ViperKey)
	adminApiPath := profileViper.GetString(profiles.PingFederateAdminApiPathOption.ViperKey)
	xBypassExternalValidationHeader := profileViper.GetBool(profiles.PingFederateXBypassExternalValidationHeaderOption.ViperKey)

	// default adminApiPath to /pf-admin-api/v1 if not set
	if adminApiPath == "" {
		adminApiPath = "/pf-admin-api/v1"
	}

	if httpsHost == "" {
		return fmt.Errorf(`failed to initialize pingfederate API client. the pingfederate https host configuration value is not set: configure this property via parameter flags, environment variables, or the tool's configuration file (default: $HOME/.pingctl/config.yaml)`)
	}

	pfClientConfig := pingfederateGoClient.NewConfiguration()
	pfClientConfig.DefaultHeader["X-Xsrf-Header"] = "PingFederate"
	pfClientConfig.DefaultHeader["X-BypassExternalValidation"] = strconv.FormatBool(xBypassExternalValidationHeader)
	pfClientConfig.Servers = pingfederateGoClient.ServerConfigurations{
		{
			URL: httpsHost + adminApiPath,
		},
	}
	httpClient := &http.Client{Transport: tr}
	pfClientConfig.HTTPClient = httpClient

	pingfederateApiClient = pingfederateGoClient.NewAPIClient(pfClientConfig)

	return nil
}

func initPingOneApiClient(ctx context.Context, version string) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Initializing PingOne API client.")

	profileViper := profiles.GetProfileViper()

	// Make sure the API client can be initialized with the required parameters
	if !profileViper.IsSet(profiles.PingOneWorkerEnvironmentIDOption.ViperKey) ||
		!profileViper.IsSet(profiles.PingOneRegionOption.ViperKey) ||
		!profileViper.IsSet(profiles.PingOneWorkerClientIDOption.ViperKey) ||
		!profileViper.IsSet(profiles.PingOneWorkerClientSecretOption.ViperKey) {
		return fmt.Errorf(`failed to initialize pingone API client. one of worker environment ID, worker client ID, worker client secret, and/or pingone region is not set. configure these properties via parameter flags, environment variables, or the tool's configuration file (default: $HOME/.pingctl/config.yaml)`)
	}

	pingoneApiClientId = profileViper.GetString(profiles.PingOneWorkerClientIDOption.ViperKey)
	clientSecret := profileViper.GetString(profiles.PingOneWorkerClientSecretOption.ViperKey)
	environmentID := profileViper.GetString(profiles.PingOneWorkerEnvironmentIDOption.ViperKey)
	region := profileViper.Get(profiles.PingOneRegionOption.ViperKey)

	var regionStr string
	switch regionVal := region.(type) {
	case string:
		regionStr = regionVal
	case customtypes.PingOneRegion:
		regionStr = string(regionVal)
	default:
		return fmt.Errorf("failed to initialize pingone API client. unrecognized pingone region variable type: %T", region)
	}

	switch regionStr {
	case customtypes.ENUM_PINGONE_REGION_AP, customtypes.ENUM_PINGONE_REGION_CA, customtypes.ENUM_PINGONE_REGION_EU, customtypes.ENUM_PINGONE_REGION_NA:
		l.Debug().Msgf("pingone region '%s' validated.", regionStr)
	default:
		return fmt.Errorf("failed to initialize pingone API client. unrecognized pingone region: '%s'. Must be one of: %s", regionStr, strings.Join(customtypes.PingOneRegionValidValues(), ", "))
	}

	// Make sure the client credentials are not empty
	if pingoneApiClientId == "" || clientSecret == "" || environmentID == "" {
		return fmt.Errorf(`failed to initialize pingone API client. one of worker client ID, worker client secret, and/or worker environment ID is empty. configure these properties via parameter flags, environment variables, or the tool's configuration file (default: $HOME/.pingctl/config.yaml)`)
	}

	userAgent := fmt.Sprintf("pingctl/%s", version)

	if v := strings.TrimSpace(os.Getenv("PINGCTL_PINGONE_APPEND_USER_AGENT")); v != "" {
		userAgent += fmt.Sprintf(" %s", v)
	}

	apiConfig := &pingoneGoClient.Config{
		ClientID:        &pingoneApiClientId,
		ClientSecret:    &clientSecret,
		EnvironmentID:   &environmentID,
		Region:          regionStr,
		UserAgentSuffix: &userAgent,
	}

	pingoneApiClient, err = apiConfig.APIClient(ctx)
	if err != nil {
		// If logging level is DEBUG or TRACE, include the worker client secret in the error message
		var clientSecretErrorMessage string
		if l.GetLevel() <= zerolog.DebugLevel {
			clientSecretErrorMessage = fmt.Sprintf("worker client secret - %s", clientSecret)
		} else {
			clientSecretErrorMessage = "worker client secret - (use DEBUG or TRACE logging level to view the client secret in the error message)" // #nosec G101
		}
		initFailErrFormatMessage := `failed to initialize pingone API client.
%s

configuration values used for client initialization:
worker client ID - %s
worker environment ID - %s
pingone region - %s
%s`
		return fmt.Errorf(initFailErrFormatMessage, err.Error(), pingoneApiClientId, environmentID, region, clientSecretErrorMessage)
	}

	return nil
}

func fixEmptyOutputDirVar(outputDir string) (newOutputDir string, err error) {
	if outputDir == "" {
		// Default the outputDir variable to the user's present working directory.
		outputDir, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to determine user's present working directory: %s", err.Error())
		}

		// Append "export" to the output directory as export needs an empty directory to write to
		outputDir = filepath.Join(outputDir, "export")

		output.Print(output.Opts{
			Message: fmt.Sprintf("Defaulting 'platform export' command output directory to '%s'", outputDir),
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})
	}

	return outputDir, nil
}

func createOrValidateOutputDir(outputDir string, overwriteExport bool) (err error) {
	l := logger.Get()

	// Check if outputDir exists
	// If not, create the directory
	l.Debug().Msgf("Validating export output directory '%s'", outputDir)
	_, err = os.Stat(outputDir)
	if err != nil {
		output.Print(output.Opts{
			Message: fmt.Sprintf("failed to find 'platform export' output directory. creating new output directory at filepath '%s'", outputDir),
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})

		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create 'platform export' output directory '%s': %s", outputDir, err.Error())
		}

		output.Print(output.Opts{
			Message: fmt.Sprintf("new 'platform export' output directory '%s' created", outputDir),
			Result:  output.ENUM_RESULT_SUCCESS,
		})
	} else {
		// Check if the output directory is empty
		// If not, default behavior is to exit and not overwrite.
		// This can be changed with the --overwrite export parameter
		if !overwriteExport {
			dirEntries, err := os.ReadDir(outputDir)
			if err != nil {
				return fmt.Errorf("failed to read contents of 'platform export' output directory '%s': %s", outputDir, err.Error())
			}

			if len(dirEntries) > 0 {
				return fmt.Errorf("'platform export' output directory '%s' is not empty. Use --overwrite to overwrite existing export data", outputDir)
			}
		}
	}

	return nil
}

func getPingOneExportEnvID() (err error) {
	profileViper := profiles.GetProfileViper()

	// Find the env ID to export. Default to worker env id if not provided by user.
	pingoneExportEnvID = profileViper.GetString(profiles.PingOneExportEnvironmentIDOption.ViperKey)
	if pingoneExportEnvID == "" {
		pingoneExportEnvID = profileViper.GetString(profiles.PingOneWorkerEnvironmentIDOption.ViperKey)

		// if the exportEnvID is still empty, this is a problem. Return error.
		if pingoneExportEnvID == "" {
			return fmt.Errorf("failed to determine pingone export environment ID")
		}

		output.Print(output.Opts{
			Message: "No target PingOne export environment ID specified. Defaulting export environment ID to the Worker App environment ID.",
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})
	}

	return nil
}

func validatePingOneExportEnvID(ctx context.Context) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Validating export environment ID...")

	if ctx == nil {
		return fmt.Errorf("failed to validate pingone environment ID '%s'. context is nil", pingoneExportEnvID)
	}

	if pingoneApiClient == nil {
		return fmt.Errorf("failed to validate pingone environment ID '%s'. apiClient is nil", pingoneExportEnvID)
	}

	environment, response, err := pingoneApiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(ctx, pingoneExportEnvID).Execute()
	err = common.HandleClientResponse(response, err, "ReadOneEnvironment", "pingone_environment")
	if err != nil {
		return err
	}

	if environment == nil {
		return fmt.Errorf("failed to validate pingone environment ID '%s'. environment matching ID does not exist", pingoneExportEnvID)
	}

	return nil
}

func getExportableConnectors(multiService *customtypes.MultiService) (exportableConnectors *[]connector.Exportable) {
	// Using the --service parameter(s) provided by user, build list of connectors to export
	connectors := []connector.Exportable{}

	if multiService == nil {
		return &connectors
	}

	for _, service := range *multiService.GetServices() {
		switch service {
		case customtypes.ENUM_SERVICE_PINGONE_PLATFORM:
			connectors = append(connectors, platform.PlatformConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_SERVICE_PINGONE_SSO:
			connectors = append(connectors, sso.SSOConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_SERVICE_PINGONE_MFA:
			connectors = append(connectors, mfa.MFAConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_SERVICE_PINGONE_PROTECT:
			connectors = append(connectors, protect.ProtectConnector(pingoneContext, pingoneApiClient, &pingoneApiClientId, pingoneExportEnvID))
		case customtypes.ENUM_SERVICE_PINGFEDERATE:
			connectors = append(connectors, pingfederate.PFConnector(pingfederateContext, pingfederateApiClient))
			// default:
			// This unrecognized service condition is handled by cobra with the custom type MultiService
		}
	}

	return &connectors
}

func exportConnectors(exportableConnectors *[]connector.Exportable, exportFormat, outputDir string, overwriteExport bool) (err error) {
	if exportableConnectors == nil {
		return fmt.Errorf("failed to export services. exportable connectors list is nil")
	}

	// Loop through user defined exportable connectors and export them
	for _, connector := range *exportableConnectors {
		output.Print(output.Opts{
			Message: fmt.Sprintf("Exporting %s service...", connector.ConnectorServiceName()),
			Result:  output.ENUM_RESULT_NIL,
		})

		err := connector.Export(string(exportFormat), outputDir, overwriteExport)
		if err != nil {
			return fmt.Errorf("failed to export '%s' service: %s", connector.ConnectorServiceName(), err.Error())
		}
	}

	return nil
}
