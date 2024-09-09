package config

import (
	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigListProfilesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               `pingctl config list-profiles`,
		Long:                  `List all configuration profiles from pingctl.`,
		RunE:                  configListProfilesRunE,
		Short:                 "List all configuration profiles from pingctl.",
		Use:                   "list-profiles",
	}

	return cmd
}

func configListProfilesRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config list-profiles Subcommand Called.")

	config_internal.RunInternalConfigListProfiles()

	return nil
}
