package bun

import (
	"context"
	"database/sql"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/models"
	"github.com/uptrace/bun"
)

type userDB struct {
	bun *Bun
}

func (u *userDB) newUserQ(user *models.User) *bun.SelectQuery {
	return u.bun.
		NewSelect().
		Model(user)
}

func (u *userDB) ReadUserByID(ctx context.Context, id string) (*models.User, db.Error) {
	return u.getUser(
		ctx,
		func() (*models.User, bool) {
			return nil, false
		},
		func(user *models.User) error {
			return u.newUserQ(user).Where("id = ?", id).Scan(ctx)
		},
	)
}

func (u *userDB) ReadUserByEmail(ctx context.Context, email string) (*models.User, db.Error) {
	return u.getUser(
		ctx,
		func() (*models.User, bool) {
			return nil, false
		},
		func(user *models.User) error {
			return u.newUserQ(user).Where("email = ?", email).Scan(ctx)
		},
	)
}

func (u *userDB) getUser(ctx context.Context, cacheGet func() (*models.User, bool), dbQuery func(*models.User) error) (*models.User, db.Error) {
	// Attempt to fetch cached account
	user, cached := cacheGet()

	if !cached {
		user = &models.User{}

		// Not cached! Perform database query
		err := dbQuery(user)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, u.bun.ProcessError(err)
		}

		// Place in the cache
		// TODO: u.cache.Put(account)
	}

	return user, nil
}
