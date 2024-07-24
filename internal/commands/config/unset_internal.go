package config_internal

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigUnset(viperKey string) error {
	// Check if the key is a valid viper configuration key
	validKeys := profiles.ProfileKeys()
	if !slices.ContainsFunc(validKeys, func(v string) bool {
		return strings.EqualFold(v, viperKey)
	}) {
		slices.Sort(validKeys)
		validKeysStr := strings.Join(validKeys, ", ")
		return fmt.Errorf("unable to unset configuration: key '%s' is not recognized as a valid configuration key. Valid keys: %s", viperKey, validKeysStr)
	}

	valueType, ok := profiles.OptionTypeFromViperKey(viperKey)
	if !ok {
		return fmt.Errorf("failed to unset configuration: value type for key %s unrecognized", viperKey)
	}

	profiles.GetProfileViper().Set(viperKey, profiles.GetDefaultValue(valueType))

	if err := profiles.SaveProfileViperToFile(); err != nil {
		return err
	}

	if err := PrintConfig(); err != nil {
		return err
	}

	return nil
}
