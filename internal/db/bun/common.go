package bun

import (
	"context"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/db/bun/migrations"
	"github.com/tyrm/megabot/internal/models"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/migrate"
	"os"
)

type commonDB struct {
	bun *Bun
}

// Close closes the bun db connection
func (c *commonDB) Close(ctx context.Context) db.Error {
	l := logger.WithField("func", "Close")
	l.Info("closing db connection")
	return c.bun.Close()
}

// Create inserts an object into the database
func (c *commonDB) Create(ctx context.Context, i interface{}) db.Error {
	_, err := c.bun.NewInsert().Model(i).Exec(ctx)
	return c.bun.ProcessError(err)
}

// DoMigration runs schema migrations on the database
func (c *commonDB) DoMigration(ctx context.Context) db.Error {
	l := logger.WithField("func", "DoMigration")

	migrator := migrate.NewMigrator(c.bun.DB, migrations.Migrations)

	if err := migrator.Init(ctx); err != nil {
		return err
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		if err.Error() == "migrate: there are no any migrations" {
			return nil
		}
		return err
	}

	if group.ID == 0 {
		l.Info("there are no new migrations to run")
		return nil
	}

	l.Infof("migrated database to %s", group)
	return nil
}

func (c *commonDB) LoadTestData(ctx context.Context) db.Error {
	l := logger.WithField("func", "DoMigration")

	// Register models before loading fixtures.
	c.bun.RegisterModel(
		(*models.User)(nil),
		(*models.GroupMembership)(nil),
	)

	// Automatically create tables.
	fixture := dbfixture.New(c.bun.DB, dbfixture.WithTruncateTables())

	// Load fixtures.
	if err := fixture.Load(ctx, os.DirFS("test"), "fixture.yml"); err != nil {
		l.Errorf("loading test data: %s", err.Error())
		return err
	}

	return nil
}
