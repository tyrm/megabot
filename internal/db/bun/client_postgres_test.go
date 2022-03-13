//go:build postgres

package bun

import (
	"context"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"testing"
)

func TestPgConn(t *testing.T) {
	dbAddress := "postgres"
	dbDatabase := "test"
	dbPassword := "test"
	dbPort := 5432
	dbTLSMode := dbTLSModeDisable
	dbUser := "test"

	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbAddress, dbAddress)
	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbTLSMode, dbTLSMode)
	viper.Set(config.Keys.DbUser, dbUser)

	bun, err := pgConn(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing pg connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}
