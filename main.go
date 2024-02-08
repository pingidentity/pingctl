package main

import (
	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/logger"
)

func main() {
	l := logger.Get()

	rootCmd := cmd.NewRootCommand()

	err := rootCmd.Execute()
	if err != nil {
		l.Fatal().Err(err).Msgf("Failed to execute pingctl")
	}
}
