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
	"github.com/pingidentity/pingctl/internal/configuration"
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

func RunInternalExport(ctx context.Context, commandVersion string) (err error) {
	if ctx == nil {
		return fmt.Errorf("failed to run 'platform export' command. context is nil")
	}

	exportFormat, err := profiles.GetOptionValue(configuration.PlatformExportExportFormatOption)
	if err != nil {
		return err
	}
	multiService, err := profiles.GetOptionValue(configuration.PlatformExportServiceOption)
	if err != nil {
		return err
	}
	outputDir, err := profiles.GetOptionValue(configuration.PlatformExportOutputDirectoryOption)
	if err != nil {
		return err
	}
	overwriteExport, err := profiles.GetOptionValue(configuration.PlatformExportOverwriteOption)
	if err != nil {
		return err
	}

	ms := customtypes.NewMultiService()
	if err = ms.Set(multiService); err != nil {
		return err
	}

	if ms.ContainsPingOneService() {
		if err = initPingOneServices(ctx, commandVersion); err != nil {
			return err
		}
	}

	if ms.ContainsPingFederateService() {
		if err = initPingFederateServices(ctx, commandVersion); err != nil {
			return err
		}
	}

	overwriteExportBool, err := strconv.ParseBool(overwriteExport)
	if err != nil {
		return err
	}
	if err := createOrValidateOutputDir(outputDir, overwriteExportBool); err != nil {
		return err
	}

	exportableConnectors := getExportableConnectors(ms)

	if err := exportConnectors(exportableConnectors, exportFormat, outputDir, overwriteExportBool); err != nil {
		return err
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Export to directory '%s' complete.", outputDir),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}

func initPingFederateServices(ctx context.Context, pingctlVersion string) (err error) {
	if ctx == nil {
		return fmt.Errorf("failed to initialize PingFederate services. context is nil")
	}

	// Get all the PingFederate configuration values
	pfClientID, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateClientIDOption)
	if err != nil {
		return err
	}
	pfClientSecret, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateClientSecretOption)
	if err != nil {
		return err
	}
	pfTokenUrl, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateTokenURLOption)
	if err != nil {
		return err
	}
	pfScopes, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateScopesOption)
	if err != nil {
		return err
	}
	pfAccessToken, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateAccessTokenOption)
	if err != nil {
		return err
	}
	pfUsername, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateUsernameOption)
	if err != nil {
		return err
	}
	pfPassword, err := profiles.GetOptionValue(configuration.PlatformExportPingfederatePasswordOption)
	if err != nil {
		return err
	}
	pfInsecureTrustAllTLS, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateInsecureTrustAllTLSOption)
	if err != nil {
		return err
	}
	caCertPemFiles, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateCACertificatePemFilesOption)
	if err != nil {
		return err
	}

	caCertPool := x509.NewCertPool()
	for _, caCertPemFile := range strings.Split(caCertPemFiles, ",") {
		if caCertPemFile == "" {
			continue
		}
		caCertPemFile := filepath.Clean(caCertPemFile)
		caCert, err := os.ReadFile(caCertPemFile)
		if err != nil {
			return fmt.Errorf("failed to read CA certificate PEM file '%s': %v", caCertPemFile, err)
		}

		ok := caCertPool.AppendCertsFromPEM(caCert)
		if !ok {
			return fmt.Errorf("failed to parse CA certificate PEM file '%s' to certificate pool", caCertPemFile)
		}
	}

	pfInsecureTrustAllTLSBool, err := strconv.ParseBool(pfInsecureTrustAllTLS)
	if err != nil {
		return err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: pfInsecureTrustAllTLSBool, //#nosec G402 -- This is defined by the user (default false), and warned as inappropriate in production.
			RootCAs:            caCertPool,
		},
	}

	if err = initPingFederateApiClient(tr, pingctlVersion); err != nil {
		return err
	}

	switch {
	case configuration.PlatformExportPingfederateUsernameOption.Flag.Changed && configuration.PlatformExportPingfederatePasswordOption.Flag.Changed:
		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextBasicAuth, pingfederateGoClient.BasicAuth{
			UserName: pfUsername,
			Password: pfPassword,
		})
	case configuration.PlatformExportPingfederateAccessTokenOption.Flag.Changed:
		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextAccessToken, pfAccessToken)
	case pfClientID != "" && pfClientSecret != "" && pfTokenUrl != "":
		pingfederateContext = context.WithValue(ctx, pingfederateGoClient.ContextOAuth2, pingfederateGoClient.OAuthValues{
			Transport:    tr,
			TokenUrl:     pfTokenUrl,
			ClientId:     pfClientID,
			ClientSecret: pfClientSecret,
			Scopes:       strings.Split(pfScopes, ","),
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

func initPingFederateApiClient(tr *http.Transport, pingctlVersion string) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Initializing PingFederate API client.")

	if tr == nil {
		return fmt.Errorf("failed to initialize pingfederate API client. http transport is nil")
	}

	httpsHost, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateHTTPSHostOption)
	if err != nil {
		return err
	}
	adminApiPath, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateAdminAPIPathOption)
	if err != nil {
		return err
	}
	xBypassExternalValidationHeader, err := profiles.GetOptionValue(configuration.PlatformExportPingfederateXBypassExternalValidationHeaderOption)
	if err != nil {
		return err
	}

	// default adminApiPath to /pf-admin-api/v1 if not set
	if adminApiPath == "" {
		adminApiPath = "/pf-admin-api/v1"
	}

	if httpsHost == "" {
		return fmt.Errorf(`failed to initialize pingfederate API client. the pingfederate https host configuration value is not set: configure this property via parameter flags, environment variables, or the tool's configuration file (default: $HOME/.pingctl/config.yaml)`)
	}

	userAgent := fmt.Sprintf("pingctl/%s", pingctlVersion)

	if v := strings.TrimSpace(os.Getenv("PINGCTL_PINGFEDERATE_APPEND_USER_AGENT")); v != "" {
		userAgent += fmt.Sprintf(" %s", v)
	}

	pfClientConfig := pingfederateGoClient.NewConfiguration()
	pfClientConfig.UserAgentSuffix = &userAgent
	pfClientConfig.DefaultHeader["X-Xsrf-Header"] = "PingFederate"
	pfClientConfig.DefaultHeader["X-BypassExternalValidation"] = xBypassExternalValidationHeader
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

func initPingOneApiClient(ctx context.Context, pingctlVersion string) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Initializing PingOne API client.")

	if ctx == nil {
		return fmt.Errorf("failed to initialize pingone API client. context is nil")
	}

	pingoneApiClientId, err = profiles.GetOptionValue(configuration.PlatformExportPingoneWorkerClientIDOption)
	if err != nil {
		return err
	}
	clientSecret, err := profiles.GetOptionValue(configuration.PlatformExportPingoneWorkerClientSecretOption)
	if err != nil {
		return err
	}
	environmentID, err := profiles.GetOptionValue(configuration.PlatformExportPingoneWorkerEnvironmentIDOption)
	if err != nil {
		return err
	}
	region, err := profiles.GetOptionValue(configuration.PlatformExportPingoneRegionOption)
	if err != nil {
		return err
	}

	if pingoneApiClientId == "" || clientSecret == "" || environmentID == "" || region == "" {
		return fmt.Errorf(`failed to initialize pingone API client. one of worker client ID, worker client secret, pingone region, and/or worker environment ID is empty. configure these properties via parameter flags, environment variables, or the tool's configuration file (default: $HOME/.pingctl/config.yaml)`)
	}

	userAgent := fmt.Sprintf("pingctl/%s", pingctlVersion)

	if v := strings.TrimSpace(os.Getenv("PINGCTL_PINGONE_APPEND_USER_AGENT")); v != "" {
		userAgent += fmt.Sprintf(" %s", v)
	}

	apiConfig := &pingoneGoClient.Config{
		ClientID:        &pingoneApiClientId,
		ClientSecret:    &clientSecret,
		EnvironmentID:   &environmentID,
		Region:          region,
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
	pingoneExportEnvID, err = profiles.GetOptionValue(configuration.PlatformExportPingoneExportEnvironmentIDOption)
	if err != nil {
		return err
	}

	if pingoneExportEnvID == "" {
		pingoneExportEnvID, err = profiles.GetOptionValue(configuration.PlatformExportPingoneWorkerEnvironmentIDOption)
		if err != nil {
			return err
		}
		if pingoneExportEnvID == "" {
			return fmt.Errorf("failed to determine pingone export environment ID.")
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

	for _, service := range multiService.GetServices() {
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

		err := connector.Export(exportFormat, outputDir, overwriteExport)
		if err != nil {
			return fmt.Errorf("failed to export '%s' service: %s", connector.ConnectorServiceName(), err.Error())
		}
	}

	return nil
}
