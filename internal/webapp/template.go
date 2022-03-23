package webapp

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/tyrm/megabot"
	"github.com/tyrm/megabot/internal/language"
	"github.com/tyrm/megabot/internal/models"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type templateVars interface {
	AddHeadLink(l templateHeadLink)
	AddFooterScript(s templateScript)
	SetLanguage(l string)
	SetLocalizer(l *language.Localizer)
	SetNavbar(nodes []templateNavbarNode)
	SetUser(user *models.User)
}

type templateCommon struct {
	Language  string
	Localizer *language.Localizer

	FooterScripts []templateScript
	HeadLinks     []templateHeadLink
	NavBar        []templateNavbarNode
	PageTitle     string
	User          *models.User
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

func (t *templateCommon) SetLanguage(l string) {
	t.Language = l
	return
}

func (t *templateCommon) SetLocalizer(l *language.Localizer) {
	t.Localizer = l
	return
}

func (t *templateCommon) SetNavbar(nodes []templateNavbarNode) {
	t.NavBar = nodes
	return
}

func (t *templateCommon) SetUser(user *models.User) {
	t.User = user
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

type templateImage struct {
	Src    string
	Alt    string
	Height int
	Width  int
	Class  string
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

func compileTemplates(path string, suffix string, funcs *template.FuncMap) (*template.Template, error) {
	///l := logger.WithField("func", "compileTemplates")

	tpl := template.New("")
	if funcs != nil {
		tpl.Funcs(*funcs)
	}

	dir, err := megabot.Files.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, d := range dir {
		filePath := path + "/" + d.Name()
		if d.IsDir() || !strings.HasSuffix(d.Name(), suffix) {
			continue
		}

		// open it
		file, err := megabot.Files.Open(filePath)
		if err != nil {
			return nil, err
		}

		// read it
		tmplData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(tmplData))
		if err != nil {
			return nil, err
		}
	}

	return tpl, err
}

func (m *Module) initTemplate(w http.ResponseWriter, r *http.Request, tmpl templateVars) error {
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
	tmpl.SetNavbar(*navbar)

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
	err := m.templates.ExecuteTemplate(b, name, tmplVars)
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

func makeNavbar(r *http.Request, l *language.Localizer) *[]templateNavbarNode {

	// create navbar
	newNavbar := []templateNavbarNode{
		{
			Text:     l.TextHomeShort().String(),
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
