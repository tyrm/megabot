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

func TestCommonDB_Create(t *testing.T) {
	client, err := testNewTestClient()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}

	tables := []struct {
		object db.Creatable
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
				UserID:  "01FX740C6CFRQYW5QP0JEJF20K",
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

	if user.ID == "" {
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

	if user.ID == "" {
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
