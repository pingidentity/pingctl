package platform

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	pingoneExportEnvironmentIdParamName      = "pingone-export-environment-id"
	pingoneExportEnvironmentIdParamConfigKey = "pingone.export-environment-id"

	pingoneWorkerEnvironmentIdParamName      = "pingone-worker-environment-id"
	pingoneWorkerEnvironmentIdParamConfigKey = "pingone.worker-environment-id"

	pingoneWorkerClientIdParamName      = "pingone-worker-client-id"
	pingoneWorkerClientIdParamConfigKey = "pingone.worker-client-id"

	pingoneWorkerClientSecretParamName      = "pingone-worker-client-secret"
	pingoneWorkerClientSecretParamConfigKey = "pingone.worker-client-secret"

	pingoneRegionParamName      = "pingone-region"
	pingoneRegionParamConfigKey = "pingone.region"
)

var (
	multiService MultiService = *NewMultiService()

	exportFormat    ExportFormat = connector.ENUMEXPORTFORMAT_HCL
	pingoneRegion   PingOneRegion
	outputDir       string
	overwriteExport bool
	apiClient       *sdk.Client

	exportConfigurationParamMapping = map[string]string{
		pingoneWorkerEnvironmentIdParamName: pingoneWorkerEnvironmentIdParamConfigKey,
		pingoneExportEnvironmentIdParamName: pingoneExportEnvironmentIdParamConfigKey,
		pingoneWorkerClientIdParamName:      pingoneWorkerClientIdParamConfigKey,
		pingoneWorkerClientSecretParamName:  pingoneWorkerClientSecretParamConfigKey,
		pingoneRegionParamName:              pingoneRegionParamConfigKey,
	}
)

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export configuration-as-code packages for the Ping Platform.",
		Long:  `Export configuration-as-code packages for the Ping Platform.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			l := logger.Get()

			l.Debug().Msgf("Export Subcommand Called.")

			apiClient, err := initApiClient(cmd.Context(), cmd.Root().Version)
			if err != nil {
				return fmt.Errorf("failed to initialize PingOne SDK client: %s", err.Error())
			}
			apiClientId := viper.GetString(pingoneWorkerClientIdParamConfigKey)

			if outputDir == "" {
				// Default the outputDir variable to the user's present working directory.
				outputDir, err = os.Getwd()
				if err != nil {
					return fmt.Errorf("failed to determine user's present working directory: %s", err.Error())
				}

				// Append "export" to the output directory as export needs an empty directory to write to
				outputDir = filepath.Join(outputDir, "export")

				l.Debug().Msgf("Defaulting export command output directory to %q...", outputDir)
			}

			// Check if outputDir exists
			// If not, create the directory
			l.Debug().Msgf("Validating export output directory %q...", outputDir)
			_, err = os.Stat(outputDir)
			if err != nil {
				output.Format(cmd, output.CommandOutput{
					Message: fmt.Sprintf("Failed to find or validate export output directory %q. Creating new output directory...", outputDir),
					Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
				})

				err = os.MkdirAll(outputDir, os.ModePerm)
				if err != nil {
					return fmt.Errorf("failed to create export output directory %q: %s", outputDir, err.Error())
				}

				l.Debug().Msgf("New export output directory %q created.", outputDir)
			} else {
				// Check if the output directory is empty
				// If not, default behavior is to exit and not overwrite.
				// This can be changed with the --overwrite export parameter
				if !overwriteExport {
					dirEntries, err := os.ReadDir(outputDir)
					if err != nil {
						return fmt.Errorf("failed to read contents of export directory %q: %s", outputDir, err.Error())
					}

					if len(dirEntries) > 0 {
						return fmt.Errorf("export directory %q is not empty. Use --overwrite to overwrite existing export data", outputDir)
					}
				}
			}

			// Find the env ID to export. Default to worker env id if not provided by user.
			exportEnvID := viper.GetString(pingoneExportEnvironmentIdParamConfigKey)
			if exportEnvID == "" {
				exportEnvID = viper.GetString(pingoneWorkerEnvironmentIdParamConfigKey)

				// if the exportEnvID is still empty, this is a problem. Return error.
				if exportEnvID == "" {
					return fmt.Errorf("failed to determine export environment ID")
				}

				output.Format(cmd, output.CommandOutput{
					Message: "No target export environment ID specified. Defaulting export environment ID to the Worker App environment ID.",
					Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
				})
			}

			l.Debug().Msgf("Validating export environment ID...")

			environment, response, err := apiClient.ManagementAPIClient.EnvironmentsApi.ReadOneEnvironment(cmd.Context(), exportEnvID).Execute()
			defer response.Body.Close()
			if err != nil {
				return fmt.Errorf("failed to read environment.\nReadOneEnvironment Response Code: %s\nResponse Body: %s\n Error: %s", response.Status, response.Body, err.Error())
			}

			if environment == nil {
				if response.StatusCode == 404 {
					return fmt.Errorf("failed to fetch environment. the provided environment id %q does not exist.\nReadOneEnvironment Response Code: %s\nResponse Body: %s", exportEnvID, response.Status, response.Body)
				} else {
					return fmt.Errorf("failed to fetch environment %q via ReadOneEnvironment()\nReadOneEnvironment Response Code: %s\nResponse Body: %s", exportEnvID, response.Status, response.Body)
				}
			}

			// Using the --service parameter(s) provided by user, build list of connectors to export
			exportableConnectors := []connector.Exportable{}

			for _, service := range *multiService.GetServices() {
				switch service {
				case serviceEnumPlatform:
					exportableConnectors = append(exportableConnectors, platform.PlatformConnector(cmd.Context(), apiClient, &apiClientId, exportEnvID))
				case serviceEnumSSO:
					exportableConnectors = append(exportableConnectors, sso.SSOConnector(cmd.Context(), apiClient, &apiClientId, exportEnvID))
					// default:
					// This unrecognized service condition is handled by cobra with the custom type MultiService
				}
			}

			// Loop through user defined exportable connectors and export them
			for _, connector := range exportableConnectors {
				output.Format(cmd, output.CommandOutput{
					Message: fmt.Sprintf("Exporting %s service...", connector.ConnectorServiceName()),
					Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
				})

				err := connector.Export(string(exportFormat), outputDir, overwriteExport)
				if err != nil {
					return fmt.Errorf("failed to export %s service: %s", connector.ConnectorServiceName(), err.Error())
				}
			}

			output.Format(cmd, output.CommandOutput{
				Message: fmt.Sprintf("Export to directory %q complete.", outputDir),
				Result:  output.ENUMCOMMANDOUTPUTRESULT_SUCCESS,
			})
			return nil
		},
	}

	// Add flags that are not tracked in the viper configuration file
	cmd.Flags().Var(&exportFormat, "export-format", fmt.Sprintf("Specifies export format\nAllowed: %q", connector.ENUMEXPORTFORMAT_HCL))
	cmd.Flags().Var(&multiService, "service", fmt.Sprintf("Specifies service(s) to export. Allowed: %s", multiService.String()))
	cmd.Flags().StringVar(&outputDir, "output-directory", "", "Specifies output directory for export (Default: Present working directory)")
	cmd.Flags().BoolVar(&overwriteExport, "overwrite", false, "Overwrite existing generated exports if set.")

	// Add flags that are bound to configuration file keys
	cmd.Flags().String(pingoneWorkerEnvironmentIdParamName, "", "The ID of the PingOne environment that contains the worker token client used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")
	cmd.Flags().String(pingoneExportEnvironmentIdParamName, "", "The ID of the PingOne environment to export. (Default: The PingOne worker environment ID)\nAlso configurable via environment variable PINGCTL_PINGONE_EXPORT_ENVIRONMENT_ID")
	cmd.Flags().String(pingoneWorkerClientIdParamName, "", "The ID of the worker app (also the client ID) used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_CLIENT_ID")
	cmd.Flags().String(pingoneWorkerClientSecretParamName, "", "The client secret of the worker app used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_CLIENT_SECRET")
	cmd.Flags().Var(&pingoneRegion, pingoneRegionParamName, fmt.Sprintf("The region of the service. Allowed: %q, %q, %q, %q\nAlso configurable via environment variable PINGCTL_PINGONE_REGION", connector.ENUMREGION_AP, connector.ENUMREGION_CA, connector.ENUMREGION_EU, connector.ENUMREGION_NA))

	cmd.MarkFlagsRequiredTogether(pingoneWorkerEnvironmentIdParamName, pingoneWorkerClientIdParamName, pingoneWorkerClientSecretParamName, pingoneRegionParamName)

	// Bind the newly created flags to viper configuration file
	if err := bindFlags(exportConfigurationParamMapping, cmd); err != nil {
		output.Format(cmd, output.CommandOutput{
			Message: "Error binding export command flag parameters. Flag values may not be recognized.",
			Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
			Error:   err,
		})
	}

	return cmd
}

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Export Subcommand...")
}

func initApiClient(ctx context.Context, version string) (*sdk.Client, error) {
	l := logger.Get()

	if apiClient != nil {
		return apiClient, nil
	}

	l.Debug().Msgf("Initialising API client..")

	// Make sure the API client can be initialized with the required parameters
	if !viper.IsSet(pingoneWorkerClientIdParamConfigKey) || !viper.IsSet(pingoneWorkerClientSecretParamConfigKey) ||
		!viper.IsSet(pingoneWorkerEnvironmentIdParamConfigKey) || !viper.IsSet(pingoneRegionParamConfigKey) {
		return nil, fmt.Errorf(`unable to initialize PingOne API client.
		One of worker environment ID, worker client ID, worker client secret,
		and/or pingone region is not set.
		Configure these properties via parameter flags, environment variables,
		or the tool's configuration file (Default: $HOME/.pingctl/config.yaml)`)
	}

	clientID := viper.GetString(pingoneWorkerClientIdParamConfigKey)
	clientSecret := viper.GetString(pingoneWorkerClientSecretParamConfigKey)
	environmentID := viper.GetString(pingoneWorkerEnvironmentIdParamConfigKey)
	region := viper.GetString(pingoneRegionParamConfigKey)

	switch region {
	case connector.ENUMREGION_AP, connector.ENUMREGION_CA, connector.ENUMREGION_EU, connector.ENUMREGION_NA:
		l.Debug().Msgf("PingOne region %q validated.", region)
	default:
		return nil, fmt.Errorf("unrecognized PingOne Region: %q. Must be one of: %q, %q, %q, %q", region, connector.ENUMREGION_AP, connector.ENUMREGION_CA, connector.ENUMREGION_EU, connector.ENUMREGION_NA)
	}

	userAgent := fmt.Sprintf("pingctl/%s", version)

	if v := strings.TrimSpace(os.Getenv("PINGCTL_PINGONE_APPEND_USER_AGENT")); v != "" {
		userAgent += fmt.Sprintf(" %s", v)
	}

	apiConfig := &sdk.Config{
		ClientID:        &clientID,
		ClientSecret:    &clientSecret,
		EnvironmentID:   &environmentID,
		Region:          region,
		UserAgentSuffix: &userAgent,
	}

	client, err := apiConfig.APIClient(ctx)
	if err != nil {
		// If logging level is DEBUG or TRACE, include the worker client secret in the error message
		var clientSecretErrorMessage string
		if l.GetLevel() <= zerolog.DebugLevel {
			clientSecretErrorMessage = fmt.Sprintf("Worker Client Secret - %s", clientSecret)
		} else {
			clientSecretErrorMessage = "Worker Client Secret - (Use DEBUG or TRACE logging level to view the client secret in the error message)" // #nosec G101
		}
		return nil, fmt.Errorf("%s\n\nConfiguration values used for client initialization:\nWorker Client ID - %s\nWorker Environment ID - %s\nPingOne Region - %s\n%s", err.Error(), clientID, environmentID, region, clientSecretErrorMessage)
	}

	return client, nil
}

func bindFlags(paramlist map[string]string, command *cobra.Command) error {
	for k, v := range paramlist {
		err := viper.BindPFlag(v, command.Flags().Lookup(k))
		if err != nil {
			return err
		}
	}

	return nil
}
