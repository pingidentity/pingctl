package profiles

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Binding struct {
	Option Option
	Flag   *pflag.Flag
}

var (
	profileViper   *viper.Viper
	profileName    string
	flagBindings   []Binding
	pFlagBindings  []Binding
	envVarBindings []Option
)

func SetProfileViperWithProfile(pName string) (err error) {
	if err := ValidateExistingProfileName(pName); err != nil {
		return err
	}

	subViper := mainViper.Sub(pName)
	if subViper == nil {
		return fmt.Errorf("profile '%s' not found in configuration file: %s", pName, mainViper.ConfigFileUsed())
	}

	profileViper = subViper
	profileName = pName
	return nil
}

func SetProfileViperWithViper(v *viper.Viper, pName string) {
	profileViper = v
	profileName = pName
}

func GetProfileViper() *viper.Viper {
	return profileViper
}

func AddFlagBinding(binding Binding) {
	flagBindings = append(flagBindings, binding)
}

func AddPFlagBinding(binding Binding) {
	pFlagBindings = append(pFlagBindings, binding)
}

func AddEnvVarBinding(opt Option) {
	envVarBindings = append(envVarBindings, opt)
}

func ApplyBindingsToProfileViper() (err error) {
	for _, binding := range flagBindings {
		err = profileViper.BindPFlag(binding.Option.ViperKey, binding.Flag)
		if err != nil {
			return err
		}
	}

	for _, binding := range pFlagBindings {
		err = profileViper.BindPFlag(binding.Option.ViperKey, binding.Flag)
		if err != nil {
			return err
		}
	}

	for _, opt := range envVarBindings {
		err = profileViper.BindEnv(opt.ViperKey, opt.EnvVar)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateNewProfile(pName, desc string, setActive bool) (err error) {
	err = ValidateNewProfileName(pName)
	if err != nil {
		return err
	}

	oldProfileName := profileName
	newProfileViper := viper.New()

	for _, opt := range ConfigOptions.Options {
		if opt.ViperKey == ProfileOption.ViperKey {
			continue
		}

		// Set all options to their default state
		newProfileViper.Set(opt.ViperKey, GetDefaultValue(opt.Type))
	}

	// set the new profile description
	newProfileViper.Set(ProfileDescriptionOption.ViperKey, desc)

	// set the new viper as the profile viper
	SetProfileViperWithViper(newProfileViper, pName)

	// save the new profile to the configuration file
	if err := SaveProfileViperToFile(); err != nil {
		return fmt.Errorf("failed to create new profile: %v", err)
	}

	// set the profile viper back to the old profile if it existed
	if oldProfileName != "" {
		if err = SetProfileViperWithProfile(oldProfileName); err != nil {
			return fmt.Errorf("failed to create new profile: %v", err)
		}
	}

	// set the new profile as the active profile if applicable
	if setActive {
		if err = SetConfigActiveProfile(pName); err != nil {
			return fmt.Errorf("failed to create new profile: %v", err)
		}
	}

	return nil
}
