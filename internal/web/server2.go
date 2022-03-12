package web

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/megabot/internal/db"
	"net/http"
	"time"
)

// Server2 is the web server
type Server2 struct {
	router *mux.Router
	srv    *http.Server
}

// HandleFunc attaches a function to a path
func (r *Server2) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.router.HandleFunc(path, f)
}

// PathPrefix attaches a new route url path prefix
func (r *Server2) PathPrefix(tpl string) *mux.Route {
	return r.router.PathPrefix(tpl)
}

// Start starts the web server
func (r *Server2) Start() error {
	logrus.Infof("listening on %s", r.srv.Addr)
	return r.srv.ListenAndServe()
}

// Stop shuts down the web server
func (r *Server2) Stop(ctx context.Context) error {
	return r.srv.Shutdown(ctx)
}

// New2 creates a new http2 web server
func New2(ctx context.Context, db db.DB) (Server, error) {
	r := mux.NewRouter()
	r.Use(handlers.CompressHandler)

	s := &http.Server{
		Addr:         ":5000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &Server2{
		router: r,
		srv:    s,
	}, nil
}
