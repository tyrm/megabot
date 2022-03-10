package webapp

import (
	"fmt"
	"net/http"
)

// ErrorPageTemplate contains the variables for the "error" template.
type ErrorPageTemplate struct {
	templateCommon

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
	signature, err := m.getSignatureCached(fmt.Sprintf("%s/%s", staticDir, pathFileErrorCSS))
	if err != nil {
		l.Errorf("getting signature for %s: %s", pathFileErrorCSS, err.Error())
	}
	tmplVars.AddHeadLink(templateHeadLink{
		HRef:        fmt.Sprintf("%s%s", pathStatic, pathFileErrorCSS),
		Rel:         "stylesheet",
		CrossOrigin: "anonymous",
		Integrity:   signature,
	})

	// set bot image
	switch code {
	case http.StatusBadRequest:
		// 400
		tmplVars.BotImage = fmt.Sprintf("%s%s", pathStatic, pathFileBotConfused)
	case http.StatusUnauthorized:
		// 401
		tmplVars.BotImage = fmt.Sprintf("%s%s", pathStatic, pathFileBotAngry)
	case http.StatusForbidden:
		// 403
		tmplVars.BotImage = fmt.Sprintf("%s%s", pathStatic, pathFileBotMad)
	case http.StatusNotFound:
		// 404
		tmplVars.BotImage = fmt.Sprintf("%s%s", pathStatic, pathFileBotConfused)
	case http.StatusMethodNotAllowed:
		// 405
		tmplVars.BotImage = fmt.Sprintf("%s%s", pathStatic, pathFileBotMad)
	default:
		tmplVars.BotImage = fmt.Sprintf("%s%s", pathStatic, pathFileBotOffline)
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
