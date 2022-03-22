package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/megabot/internal/config"
)

// Server adds all flags for running the server.
func Server(cmd *cobra.Command, values config.Values) {
	Redis(cmd, values)

	// server
	cmd.PersistentFlags().String(config.Keys.ServerExternalHostname, values.ServerExternalHostname, usage.ServerExternalHostname)
	cmd.PersistentFlags().Bool(config.Keys.ServerHTTP2, values.ServerHTTP2, usage.ServerHTTP2)
	cmd.PersistentFlags().String(config.Keys.ServerHTTP2Bind, values.ServerHTTP2Bind, usage.ServerHTTP2Bind)
	cmd.PersistentFlags().Bool(config.Keys.ServerHTTP3, values.ServerHTTP3, usage.ServerHTTP3)
	cmd.PersistentFlags().String(config.Keys.ServerHTTP3Bind, values.ServerHTTP3Bind, usage.ServerHTTP3Bind)
	cmd.PersistentFlags().Bool(config.Keys.ServerMinifyHTML, values.ServerMinifyHTML, usage.ServerMinifyHTML)
	cmd.PersistentFlags().StringArray(config.Keys.ServerRoles, values.ServerRoles, usage.ServerRoles)

	// auth
	cmd.PersistentFlags().Duration(config.Keys.AccessExpiration, values.AccessExpiration, usage.AccessExpiration)
	cmd.PersistentFlags().String(config.Keys.AccessSecret, values.AccessSecret, usage.AccessSecret)
	cmd.PersistentFlags().Duration(config.Keys.RefreshExpiration, values.RefreshExpiration, usage.RefreshExpiration)
	cmd.PersistentFlags().String(config.Keys.RefreshSecret, values.RefreshSecret, usage.RefreshSecret)
}
