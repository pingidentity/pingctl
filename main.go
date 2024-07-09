package main

import (
	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/output"
)

func main() {
	rootCmd := cmd.NewRootCommand()

	err := rootCmd.Execute()
	if err != nil {
		output.Print(output.Opts{
			ErrorMessage: err.Error(),
			Message:      "Failed to execute pingctl",
			Result:       output.ENUM_RESULT_FAILURE,
		})
	}
}
