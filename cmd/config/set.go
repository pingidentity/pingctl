package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Config Set Subcommand...")
}

func NewConfigSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set pingctl configuration settings.",
		Long: `Set pingctl configuration settings.

Example command usage: 'pingctl config set pingctl.color=false'`,
		RunE: ConfigSetRunE,
	}

	return cmd
}
func ConfigSetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	// Parse the key=value pair from the command line arguments
	viperKey, value, err := parseSetArgs(args, cmd)
	if err != nil {
		return err
	}

	// Check if the key is a valid viper configuration key
	if !viperconfig.IsValidViperKey(viperKey) {
		validKeys := strings.Join(viperconfig.GetViperConfigKeys(), ", ")
		return fmt.Errorf("unable to unset configuration: key '%s' is not recognized as a valid configuration key. \nValid keys: %s", viperKey, validKeys)
	}

	// Make sure value is not empty, and suggest unset command if it is
	if value == "" {
		return fmt.Errorf("failed to set configuration: value for key '%s' is empty. Use 'pingctl config unset %s' to unset the key", viperKey, viperKey)
	}

	valueType, ok := viperconfig.GetValueTypeFromViperKey(viperKey)
	if !ok {
		return fmt.Errorf("failed to set configuration: value type for key %s unrecognized", viperKey)
	}

	if err := setValue(viperKey, value, valueType); err != nil {
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

func parseSetArgs(args []string, cmd *cobra.Command) (string, string, error) {
	if len(args) == 0 {
		return "", "", fmt.Errorf("failed to set configuration: no 'key=value' assignment given in set command")
	}

	if len(args) > 1 {
		output.Format(cmd, output.CommandOutput{
			Message: fmt.Sprintf("'pingctl config set' only sets one key-value pair per commands. Ignoring extra arguments: %s", strings.Join(args[1:], " ")),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})
	}

	// Assume viper configuration key=value pair is args[0] and ignore any other input
	parsedInput := strings.SplitN(args[0], "=", 2)
	if len(parsedInput) != 2 {
		return "", "", fmt.Errorf("failed to set configuration: invalid assignment format '%s'. Expect 'key=value' format", args[0])
	}

	return parsedInput[0], parsedInput[1], nil
}

func setValue(viperKey, value string, valueType viperconfig.ConfigType) error {
	switch valueType {
	case viperconfig.ENUM_BOOL:
		return setBool(viperKey, value)
	case viperconfig.ENUM_ID:
		return setUUID(viperKey, value)
	case viperconfig.ENUM_OUTPUT_FORMAT:
		return setOutputFormat(viperKey, value)
	case viperconfig.ENUM_PINGONE_REGION:
		return setPingOneRegion(viperKey, value)
	case viperconfig.ENUM_STRING:
		viper.Set(viperKey, string(value))
		return nil
	default:
		return fmt.Errorf("unable to set configuration: variable type for key '%s' is not recognized", viperKey)
	}
}

func setBool(viperKey string, value string) error {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("failed to set configuration: value for key '%s' must be a boolean. Use 'true' or 'false'", viperKey)
	}

	viper.Set(viperKey, boolValue)

	return nil
}

func setUUID(viperKey string, value string) error {
	// Check string is in the form of a UUID
	if _, err := uuid.ParseUUID(value); err != nil {
		return fmt.Errorf("failed to set configuration: value for key '%s' must be a valid UUID", viperKey)
	}

	viper.Set(viperKey, string(value))

	return nil
}

func setOutputFormat(viperKey string, value string) error {
	outputFormat := customtypes.OutputFormat("")
	if err := outputFormat.Set(value); err != nil {
		return fmt.Errorf("failed to set configuration: %s", err.Error())
	}

	viper.Set(viperKey, outputFormat)

	return nil
}

func setPingOneRegion(viperKey string, value string) error {
	region := customtypes.PingOneRegion("")
	if err := region.Set(value); err != nil {
		return fmt.Errorf("failed to set configuration: %s", err.Error())
	}

	viper.Set(viperKey, region)

	return nil
}
