package profile

import (
	"github.com/spf13/cobra"
)

func NewConfigProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Command to add, list, describe, delete, and set active configuration profiles.",
		Long:  `Command to add, list, describe, delete, and set active configuration profiles.`,
	}

	cmd.AddCommand(
		NewConfigProfileAddCommand(),
		NewConfigProfileDeleteCommand(),
		NewConfigProfileDescribeCommand(),
		NewConfigProfileListCommand(),
		NewConfigProfileSetActiveCommand(),
	)

	return cmd
}
