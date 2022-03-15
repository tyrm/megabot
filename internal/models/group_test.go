package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
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

func TestGroupMembership_GenID(t *testing.T) {
	validate := validator.New()

	for n := 0; n < 5; n++ {
		i := n

		name := fmt.Sprintf("[%d] Running user GenID", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gm := GroupMembership{}

			err := gm.GenID()
			if err != nil {
				t.Errorf("[%d] got error generating id; %s", i, err)
				return
			}

			err = validate.Var(gm.ID, "required,ulid")
			if err != nil {
				t.Errorf("[%d] id '%s' is invalid: %s", i, gm.ID, err.Error())
				return
			}
		})
	}
}

func TestGroupMemberships_GenID(t *testing.T) {
	validate := validator.New()

	for n := 0; n < 5; n++ {
		i := n

		name := fmt.Sprintf("[%d] Running user GenID", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gms := GroupMemberships{
				{},
				{},
				{},
				{},
				{},
			}

			err := gms.GenID()
			if err != nil {
				t.Errorf("[%d] got error generating id; %s", i, err)
				return
			}

			for _, gm := range gms {
				err = validate.Var(gm.ID, "required,ulid")
				if err != nil {
					t.Errorf("[%d] id '%s' is invalid: %s", i, gm.ID, err.Error())
					return
				}
			}
		})
	}
}
