package platform

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone_platform"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
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
	exportFormat ExportFormat = connector.ENUMEXPORTFORMAT_HCL
	multiService MultiService = MultiService{
		services: &[]string{
			serviceEnumPlatform,
		},
	}
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
		Use: "export",
		//TODO add command short and long description
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			l := logger.Get()

			l.Debug().Msgf("Export Subcommand Called.")

			apiClient, err := initApiClient(cmd.Context())
			if err != nil {
				output.Format(cmd, output.CommandOutput{
					Message: "Unable to initialize PingOne SDK client",
					Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
				})
				return err
			}

			if outputDir == "" {
				// Default the outputDir variable to the user's present working directory.
				pwd, err := os.Getwd()
				if err != nil {
					output.Format(cmd, output.CommandOutput{
						Message: "Failed to determine user's present working directory",
						Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
					})
					return err
				}

				l.Debug().Msgf("Defaulting export command output directory to %q...", pwd)

				outputDir = pwd
			}

			// Find the env ID to export. Default to worker env id if not provided by user.
			exportEnvID := viper.GetString(pingoneExportEnvironmentIdParamConfigKey)
			if exportEnvID == "" {
				exportEnvID = viper.GetString(pingoneWorkerEnvironmentIdParamConfigKey)

				// if the exportEnvID is still empty, this is a problem. Return error.
				if exportEnvID == "" {
					output.Format(cmd, output.CommandOutput{
						Message: "Failed to determine export environment ID",
						Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
					})
					return fmt.Errorf("failed to determine export environment ID")
				}
			}

			// Using the --service parameter(s) provided by user, build list of connectors to export
			exportableConnectors := []connector.Exportable{}
			for _, service := range *multiService.services {
				switch service {
				case serviceEnumPlatform:
					exportableConnectors = append(exportableConnectors, pingone_platform.Connector(cmd.Context(), apiClient, exportEnvID))
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
					output.Format(cmd, output.CommandOutput{
						Message: fmt.Sprintf("Export failed for service: %s.", connector.ConnectorServiceName()),
						Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
					})
					return err
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
	cmd.Flags().Var(&multiService, "service", fmt.Sprintf("Specifies service(s) to export. Allowed: %q", serviceEnumPlatform))
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

func initApiClient(ctx context.Context) (*sdk.Client, error) {
	l := logger.Get()

	if apiClient != nil {
		return apiClient, nil
	}

	l.Debug().Msgf("Initialising API client..")

	if !viper.IsSet(pingoneWorkerClientIdParamConfigKey) || !viper.IsSet(pingoneWorkerClientSecretParamConfigKey) ||
		!viper.IsSet(pingoneWorkerEnvironmentIdParamConfigKey) || !viper.IsSet(pingoneRegionParamConfigKey) {
		return nil, fmt.Errorf(`unable to initialize PingOne API client.
		One of environment ID, client ID, client secret, and region is not set.
		Configure these properties via parameter flags, environment variables, or configuration file`)
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

	apiConfig := &sdk.Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		Region:        region,
	}

	client, err := apiConfig.APIClient(ctx)
	if err != nil {
		return nil, err
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
