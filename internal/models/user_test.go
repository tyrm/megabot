package models

import (
	"fmt"
	"testing"
)

func TestUserPasswordHash(t *testing.T) {
	tables := []struct {
		x string
		y string
		n bool
	}{
		{"", "", true},
		{"password", "password", true},
		{"i'm a super long password with $p3c!@L characters!!!!", "i'm a super long password with $p3c!@L characters!!!!", true},
		{"password", "Password", false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Comparing %s to %s", i, table.x, table.y)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			u := User{}

			err := u.SetPassword(table.x)
			if err != nil {
				t.Errorf("[%d] got error setting password %s: %s", i, table.x, err.Error())
				return
			}

			valid := u.CheckPasswordHash(table.y)
			if valid != table.n {
				t.Errorf("[%d] check password failed matching %s to %s, got: %v, want: %v,", i, table.x, table.y, valid, table.n)
			}
		})
	}
}
