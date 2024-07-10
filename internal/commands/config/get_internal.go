package config_internal

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
	"gopkg.in/yaml.v3"
)

func RunInternalConfigGet(args []string) error {
	// Write the profile configuration to file,
	// even though no configuration change is happening here
	// This handles the edge case where the config.yaml file was generated for
	// the first time, but no configuration changes were made and parameters/env vars were used
	if err := profiles.SaveProfileViperToFile(); err != nil {
		return err
	}

	viperKey, err := parseGetArgs(args)
	if err != nil {
		return err
	}

	// If the viper key is empty,
	// the parseGetArgs() function already printed the entire configuration
	if viperKey == "" {
		return nil
	}

	// The only valid configuration keys are those defined in profiles/types.go,
	// and their parent keys
	validKeys := profiles.ExpandedProfileKeys()
	if !slices.ContainsFunc(validKeys, func(v string) bool {
		return strings.EqualFold(v, viperKey)
	}) {
		validKeyStr := strings.Join(validKeys, ", ")
		return fmt.Errorf("unable to get configuration: value '%s' is not recognized as a valid configuration key. Valid keys: %s", viperKey, validKeyStr)
	}

	// Check if the viper configuration key is set
	if !profiles.GetProfileViper().IsSet(viperKey) {
		output.Print(output.Opts{
			Result:  output.ENUM_RESULT_NOACTION_WARN,
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
		output.Print(output.Opts{
			Message: fmt.Sprintf("'pingctl config get' only gets one key per command. Ignoring extra arguments: %s", strings.Join(args[1:], " ")),
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})
	}

	// Assume viper configuration key is args[0] and ignore any other input
	return args[0], nil
}

func PrintConfig() error {
	// Print the updated configuration
	yaml, err := yaml.Marshal(profiles.GetProfileViper().AllSettings())
	if err != nil {
		return fmt.Errorf("failed to yaml marshal pingctl configuration: %s", err.Error())
	}
	output.Print(output.Opts{
		Result:  output.ENUM_RESULT_NIL,
		Message: string(yaml),
	})

	return nil
}

func printConfigFromKey(viperKey string) error {
	// Print the updated configuration
	yaml, err := yaml.Marshal(profiles.GetProfileViper().Get(viperKey))
	if err != nil {
		return fmt.Errorf("failed to yaml marshal viper configuration: %s", err.Error())
	}
	output.Print(output.Opts{
		Result:  output.ENUM_RESULT_NIL,
		Message: string(yaml),
	})

	return nil
}
