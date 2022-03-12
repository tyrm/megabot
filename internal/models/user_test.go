package models

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUserInGroup(t *testing.T) {
	user1 := &User{
		Groups: []*GroupMembership{},
	}
	user2 := &User{
		Groups: []*GroupMembership{
			{
				GroupID: uuid.MustParse("35326c82-adac-43f6-a03f-000000000001"),
			},
		},
	}
	user3 := &User{
		Groups: []*GroupMembership{
			{
				GroupID: uuid.MustParse("35326c82-adac-43f6-a03f-000000000001"),
			},
			{
				GroupID: uuid.MustParse("35326c82-adac-43f6-a03f-000000000002"),
			},
			{
				GroupID: uuid.MustParse("35326c82-adac-43f6-a03f-000000000003"),
			},
			{
				GroupID: uuid.MustParse("35326c82-adac-43f6-a03f-000000000004"),
			},
			{
				GroupID: uuid.MustParse("35326c82-adac-43f6-a03f-000000000005"),
			},
		},
	}

	tables := []struct {
		x *User
		y []uuid.UUID
		n bool
	}{
		{user1, []uuid.UUID{}, false},
		{user2, []uuid.UUID{}, false},
		{user3, []uuid.UUID{}, false},
		{user1, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000001")}, false},
		{user2, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000001")}, true},
		{user3, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000001")}, true},
		{user1, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000002")}, false},
		{user2, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000002")}, false},
		{user3, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000002")}, true},
		{user1, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000008")}, false},
		{user2, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000008")}, false},
		{user3, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000008")}, false},
		{user1, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000008"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000009")}, false},
		{user2, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000008"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000009")}, false},
		{user3, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000008"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000009")}, false},
		{user1, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000007"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000010"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000001")}, false},
		{user2, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000007"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000010"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000001")}, true},
		{user3, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000006"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000010"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000001")}, true},
		{user1, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000006"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000002"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000011")}, false},
		{user2, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000006"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000002"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000011")}, false},
		{user3, []uuid.UUID{uuid.MustParse("35326c82-adac-43f6-a03f-000000000006"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000002"), uuid.MustParse("35326c82-adac-43f6-a03f-000000000011")}, true},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running InGroup %v to %s", i, table.x, table.y)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := table.x.InGroup(table.y...)
			if result != table.n {
				t.Errorf("[%d] InGroup wrong for %v, got: %v, want: %v,", i, table.y, result, table.n)
			}

		})
	}
}

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

func BenchmarkUser_CheckPasswordHash(b *testing.B) {
	user := User{
		EncryptedPassword: "$2a$14$iU.0NmiiQ5vdQefC77RTMeWpEbBUFsmyOOddo0srZHqXJF7oVC7ye",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !user.CheckPasswordHash("password") {
			panic("wrong answer")
		}
	}
}

func BenchmarkUser_SetPassword(b *testing.B) {
	user := User{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := user.SetPassword("password")
		if err != nil {
			panic(err)
		}
	}
}
