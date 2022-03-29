package webapp

import (
	"github.com/tyrm/megabot/internal/language"
	"github.com/tyrm/megabot/internal/web/template"
	"net/http"
	"regexp"
)

func makeChatbotSidebar(r *http.Request) template.Sidebar {
	// get localizer
	localizer := r.Context().Value(localizerKey).(*language.Localizer)

	// create sidebar
	newSidebar := template.Sidebar{
		{
			Children: []template.SidebarNode{
				{
					Text:     localizer.TextDashboard().String(),
					MatchStr: regexp.MustCompile("^" + pathBase + pathChatbot + "$"),
					Icon:     "home",
					URL:      pathBase + pathChatbot,
				},
				{
					Text:     localizer.TextServices().String(),
					MatchStr: regexp.MustCompile("^" + pathBase + pathChatbot + pathChatbotServices + "$"),
					Icon:     "person-digging",
					URL:      pathBase + pathChatbot + pathChatbotServices,
				},
			},
		},
	}

	newSidebar.ActivateFromPath(r.URL.Path)

	return newSidebar
}

// ChatbotGetHandler serves the chatbot main page
func (m *Module) ChatbotGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "ChatbotGetHandler")

	// Init template variables
	tmplVars := &template.ChatbotTemplate{
		Common: template.Common{
			PageTitle: "Chatbot",
		},
	}

	tmplVars.Sidebar = makeChatbotSidebar(r)

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
