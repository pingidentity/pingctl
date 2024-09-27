package config

import (
	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

const (
	configGetCommandExamples = `  pingctl config get pingone
  pingctl config get --profile myProfile color
  pingctl config get service.pingone.authentication.worker.environmentID`
)

func NewConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               configGetCommandExamples,
		Long:                  `Get pingctl configuration settings.`,
		RunE:                  configGetRunE,
		Short:                 "Get pingctl configuration settings.",
		Use:                   "get [flags] key",
	}

	cmd.Flags().AddFlag(options.ConfigGetProfileOption.Flag)

	return cmd
}

func configGetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	if err := config_internal.RunInternalConfigGet(args[0]); err != nil {
		return err
	}

	return nil
}
