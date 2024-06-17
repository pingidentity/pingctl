package auth

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Auth Logout Subcommand...")
}

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
