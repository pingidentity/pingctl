package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pingidentity/pingctl/cmd/auth"
	"github.com/pingidentity/pingctl/cmd/platform"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
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
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pingctl",
		Version: "v0.0.1",
		//TODO add command short and long description
		Short: "",
		Long:  ``,
	}

	cmd.AddCommand(
		platform.NewPlatformCommand(),
		NewFeedbackCommand(),
		auth.NewAuthCommand(),
	)

	cmd.PersistentFlags().StringVar(&cfgFile, configParamName, "", "Configuration file location\nDefault: $HOME/.pingctl/config.yaml")
	cmd.PersistentFlags().StringVar(&outputFormat, outputParamName, "text", "Specifies output format\nValid output options: 'text', 'json'")
	cmd.PersistentFlags().BoolVar(&colorizeOutput, colorParamName, true, "Use colorized output")

	if err := bindPersistentFlags(rootConfigurationParamMapping, cmd); err != nil {
		output.Format(cmd, output.CommandOutput{
			Message: "Error binding flag parameters. Flag values may not be recognized.",
			Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
			Error:   err,
		})
	}

	return cmd
}

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Root command...")

	cobra.OnInitialize(initConfig)
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

	//Use viper env string replacer for dashes and dots in var name
	replacer := strings.NewReplacer("-", "_", ".", "_")
	viper.SetEnvKeyReplacer(replacer)

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
