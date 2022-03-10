package language

import (
	"github.com/BurntSushi/toml"
	"github.com/markbates/pkger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var defaultLanguage = language.English

var translationFiles = map[string]string{
	"active.es.toml": pkger.Include("/active.es.toml"),
}

// Module represent the language module for translating text
type Module struct {
	lang       language.Tag
	langBundle *i18n.Bundle
}

// New creates a new language module
func New() (*Module, error) {
	l := logger.WithField("func", "New")

	module := Module{
		lang:       defaultLanguage,
		langBundle: i18n.NewBundle(defaultLanguage),
	}

	module.langBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for filename, file := range translationFiles {
		langFile, err := pkger.Open(file)
		if err != nil {
			l.Errorf("opening file: %s", err.Error())
			return nil, err
		}
		defer langFile.Close()

		fileinfo, err := langFile.Stat()
		if err != nil {
			l.Errorf("stating file: %s", err.Error())
			return nil, err
		}

		filesize := fileinfo.Size()
		buffer := make([]byte, filesize)

		_, err = langFile.Read(buffer)
		if err != nil {
			l.Errorf("reading buffer: %s", err.Error())
			return nil, err
		}

		module.langBundle.MustParseMessageFileBytes(buffer, filename)
	}

	return &module, nil
}

// Language returns the default language
func (m Module) Language() language.Tag { return m.lang }
