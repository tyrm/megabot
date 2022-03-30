package chatbot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tyrm/megabot/internal/chatbot"
	"io/ioutil"
	"net/http"
	"sync"
)

// Mock is a fake chatbot service
type Mock struct {
	id int64

	// worker
	updateChan      chan tgbotapi.Update
	workerState     chatbot.WorkerState
	workerStateLock sync.RWMutex
}

// New creates a new telegram service
func New(id int64) (chatbot.Service, error) {
	l := logger.WithField("func", "New")
	l.Warnf("[%d] mock is for testing purposes only", id)
	return &Mock{
		id: id,
	}, nil
}

// GetWorkerState returns the current state of the workers
func (m *Mock) GetWorkerState() chatbot.WorkerState {
	m.workerStateLock.RLock()
	w := m.workerState
	m.workerStateLock.RUnlock()

	return w
}

// HandleWebhook handles an incoming message from telegram
func (m *Mock) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "StartWorkers")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Errorf("[%d] can't read body", m.id)
		return
	}

	l.Infof("[%d] webhook called: %s", m.id, body)
	return
}

// StartWorkers starts a number of workers for the given service
func (m *Mock) StartWorkers(ctx context.Context, workerCount int) chatbot.Error {
	l := logger.WithField("func", "StartWorkers")

	if !m.setWorkerStateIf(chatbot.StateStarting, chatbot.StateStopped) {
		l.Infof("[%d](mock) can't start workers because they're already running", m.id)
		return chatbot.ErrAlreadyRunning
	}

	m.setWorkerState(chatbot.StateRunning)
	return nil
}

// StopWorkers stops the running workers
func (m *Mock) StopWorkers(ctx context.Context) chatbot.Error {
	l := logger.WithField("func", "StopWorkers")

	// skip if stopped
	if m.GetWorkerState() == chatbot.StateStopped {
		l.Infof("[%d](mock) tried to stop workers but they are not running", m.id)
		return nil
	}

	l.Infof("[%d](mock) stopping workers", m.id)
	m.setWorkerState(chatbot.StateStopped)
	return nil
}

func (m *Mock) setWorkerState(state chatbot.WorkerState) {
	l := logger.WithField("func", "setWorkerState")
	l.Infof("setting worker state: %s", state)
	m.workerStateLock.Lock()
	m.workerState = state
	m.workerStateLock.Unlock()
	return
}

func (m *Mock) setWorkerStateIf(state, ifState chatbot.WorkerState) bool {
	m.workerStateLock.Lock()

	// if current state doesn't match if state, return unsuccessful
	if m.workerState != ifState {
		m.workerStateLock.Unlock()
		return false
	}

	// update state and return success
	m.workerState = state
	m.workerStateLock.Unlock()
	return true
}
