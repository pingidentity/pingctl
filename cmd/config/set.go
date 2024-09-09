package config

import (
	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl config set pingctl.color=true
pingctl config set --profile myProfile pingone.region=AsiaPacific`,
		Long:  `Set pingctl configuration settings.`,
		RunE:  configSetRunE,
		Short: "Set pingctl configuration settings.",
		Use:   "set [flags] key=value",
	}

	cmd.Flags().AddFlag(configuration.ConfigSetProfileOption.Flag)

	return cmd
}
func configSetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config set Subcommand Called.")

	if err := config_internal.RunInternalConfigSet(args[0]); err != nil {
		return err
	}

	return nil
}
