package bun

import (
	"context"
	"errors"
	"fmt"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/db/bun/migrations"
	"github.com/tyrm/megabot/internal/models"
	"github.com/tyrm/megabot/internal/models/testdata"
	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/migrate"
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
func (c *commonDB) Create(ctx context.Context, i any) db.Error {
	_, err := c.bun.NewInsert().Model(i).Exec(ctx)
	if err != nil {
		logger.WithField("func", "Create").Errorf("db: %s", err.Error())
	}
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
	l := logger.WithField("func", "LoadTestData")
	l.Debugf("adding test data")

	// Truncate
	modelList := []interface{}{
		&models.User{},
		&models.GroupMembership{},
		&models.ChatbotService{},
	}

	for _, m := range modelList {
		l.Debugf("truncating %T", m)
		_, err := c.bun.NewTruncateTable().Model(m).Exec(ctx)
		if err != nil {
			l.Errorf("truncating %T: %s", m, err.Error())
			return err
		}
	}

	// Create Users
	l.Debugf("creating %d users", len(testdata.TestUsers))
	for i := 0; i < len(testdata.TestUsers); i++ {
		err := c.Create(ctx, testdata.TestUsers[i])
		if err != nil {
			l.Errorf("[%d] creating user: %s", i, err.Error())
			return err
		}
	}

	// Create GroupMembership
	l.Debugf("creating %d group memeberships", len(testdata.TestGroupMembership))
	for i := 0; i < len(testdata.TestGroupMembership); i++ {
		err := c.Create(ctx, testdata.TestGroupMembership[i])
		if err != nil {
			l.Errorf("[%d] creating group membership: %s", i, err.Error())
			return err
		}
	}

	// Create ChatbotServices
	l.Debugf("creating %d chatbot services", len(testdata.TestChatbotServices))
	for i := 0; i < len(testdata.TestChatbotServices); i++ {
		err := c.Create(ctx, testdata.TestChatbotServices[i])
		if err != nil {
			l.Errorf("[%d] creating chatbot srevice: %s", i, err.Error())
			return err
		}
	}

	// fix sequences
	sequences := []struct {
		table        string
		currentValue int
	}{
		{
			table:        "users",
			currentValue: len(testdata.TestUsers),
		},
		{
			table:        "group_memberships",
			currentValue: len(testdata.TestGroupMembership),
		},
		{
			table:        "chatbot_services",
			currentValue: len(testdata.TestChatbotServices),
		},
	}

	switch c.bun.Dialect().Name() {
	case dialect.SQLite:
		// nothing to do
	case dialect.PG:
		for _, s := range sequences {
			_, err := c.bun.Exec("SELECT setval(?, ?, true);", fmt.Sprintf("%s_id_seq", s.table), s.currentValue)
			if err != nil {
				l.Errorf("can't update sequence for %s: %s", s.table, err.Error())
				return err
			}
		}
	case dialect.MySQL:
		for _, s := range sequences {
			_, err := c.bun.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = ?;", s.table), s.currentValue)
			if err != nil {
				l.Errorf("can't update sequence for %s: %s", s.table, err.Error())
				return err
			}
		}
	case dialect.MSSQL:
	default:
		return errors.New("unknown dialect")
	}

	return nil
}
