package profile

import (
	"github.com/spf13/cobra"
)

func NewConfigProfileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long:  `Command to add, list, describe, delete, and set active configuration profiles.`,
		Short: "Command to add, list, describe, delete, and set active configuration profiles.",
		Use:   "profile",
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
