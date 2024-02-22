package auth

import (
	"github.com/spf13/cobra"
)

func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "login",
		//TODO more fleshed-out descriptions
		Short: "Login with Ping",
		Long:  "Login with Ping",
		RunE: func(cmd *cobra.Command, args []string) error {
			// authConnectors := []connector.Authenticatable{}
			return nil
		},
	}

	return cmd
}
