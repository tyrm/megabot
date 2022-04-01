package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/cmd/megabot/action"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db/bun"
	"github.com/tyrm/megabot/internal/models"
	"strings"
)

// Show displays info about a user
var Show action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Show")

	l.Infof("reading user %s", viper.GetString(config.Keys.UserEmail))
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

	user, err := dbClient.ReadUserByEmail(ctx, viper.GetString(config.Keys.UserEmail))
	if err != nil {
		l.Errorf("readng : %s", err.Error())
		return err
	}

	l.Infof("ID: %d", user.ID)
	l.Infof("Email: %s", user.Email)

	groupSlice := make([]string, len(user.Groups))
	for i, g := range user.Groups {
		var gid uuid.UUID
		var err error
		if gid, err = g.GetGroupID(); err != nil {
			return err
		}
		groupSlice[i] = models.GroupTitle(gid)
	}
	l.Infof("Groups: %s", strings.Join(groupSlice, ", "))

	return nil
}
