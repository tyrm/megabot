package models

import (
	"github.com/google/uuid"
	"time"
)

type GroupMembership struct {
	ID        string    `validate:"required,ulid" bun:"type:CHAR(26),pk,nullzero,notnull,unique"`
	CreatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UserID    string    `validate:"required,ulid" bun:"type:CHAR(26),unique:groupmembership,notnull,nullzero"`
	User      *User     `validate:"-" bun:"rel:belongs-to"`
	GroupID   uuid.UUID `validate:"required,uuid" bun:",unique:groupmembership,notnull,nullzero"`
}
