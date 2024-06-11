package config

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
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
	if len(args) == 0 {
		return fmt.Errorf("unable to unset configuration: no key given in unset command")
	}

	// Assume viper configuration key is args[0] and ignore any other input
	viperKey := args[0]

	// The only valid configuration keys are those that are already set in the
	// configuration file. This is ensured by viper.SetDefault calls on command's init.
	if !viper.InConfig(viperKey) {
		validKeys := strings.Join(viper.AllKeys(), ", ")
		return fmt.Errorf("unable to unset configuration: value '%s' is not recognized as a valid configuration key. \nValid keys: %s", viperKey, validKeys)
	}

	// Check is viper key maps to a map object
	// If it does, do not allow setting the key directly
	if len(viper.GetStringMap(viperKey)) > 0 {
		return fmt.Errorf("unable to unset configuration: key '%s' is an object and cannot be unset. Use 'pingctl config unset %s.<key>' to unset a specific configuration setting", viperKey, viperKey)
	}

	viper.Set(viperKey, "")

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
