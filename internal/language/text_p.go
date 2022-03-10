package language

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// TextPassword returns a translated phrase.
func (l *Localizer) TextPassword() string {
	lg := logger.WithField("func", "TextPassword")

	text, err := l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Password",
			Description: "the common phrase for password",
			Other:       "Password",
		},
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return text
}
