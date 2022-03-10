package language

import (
	"fmt"
	"testing"
)

func TestTextLogin(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x string
		n string
	}{
		{"en", "Login"},
		{"es", "Iniciar Sesión"},
		{"in", "Login"},
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

			result := localizer.TextLogin()
			if result != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result, table.n)
			}
		})
	}
}

func TestTextLoginInvalid(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x string
		n string
	}{
		{"en", "email/password combo invalid"},
		{"es", "combinación de correo electrónico/contraseña no válida"},
		{"in", "email/password combo invalid"},
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

			result := localizer.TextLoginInvalid()
			if result != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result, table.n)
			}
		})
	}
}

func TestTextLoginShort(t *testing.T) {
	langMod, _ := New()

	tables := []struct {
		x string
		n string
	}{
		{"en", "Login"},
		{"es", "Iniciar"},
		{"in", "Login"},
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

			result := localizer.TextLoginShort()
			if result != table.n {
				t.Errorf("[%d] got invalid translation for %s, got: %v, want: %v,", i, table.x, result, table.n)
			}
		})
	}
}
