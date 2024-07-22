package profile_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileSetActive(args []string) error {
	profileName, err := parseSetActiveArgs(args)
	if err != nil {
		return fmt.Errorf("failed to set active profile: %v", err)
	}

	if err = profiles.ValidateExistingProfileName(profileName); err != nil {
		return fmt.Errorf("failed to set active profile: %v", err)
	}

	if err = profiles.SetConfigActiveProfile(profileName); err != nil {
		return fmt.Errorf("failed to set active profile: %v", err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Active configuration profile set to '%s'", profileName),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}

func parseSetActiveArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("profile name is required")
	}

	if len(args) > 1 {
		output.Print(output.Opts{
			Message: fmt.Sprintf("'pingctl config profile set-active' only sets one profile as active per command. Ignoring extra arguments: %s", strings.Join(args[1:], " ")),
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})
	}

	profileName := args[0]

	return profileName, nil
}
