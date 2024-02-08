package main

import (
	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/output"
)

func main() {
	rootCmd := cmd.NewRootCommand()

	err := rootCmd.Execute()
	if err != nil {
		output.Format(rootCmd, output.CommandOutput{
			Fatal:   err,
			Message: "Failed to execute pingctl",
			Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
		})
	}
}
