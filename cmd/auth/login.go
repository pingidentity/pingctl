package auth

import (
	"github.com/spf13/cobra"
)

func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login user to the CLI",
		Long:  "Login user to the CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			// authConnectors := []connector.Authenticatable{}
			return nil
		},
	}

	return cmd
}
