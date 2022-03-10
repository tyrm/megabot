package language

import (
	"fmt"
	"testing"
)

func TestTextEmail(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x string
		n string
	}{
		{"en", "Email"},
		{"es", "Correo Electrónico"},
		{"in", "Email"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Translating to %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			localizer, err := langMod.NewLocalizer(table.x)
			if err != nil {
				t.Errorf("[%d] can't get localizer for %s: %s", i, table.x, err.Error())
				return
			}

			result := localizer.TextEmail()
			if result != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result, table.n)
			}
		})
	}
}
