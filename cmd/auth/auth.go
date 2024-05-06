package auth

import (
	"github.com/spf13/cobra"
)

func NewAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate the CLI with configured Ping connections",
		Long:  "Authenticate the CLI with configured Ping connections",
	}

	cmd.AddCommand(NewLoginCommand(), NewLogoutCommand())

	return cmd
}
