package models

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tmthrgd/go-hex"
	"github.com/uptrace/bun"
	"testing"
	"time"
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

func TestGroupMembership_BeforeAppendModel_Insert(t *testing.T) {
	obj := &GroupMembership{
		UserID: 1,
	}
	err := obj.SetGroupID(uuid.MustParse("957bb260-2a48-464e-91ba-6ac7f7863825"))
	if err != nil {
		t.Errorf("got error setting id: %s", err.Error())
		return
	}

	err = obj.BeforeAppendModel(context.Background(), &bun.InsertQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	emptyTime := time.Time{}
	if obj.CreatedAt == emptyTime {
		t.Errorf("invalid created at time: %s", obj.CreatedAt.String())
	}
	if obj.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}

func TestGroupMembership_BeforeAppendModel_Update(t *testing.T) {
	obj := &GroupMembership{
		UserID: 2,
	}
	err := obj.SetGroupID(uuid.MustParse("957bb260-2a48-464e-91ba-6ac7f7863825"))
	if err != nil {
		t.Errorf("got error setting id: %s", err.Error())
		return
	}

	err = obj.BeforeAppendModel(context.Background(), &bun.UpdateQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	emptyTime := time.Time{}
	if obj.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}

func TestGroupMembership_GetGroupID(t *testing.T) {
	tables := []struct {
		x []byte
		y uuid.UUID
		e string
	}{
		{hex.MustDecodeString("57261eb2a2224497ae76f2b18a5da681"), uuid.MustParse("57261eb2-a222-4497-ae76-f2b18a5da681"), ""},
		{hex.MustDecodeString("11"), uuid.Nil, "invalid UUID (got 1 bytes)"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Comparing %s to %s", i, table.x, table.y)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			u := GroupMembership{
				GroupID: table.x,
			}

			gid, err := u.GetGroupID()
			if table.e == "" {
				if err != nil {
					t.Errorf("[%d] got error getting group id %x: %s", i, table.x, err.Error())
					return
				}
			} else {
				if err == nil {
					t.Errorf("[%d] expected error getting group id %x, got: 'nil', want: '%v'", i, table.x, table.e)
					return
				}
				if err.Error() != table.e {
					t.Errorf("[%d] unexpected error getting group id %x, got: '%s', want: '%s'", i, table.x, err.Error(), table.e)
					return
				}
			}
			if gid != table.y {
				t.Errorf("[%d] got bad data for %x, got: %s, want: %s", i, table.x, gid.String(), table.y.String())
			}
		})
	}
}

func TestGroupMembership_SetGroupID(t *testing.T) {
	tables := []struct {
		x uuid.UUID
		y []byte
		e string
	}{
		{uuid.MustParse("57261eb2-a222-4497-ae76-f2b18a5da681"), hex.MustDecodeString("57261eb2a2224497ae76f2b18a5da681"), ""},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Comparing %s to %s", i, table.x, table.y)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			u := GroupMembership{}

			err := u.SetGroupID(table.x)
			if table.e == "" {
				if err != nil {
					t.Errorf("[%d] got error setting group id %x: %s", i, table.x, err.Error())
					return
				}
			} else {
				if err == nil {
					t.Errorf("[%d] expected error setting group id %x, got: 'nil', want: '%v'", i, table.x, table.e)
					return
				}
				if err.Error() != table.e {
					t.Errorf("[%d] unexpected error setting group id %x, got: '%s', want: '%s'", i, table.x, err.Error(), table.e)
					return
				}
			}
			if bytes.Compare(u.GroupID, table.y) > 0 {
				t.Errorf("[%d] got bad data for %s, got: %x, want: %x,", i, table.x.String(), u.GroupID, table.y)
			}
		})
	}
}
