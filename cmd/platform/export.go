package platform

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/cmd/common"
	platform_internal "github.com/pingidentity/pingctl/internal/commands/platform"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/spf13/cobra"
)

var (
	multiService customtypes.MultiService = *customtypes.NewMultiService()

	exportFormat    customtypes.ExportFormat = connector.ENUMEXPORTFORMAT_HCL
	pingoneRegion   customtypes.PingOneRegion
	outputDir       string
	overwriteExport bool
)

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl platform export
pingctl platform export --service pingone-platform --service pingone-sso
pingctl platform export --output-directory dir --overwrite
pingctl platform export --export-format HCL`,
		Long:  `Export configuration-as-code packages for the Ping Platform.`,
		Short: "Export configuration-as-code packages for the Ping Platform.",
		RunE:  exportRunE,
		Use:   "export [flags]",
	}

	// Add flags that are not tracked in the viper configuration file
	cmd.Flags().VarP(&exportFormat, "export-format", "e", fmt.Sprintf("Specifies export format\nAllowed: %q", connector.ENUMEXPORTFORMAT_HCL))
	cmd.Flags().VarP(&multiService, "service", "s", fmt.Sprintf("Specifies service(s) to export. Allowed services: %s", multiService.String()))
	cmd.Flags().StringVarP(&outputDir, "output-directory", "d", "", "Specifies output directory for export (Default: Present working directory)")
	cmd.Flags().BoolVarP(&overwriteExport, "overwrite", "o", false, "Overwrite existing generated exports if set.")

	// Add flags that are bound to configuration file keys
	cmd.Flags().String(profiles.WorkerEnvironmentIDOption.CobraParamName, "", "The ID of the PingOne environment that contains the worker token client used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID")
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.WorkerEnvironmentIDOption,
		Flag:   cmd.Flags().Lookup(profiles.WorkerEnvironmentIDOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.WorkerEnvironmentIDOption)

	cmd.Flags().String(profiles.ExportEnvironmentIDOption.CobraParamName, "", "The ID of the PingOne environment to export. (Default: The PingOne worker environment ID)\nAlso configurable via environment variable PINGCTL_PINGONE_EXPORT_ENVIRONMENT_ID")
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.ExportEnvironmentIDOption,
		Flag:   cmd.Flags().Lookup(profiles.ExportEnvironmentIDOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.ExportEnvironmentIDOption)

	cmd.Flags().String(profiles.WorkerClientIDOption.CobraParamName, "", "The ID of the worker app (also the client ID) used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_CLIENT_ID")
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.WorkerClientIDOption,
		Flag:   cmd.Flags().Lookup(profiles.WorkerClientIDOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.WorkerClientIDOption)

	cmd.Flags().String(profiles.WorkerClientSecretOption.CobraParamName, "", "The client secret of the worker app used to authenticate.\nAlso configurable via environment variable PINGCTL_PINGONE_WORKER_CLIENT_SECRET")
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.WorkerClientSecretOption,
		Flag:   cmd.Flags().Lookup(profiles.WorkerClientSecretOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.WorkerClientSecretOption)

	cmd.Flags().Var(&pingoneRegion, profiles.RegionOption.CobraParamName, fmt.Sprintf("The region of the service. Allowed: %s\nAlso configurable via environment variable PINGCTL_PINGONE_REGION", strings.Join(customtypes.PingOneRegionValidValues(), ", ")))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.RegionOption,
		Flag:   cmd.Flags().Lookup(profiles.RegionOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.RegionOption)

	cmd.MarkFlagsRequiredTogether(profiles.WorkerEnvironmentIDOption.CobraParamName, profiles.WorkerClientIDOption.CobraParamName, profiles.WorkerClientSecretOption.CobraParamName, profiles.RegionOption.CobraParamName)

	return cmd
}

func exportRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()

	l.Debug().Msgf("Platform Export Subcommand Called.")

	return platform_internal.RunInternalExport(cmd, outputDir, string(exportFormat), overwriteExport, &multiService)
}
