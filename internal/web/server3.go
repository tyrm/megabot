package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/megabot/internal/db"
	"net/http"
	"time"
)

// Server3 is a http 3 web server
type Server3 struct {
	router *mux.Router
	srv    *http3.Server
}

// HandleFunc attaches a function to a path
func (r *Server3) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.router.HandleFunc(path, f)
}

// PathPrefix attaches a new route url path prefix
func (r *Server3) PathPrefix(tpl string) *mux.Route {
	return r.router.PathPrefix(tpl)
}

// Start starts the web server
func (r *Server3) Start() error {
	logrus.Infof("listening on %s", r.srv.Addr)
	return r.srv.ListenAndServe()
}

// Stop shuts down the web server
func (r *Server3) Stop(ctx context.Context) error {
	return r.srv.Shutdown(ctx)
}

// New3 creates a new http 3 web server
func New3(ctx context.Context, db db.DB) (Server, error) {
	r := mux.NewRouter()

	quicConf := &quic.Config{}

	s := &http3.Server{
		Server: &http.Server{
			Addr:         ":5000",
			Handler:      r,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
		QuicConfig: quicConf,
	}

	return &Server3{
		router: r,
		srv:    s,
	}, nil
}
