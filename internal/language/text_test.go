package language

import (
	"fmt"
	"testing"
)

func TestTextHelloWorld(t *testing.T) {
	langMod, err := New()
	if err != nil {
		t.Errorf("can't get language module: %s", err.Error())
		return
	}

	tables := []struct {
		x string
		n string
	}{
		{"en", "Hello World!"},
		{"es", "¡Hola mundo!"},
		{"in", "Hello World!"},
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

			result := localizer.TextHelloWorld()
			if result != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result, table.n)
			}
		})
	}
}
