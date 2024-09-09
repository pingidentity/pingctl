package config

import (
	"os"

	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigSetActiveProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl config set-active-profile
pingctl config set-active-profile --profile myprofile`,
		Long:  `Set a configuration profile as the in-use profile for pingctl.`,
		RunE:  configSetActiveProfileRunE,
		Short: "Set a configuration profile as the in-use profile for pingctl.",
		Use:   "set-active-profile [flags]",
	}

	cmd.Flags().AddFlag(configuration.ConfigSetActiveProfileOption.Flag)

	return cmd
}

func configSetActiveProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config set-active-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigSetActiveProfile(os.Stdin); err != nil {
		return err
	}

	return nil
}
