package config

import (
	"os"

	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl config
pingctl config --profile myprofile
pingctl config --name myprofile --description "My Profile"`,
		Long:  `Update a configuration profile's name and description. See subcommands for more profile configuration management.`,
		RunE:  configRunE,
		Short: "Update a configuration profile's name and description. See subcommands for more profile configuration management.",
		Use:   "config [flags]",
	}

	// Add subcommands
	cmd.AddCommand(
		NewConfigAddProfileCommand(),
		NewConfigDeleteProfileCommand(),
		NewConfigViewProfileCommand(),
		NewConfigListProfilesCommand(),
		NewConfigSetActiveProfileCommand(),
		NewConfigGetCommand(),
		NewConfigSetCommand(),
		NewConfigUnsetCommand(),
	)

	cmd.Flags().AddFlag(configuration.ConfigProfileOption.Flag)
	cmd.Flags().AddFlag(configuration.ConfigNameOption.Flag)
	cmd.Flags().AddFlag(configuration.ConfigDescriptionOption.Flag)

	return cmd
}

func configRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Subcommand Called.")

	if err := config_internal.RunInternalConfig(os.Stdin); err != nil {
		return err
	}

	return nil
}
