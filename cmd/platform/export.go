package platform

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

var (
	multiService customtypes.MultiService = *customtypes.NewMultiService()

	exportFormat    customtypes.ExportFormat = connector.ENUMEXPORTFORMAT_HCL
	pingoneRegion   customtypes.PingOneRegion
	outputDir       string
	overwriteExport bool
	apiClient       *sdk.Client
	apiClientId     string

	cobraParamNames = []viperconfig.ConfigCobraParam{
		viperconfig.ExportPingoneExportEnvironmentIdParamName,
		viperconfig.ExportPingoneWorkerEnvironmentIdParamName,
		viperconfig.ExportPingoneWorkerClientIdParamName,
		viperconfig.ExportPingoneWorkerClientSecretParamName,
		viperconfig.ExportPingoneRegionParamName,
	}
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Platform Export Subcommand...")
}

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export configuration-as-code packages for the Ping Platform.",
		Long:  `Export configuration-as-code packages for the Ping Platform.`,
		RunE:  ExportRunE,
	}

	// Add flags that are not tracked in the viper configuration file
	cmd.Flags().Var(&exportFormat, "export-format", fmt.Sprintf("Specifies export format\nAllowed: %q", connector.ENUMEXPORTFORMAT_HCL))
	cmd.Flags().Var(&multiService, "service", fmt.Sprintf("Specifies service(s) to export. Allowed: %s", multiService.String()))
	cmd.Flags().StringVar(&outputDir, "output-directory", "", "Specifies output directory for export (Default: Present working directory)")
	cmd.Flags().BoolVar(&overwriteExport, "overwrite", false, "Overwrite existing generated exports if set.")

	// Add flags that are bound to configuration file keys
	cmd.Flags().String(string(viperconfig.ExportPingoneWorkerEnvironmentIdParamName), "", "The ID of the PingOne environment that contains the worker token client used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")
	cmd.Flags().String(string(viperconfig.ExportPingoneExportEnvironmentIdParamName), "", "The ID of the PingOne environment to export. (Default: The PingOne worker environment ID)\nAlso configurable via environment variable PINGCTL_PINGONE_EXPORT_ENVIRONMENT_ID")
	cmd.Flags().String(string(viperconfig.ExportPingoneWorkerClientIdParamName), "", "The ID of the worker app (also the client ID) used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_CLIENT_ID")
	cmd.Flags().String(string(viperconfig.ExportPingoneWorkerClientSecretParamName), "", "The client secret of the worker app used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_CLIENT_SECRET")
	cmd.Flags().Var(&pingoneRegion, string(viperconfig.ExportPingoneRegionParamName), fmt.Sprintf("The region of the service. Allowed: %s\nAlso configurable via environment variable PINGCTL_PINGONE_REGION", customtypes.PingOneRegionValidValues()))

	cmd.MarkFlagsRequiredTogether(string(viperconfig.ExportPingoneWorkerEnvironmentIdParamName), string(viperconfig.ExportPingoneWorkerClientIdParamName), string(viperconfig.ExportPingoneWorkerClientSecretParamName), string(viperconfig.ExportPingoneRegionParamName))

	// Bind the newly created flags to viper configuration file
	if err := viperconfig.BindFlags(cobraParamNames, cmd); err != nil {
		output.Format(cmd, output.CommandOutput{
			Message:      "Error binding export command flag parameters. Flag values may not be recognized.",
			Result:       output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
			ErrorMessage: err.Error(),
		})
	}

	if err := viperconfig.BindEnvVars(cobraParamNames); err != nil {
		output.Format(cmd, output.CommandOutput{
			Message:      "Error binding environment varibales. Environment Variable values may not be recognized.",
			Result:       output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
			ErrorMessage: err.Error(),
		})
	}

	return cmd
}

func ExportRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()

	l.Debug().Msgf("Platform Export Subcommand Called.")

	if err := initApiClient(cmd.Context(), cmd.Root().Version); err != nil {
		return err
	}

	if err := fixEmptyOutputDirVar(cmd); err != nil {
		return err
	}

	if err := createOrValidateOutputDir(cmd); err != nil {
		return err
	}

	exportEnvID, err := getExportEnvID(cmd)
	if err != nil {
		return err
	}

	if err := validateExportEnvID(cmd, exportEnvID); err != nil {
		return err
	}

	exportableConnectors := getExportableConnectors(exportEnvID, cmd)

	if err := exportConnectors(cmd, exportableConnectors); err != nil {
		return err
	}

	output.Format(cmd, output.CommandOutput{
		Message: fmt.Sprintf("Export to directory '%s' complete.", outputDir),
		Result:  output.ENUMCOMMANDOUTPUTRESULT_SUCCESS,
	})
	return nil
}

func initApiClient(ctx context.Context, version string) error {
	l := logger.Get()
	l.Debug().Msgf("Initialising API client..")

	if apiClient != nil && apiClientId != "" {
		return nil
	}

	// Make sure the API client can be initialized with the required parameters
	if !viper.IsSet(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerClientIdParamName].ViperConfigKey) ||
		!viper.IsSet(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerClientSecretParamName].ViperConfigKey) ||
		!viper.IsSet(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerEnvironmentIdParamName].ViperConfigKey) ||
		!viper.IsSet(viperconfig.ConfigOptions[viperconfig.ExportPingoneRegionParamName].ViperConfigKey) {
		return fmt.Errorf(`failed to initialize pingone API client.
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
		return fmt.Errorf("failed to initialize pingone API client. unrecognized pingone region variable type: %T", region)
	}

	switch regionStr {
	case customtypes.ENUM_PINGONE_REGION_AP, customtypes.ENUM_PINGONE_REGION_CA, customtypes.ENUM_PINGONE_REGION_EU, customtypes.ENUM_PINGONE_REGION_NA:
		l.Debug().Msgf("pingone region '%s' validated.", regionStr)
	default:
		return fmt.Errorf("failed to initialize pingone API client. unrecognized pingone region: '%s'. Must be one of: %s", regionStr, customtypes.PingOneRegionValidValues())
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

	var err error
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
		return fmt.Errorf(initFailErrFormatMessage, err.Error(), apiClientId, environmentID, region, clientSecretErrorMessage)
	}

	return nil
}

func fixEmptyOutputDirVar(cmd *cobra.Command) error {
	if outputDir == "" {
		// Default the outputDir variable to the user's present working directory.
		var err error
		outputDir, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to determine user's present working directory: %s", err.Error())
		}

		// Append "export" to the output directory as export needs an empty directory to write to
		outputDir = filepath.Join(outputDir, "export")

		output.Format(cmd, output.CommandOutput{
			Message: fmt.Sprintf("Defaulting 'platform export' command output directory to '%s'", outputDir),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})
	}

	return nil
}

func createOrValidateOutputDir(cmd *cobra.Command) error {
	l := logger.Get()

	// Check if outputDir exists
	// If not, create the directory
	l.Debug().Msgf("Validating export output directory '%s'", outputDir)
	_, err := os.Stat(outputDir)
	if err != nil {
		output.Format(cmd, output.CommandOutput{
			Message: fmt.Sprintf("failed to find 'platform export' output directory. creating new output directory at filepath '%s'", outputDir),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})

		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create 'platform export' output directory '%s': %s", outputDir, err.Error())
		}

		output.Format(cmd, output.CommandOutput{
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

func getExportEnvID(cmd *cobra.Command) (string, error) {
	// Find the env ID to export. Default to worker env id if not provided by user.
	exportEnvID := viper.GetString(viperconfig.ConfigOptions[viperconfig.ExportPingoneExportEnvironmentIdParamName].ViperConfigKey)
	if exportEnvID == "" {
		exportEnvID = viper.GetString(viperconfig.ConfigOptions[viperconfig.ExportPingoneWorkerEnvironmentIdParamName].ViperConfigKey)

		// if the exportEnvID is still empty, this is a problem. Return error.
		if exportEnvID == "" {
			return "", fmt.Errorf("failed to determine export environment ID")
		}

		output.Format(cmd, output.CommandOutput{
			Message: "No target export environment ID specified. Defaulting export environment ID to the Worker App environment ID.",
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})
	}

	return exportEnvID, nil
}

func validateExportEnvID(cmd *cobra.Command, exportEnvID string) error {
	l := logger.Get()
	l.Debug().Msgf("Validating export environment ID...")

	environment, response, err := apiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(cmd.Context(), exportEnvID).Execute()
	err = common.HandleClientResponse(response, err, "ReadOneEnvironment", "pingone_environment")
	if err != nil {
		return err
	}

	if environment == nil {
		return fmt.Errorf("failed to validate environment ID '%s'. environment matching ID does not exist", exportEnvID)
	}

	return nil
}

func getExportableConnectors(exportEnvID string, cmd *cobra.Command) *[]connector.Exportable {
	// Using the --service parameter(s) provided by user, build list of connectors to export
	exportableConnectors := []connector.Exportable{}

	for _, service := range *multiService.GetServices() {
		switch service {
		case customtypes.ENUM_SERVICE_PLATFORM:
			exportableConnectors = append(exportableConnectors, platform.PlatformConnector(cmd.Context(), apiClient, &apiClientId, exportEnvID))
		case customtypes.ENUM_SERVICE_SSO:
			exportableConnectors = append(exportableConnectors, sso.SSOConnector(cmd.Context(), apiClient, &apiClientId, exportEnvID))
		case customtypes.ENUM_SERVICE_MFA:
			exportableConnectors = append(exportableConnectors, mfa.MFAConnector(cmd.Context(), apiClient, &apiClientId, exportEnvID))
		case customtypes.ENUM_SERVICE_PROTECT:
			exportableConnectors = append(exportableConnectors, protect.ProtectConnector(cmd.Context(), apiClient, &apiClientId, exportEnvID))
			// default:
			// This unrecognized service condition is handled by cobra with the custom type MultiService
		}
	}

	return &exportableConnectors
}

func exportConnectors(cmd *cobra.Command, exportableConnectors *[]connector.Exportable) error {
	// Loop through user defined exportable connectors and export them
	for _, connector := range *exportableConnectors {
		output.Format(cmd, output.CommandOutput{
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
