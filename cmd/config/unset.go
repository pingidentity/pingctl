package config

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Config Unset Subcommand...")
}

func NewConfigUnsetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unset",
		Short: "Unset pingctl configuration settings.",
		Long: `Unset pingctl configuration settings.
		
Example command usage: 'pingctl config unset pingctl.color'`,
		RunE: ConfigUnsetRunE,
	}

	return cmd
}
func ConfigUnsetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	// Parse the viper key from the command line arguments
	viperKey, err := parseUnsetArgs(args)
	if err != nil {
		return err
	}

	// Check if the key is a valid viper configuration key
	if !viperconfig.IsValidViperKey(viperKey) {
		validKeys := strings.Join(viperconfig.GetViperConfigKeys(), ", ")
		return fmt.Errorf("unable to unset configuration: key '%s' is not recognized as a valid configuration key. \nValid keys: %s", viperKey, validKeys)
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

	if err := printConfig(cmd); err != nil {
		return err
	}

	return nil
}

func parseUnsetArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("unable to unset configuration: no key given in unset command")
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
