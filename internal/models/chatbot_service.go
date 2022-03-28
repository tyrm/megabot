package models

import (
	"context"
	"github.com/tyrm/megabot/internal/chatbot"
	"github.com/tyrm/megabot/internal/id"
	"github.com/uptrace/bun"
	"time"
)

// ChatbotService represents a
type ChatbotService struct {
	ID          string              `validate:"required,ulid" bun:"type:CHAR(26),pk,nullzero,notnull,unique"`
	CreatedAt   time.Time           `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time           `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	ServiceType chatbot.ChatService `validate:"min=1,max=1" bun:",nullzero,notnull"`
	Config      EncryptedString     `validate:"-" bun:",nullzero,notnull"`
}

var _ bun.BeforeAppendModelHook = (*ChatbotService)(nil)

// BeforeAppendModel runs before a bun append operation
func (c *ChatbotService) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if c.ID == "" {
			newID, err := id.NewULID()
			if err != nil {
				return err
			}
			c.ID = newID
		}

		now := time.Now()
		c.CreatedAt = now
		c.UpdatedAt = now

		err := validate.Struct(c)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		c.UpdatedAt = time.Now()

		err := validate.Struct(c)
		if err != nil {
			return err
		}
	}
	return nil
}
