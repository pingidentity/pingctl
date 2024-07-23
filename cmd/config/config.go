package config

import (
	"github.com/pingidentity/pingctl/cmd/config/profile"
	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long:  `Command to get, set, and unset pingctl configuration settings.`,
		Short: "Command to get, set, and unset pingctl configuration settings.",
		Use:   "config",
	}

	cmd.AddCommand(NewConfigGetCommand())
	cmd.AddCommand(NewConfigSetCommand())
	cmd.AddCommand(NewConfigUnsetCommand())
	cmd.AddCommand(profile.NewConfigProfileCommand())

	return cmd
}
