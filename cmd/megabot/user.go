package main

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/megabot/cmd/megabot/action/user"
	"github.com/tyrm/megabot/cmd/megabot/flag"
	"github.com/tyrm/megabot/internal/config"
)

// serverCommands returns the 'server' subcommand
func userCommands() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "manage users",
	}
	flag.User(userCmd, config.Defaults)

	userAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add a user",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), user.Add)
		},
	}
	flag.UserAdd(userAddCmd, config.Defaults)
	userCmd.AddCommand(userAddCmd)

	return userCmd
}
