package webapp

import (
	"github.com/tyrm/megabot/internal/web/template"
	"net/http"
	"regexp"
)

// ChatbotGetHandler serves the chatbot main page
func (m *Module) ChatbotGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "ChatbotGetHandler")

	// Init template variables
	tmplVars := &template.ChatbotTemplate{
		Common: template.Common{
			PageTitle: "Chatbot",
		},
	}

	tmplVars.Sidebar = []template.SidebarNode{
		{
			Children: []template.SidebarNode{
				{
					Text:     "Dashboard",
					MatchStr: regexp.MustCompile("^/app/chatbot$"),
					Icon:     "home",
					URL:      "#",
				},
				{
					Text: "Orders",
					Icon: "file",
					URL:  "#",
				},
				{
					Text: "Products",
					Icon: "cart-shopping",
					URL:  "#",
				},
				{
					Text: "Customers",
					URL:  "#",
				},
				{
					Text: "Reports",
					URL:  "#",
				},
				{
					Text: "Integrations",
					URL:  "#",
				},
			},
		},
		{
			Label: "Saved reports",
			Children: []template.SidebarNode{
				{
					Text: "Current month",
					URL:  "#",
				},
				{
					Text: "Last Quarter",
					URL:  "#",
				},
				{
					Text: "Social engagement",
					URL:  "#",
				},
				{
					Text: "Year-end sale",
					URL:  "#",
				},
			},
		},
	}

	tmplVars.Sidebar.ActivateFromPath(r.URL.Path)

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
