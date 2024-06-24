package config_internal

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/viper"
)

func RunInternalConfigUnset(args []string) error {
	// Parse the viper key from the command line arguments
	viperKey, err := parseUnsetArgs(args)
	if err != nil {
		return err
	}

	// Check if the key is a valid viper configuration key
	if !viperconfig.IsValidViperKey(viperKey) {
		validKeys := viperconfig.GetViperConfigKeys()
		slices.Sort(validKeys)
		validKeysStr := strings.Join(validKeys, ", ")
		return fmt.Errorf("unable to unset configuration: key '%s' is not recognized as a valid configuration key. Valid keys: %s", viperKey, validKeysStr)
	}

	valueType, ok := viperconfig.GetValueTypeFromViperKey(viperKey)
	if !ok {
		return fmt.Errorf("failed to unset configuration: value type for key %s unrecognized", viperKey)
	}

	if err := unsetValue(viperKey, valueType); err != nil {
		return err
	}

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write pingctl configuration to file '%s': %s", viper.ConfigFileUsed(), err.Error())
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
		output.Format(output.CommandOutput{
			Message: fmt.Sprintf("'pingctl config unset' can only unset one key per command. Ignoring extra arguments: %s", strings.Join(args[1:], " ")),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})
	}

	// Assume viper configuration key is args[0] and ignore any other input
	return args[0], nil
}

func unsetValue(viperKey string, valueType viperconfig.ConfigType) error {
	switch valueType {
	case viperconfig.ENUM_BOOL:
		viper.Set(viperKey, false)
	case viperconfig.ENUM_ID:
		viper.Set(viperKey, string(""))
	case viperconfig.ENUM_OUTPUT_FORMAT:
		viper.Set(viperKey, customtypes.OutputFormat(""))
	case viperconfig.ENUM_PINGONE_REGION:
		viper.Set(viperKey, customtypes.PingOneRegion(""))
	case viperconfig.ENUM_STRING:
		viper.Set(viperKey, string(""))
	default:
		return fmt.Errorf("unable to unset configuration: variable type for key '%s' is not recognized", viperKey)
	}
	return nil
}
