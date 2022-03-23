package filestore

import (
	"fmt"
	"testing"
)

func TestGetSuffix(t *testing.T) {
	tables := []struct {
		x string
		n string
		e error
	}{
		{"image/jpeg", "jpg", nil},
		{"image/png", "png", nil},
		{"application/zip", "", ErrUnknownType},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] running GetSuffix on %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res, err := GetSuffix(table.x)
			if err != table.e {
				t.Errorf("[%d] wrong error, got: %v, want: %v", i, err, table.e)
				return
			}
			if res != table.n {
				t.Errorf("[%d] wrong suffix, got: %s, want: %vs", i, table.x, table.n)
			}
		})
	}
}
