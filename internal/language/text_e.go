package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextEmail returns a translated phrase.
func (l *Localizer) TextEmail() string {
	lg := logger.WithField("func", "TextEmail")

	text, err := l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Email",
			Description: "the common phrase for email",
			Other:       "Email",
		},
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return text
}
