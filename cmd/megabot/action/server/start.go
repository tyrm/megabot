package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/megabot/cmd/megabot/action"
	"github.com/tyrm/megabot/internal/db/bun"
)

var Start action.Action = func(ctx context.Context) error {
	logrus.Infof("starting")
	dbClient, err := bun.NewClient(ctx)
	if err != nil {
		logrus.Errorf("db: %s", err.Error())
		return err
	}
	defer func() {
		err := dbClient.Close(ctx)
		if err != nil {
			logrus.Errorf("closing db: %s", err.Error())
		}
	}()

	_ = dbClient

	return nil
}
