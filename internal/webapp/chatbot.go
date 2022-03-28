package webapp

import (
	"github.com/tyrm/megabot/internal/web/template"
	"net/http"
)

// ChatbotGetHandler serves the chatbot main page
func (m *Module) ChatbotGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "ChatbotGetHandler")

	// Init template variables
	tmplVars := &template.HomeTemplate{
		Common: template.Common{
			PageTitle: "Chatbot",
		},
	}

	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.executeTemplate(w, "chatbot", tmplVars)
	if err != nil {
		l.Errorf("could not render home template: %s", err.Error())
	}

}
