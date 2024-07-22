package profiles

/* The main viper instance should ONLY interact with the configuration file
on disk. No viper overrides, environment variable bindings, or pflag
bindings should be used with this viper instance. This keeps the config
file as the ONLY source of truth for the main viper instance, and prevents
profile drift, as well as active profile drift and other niche bugs. As a
result, much of the logic in this file avoids the use of mainViper.Set(), and
goes out of the way to modify the config file.*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/spf13/viper"
)

var (
	mainViper *viper.Viper = viper.New()
)

func GetMainViper() *viper.Viper {
	return mainViper
}

func GetConfigActiveProfile() string {
	return mainViper.GetString(ProfileOption.ViperKey)
}

func SetConfigActiveProfile(pName string) (err error) {
	tempViper := viper.New()
	tempViper.SetConfigFile(mainViper.ConfigFileUsed())
	if err = tempViper.ReadInConfig(); err != nil {
		return err
	}

	tempViper.Set(ProfileOption.ViperKey, pName)
	if err = tempViper.WriteConfig(); err != nil {
		return err
	}

	if err = mainViper.ReadInConfig(); err != nil {
		return err
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
	if slices.ContainsFunc(pNames, func(n string) bool {
		return strings.EqualFold(n, pName)
	}) {
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
	if !slices.ContainsFunc(pNames, func(n string) bool {
		return strings.EqualFold(n, pName)
	}) {
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

// Viper gives no built-in delete or unset method for keys
// Using this "workaround" described here: https://github.com/spf13/viper/issues/632
func DeleteConfigProfile(pName string) (err error) {
	err = ValidateExistingProfileName(pName)
	if err != nil {
		return err
	}

	if pName == GetConfigActiveProfile() {
		return fmt.Errorf("'%s' is the active profile and cannot be deleted", pName)
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

func SaveProfileViperToFile() (err error) {
	tempViper := viper.New()
	tempViper.SetConfigFile(mainViper.ConfigFileUsed())
	if err = tempViper.ReadInConfig(); err != nil {
		return err
	}

	profileKeys := profileViper.AllKeys()
	for _, key := range profileKeys {
		tempViper.Set(fmt.Sprintf("%s.%s", profileName, key), profileViper.Get(key))
	}

	if err := tempViper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write pingctl configuration to file '%s': %v", mainViper.ConfigFileUsed(), err)
	}

	if err := mainViper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read pingctl configuration from file '%s': %v", mainViper.ConfigFileUsed(), err)
	}

	return nil
}
