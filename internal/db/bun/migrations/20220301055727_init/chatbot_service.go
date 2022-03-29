package models

import (
	"github.com/tyrm/megabot/internal/chatbot"
	"github.com/tyrm/megabot/internal/models"
	"time"
)

// ChatbotService represents a
type ChatbotService struct {
	ID          int64                  `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time              `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time              `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	ServiceType chatbot.ChatService    `validate:"min=1,max=1" bun:",nullzero,notnull"`
	Config      models.EncryptedString `validate:"-" bun:",nullzero,notnull"`
}
