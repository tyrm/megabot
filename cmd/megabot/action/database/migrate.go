package database

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/cmd/megabot/action"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db/bun"
)

// Migrate runs database migrations
var Migrate action.Action = func(ctx context.Context) error {
	l := logrus.WithField("func", "Migrate")

	l.Info("running database migration")
	dbClient, err := bun.New(ctx)
	if err != nil {
		l.Errorf("db: %s", err.Error())
		return err
	}
	defer func() {
		err := dbClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	err = dbClient.DoMigration(ctx)
	if err != nil {
		l.Errorf("migration: %s", err.Error())
		return err
	}

	if viper.GetBool(config.Keys.DbLoadTestData) {
		err = dbClient.LoadTestData(ctx)
		if err != nil {
			l.Errorf("migration: %s", err.Error())
			return err
		}
	}

	return nil
}
