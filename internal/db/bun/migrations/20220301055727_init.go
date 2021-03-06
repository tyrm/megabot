package migrations

import (
	"context"
	models "github.com/tyrm/megabot/internal/db/bun/migrations/20220301055727_init"
	"github.com/uptrace/bun"
)

func init() {
	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			modelList := []interface{}{
				&models.User{},
				&models.GroupMembership{},
			}
			for _, i := range modelList {
				if _, err := tx.NewCreateTable().Model(i).IfNotExists().Exec(ctx); err != nil {
					return err
				}
			}

			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}
