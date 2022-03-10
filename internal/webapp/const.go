package webapp

import "golang.org/x/text/language"

type contextKey int

const (
	sessionKey   contextKey = 0
	localizerKey contextKey = 1
	languageKey  contextKey = 2
)

var defaultLanguage = language.English
