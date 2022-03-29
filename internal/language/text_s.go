package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextServices returns a translated phrase.
func (l *Localizer) TextServices() *LocalizedString {
	lg := logger.WithField("func", "TextChatbot")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Services",
			Description: "the common phrase for services",
			Other:       "Services",
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
