//go:build mysql

package bun

import (
	"context"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"testing"
)

func TestCommonDB_Create_Mysql(t *testing.T) {
	client, err := testNewMysqlClient()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	testCommonDBCreate(t, client)
}

func TestCommonDB_DoMigration_Mysql(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "mysql")

	viper.Set(config.Keys.DbAddress, "mariadb")
	viper.Set(config.Keys.DbDatabase, "test")
	viper.Set(config.Keys.DbPassword, "test")
	viper.Set(config.Keys.DbPort, 3306)
	viper.Set(config.Keys.DbUser, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing bun connection: %s", err.Error())
		return
	}

	err = client.DoMigration(context.Background())
	if err != nil {
		t.Errorf("unexpected error running migration: %s", err.Error())
		return
	}

	err = client.DoMigration(context.Background())
	if err != nil {
		t.Errorf("unexpected error running migration twice: %s", err.Error())
		return
	}
}
