package webapp

import (
	"fmt"
	template2 "github.com/tyrm/megabot/internal/web/template"
	"net/http"
)

// ErrorPageTemplate contains the variables for the "error" template.
type ErrorPageTemplate struct {
	template2.Common

	BotImage    string
	Header      string
	SubHeader   string
	Paragraph   string
	ButtonHRef  string
	ButtonLabel string
}

func (m *Module) returnErrorPage(w http.ResponseWriter, r *http.Request, code int, errStr string) {
	l := logger.WithField("func", "returnErrorPage")

	// Init template variables
	tmplVars := &ErrorPageTemplate{}
	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// add error css file
	signature, err := m.getSignatureCached(staticDir + pathFileErrorCSS)
	if err != nil {
		l.Errorf("getting signature for %s: %s", pathFileErrorCSS, err.Error())
	}
	tmplVars.AddHeadLink(template2.HeadLink{
		HRef:        pathStatic + pathFileErrorCSS,
		Rel:         "stylesheet",
		CrossOrigin: "anonymous",
		Integrity:   signature,
	})

	// set bot image
	switch code {
	case http.StatusBadRequest:
		// 400
		tmplVars.BotImage = pathStatic + pathFileBotConfused
	case http.StatusUnauthorized:
		// 401
		tmplVars.BotImage = pathStatic + pathFileBotAngry
	case http.StatusForbidden:
		// 403
		tmplVars.BotImage = pathStatic + pathFileBotMad
	case http.StatusNotFound:
		// 404
		tmplVars.BotImage = pathStatic + pathFileBotConfused
	case http.StatusMethodNotAllowed:
		// 405
		tmplVars.BotImage = pathStatic + pathFileBotMad
	default:
		tmplVars.BotImage = pathStatic + pathFileBotOffline
	}

	// set text
	tmplVars.Header = fmt.Sprintf("%d", code)
	tmplVars.SubHeader = http.StatusText(code)
	tmplVars.PageTitle = fmt.Sprintf("%d - %s", code, http.StatusText(code))
	tmplVars.Paragraph = errStr

	// set top button
	switch code {
	case http.StatusUnauthorized:
		tmplVars.ButtonHRef = "/login"
		tmplVars.ButtonLabel = "Login"
	default:
		tmplVars.ButtonHRef = "/app/"
		tmplVars.ButtonLabel = "Home"
	}

	w.WriteHeader(code)
	err = m.executeTemplate(w, "error", tmplVars)
	if err != nil {
		logger.Errorf("could not render error template: %s", err.Error())
	}
}

func (m *Module) methodNotAllowedHandler() http.Handler {
	// wrap in middleware since middlware isn't run on error pages
	return m.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.returnErrorPage(w, r, http.StatusMethodNotAllowed, r.Method)
	}))
}

func (m *Module) notFoundHandler() http.Handler {
	// wrap in middleware since middlware isn't run on error pages
	return m.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.returnErrorPage(w, r, http.StatusNotFound, fmt.Sprintf("page not found: %s", r.URL.Path))
	}))
}
