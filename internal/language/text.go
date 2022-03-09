package language

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sirupsen/logrus"
)

// TextHelloWorld returns a translated phrase.
func (l *Localizer) TextHelloWorld() string {
	text, err := l.localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "HelloWorld",
			Description: "the phrase: Hello World!",
			Other:       "Hello World!",
		},
	})
	if err != nil {
		logrus.Warningf("missing translation: %s", err.Error())
	}
	return text
}
