package telegram

import (
	"context"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/chatbot"
	"github.com/tyrm/megabot/internal/config"
	"net/http"
	"net/url"
)

// GetWorkerState returns the current state of the workers
func (s *Service) GetWorkerState() chatbot.WorkerState {
	s.workerStateLock.RLock()
	defer s.workerStateLock.RUnlock()

	return s.workerState
}

// HandleWebhook handles an incoming message from telegram
func (s *Service) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	update, err := s.botapi.HandleUpdate(r)
	if err != nil {
		errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(errMsg)
		return
	}

	s.updateChan <- *update
}

// StartWorkers starts a number of workers for the given service
func (s *Service) StartWorkers(ctx context.Context, workerCount int) chatbot.Error {
	l := logger.WithField("func", "StartWorkers")

	if !s.setWorkerStateIf(chatbot.StateStarting, chatbot.StateStopped) {
		return chatbot.ErrAlreadyRunning
	}

	// create new update channel
	s.updateChan = make(chan tgbotapi.Update, 100)

	// start Worker
	for w := 1; w <= workerCount; w++ {
		go s.worker(ctx, w)
	}

	// enable telegram webrook
	wh := &tgbotapi.WebhookConfig{
		URL: &url.URL{
			Scheme: "https",
			Host:   viper.GetString(config.Keys.ServerExternalHostname),
			Path:   chatbot.PathWebhook(s.id),
		},
	}
	_, err := s.botapi.Request(wh)
	if err != nil {
		l.Errorf("error creating webhook for %d (telegram)", s.id)

		l.Debugf("closing update channel")
		close(s.updateChan)

		s.setWorkerState(chatbot.StateStopped)

		return chatbot.ErrAPIError
	}

	s.setWorkerState(chatbot.StateRunning)
	return nil
}

// StopWorkers stops the running workers
func (s *Service) StopWorkers(ctx context.Context) chatbot.Error {
	l := logger.WithField("func", "StopWorkers")

	// skip if stopped
	if s.GetWorkerState() == chatbot.StateStopped {
		l.Warnf("tried to stop workers on %d (telegram) but they are not running", s.id)
		return nil
	}

	// ask telegram to stop calling webhook
	_, err := s.botapi.Request(&tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		l.Errorf("error deleting webhook config for %d (telegram)", s.id)
		return chatbot.ErrAPIError
	}

	// close update channel
	l.Debugf("closing update channel")
	close(s.updateChan)

	// wait
	l.Debugf("waiting for workers to finish")
	s.workerWG.Wait()

	s.setWorkerState(chatbot.StateStopped)

	return nil
}

func (s *Service) worker(ctx context.Context, tid int) {
	l := logger.WithField("func", "worker").WithField("thread", tid)
	l.Infof("starting Worker %d", tid)
	s.workerWG.Add(1)

	defer func() {
		l.Infof("Worker %d stopped", tid)
		s.workerWG.Done()
	}()

	for {
		select {
		case update, ok := <-s.updateChan:
			if !ok {
				return // closed
			}

			l.Infof("got update %d: %v", update.UpdateID, update)
		case <-ctx.Done():
			l.Infof("cancelled Worker. Error detail: %v\n", ctx.Err())
			return
		}
	}
}

func (s *Service) setWorkerState(state chatbot.WorkerState) {
	s.workerStateLock.Lock()
	defer s.workerStateLock.Unlock()

	s.workerState = state
	return
}

func (s *Service) setWorkerStateIf(state, ifState chatbot.WorkerState) bool {
	s.workerStateLock.Lock()
	defer s.workerStateLock.Unlock()

	// if current state doesn't match if state, return unsuccessful
	if s.workerState != ifState {
		return false
	}

	// update state and return success
	s.workerState = state
	return true
}
