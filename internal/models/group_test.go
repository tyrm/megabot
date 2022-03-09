package models

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestGroup(t *testing.T) {
	tables := []struct {
		x func() uuid.UUID
		n uuid.UUID
	}{
		{GroupSuperAdmin, uuid.Must(uuid.Parse("11a08aec-b7e0-46b4-ba53-e95a858d4cad"))},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Getting id", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			groupID := table.x()
			if groupID != table.n {
				t.Errorf("[%d] got bad id, got: %v, want: %v,", i, groupID, table.n)
			}
		})
	}
}

func TestGroupName(t *testing.T) {
	tables := []struct {
		x string
		n uuid.UUID
	}{
		{"", uuid.Nil},
		{"admin", uuid.Must(uuid.Parse("11a08aec-b7e0-46b4-ba53-e95a858d4cad"))},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Getting name for %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			title := GroupName(table.x)
			if title != table.n {
				t.Errorf("[%d] got bad title for %s, got: %v, want: %v,", i, table.x, title, table.n)
			}
		})
	}
}

func TestGroupTitle(t *testing.T) {
	tables := []struct {
		x uuid.UUID
		n string
	}{
		{uuid.Nil, ""},
		{uuid.Must(uuid.Parse("11a08aec-b7e0-46b4-ba53-e95a858d4cad")), "Super Admin"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Getting title for %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			title := GroupTitle(table.x)
			if title != table.n {
				t.Errorf("[%d] got bad title for %s, got: %v, want: %v,", i, table.x, title, table.n)
			}
		})
	}
}
