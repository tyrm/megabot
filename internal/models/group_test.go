package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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
