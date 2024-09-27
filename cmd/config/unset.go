package config

import (
	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configUnsetCommandExamples = `  pingctl config unset color
  pingctl config unset --profile myProfile service.pingone.regionCode`
)

func NewConfigUnsetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configUnsetCommandExamples,
		Long:                  `Unset pingctl configuration settings.`,
		RunE:                  configUnsetRunE,
		Short:                 "Unset pingctl configuration settings.",
		Use:                   "unset [flags] key",
	}

	cmd.Flags().AddFlag(options.ConfigUnsetProfileOption.Flag)

	return cmd
}
func configUnsetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config unset Subcommand Called.")

	if err := config_internal.RunInternalConfigUnset(args[0]); err != nil {
		return err
	}

	return nil
}
