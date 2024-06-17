package viperconfig

import (
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/spf13/viper"
)

func ValidateViperConfig() error {
	if err := validateConfigFileKeys(); err != nil {
		return err
	}

	// Go through all viper configuration keys,
	// and set their values by value type set in the common.ConfigOptions map
	// This will validate the viper configuration heirarchy.
	// NOTE: IF there are invalid values in the config file, but they are overwritten by
	// an env var or flag, the invalid values will not be caught here.
	for _, configOption := range ConfigOptions {
		viperKey := configOption.ViperConfigKey
		switch configOption.VariableType {
		case ENUM_BOOL:
			if err := validateBool(viperKey); err != nil {
				return err
			}
		case ENUM_ID:
			if err := validateUUID(viperKey); err != nil {
				return err
			}
		case ENUM_OUTPUT_FORMAT:
			if err := validateOutputFormat(viperKey); err != nil {
				return err
			}
		case ENUM_PINGONE_REGION:
			if err := validatePingOneRegion(viperKey); err != nil {
				return err
			}
		case ENUM_STRING:
			if err := validateString(viperKey); err != nil {
				return err
			}
		default:
			return fmt.Errorf("failed to validate pingctl configuration: variable type for key '%s' is not recognized", viperKey)
		}
	}
	return nil
}

func validateConfigFileKeys() error {
	// Get all internal viper configuration keys
	viperKeys := []string{}
	for _, configOption := range ConfigOptions {
		viperKey := configOption.ViperConfigKey
		viperKeyToLower := strings.ToLower(viperKey)
		viperKeys = append(viperKeys, viperKeyToLower)
	}

	// Get all keys viper has loaded from config file.
	// If a key found in the config file is not in the viperKeys list,
	// it is an invalid key.
	// Match against lowercase keys as Viper is case insensitive.
	invalidKeys := []string{}
	for _, key := range viper.AllKeys() {
		keyToLower := strings.ToLower(key)
		if !slices.Contains(viperKeys, keyToLower) {
			invalidKeys = append(invalidKeys, key)
		}
	}
	if len(invalidKeys) > 0 {
		invalidKeysStr := strings.Join(invalidKeys, ", ")
		validKeysStr := strings.Join(viperKeys, ", ")
		return fmt.Errorf("failed to validate pingctl configuration: invalid configuration key(s) found in config file: %s\nMust use one of: %s", invalidKeysStr, validKeysStr)
	}
	return nil
}

func validateBool(viperKey string) error {
	value := viper.Get(viperKey)
	switch valueBool := value.(type) {
	case bool:
		viper.Set(viperKey, valueBool)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: variable type for key '%s' is not a boolean value", viperKey)
	}
	return nil
}

func validateUUID(viperKey string) error {
	value := viper.Get(viperKey)
	switch valueUUID := value.(type) {
	case string:
		// Check string is in the form of a UUID or empty
		if valueUUID == "" {
			viper.Set(viperKey, valueUUID)
			return nil
		}

		if _, err := uuid.ParseUUID(valueUUID); err != nil {
			return fmt.Errorf("failed to validate pingctl configuration: value for key '%s' must be a valid UUID", viperKey)
		}
		viper.Set(viperKey, valueUUID)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: value for key '%s' is not a valid UUID", viperKey)
	}
	return nil
}

func validateOutputFormat(viperKey string) error {
	value := viper.Get(viperKey)
	switch valueOutputFormat := value.(type) {
	case customtypes.OutputFormat:
		viper.Set(viperKey, valueOutputFormat)
	case string:
		outputFormat := customtypes.OutputFormat("")
		if err := outputFormat.Set(valueOutputFormat); err != nil {
			return fmt.Errorf("failed to validate pingctl configuration: %s", err.Error())
		}
		viper.Set(viperKey, outputFormat)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: value for key '%s' is not a valid output format. Must use one of: %s", viperKey, strings.Join(customtypes.OutputFormatValidValues(), ", "))
	}
	return nil
}

func validatePingOneRegion(viperKey string) error {
	value := viper.Get(viperKey)
	switch valuePingoneRegion := value.(type) {
	case customtypes.PingOneRegion:
		viper.Set(viperKey, valuePingoneRegion)
	case string:
		region := customtypes.PingOneRegion("")
		if err := region.Set(valuePingoneRegion); err != nil {
			return fmt.Errorf("failed to validate pingctl configuration: %s", err.Error())
		}
		viper.Set(viperKey, region)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: value for key '%s' is not a valid PingOne region. Must use one of: %s", viperKey, strings.Join(customtypes.PingOneRegionValidValues(), ", "))
	}
	return nil
}

func validateString(viperKey string) error {
	value := viper.Get(viperKey)
	switch valueString := value.(type) {
	case string:
		viper.Set(viperKey, valueString)
	default:
		return fmt.Errorf("failed to validate pingctl configuration: variable type for key '%s' is not string", viperKey)
	}
	return nil
}
