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
func TestFeedbackCmd_Execute(t *testing.T) {
	// Create the root command
	feedbackCmd := cmd.NewFeedbackCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	feedbackCmd.SetOut(&stdout)
	feedbackCmd.SetErr(&stdout)

	// Execute the root command
	err := feedbackCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
