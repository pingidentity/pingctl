package viperconfig

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func BindPersistentFlags(cobraParamNames []ConfigCobraParam, command *cobra.Command) error {
	for _, cobraParamName := range cobraParamNames {
		viperConfigKey := ConfigOptions[cobraParamName].ViperConfigKey
		err := viper.BindPFlag(viperConfigKey, command.PersistentFlags().Lookup(string(cobraParamName)))
		if err != nil {
			return err
		}
	}

	return nil
}

func BindFlags(cobraParamNames []ConfigCobraParam, command *cobra.Command) error {
	for _, cobraParamName := range cobraParamNames {
		viperConfigKey := ConfigOptions[cobraParamName].ViperConfigKey
		err := viper.BindPFlag(viperConfigKey, command.Flags().Lookup(string(cobraParamName)))
		if err != nil {
			return err
		}
	}

	return nil
}

func BindEnvVars(cobraParamNames []ConfigCobraParam) error {
	for _, cobraParamName := range cobraParamNames {
		viperConfigKey := ConfigOptions[cobraParamName].ViperConfigKey
		envVar := ConfigOptions[cobraParamName].EnvVar
		err := viper.BindEnv(viperConfigKey, envVar)
		if err != nil {
			return err
		}
	}

	return nil
}
