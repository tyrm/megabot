package bun

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/megabot/internal/db"
)

type commonDB struct {
	bun *Bun
}

// Close closes the bun db connection
func (c *commonDB) Close(ctx context.Context) db.Error {
	logrus.Info("closing db connection")
	return c.bun.Close()
}

// Create inserts an object into the database
func (c *commonDB) Create(ctx context.Context, i interface{}) db.Error {
	_, err := c.bun.NewInsert().Model(i).Exec(ctx)
	return c.bun.ProcessError(err)
}
