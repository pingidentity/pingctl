package profile_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileDescribe(args []string) error {
	pName, err := parseDescArgs(args)
	if err != nil {
		return fmt.Errorf("failed to describe profile: %v", err)
	}

	err = profiles.ValidateExistingProfileName(pName)
	if err != nil {
		return fmt.Errorf("failed to describe profile: %v", err)
	}

	// Create temp sub viper for profile
	mainViper := profiles.GetMainViper()
	tempViper := mainViper.Sub(pName)

	descStr := fmt.Sprintf("Profile Name: %s\n", pName)
	descStr += fmt.Sprintf("Description: %s\n\n", tempViper.GetString(profiles.ProfileDescriptionOption.ViperKey))

	setStr := ""
	unsetStr := ""
	for _, opt := range profiles.ConfigOptions.Options {
		if opt.ViperKey == profiles.ProfileDescriptionOption.ViperKey || opt.ViperKey == profiles.ProfileOption.ViperKey {
			continue
		}

		vValue := tempViper.Get(opt.ViperKey)
		if vValue == nil || vValue == "" {
			unsetStr += fmt.Sprintf(" - %s\n", opt.ViperKey)
		} else {
			setStr += fmt.Sprintf(" - %s: %v\n", opt.ViperKey, vValue)
		}

	}

	if setStr != "" {
		descStr += fmt.Sprintf("Set Options:\n%s\n", setStr)
	}

	if unsetStr != "" {
		descStr += fmt.Sprintf("Unset Options:\n%s\n", unsetStr)
	}

	output.Print(output.Opts{
		Message: descStr,
		Result:  output.ENUM_RESULT_NIL,
	})

	return nil
}

func parseDescArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("profile name is required")
	}

	if len(args) > 1 {
		output.Print(output.Opts{
			Message: fmt.Sprintf("'pingctl config profile describe' takes only one argument. Ignoring extra arguments: %s", strings.Join(args[1:], " ")),
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})
	}

	return args[0], nil
}