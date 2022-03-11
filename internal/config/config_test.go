package config

import (
	"github.com/spf13/pflag"
	"testing"
)

func TestInit(t *testing.T) {
	flags := &pflag.FlagSet{}
	err := Init(flags)

	if err != nil {
		t.Errorf("unexpected error initializing config: %s", err.Error())
	}
}

func TestReadConfigFile(t *testing.T) {
	flags := &pflag.FlagSet{}
	flags.Set("config-path", "test/test-config.yml")
	Init(flags)

	err := ReadConfigFile()
	if err != nil {
		t.Errorf("unexpected error reading config: %s", err.Error())
	}
}
