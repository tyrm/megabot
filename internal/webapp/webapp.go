package webapp

import (
	"github.com/markbates/pkger"
	"github.com/tyrm/go-util/pkgerutil"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/web"
	"html/template"
	"net/http"
)

var tmplFuncs = template.FuncMap{
	"dec": func(i int) int {
		i--
		return i
	},
	"htmlSafe": func(html string) template.HTML {
		/* #nosec G203 */
		return template.HTML(html)
	},
	"inc": func(i int) int {
		i++
		return i
	},
}

// Module contains a webapp module for the web server. Implements web.Module
type Module struct {
	db        db.DB
	templates *template.Template
}

// New returns a new webapp module
func New(db db.DB) (web.Module, error) {
	// Load Templates
	t, err := pkgerutil.CompileTemplates(pkger.Include("/web/template"), "", &tmplFuncs)
	if err != nil {
		return nil, err
	}

	return &Module{
		db:        db,
		templates: t,
	}, nil
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleWebapp
}

// Route attaches routes to the web server
func (m *Module) Route(s *web.Server) error {
	// Static Files
	s.PathPrefix("/static/").Handler(http.StripPrefix(
		"/static/", http.FileServer(pkger.Dir("/web/static"))))

	webapp := s.PathPrefix("/app/").Subrouter()
	webapp.Use(m.Middleware)
	webapp.HandleFunc(pathHome, m.HomeGetHandler).Methods("GET")

	return nil
}
