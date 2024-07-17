package config

import (
	"github.com/pingidentity/pingctl/cmd/config/profile"
	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Command to get, set, and unset pingctl configuration settings.",
		Long:  `Command to get, set, and unset pingctl configuration settings.`,
	}

	cmd.AddCommand(NewConfigGetCommand())
	cmd.AddCommand(NewConfigSetCommand())
	cmd.AddCommand(NewConfigUnsetCommand())
	cmd.AddCommand(profile.NewConfigProfileCommand())

	return cmd
}
