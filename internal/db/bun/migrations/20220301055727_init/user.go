package models

import "time"

type User struct {
	ID                string    `validate:"required,ulid" bun:"type:CHAR(26),pk,nullzero,notnull,unique"`
	CreatedAt         time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt         time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Email             string    `validate:"required_with=ConfirmedAt" bun:",nullzero,unique"`
	EncryptedPassword string    `validate:"required" bun:",nullzero,notnull"`
	SignInCount       int       `validate:"min=0" bun:",notnull,default:0"`
	Admin             bool      `validate:"-" bun:",notnull,default:false"`
	Disabled          bool      `validate:"-" bun:",notnull,default:false"`
}