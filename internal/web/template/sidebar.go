package template

import "regexp"

// SidebarSection is a section on a sidebar
type SidebarSection struct {
	Text string

	ButtonIcon  string
	ButtonLabel string
	ButtonURL   string

	Children []SidebarNode
}

// SidebarNode is an entry on a sidebar
type SidebarNode struct {
	Text     string
	URL      string
	MatchStr *regexp.Regexp
	Icon     string

	Active   bool
	Disabled bool
}
