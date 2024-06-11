package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pingidentity/pingctl/cmd/common"
	"github.com/pingidentity/pingctl/cmd/config"
	"github.com/pingidentity/pingctl/cmd/platform"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configParamName      = "config"
	configParamConfigKey = "pingctl.config"
	configParamEnvVar    = "PINGCTL_CONFIG"

	outputParamName      = "output"
	outputParamConfigKey = "pingctl.output"
	outputParamEnvVar    = "PINGCTL_OUTPUT"

	colorParamName      = "color"
	colorParamConfigKey = "pingctl.color"
	colorParamEnvVar    = "PINGCTL_COLOR"
)

var (
	cfgFile        string
	defaultCfgFile string
	outputFormat   string
	colorizeOutput bool

	cobraParamToViperConfigKeyMapping = map[string]string{
		configParamName: configParamConfigKey,
		outputParamName: outputParamConfigKey,
		colorParamName:  colorParamConfigKey,
	}

	viperConfigKeyToEnvVarMapping = map[string]string{
		configParamConfigKey: configParamEnvVar,
		outputParamConfigKey: outputParamEnvVar,
		colorParamConfigKey:  colorParamEnvVar,
	}
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Root command...")

	// Determine the default configuration file location
	home, err := os.UserHomeDir()
	if err != nil {
		l.Fatal().Err(err).Msgf("Failed to determine user's home directory")
	}

	// Default the config in $home/.pingctl directory with name "config.yaml".
	defaultCfgFile = fmt.Sprintf("%s/.pingctl/config.yaml", home)

	// Set config defaults
	viper.SetDefault(configParamConfigKey, defaultCfgFile)
	viper.SetDefault(outputParamConfigKey, "text")
	viper.SetDefault(colorParamConfigKey, true)

	cobra.OnInitialize(initViperConfigFile)
}

// rootCmd represents the base command when called without any subcommands
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "pingctl",
		Version:       "v2.0.0-alpha.3",
		Short:         "A CLI tool for managing Ping Identity products.",
		Long:          `A CLI tool for managing Ping Identity products.`,
		SilenceErrors: true, // Upon error in RunE method, let output package in main.go handle error output
	}

	cmd.AddCommand(
		platform.NewPlatformCommand(),
		NewFeedbackCommand(),
		config.NewConfigCommand(),
		// auth.NewAuthCommand(),
	)

	cmd.PersistentFlags().StringVar(&cfgFile, configParamName, "", "Configuration file location\nDefault: $HOME/.pingctl/config.yaml")
	cmd.PersistentFlags().StringVar(&outputFormat, outputParamName, "text", "Specifies output format\nValid output options: 'text', 'json'")
	cmd.PersistentFlags().BoolVar(&colorizeOutput, colorParamName, true, "Use colorized output")

	if err := common.BindPersistentFlags(cobraParamToViperConfigKeyMapping, cmd); err != nil {
		output.Format(cmd, output.CommandOutput{
			Message: "Error binding flag parameters. Flag values may not be recognized.",
			Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
			Error:   err,
		})
	}

	if err := common.BindEnvVars(viperConfigKeyToEnvVarMapping); err != nil {
		output.Format(cmd, output.CommandOutput{
			Message: "Error binding environment varibales. Environment Variable values may not be recognized.",
			Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
			Error:   err,
		})
	}

	return cmd
}

// initViperConfigFile reads in config file and ENV variables if set.
func initViperConfigFile() {
	l := logger.Get()

	if cfgFile == "" {
		l.Debug().Msgf("No configuration file specified. Determining default configuration file location: $HOME/.pingctl/config.yaml")
		cfgFile = defaultCfgFile
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

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		l.Fatal().Err(err).Msgf("Failed to read configuration from file: %s", cfgFile)
	} else {
		l.Info().Msgf("Using configuration file: %s", viper.ConfigFileUsed())
	}
}
