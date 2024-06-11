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
	// If no configuration key is supplied via args, return all configuration settings as YAML
	if len(args) == 0 {
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

	// Assume viper configuration key is args[0] and ignore any other input
	viperKey := args[0]

	// The only valid configuration keys are those that are already set in the
	// configuration file. This is ensured by viper.SetDefault calls on command's init.
	if !viper.InConfig(viperKey) {
		validKeys := strings.Join(viper.AllKeys(), ", ")
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
