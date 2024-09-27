package config

import (
	"os"

	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

const (
	deleteProfileCommandExamples = `  pingctl config delete-profile
  pingctl config delete-profile --profile myprofile`
)

func NewConfigDeleteProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               deleteProfileCommandExamples,
		Long:                  `Delete a configuration profile from pingctl.`,
		RunE:                  configDeleteProfileRunE,
		Short:                 "Delete a configuration profile from pingctl.",
		Use:                   "delete-profile [flags]",
	}

	cmd.Flags().AddFlag(options.ConfigDeleteProfileOption.Flag)

	return cmd
}

func configDeleteProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config delete-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigDeleteProfile(os.Stdin); err != nil {
		return err
	}

	return nil
}
