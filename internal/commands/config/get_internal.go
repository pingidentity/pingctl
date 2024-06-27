package config_internal

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func RunInternalConfigGet(args []string) error {
	viperKey, err := parseGetArgs(args)
	if err != nil {
		return err
	}

	// If the viper key is empty,
	// the parseGetArgs() function already printed the entire configuration
	if viperKey == "" {
		return nil
	}

	// The only valid configuration keys are those defined in viperconfig,
	// and their parent keys
	validKeys := getValidGetKeys()
	if !slices.Contains(validKeys, viperKey) {
		validKeyStr := strings.Join(validKeys, ", ")
		return fmt.Errorf("unable to get configuration: value '%s' is not recognized as a valid configuration key. Valid keys: %s", viperKey, validKeyStr)
	}

	// Check if the viper configuration key is set
	if !viper.IsSet(viperKey) {
		output.Format(output.CommandOutput{
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
			Message: fmt.Sprintf("Configuration key '%s' is not set", viperKey),
		})
		return nil
	}

	if err := printConfigFromKey(viperKey); err != nil {
		return err
	}

	return nil
}

func parseGetArgs(args []string) (string, error) {
	// If no configuration key is supplied via args, return all configuration settings as YAML
	if len(args) == 0 {
		if err := PrintConfig(); err != nil {
			return "", err
		}
		return "", nil
	}

	if len(args) > 1 {
		output.Format(output.CommandOutput{
			Message: fmt.Sprintf("'pingctl config get' only gets one key per command. Ignoring extra arguments: %s", strings.Join(args[1:], " ")),
			Result:  output.ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN,
		})
	}

	// Assume viper configuration key is args[0] and ignore any other input
	return args[0], nil
}

func PrintConfig() error {
	// Print the updated configuration
	yaml, err := yaml.Marshal(viper.AllSettings())
	if err != nil {
		return fmt.Errorf("failed to yaml marshal viper configuration: %s", err.Error())
	}
	output.Format(output.CommandOutput{
		Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
		Message: string(yaml),
	})

	return nil
}

func printConfigFromKey(viperKey string) error {
	// Print the updated configuration
	yaml, err := yaml.Marshal(viper.Get(viperKey))
	if err != nil {
		return fmt.Errorf("failed to yaml marshal viper configuration: %s", err.Error())
	}
	output.Format(output.CommandOutput{
		Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
		Message: string(yaml),
	})

	return nil
}

func getValidGetKeys() []string {
	// for each leaf key, add parent keys by splitting on the "." character
	leafKeys := viperconfig.GetViperConfigKeys()
	allKeys := []string{}
	for _, key := range leafKeys {
		keySplit := strings.Split(key, ".")
		for i := 0; i < len(keySplit); i++ {
			curKey := strings.Join(keySplit[:i+1], ".")
			if !slices.Contains(allKeys, curKey) {
				allKeys = append(allKeys, curKey)
			}
		}
	}

	slices.Sort(allKeys)

	return allKeys
}
