package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/megabot/internal/db"
	"net/http"
	"time"
)

type Server interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
	// Start the web
	Start() error
	// Stop the web
	Stop(ctx context.Context) error
}

type server struct {
	router *mux.Router
	srv    *http.Server
}

func (r *server) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.router.HandleFunc(path, f)
}

func (r *server) Start() error {
	logrus.Infof("listening on %s", r.srv.Addr)
	return r.srv.ListenAndServe()
}

func (r *server) Stop(ctx context.Context) error {
	return r.srv.Shutdown(ctx)
}

func NewServer(ctx context.Context, db db.DB) (Server, error) {
	r := mux.NewRouter()

	s := &http.Server{
		Addr:         ":5000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &server{
		router: r,
		srv:    s,
	}, nil
}
