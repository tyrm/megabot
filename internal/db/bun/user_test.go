package bun

import (
	"context"
	"fmt"
	"github.com/tyrm/megabot/internal/models"
	"testing"
)

var testUser1 = models.User{
	ID:                "01FX66G84MXP9DA4N3P9RVXVMB",
	Email:             "test@test.com",
	EncryptedPassword: "$2a$14$iU.0NmiiQ5vdQefC77RTMeWpEbBUFsmyOOddo0srZHqXJF7oVC7ye",
}
var testUser2 = models.User{
	ID:                "01FX740C6CFRQYW5QP0JEJF20K",
	Email:             "test2@example.com",
	EncryptedPassword: "$2a$14$gleBixsHuNkr/TJGYbkTiOrci1J33778f/Nq39EAn7mlirR87XIx.",
}

func TestUserDB_ReadUserByID(t *testing.T) {
	client, err := testNewTestClient()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}

	tables := []struct {
		user   *models.User
		exists bool
	}{
		{
			user:   &testUser1,
			exists: true,
		},
		{
			user:   &testUser2,
			exists: true,
		},
		{
			user: &models.User{
				ID: "01FY0TJFHGY8ESHJ1CGB7CV61F",
			},
			exists: false,
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running ReadUserByID %v", i, table.user.ID)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			user, err := client.ReadUserByID(context.Background(), table.user.ID)
			if err != nil {
				t.Errorf("[%d] got error reading user %s: %s", i, table.user.ID, err.Error())
				return
			}

			if table.exists {
				if user == nil {
					t.Errorf("[%d] expected user: got 'nil'", i)
					return
				}

				if user.ID != table.user.ID {
					t.Errorf("[%d] wrong id for user: got '%s', want '%s'", i, user.ID, table.user.ID)
				}
				if user.Email != table.user.Email {
					t.Errorf("[%d] wrong email for user: got '%s', want '%s'", i, user.Email, table.user.Email)
				}
				if user.EncryptedPassword != table.user.EncryptedPassword {
					t.Errorf("[%d] wrong id for user: got '%s', want '%s'", i, user.EncryptedPassword, table.user.EncryptedPassword)
				}
			} else {
				if user != nil {
					t.Errorf("[%d] unexpected user: got '%v'", i, user)
					return
				}
			}

		})
	}
}
