package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextChatbot returns a translated phrase.
func (l *Localizer) TextChatbot() *LocalizedString {
	lg := logger.WithField("func", "TextChatbot")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Chatbot",
			Description: "Chatbot common phrase for chatbot",
			Other:       "Chatbot",
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
