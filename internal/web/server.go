package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/megabot/internal/db"
	"net/http"
	"time"
)

// Server is the web server
type Server struct {
	router *mux.Router
	srv    *http.Server
}

// HandleFunc attaches a function to a path
func (r *Server) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.router.HandleFunc(path, f)
}

// Start starts the web server
func (r *Server) Start() error {
	logrus.Infof("listening on %s", r.srv.Addr)
	return r.srv.ListenAndServe()
}

// Stop shuts down the web server
func (r *Server) Stop(ctx context.Context) error {
	return r.srv.Shutdown(ctx)
}

// New creates a new web server
func New(ctx context.Context, db db.DB) (*Server, error) {
	r := mux.NewRouter()

	s := &http.Server{
		Addr:         ":5000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &Server{
		router: r,
		srv:    s,
	}, nil
}
