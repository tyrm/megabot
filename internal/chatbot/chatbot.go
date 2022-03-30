package chatbot

import (
	"github.com/gorilla/mux"
	"github.com/tyrm/megabot/internal/web"
	"net/http"
	"sync"
)

// Module is a chatbot module which receives messages from a chat service and processes them
type Module struct {
	serviceWorkers     map[string]Service
	serviceWorkersLock sync.RWMutex
}

// New creates a new chatbot module
func New() (*Module, error) {
	return &Module{
		serviceWorkers: make(map[string]Service),
	}, nil
}

// AddServiceWorker adds a worker to the chatbot for handling incoming messages
func (m *Module) AddServiceWorker(token string, sw Service) error {
	m.serviceWorkersLock.Lock()
	defer m.serviceWorkersLock.Unlock()

	// don't re-add a service
	if _, ok := m.serviceWorkers[token]; ok {
		return ErrIDExists
	}

	// add service
	m.serviceWorkers[token] = sw
	return nil
}

// GetServiceWorker retrieves a service worker
func (m *Module) GetServiceWorker(token string) (Service, bool) {
	m.serviceWorkersLock.RLock()
	defer m.serviceWorkersLock.RUnlock()

	sw, ok := m.serviceWorkers[token]
	return sw, ok
}

// Route attaches routes to the web server
func (m *Module) Route(ws web.Server) error {
	ws.HandleFunc(pathBase+"/{token}", m.handleWebhook).Methods("POST")
	return nil
}

func (m *Module) handleWebhook(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	sw, ok := m.GetServiceWorker(token)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sw.HandleWebhook(w, r)
}
