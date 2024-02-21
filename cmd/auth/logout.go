package auth

import (
	"github.com/spf13/cobra"
)

func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "logout",
		//TODO more fleshed-out descriptions
		Short: "Logout with Ping",
		Long:  "Logout with Ping",
		RunE: func(cmd *cobra.Command, args []string) error {
			// authConnectors := []connector.Authenticatable{}
			return nil
		},
	}

	return cmd
}
