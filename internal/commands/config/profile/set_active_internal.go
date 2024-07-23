package profile_internal

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileSetActive(profileName string) (err error) {
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
