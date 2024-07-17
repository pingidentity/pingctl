package profile_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileDelete(args []string) error {
	pName, err := parseDeleteArgs(args)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %v", err)
	}

	err = profiles.DeleteConfigProfile(pName)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %v", err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Profile '%s' deleted successfully", pName),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}

func parseDeleteArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("profile name is required")
	}

	if len(args) > 1 {
		output.Print(output.Opts{
			Message: fmt.Sprintf("'pingctl config profile delete' takes only one argument. Ignoring extra arguments: %s", strings.Join(args[1:], " ")),
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})
	}

	return args[0], nil
}
