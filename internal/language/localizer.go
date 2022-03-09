package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Localizer returns translated phrases
type Localizer struct {
	localizer *i18n.Localizer
}

// NewLocalizer returns a localizer which will return translated phrases based on the provided languages
func (m Module) NewLocalizer(langs ...string) (*Localizer, error) {
	return &Localizer{
		localizer: i18n.NewLocalizer(m.langBundle, langs...),
	}, nil
}
