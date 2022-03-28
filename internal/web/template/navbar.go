package template

import (
	"regexp"
)

// Navbar is a navbar that can be added to a page
type Navbar []*NavbarNode

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

// SetActive sets the active bool based on the match regex
func (n *Navbar) SetActive(path string) {
	for i := 0; i < len(*n); i++ {
		if (*n)[i].MatchStr != nil {
			if (*n)[i].MatchStr.Match([]byte(path)) {
				(*n)[i].Active = true
			}
		}
		for j := 0; j < len((*n)[i].Children); j++ {
			if (*n)[i].Children[j].MatchStr != nil {
				if (*n)[i].Children[j].MatchStr.Match([]byte(path)) {
					(*n)[i].Active = true
					(*n)[i].Children[j].Active = true
				}
			}
		}
	}
}
