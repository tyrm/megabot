//go:build mysql

package bun

import (
	"context"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db"
	"testing"
)

func TestNew_Mysql(t *testing.T) {
	dbAddress := "mariadb"
	dbDatabase := "test"
	dbPassword := "test"
	dbPort := 3306
	dbTLSMode := dbTLSModeDisable
	dbUser := "test"

	viper.Reset()

	viper.Set(config.Keys.DbType, "mysql")

	viper.Set(config.Keys.DbAddress, dbAddress)
	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbTLSMode, dbTLSMode)
	viper.Set(config.Keys.DbUser, dbUser)

	bun, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing bun connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}

func TestMyConn(t *testing.T) {
	dbAddress := "mariadb"
	dbDatabase := "test"
	dbPassword := "test"
	dbPort := 3306
	dbTLSMode := dbTLSModeDisable
	dbUser := "test"

	viper.Reset()

	viper.Set(config.Keys.DbType, "mysql")

	viper.Set(config.Keys.DbAddress, dbAddress)
	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbTLSMode, dbTLSMode)
	viper.Set(config.Keys.DbUser, dbUser)

	bun, err := myConn(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing pg connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}

func testNewMysqlClient() (db.DB, error) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "mysql")

	viper.Set(config.Keys.DbAddress, "mariadb")
	viper.Set(config.Keys.DbDatabase, "test")
	viper.Set(config.Keys.DbPassword, "test")
	viper.Set(config.Keys.DbPort, 3306)
	viper.Set(config.Keys.DbUser, "test")

	client, err := New(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.DoMigration(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.LoadTestData(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
