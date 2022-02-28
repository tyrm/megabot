package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Init starts config collection
func Init(flags *pflag.FlagSet) error {
	viper.SetEnvPrefix("mb")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	err := viper.BindPFlags(flags)
	if err != nil {
		return err
	}
	return nil
}