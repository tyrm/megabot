package language

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// TextHelloWorld returns a translated phrase.
func (l *Localizer) TextHelloWorld() string {
	lg := logger.WithField("func", "TextHelloWorld")

	text, err := l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "HelloWorld",
			Description: "the phrase: Hello World!",
			Other:       "Hello World!",
		},
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return text
}

// TextLogin returns a translated phrase.
func (l *Localizer) TextLogin() string {
	lg := logger.WithField("func", "TextLogin")

	text, err := l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Login",
			Description: "the word 'Login'",
			Other:       "Login",
		},
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return text
}

// TextLoginInvalid returns a translated phrase.
func (l *Localizer) TextLoginInvalid() string {
	lg := logger.WithField("func", "TextLoginInvalid")

	text, err := l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "LoginInvalid",
			Description: "Explains to the user that the email and password combination are invalid.",
			Other:       "email/password combo invalid",
		},
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return text
}
