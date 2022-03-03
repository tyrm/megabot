package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User represents a human user.
type User struct {
	ID                string             `validate:"required,ulid" bun:"type:CHAR(26),pk,nullzero,notnull,unique"`
	CreatedAt         time.Time          `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt         time.Time          `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Email             string             `validate:"-" bun:",nullzero,notnull,unique"`
	EncryptedPassword string             `validate:"-" bun:""`
	SignInCount       int                `validate:"min=0" bun:",notnull,default:0"`
	Groups            []*GroupMembership `validate:"-" bun:"rel:has-many,join:id=user_id"`
	Disabled          bool               `validate:"-" bun:",notnull,default:false"`
}

// CheckPasswordHash is used to validate that a given password matches the stored hash
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	return err == nil
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
