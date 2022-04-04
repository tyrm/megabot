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

	// get localizer
	localizer := r.Context().Value(localizerKey).(*language.Localizer)

	// Init template variables
	tmplVars := &template.ChatbotServiceTemplate{}
	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmplVars.PageTitle = localizer.TextChatbot().String()
	tmplVars.Sidebar = makeChatbotSidebar(r)

	err = m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = m.executeTemplate(w, "chatbot", tmplVars)
	if err != nil {
		l.Errorf("could not render home template: %s", err.Error())
	}
}

// ChatbotServiceGetHandler serves the chatbot services page
func (m *Module) ChatbotServiceGetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.WithField("func", "ChatbotServiceGetHandler")

	// get localizer
	localizer := r.Context().Value(localizerKey).(*language.Localizer)

	// Init template variables
	tmplVars := &template.ChatbotServiceTemplate{}
	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmplVars.PageTitle = localizer.TextChatbot().String()
	tmplVars.Sidebar = makeChatbotSidebar(r)

	count, err := m.db.CountChatbotServices(r.Context())
	if err != nil {
		l.Errorf("getting count: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page, displayCount, countInURL := getPaginationFromURL(r, defaultCount)
	chatbotServices, err := m.db.ReadChatbotServicesPage(r.Context(), page-1, displayCount)
	if err != nil {
		l.Errorf("getting chatbot page: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmplVars.ChatbotServices = chatbotServices

	hrefCount := 0
	if countInURL {
		hrefCount = displayCount
	}
	pagination := template.MakePagination(page, int(count), displayCount, pathBase+pathChatbot+pathChatbotServices, hrefCount)
	tmplVars.ChatbotServicesPagination = &pagination

	err = m.executeTemplate(w, "chatbot_services", tmplVars)
	if err != nil {
		l.Errorf("could not render home template: %s", err.Error())
	}
}
