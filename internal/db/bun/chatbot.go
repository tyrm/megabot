package bun

import (
	"context"
	"database/sql"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/models"
	"github.com/uptrace/bun"
)

type chatbotDB struct {
	bun *Bun
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
	var chatbotServices *[]models.ChatbotService
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
