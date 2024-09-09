package configuration

import (
	"fmt"
	"os"
	"strings"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/pflag"
)

// Options
var (
	RootActiveProfileOption Option
	RootColorOption         Option
	RootConfigOption        Option
	RootOutputFormatOption  Option
)

func initRootOptions() {
	initActiveProfileOption()
	initColorOption()
	initConfigOption()
	initOutputFormatOption()
}

func initActiveProfileOption() {
	cobraParamName := "active-profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("default")

	RootActiveProfileOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCTL_ACTIVE_PROFILE",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "P",
			Usage:     "Profile to use from configuration file",
			Value:     cobraValue,
			DefValue:  "default",
		},
		Type:     ENUM_STRING,
		ViperKey: "activeProfile",
	}
}

func initColorOption() {
	cobraParamName := "color"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(true)

	RootColorOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCTL_COLOR",
		Flag: &pflag.Flag{
			Name:     cobraParamName,
			Usage:    "Use colorized output",
			Value:    cobraValue,
			DefValue: "true",
		},
		Type:     ENUM_BOOL,
		ViperKey: "pingctl.color",
	}
}

func initConfigOption() {
	cobraParamName := "config"
	cobraValue := new(customtypes.String)
	defaultValue := getDefaultConfigFilepath()

	RootConfigOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    defaultValue,
		EnvVar:          "PINGCTL_CONFIG",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "C",
			Usage:     "Configuration file location",
			Value:     cobraValue,
			DefValue:  "\"$HOME/.pingctl/config.yaml\"",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initOutputFormatOption() {
	cobraParamName := "output-format"
	cobraValue := new(customtypes.OutputFormat)
	defaultValue := customtypes.OutputFormat(customtypes.ENUM_OUTPUT_FORMAT_TEXT)

	RootOutputFormatOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "PINGCTL_OUTPUT_FORMAT",
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "O",
			Usage:     fmt.Sprintf("Specifies pingctl's console output format. Allowed: %s", strings.Join(customtypes.OutputFormatValidValues(), ", ")),
			Value:     cobraValue,
			DefValue:  customtypes.ENUM_OUTPUT_FORMAT_TEXT,
		},
		Type:     ENUM_OUTPUT_FORMAT,
		ViperKey: "pingctl.outputFormat",
	}
}

func getDefaultConfigFilepath() (defaultConfigFilepath *customtypes.String) {
	l := logger.Get()

	defaultConfigFilepath = new(customtypes.String)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		l.Err(err).Msg("Failed to determine user's home directory")
		return nil
	}

	err = defaultConfigFilepath.Set(fmt.Sprintf("%s/.pingctl/config.yaml", homeDir))
	if err != nil {
		l.Err(err).Msg("Failed to set default config file path")
		return nil
	}

	return defaultConfigFilepath
}
