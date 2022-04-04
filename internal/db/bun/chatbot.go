package bun

import (
	"context"
	"database/sql"
	"github.com/allegro/bigcache/v3"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/models"
	"github.com/uptrace/bun"
)

const countChatbotServicesKey = "choatbot_services"

type chatbotDB struct {
	bun        *Bun
	countCache *bigcache.BigCache
}

func (c *chatbotDB) CountChatbotServices(ctx context.Context) (int64, db.Error) {
	l := logger.WithField("func", "CountChatbotServices")

	return c.countChatbotServices(
		ctx,
		func() (int64, bool) {
			count, err := c.countCache.Get(countChatbotServicesKey)
			if err == bigcache.ErrEntryNotFound {
				return 0, false
			} else if err != nil {
				l.Errorf("cache error: %s", err.Error())
				return 0, false
			}
			return bytes2int(count), true
		},
		func() (int64, error) {
			count, err := c.newChatbotServiceQ((*models.ChatbotService)(nil)).Count(ctx)
			return int64(count), err
		},
	)

}

func (c *chatbotDB) ReadChatbotServiceByID(ctx context.Context, id int64) (*models.ChatbotService, db.Error) {
	return c.getChatbotService(
		ctx,
		func() (*models.ChatbotService, bool) {
			return nil, false
		},
		func(chatbotService *models.ChatbotService) error {
			return c.newChatbotServiceQ(chatbotService).Where("id = ?", id).Scan(ctx)
		},
	)
}

func (c *chatbotDB) ReadChatbotServicesPage(ctx context.Context, index, count int) (*[]models.ChatbotService, db.Error) {
	chatbotServices := new([]models.ChatbotService)
	err := c.bun.
		NewSelect().
		Model(chatbotServices).
		Limit(count).
		Offset(offset(index, count)).
		Scan(ctx)

	if err != nil {
		return nil, c.bun.ProcessError(err)
	}

	return chatbotServices, nil
}

func (c *chatbotDB) newChatbotServiceQ(chatbotService *models.ChatbotService) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(chatbotService)
}

func (c *chatbotDB) countChatbotServices(ctx context.Context, cacheGet func() (int64, bool), dbQuery func() (int64, error)) (int64, db.Error) {
	l := logger.WithField("func", "countChatbotServices")

	// Attempt to fetch cached account
	chatbotServiceCount, cached := cacheGet()

	if !cached {
		// Not cached! Perform database query
		count, err := dbQuery()
		if err == sql.ErrNoRows {
			return 0, nil
		}
		if err != nil {
			return 0, c.bun.ProcessError(err)
		}
		chatbotServiceCount = count

		// Place in the cache
		err = c.countCache.Set(countChatbotServicesKey, int2bytes(chatbotServiceCount))
		if err != nil {
			l.Errorf("can't update count cache: %s", err.Error())
		}
	}

	return chatbotServiceCount, nil
}

func (c *chatbotDB) getChatbotService(ctx context.Context, cacheGet func() (*models.ChatbotService, bool), dbQuery func(*models.ChatbotService) error) (*models.ChatbotService, db.Error) {
	// Attempt to fetch cached account
	chatbotService, cached := cacheGet()

	if !cached {
		chatbotService = &models.ChatbotService{}

		// Not cached! Perform database query
		err := dbQuery(chatbotService)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, c.bun.ProcessError(err)
		}

		// Place in the cache
		// TODO: u.cache.Put(account)
	}

	return chatbotService, nil
}
