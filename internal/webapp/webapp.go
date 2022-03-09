package webapp

import (
	"github.com/markbates/pkger"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/web"
	"net/http"
)

// Module contains a webapp module for the web server. Implements web.Module
type Module struct {
	db db.DB
}

// New returns a new graphql module
func New(db db.DB) web.Module {
	return &Module{
		db: db,
	}
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleWebapp
}

// Route attaches routes to the web server
func (m *Module) Route(s *web.Server) error {
	// Static Files
	s.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(pkger.Dir("/web/static"))))

	webapp := s.PathPrefix("/app/").Subrouter()
	webapp.Use(m.Middleware)

	return nil
}
