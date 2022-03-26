package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tyrm/megabot/internal/chatbot"
	"sync"
)

// Service works with the telegram service
type Service struct {
	id     string
	botapi *tgbotapi.BotAPI

	// worker
	bufferSize      int
	updateChan      chan tgbotapi.Update
	workerState     chatbot.WorkerState
	workerStateLock sync.RWMutex
	workerWG        sync.WaitGroup
}

// New creates a new telegram service
func New(id, tgToken string) (chatbot.Service, error) {
	botapi, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		return nil, err
	}

	return &Service{
		id:     id,
		botapi: botapi,

		bufferSize:  100,
		workerState: chatbot.StateStopped,
	}, nil
}
