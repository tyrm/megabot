package webapp

import (
	"github.com/tyrm/megabot/internal/web/template"
	"net/http"
)

// HomeGetHandler serves the home page
func (m *Module) HomeGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "HomeGetHandler")

	// Init template variables
	tmplVars := &template.HomeTemplate{
		Common: template.Common{
			PageTitle: "Home",
		},
	}

	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.executeTemplate(w, "home", tmplVars)
	if err != nil {
		l.Errorf("could not render home template: %s", err.Error())
	}

}
