package webapp

import (
	"bytes"
	"errors"
	"github.com/gorilla/sessions"
	"github.com/tyrm/megabot/internal/language"
	"io"
	"net/http"
	"regexp"
)

type templateVars interface {
	AddHeadLink(l templateHeadLink)
	AddFooterScript(s templateScript)
	SetLocalizer(l *language.Localizer)
	SetNavbar(nodes []templateNavbarNode)
}

type templateCommon struct {
	Localizer *language.Localizer

	HeadLinks     []templateHeadLink
	NavBar        []templateNavbarNode
	PageTitle     string
	FooterScripts []templateScript
}

func (t *templateCommon) AddHeadLink(l templateHeadLink) {
	if t.HeadLinks == nil {
		t.HeadLinks = []templateHeadLink{}
	}
	t.HeadLinks = append(t.HeadLinks, l)
	return
}

func (t *templateCommon) AddFooterScript(s templateScript) {
	if t.FooterScripts == nil {
		t.FooterScripts = []templateScript{}
	}
	t.FooterScripts = append(t.FooterScripts, s)
	return
}

func (t *templateCommon) SetNavbar(nodes []templateNavbarNode) {
	t.NavBar = nodes
	return
}

func (t *templateCommon) SetLocalizer(l *language.Localizer) {
	t.Localizer = l
	return
}

type templateHeadLink struct {
	HRef        string
	Rel         string
	Integrity   string
	CrossOrigin string
	Sizes       string
	Type        string
}

type templateNavbarNode struct {
	Text     string
	URL      string
	MatchStr *regexp.Regexp
	FAIcon   string

	Active   bool
	Disabled bool

	Children []templateNavbarNode
}

type templateScript struct {
	Src         string
	Integrity   string
	CrossOrigin string
}

func (m *Module) initTemplate(w http.ResponseWriter, r *http.Request, tmpl templateVars) error {
	l := logger.WithField("func", "initTemplate")

	// add css
	for _, link := range m.headLinks {
		tmpl.AddHeadLink(link)
	}

	// add scripts
	for _, script := range m.footerScripts {
		tmpl.AddFooterScript(script)
	}

	// navbar
	navbar := makeNavbar(r)
	tmpl.SetNavbar(*navbar)

	// set text handler
	if r.Context().Value(localizerKey) == nil {
		return errors.New("could not get localizer")
	}
	localizer := r.Context().Value(localizerKey).(*language.Localizer)
	tmpl.SetLocalizer(localizer)

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

func (m *Module) executeTemplate(w io.Writer, name string, tmplVars interface{}) error {
	if m.minify == nil {
		return m.templates.ExecuteTemplate(w, name, tmplVars)
	}

	b := new(bytes.Buffer)
	err := m.templates.ExecuteTemplate(b, name, tmplVars)
	if err != nil {
		return err
	}

	return m.minify.Minify("text/html", w, b)
}

func makeNavbar(r *http.Request) *[]templateNavbarNode {

	// create navbar
	newNavbar := []templateNavbarNode{
		{
			Text:     "Home",
			MatchStr: regexp.MustCompile("^/app/$"),
			FAIcon:   "home",
			URL:      "/app/",
		},
	}

	for i := 0; i < len(newNavbar); i++ {
		if newNavbar[i].MatchStr != nil {
			if newNavbar[i].MatchStr.Match([]byte(r.URL.Path)) {
				newNavbar[i].Active = true
			}
		}
		for j := 0; j < len(newNavbar[i].Children); j++ {
			if newNavbar[i].Children[j].MatchStr != nil {
				if newNavbar[i].Children[j].MatchStr.Match([]byte(r.URL.Path)) {
					newNavbar[i].Active = true
					newNavbar[i].Children[j].Active = true
				}
			}
		}
	}

	return &newNavbar
}
