package boostrapcheatsheet

import (
	"context"
	"github.com/tyrm/megabot"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/web"
	"io/fs"
	"net/http"
)

// Module contains a test page for bootstrap. Implements web.Module
type Module struct{}

// New returns a new webapp module
func New(_ context.Context) (web.Module, error) {
	//l := logger.WithField("func", "New")
	return &Module{}, nil
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleBootstrap
}

// Route attaches routes to the web server
func (m *Module) Route(s web.Server) error {
	staticFS, err := fs.Sub(megabot.Files, cheatsheetDir)
	if err != nil {
		return err
	}

	// Static Files
	s.PathPrefix(cheatsheetBase + "/").Handler(http.StripPrefix(cheatsheetBase+"/", http.FileServer(http.FS(staticFS))))

	return nil
}
