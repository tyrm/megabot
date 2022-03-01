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
	GroupID   uuid.UUID `validate:"required,uuid" bun:"type:CHAR(26),unique:groupmembership,notnull,nullzero"`
}

// GroupSuperAdmin is the uuid of the Super Administrators group
var groupSuperAdmin = uuid.Must(uuid.Parse("11a08aec-b7e0-46b4-ba53-e95a858d4cad"))

// GroupTitle contains the titles of the groups.
var groupTitle = map[uuid.UUID]string{
	groupSuperAdmin: "Super Admin",
}

// GroupSuperAdmin returns the uuid for the Super Admin
func GroupSuperAdmin() uuid.UUID {
	return groupSuperAdmin
}

// GroupTitle return a pretty text name for the group
func GroupTitle(g uuid.UUID) string {
	if s, ok := groupTitle[g]; ok {
		return s
	}
	return ""
}
