package chatbot

import (
	"context"
	"net/http"
)

// ChatService represents the chat service the service connects to
type ChatService int

const (
	// ServiceUnknown is an unknown chat service
	ServiceUnknown ChatService = iota
	// ServiceTelegram is the telegram service (https://telegram.org/)
	ServiceTelegram
)

func (s ChatService) String() string {
	switch s {
	case ServiceTelegram:
		return "telegram"
	default:
		return "unknown"
	}
}

// WorkerState represents the current state of the
type WorkerState int

const (
	// StateUnknown worker is in unknown state
	StateUnknown WorkerState = iota
	// StateStopped worker is stopped
	StateStopped
	// StateStarting worker is starting
	StateStarting
	// StateRunning worker is running
	StateRunning
)

func (s WorkerState) String() string {
	switch s {
	case StateStopped:
		return "stopped"
	case StateStarting:
		return "starting"
	case StateRunning:
		return "running"
	default:
		return "unknown"
	}
}

// Service is a fully functional chat service interface
type Service interface {
	ServiceRequester
	ServiceWorker
}

// ServiceRequester functions make calls to a service without a worker running
type ServiceRequester interface {
}

// ServiceWorker functions start/stop and send signals to the service worker
type ServiceWorker interface {
	GetWorkerState() WorkerState
	HandleWebhook(w http.ResponseWriter, r *http.Request)
	StartWorkers(ctx context.Context, workerCount int) Error
	StopWorkers(ctx context.Context) Error
}
