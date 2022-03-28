package models

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/tyrm/megabot/internal/chatbot"
	"github.com/uptrace/bun"
	"testing"
	"time"
)

func TestChatbotService_BeforeAppendModel_Insert(t *testing.T) {
	obj := &ChatbotService{
		ServiceType: chatbot.ServiceTelegram,
	}

	err := obj.BeforeAppendModel(context.Background(), &bun.InsertQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	emptyTime := time.Time{}
	err = validator.New().Var(obj.ID, "required,ulid")
	if err != nil {
		t.Errorf("invalid id: %s", err.Error())
	}
	if obj.CreatedAt == emptyTime {
		t.Errorf("invalid created at time: %s", obj.CreatedAt.String())
	}
	if obj.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}

func TestChatbotService_BeforeAppendModel_Update(t *testing.T) {
	obj := &ChatbotService{
		ID:          "01FYFXS49Z22W6K1NPBAQ9M0GB",
		ServiceType: chatbot.ServiceTelegram,
	}

	err := obj.BeforeAppendModel(context.Background(), &bun.UpdateQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	emptyTime := time.Time{}
	if obj.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}
