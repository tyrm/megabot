package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/megabot/internal/config"
)

// Server adds all flags for running the server.
func Server(cmd *cobra.Command, values config.Values) {
	// application
	cmd.PersistentFlags().StringArray(config.Keys.ServerRoles, values.ServerRoles, usage.ServerRoles)
}
