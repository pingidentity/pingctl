package auth

import (
	"github.com/spf13/cobra"
)

func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout user from the CLI",
		Long:  "Logout user from the CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			// authConnectors := []connector.Authenticatable{}
			return nil
		},
	}

	return cmd
}
