package config

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Config Get Subcommand...")
}

func NewConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get pingctl configuration settings.",
		Long: `Get pingctl configuration settings.
		
Example command usage: 'pingctl config get pingctl.color'`,
		RunE: ConfigGetRunE,
	}

	return cmd
}

func ConfigGetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	viperKey, err := parseGetArgs(args, cmd)
	if err != nil {
		return err
	}

	// If the viper key is empty,
	// the parseGetArgs() function already printed the entire configuration
	if viperKey == "" {
		return nil
	}

	// The only valid configuration keys are those that are already set in the
	// configuration file. If the key is not recognized, return an error.
	if !viper.InConfig(viperKey) {
		validKeys := strings.Join(viperconfig.GetViperConfigKeys(), ", ")
		return fmt.Errorf("unable to get configuration: value '%s' is not recognized as a valid configuration key. \nValid keys: %s", viperKey, validKeys)
	}

	// Check if the viper configuration key is set
	if !viper.IsSet(viperKey) {
		output.Format(cmd, output.CommandOutput{
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
			Message: fmt.Sprintf("Configuration key '%s' is not set", viperKey),
		})
		return nil
	}

	if err := printConfigFromKey(cmd, viperKey); err != nil {
		return err
	}

	return nil
}

func parseGetArgs(args []string, cmd *cobra.Command) (string, error) {
	// If no configuration key is supplied via args, return all configuration settings as YAML
	if len(args) == 0 {
		if err := printConfig(cmd); err != nil {
			return "", err
		}
		return "", nil
	}

	// Assume viper configuration key is args[0] and ignore any other input
	return args[0], nil
}

func printConfig(cmd *cobra.Command) error {
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

func printConfigFromKey(cmd *cobra.Command, viperKey string) error {
	// Print the updated configuration
	yaml, err := yaml.Marshal(viper.Get(viperKey))
	if err != nil {
		return fmt.Errorf("failed to yaml marshal viper configuration: %s", err.Error())
	}
	output.Format(cmd, output.CommandOutput{
		Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
		Message: string(yaml),
	})

	return nil
}
