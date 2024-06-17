package platform_internal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform"
	"github.com/pingidentity/pingctl/internal/connector/pingone/protect"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunInternalExport(cmd *cobra.Command, outputDir, exportFormat string, overwriteExport bool, multiService *customtypes.MultiService) (err error) {
	apiClient, apiClientId, err := initApiClient(cmd.Context(), cmd.Root().Version)
	if err != nil {
		return err
	}

	outputDir, err = fixEmptyOutputDirVar(outputDir)
	if err != nil {
		return err
	}

	if err := createOrValidateOutputDir(outputDir, overwriteExport); err != nil {
		return err
	}

	exportEnvID, err := getExportEnvID()
	if err != nil {
		return err
	}

	if err := validateExportEnvID(cmd.Context(), exportEnvID, apiClient); err != nil {
		return err
	}

	exportableConnectors := getExportableConnectors(exportEnvID, apiClientId, cmd.Context(), multiService, apiClient)

	if err := exportConnectors(exportableConnectors, exportFormat, outputDir, overwriteExport); err != nil {
		return err
	}

	output.Format(output.CommandOutput{
		Message: fmt.Sprintf("Export to directory '%s' complete.", outputDir),
		Result:  output.ENUMCOMMANDOUTPUTRESULT_SUCCESS,
	})
	return nil
}

func initApiClient(ctx context.Context, version string) (apiClient *sdk.Client, apiClientId string, err error) {
	l := logger.Get()
	l.Debug().Msgf("Initialising API client..")

	// Make sure the API client can be initialized with the required parameters
	if !viper.IsSet(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerClientIdParamName].ViperConfigKey) ||
		!viper.IsSet(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerClientSecretParamName].ViperConfigKey) ||
		!viper.IsSet(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerEnvironmentIdParamName].ViperConfigKey) ||
		!viper.IsSet(viperconfig.ConfigOptions[viperconfig.ExportPingoneRegionParamName].ViperConfigKey) {
		return nil, "", fmt.Errorf(`failed to initialize pingone API client.
		one of worker environment ID, worker client ID, worker client secret,
		and/or pingone region is not set.
		configure these properties via parameter flags, environment variables,
		or the tool's configuration file (default: $HOME/.pingctl/config.yaml)`)
	}

	apiClientId = viper.GetString(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerClientIdParamName].ViperConfigKey)
	clientSecret := viper.GetString(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerClientSecretParamName].ViperConfigKey)
	environmentID := viper.GetString(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerEnvironmentIdParamName].ViperConfigKey)
	region := viper.Get(viperconfig.ConfigOptions[viperconfig.ExportPingoneRegionParamName].ViperConfigKey)

	var regionStr string
	switch regionVal := region.(type) {
	case string:
		regionStr = regionVal
	case customtypes.PingOneRegion:
		regionStr = string(regionVal)
	default:
		return nil, "", fmt.Errorf("failed to initialize pingone API client. unrecognized pingone region variable type: %T", region)
	}

	switch regionStr {
	case customtypes.ENUM_PINGONE_REGION_AP, customtypes.ENUM_PINGONE_REGION_CA, customtypes.ENUM_PINGONE_REGION_EU, customtypes.ENUM_PINGONE_REGION_NA:
		l.Debug().Msgf("pingone region '%s' validated.", regionStr)
	default:
		return nil, "", fmt.Errorf("failed to initialize pingone API client. unrecognized pingone region: '%s'. Must be one of: %s", regionStr, strings.Join(customtypes.PingOneRegionValidValues(), ", "))
	}

	userAgent := fmt.Sprintf("pingctl/%s", version)

	if v := strings.TrimSpace(os.Getenv("PINGCTL_PINGONE_APPEND_USER_AGENT")); v != "" {
		userAgent += fmt.Sprintf(" %s", v)
	}

	apiConfig := &sdk.Config{
		ClientID:        &apiClientId,
		ClientSecret:    &clientSecret,
		EnvironmentID:   &environmentID,
		Region:          regionStr,
		UserAgentSuffix: &userAgent,
	}

	apiClient, err = apiConfig.APIClient(ctx)
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
		return nil, "", fmt.Errorf(initFailErrFormatMessage, err.Error(), apiClientId, environmentID, region, clientSecretErrorMessage)
	}

	return apiClient, apiClientId, nil
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

		output.Format(output.CommandOutput{
			Message: fmt.Sprintf("Defaulting 'platform export' command output directory to '%s'", outputDir),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
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
		output.Format(output.CommandOutput{
			Message: fmt.Sprintf("failed to find 'platform export' output directory. creating new output directory at filepath '%s'", outputDir),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})

		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create 'platform export' output directory '%s': %s", outputDir, err.Error())
		}

		output.Format(output.CommandOutput{
			Message: fmt.Sprintf("new 'platform export' output directory '%s' created", outputDir),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_SUCCESS,
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

func getExportEnvID() (exportEnvID string, err error) {
	// Find the env ID to export. Default to worker env id if not provided by user.
	exportEnvID = viper.GetString(viperconfig.ConfigOptions[viperconfig.ExportPingoneExportEnvironmentIdParamName].ViperConfigKey)
	if exportEnvID == "" {
		exportEnvID = viper.GetString(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerEnvironmentIdParamName].ViperConfigKey)

		// if the exportEnvID is still empty, this is a problem. Return error.
		if exportEnvID == "" {
			return "", fmt.Errorf("failed to determine export environment ID")
		}

		output.Format(output.CommandOutput{
			Message: "No target export environment ID specified. Defaulting export environment ID to the Worker App environment ID.",
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})
	}

	return exportEnvID, nil
}

func validateExportEnvID(ctx context.Context, exportEnvID string, apiClient *sdk.Client) (err error) {
	l := logger.Get()
	l.Debug().Msgf("Validating export environment ID...")

	if ctx == nil {
		return fmt.Errorf("failed to validate environment ID '%s'. context is nil", exportEnvID)
	}

	if apiClient == nil {
		return fmt.Errorf("failed to validate environment ID '%s'. apiClient is nil", exportEnvID)
	}

	environment, response, err := apiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(ctx, exportEnvID).Execute()
	err = common.HandleClientResponse(response, err, "ReadOneEnvironment", "pingone_environment")
	if err != nil {
		return err
	}

	if environment == nil {
		return fmt.Errorf("failed to validate environment ID '%s'. environment matching ID does not exist", exportEnvID)
	}

	return nil
}

func getExportableConnectors(exportEnvID, apiClientId string, ctx context.Context, multiService *customtypes.MultiService, apiClient *sdk.Client) (exportableConnectors *[]connector.Exportable) {
	// Using the --service parameter(s) provided by user, build list of connectors to export
	connectors := []connector.Exportable{}

	if multiService == nil {
		return &connectors
	}

	for _, service := range *multiService.GetServices() {
		switch service {
		case customtypes.ENUM_SERVICE_PLATFORM:
			connectors = append(connectors, platform.PlatformConnector(ctx, apiClient, &apiClientId, exportEnvID))
		case customtypes.ENUM_SERVICE_SSO:
			connectors = append(connectors, sso.SSOConnector(ctx, apiClient, &apiClientId, exportEnvID))
		case customtypes.ENUM_SERVICE_MFA:
			connectors = append(connectors, mfa.MFAConnector(ctx, apiClient, &apiClientId, exportEnvID))
		case customtypes.ENUM_SERVICE_PROTECT:
			connectors = append(connectors, protect.ProtectConnector(ctx, apiClient, &apiClientId, exportEnvID))
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
		output.Format(output.CommandOutput{
			Message: fmt.Sprintf("Exporting %s service...", connector.ConnectorServiceName()),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
		})

		err := connector.Export(string(exportFormat), outputDir, overwriteExport)
		if err != nil {
			return fmt.Errorf("failed to export '%s' service: %s", connector.ConnectorServiceName(), err.Error())
		}
	}

	return nil
}
