package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextHelloWorld returns a translated phrase.
func (l *Localizer) TextHelloWorld() *LocalizedString {
	lg := logger.WithField("func", "TextHelloWorld")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "HelloWorld",
			Description: "the phrase: Hello World!",
			Other:       "Hello World!",
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

// TextHomeShort returns a translated phrase.
func (l *Localizer) TextHomeShort() *LocalizedString {
	lg := logger.WithField("func", "TextHomeShort")

	text, tag, err := l.localizer.LocalizeWithTag(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "HomeShort",
			Description: "a single word representation of home, as in home page.",
			Other:       "Home",
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
