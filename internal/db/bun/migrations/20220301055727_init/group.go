package models

import (
	"github.com/google/uuid"
	"time"
)

// GroupMembership as of 20220301055727
type GroupMembership struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UserID    int64     `validate:"min=1" bun:",unique:groupmembership,notnull,nullzero"`
	User      *User     `validate:"-" bun:"rel:belongs-to,join:user_id=id"`
	GroupID   uuid.UUID `validate:"required" bun:",unique:groupmembership,notnull,nullzero"`
}
