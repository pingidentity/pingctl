package profile_internal

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileAdd(profileName, description string, setActive, setActiveChanged bool) (err error) {
	if profileName == "" {
		profileName, err = promptForName()
		if err != nil {
			return err
		}
	}

	if description == "" {
		description, err = promptForDescription()
		if err != nil {
			return err
		}
	}

	if !setActiveChanged {
		setActive, err = promptForSetActive()
		if err != nil {
			return err
		}
	}

	if err = profiles.CreateNewProfile(profileName, description, setActive); err != nil {
		return err
	}

	return nil
}

func promptForName() (string, error) {
	prompt := promptui.Prompt{
		Label:    "New Profile Name",
		Validate: profiles.ValidateNewProfileName,
	}

	return prompt.Run()
}

func promptForDescription() (string, error) {
	prompt := promptui.Prompt{
		Label: "New Profile Description",
	}

	return prompt.Run()
}

func promptForSetActive() (bool, error) {
	prompt := promptui.Prompt{
		Label:     "Set Profile as Active Profile",
		IsConfirm: true,
		Default:   "n",
	}

	// This is odd behavior discussed in https://github.com/manifoldco/promptui/issues/81
	// If err is type promptui.ErrAbort, the user can be assumed to have responded "No"
	_, err := prompt.Run()
	if err != nil {
		if errors.Is(err, promptui.ErrAbort) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
