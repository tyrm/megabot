package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/megabot/internal/config"
)

// Global adds flags that are common to all commands.
func Global(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.ConfigPath, values.ConfigPath, usage.ConfigPath)
	cmd.PersistentFlags().String(config.Keys.LogLevel, values.LogLevel, usage.LogLevel)

	// application
	cmd.PersistentFlags().String(config.Keys.ApplicationName, values.ApplicationName, usage.ApplicationName)
	cmd.PersistentFlags().String(config.Keys.TokenSalt, values.TokenSalt, usage.TokenSalt)

	// database
	cmd.PersistentFlags().String(config.Keys.DbType, values.DbType, usage.DbType)
	cmd.PersistentFlags().String(config.Keys.DbAddress, values.DbAddress, usage.DbAddress)
	cmd.PersistentFlags().Int(config.Keys.DbPort, values.DbPort, usage.DbPort)
	cmd.PersistentFlags().String(config.Keys.DbUser, values.DbUser, usage.DbUser)
	cmd.PersistentFlags().String(config.Keys.DbPassword, values.DbPassword, usage.DbPassword)
	cmd.PersistentFlags().String(config.Keys.DbDatabase, values.DbDatabase, usage.DbDatabase)
	cmd.PersistentFlags().String(config.Keys.DbTLSMode, values.DbTLSMode, usage.DbTLSMode)
	cmd.PersistentFlags().String(config.Keys.DbTLSCACert, values.DbTLSCACert, usage.DbTLSCACert)
	cmd.PersistentFlags().String(config.Keys.DbEncryptionKey, values.DbEncryptionKey, usage.DbEncryptionKey)
}
