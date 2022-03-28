package webapp

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/tyrm/megabot/internal/language"
	"github.com/tyrm/megabot/internal/models"
	"github.com/tyrm/megabot/internal/web/template"
	"net/http"
	"regexp"
)

func (m *Module) initTemplate(w http.ResponseWriter, r *http.Request, tmpl template.InitTemplate) error {
	l := logger.WithField("func", "initTemplate")

	// set text handler
	localizer := r.Context().Value(localizerKey).(*language.Localizer)
	tmpl.SetLocalizer(localizer)

	// set language
	lang := r.Context().Value(languageKey).(string)
	tmpl.SetLanguage(lang)

	// add css
	for _, link := range m.headLinks {
		tmpl.AddHeadLink(link)
	}

	// add scripts
	for _, script := range m.footerScripts {
		tmpl.AddFooterScript(script)
	}

	// navbar
	navbar := makeNavbar(r, localizer)
	tmpl.SetNavbar(navbar)

	if r.Context().Value(userKey) != nil {
		user := r.Context().Value(userKey).(*models.User)
		tmpl.SetUser(user)
	}

	// try to read session data
	if r.Context().Value(sessionKey) == nil {
		return nil
	}

	us := r.Context().Value(sessionKey).(*sessions.Session)
	saveSession := false

	if saveSession {
		err := us.Save(r, w)
		if err != nil {
			l.Warningf("initTemplate could not save session: %s", err.Error())
			return err
		}
	}

	return nil
}

func (m *Module) executeTemplate(w http.ResponseWriter, name string, tmplVars interface{}) error {
	b := new(bytes.Buffer)
	err := template.Templates.ExecuteTemplate(b, name, tmplVars)
	if err != nil {
		return err
	}

	h := sha512.New()
	h.Write(b.Bytes())
	w.Header().Set("Digest", fmt.Sprintf("sha-512=%s", base64.StdEncoding.EncodeToString(h.Sum(nil))))

	if m.minify == nil {
		_, err := w.Write(b.Bytes())
		return err
	}
	return m.minify.Minify("text/html", w, b)
}

func makeNavbar(r *http.Request, l *language.Localizer) []*template.NavbarNode {
	// create navbar
	newNavbar := []*template.NavbarNode{
		{
			Text:     l.TextHomeShort().String(),
			MatchStr: regexp.MustCompile("^/app/$"),
			FAIcon:   "home",
			URL:      "/app/",
		},
		{
			Text:     l.TextChatbot().String(),
			MatchStr: regexp.MustCompile("^/app/chatbot"),
			FAIcon:   "home",
			URL:      "/app/chatbot",
		},
	}

	template.MakeNavbar(r, newNavbar)

	return newNavbar
}
