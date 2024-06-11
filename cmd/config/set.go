package config

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
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
	if len(args) == 0 {
		return fmt.Errorf("unable to set configuration: no 'key=value' assignment given in set command")
	}

	// Assume viper configuration key=value pair is args[0] and ignore any other input
	parsedInput := strings.SplitN(args[0], "=", 2)
	if len(parsedInput) != 2 {
		return fmt.Errorf("unable to set configuration: invalid assignment format '%s'. Expect 'key=value' format", args[0])
	}

	viperKey := parsedInput[0]
	value := parsedInput[1]

	// The only valid configuration keys are those that are already set in the
	// configuration file. This is ensured by viper.SetDefault calls on command's init.
	if !viper.InConfig(viperKey) {
		validKeys := strings.Join(viper.AllKeys(), ", ")
		return fmt.Errorf("unable to set configuration: value '%s' is not recognized as a valid configuration key. \nValid keys: %s", viperKey, validKeys)
	}

	// Check if viperKey is in viper.AllKeys()
	if !slices.Contains(viper.AllKeys(), viperKey) {
		return fmt.Errorf("unable to set configuration: key '%s' is an object and cannot be set. Use 'pingctl config set %s.<key>=%s' to set a specific configuration setting", viperKey, viperKey, value)
	}

	// Make sure value is not empty, and suggest unset command if it is
	if value == "" {
		return fmt.Errorf("unable to set configuration: value for key '%s' is empty. Use 'pingctl config unset %s' to unset the key", viperKey, viperKey)
	}

	viper.Set(viperKey, value)

	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("failed to write configuration to file: %s", err.Error())
	}

	// Print the updated configuration
	yaml, err := yaml.Marshal(viper.AllSettings())
	if err != nil {
		return fmt.Errorf("failed to yaml marshal viper configuration: %s", err.Error())
	}
	output.Format(cmd, output.CommandOutput{
		Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
		Message: string(yaml),
	})

	return nil
}
