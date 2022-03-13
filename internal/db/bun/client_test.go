package bun

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"testing"
)

func TestDeriveBunDBPGOptions(t *testing.T) {
	dbDatabase := "database"
	dbPassword := "password"
	dbPort := 5432
	dbUser := "user"

	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbUser, dbUser)

	opts, err := deriveBunDBPGOptions()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}
	if opts == nil {
		t.Errorf("opts is nil")
		return
	}

	if opts.Database != dbDatabase {
		t.Errorf("unexpected value for database, got: '%s', want: '%s'", opts.Database, dbDatabase)
	}
	if opts.Password != dbPassword {
		t.Errorf("unexpected value for password, got: '%s', want: '%s'", opts.Password, dbPassword)
	}
	if opts.Port != uint16(dbPort) {
		t.Errorf("unexpected value for port, got: '%d', want: '%d'", opts.Port, dbPort)
	}
	if opts.User != dbUser {
		t.Errorf("unexpected value for user, got: '%s', want: '%s'", opts.User, dbUser)
	}
	// tls
	if opts.TLSConfig != nil {
		t.Errorf("unexpected value for tls config, got: '%v', want: '%v'", opts.User, nil)
	}
}

func TestDeriveBunDBPGOptions_NoDatabase(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	_, err := deriveBunDBPGOptions()
	errText := "no database set"
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

func TestDeriveBunDBPGOptions_NoType(t *testing.T) {
	viper.Reset()

	_, err := deriveBunDBPGOptions()
	errText := "expected bun type of POSTGRES but got "
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

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
