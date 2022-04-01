package models

import (
	"context"
	"github.com/tyrm/megabot/internal/chatbot"
	"github.com/uptrace/bun"
	"time"
)

// ChatbotService represents a
type ChatbotService struct {
	ID          int64               `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time           `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time           `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Description string              `validate:"required" bun:",nullzero,notnull"`
	ServiceType chatbot.ChatService `validate:"min=1,max=2" bun:",nullzero,notnull"`
	Config      []byte              `validate:"-" bun:",nullzero,notnull"`
}

var _ bun.BeforeAppendModelHook = (*ChatbotService)(nil)

// BeforeAppendModel runs before a bun append operation
func (c *ChatbotService) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
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

// GetConfig returns unencrypted config
func (c *ChatbotService) GetConfig() (string, error) {
	data, err := decrypt(c.Config)
	return string(data), err
}

// SetConfig sets encrypted config
func (c *ChatbotService) SetConfig(s string) error {
	data, err := encrypt([]byte(s))
	if err != nil {
		return err
	}
	c.Config = data
	return nil
}
