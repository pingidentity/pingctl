package auth

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Auth Login Subcommand...")
}

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
