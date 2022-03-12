package web

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

// Server represents a http server
type Server interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
	PathPrefix(tpl string) *mux.Route
	Start() error
	Stop(ctx context.Context) error
}
