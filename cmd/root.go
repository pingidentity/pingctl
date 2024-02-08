/*
Copyright © 2024 Ping Identity Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pingidentity/pingctl/cmd/auth"
	"github.com/pingidentity/pingctl/cmd/platform"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configParamName      = "config"
	configParamConfigKey = "config"

	outputParamName      = "output"
	outputParamConfigKey = "output"

	colorParamName      = "color"
	colorParamConfigKey = "color"
)

var (
	cfgFile        string
	outputFormat   string
	colorizeOutput bool

	rootConfigurationParamMapping = map[string]string{
		configParamName: configParamConfigKey,
		outputParamName: outputParamConfigKey,
		colorParamName:  colorParamConfigKey,
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pingctl",
	Version: "v0.0.1",
	//TODO add command short and long description
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	l := logger.Get()
	err := rootCmd.Execute()
	if err != nil {
		l.Fatal().Err(err).Msgf("")
	}
}

// init adds all child commands to the root command and sets flags appropriately.
func init() {
	l := logger.Get()

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(
		platform.PlatformCmd,
		feedbackCmd,
		auth.NewAuthCommand(),
	)

	rootCmd.PersistentFlags().StringVar(&cfgFile, configParamName, "", "Configuration file location\nDefault: $HOME/.pingctl/config.yaml")
	rootCmd.PersistentFlags().StringVar(&outputFormat, outputParamName, "text", "Specifies output format\nValid output options: 'text', 'json'")
	rootCmd.PersistentFlags().BoolVar(&colorizeOutput, colorParamName, true, "Use colorized output")

	if err := bindPersistentFlags(rootConfigurationParamMapping, rootCmd); err != nil {
		l.Error().Err(err).Msgf("Error binding flag parameters. Flag values may not be recognized.")
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	l := logger.Get()
	if cfgFile == "" {
		l.Debug().Msgf("No configuration file specified. Determining default configuration file location: $HOME/.pingctl/config.yaml")
		home, err := os.UserHomeDir()
		if err != nil {
			l.Fatal().Err(err).Msgf("Failed to determine user's home directory")
		}

		// Default the config in $home/.pingctl directory with name "config.yaml".
		cfgFile = fmt.Sprintf("%s/.pingctl/config.yaml", home)
		l.Debug().Msgf("Determined default configuration file location: %s", cfgFile)

		// Make sure the default config file exists, and if not, seed a new file
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			l.Debug().Msgf("Default configuration file does not exists. Seeding a new file at location: %s", cfgFile)

			// MkdirAll does nothing if directories already exist. Create needed directories for config file location.
			err := os.MkdirAll(filepath.Dir(cfgFile), os.ModePerm)
			if err != nil {
				l.Fatal().Err(err).Msgf("Failed to make directories needed for filepath: %s", cfgFile)
			}

			// SafeWriteConfigAs writes current configuration to a given filename if it does not exist.
			err = viper.SafeWriteConfigAs(cfgFile)
			if err != nil {
				l.Fatal().Err(err).Msgf("Failed to create configuration file at: %s", cfgFile)
			}
		}
	}

	// Use config file from the flag.
	viper.SetConfigFile(cfgFile)

	//Only use environment variabes with the "PINGCTL" prefix
	viper.SetEnvPrefix("PINGCTL")

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		l.Fatal().Err(err).Msgf("Failed to read configuration from file: %s", cfgFile)
	} else {
		l.Info().Msgf("Using configuration file: %s", viper.ConfigFileUsed())
	}
}

func bindPersistentFlags(paramlist map[string]string, command *cobra.Command) error {
	for k, v := range paramlist {
		err := viper.BindPFlag(v, command.PersistentFlags().Lookup(k))
		if err != nil {
			return err
		}
	}

	return nil
}
