/*
Copyright © 2024 Ping Identity Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
)

// Test Root Command Executes without issue
func TestRootCmd_Execute(t *testing.T) {
	// Create the root command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
}

// Test Root Command Executes output does not change with output=json
func TestRootCmd_JSONOutput(t *testing.T) {
	// Create the root command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	outputWithoutJSON := stdout.String()

	// Create the root command
	rootCmd = cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	stdout = bytes.Buffer{}
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)
	rootCmd.SetArgs([]string{"--output=json"})

	// Execute the root command
	err = rootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	outputWithJSON := stdout.String()

	//expect both outputs to be the same
	if outputWithJSON != outputWithoutJSON {
		t.Errorf("Expected no change on output with json specified, got %q VS %q", outputWithoutJSON, outputWithJSON)
	}
}
