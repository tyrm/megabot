package bun

import (
	"context"
	"fmt"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/models"
	"reflect"
	"testing"
	"time"
)

func TestCommonDB_Close(t *testing.T) {
	client, err := testNewSqliteClient()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	err = client.Close(context.Background())
	if err != nil {
		t.Errorf("unexpected error closing: %s", err.Error())
		return
	}
}

func TestCommonDB_Create(t *testing.T) {
	client, err := testNewSqliteClient()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	tables := []struct {
		object any
		check  func(t *testing.T, o interface{}, b db.DB)
	}{
		{
			object: &models.User{
				Email:             "new@test.com",
				EncryptedPassword: "$2a$14$iU.0NmiiQ5vdQefC77RTMeWpEbBUFsmyOOddo0srZHqXJF7oVC7ye",
			},
			check: testCreateUser,
		},
		{
			object: &models.GroupMembership{
				UserID:  2,
				GroupID: models.GroupSuperAdmin(),
			},
			check: testCreateGroupMembership,
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running create for %s", i, reflect.TypeOf(table.object))
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := client.Create(context.Background(), table.object)
			if err != nil {
				t.Errorf("[%d] got error creating %s: %s", i, reflect.TypeOf(table.object), err.Error())
				return
			}

			table.check(t, table.object, client)
		})
	}
}

func testCreateGroupMembership(t *testing.T, o interface{}, b db.DB) {
	user := o.(*models.GroupMembership)
	emptyTime := time.Time{}

	if user.ID == 0 {
		t.Errorf("invalid ID")
	}
	if user.CreatedAt == emptyTime {
		t.Errorf("invalid created at time, got: '%s', want: '%s'", user.CreatedAt.String(), emptyTime.String())
	}
	if user.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time, got: '%s', want: '%s'", user.UpdatedAt.String(), emptyTime.String())
	}
}

func testCreateUser(t *testing.T, o interface{}, b db.DB) {
	user := o.(*models.User)
	emptyTime := time.Time{}

	if user.ID == 0 {
		t.Errorf("invalid ID")
	}
	if user.Disabled {
		t.Errorf("invalid disabled value, got: '%v', want: 'false'", user.CreatedAt.String())
	}
	if user.CreatedAt == emptyTime {
		t.Errorf("invalid created at time, got: '%s', want: '%s'", user.CreatedAt.String(), emptyTime.String())
	}
	if user.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time, got: '%s', want: '%s'", user.UpdatedAt.String(), emptyTime.String())
	}
}
