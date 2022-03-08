package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/megabot/internal/config"
)

// Database adds all flags for running the database command.
func Database(cmd *cobra.Command, values config.Values) {
}

// DatabaseMigrate adds all flags for running the database migrate command.
func DatabaseMigrate(cmd *cobra.Command, values config.Values) {
	Database(cmd, values)
	cmd.PersistentFlags().Bool(config.Keys.DbLoadTestData, values.DbLoadTestData, usage.DbLoadTestData)
}
