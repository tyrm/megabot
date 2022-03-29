package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextDashboard returns a translated phrase.
func (l *Localizer) TextDashboard() *LocalizedString {
	lg := logger.WithField("func", "TextChatbot")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Dashboard",
			Description: "the common phrase for dashboard",
			Other:       "Dashboard",
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
