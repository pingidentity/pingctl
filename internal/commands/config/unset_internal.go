package config_internal

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigUnset(args []string) error {
	// Parse the viper key from the command line arguments
	viperKey, err := parseUnsetArgs(args)
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

func parseUnsetArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("unable to unset configuration: no key given in unset command")
	}

	if len(args) > 1 {
		output.Print(output.Opts{
			Message: fmt.Sprintf("'pingctl config unset' can only unset one key per command. Ignoring extra arguments: %s", strings.Join(args[1:], " ")),
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})
	}

	// Assume viper configuration key is args[0] and ignore any other input
	return args[0], nil
}
