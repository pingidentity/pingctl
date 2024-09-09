package config

import (
	"os"

	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigAddProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl config add-profile
pingctl config add-profile --name myprofile --description "My Profile desc"
pingctl config add-profile --set-active=true`,
		Long:  `Add a new configuration profile to pingctl.`,
		RunE:  configAddProfileRunE,
		Short: "Add a new configuration profile to pingctl.",
		Use:   "add-profile [flags]",
	}

	cmd.Flags().AddFlag(configuration.ConfigAddProfileNameOption.Flag)
	cmd.Flags().AddFlag(configuration.ConfigAddProfileDescriptionOption.Flag)
	cmd.Flags().AddFlag(configuration.ConfigAddProfileSetActiveOption.Flag)

	return cmd
}

func configAddProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config add-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigAddProfile(os.Stdin); err != nil {
		return err
	}

	return nil
}
