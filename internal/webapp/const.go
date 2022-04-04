package webapp

import "golang.org/x/text/language"

type contextKey int

const (
	sessionKey contextKey = iota
	localizerKey
	languageKey
	userKey
)

var defaultLanguage = language.English

const (
	defaultCount = 10
)
