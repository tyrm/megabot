package template

import (
	"net/http"
	"regexp"
)

// NavbarNode is an entry on a navbar, can be nested one level
type NavbarNode struct {
	Text     string
	URL      string
	MatchStr *regexp.Regexp
	FAIcon   string

	Active   bool
	Disabled bool

	Children []NavbarNode
}

// MakeNavbar sets the active attribute based on the request url
func MakeNavbar(r *http.Request, newNavbar []*NavbarNode) {
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
	return
}
