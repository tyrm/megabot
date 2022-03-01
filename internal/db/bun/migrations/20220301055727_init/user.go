package models

import "time"

type User struct {
	ID                string             `validate:"required,ulid" bun:"type:CHAR(26),pk,nullzero,notnull,unique"`
	CreatedAt         time.Time          `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt         time.Time          `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Email             string             `validate:"-" bun:",nullzero,notnull,unique"`
	EncryptedPassword []byte             `validate:"-" bun:""`
	SignInCount       int                `validate:"min=0" bun:",notnull,default:0"`
	Groups            []*GroupMembership `validate:"-" bun:"rel:has-many,join:id=user_id"`
	Admin             bool               `validate:"-" bun:",notnull,default:false"`
	Disabled          bool               `validate:"-" bun:",notnull,default:false"`
}
