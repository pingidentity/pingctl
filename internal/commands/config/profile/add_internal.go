package profile_internal

import (
	"fmt"
	"io"

	"github.com/pingidentity/pingctl/internal/input"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileAdd(profileName, description string, setActive, setActiveChanged bool, r io.ReadCloser) (err error) {
	if profileName == "" {
		if profileName, err = input.RunPrompt("New Profile Name", profiles.ValidateNewProfileName, r); err != nil {
			return err
		}
	} else {
		if err = profiles.ValidateNewProfileName(profileName); err != nil {
			return err
		}
	}

	if description == "" {
		if description, err = input.RunPrompt("New Profile Description", nil, r); err != nil {
			return err
		}
	}

	if !setActiveChanged {
		if setActive, err = input.RunPromptConfirm("Set Profile as Active Profile", r); err != nil {
			return err
		}
	}

	if err = profiles.CreateNewProfile(profileName, description, setActive); err != nil {
		return err
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Profile '%s' added successfully", profileName),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}
