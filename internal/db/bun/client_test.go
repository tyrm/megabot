package bun

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"testing"
)

func TestSqliteConn(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbAddress, ":memory:")

	bun, err := sqliteConn(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing sqlite connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}

func TestSqliteConn_NoConfig(t *testing.T) {
	viper.Reset()

	_, err := sqliteConn(context.Background())
	errText := fmt.Sprintf("'%s' was not set when attempting to start sqlite", config.Keys.DbAddress)
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

func TestSqliteConn_BadPath(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbAddress, "invalidir/db.sqlite")

	_, err := sqliteConn(context.Background())
	errText := "sqlite ping: Unable to open the database file (SQLITE_CANTOPEN)"
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}
