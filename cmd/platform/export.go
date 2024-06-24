package platform

import (
	"fmt"
	"strings"

	platform_internal "github.com/pingidentity/pingctl/internal/commands/platform"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/cobra"
)

var (
	multiService customtypes.MultiService = *customtypes.NewMultiService()

	exportFormat    customtypes.ExportFormat = connector.ENUMEXPORTFORMAT_HCL
	pingoneRegion   customtypes.PingOneRegion
	outputDir       string
	overwriteExport bool

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
	cmd.Flags().Var(&multiService, "service", fmt.Sprintf("Specifies service(s) to export. Allowed services: %s", multiService.String()))
	cmd.Flags().StringVar(&outputDir, "output-directory", "", "Specifies output directory for export (Default: Present working directory)")
	cmd.Flags().BoolVar(&overwriteExport, "overwrite", false, "Overwrite existing generated exports if set.")

	// Add flags that are bound to configuration file keys
	cmd.Flags().String(string(viperconfig.ExportPingoneWorkerEnvironmentIdParamName), "", "The ID of the PingOne environment that contains the worker token client used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")
	cmd.Flags().String(string(viperconfig.ExportPingoneExportEnvironmentIdParamName), "", "The ID of the PingOne environment to export. (Default: The PingOne worker environment ID)\nAlso configurable via environment variable PINGCTL_PINGONE_EXPORT_ENVIRONMENT_ID")
	cmd.Flags().String(string(viperconfig.ExportPingoneWorkerClientIdParamName), "", "The ID of the worker app (also the client ID) used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_CLIENT_ID")
	cmd.Flags().String(string(viperconfig.ExportPingoneWorkerClientSecretParamName), "", "The client secret of the worker app used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_CLIENT_SECRET")
	cmd.Flags().Var(&pingoneRegion, string(viperconfig.ExportPingoneRegionParamName), fmt.Sprintf("The region of the service. Allowed: %s\nAlso configurable via environment variable PINGCTL_PINGONE_REGION", strings.Join(customtypes.PingOneRegionValidValues(), ", ")))

	cmd.MarkFlagsRequiredTogether(string(viperconfig.ExportPingoneWorkerEnvironmentIdParamName), string(viperconfig.ExportPingoneWorkerClientIdParamName), string(viperconfig.ExportPingoneWorkerClientSecretParamName), string(viperconfig.ExportPingoneRegionParamName))

	// Bind the newly created flags to viper configuration file
	if err := viperconfig.BindFlags(cobraParamNames, cmd); err != nil {
		output.Format(output.CommandOutput{
			Message:      "Error binding export command flag parameters. Flag values may not be recognized.",
			Result:       output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
			ErrorMessage: err.Error(),
		})
	}

	if err := viperconfig.BindEnvVars(cobraParamNames); err != nil {
		output.Format(output.CommandOutput{
			Message:      "Error binding environment variables. Environment Variable values may not be recognized.",
			Result:       output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
			ErrorMessage: err.Error(),
		})
	}

	return cmd
}

func ExportRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()

	l.Debug().Msgf("Platform Export Subcommand Called.")

	return platform_internal.RunInternalExport(cmd, outputDir, string(exportFormat), overwriteExport, &multiService)
}
