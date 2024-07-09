package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pingidentity/pingctl/cmd/config"
	"github.com/pingidentity/pingctl/cmd/feedback"
	"github.com/pingidentity/pingctl/cmd/platform"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	defaultCfgFile string
	outputFormat   customtypes.OutputFormat = customtypes.OutputFormat(customtypes.ENUM_OUTPUT_FORMAT_TEXT)
	colorizeOutput bool

	cobraParamNames = []viperconfig.ConfigCobraParam{
		viperconfig.RootOutputParamName,
		viperconfig.RootColorParamName,
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

	cobra.OnInitialize(initViperConfigFile)
}

// rootCmd represents the base command when called without any subcommands
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "pingctl",
		Version:           "v2.0.0-alpha.3",
		Short:             "A CLI tool for managing Ping Identity products.",
		Long:              `A CLI tool for managing Ping Identity products.`,
		SilenceErrors:     true, // Upon error in RunE method, let output package in main.go handle error output
		PersistentPreRunE: RootPersistentPreRunE,
	}

	cmd.AddCommand(
		platform.NewPlatformCommand(),
		feedback.NewFeedbackCommand(),
		config.NewConfigCommand(),
		// auth.NewAuthCommand(),
	)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Configuration file location\nDefault: $HOME/.pingctl/config.yaml")
	cmd.PersistentFlags().Var(&outputFormat, string(viperconfig.RootOutputParamName), fmt.Sprintf("Specifies output format\nValid output options: %s", strings.Join(customtypes.OutputFormatValidValues(), ", ")))
	cmd.PersistentFlags().BoolVar(&colorizeOutput, string(viperconfig.RootColorParamName), true, "Use colorized output")

	if err := viperconfig.BindPersistentFlags(cobraParamNames, cmd); err != nil {
		output.Print(output.Opts{
			Message:      "Error binding flag parameters. Flag values may not be recognized.",
			Result:       output.ENUM_RESULT_FAILURE,
			ErrorMessage: err.Error(),
		})
	}

	if err := viperconfig.BindEnvVars(cobraParamNames); err != nil {
		output.Print(output.Opts{
			Message:      "Error binding environment variables. Environment Variable values may not be recognized.",
			Result:       output.ENUM_RESULT_FAILURE,
			ErrorMessage: err.Error(),
		})
	}

	// Make sure cobra is outputting to stdout and stderr
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

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

func RootPersistentPreRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()

	// Validate viper config
	if err := viperconfig.ValidateViperConfig(); err != nil {
		return err
	}

	l.Info().Msgf("Successfully validated pingctl configuration")
	return nil
}
