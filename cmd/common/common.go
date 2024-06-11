package common

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func BindPersistentFlags(paramlist map[string]string, command *cobra.Command) error {
	for cobraParam, viperConfigKey := range paramlist {
		err := viper.BindPFlag(viperConfigKey, command.PersistentFlags().Lookup(cobraParam))
		if err != nil {
			return err
		}
	}

	return nil
}

func BindFlags(paramlist map[string]string, command *cobra.Command) error {
	for k, v := range paramlist {
		err := viper.BindPFlag(v, command.Flags().Lookup(k))
		if err != nil {
			return err
		}
	}

	return nil
}

func BindEnvVars(paramlist map[string]string) error {
	for viperConfigKey, envVar := range paramlist {
		err := viper.BindEnv(viperConfigKey, envVar)
		if err != nil {
			return err
		}
	}

	return nil
}
