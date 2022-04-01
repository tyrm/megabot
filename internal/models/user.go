package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User represents a human user.
type User struct {
	ID                int64              `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt         time.Time          `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt         time.Time          `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Email             string             `validate:"-" bun:",nullzero,notnull,unique"`
	EncryptedPassword string             `validate:"-" bun:""`
	SignInCount       int                `validate:"min=0" bun:",notnull,default:0"`
	Groups            []*GroupMembership `validate:"-" bun:"rel:has-many,join:id=user_id"`
	Disabled          bool               `validate:"-" bun:",notnull,default:false"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

// BeforeAppendModel runs before a bun append operation
func (u *User) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		u.CreatedAt = now
		u.UpdatedAt = now

		err := validate.Struct(u)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()

		err := validate.Struct(u)
		if err != nil {
			return err
		}
	}
	return nil
}

// CheckPasswordHash is used to validate that a given password matches the stored hash
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	return err == nil
}

// InGroup returns true if user is member of group
func (u *User) InGroup(g ...uuid.UUID) bool {
	l := logger.WithField("func", "InGroup").WithField("model", "User")
	for _, group := range u.Groups {
		for _, needle := range g {
			var gid uuid.UUID
			var err error
			if gid, err = group.GetGroupID(); err != nil {
				l.Warnf("problem reading uuid for group membership %d: %s", group.ID, err.Error())
				continue
			}
			if gid == needle {
				return true
			}
		}
	}
	return false
}

// SetPassword updates the user object's password hash
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u.EncryptedPassword = string(bytes)
	return nil
}
