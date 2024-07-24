package config_internal

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigSet(kvPair string) error {
	// Parse the key=value pair from the command line arguments
	viperKey, value, err := parseKeyValuePair(kvPair)
	if err != nil {
		return err
	}

	// Check if the key is a valid viper configuration key
	validKeys := profiles.ProfileKeys()
	if !slices.ContainsFunc(validKeys, func(v string) bool {
		return strings.EqualFold(v, viperKey)
	}) {
		slices.Sort(validKeys)
		validKeysStr := strings.Join(validKeys, ", ")
		return fmt.Errorf("failed to set configuration: key '%s' is not recognized as a valid configuration key. Valid keys: %s", viperKey, validKeysStr)
	}

	// Make sure value is not empty, and suggest unset command if it is
	if value == "" {
		return fmt.Errorf("failed to set configuration: value for key '%s' is empty. Use 'pingctl config unset %s' to unset the key", viperKey, viperKey)
	}

	valueType, ok := profiles.OptionTypeFromViperKey(viperKey)
	if !ok {
		return fmt.Errorf("failed to set configuration: value type for key %s unrecognized", viperKey)
	}

	if err := setValue(viperKey, value, valueType); err != nil {
		return err
	}

	if err := profiles.SaveProfileViperToFile(); err != nil {
		return err
	}

	if err := PrintConfig(); err != nil {
		return err
	}

	return nil
}

func parseKeyValuePair(kvPair string) (string, string, error) {
	parsedInput := strings.SplitN(kvPair, "=", 2)
	if len(parsedInput) != 2 {
		return "", "", fmt.Errorf("failed to set configuration: invalid assignment format '%s'. Expect 'key=value' format", kvPair)
	}

	return parsedInput[0], parsedInput[1], nil
}

func setValue(viperKey, value string, valueType profiles.OptionType) error {
	switch valueType {
	case profiles.ENUM_BOOL:
		return setBool(viperKey, value)
	case profiles.ENUM_ID:
		return setUUID(viperKey, value)
	case profiles.ENUM_OUTPUT_FORMAT:
		return setOutputFormat(viperKey, value)
	case profiles.ENUM_PINGONE_REGION:
		return setPingOneRegion(viperKey, value)
	case profiles.ENUM_STRING:
		profiles.GetProfileViper().Set(viperKey, string(value))
		return nil
	default:
		return fmt.Errorf("failed to set configuration: variable type for key '%s' is not recognized", viperKey)
	}
}

func setBool(viperKey string, value string) error {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("failed to set configuration: value for key '%s' must be a boolean. Use 'true' or 'false'", viperKey)
	}

	profiles.GetProfileViper().Set(viperKey, boolValue)

	return nil
}

func setUUID(viperKey string, value string) error {
	// Check string is in the form of a UUID
	if _, err := uuid.ParseUUID(value); err != nil {
		return fmt.Errorf("failed to set configuration: value for key '%s' must be a valid UUID", viperKey)
	}

	profiles.GetProfileViper().Set(viperKey, string(value))

	return nil
}

func setOutputFormat(viperKey string, value string) error {
	outputFormat := customtypes.OutputFormat("")
	if err := outputFormat.Set(value); err != nil {
		return fmt.Errorf("failed to set configuration: %s", err.Error())
	}

	profiles.GetProfileViper().Set(viperKey, outputFormat)

	return nil
}

func setPingOneRegion(viperKey string, value string) error {
	region := customtypes.PingOneRegion("")
	if err := region.Set(value); err != nil {
		return fmt.Errorf("failed to set configuration: %s", err.Error())
	}

	profiles.GetProfileViper().Set(viperKey, region)

	return nil
}
