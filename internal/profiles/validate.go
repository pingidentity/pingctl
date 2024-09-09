package profiles

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/spf13/viper"
)

func Validate() error {
	// Get a slice of all profile names configured in the config.yaml file
	profileNames := GetMainConfig().ProfileNames()

	// Validate profile names
	if err := validateProfileNames(profileNames); err != nil {
		return err
	}

	// Make sure selected active profile is in the configuration file
	activeProfile := GetMainConfig().ActiveProfile().Name()
	if !slices.Contains(profileNames, activeProfile) {
		return fmt.Errorf("failed to validate pingctl configuration: active profile '%s' not found in configuration file %s", activeProfile, GetMainConfig().ViperInstance().ConfigFileUsed())
	}

	// for each profile key, set the profile based on mainViper.Sub() and validate the profile
	for _, pName := range profileNames {
		subViper := GetMainConfig().ViperInstance().Sub(pName)

		if err := validateProfileKeys(pName, subViper); err != nil {
			return fmt.Errorf("failed to validate pingctl configuration: %v", err)
		}

		if err := validateProfileValues(pName, subViper); err != nil {
			return fmt.Errorf("failed to validate pingctl configuration: %v", err)
		}
	}

	return nil
}

func validateProfileNames(profileNames []string) error {
	for _, profileName := range profileNames {
		if err := GetMainConfig().ValidateProfileNameFormat(profileName); err != nil {
			return err
		}
	}
	return nil
}

func validateProfileKeys(profileName string, profileViper *viper.Viper) error {
	validProfileKeys := configuration.ViperKeys()

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
		return fmt.Errorf("invalid configuration key(s) found in profile %s: %s\nMust use one of: %s", profileName, invalidKeysStr, validKeysStr)
	}
	return nil
}

func validateProfileValues(pName string, profileViper *viper.Viper) (err error) {
	for _, key := range profileViper.AllKeys() {
		opt, err := configuration.OptionFromViperKey(key)
		if err != nil {
			return err
		}

		vValue := profileViper.Get(key)

		switch opt.Type {
		case configuration.ENUM_BOOL:
			switch typedValue := vValue.(type) {
			case *customtypes.Bool:
				continue
			case string:
				b := customtypes.Bool(false)
				if err = b.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a boolean value: %v", pName, typedValue, key, err)
				}
			case bool:
				continue
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a boolean value", pName, typedValue, key)
			}
		case configuration.ENUM_UUID:
			switch typedValue := vValue.(type) {
			case *customtypes.UUID:
				continue
			case string:
				u := customtypes.UUID("")
				if err = u.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a UUID value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a UUID value", pName, typedValue, key)
			}
		case configuration.ENUM_OUTPUT_FORMAT:
			switch typedValue := vValue.(type) {
			case *customtypes.OutputFormat:
				continue
			case string:
				o := customtypes.OutputFormat("")
				if err = o.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not an output format value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not an output format value", pName, typedValue, key)
			}
		case configuration.ENUM_PINGONE_REGION:
			switch typedValue := vValue.(type) {
			case *customtypes.PingOneRegion:
				continue
			case string:
				p := customtypes.PingOneRegion("")
				if err = p.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a PingOne region value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a PingOne region value", pName, typedValue, key)
			}
		case configuration.ENUM_STRING:
			switch typedValue := vValue.(type) {
			case *customtypes.String:
				continue
			case string:
				s := customtypes.String("")
				if err = s.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a string value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a string value", pName, typedValue, key)
			}
		case configuration.ENUM_STRING_SLICE:
			switch typedValue := vValue.(type) {
			case *customtypes.StringSlice:
				continue
			case string:
				ss := customtypes.StringSlice([]string{})
				if err = ss.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a string slice value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a string slice value", pName, typedValue, key)
			}
		case configuration.ENUM_MULTI_SERVICE:
			switch typedValue := vValue.(type) {
			case *customtypes.MultiService:
				continue
			case string:
				ms := customtypes.NewMultiService()
				if err = ms.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not a multi-service value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not a multi-service value", pName, typedValue, key)
			}
		case configuration.ENUM_EXPORT_FORMAT:
			switch typedValue := vValue.(type) {
			case *customtypes.ExportFormat:
				continue
			case string:
				ef := customtypes.ExportFormat("")
				if err = ef.Set(typedValue); err != nil {
					return fmt.Errorf("profile '%s': variable type '%T' for key '%s' is not an export format value: %v", pName, typedValue, key, err)
				}
			default:
				return fmt.Errorf("profile '%s': variable type %T for key '%s' is not an export format value", pName, typedValue, key)
			}
		default:
			return fmt.Errorf("profile '%s': variable type '%s' for key '%s' is not recognized", pName, opt.Type, key)
		}
	}

	return nil
}
