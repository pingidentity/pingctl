package profiles

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/spf13/viper"
)

func Validate() error {
	// Get a slice of all profile names configured in the config.yaml file
	profileNames := configProfileNames()

	// Validate profile names
	if err := validateProfileNames(profileNames); err != nil {
		return err
	}

	// for each profile key, set the profile based on mainViper.Sub() and validate the profile
	for _, name := range profileNames {
		v := mainViper.Sub(name)

		if err := validateProfileKeys(name, v); err != nil {
			return err
		}

		if err := validateProfileValues(v); err != nil {
			return err
		}
	}

	return nil
}

func validateProfileNames(profileKeys []string) error {
	for _, profileKey := range profileKeys {
		re := regexp.MustCompile(`^[a-zA-Z0-9\_\-]+$`)
		if !re.MatchString(profileKey) {
			return fmt.Errorf("failed to validate pingctl configuration: profile name '%s' must contain only alphanumeric characters, underscores, and dashes", profileKey)
		}
	}
	return nil
}

func configProfileNames() []string {
	allKeys := mainViper.AllKeys()

	// Get a slice of all profile keys to validate profiles individually
	var profileKeys []string
	for _, key := range allKeys {
		//remove "activeProfile" from profileKeys
		if strings.EqualFold(key, ProfileOption.ViperKey) {
			continue
		}

		rootKey := strings.Split(key, ".")[0]
		if !slices.Contains(profileKeys, rootKey) {
			profileKeys = append(profileKeys, rootKey)
		}
	}
	return profileKeys
}

func validateProfileKeys(profileName string, profileViper *viper.Viper) error {
	validProfileKeys := ProfileKeys()

	// Get all keys viper has loaded from config file.
	// If a key found in the config file is not in the viperKeys list,
	// it is an invalid key.
	var invalidKeys []string
	for _, key := range profileViper.AllKeys() {
		if !slices.ContainsFunc(validProfileKeys, func(v string) bool {
			return strings.EqualFold(v, key)
		}) {
			invalidKeys = append(invalidKeys, key)
		}
	}

	if len(invalidKeys) > 0 {
		invalidKeysStr := strings.Join(invalidKeys, ", ")
		validKeysStr := strings.Join(validProfileKeys, ", ")
		return fmt.Errorf("failed to validate pingctl configuration: invalid configuration key(s) found in profile %s: %s\nMust use one of: %s", profileName, invalidKeysStr, validKeysStr)
	}
	return nil
}

func validateProfileValues(profileViper *viper.Viper) error {
	// Go through all valid profile keys,
	// and set their values by value type set in the ConfigOptions map
	// This will validate the viper configuration hierarchy.
	// NOTE: IF there are invalid values in the config file, but they are overwritten by
	// an env var or flag, the invalid values will not be caught here.
	for _, opt := range ConfigOptions.Options {
		viperKey := opt.ViperKey
		if viperKey == ProfileOption.ViperKey {
			continue
		}
		switch opt.Type {
		case ENUM_BOOL:
			if err := validateBool(profileViper, viperKey); err != nil {
				return err
			}
		case ENUM_ID:
			if err := validateUUID(profileViper, viperKey); err != nil {
				return err
			}
		case ENUM_OUTPUT_FORMAT:
			if err := validateOutputFormat(profileViper, viperKey); err != nil {
				return err
			}
		case ENUM_PINGONE_REGION:
			if err := validatePingOneRegion(profileViper, viperKey); err != nil {
				return err
			}
		case ENUM_STRING:
			if err := validateString(profileViper, viperKey); err != nil {
				return err
			}
		default:
			return fmt.Errorf("failed to validate pingctl configuration: variable type for key '%s' is not recognized", viperKey)
		}
	}
	return nil
}

func validateBool(profileViper *viper.Viper, viperKey string) error {
	value := profileViper.Get(viperKey)
	switch valueBool := value.(type) {
	case bool:
		profileViper.Set(viperKey, valueBool)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: variable type for key '%s' is not a boolean value", viperKey)
	}
	return nil
}

func validateUUID(profileViper *viper.Viper, viperKey string) error {
	value := profileViper.Get(viperKey)
	switch valueUUID := value.(type) {
	case string:
		// Check string is in the form of a UUID or empty
		if valueUUID == "" {
			profileViper.Set(viperKey, valueUUID)
			return nil
		}

		if _, err := uuid.ParseUUID(valueUUID); err != nil {
			return fmt.Errorf("failed to validate pingctl configuration: value for key '%s' must be a valid UUID", viperKey)
		}
		profileViper.Set(viperKey, valueUUID)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: value for key '%s' is not a valid UUID", viperKey)
	}
	return nil
}

func validateOutputFormat(profileViper *viper.Viper, viperKey string) error {
	value := profileViper.Get(viperKey)
	switch valueOutputFormat := value.(type) {
	case customtypes.OutputFormat:
		profileViper.Set(viperKey, valueOutputFormat)
	case string:
		outputFormat := customtypes.OutputFormat("")
		if valueOutputFormat != "" { // Allow empty string for output format validation
			if err := outputFormat.Set(valueOutputFormat); err != nil {
				return fmt.Errorf("failed to validate pingctl configuration: %s", err.Error())
			}
		}
		profileViper.Set(viperKey, outputFormat)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: value for key '%s' is not a valid output format. Must use one of: %s", viperKey, strings.Join(customtypes.OutputFormatValidValues(), ", "))
	}
	return nil
}

func validatePingOneRegion(profileViper *viper.Viper, viperKey string) error {
	value := profileViper.Get(viperKey)
	switch valuePingoneRegion := value.(type) {
	case customtypes.PingOneRegion:
		profileViper.Set(viperKey, valuePingoneRegion)
	case string:
		region := customtypes.PingOneRegion("")
		if valuePingoneRegion != "" { // Allow empty string for pingone region validation
			if err := region.Set(valuePingoneRegion); err != nil {
				return fmt.Errorf("failed to validate pingctl configuration: %s", err.Error())
			}
		}
		profileViper.Set(viperKey, region)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: value for key '%s' is not a valid PingOne region. Must use one of: %s", viperKey, strings.Join(customtypes.PingOneRegionValidValues(), ", "))
	}
	return nil
}

func validateString(profileViper *viper.Viper, viperKey string) error {
	value := profileViper.Get(viperKey)
	switch valueString := value.(type) {
	case string:
		profileViper.Set(viperKey, valueString)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: variable type for key '%s' is not string", viperKey)
	}
	return nil
}
