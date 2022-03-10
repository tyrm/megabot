package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextEmail returns a translated phrase.
func (l *Localizer) TextEmail() *LocalizedString {
	lg := logger.WithField("func", "TextEmail")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Email",
			Description: "the common phrase for email",
			Other:       "Email",
		},
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return &LocalizedString{
		language: tag,
		string:   text,
	}
}
