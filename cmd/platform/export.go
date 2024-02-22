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
	outputDir       string
	overwriteExport bool
	apiClient       *sdk.Client

	exportConfigurationParamMapping = map[string]string{
		pingoneWorkerEnvironmentIdParamName: pingoneWorkerEnvironmentIdParamConfigKey,
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

			apiClient, err := initApiClient(cmd.Context(), cmd)
			if err != nil {
				output.Format(cmd, output.CommandOutput{
					Message: "Unable to initialize PingOne SDK client",
					Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
				})
				return err
			}

			// Using the --service parameter(s) provided by user, build list of connectors to export
			exportableConnectors := []connector.Exportable{}
			for _, service := range *multiService.services {
				switch service {
				case serviceEnumPlatform:
					exportableConnectors = append(exportableConnectors, pingone_platform.Connector(cmd.Context(), apiClient, viper.GetString(pingoneWorkerEnvironmentIdParamConfigKey)))
					// default:
					// This unrecognized service condition is handled by cobra with the custom type MultiService
				}
			}

			// Loop through user defined exportable connectors and export them
			for _, connector := range exportableConnectors {
				l.Debug().Msgf("Exporting %s service...", connector.ConnectorServiceName())

				err := connector.Export(string(exportFormat), outputDir, overwriteExport)
				if err != nil {
					output.Format(cmd, output.CommandOutput{
						Message: fmt.Sprintf("Export failed for service: %s.", connector.ConnectorServiceName()),
						Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
					})
					return err
				}
			}
			return nil
		},
	}

	// Add flags that are not tracked in the viper configuration file
	cmd.Flags().Var(&exportFormat, "export-format", fmt.Sprintf("Specifies export format\nAllowed: %q", connector.ENUMEXPORTFORMAT_HCL))
	cmd.Flags().Var(&multiService, "service", fmt.Sprintf("Specifies service(s) to export. Allowed: %q", serviceEnumPlatform))
	cmd.Flags().StringVar(&outputDir, "output-directory", "", "Specifies output directory for export (Default: Present working directory)")
	cmd.Flags().BoolVar(&overwriteExport, "overwrite", false, "Overwrite existing generated exports if set.")

	// Add flags that are bound to configuration file keys
	cmd.Flags().String(pingoneWorkerEnvironmentIdParamName, os.Getenv("PINGONE_ENVIRONMENT_ID"), "The ID of the PingOne environment that contains the worker token client used to authenticate.\nAlso configurable via environment variable PINGONE_ENVIRONMENT_ID")
	cmd.Flags().String(pingoneWorkerClientIdParamName, os.Getenv("PINGONE_CLIENT_ID"), "The ID of the worker app (also the client ID) used to authenticate.\nAlso configurable via environment variable PINGONE_CLIENT_ID")
	cmd.Flags().String(pingoneWorkerClientSecretParamName, os.Getenv("PINGONE_CLIENT_SECRET"), "The client secret of the worker app used to authenticate.\nAlso configurable via environment variable PINGONE_CLIENT_SECRET")
	cmd.Flags().String(pingoneRegionParamName, os.Getenv("PINGONE_REGION"), "The region code of the service (NA, EU, AP, CA).\nAlso configurable via environment variable PINGONE_REGION")

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

	if outputDir == "" {
		// Default the outputDir variable to the user's present working directory.
		pwd, err := os.Getwd()
		if err != nil {
			l.Fatal().Err(err).Msgf("Failed to determine user's present working directory")
		}

		outputDir = pwd
	}
}

func initApiClient(ctx context.Context, cmd *cobra.Command) (*sdk.Client, error) {
	l := logger.Get()

	if apiClient != nil {
		return apiClient, nil
	}

	l.Debug().Msgf("Initialising API client..")

	clientID := viper.GetString(pingoneWorkerClientIdParamConfigKey)
	clientSecret := viper.GetString(pingoneWorkerClientSecretParamConfigKey)
	environmentID := viper.GetString(pingoneWorkerEnvironmentIdParamConfigKey)

	var region string
	switch viper.GetString(pingoneRegionParamConfigKey) {
	case "NA":
		region = "NorthAmerica"
	case "EU":
		region = "Europe"
	case "AP":
		region = "AsiaPacific"
	case "CA":
		region = "Canada"
	default:
		return nil, fmt.Errorf("provided pingone region code not recognized: %s", viper.GetString(pingoneRegionParamConfigKey))
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
