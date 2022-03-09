package webapp

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type templateVars interface {
	AddHeadLink(l templateHeadLink)
}

type templateCommon struct {
	HeadLinks []templateHeadLink
	PageTitle string
}

func (t *templateCommon) AddHeadLink(l templateHeadLink) {
	if t.HeadLinks == nil {
		t.HeadLinks = []templateHeadLink{}
	}
	t.HeadLinks = append(t.HeadLinks, l)
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

func (m *Module) initTemplate(w http.ResponseWriter, r *http.Request, tmpl templateVars) error {
	l := logger.WithField("func", "initTemplate")

	// add css
	for _, link := range m.headLinks {
		tmpl.AddHeadLink(link)
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
