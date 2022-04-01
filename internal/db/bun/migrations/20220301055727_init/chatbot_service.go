package models

import (
	"github.com/tyrm/megabot/internal/chatbot"
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
