package config_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/spf13/viper"
)

func RunInternalConfigSet(kvPair string) (err error) {
	pName, vKey, vValue, err := readConfigSetOptions(kvPair)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	if err = configuration.ValidateViperKey(vKey); err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	// Make sure value is not empty, and suggest unset command if it is
	if vValue == "" {
		return fmt.Errorf("failed to set configuration: value for key '%s' is empty. Use 'pingctl config unset %s' to unset the key", vKey, vKey)
	}

	if err = profiles.GetMainConfig().ValidateExistingProfileName(pName); err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	subViper := profiles.GetMainConfig().ViperInstance().Sub(pName)

	opt, err := configuration.OptionFromViperKey(vKey)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	if err = setValue(subViper, vKey, vValue, opt.Type); err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	if err = profiles.GetMainConfig().SaveProfile(pName, subViper); err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	yamlStr, err := profiles.GetMainConfig().ProfileToString(pName)
	if err != nil {
		return fmt.Errorf("failed to set configuration: %v", err)
	}

	output.Print(output.Opts{
		Message: "Configuration set successfully",
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	output.Print(output.Opts{
		Message: yamlStr,
		Result:  output.ENUM_RESULT_NIL,
	})

	return nil
}

func readConfigSetOptions(kvPair string) (pName string, vKey string, vValue string, err error) {
	if pName, err = readConfigSetProfileName(); err != nil {
		return pName, vKey, vValue, err
	}

	if vKey, vValue, err = parseKeyValuePair(kvPair); err != nil {
		return pName, vKey, vValue, err
	}

	return pName, vKey, vValue, nil
}

func readConfigSetProfileName() (pName string, err error) {
	if !configuration.ConfigSetProfileOption.Flag.Changed {
		pName, err = profiles.GetOptionValue(configuration.RootActiveProfileOption)
	} else {
		pName, err = profiles.GetOptionValue(configuration.ConfigSetProfileOption)
	}

	if err != nil {
		return pName, err
	}

	if pName == "" {
		return pName, fmt.Errorf("unable to determine profile to set configuration to")
	}

	return pName, nil
}

func parseKeyValuePair(kvPair string) (string, string, error) {
	parsedInput := strings.SplitN(kvPair, "=", 2)
	if len(parsedInput) < 2 {
		return "", "", fmt.Errorf("invalid assignment format '%s'. Expect 'key=value' format", kvPair)
	}

	return parsedInput[0], parsedInput[1], nil
}

func setValue(profileViper *viper.Viper, vKey, vValue string, valueType configuration.OptionType) (err error) {
	switch valueType {
	case configuration.ENUM_BOOL:
		var bool customtypes.Bool
		if err = bool.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a boolean. Allowed [true, false]: %v", vKey, err)
		}
		profileViper.Set(vKey, bool)
	case configuration.ENUM_EXPORT_FORMAT:
		var exportFormat customtypes.ExportFormat
		if err = exportFormat.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid export format. Allowed [%s]: %v", vKey, strings.Join(customtypes.ExportFormatValidValues(), ", "), err)
		}
		profileViper.Set(vKey, exportFormat)
	case configuration.ENUM_MULTI_SERVICE:
		var multiService customtypes.MultiService
		if err = multiService.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid multi-service. Allowed [%s]: %v", vKey, strings.Join(customtypes.MultiServiceValidValues(), ", "), err)
		}
		profileViper.Set(vKey, multiService)
	case configuration.ENUM_OUTPUT_FORMAT:
		var outputFormat customtypes.OutputFormat
		if err = outputFormat.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid output format. Allowed [%s]: %v", vKey, strings.Join(customtypes.OutputFormatValidValues(), ", "), err)
		}
		profileViper.Set(vKey, outputFormat)
	case configuration.ENUM_PINGONE_REGION:
		var region customtypes.PingOneRegion
		if err = region.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid PingOne region. Allowed [%s]: %v", vKey, strings.Join(customtypes.PingOneRegionValidValues(), ", "), err)
		}
		profileViper.Set(vKey, region)
	case configuration.ENUM_STRING:
		var str customtypes.String
		if err = str.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a string: %v", vKey, err)
		}
		profileViper.Set(vKey, str)
	case configuration.ENUM_STRING_SLICE:
		var strSlice customtypes.StringSlice
		if err = strSlice.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a string slice: %v", vKey, err)
		}
		profileViper.Set(vKey, strSlice)
	case configuration.ENUM_UUID:
		var uuid customtypes.UUID
		if err = uuid.Set(vValue); err != nil {
			return fmt.Errorf("value for key '%s' must be a valid UUID: %v", vKey, err)
		}
		profileViper.Set(vKey, uuid)
	default:
		return fmt.Errorf("failed to set configuration: variable type for key '%s' is not recognized", vKey)
	}

	return nil
}
