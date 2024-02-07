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
package cmd

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/spf13/cobra"
)

func NewFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feedback",
		Short: "Information on tool feedback",
		Long: `A command to provide the user information
			on how to give feedback or get help with the tool
			through the use of the GitHub repository's issue tracker.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			l := logger.Get()

			l.Debug().Msgf("Feedback command called.")

			feedbackMessage := `Thank you for participating in early adoption of the refreshed Ping Identity universal CLI.

We appreciate your feedback and information regarding your expirience with the CLI.

Please visit the following URL in your browser and let us know of any feedback or issues related
to the tool:

	https://github.com/pingidentity/pingctl/issues/new

`
			output.Format(output.CommandOutput{
				Message: feedbackMessage,
				Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
				Command: cmd,
			})

			return nil
		},
	}

	return cmd
}
