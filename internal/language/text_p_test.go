package language

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
)

func TestTextPassword(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x language.Tag
		n string
		l language.Tag
	}{
		{language.English, "Password", language.English},
		{language.Spanish, "Contraseña", language.Spanish},
		{language.Hindi, "Password", language.English},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Translating to %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			localizer, err := langMod.NewLocalizer(table.x.String())
			if err != nil {
				t.Errorf("[%d] can't get localizer for %s: %s", i, table.x, err.Error())
				return
			}

			result := localizer.TextPassword()
			if result.String() != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result.String(), table.n)
			}
			if result.Language() != table.l {
				t.Errorf("[%d] got invalid language for %s, got: %v, want: %v,", i, table.x, result.Language(), table.l)
			}
		})
	}
}
