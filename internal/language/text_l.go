package language

import "github.com/nicksnyder/go-i18n/v2/i18n"

// TextLogin returns a translated phrase.
func (l *Localizer) TextLogin() string {
	lg := logger.WithField("func", "TextLogin")

	text, err := l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "Login",
			Description: "the common phrase for logging in",
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

// TextLoginShort returns a translated phrase.
func (l *Localizer) TextLoginShort() string {
	lg := logger.WithField("func", "TextLoginShort")

	text, err := l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "LoginShort",
			Description: "a single word representation of logging in",
			Other:       "Login",
		},
	})
	if err != nil {
		lg.Warningf("missing translation: %s", err.Error())
	}
	return text
}
