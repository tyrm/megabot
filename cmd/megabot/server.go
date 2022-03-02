package main

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/megabot/cmd/megabot/action/server"
	"github.com/tyrm/megabot/cmd/megabot/flag"
	"github.com/tyrm/megabot/internal/config"
)

// serverCommands returns the 'server' subcommand
func serverCommands() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "run a megabot server",
	}

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start the megabot server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), server.Start)
		},
	}

	flag.Server(serverStartCmd, config.Defaults)

	serverCmd.AddCommand(serverStartCmd)

	return serverCmd
}
