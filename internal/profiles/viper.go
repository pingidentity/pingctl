package profiles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Binding struct {
	Option Option
	Flag   *pflag.Flag
}

var (
	mainViper      *viper.Viper = viper.New()
	profileViper   *viper.Viper
	profileName    string
	flagBindings   []Binding
	pFlagBindings  []Binding
	envVarBindings []Option
)

func GetMainViper() *viper.Viper {
	return mainViper
}

func GetConfigActiveProfile() string {
	return mainViper.GetString(ProfileOption.ViperKey)
}

func SetConfigActiveProfile(pName string) (err error) {
	mainViper.Set(ProfileOption.ViperKey, pName)
	if err = mainViper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to set active profile: %v", err)
	}

	return nil
}

func SetProfileViperWithProfile(pName string) (err error) {
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

func SaveProfileViperToFile() error {
	profileKeys := profileViper.AllKeys()
	for _, key := range profileKeys {
		mainViper.Set(fmt.Sprintf("%s.%s", profileName, key), profileViper.Get(key))
	}

	if err := mainViper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write pingctl configuration to file '%s': %v", mainViper.ConfigFileUsed(), err)
	}

	return nil
}

// Get all profile names from config.yaml configuration file
func ConfigProfileNames() (profileKeys []string) {
	keySet := make(map[string]struct{})
	allKeys := mainViper.AllKeys()

	for _, key := range allKeys {
		//remove "activeProfile" from profileKeys
		if strings.EqualFold(key, ProfileOption.ViperKey) {
			continue
		}

		rootKey := strings.Split(key, ".")[0]
		if _, ok := keySet[rootKey]; !ok {
			keySet[rootKey] = struct{}{}
			profileKeys = append(profileKeys, rootKey)
		}
	}

	slices.Sort(profileKeys)
	return profileKeys
}

// The profile name must contain only alphanumeric characters, underscores, and dashes
// The profile name cannot be empty
// The profile name must be unique
func ValidateNewProfileName(pName string) (err error) {
	if pName == "" {
		return fmt.Errorf("invalid profile name: profile name cannot be empty")
	}

	if err := ValidateProfileNameFormat(pName); err != nil {
		return err
	}

	pNames := ConfigProfileNames()
	if slices.Contains(pNames, pName) {
		return fmt.Errorf("invalid profile name: '%s' profile already exists", pName)
	}

	return nil
}

func ValidateExistingProfileName(pName string) (err error) {
	if pName == "" {
		return fmt.Errorf("invalid profile name: profile name cannot be empty")
	}

	if err := ValidateProfileNameFormat(pName); err != nil {
		return err
	}

	pNames := ConfigProfileNames()
	if !slices.Contains(pNames, pName) {
		return fmt.Errorf("invalid profile name: '%s' profile does not exist", pName)
	}

	return nil
}

func ValidateProfileNameFormat(pName string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9\_\-]+$`)
	if !re.MatchString(pName) {
		return fmt.Errorf("invalid profile name: '%s'. name must contain only alphanumeric characters, underscores, and dashes", pName)
	}

	return nil
}

func CreateNewProfile(pName, desc string, setActive bool) (err error) {
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

// Viper gives no built-in delete or unset method for keys
// Using this "hack" described here: https://github.com/spf13/viper/issues/632
func DeleteConfigProfile(pName string) (err error) {
	err = ValidateExistingProfileName(pName)
	if err != nil {
		return err
	}

	if pName == GetConfigActiveProfile() {
		return fmt.Errorf("cannot delete active profile")
	}

	configMap := mainViper.AllSettings()
	delete(configMap, pName)

	encodedConfig, err := json.MarshalIndent(configMap, "", " ")
	if err != nil {
		return err
	}

	err = mainViper.ReadConfig(bytes.NewReader(encodedConfig))
	if err != nil {
		return err
	}

	err = mainViper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}
