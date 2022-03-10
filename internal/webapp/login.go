package webapp

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/tyrm/megabot/internal/language"
	"net/http"
)

// LoginTemplate contains the variables for the "login" template.
type LoginTemplate struct {
	templateCommon

	BotImage string

	FormError    string
	FormEmail    string
	FormPassword string
}

// LoginGetHandler serves the login page
func (m *Module) LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	m.displayLoginPage(w, r, "", "", pathFileBotHappy, "")
}

// LoginPostHandler attempts a login
func (m *Module) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	// get localizer
	localizer := r.Context().Value(localizerKey).(*language.Localizer)

	// parse form data
	err := r.ParseForm()
	if err != nil {
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// check if user exists
	formEmail := r.Form.Get("email")
	formPassword := r.Form.Get("password")
	user, err := m.db.ReadUserByEmail(r.Context(), formEmail)
	if err != nil {
		logger.Errorf("db error: %s", err.Error())
		m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		m.displayLoginPage(w, r, formEmail, formPassword, pathFileBotMad, localizer.TextLoginInvalid())
		return
	}

	// check password validity
	passValid := user.CheckPasswordHash(formPassword)
	if passValid == false {
		m.displayLoginPage(w, r, formEmail, formPassword, pathFileBotMad, localizer.TextLoginInvalid())
		return
	}

	// Init Session
	us := r.Context().Value(sessionKey).(*sessions.Session)
	us.Values["user"] = user
	err = us.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// redirect to last page
	val := us.Values["login-redirect"]
	var loginRedirect string
	var ok bool
	if loginRedirect, ok = val.(string); !ok {
		// redirect home page if no login-redirect
		http.Redirect(w, r, "/app/", http.StatusFound)
		return
	}

	http.Redirect(w, r, loginRedirect, http.StatusFound)
	return
}

func (m *Module) displayLoginPage(w http.ResponseWriter, r *http.Request, email, password, botImage, formError string) {
	l := logger.WithField("func", "LoginGetHandler")

	// Init template variables
	tmplVars := &LoginTemplate{}
	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// add error css file
	signature, err := m.getSignatureCached(fmt.Sprintf("%s/%s", staticDir, pathFileLoginCSS))
	if err != nil {
		l.Errorf("getting signature for %s: %s", pathFileLoginCSS, err.Error())
	}
	tmplVars.AddHeadLink(templateHeadLink{
		HRef:        fmt.Sprintf("%s%s", pathStatic, pathFileLoginCSS),
		Rel:         "stylesheet",
		CrossOrigin: "anonymous",
		Integrity:   signature,
	})

	tmplVars.PageTitle = "Login"

	// set bot image
	tmplVars.BotImage = fmt.Sprintf("%s%s", pathStatic, botImage)

	// set form values
	tmplVars.FormError = formError
	tmplVars.FormEmail = email
	tmplVars.FormPassword = password

	err = m.executeTemplate(w, "login", tmplVars)
	if err != nil {
		l.Errorf("could not render login template: %s", err.Error())
	}
}
