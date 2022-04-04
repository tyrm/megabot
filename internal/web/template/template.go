package template

import (
	"github.com/google/uuid"
	"github.com/tyrm/megabot"
	"github.com/tyrm/megabot/internal/language"
	"github.com/tyrm/megabot/internal/models"
	"github.com/tyrm/megabot/internal/token"
	"html/template"
	"io/ioutil"
	"strings"
)

const templateDir = "web/template"

// InitTemplate are the functions a template implementing Common will have
type InitTemplate interface {
	AddHeadLink(l HeadLink)
	AddFooterScript(s Script)
	SetLanguage(l string)
	SetLocalizer(l *language.Localizer)
	SetNavbar(nodes Navbar)
	SetUser(user *models.User)
}

// New creates a new tokenizer
func New(t *token.Tokenizer) (*template.Template, error) {
	tpl := template.New("")
	tpl.Funcs(template.FuncMap{
		"dec": func(i int) int {
			i--
			return i
		},
		"groupSuperAdmin": func() uuid.UUID {
			return models.GroupSuperAdmin()
		},
		"htmlSafe": func(html string) template.HTML {
			/* #nosec G203 */
			return template.HTML(html)
		},
		"inc": func(i int) int {
			i++
			return i
		},
		"token": t.GetToken,
	})

	dir, err := megabot.Files.ReadDir(templateDir)
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		filePath := templateDir + "/" + d.Name()
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".gohtml") {
			continue
		}

		// open it
		file, err := megabot.Files.Open(filePath)
		if err != nil {
			return nil, err
		}

		// read it
		tmplData, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(tmplData))
		if err != nil {
			return nil, err
		}
	}

	return tpl, nil
}
