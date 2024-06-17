package auth

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Auth Subcommand...")
}

func NewAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate the CLI with configured Ping connections",
		Long:  "Authenticate the CLI with configured Ping connections",
	}

	cmd.AddCommand(NewLoginCommand(), NewLogoutCommand())

	return cmd
}
