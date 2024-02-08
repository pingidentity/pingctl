package auth

import (
	"github.com/spf13/cobra"
)

func NewAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "auth",
		//TODO more fleshed-out descriptions
		Short: "Authenticate with Ping",
		Long:  "Authenticate with Ping",
	}

	cmd.AddCommand(NewLoginCommand(), NewLogoutCommand())

	return cmd
}
