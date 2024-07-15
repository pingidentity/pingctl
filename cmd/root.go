package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pingidentity/pingctl/cmd/config"
	"github.com/pingidentity/pingctl/cmd/feedback"
	"github.com/pingidentity/pingctl/cmd/platform"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultProfileName string = "default"
)

var (
	cfgFile        string
	defaultCfgFile string
	profileName    string

	// Custom pflag.Value types
	outputFormat customtypes.OutputFormat = customtypes.OutputFormat(customtypes.ENUM_OUTPUT_FORMAT_TEXT)
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

	cobra.OnInitialize(initViperAndProfile)
}

// rootCmd represents the base command when called without any subcommands
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "pingctl",
		Version:       "v2.0.0-alpha.4",
		Short:         "A CLI tool for managing Ping Identity products.",
		Long:          `A CLI tool for managing Ping Identity products.`,
		SilenceErrors: true, // Upon error in RunE method, let output package in main.go handle error output
	}

	cmd.AddCommand(
		platform.NewPlatformCommand(),
		feedback.NewFeedbackCommand(),
		config.NewConfigCommand(),
		// auth.NewAuthCommand(),
	)

	// flags used within this file assigned to variables
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Configuration file location (default \"$HOME/.pingctl/config.yaml\"")
	cmd.PersistentFlags().StringVar(&profileName, profiles.ProfileOption.CobraParamName, "", "Profile to use from configuration file (default \"default\")")

	// custom pflag.Value types use Var() method
	cmd.PersistentFlags().Var(&outputFormat, profiles.OutputOption.CobraParamName, fmt.Sprintf("Specifies output format\nValid output options: %s", strings.Join(customtypes.OutputFormatValidValues(), ", ")))
	profiles.AddPFlagBinding(profiles.Binding{
		Option: profiles.OutputOption,
		Flag:   cmd.PersistentFlags().Lookup(profiles.OutputOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.OutputOption)

	// flags where values are stored and accessed via viper
	cmd.PersistentFlags().Bool(profiles.ColorOption.CobraParamName, true, "Use colorized output")
	profiles.AddPFlagBinding(profiles.Binding{
		Option: profiles.ColorOption,
		Flag:   cmd.PersistentFlags().Lookup(profiles.ColorOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.ColorOption)

	// Make sure cobra is outputting to stdout and stderr
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	return cmd
}

func initViperAndProfile() {
	l := logger.Get()

	// If no configuration file location is specified, use the default configuration file location
	if cfgFile == "" {
		initDefaultConfigFile()
	}

	//Configure the main viper instance
	initMainViper()

	// Prefer parameter, then environment variable, then configuration file
	// like with viper hierarchy. Finally default to @defaultProfileName if not found
	// NOTE: this is needed because parameter and env var are not bound to
	// the main viper instance
	if profileName == "" {
		profileName = os.Getenv(profiles.ProfileOption.EnvVar)
	}
	if profileName == "" {
		mainViper := profiles.GetMainViper()
		profileName = mainViper.GetString(profiles.ProfileOption.ViperKey)
	}
	if profileName == "" {
		profileName = defaultProfileName
	}

	l.Debug().Msgf("Using configuration profile: %s", profileName)

	// Configure the profile viper instance
	if err := profiles.SetProfileViperWithProfile(profileName); err != nil {
		output.Print(output.Opts{
			Message:      "Failed to set profile viper",
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}

	// All bindings have been set by NewXCommand() functions previous to OnInitialize()
	if err := profiles.ApplyBindingsToProfileViper(); err != nil {
		output.Print(output.Opts{
			Message:      "Failed to apply bindings to profile viper",
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}
}

func initMainViper() {
	l := logger.Get()

	mainViper := profiles.GetMainViper()
	// Use config file from the flag.
	mainViper.SetConfigFile(cfgFile)

	// If a config file is found, read it in.
	if err := mainViper.ReadInConfig(); err != nil {
		output.Print(output.Opts{
			Message:      fmt.Sprintf("Failed to read configuration from file: %s", cfgFile),
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	} else {
		l.Info().Msgf("Using configuration file: %s", mainViper.ConfigFileUsed())
	}

	// Validate the configuration
	if err := profiles.Validate(); err != nil {
		output.Print(output.Opts{
			Message:      "Failed to validate pingctl configuration",
			Result:       output.ENUM_RESULT_FAILURE,
			FatalMessage: err.Error(),
		})
	}
}

func initDefaultConfigFile() {
	l := logger.Get()

	l.Debug().Msgf("No configuration file location specified. Using default configuration file location: %s", defaultCfgFile)
	cfgFile = defaultCfgFile

	// Make sure the default config file exists, and if not, seed a new file
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		l.Debug().Msgf("Default configuration file does not exist. Seeding a new file at location: %s", cfgFile)

		// MkdirAll does nothing if directories already exist. Create needed directories for config file location.
		err := os.MkdirAll(filepath.Dir(cfgFile), os.ModePerm)
		if err != nil {
			output.Print(output.Opts{
				Message:      fmt.Sprintf("Failed to make directories needed for filepath: %s", cfgFile),
				Result:       output.ENUM_RESULT_FAILURE,
				FatalMessage: err.Error(),
			})
		}

		// No viper instance is configured yet, so to create a valid configuration file,
		// we need to create a new viper instance and set the configuration options to their default values.
		tempViper := viper.New()
		profiles.SetProfileViperWithViper(tempViper)
		for _, opt := range profiles.ConfigOptions.Options {
			if opt.ViperKey == profiles.ProfileOption.ViperKey {
				tempViper.Set(opt.ViperKey, defaultProfileName)
				continue
			}
			if err := config_internal.UnsetValue(fmt.Sprintf("%s.%s", defaultProfileName, opt.ViperKey), opt.Type); err != nil {
				output.Print(output.Opts{
					Message:      "Failed to set default configuration value",
					Result:       output.ENUM_RESULT_FAILURE,
					FatalMessage: err.Error(),
				})
			}
		}

		// Ensure this temporary viper instance cannot be used elsewhere
		profiles.SetProfileViperWithViper(nil)

		// SafeWriteConfigAs writes current configuration to a given filename if it does not exist.
		// Use global viper instance as main viper instance is not yet configured.
		err = tempViper.SafeWriteConfigAs(cfgFile)
		if err != nil {
			output.Print(output.Opts{
				Message:      fmt.Sprintf("Failed to create configuration file at: %s", cfgFile),
				Result:       output.ENUM_RESULT_FAILURE,
				FatalMessage: err.Error(),
			})
		}
	}
}
