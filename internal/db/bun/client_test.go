package bun

import (
	"context"
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
