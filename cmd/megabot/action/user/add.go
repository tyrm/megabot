package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/cmd/megabot/action"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db/bun"
	"github.com/tyrm/megabot/internal/models"
)

// Add adds a user from the command line
var Add action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Add")

	l.Infof("adding user %s", viper.GetString(config.Keys.UserEmail))
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

	// create user
	newUser := models.User{
		Email: viper.GetString(config.Keys.UserEmail),
	}
	err = newUser.SetPassword(viper.GetString(config.Keys.UserPassword))
	if err != nil {
		l.Errorf("setting password: %s", err.Error())
		return err
	}

	// add groups
	for _, group := range viper.GetStringSlice(config.Keys.UserGroups) {
		groupID := models.GroupName(group)
		if groupID == uuid.Nil {
			msg := fmt.Sprintf("unknown group: %s", group)
			l.Errorf(msg)
			return errors.New(msg)
		}

		groupMem := &models.GroupMembership{
			GroupID: groupID,
		}
		l.Debugf("adding group: %s", group)
		newUser.Groups = append(newUser.Groups, groupMem)
	}

	err = dbClient.Create(ctx, &newUser)
	if err != nil {
		l.Errorf("db: %s", err.Error())
		return err
	}

	for _, g := range newUser.Groups {
		g.UserID = newUser.ID
	}

	if len(newUser.Groups) > 0 {
		err = dbClient.Create(ctx, &newUser.Groups)
		if err != nil {
			l.Errorf("db: %s", err.Error())
			return err
		}
	}

	l.Infof("added user %s", newUser.Email)
	return nil
}
