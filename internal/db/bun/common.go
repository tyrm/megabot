package bun

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/megabot/internal/db"
)

type commonDB struct {
	bun *Bun
}

func (b *commonDB) Close(ctx context.Context) db.Error {
	logrus.Info("closing db connection")
	return b.bun.Close()
}

func (c *commonDB) Put(ctx context.Context, i interface{}) db.Error {
	_, err := c.bun.NewInsert().Model(i).Exec(ctx)
	return c.bun.ProcessError(err)
}
