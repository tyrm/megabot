package models

import (
	"time"
)

// GroupMembership as of 20220301055727
type GroupMembership struct {
	ID        int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UserID    int64     `validate:"min=1" bun:",unique:groupmembership,notnull,nullzero"`
	User      *User     `validate:"-" bun:"rel:belongs-to,join:user_id=id"`
	GroupID   []byte    `validate:"required" bun:",unique:groupmembership,notnull,nullzero"`
}
