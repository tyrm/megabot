package models

import (
	"context"
	"github.com/tyrm/megabot/internal/chatbot"
	"github.com/uptrace/bun"
	"testing"
	"time"
)

func TestChatbotService_BeforeAppendModel_Insert(t *testing.T) {
	obj := &ChatbotService{
		Description: "test 1",
		ServiceType: chatbot.ServiceTelegram,
	}

	err := obj.BeforeAppendModel(context.Background(), &bun.InsertQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	emptyTime := time.Time{}
	if obj.CreatedAt == emptyTime {
		t.Errorf("invalid created at time: %s", obj.CreatedAt.String())
	}
	if obj.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}

func TestChatbotService_BeforeAppendModel_Update(t *testing.T) {
	obj := &ChatbotService{
		Description: "test 1",
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
