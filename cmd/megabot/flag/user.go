package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/megabot/internal/config"
)

// User adds all flags for running the user command.
func User(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.UserEmail, values.UserEmail, usage.UserEmail)
}

// UserAdd adds all flags for running the user add command.
func UserAdd(cmd *cobra.Command, values config.Values) {
	User(cmd, values)
	cmd.PersistentFlags().StringArray(config.Keys.UserGroups, values.UserGroups, usage.UserGroups)
	cmd.PersistentFlags().String(config.Keys.UserPassword, values.UserPassword, usage.UserPassword)
}