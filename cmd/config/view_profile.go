package config

import (
	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

const (
	viewProfileCommandExamples = `  pingctl config view-profile
  pingctl config view-profile --profile myprofile`
)

func NewConfigViewProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               viewProfileCommandExamples,
		Long:                  `View a configuration profile from pingctl.`,
		RunE:                  configViewProfileRunE,
		Short:                 "View a configuration profile from pingctl.",
		Use:                   "view-profile [flags]",
	}

	cmd.Flags().AddFlag(options.ConfigViewProfileOption.Flag)

	return cmd
}

func configViewProfileRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config view-profile Subcommand Called.")

	if err := config_internal.RunInternalConfigViewProfile(); err != nil {
		return err
	}

	return nil
}
