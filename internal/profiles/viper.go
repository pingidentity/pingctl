package profiles

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Binding struct {
	Option Option
	Flag   *pflag.Flag
}

var (
	mainViper      *viper.Viper = viper.New()
	profileViper   *viper.Viper
	profileName    string
	flagBindings   []Binding
	pFlagBindings  []Binding
	envVarBindings []Option
)

func GetMainViper() *viper.Viper {
	return mainViper
}

func SetProfileViperWithProfile(pName string) (err error) {
	subViper := mainViper.Sub(pName)
	if subViper == nil {
		return fmt.Errorf("profile '%s' not found in configuration file: %s", pName, mainViper.ConfigFileUsed())
	}

	profileViper = subViper
	profileName = pName
	return nil
}

func SetProfileViperWithViper(v *viper.Viper) {
	profileViper = v
}

func GetProfileViper() *viper.Viper {
	return profileViper
}

func AddFlagBinding(binding Binding) {
	flagBindings = append(flagBindings, binding)
}

func AddPFlagBinding(binding Binding) {
	pFlagBindings = append(pFlagBindings, binding)
}

func AddEnvVarBinding(opt Option) {
	envVarBindings = append(envVarBindings, opt)
}

func ApplyBindingsToProfileViper() (err error) {
	for _, binding := range flagBindings {
		err = profileViper.BindPFlag(binding.Option.ViperKey, binding.Flag)
		if err != nil {
			return err
		}
	}

	for _, binding := range pFlagBindings {
		err = profileViper.BindPFlag(binding.Option.ViperKey, binding.Flag)
		if err != nil {
			return err
		}
	}

	for _, opt := range envVarBindings {
		err = profileViper.BindEnv(opt.ViperKey, opt.EnvVar)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveProfileViperToFile() error {
	profileKeys := profileViper.AllKeys()
	for _, key := range profileKeys {
		mainViper.Set(fmt.Sprintf("%s.%s", profileName, key), profileViper.Get(key))
	}

	if err := mainViper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write pingctl configuration to file '%s': %v", mainViper.ConfigFileUsed(), err)
	}

	return nil
}
