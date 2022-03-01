package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/cmd/megabot/action"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/log"
)

// Version is the software version.
var Version string

// Commit is the git commit.
var Commit string

func main() {
	var v string
	if len(Commit) < 7 {
		v = Version
	} else {
		v = Version + " " + Commit[:7]
	}

	// set software version
	viper.Set(config.Keys.SoftwareVersion, v)

	rootCmd := &cobra.Command{
		Use:   "megabot",
		Short: "MegaBot - a robot",
		//TODO Long:          "",
		Version:       v,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	// add commands
	rootCmd.AddCommand(serverCommands())

	err := rootCmd.Execute()
	if err != nil {
		logrus.Fatalf("error executing command: %s", err)
	}
}

func preRun(cmd *cobra.Command) error {
	if err := config.Init(cmd.Flags()); err != nil {
		return fmt.Errorf("error initializing config: %s", err)
	}

	if err := config.ReadConfigFile(); err != nil {
		return fmt.Errorf("error reading config: %s", err)
	}

	return nil
}

func run(ctx context.Context, action action.Action) error {
	if err := log.Init(); err != nil {
		return fmt.Errorf("error initializing log: %s", err)
	}

	return action(ctx)
}
