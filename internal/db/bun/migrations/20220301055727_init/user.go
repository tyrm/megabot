package models

import "time"

// User as of 20220301055727
type User struct {
	ID                int64              `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt         time.Time          `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt         time.Time          `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Email             string             `validate:"-" bun:",nullzero,notnull,unique"`
	EncryptedPassword string             `validate:"-" bun:""`
	SignInCount       int                `validate:"min=0" bun:",notnull,default:0"`
	Groups            []*GroupMembership `validate:"-" bun:"rel:has-many,join:id=user_id"`
	Disabled          bool               `validate:"-" bun:",notnull,default:false"`
}
